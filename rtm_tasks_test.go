package rtm

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func sortByIDs(list map[string][]TaskSeries) {
	for k := range list {
		sort.Slice(list[k], func(i, j int) bool { return list[k][i].ID < list[k][j].ID })
	}
}

func TestTasks(t *testing.T) {
	t.Run("GetList", func(t *testing.T) {
		expected := map[string][]TaskSeries{
			"43911488": {
				{
					ID:         "358441579",
					Created:    parseDateTime(t, "2018-08-05T09:30:53Z"),
					Modified:   parseDateTime(t, "2020-08-01T13:05:09Z"),
					Name:       "Task 1",
					Source:     "js",
					URL:        "https://github.com/AlekSi/rtm",
					LocationID: "1265394",
					Tags:       []string{"tag1", "tag2"},
					Notes: []Note{{
						ID:       "81657666",
						Created:  parseDateTime(t, "2020-08-01T13:05:09Z"),
						Modified: parseDateTime(t, "2020-08-01T13:05:09Z"),
						Text:     "Note 2",
					}, {
						ID:       "81657665",
						Created:  parseDateTime(t, "2020-08-01T13:05:05Z"),
						Modified: parseDateTime(t, "2020-08-01T13:05:05Z"),
						Text:     "Note 1",
					}},
					Task: []Task{{
						ID:       "622345829",
						Due:      parseDateTime(t, "2018-08-06T20:30:00Z"),
						Added:    parseDateTime(t, "2018-08-05T09:30:53Z"),
						Priority: "N",
						Start:    parseDateTime(t, "2018-08-04T21:30:00Z"),
					}},
				}, {
					ID:       "358441583",
					Created:  parseDateTime(t, "2018-08-05T09:30:56Z"),
					Modified: parseDateTime(t, "2018-08-05T13:42:43Z"),
					Name:     "Задача 2",
					Source:   "js",
					Task: []Task{{
						ID:        "622345833",
						Due:       parseDateTime(t, "2018-08-05"),
						Added:     parseDateTime(t, "2018-08-05T09:30:56Z"),
						Completed: parseDateTime(t, "2018-08-05T13:42:43Z"),
						Priority:  "2",
						Estimate:  parseDuration(t, "PT1H12M"),
						Start:     parseDateTime(t, "2018-08-04"),
					}},
				},
			},
		}

		t.Run("Unmarshal", func(t *testing.T) {
			b := readTestdataFile(t, "rtm.tasks.getList.json")
			actual, err := new(Client).Tasks().client.Tasks().getListUnmarshal(b)
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})

		t.Run("Real", func(t *testing.T) {
			actual, err := GetClient(t).Tasks().GetList(Ctx, &TasksGetListParams{
				ListID: "43911488",
			})
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("Add", func(t *testing.T) {
		expectedListID := "43911488"
		expected := &TaskSeries{
			ID:       "467252622",
			Created:  parseDateTime(t, "2022-01-22T10:12:43Z"),
			Modified: parseDateTime(t, "2022-01-22T10:12:43Z"),
			Name:     "TestTasks/Add/Real",
			Source:   "api:XXX",
			Task: []Task{{
				ID:       "856005635",
				Added:    parseDateTime(t, "2022-01-22T10:12:43Z"),
				Priority: "N",
			}},
		}

		t.Run("Unmarshal", func(t *testing.T) {
			b := readTestdataFile(t, "rtm.tasks.add.json")
			actualListID, actual, err := new(Client).Tasks().addUnmarshal(b)
			require.NoError(t, err)
			assert.Equal(t, expectedListID, actualListID)
			assert.Equal(t, expected, actual)
		})

		t.Run("Real", func(t *testing.T) {
			timeline, err := GetClient(t).Timelines().Create(Ctx)
			require.NoError(t, err)

			listID, task, err := GetClient(t).Tasks().Add(Ctx, timeline, &TasksAddParams{
				ListID: "43911488",
				Name:   t.Name(),
			})
			require.NoError(t, err)
			t.Log(task)
			require.NotEmpty(t, task.Task)
			require.Equal(t, listID, "43911488")

			err = GetClient(t).Tasks().Delete(Ctx, timeline, &TasksDeleteParams{
				ListID:       "43911488",
				TaskSeriesID: task.ID,
				TaskID:       task.Task[0].ID,
			})
			require.NoError(t, err)
		})
	})
}
