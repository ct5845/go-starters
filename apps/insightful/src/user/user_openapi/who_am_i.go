package user_openapi

import (
	"context"
	"insightful/src/user"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type WhoAmIRequest struct{}
type WhoAMIResponse struct {
	Body *user.User
}

func WhoAmIRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "whoami",
		Method:        http.MethodGet,
		Path:          "/whoami",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *WhoAmIRequest) (*WhoAMIResponse, error) {
		return nil, nil
	})
}
