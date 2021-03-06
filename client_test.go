package rtm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	t.Run("Call", func(t *testing.T) {
		t.Run("Error", func(t *testing.T) {
			client := GetClient(t)
			b, err := client.Call(Ctx, "no.such.method", nil)
			assert.Equal(t, &Error{Code: 112, Msg: `Method "no.such.method" not found`}, err)
			assert.EqualError(t, err, `112: Method "no.such.method" not found`)
			assert.Empty(t, b)
		})
	})
}
