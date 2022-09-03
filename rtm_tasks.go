package rtm

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type TasksService struct {
	client *Client
}

type Note struct {
	ID       string
	Created  DateTime
	Modified DateTime
	Text     string
}

type TaskSeries struct {
	ID         string
	Created    DateTime
	Modified   DateTime
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
	ID        string
	Due       DateTime
	Added     DateTime
	Completed DateTime
	Deleted   DateTime
	Priority  Priority
	Estimate  time.Duration
	Start     DateTime
}

type taskResp struct {
	ID           string      `json:"id"`
	Due          DateTime    `json:"due"`
	HasDueTime   rtmBool     `json:"has_due_time"`
	Added        DateTime    `json:"added"`
	Completed    DateTime    `json:"completed"`
	Deleted      DateTime    `json:"deleted"`
	Priority     Priority    `json:"priority"`
	Estimate     rtmDuration `json:"estimate"`
	Start        DateTime    `json:"start"`
	HasStartTime rtmBool     `json:"has_start_time"`
}

func (t *taskResp) parseTask() *Task {
	if !t.HasDueTime {
		t.Due.stripTime()
	}
	if !t.HasStartTime {
		t.Start.stripTime()
	}

	return &Task{
		ID:        t.ID,
		Due:       t.Due,
		Added:     t.Added,
		Completed: t.Completed,
		Deleted:   t.Deleted,
		Priority:  t.Priority,
		Estimate:  t.Estimate.Duration,
		Start:     t.Start,
	}
}

type taskSeriesResp struct {
	ID         string          `json:"id"`
	Created    DateTime        `json:"created"`
	Modified   DateTime        `json:"modified"`
	Name       string          `json:"name"`
	Source     string          `json:"source"`
	URL        string          `json:"url"`
	LocationID string          `json:"location_id"`
	Tags       json.RawMessage `json:"tags"`
	Notes      json.RawMessage `json:"notes"`
	Task       []taskResp      `json:"task"`
}

func (ts *taskSeriesResp) parseTags() ([]string, error) {
	// in RTM JSON API, tags are returned as either empty JSON array (when there are no tags),
	// or as an object with a single field "tag" containing array
	if string(ts.Tags) == "[]" {
		return nil, nil
	}

	var tagsResp struct {
		Tags []string `json:"tag"`
	}
	if err := json.Unmarshal(ts.Tags, &tagsResp); err != nil {
		return nil, err
	}

	return tagsResp.Tags, nil
}

func (ts *taskSeriesResp) parseNotes() ([]Note, error) {
	// in RTM JSON API, notes are returned as either empty JSON array (when there are no notes),
	// or as an object with a single field "note" containing array
	if string(ts.Notes) == "[]" {
		return nil, nil
	}

	var notesResp struct {
		Note []struct {
			ID       string   `json:"id"`
			Created  DateTime `json:"created"`
			Modified DateTime `json:"modified"`
			Text     string   `json:"$t"`
		} `json:"note"`
	}
	if err := json.Unmarshal(ts.Notes, &notesResp); err != nil {
		return nil, err
	}

	res := make([]Note, len(notesResp.Note))
	for j, n := range notesResp.Note {
		res[j] = Note{
			ID:       n.ID,
			Created:  n.Created,
			Modified: n.Modified,
			Text:     n.Text,
		}
	}

	return res, nil
}

func (ts *taskSeriesResp) parseTaskSeries() (*TaskSeries, error) {
	tasks := make([]Task, len(ts.Task))
	for j, t := range ts.Task {
		tasks[j] = *t.parseTask()
	}

	tags, err := ts.parseTags()
	if err != nil {
		return nil, err
	}

	notes, err := ts.parseNotes()
	if err != nil {
		return nil, err
	}

	return &TaskSeries{
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
	}, nil
}

type TasksGetListParams struct {
	ListID   string
	Filter   string
	LastSync DateTime
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

	return t.getListUnmarshal(b)
}

func (t *TasksService) getListUnmarshal(b []byte) (map[string][]TaskSeries, error) {
	var resp struct {
		Rsp struct {
			Tasks struct {
				Rev  string `json:"rev"`
				List []struct {
					ID         string           `json:"id"`
					TaskSeries []taskSeriesResp `json:"taskseries"`
				} `json:"list"`
			} `json:"tasks"`
		} `json:"rsp"`
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	res := make(map[string][]TaskSeries, len(resp.Rsp.Tasks.List))
	for _, l := range resp.Rsp.Tasks.List {
		taskSeries := make([]TaskSeries, len(l.TaskSeries))
		for i, ts := range l.TaskSeries {
			p, err := ts.parseTaskSeries()
			if err != nil {
				return nil, err
			}

			taskSeries[i] = *p
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

// https://www.rememberthemilk.com/services/api/methods/rtm.tasks.add.rtm
func (t *TasksService) Add(ctx context.Context, timeline string, params *TasksAddParams) (*TaskSeries, error) {
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

	return t.addUnmarshal(b)
}

func (t *TasksService) addUnmarshal(b []byte) (*TaskSeries, error) {
	var resp struct {
		Rsp struct {
			Transaction struct {
				ID       string  `json:"id"`
				Undoable rtmBool `json:"undoable"`
			} `json:"transaction"`
			List struct {
				ID         string           `json:"id"`
				TaskSeries []taskSeriesResp `json:"taskseries"`
			} `json:"list"`
		}
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	ts := resp.Rsp.List.TaskSeries
	if l := len(ts); l != 1 {
		return nil, fmt.Errorf("expected 1 taskseries, got %d", l)
	}

	return ts[0].parseTaskSeries()
}

type TasksDeleteParams struct {
	ListID       string
	TaskSeriesID string
	TaskID       string
}

// https://www.rememberthemilk.com/services/api/methods/rtm.tasks.delete.rtm
func (t *TasksService) Delete(ctx context.Context, timeline string, params *TasksDeleteParams) error {
	args := Args{
		"timeline":      timeline,
		"list_id":       params.ListID,
		"taskseries_id": params.TaskSeriesID,
		"task_id":       params.TaskID,
	}

	_, err := t.client.Call(ctx, "rtm.tasks.delete.rtm", args)
	return err
}
