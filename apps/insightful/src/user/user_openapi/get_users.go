package user_openapi

import (
	"context"
	"insightful/src/user"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetUsersRequest struct {
	Filter string `path:"filter" docs:"String based filter for filtered users"`
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
		return nil, nil
	})
}
