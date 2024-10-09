package user_openapi

import (
	"context"
	"insightful/src/user"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

type GetUsersFilter struct {
	Sub string `query:"sub" required:"false" doc:"User Sub/Id"`
}

type GetUsersRequest struct {
	IncludeGroups bool   `query:"includeGroups" doc:"Include the users groups in the response"`
	Filter        string `query:"filter" doc:"KeyValue pair of UserAttributes and Values"`
}
type GetUsersResponse struct {
	Body *[]user.User
}

func GetUsersRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "get-users",
		Method:        http.MethodGet,
		Path:          "/user",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *GetUsersRequest) (*GetUsersResponse, error) {
		service := user.Context.UserService

		filter := ""
		if input.Filter != "" {
			parts := strings.SplitN(input.Filter, "=", 2)
			if len(parts) == 2 {
				filter = strings.Trim(parts[1], `"`)
			}
		}

		users, err := service.ListUsers(filter, input.IncludeGroups)

		if err != nil {
			return nil, err
		}

		return &GetUsersResponse{Body: &users}, nil
	})
}
