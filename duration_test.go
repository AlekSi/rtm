package rtm

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseDuration(tb testing.TB, s string) time.Duration {
	tb.Helper()

	var d rtmDuration
	err := json.Unmarshal([]byte(`"`+s+`"`), &d)
	require.NoError(tb, err)
	return d.Duration
}

func TestDuration(t *testing.T) {
	for j, expected := range map[string]rtmDuration{
		`""`:        {0},
		`"PT5M"`:    {5 * time.Minute},
		`"PT1H"`:    {time.Hour},
		`"PT1H12M"`: {time.Hour + 12*time.Minute},
		`"PT48H"`:   {48 * time.Hour},
	} {
		t.Run(j, func(t *testing.T) {
			b := []byte(`{"estimate":` + j + `}`)
			var actual struct {
				Estimate rtmDuration `json:"estimate"`
			}
			actual.Estimate = rtmDuration{parseDuration(t, "PT1H02M")}

			err := json.Unmarshal(b, &actual)
			require.NoError(t, err)
			assert.Equal(t, expected, actual.Estimate)

			actualB, err := json.Marshal(actual)
			require.NoError(t, err)
			assert.Equal(t, string(b), string(actualB))
		})
	}
}
