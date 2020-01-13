package rtm

import (
	"context"
	"encoding/xml"
)

type ListsService struct {
	client *Client
}

type List struct {
	ID       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Locked   bool   `xml:"locked,attr"`
	Archived bool   `xml:"archived,attr"`
	Position int    `xml:"position,attr"`
	Smart    bool   `xml:"smart,attr"`
}

// https://www.rememberthemilk.com/services/api/methods/rtm.lists.getList.rtm
func (l *ListsService) GetList(ctx context.Context) ([]List, error) {
	b, err := l.client.Call(ctx, "rtm.lists.getList", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		XMLName xml.Name `xml:"lists"`
		Lists   []List   `xml:"list"`
	}
	if err = xml.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return resp.Lists, nil
}
