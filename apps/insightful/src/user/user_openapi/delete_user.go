package user_openapi

import (
	"context"
	"insightful/src/user"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeleteUserRequest struct {
	Id string `path:"id"`
}
type DeleteUserResponse struct{}

func DeleteUserRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "delete-user",
		Method:        http.MethodDelete,
		Path:          "/user/{id}",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *DeleteUserRequest) (*DeleteUserResponse, error) {
		service := user.Context.UserService

		err := service.DeleteUser(input.Id)

		if err != nil {
			return nil, err
		}

		return nil, nil
	})
}
