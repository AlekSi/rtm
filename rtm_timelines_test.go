package rtm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimelines(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		expected := "1408177753"

		t.Run("Unmarshal", func(t *testing.T) {
			b := readTestdataFile(t, "rtm.timelines.create.json")
			actual, err := new(Client).Timelines().createUnmarshal(b)
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})

		t.Run("Real", func(t *testing.T) {
			actual, err := GetClient(t).Timelines().Create(Ctx)
			require.NoError(t, err)
			assert.NotEmpty(t, actual)
		})
	})
}
