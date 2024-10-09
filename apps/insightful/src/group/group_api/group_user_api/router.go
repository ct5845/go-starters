package group_user_api

import (
	"sync"

	"github.com/danielgtaylor/huma/v2"
)

var initRouter sync.Once

func Router(api huma.API) {
	initRouter.Do(func() {
		AddGroupUserRoute(api)
		DeleteGroupUserRoute(api)
		GetGroupUsersRoute(api)
	})
}
