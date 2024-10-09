package group_api

import (
	"context"
	"insightful/src/group"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type CreateGroupRequest struct {
	Name string `path:"name"`
	Body struct {
		Description string `json:"description"`
	}
}
type CreateGroupResponse struct {
	Body *group.Group
}

func CreateGroupRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "create-group",
		Method:        http.MethodPut,
		Path:          "/group/{name}",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"Groups"},
	}, func(ctx context.Context, input *CreateGroupRequest) (*CreateGroupResponse, error) {
		service := *group.GetGroupService()

		group, err := service.CreateGroup(input.Name, input.Body.Description)

		if err != nil {
			return nil, err
		}

		return &CreateGroupResponse{Body: group}, nil
	})
}
