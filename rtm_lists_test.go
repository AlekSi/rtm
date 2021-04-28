package rtm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLists(t *testing.T) {
	t.Run("GetList", func(t *testing.T) {
		expected := []List{{
			ID:       "43911488",
			Name:     "Inbox",
			Locked:   true,
			Position: -1,
		}, {
			ID:       "43911489",
			Name:     "Sent",
			Locked:   true,
			Position: 1,
		}, {
			ID:   "43911490",
			Name: "Personal",
		}, {
			ID:   "43911491",
			Name: "Business",
		}, {
			ID:    "43999477",
			Name:  "Smart",
			Smart: true,
		}, {
			ID:       "43999479",
			Name:     "Archived",
			Archived: true,
		}}

		t.Run("Unmarshal", func(t *testing.T) {
			b := readTestdataFile(t, "rtm.lists.getList.json")
			actual, err := new(Client).Lists().getListUnmarshal(b)
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})

		t.Run("Real", func(t *testing.T) {
			actual, err := GetClient(t).Lists().GetList(Ctx)
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	})
}
