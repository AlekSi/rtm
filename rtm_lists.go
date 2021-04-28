package rtm

import (
	"context"
	"encoding/json"
	"strconv"
)

type ListsService struct {
	client *Client
}

type List struct {
	ID       string
	Name     string
	Locked   bool
	Archived bool
	Position int
	Smart    bool
}

// https://www.rememberthemilk.com/services/api/methods/rtm.lists.getList.rtm
func (l *ListsService) GetList(ctx context.Context) ([]List, error) {
	b, err := l.client.CallJSON(ctx, "rtm.lists.getList", nil)
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
					ID       string `json:"id"`
					Name     string `json:"name"`
					Locked   string `json:"locked"`
					Archived string `json:"archived"`
					Position string `json:"position"`
					Smart    string `json:"smart"`
				} `json:"list"`
			} `json:"lists"`
		} `json:"rsp"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	res := make([]List, len(resp.Rsp.Lists.List))
	for i, l := range resp.Rsp.Lists.List {
		locked, _ := strconv.ParseBool(l.Locked)
		archived, _ := strconv.ParseBool(l.Archived)
		position, _ := strconv.Atoi(l.Position)
		smart, _ := strconv.ParseBool(l.Smart)
		res[i] = List{
			ID:       l.ID,
			Name:     l.Name,
			Locked:   locked,
			Archived: archived,
			Position: position,
			Smart:    smart,
		}
	}
	return res, nil
}
