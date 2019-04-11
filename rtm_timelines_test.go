package rtm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimelines(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		id, err := GetClient(t).Timelines().Create(Ctx)
		require.NoError(t, err)
		assert.NotEmpty(t, id)
	})
}
