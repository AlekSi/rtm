package rtm

import (
	"context"
	"encoding/json"
)

type ListsService struct {
	client *Client
}

type List struct {
	ID       string
	Name     string
	Position int
	Locked   bool
	Archived bool
	Deleted  bool
	Smart    bool
}

// https://www.rememberthemilk.com/services/api/methods/rtm.lists.getList.rtm
func (l *ListsService) GetList(ctx context.Context) ([]List, error) {
	b, err := l.client.Call(ctx, "rtm.lists.getList", nil)
	if err != nil {
		return nil, err
	}

	return l.getListUnmarshal(b)
}

func (l *ListsService) getListUnmarshal(b []byte) ([]List, error) {
	var resp struct {
		Rsp struct {
			Lists struct {
				List []struct {
					ID       string  `json:"id"`
					Name     string  `json:"name"`
					Deleted  rtmBool `json:"deleted"`
					Locked   rtmBool `json:"locked"`
					Archived rtmBool `json:"archived"`
					Position int     `json:"position,string"`
					Smart    rtmBool `json:"smart"`
				} `json:"list"`
			} `json:"lists"`
		} `json:"rsp"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	res := make([]List, len(resp.Rsp.Lists.List))
	for i, l := range resp.Rsp.Lists.List {
		res[i] = List{
			ID:       l.ID,
			Name:     l.Name,
			Position: l.Position,
			Locked:   bool(l.Locked),
			Archived: bool(l.Archived),
			Deleted:  bool(l.Deleted),
			Smart:    bool(l.Smart),
		}
	}

	return res, nil
}
