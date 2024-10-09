package group_user_api

import (
	"context"
	"insightful/src/group"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeleteGroupUserRequest struct {
	Group string `path:"group"`
	User  string `path:"user"`
}
type DeleteGroupUserResponse struct{}

func DeleteGroupUserRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "delete-group-user",
		Method:        http.MethodDelete,
		Path:          "/group/{group}/user/{user}",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"Group Users"},
	}, func(ctx context.Context, input *DeleteGroupUserRequest) (*DeleteGroupUserResponse, error) {
		service := *group.GetGroupService()

		err := service.RemoveUserFromGroup(input.Group, input.User)

		if err != nil {
			return nil, err
		}

		return &DeleteGroupUserResponse{}, nil
	})
}
