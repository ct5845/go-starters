package group_user_api

import (
	"context"
	"insightful/src/group"
	"insightful/src/user"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetGroupUsersRequest struct {
	Group string `path:"group"`
}
type GetGroupUsersResponse struct {
	Body *[]user.User
}

func GetGroupUsersRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "get-group-users",
		Method:        http.MethodGet,
		Path:          "/group/{group}/user",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"Group Users"},
	}, func(ctx context.Context, input *GetGroupUsersRequest) (*GetGroupUsersResponse, error) {
		service := *group.GetGroupService()

		users, err := service.ListUsersInGroup(input.Group)

		if err != nil {
			return nil, err
		}

		return &GetGroupUsersResponse{Body: users}, nil
	})
}
