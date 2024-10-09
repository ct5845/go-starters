package group_api

import (
	"insightful/src/group/group_api/group_user_api"
	"sync"

	"github.com/danielgtaylor/huma/v2"
)

var createRoutes sync.Once

func Router(api huma.API) {
	createRoutes.Do(func() {
		CreateGroupRoute(api)
		DeleteGroupRoute(api)
		GetGroupsRoute(api)

		group_user_api.Router(api)
	})
}
