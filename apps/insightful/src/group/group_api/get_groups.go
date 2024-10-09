package group_api

import (
	"context"
	"insightful/src/group"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

type GetGroupsFilter struct {
	Sub string `query:"sub" required:"false" doc:"Group Sub/Id"`
}

type GetGroupsRequest struct {
	Filter string `query:"filter" doc:"KeyValue pair of Attributes and Values"`
}
type GetGroupsResponse struct {
	Body *[]group.Group
}

func GetGroupsRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "get-groups",
		Method:        http.MethodGet,
		Path:          "/group",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"Groups"},
	}, func(ctx context.Context, input *GetGroupsRequest) (*GetGroupsResponse, error) {
		groupService := *group.GetGroupService()

		filter := ""
		if input.Filter != "" {
			parts := strings.SplitN(input.Filter, "=", 2)
			if len(parts) == 2 {
				filter = strings.Trim(parts[1], `"`)
			}
		}

		groups, err := groupService.ListGroups(&filter)

		if err != nil {
			return nil, err
		}

		return &GetGroupsResponse{Body: groups}, nil
	})
}
