package rtm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	t.Run("CheckToken", func(t *testing.T) {
		expected := &AuthInfo{
			Token: "XXX",
			Perms: "delete",
			User: AuthInfoUser{
				ID:       "6561679",
				UserName: "alexey.tester",
				FullName: "Alexey Tester",
			},
		}

		t.Run("Unmarshal", func(t *testing.T) {
			b := readTestdataFile(t, "rtm.auth.checkToken.json")
			actual, err := new(Client).Auth().checkTokenUnmarshal(b)
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})

		t.Run("Real", func(t *testing.T) {
			actual, err := GetClient(t).Auth().CheckToken(Ctx)
			require.NoError(t, err)
			actual.Token = "XXX"
			assert.Equal(t, expected, actual)
		})
	})
}
