package rtm

import (
	"context"
	"encoding/json"
	"sort"
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

	return r.getMethodInfoUnmarshal(b)
}

func (r *ReflectionService) getMethodInfoUnmarshal(b []byte) (*MethodInfo, error) {
	var resp struct {
		Rsp struct {
			Method struct {
				Name       string  `json:"name"`
				NeedsLogin rtmBool `json:"needslogin"`
			} `json:"method"`
		} `json:"rsp"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return &MethodInfo{
		Name:       resp.Rsp.Method.Name,
		NeedsLogin: bool(resp.Rsp.Method.NeedsLogin),
	}, nil
}

// https://www.rememberthemilk.com/services/api/methods/rtm.reflection.getMethods.rtm
func (r *ReflectionService) GetMethods(ctx context.Context) ([]string, error) {
	b, err := r.client.Call(ctx, "rtm.reflection.getMethods", nil)
	if err != nil {
		return nil, err
	}

	return r.getMethodsUnmarshal(b)
}

func (r *ReflectionService) getMethodsUnmarshal(b []byte) ([]string, error) {
	var resp struct {
		Rsp struct {
			Methods struct {
				Method []string `json:"method"`
			} `json:"methods"`
		} `json:"rsp"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	sort.Strings(resp.Rsp.Methods.Method)

	return resp.Rsp.Methods.Method, nil
}
