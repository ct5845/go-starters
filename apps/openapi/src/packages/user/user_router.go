package user

import (
	"github.com/danielgtaylor/huma/v2"
)

func Router(api huma.API) {
	WhoAmIRoute(api)
	WhatCanIDoRoute(api)
	GetUserRoute(api)
}
