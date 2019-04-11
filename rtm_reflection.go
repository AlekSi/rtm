package rtm

import (
	"context"
	"encoding/xml"
)

type ReflectionService struct {
	client *Client
}

type MethodInfo struct {
	Name       string
	NeedsLogin bool
}

// https://www.rememberthemilk.com/services/api/methods/rtm.reflection.getMethodInfo.rtm
func (r *ReflectionService) GetMethodInfo(ctx context.Context, method string) (*MethodInfo, error) {
	b, err := r.client.Call(ctx, "rtm.reflection.getMethodInfo", Args{"method_name": method})
	if err != nil {
		return nil, err
	}

	var resp struct {
		XMLName    xml.Name `xml:"method"`
		Name       string   `xml:"name,attr"`
		NeedsLogin bool     `xml:"needslogin,attr"`
	}
	if err = xml.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return &MethodInfo{
		Name:       resp.Name,
		NeedsLogin: resp.NeedsLogin,
	}, nil
}

// https://www.rememberthemilk.com/services/api/methods/rtm.reflection.getMethods.rtm
func (r *ReflectionService) GetMethods(ctx context.Context) ([]string, error) {
	b, err := r.client.Call(ctx, "rtm.reflection.getMethods", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		XMLName xml.Name `xml:"methods"`
		Methods []string `xml:"method"`
	}
	if err = xml.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return resp.Methods, nil
}
