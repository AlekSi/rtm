package rtm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReflection(t *testing.T) {
	t.Run("GetMethodInfo", func(t *testing.T) {
		actual, err := GetClient(t).Reflection().GetMethodInfo(Ctx, "rtm.test.login")
		require.NoError(t, err)
		expected := &MethodInfo{
			Name:       "rtm.test.login",
			NeedsLogin: true,
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("GetMethods", func(t *testing.T) {
		actual, err := GetClient(t).Reflection().GetMethods(Ctx)
		require.NoError(t, err)
		expected := []string{
			"rtm.auth.checkToken",
			"rtm.auth.getFrob",
			"rtm.auth.getToken",
			"rtm.contacts.add",
			"rtm.contacts.delete",
			"rtm.contacts.getList",
			"rtm.groups.add",
			"rtm.groups.addContact",
			"rtm.groups.delete",
			"rtm.groups.getList",
			"rtm.groups.removeContact",
			"rtm.lists.add",
			"rtm.lists.archive",
			"rtm.lists.delete",
			"rtm.lists.getList",
			"rtm.lists.setDefaultList",
			"rtm.lists.setName",
			"rtm.lists.unarchive",
			"rtm.locations.getList",
			"rtm.reflection.getMethodInfo",
			"rtm.reflection.getMethods",
			"rtm.settings.getList",
			"rtm.tasks.add",
			"rtm.tasks.addTags",
			"rtm.tasks.complete",
			"rtm.tasks.delete",
			"rtm.tasks.getList",
			"rtm.tasks.movePriority",
			"rtm.tasks.moveTo",
			"rtm.tasks.notes.add",
			"rtm.tasks.notes.delete",
			"rtm.tasks.notes.edit",
			"rtm.tasks.postpone",
			"rtm.tasks.removeTags",
			"rtm.tasks.setDueDate",
			"rtm.tasks.setEstimate",
			"rtm.tasks.setLocation",
			"rtm.tasks.setName",
			"rtm.tasks.setParentTask",
			"rtm.tasks.setPriority",
			"rtm.tasks.setRecurrence",
			"rtm.tasks.setStartDate",
			"rtm.tasks.setTags",
			"rtm.tasks.setURL",
			"rtm.tasks.uncomplete",
			"rtm.test.echo",
			"rtm.test.login",
			"rtm.time.convert",
			"rtm.time.parse",
			"rtm.timelines.create",
			"rtm.timezones.getList",
			"rtm.transactions.undo",
		}
		assert.Equal(t, expected, actual)
	})
}
