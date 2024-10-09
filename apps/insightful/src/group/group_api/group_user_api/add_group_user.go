package group_user_api

import (
	"context"
	"insightful/src/group"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type AddGroupUserRequest struct {
	Group string `path:"group"`
	User  string `path:"user"`
}
type AddGroupUserResponse struct{}

func AddGroupUserRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "add-group-user",
		Method:        http.MethodPut,
		Path:          "/group/{group}/user/{user}",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"Group Users"},
	}, func(ctx context.Context, input *AddGroupUserRequest) (*AddGroupUserResponse, error) {
		service := *group.GetGroupService()

		err := service.AddUserToGroup(input.Group, input.User)

		if err != nil {
			return nil, err
		}

		return &AddGroupUserResponse{}, nil
	})
}
