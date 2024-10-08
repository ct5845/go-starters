package user_openapi

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UpdateUserRequest struct{}
type UpdateUserResponse struct{}

func UpdateUserRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "update-user",
		Method:        http.MethodPatch,
		Path:          "/user/{id}",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *UpdateUserRequest) (*UpdateUserResponse, error) {
		return nil, nil
	})
}
