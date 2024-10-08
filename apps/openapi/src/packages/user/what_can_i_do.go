package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UserPermissions = []string

type WhatCanIDoReponse struct {
	Body *UserPermissions
}

func WhatCanIDo(ctx context.Context) (*UserPermissions, error) {
	// TODO get user context from ?
	return &UserPermissions{}, nil
}

func WhatCanIDoRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "whatcanido",
		Method:        http.MethodGet,
		Path:          "/whatcanido",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	},
		func(ctx context.Context, input *struct{}) (*WhatCanIDoReponse, error) {
			result, err := WhatCanIDo(ctx)

			if err != nil {
				return nil, err
			}

			return &WhatCanIDoReponse{Body: result}, nil
		})
}
