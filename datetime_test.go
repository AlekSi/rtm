package rtm

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseDateTime(tb testing.TB, s string) DateTime {
	tb.Helper()

	var t DateTime
	err := json.Unmarshal([]byte(`"`+s+`"`), &t)
	require.NoError(tb, err)
	return t
}

func TestDateTime(t *testing.T) {
	for j, expected := range map[string]DateTime{
		`"2019-01-20T00:20:58Z"`: {Time: time.Date(2019, 1, 20, 0, 20, 58, 0, time.UTC), hasTime: true},
		`"2018-08-05T23:30:00Z"`: {Time: time.Date(2018, 8, 5, 23, 30, 0, 0, time.UTC), hasTime: true},
		`"2019-01-20"`:           {Time: time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC), hasTime: false},
		`"2018-08-05"`:           {Time: time.Date(2018, 8, 5, 0, 0, 0, 0, time.UTC), hasTime: false},
		`""`:                     {},
	} {
		t.Run(j, func(t *testing.T) {
			b := []byte(`{"completed":` + j + `}`)
			var actual struct {
				Completed DateTime `json:"completed"`
			}
			actual.Completed = parseDateTime(t, "2020-01-02T03:04:05Z")

			err := json.Unmarshal(b, &actual)
			require.NoError(t, err)
			assert.Equal(t, expected, actual.Completed)

			actualB, err := json.Marshal(actual)
			require.NoError(t, err)
			assert.Equal(t, string(b), string(actualB))
		})
	}
}
