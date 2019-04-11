package rtm

import (
	"context"
)

type TestService struct {
	client *Client
}

// https://www.rememberthemilk.com/services/api/methods/rtm.test.echo.rtm
func (t *TestService) Echo(ctx context.Context, args Args) ([]byte, error) {
	return t.client.Call(ctx, "rtm.test.echo", args)
}

// https://www.rememberthemilk.com/services/api/methods/rtm.test.login.rtm
func (t *TestService) Login(ctx context.Context) error {
	_, err := t.client.Call(ctx, "rtm.test.login", nil)
	return err
}
