package rtm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	t.Run("CheckToken", func(t *testing.T) {
		c := GetClient(t)
		info, err := c.Auth().CheckToken(Ctx)
		require.NoError(t, err)
		assert.Equal(t, c.AuthToken, info.Token)
		assert.Equal(t, Delete, info.Perms)
		assert.NotEmpty(t, info.User.ID)
		assert.NotEmpty(t, info.User.UserName)
		assert.NotEmpty(t, info.User.FullName)
	})
}
