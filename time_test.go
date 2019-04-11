package rtm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	for text, expected := range map[string]time.Time{
		"2019-01-20T09:20:58Z": time.Date(2019, 1, 20, 9, 20, 58, 0, time.UTC),
	} {
		t.Run(text, func(t *testing.T) {
			actual := Time{}
			err := actual.UnmarshalText([]byte(text))
			require.NoError(t, err)
			assert.Equal(t, Time{expected}, actual)

			actualText, err := actual.MarshalText()
			require.NoError(t, err)
			assert.Equal(t, text, string(actualText))
		})
	}
}
