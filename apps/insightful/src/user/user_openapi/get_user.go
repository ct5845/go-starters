package user_openapi

import (
	"context"
	"insightful/src/user"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetUserRequest struct {
	Id string `path:"id" doc:"Email, Sub/Id, or Username of the User"`
}
type GetUserResponse struct {
	Body *user.User
}

func GetUserRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "get-user",
		Method:        http.MethodGet,
		Path:          "/user/{id}",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *GetUserRequest) (*GetUserResponse, error) {
		service := user.Context.UserService

		user, err := service.GetUser(input.Id)

		if err != nil {
			return nil, err
		}

		return &GetUserResponse{Body: user}, nil
	})
}
