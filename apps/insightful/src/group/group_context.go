package group

import "sync"

type GroupContext struct {
	GroupService *GroupService
}

var groupContext GroupContext
var initContext sync.Once

func GetContext() *GroupContext {
	initContext.Do(func() {
		groupContext = GroupContext{GroupService: GetGroupCognitoService()}
	})

	return &groupContext
}

func GetGroupService() *GroupService {
	context := GetContext()

	return context.GroupService
}
