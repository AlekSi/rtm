package rtm

import (
	"context"
	"encoding/xml"
	"sort"
)

type TasksService struct {
	client *Client
}

type TaskSeries struct {
	ID         string `xml:"id,attr"`
	Created    Time   `xml:"created,attr"`
	Modified   Time   `xml:"modified,attr"`
	Name       string `xml:"name,attr"`
	Source     string `xml:"source,attr"`
	URL        string `xml:"url,attr"`
	LocationID string `xml:"location_id,attr"`
	// Tags         []interface{} `json:"tags"`
	// Notes        []interface{} `json:"notes"`
	Task []Task `xml:"task"`
}

type Priority string

type Task struct {
	ID         string `xml:"id,attr"`
	Due        Time   `xml:"due,attr"`
	HasDueTime bool   `xml:"has_due_time,attr"`
	Added      Time   `xml:"added,attr"`
	Completed  Time   `xml:"completed,attr,omitempty"`
	// Deleted      string    `xml:"deleted,attr"`
	Priority     Priority `xml:"priority,attr"`
	Estimate     Duration `xml:"estimate,attr"`
	Start        Time     `xml:"start,attr"`
	HasStartTime bool     `xml:"has_start_time,attr"`
}

type TasksGetListParams struct {
	ListID   string
	Filter   string
	LastSync Time
}

// https://www.rememberthemilk.com/services/api/methods/rtm.tasks.getList.rtm
func (t *TasksService) GetList(ctx context.Context, params *TasksGetListParams) (map[string][]TaskSeries, error) {
	args := make(Args)
	if params != nil {
		if params.ListID != "" {
			args["list_id"] = params.ListID
		}
		if params.Filter != "" {
			args["filter"] = params.Filter
		}
		if !params.LastSync.IsZero() {
			args["last_sync"] = params.LastSync.String()
		}
	}

	b, err := t.client.Call(ctx, "rtm.tasks.getList", args)
	if err != nil {
		return nil, err
	}

	var resp struct {
		XMLName xml.Name `xml:"tasks"`
		Lists   []struct {
			XMLName    xml.Name     `xml:"list"`
			ID         string       `xml:"id,attr"`
			TaskSeries []TaskSeries `xml:"taskseries"`
		} `xml:"list"`
	}
	if err = xml.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	res := make(map[string][]TaskSeries, len(resp.Lists))
	for _, list := range resp.Lists {
		sort.Slice(list.TaskSeries, func(i, j int) bool { return list.TaskSeries[i].ID < list.TaskSeries[j].ID })

		// for i, series := range list.TaskSeries {
		// 	for j, task := range series.Task {
		// 		if !task.HasDueTime {
		// 			list.TaskSeries[i].Task[j].Due = task.Due.withoutTime()
		// 		}
		// 		if !task.HasStartTime {
		// 			list.TaskSeries[i].Task[j].Start = task.Start.withoutTime()
		// 		}
		// 	}
		// }
		res[list.ID] = list.TaskSeries
	}
	return res, nil
}

type TasksAddParams struct {
	ListID       string
	Name         string
	Parse        bool
	ParentTaskID string
}

type tasksAddResponse struct {
	XMLName     xml.Name
	Transaction transaction `xml:"transaction"`
	// TransactionID       string     `xml:"transaction,attr"`
	// TransactionUndoable bool       `xml:"transaction,attr"`
	TaskSeries TaskSeries `xml:"taskseries"`
}

// https://www.rememberthemilk.com/services/api/methods/rtm.tasks.add.rtm
func (t *TasksService) Add(ctx context.Context, timeline string, params TasksAddParams) (*TaskSeries, error) {
	args := Args{
		"timeline": timeline,
		"name":     params.Name,
	}
	if params.ListID != "" {
		args["list_id"] = params.ListID
	}
	if params.Parse {
		args["parse"] = "1"
	}
	if params.ParentTaskID != "" {
		args["parent_task_id"] = params.ParentTaskID
	}

	b, err := t.client.Call(ctx, "rtm.tasks.add", args)
	if err != nil {
		return nil, err
	}

	var resp tasksAddResponse
	if err = xml.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	return &resp.TaskSeries, nil
}

type TasksDeleteParams struct {
	ListID       string
	TaskSeriesID string
	TaskID       string
}

// https://www.rememberthemilk.com/services/api/methods/rtm.tasks.delete.rtm
func (t *TasksService) Delete(ctx context.Context, timeline string, params TasksDeleteParams) error {
	args := Args{
		"timeline":      timeline,
		"list_id":       params.ListID,
		"taskseries_id": params.TaskSeriesID,
		"task_id":       params.TaskID,
	}

	_, err := t.client.Call(ctx, "rtm.tasks.delete.rtm", args)
	return err
}
