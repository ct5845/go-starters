package user_openapi

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type CreateUserRequest struct{}
type CreateUserResponse struct{}

func CreateUserRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "create-user",
		Method:        http.MethodPost,
		Path:          "/user",
		DefaultStatus: http.StatusCreated,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *CreateUserRequest) (*CreateUserResponse, error) {
		return nil, nil
	})
}
