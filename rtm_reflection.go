package rtm

import (
	"context"
	"encoding/json"
	"sort"
	"strconv"
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
	b, err := r.client.CallJSON(ctx, "rtm.reflection.getMethodInfo", Args{"method_name": method})
	if err != nil {
		return nil, err
	}

	return r.getMethodInfoUnmarshal(b)
}

func (r *ReflectionService) getMethodInfoUnmarshal(b []byte) (*MethodInfo, error) {
	var res struct {
		Rsp struct {
			Method struct {
				Name       string `json:"name"`
				NeedsLogin string `json:"needslogin"`
			} `json:"method"`
		} `json:"rsp"`
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	needsLogin, _ := strconv.ParseBool(res.Rsp.Method.NeedsLogin)

	return &MethodInfo{
		Name:       res.Rsp.Method.Name,
		NeedsLogin: needsLogin,
	}, nil
}

// https://www.rememberthemilk.com/services/api/methods/rtm.reflection.getMethods.rtm
func (r *ReflectionService) GetMethods(ctx context.Context) ([]string, error) {
	b, err := r.client.CallJSON(ctx, "rtm.reflection.getMethods", nil)
	if err != nil {
		return nil, err
	}

	return r.getMethodsUnmarshal(b)
}

func (r *ReflectionService) getMethodsUnmarshal(b []byte) ([]string, error) {
	var res struct {
		Rsp struct {
			Methods struct {
				Method []string `json:"method"`
			} `json:"methods"`
		} `json:"rsp"`
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	sort.Strings(res.Rsp.Methods.Method)

	return res.Rsp.Methods.Method, nil
}
