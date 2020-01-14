package rtm

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTasks(t *testing.T) {
	t.Run("GetList", func(t *testing.T) {
		params := &TasksGetListParams{
			ListID: "43911488",
		}
		lists, err := GetClient(t).Tasks().GetList(Ctx, params)
		require.NoError(t, err)
		expected := map[string][]TaskSeries{
			"43911488": {
				{
					ID:         "358441579",
					Created:    Time{time.Date(2018, 8, 5, 9, 30, 53, 0, time.UTC)},
					Modified:   Time{time.Date(2019, 4, 11, 4, 28, 50, 0, time.UTC)},
					Name:       "Task 1",
					Source:     "js",
					URL:        "https://github.com/AlekSi/rtm",
					LocationID: "1265394",
					Task: []Task{
						{
							ID:           "622345829",
							Due:          Time{time.Date(2018, 8, 6, 20, 30, 0, 0, time.UTC)},
							HasDueTime:   true,
							Added:        Time{time.Date(2018, 8, 5, 9, 30, 53, 0, time.UTC)},
							Completed:    Time{},
							Priority:     "N",
							Estimate:     Duration{},
							Start:        Time{time.Date(2018, 8, 4, 21, 30, 0, 0, time.UTC)},
							HasStartTime: true,
						},
					},
				}, {
					ID:       "358441583",
					Created:  Time{time.Date(2018, 8, 5, 9, 30, 56, 0, time.UTC)},
					Modified: Time{time.Date(2018, 8, 5, 13, 42, 43, 0, time.UTC)},
					Name:     "Задача 2",
					Source:   "js",
					Task: []Task{
						{
							ID:           "622345833",
							Due:          Time{time.Date(2018, 8, 5, 21, 0, 0, 0, time.UTC)},
							HasDueTime:   false,
							Added:        Time{time.Date(2018, 8, 5, 9, 30, 56, 0, time.UTC)},
							Completed:    Time{time.Date(2018, 8, 5, 13, 42, 43, 0, time.UTC)},
							Priority:     "2",
							Estimate:     Duration{time.Hour + 12*time.Minute},
							Start:        Time{time.Date(2018, 8, 4, 21, 0, 0, 0, time.UTC)},
							HasStartTime: false,
						},
					},
				},
			},
		}

		assert.Equal(t, expected, lists)

		// for _, series := range expected {
		// 	for _, s := range series {
		// 		for _, task := range s.Task {
		// 			assert.Equal(t, task.HasDueTime, task.Due.HasTime(), "due time for task %+v", task)
		// 			assert.Equal(t, task.HasStartTime, task.Start.HasTime(), "start time for task %+v", task)
		// 		}
		// 	}
		// }
	})

	t.Run("AddDecode", func(t *testing.T) {
		var actual tasksAddResponse
		unmarshalFile(t, "rtm.tasks.add.xml", &actual)
		expected := tasksAddResponse{
			Transaction: &Transaction{
				XMLName: xml.Name{Local: "transaction"},
				ID:      "4613647463",
			},
			TasksAddResponseList: tasksAddResponseList{
				XMLName: xml.Name{Local: "list"},
				ListID:  "43911488",
				TaskSeries: TaskSeries{
					ID:       "403924673",
					Created:  parseTime("2020-01-14T06:23:58Z"),
					Modified: parseTime("2020-01-14T06:23:58Z"),
					Name:     "TestTasks/AddDelete",
					Source:   "api:XXX",
					Task: []Task{{
						ID:       "706419831",
						Added:    parseTime("2020-01-14T06:23:58Z"),
						Priority: "N",
					}},
				},
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("AddDelete", func(t *testing.T) {
		timeline, err := GetClient(t).Timelines().Create(Ctx)
		require.NoError(t, err)

		task, err := GetClient(t).Tasks().Add(Ctx, timeline, TasksAddParams{
			ListID: "43911488",
			Name:   t.Name(),
		})
		require.NoError(t, err)
		t.Log(task)
		require.NotEmpty(t, task.Task)

		err = GetClient(t).Tasks().Delete(Ctx, timeline, TasksDeleteParams{
			ListID:       "43911488",
			TaskSeriesID: task.ID,
			TaskID:       task.Task[0].ID,
		})
		require.NoError(t, err)
	})
}
