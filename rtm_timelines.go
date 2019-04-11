package rtm

import (
	"context"
	"encoding/xml"
)

type TimelinesService struct {
	client *Client
}

// https://www.rememberthemilk.com/services/api/methods/rtm.timelines.create.rtm
func (t *TimelinesService) Create(ctx context.Context) (string, error) {
	b, err := t.client.Call(ctx, "rtm.timelines.create", nil)
	if err != nil {
		return "", err
	}

	var resp struct {
		XMLName xml.Name `xml:"timeline"`
		ID      string   `xml:",chardata"`
	}
	if err = xml.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	return resp.ID, nil
}
