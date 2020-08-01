package rtm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListsGetListDecode(t *testing.T) {
	var actual listsGetListResponse
	unmarshalTestdataFile(t, "rtm.lists.getList.xml", &actual)
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
	assert.Equal(t, expected, actual.Lists)
}

func TestListsGetList(t *testing.T) {
	lists, err := GetClient(t).Lists().GetList(Ctx)
	require.NoError(t, err)
	expected := []List{
		{
			ID:       "43911488",
			Name:     "Inbox",
			Locked:   true,
			Archived: false,
			Position: -1,
			Smart:    false,
		}, {
			ID:       "43911489",
			Name:     "Sent",
			Locked:   true,
			Archived: false,
			Position: 1,
			Smart:    false,
		}, {
			ID:       "43911490",
			Name:     "Personal",
			Locked:   false,
			Archived: false,
			Position: 0,
			Smart:    false,
		}, {
			ID:       "43911491",
			Name:     "Business",
			Locked:   false,
			Archived: false,
			Position: 0,
			Smart:    false,
		}, {
			ID:       "43999477",
			Name:     "Smart",
			Locked:   false,
			Archived: false,
			Position: 0,
			Smart:    true,
		}, {
			ID:       "43999479",
			Name:     "Archived",
			Locked:   false,
			Archived: true,
			Position: 0,
			Smart:    false,
		},
	}
	assert.Equal(t, expected, lists)
}
