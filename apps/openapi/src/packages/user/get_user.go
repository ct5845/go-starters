package user

import (
	"context"
	"net/http"
	"open_api/src/packages/cognito"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/danielgtaylor/huma/v2"
)

type GetUserRequest = cognitoidentityprovider.AdminGetUserInput

func mapUserFromAWS(from *cognitoidentityprovider.AdminGetUserOutput) *User {
	attributes := make(map[string]string)
	for _, attr := range from.UserAttributes {
		attributes[*attr.Name] = *attr.Value
	}

	return &User{
		Sub:          string(*from.Username),
		Email:        attributes["email"],
		Enabled:      from.Enabled,
		Status:       string(from.UserStatus),
		Created:      from.UserCreateDate,
		LastModified: from.UserLastModifiedDate,
	}
}

func GetUser(ctx context.Context, username string) (*User, error) {
	client := cognito.GetClient()
	clientConfig := cognito.GetConfig()

	request := &GetUserRequest{
		Username:   &username,
		UserPoolId: &clientConfig.UserPoolId,
	}

	awsUser, err := client.AdminGetUser(ctx, request)

	if err != nil {
		return nil, err
	}

	return mapUserFromAWS(awsUser), nil
}

type GetUserInput struct {
	Id string `path:"id" doc:"Id or Sub of the user your want to get"`
}

type GetUserResponse struct {
	Body *User
}

func GetUserRoute(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "get-user",
		Method:        http.MethodGet,
		Path:          "/user/{id}",
		DefaultStatus: http.StatusFound,
		Tags:          []string{"User", "Query"},
	},
		func(ctx context.Context, input *GetUserInput) (*GetUserResponse, error) {
			result, err := GetUser(ctx, input.Id)

			if err != nil {
				return nil, err
			}

			return &GetUserResponse{Body: result}, nil
		})
}
