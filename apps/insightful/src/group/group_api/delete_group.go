package group_api

import (
	"context"
	"insightful/src/group"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeleteGroupRequest struct {
	Name string `path:"name"`
}
type DeleteGroupResponse struct{}

func DeleteGroupRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "delete-group",
		Method:        http.MethodDelete,
		Path:          "/group/{name}",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"Groups"},
	}, func(ctx context.Context, input *DeleteGroupRequest) (*DeleteGroupResponse, error) {
		service := *group.GetGroupService()

		err := service.DeleteGroup(input.Name)

		if err != nil {
			return nil, err
		}

		return &DeleteGroupResponse{}, nil
	})
}
