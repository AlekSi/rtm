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
			_, err := new(Client).callJSONUnmarshal(b)
			require.NoError(t, err)
		})

		t.Run("Real", func(t *testing.T) {
			_, err := GetClient(t).CallJSON(Ctx, "rtm.test.echo", Args{"foo": "bar"})
			require.NoError(t, err)
		})
	})

	t.Run("Error", func(t *testing.T) {
		expected := &Error{
			Code: 112,
			Msg:  `Method "no.such.method" not found`,
		}

		t.Run("Unmarshal", func(t *testing.T) {
			b := readTestdataFile(t, "no.such.method.json")
			_, actual := new(Client).callJSONUnmarshal(b)
			assert.Equal(t, expected, actual)
		})

		t.Run("Real", func(t *testing.T) {
			_, actual := GetClient(t).CallJSON(Ctx, "no.such.method", nil)
			assert.Equal(t, expected, actual)
		})
	})
}
