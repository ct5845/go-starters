package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func WhoAmI(ctx context.Context) (*User, error) {
	// TODO get user context from ?
	return GetUser(ctx, "cturner@alterian.com")
}

type WhoAmIReponse struct {
	Body *User
}

func WhoAmIRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "whoami",
		Method:        http.MethodGet,
		Path:          "/whoami",
		DefaultStatus: http.StatusFound,
		Tags:          []string{"User"},
	},
		func(ctx context.Context, input *struct{}) (*WhoAmIReponse, error) {
			result, err := WhoAmI(ctx)

			if err != nil {
				return nil, err
			}

			return &WhoAmIReponse{Body: result}, nil
		})
}
