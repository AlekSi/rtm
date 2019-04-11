# rtm

[![Actions](https://img.shields.io/badge/tested%20with-actions-success.svg?logo=github)](https://github.com/AlekSi/rtm/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/AlekSi/rtm)](https://goreportcard.com/report/github.com/AlekSi/rtm)
[![GoDoc](https://godoc.org/github.com/AlekSi/rtm?status.svg)](https://godoc.org/github.com/AlekSi/rtm)

Go client for [Remember The Milk API v2](https://www.rememberthemilk.com/services/api/).

Note: this product uses the Remember The Milk API but is not endorsed or certified by Remember The Milk.

# Methods

| API method                                                                                                              | Package method
|-------------------------------------------------------------------------------------------------------------------------|---------------
| [`rtm.auth.checkToken`](https://www.rememberthemilk.com/services/api/methods/rtm.auth.checkToken.rtm)                   | [`AuthService.CheckToken`](https://godoc.org/github.com/AlekSi/rtm#AuthService.CheckToken)
| [`rtm.auth.getFrob`](https://www.rememberthemilk.com/services/api/methods/rtm.auth.getFrob.rtm)                         | [`AuthService.GetFrob`](https://godoc.org/github.com/AlekSi/rtm#AuthService.GetFrob)
| [`rtm.auth.getToken`](https://www.rememberthemilk.com/services/api/methods/rtm.auth.getToken.rtm)                       | [`AuthService.GetToken`](https://godoc.org/github.com/AlekSi/rtm#AuthService.GetToken)
| [`rtm.contacts.add`](https://www.rememberthemilk.com/services/api/methods/rtm.contacts.add.rtm)                         | TODO
| [`rtm.contacts.delete`](https://www.rememberthemilk.com/services/api/methods/rtm.contacts.delete.rtm)                   | TODO
| [`rtm.contacts.getList`](https://www.rememberthemilk.com/services/api/methods/rtm.contacts.getList.rtm)                 | TODO
| [`rtm.groups.add`](https://www.rememberthemilk.com/services/api/methods/rtm.groups.add.rtm)                             | TODO
| [`rtm.groups.addContact`](https://www.rememberthemilk.com/services/api/methods/rtm.groups.addContact.rtm)               | TODO
| [`rtm.groups.delete`](https://www.rememberthemilk.com/services/api/methods/rtm.groups.delete.rtm)                       | TODO
| [`rtm.groups.getList`](https://www.rememberthemilk.com/services/api/methods/rtm.groups.getList.rtm)                     | TODO
| [`rtm.groups.removeContact`](https://www.rememberthemilk.com/services/api/methods/rtm.groups.removeContact.rtm)         | TODO
| [`rtm.lists.add`](https://www.rememberthemilk.com/services/api/methods/rtm.lists.add.rtm)                               | TODO
| [`rtm.lists.archive`](https://www.rememberthemilk.com/services/api/methods/rtm.lists.archive.rtm)                       | TODO
| [`rtm.lists.delete`](https://www.rememberthemilk.com/services/api/methods/rtm.lists.delete.rtm)                         | TODO
| [`rtm.lists.getList`](https://www.rememberthemilk.com/services/api/methods/rtm.lists.getList.rtm)                       | [`ListsService.GetList`](https://godoc.org/github.com/AlekSi/rtm#ListsService.GetList)
| [`rtm.lists.setDefaultList`](https://www.rememberthemilk.com/services/api/methods/rtm.lists.setDefaultList.rtm)         | TODO
| [`rtm.lists.setName`](https://www.rememberthemilk.com/services/api/methods/rtm.lists.setName.rtm)                       | TODO
| [`rtm.lists.unarchive`](https://www.rememberthemilk.com/services/api/methods/rtm.lists.unarchive.rtm)                   | TODO
| [`rtm.locations.getList`](https://www.rememberthemilk.com/services/api/methods/rtm.locations.getList.rtm)               | TODO
| [`rtm.reflection.getMethodInfo`](https://www.rememberthemilk.com/services/api/methods/rtm.reflection.getMethodInfo.rtm) | [`ReflectionService.GetMethodInfo`](https://godoc.org/github.com/AlekSi/rtm#ReflectionService.GetMethodInfo)
| [`rtm.reflection.getMethods`](https://www.rememberthemilk.com/services/api/methods/rtm.reflection.getMethods.rtm)       | [`ReflectionService.GetMethods`](https://godoc.org/github.com/AlekSi/rtm#ReflectionService.GetMethods)
| [`rtm.settings.getList`](https://www.rememberthemilk.com/services/api/methods/rtm.settings.getList.rtm)                 | TODO
| [`rtm.tasks.add`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.add.rtm)                               | [`TasksService.Add`](https://godoc.org/github.com/AlekSi/rtm#TasksService.Add)
| [`rtm.tasks.addTags`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.addTags.rtm)                       | TODO
| [`rtm.tasks.complete`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.complete.rtm)                     | TODO
| [`rtm.tasks.delete`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.delete.rtm)                         | TODO
| [`rtm.tasks.getList`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.getList.rtm)                       | [`TasksService.GetList`](https://godoc.org/github.com/AlekSi/rtm#TasksService.GetList)
| [`rtm.tasks.movePriority`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.movePriority.rtm)             | TODO
| [`rtm.tasks.moveTo`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.moveTo.rtm)                         | TODO
| [`rtm.tasks.notes.add`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.notes.add.rtm)                   | TODO
| [`rtm.tasks.notes.delete`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.notes.delete.rtm)             | TODO
| [`rtm.tasks.notes.edit`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.notes.edit.rtm)                 | TODO
| [`rtm.tasks.postpone`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.postpone.rtm)                     | TODO
| [`rtm.tasks.removeTags`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.removeTags.rtm)                 | TODO
| [`rtm.tasks.setDueDate`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.setDueDate.rtm)                 | TODO
| [`rtm.tasks.setEstimate`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.setEstimate.rtm)               | TODO
| [`rtm.tasks.setLocation`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.setLocation.rtm)               | TODO
| [`rtm.tasks.setName`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.setName.rtm)                       | TODO
| [`rtm.tasks.setParentTask`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.setParentTask.rtm)           | TODO
| [`rtm.tasks.setPriority`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.setPriority.rtm)               | TODO
| [`rtm.tasks.setRecurrence`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.setRecurrence.rtm)           | TODO
| [`rtm.tasks.setStartDate`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.setStartDate.rtm)             | TODO
| [`rtm.tasks.setTags`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.setTags.rtm)                       | TODO
| [`rtm.tasks.setURL`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.setURL.rtm)                         | TODO
| [`rtm.tasks.uncomplete`](https://www.rememberthemilk.com/services/api/methods/rtm.tasks.uncomplete.rtm)                 | TODO
| [`rtm.test.echo`](https://www.rememberthemilk.com/services/api/methods/rtm.test.echo.rtm)                               | [`TestService.Echo`](https://godoc.org/github.com/AlekSi/rtm#TestService.Echo)
| [`rtm.test.login`](https://www.rememberthemilk.com/services/api/methods/rtm.test.login.rtm)                             | [`TestService.Login`](https://godoc.org/github.com/AlekSi/rtm#TestService.Login)
| [`rtm.time.convert`](https://www.rememberthemilk.com/services/api/methods/rtm.time.convert.rtm)                         | TODO
| [`rtm.time.parse`](https://www.rememberthemilk.com/services/api/methods/rtm.time.parse.rtm)                             | TODO
| [`rtm.timelines.create`](https://www.rememberthemilk.com/services/api/methods/rtm.timelines.create.rtm)                 | [`TimelinesService.Create`](https://godoc.org/github.com/AlekSi/rtm#TimelinesService.Create)
| [`rtm.timezones.getList`](https://www.rememberthemilk.com/services/api/methods/rtm.timezones.getList.rtm)               | TODO
| [`rtm.transactions.undo`](https://www.rememberthemilk.com/services/api/methods/rtm.transactions.undo.rtm)               | TODO
