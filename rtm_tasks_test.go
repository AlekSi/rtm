package rtm

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTasks(t *testing.T) {
	t.Run("GetList", func(t *testing.T) {
		expected := tasksGetListResponse{
			XMLName: xml.Name{Local: "tasks"},
			Lists: []tasksGetListResponseList{{
				XMLName: xml.Name{Local: "list"},
				ID:      "43911488",
				TaskSeries: []TaskSeries{{
					ID:         "358441579",
					Created:    parseTime(t, "2018-08-05T09:30:53Z"),
					Modified:   parseTime(t, "2020-08-01T13:05:09Z"),
					Name:       "Task 1",
					Source:     "js",
					URL:        "https://github.com/AlekSi/rtm",
					LocationID: "1265394",
					Tags:       []string{"tag1", "tag2"},
					Notes: []Note{{
						ID:       "81657666",
						Created:  parseTime(t, "2020-08-01T13:05:09Z"),
						Modified: parseTime(t, "2020-08-01T13:05:09Z"),
						Text:     "Note 2",
					}, {
						ID:       "81657665",
						Created:  parseTime(t, "2020-08-01T13:05:05Z"),
						Modified: parseTime(t, "2020-08-01T13:05:05Z"),
						Text:     "Note 1",
					}},
					Task: []Task{{
						ID:           "622345829",
						Due:          parseTime(t, "2018-08-06T20:30:00Z"),
						HasDueTime:   true,
						Added:        parseTime(t, "2018-08-05T09:30:53Z"),
						Priority:     "N",
						Start:        parseTime(t, "2018-08-04T21:30:00Z"),
						HasStartTime: true,
					}},
				}, {
					ID:       "358441583",
					Created:  parseTime(t, "2018-08-05T09:30:56Z"),
					Modified: parseTime(t, "2018-08-05T13:42:43Z"),
					Name:     "Задача 2",
					Source:   "js",
					Task: []Task{{
						ID:        "622345833",
						Due:       parseTime(t, "2018-08-05T21:00:00Z"),
						Added:     parseTime(t, "2018-08-05T09:30:56Z"),
						Completed: parseTime(t, "2018-08-05T13:42:43Z"),
						Priority:  "2",
						Estimate:  parseDuration(t, "PT1H12M"),
						Start:     parseTime(t, "2018-08-04T21:00:00Z"),
					}},
				}},
			}},
		}

		t.Run("Decode", func(t *testing.T) {
			var actual tasksGetListResponse
			unmarshalTestdataFile(t, "rtm.tasks.getList.xml", &actual)
			assert.Equal(t, expected, actual)
		})

		t.Run("Real", func(t *testing.T) {
			params := &TasksGetListParams{
				ListID: "43911488",
			}
			actual, err := GetClient(t).Tasks().GetList(Ctx, params)
			require.NoError(t, err)
			assert.Equal(t, expected.toMap(), actual)

			// for _, series := range expected {
			// 	for _, s := range series {
			// 		for _, task := range s.Task {
			// 			assert.Equal(t, task.HasDueTime, task.Due.HasTime(), "due time for task %+v", task)
			// 			assert.Equal(t, task.HasStartTime, task.Start.HasTime(), "start time for task %+v", task)
			// 		}
			// 	}
			// }
		})
	})

	t.Run("AddDecode", func(t *testing.T) {
		var actual tasksAddResponse
		unmarshalTestdataFile(t, "rtm.tasks.add.xml", &actual)
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
					Created:  parseTime(t, "2020-01-14T06:23:58Z"),
					Modified: parseTime(t, "2020-01-14T06:23:58Z"),
					Name:     "TestTasks/AddDelete",
					Source:   "api:XXX",
					Task: []Task{{
						ID:       "706419831",
						Added:    parseTime(t, "2020-01-14T06:23:58Z"),
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
