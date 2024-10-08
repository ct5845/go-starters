package user_openapi

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeleteUserRequest struct{}
type DeleteUserResponse struct{}

func DeleteUserRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "delete-user",
		Method:        http.MethodDelete,
		Path:          "/user/{id}",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *DeleteUserRequest) (*DeleteUserResponse, error) {
		return nil, nil
	})
}
