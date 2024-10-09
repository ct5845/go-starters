package user_openapi

import (
	"context"
	"insightful/src/user"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type CreateUserRequest struct {
	Body struct {
		Email    string `json:"email" minimum:"3" maximum:"255"`
		Username string `json:"preferred_username" minimum:"3" maximum:"255"`
	}
}
type CreateUserResponse struct {
	Body *user.User
}

func CreateUserRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "create-user",
		Method:        http.MethodPost,
		Path:          "/user",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *CreateUserRequest) (*CreateUserResponse, error) {
		service := user.Context.UserService

		user, err := service.CreateUser(input.Body.Email, input.Body.Username)

		if err != nil {
			return nil, err
		}

		return &CreateUserResponse{Body: user}, nil
	})
}
