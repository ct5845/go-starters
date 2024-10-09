package user_openapi

import (
	"context"
	"insightful/src/user"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type WhatCanIDoRequest struct{}
type WhatCanIDoResponse struct {
	Body *[]string
}

func WhatCanIDoRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "whatcanido",
		Method:        http.MethodGet,
		Path:          "/whatcanido",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *WhatCanIDoRequest) (*WhatCanIDoResponse, error) {
		service := user.Context.UserService

		whatcanido, err := service.WhatCanIDo()

		if err != nil {
			return nil, err
		}

		return &WhatCanIDoResponse{Body: whatcanido}, nil
	})
}
