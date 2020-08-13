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
		`"2019-01-20T09:20:58Z"`: {time.Date(2019, 1, 20, 9, 20, 58, 0, time.UTC)},
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
