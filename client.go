// Package rtm provides access to Remember The Milk API v2.
package rtm // import "github.com/AlekSi/rtm"

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"sort"
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
	Debugf     func(format string, args ...any)

	recordTestdata bool // records testdata/tmp-XXX files
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
//
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
//
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

// post calls the given method with arguments, returning body or error.
func (c *Client) post(ctx context.Context, method string, args Args) ([]byte, error) {
	q := make(url.Values)
	for k, v := range args {
		q.Set(k, v)
	}
	q.Set("v", "2")
	q.Set("method", method)
	q.Set("format", "json")
	q.Set("api_key", c.APIKey)
	if c.AuthToken != "" {
		q.Set("auth_token", c.AuthToken)
	}
	c.sign(q)

	u := restEndpoint
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	if c.Debugf != nil {
		b, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}
		c.Debugf("Request:\n%s", b)
	}

	resp, err := c.http().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c.Debugf != nil {
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		c.Debugf("Response:\n%s", b)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP status code %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.recordTestdata {
		s := string(b) // do not mutate b that we return later
		for _, p := range []string{c.APIKey, c.APISecret, c.AuthToken} {
			if p != "" {
				s = strings.ReplaceAll(s, p, "XXX")
			}
		}

		filename := filepath.Join("testdata", "tmp-"+method+".json")
		if err = ioutil.WriteFile(filename, []byte(s), 0o666); err != nil {
			return nil, err
		}
	}

	return b, nil
}

// checkErrorResponse checks response and returns error if it can't be parsed or contains error.
func checkErrorResponse(b []byte) error {
	var resp struct {
		Rsp struct {
			Stat string `json:"stat"`
			Err  *struct {
				Code int    `json:"code,string"`
				Msg  string `json:"msg"`
			} `json:"err"`
		} `json:"rsp"`
	}
	err := json.Unmarshal(b, &resp)
	switch {
	case err != nil:
		return err
	case resp.Rsp.Err != nil:
		return &Error{
			Code: resp.Rsp.Err.Code,
			Msg:  resp.Rsp.Err.Msg,
		}
	case resp.Rsp.Stat != "ok":
		return fmt.Errorf("unexpected stat %q", resp.Rsp.Stat)
	default:
		return nil
	}
}

// Call calls the given method with arguments and returns response body or error.
func (c *Client) Call(ctx context.Context, method string, args Args) ([]byte, error) {
	b, err := c.post(ctx, method, args)
	if err != nil {
		return nil, err
	}

	if err = checkErrorResponse(b); err != nil {
		return nil, err
	}

	return b, nil
}
