package user_openapi

import (
	"context"
	"insightful/src/user"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UpdateUserRequest struct {
	Id   string `path:"id"`
	Body struct {
		_        struct{} `json:"-" additionalProperties:"true"`
		Email    *string  `json:"email,omitempty" minimum:"3" maximum:"255" dependentRequired:"preferred_username" nullable:"true"`
		Username *string  `json:"preferred_username,omitempty" minimum:"3" maximum:"255"  nullable:"true"`
		Enabled  *bool    `json:"enabled,omitempty" nullable:"true"`
	}
}
type UpdateUserResponse struct{}

func UpdateUserRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "update-user",
		Method:        http.MethodPatch,
		Path:          "/user/{id}",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *UpdateUserRequest) (*UpdateUserResponse, error) {
		service := user.Context.UserService

		if input.Body.Email != nil && input.Body.Username != nil {
			err := service.UpdateUser(input.Id, *input.Body.Email, *input.Body.Username)

			if err != nil {
				return nil, err
			}
		} else if input.Body.Enabled != nil {
			if *input.Body.Enabled {
				err := service.EnableUser(input.Id)

				if err != nil {
					return nil, err
				}
			} else {
				err := service.DisableUser(input.Id)

				if err != nil {
					return nil, err
				}
			}
		}

		return nil, nil
	})
}
