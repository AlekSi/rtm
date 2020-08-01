package rtm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseDuration(tb testing.TB, s string) Duration {
	tb.Helper()

	var d Duration
	err := d.UnmarshalText([]byte(s))
	require.NoError(tb, err)
	return d
}

func TestDuration(t *testing.T) {
	for text, expected := range map[string]time.Duration{
		"":        0,
		"PT5M":    5 * time.Minute,
		"PT1H":    time.Hour,
		"PT1H12M": time.Hour + 12*time.Minute,
		"PT48H":   48 * time.Hour,
	} {
		t.Run(text, func(t *testing.T) {
			actual := Duration{-1}
			err := actual.UnmarshalText([]byte(text))
			require.NoError(t, err)
			assert.Equal(t, Duration{expected}, actual)
			actualText, err := actual.MarshalText()
			require.NoError(t, err)
			assert.Equal(t, text, string(actualText))
		})
	}
}
