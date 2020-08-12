package rtm

import (
	"context"
	"encoding/json"
	"encoding/xml"
)

type TasksService struct {
	client *Client
}

type Note struct {
	ID       string
	Created  Time
	Modified Time
	Text     string
}

type TaskSeries struct {
	ID         string
	Created    Time
	Modified   Time
	Name       string
	Source     string
	URL        string
	LocationID string
	Tags       []string
	Notes      []Note
	Task       []Task
}

type Priority string

type Task struct {
	ID           string
	Due          Time
	HasDueTime   bool
	Added        Time
	Completed    Time
	Deleted      string
	Priority     Priority
	Estimate     Duration
	Start        Time
	HasStartTime bool
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

	b, err := t.client.CallJSON(ctx, "rtm.tasks.getList", args)
	if err != nil {
		return nil, err
	}

	return t.getListUnmarshal(b)
}

func (t *TasksService) getListUnmarshal(b []byte) (map[string][]TaskSeries, error) {
	var resp struct {
		Rsp struct {
			Tasks struct {
				Rev  string `json:"rev"`
				List []struct {
					ID         string `json:"id"`
					TaskSeries []struct {
						ID         string          `json:"id"`
						Created    Time            `json:"created"`
						Modified   Time            `json:"modified"`
						Name       string          `json:"name"`
						Source     string          `json:"source"`
						URL        string          `json:"url"`
						LocationID string          `json:"location_id"`
						Tags       json.RawMessage `json:"tags"`
						Notes      json.RawMessage `json:"notes"`
						Task       []struct {
							ID  string `json:"id"`
							Due Time   `json:"due"`
							// HasDueTime bool `json:"has_due_time"`
							Added Time `json:"added"`
							// Completed Time `json:"completed"`
							// Deleted      string    `json:"deleted"`
							Priority Priority `json:"priority"`
							Estimate Duration `json:"estimate"`
							Start    Time     `json:"start"`
							// HasStartTime bool `json:"has_start_time"`
						} `json:"task"`
					} `json:"taskseries"`
				} `json:"list"`
			} `json:"tasks"`
		} `json:"rsp"`
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	res := make(map[string][]TaskSeries, len(resp.Rsp.Tasks.List))
	for _, l := range resp.Rsp.Tasks.List {
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

		taskSeries := make([]TaskSeries, len(l.TaskSeries))
		for i, ts := range l.TaskSeries {
			tasks := make([]Task, len(ts.Task))
			for j, t := range ts.Task {
				tasks[j] = Task{
					ID:    t.ID,
					Due:   t.Due,
					Added: t.Added,
					// Completed: t.Completed,
					Priority: t.Priority,
					Estimate: t.Estimate,
					Start:    t.Start,
				}
			}

			// in RTM JSON API, tags are returned as either empty JSON array (when there are no tags),
			// or as an object with a single field "tag" containing array
			var tags []string
			if string(ts.Tags) != "[]" {
				var tagsResp struct {
					Tags []string `json:"tag"`
				}
				if err := json.Unmarshal(ts.Tags, &tagsResp); err != nil {
					return nil, err
				}

				tags = tagsResp.Tags
			}

			// in RTM JSON API, notes are returned as either empty JSON array (when there are no notes),
			// or as an object with a single field "note" containing array
			var notes []Note
			if string(ts.Notes) != "[]" {
				var notesResp struct {
					Note []struct {
						ID       string `json:"id"`
						Created  Time   `json:"created"`
						Modified Time   `json:"modified"`
						Text     string `json:"$t"`
					} `json:"note"`
				}
				if err := json.Unmarshal(ts.Notes, &notesResp); err != nil {
					return nil, err
				}

				notes = make([]Note, len(notesResp.Note))
				for j, n := range notesResp.Note {
					notes[j] = Note{
						ID:       n.ID,
						Created:  n.Created,
						Modified: n.Modified,
						Text:     n.Text,
					}
				}
			}

			taskSeries[i] = TaskSeries{
				ID:         ts.ID,
				Created:    ts.Created,
				Modified:   ts.Modified,
				Name:       ts.Name,
				Source:     ts.Source,
				URL:        ts.URL,
				LocationID: ts.LocationID,
				Tags:       tags,
				Notes:      notes,
				Task:       tasks,
			}
		}
		res[l.ID] = taskSeries
	}

	return res, nil
}

type TasksAddParams struct {
	ListID       string
	Name         string
	Parse        bool
	ParentTaskID string
}

type tasksAddResponseList struct {
	XMLName    xml.Name   `json:"list"`
	ListID     string     `json:"id"`
	TaskSeries TaskSeries `json:"taskseries"`
}

type tasksAddResponse struct {
	Transaction          *Transaction         `json:"transaction"`
	TasksAddResponseList tasksAddResponseList `json:"list"`
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

	b, err := t.client.CallJSON(ctx, "rtm.tasks.add", args)
	if err != nil {
		return nil, err
	}

	var resp tasksAddResponse
	if err = xml.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return &resp.TasksAddResponseList.TaskSeries, nil
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

	_, err := t.client.CallJSON(ctx, "rtm.tasks.delete.rtm", args)
	return err
}
