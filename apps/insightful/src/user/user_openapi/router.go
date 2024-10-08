package user_openapi

import (
	"github.com/danielgtaylor/huma/v2"
)

func Router(api huma.API) {
	CreateUserRoute(api)
	DeleteUserRoute(api)
	GetUserRoute(api)
	GetUsersRoute(api)
	UpdateUserRoute(api)
	WhatCanIDoRoute(api)
	WhoAmIRoute(api)
}
