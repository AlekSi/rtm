package rtm

import (
	"context"
	"encoding/json"
	"encoding/xml"
)

type AuthService struct {
	client *Client
}

type AuthInfoUser struct {
	ID       string
	UserName string
	FullName string
}

type AuthInfo struct {
	Token string
	Perms Perms
	User  AuthInfoUser
}

// https://www.rememberthemilk.com/services/api/methods/rtm.auth.checkToken.rtm
func (a *AuthService) CheckToken(ctx context.Context) (*AuthInfo, error) {
	b, err := a.client.CallJSON(ctx, "rtm.auth.checkToken", nil)
	if err != nil {
		return nil, err
	}

	return a.checkTokenUnmarshal(b)
}

func (a *AuthService) checkTokenUnmarshal(b []byte) (*AuthInfo, error) {
	var resp struct {
		Rsp struct {
			Auth struct {
				Token string `json:"token"`
				Perms string `json:"perms"`
				User  struct {
					ID       string `json:"id"`
					UserName string `json:"username"`
					FullName string `json:"fullname"`
				} `json:"user"`
			} `json:"auth"`
		} `json:"rsp"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return &AuthInfo{
		Token: resp.Rsp.Auth.Token,
		Perms: Perms(resp.Rsp.Auth.Perms),
		User: AuthInfoUser{
			ID:       resp.Rsp.Auth.User.ID,
			UserName: resp.Rsp.Auth.User.UserName,
			FullName: resp.Rsp.Auth.User.FullName,
		},
	}, nil
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
