// Package rtm provides access to Remember The Milk API v2.
package rtm // import "github.com/AlekSi/rtm"

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	userAgent = "github.com/AlekSi/rtm"
)

var (
	authEndpoint = url.URL{
		Scheme: "https",
		Host:   "api.rememberthemilk.com",
		Path:   "/services/auth/",
	}
	restEndpoint = url.URL{
		Scheme: "https",
		Host:   "api.rememberthemilk.com",
		Path:   "/services/rest/",
	}
)

type Perms string

const (
	Read   Perms = "read"
	Write  Perms = "write"
	Delete Perms = "delete"
)

type Args map[string]string

type Client struct {
	APIKey     string
	APISecret  string
	AuthToken  string
	HTTPClient *http.Client
	Debugf     func(format string, args ...interface{})

	recordTestdata bool
}

func (c *Client) Auth() *AuthService             { return &AuthService{c} }
func (c *Client) Lists() *ListsService           { return &ListsService{c} }
func (c *Client) Reflection() *ReflectionService { return &ReflectionService{c} }
func (c *Client) Tasks() *TasksService           { return &TasksService{c} }
func (c *Client) Test() *TestService             { return &TestService{c} }
func (c *Client) Timelines() *TimelinesService   { return &TimelinesService{c} }

// http returns used HTTP client.
func (c *Client) http() *http.Client {
	if c.HTTPClient == nil {
		return http.DefaultClient
	}
	return c.HTTPClient
}

// sign adds api_sig to request parameters.
// See https://www.rememberthemilk.com/services/api/authentication.rtm, "Signing Requests".
func (c *Client) sign(q url.Values) {
	keys := make([]string, 0, len(q))
	for k := range q {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sts := c.APISecret
	for _, k := range keys {
		sts += k + q.Get(k)
	}

	sum := md5.Sum([]byte(sts))
	q.Set("api_sig", hex.EncodeToString(sum[:]))
}

// AuthenticationURL returns authentication URL for given permissions and frob (that can be empty).
// See https://www.rememberthemilk.com/services/api/authentication.rtm, "User authentication for web-based applications"
// and "User authentication for desktop applications".
func (c *Client) AuthenticationURL(perms Perms, frob string) string {
	q := make(url.Values)
	q.Set("api_key", c.APIKey)
	q.Set("perms", string(perms))
	if frob != "" {
		q.Set("frob", frob)
	}
	c.sign(q)

	u := authEndpoint
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *Client) post(ctx context.Context, method string, args Args, format string) ([]byte, error) {
	q := make(url.Values)
	for k, v := range args {
		q.Set(k, v)
	}
	q.Set("v", "2")
	q.Set("method", method)
	if format != "" {
		q.Set("format", format)
	}
	q.Set("api_key", c.APIKey)
	if c.AuthToken != "" {
		q.Set("auth_token", c.AuthToken)
	}
	c.sign(q)

	u := restEndpoint
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", userAgent)

	if c.Debugf != nil {
		b, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}
		c.Debugf("Request:\n%s", b)
	}

	resp, err := c.http().Do(req)
	if resp != nil {
		defer resp.Body.Close()

		if c.Debugf != nil {
			b, err := httputil.DumpResponse(resp, true)
			if err != nil {
				return nil, err
			}
			c.Debugf("Response:\n%s", b)
		}
	}
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP status code %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.recordTestdata {
		ext := format
		if format == "" || format == "rest" {
			ext = "xml"
		}
		filename := filepath.Join("testdata", "tmp-"+method+"."+ext)
		s := string(b)
		for _, p := range []string{c.APIKey, c.APISecret, c.AuthToken} {
			if p != "" {
				s = strings.Replace(s, p, "XXX", -1)
			}
		}
		if err = ioutil.WriteFile(filename, []byte(s), 0666); err != nil {
			return nil, err
		}
	}

	return b, nil
}

func unmarshalXMLRsp(b []byte) ([]byte, error) {
	var rsp struct {
		XMLName xml.Name `xml:"rsp"`
		Stat    string   `xml:"stat,attr"`
		Err     *Error   `xml:"err"`
		Inner   []byte   `xml:",innerxml"`
	}
	err := xml.Unmarshal(b, &rsp)
	switch {
	case err != nil:
		return nil, err
	case rsp.Err != nil:
		return nil, rsp.Err
	case rsp.Stat != "ok":
		return nil, fmt.Errorf("unexpected stat %q", rsp.Stat)
	default:
		return rsp.Inner, nil
	}
}

// TODO remove
func (c *Client) Call(ctx context.Context, method string, args Args) ([]byte, error) {
	b, err := c.post(ctx, method, args, "")
	if err != nil {
		return nil, err
	}

	return unmarshalXMLRsp(b)
}

// TODO rename to Call
func (c *Client) CallJSON(ctx context.Context, method string, args Args) ([]byte, error) {
	b, err := c.post(ctx, method, args, "json")
	if err != nil {
		return nil, err
	}

	return c.callJSONUnmarshal(b)
}

// TODO rename to callUnmarshal
func (c *Client) callJSONUnmarshal(b []byte) ([]byte, error) {
	var resp struct {
		Rsp struct {
			Stat string `json:"stat"`
			Err  *struct {
				Code string `json:"code"`
				Msg  string `json:"msg"`
			} `json:"err"`
		} `json:"rsp"`
	}
	err := json.Unmarshal(b, &resp)
	switch {
	case err != nil:
		return nil, err
	case resp.Rsp.Err != nil:
		code, _ := strconv.Atoi(resp.Rsp.Err.Code)
		return nil, &Error{
			Code: code,
			Msg:  resp.Rsp.Err.Msg,
		}
	case resp.Rsp.Stat != "ok":
		return nil, fmt.Errorf("unexpected stat %q", resp.Rsp.Stat)
	default:
		return b, nil
	}
}

// check interfaces
var (
	_ error = (*Error)(nil)
)
