package rtm

import (
	"context"
	"encoding/xml"
)

type AuthService struct {
	client *Client
}

type AuthInfo struct {
	Token string
	Perms Perms
	User  struct {
		ID       string
		UserName string
		FullName string
	}
}

// https://www.rememberthemilk.com/services/api/methods/rtm.auth.checkToken.rtm
func (a *AuthService) CheckToken(ctx context.Context) (*AuthInfo, error) {
	b, err := a.client.Call(ctx, "rtm.auth.checkToken", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		XMLName xml.Name `xml:"auth"`
		Token   string   `xml:"token"`
		Perms   string   `xml:"perms"`
		User    struct {
			ID       string `xml:"id,attr"`
			UserName string `xml:"username,attr"`
			FullName string `xml:"fullname,attr"`
		} `xml:"user"`
	}
	if err = xml.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	res := &AuthInfo{
		Token: resp.Token,
		Perms: Perms(resp.Perms),
	}
	res.User.ID = resp.User.ID
	res.User.UserName = resp.User.UserName
	res.User.FullName = resp.User.FullName
	return res, nil
}

// https://www.rememberthemilk.com/services/api/methods/rtm.auth.getFrob.rtm
func (a *AuthService) GetFrob(ctx context.Context) (string, error) {
	b, err := a.client.Call(ctx, "rtm.auth.getFrob", nil)
	if err != nil {
		return "", err
	}

	var resp struct {
		XMLName xml.Name `xml:"frob"`
		Frob    string   `xml:",chardata"`
	}
	if err = xml.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	return resp.Frob, nil
}

// https://www.rememberthemilk.com/services/api/methods/rtm.auth.getToken.rtm
func (a *AuthService) GetToken(ctx context.Context, frob string) (string, error) {
	b, err := a.client.Call(ctx, "rtm.auth.getToken", Args{"frob": frob})
	if err != nil {
		return "", err
	}

	var resp struct {
		XMLName xml.Name `xml:"auth"`
		Token   string   `xml:"token"`
	}
	if err = xml.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	return resp.Token, nil
}
