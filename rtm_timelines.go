package rtm

import (
	"context"
	"encoding/json"
)

type TimelinesService struct {
	client *Client
}

// https://www.rememberthemilk.com/services/api/methods/rtm.timelines.create.rtm
func (t *TimelinesService) Create(ctx context.Context) (string, error) {
	b, err := t.client.CallJSON(ctx, "rtm.timelines.create", nil)
	if err != nil {
		return "", err
	}

	return t.createUnmarshal(b)
}

func (t *TimelinesService) createUnmarshal(b []byte) (string, error) {
	var resp struct {
		Rsp struct {
			Timeline string `json:"timeline"`
		}
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	return resp.Rsp.Timeline, nil
}
