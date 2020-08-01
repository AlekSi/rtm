package rtm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListsGetList(t *testing.T) {
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

	t.Run("Decode", func(t *testing.T) {
		var actual listsGetListResponse
		unmarshalTestdataFile(t, "rtm.lists.getList.xml", &actual)
		assert.Equal(t, expected, actual.Lists)
	})

	t.Run("Real", func(t *testing.T) {
		actual, err := GetClient(t).Lists().GetList(Ctx)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
