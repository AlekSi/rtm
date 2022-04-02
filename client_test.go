package rtm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Run("NoError", func(t *testing.T) {
		t.Run("Unmarshal", func(t *testing.T) {
			b := readTestdataFile(t, "rtm.test.echo.json")
			err := checkErrorResponse(b)
			require.NoError(t, err)
		})

		t.Run("Real", func(t *testing.T) {
			_, err := GetClient(t).Call(Ctx, "rtm.test.echo", Args{"foo": "bar"})
			require.NoError(t, err)
		})
	})

	t.Run("Error", func(t *testing.T) {
		expectedErr := &Error{
			Code: 112,
			Msg:  `Method "no.such.method" not found`,
		}

		t.Run("Unmarshal", func(t *testing.T) {
			b := readTestdataFile(t, "no.such.method.json")
			err := checkErrorResponse(b)
			assert.Equal(t, expectedErr, err)
		})

		t.Run("Real", func(t *testing.T) {
			_, actual := GetClient(t).Call(Ctx, "no.such.method", nil)
			assert.Equal(t, expectedErr, actual)
		})
	})
}
