package rtm

import (
	"context"
	"encoding/json"
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
	b, err := a.client.Call(ctx, "rtm.auth.checkToken", nil)
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

	return a.getFrobUnmarshal(b)
}

func (a *AuthService) getFrobUnmarshal(b []byte) (string, error) {
	var resp struct {
		Rsp struct {
			Frob string `json:"frob"`
		} `json:"rsp"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	return resp.Rsp.Frob, nil
}

// https://www.rememberthemilk.com/services/api/methods/rtm.auth.getToken.rtm
func (a *AuthService) GetToken(ctx context.Context, frob string) (*AuthInfo, error) {
	b, err := a.client.Call(ctx, "rtm.auth.getToken", Args{"frob": frob})
	if err != nil {
		return nil, err
	}

	return a.getTokenUnmarshal(b)
}

func (a *AuthService) getTokenUnmarshal(b []byte) (*AuthInfo, error) {
	// responce is exactly the same
	return a.checkTokenUnmarshal(b)
}
