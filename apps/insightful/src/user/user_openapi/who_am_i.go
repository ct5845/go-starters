package user_openapi

import (
	"context"
	"insightful/src/user"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type WhoAmIRequest struct{}
type WhoAmIResponse struct {
	Body *user.User
}

func WhoAmIRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "whoami",
		Method:        http.MethodGet,
		Path:          "/whoami",
		DefaultStatus: http.StatusOK,
		Tags:          []string{"User"},
	}, func(ctx context.Context, input *WhoAmIRequest) (*WhoAmIResponse, error) {
		service := user.Context.UserService

		user, err := service.WhoAmI()

		if err != nil {
			return nil, err
		}

		return &WhoAmIResponse{Body: user}, nil
	})
}
