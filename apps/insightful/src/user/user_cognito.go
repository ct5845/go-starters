package user

import (
	"context"
	"fmt"
	"insightful/src/cognito"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type UserCognitoService struct{}

var userCognitoService UserService

func MapCognitoUser(user types.UserType, groups []string) User {
	attributes := make(map[string]string)
	for _, attr := range user.Attributes {
		attributes[aws.ToString(attr.Name)] = aws.ToString(attr.Value)
	}
	return User{
		ID:         aws.ToString(user.Username),
		Username:   attributes["preferred_username"],
		Account:    attributes["custom:accountname"],
		Status:     string(user.UserStatus),
		Email:      attributes["email"],
		Enabled:    aws.ToBool(&user.Enabled),
		CreatedAt:  user.UserCreateDate.Format(time.RFC3339),
		UpdatedAt:  user.UserLastModifiedDate.Format(time.RFC3339),
		Attributes: attributes,
		Groups:     groups,
	}
}

func (us *UserCognitoService) addAttribute(attributes []types.AttributeType, name, value string) []types.AttributeType {
	if value != "" {
		attributes = append(attributes, types.AttributeType{Name: aws.String(name), Value: aws.String(value)})
	}
	return attributes
}

func (us *UserCognitoService) ListUsers(filter string, includeGroups bool) ([]User, error) {
	var users []User
	input := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
	}

	paginator := cognitoidentityprovider.NewListUsersPaginator(cognito.GetClient(), input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, fmt.Errorf("error listing users: %v", err)
		}

		userChan := make(chan User, len(page.Users))
		errChan := make(chan error, 1)

		for _, user := range page.Users {
			go func(user types.UserType) {
				var groups []string
				if includeGroups {
					groups, err = us.listGroupsForUser(aws.ToString(user.Username))
					if err != nil {
						errChan <- err
						return
					}
				}
				mappedUser := MapCognitoUser(user, groups)
				if filter == "" || us.matchesFilter(mappedUser, filter) {
					userChan <- mappedUser
				} else {
					userChan <- User{}
				}
			}(user)
		}

		for range page.Users {
			select {
			case user := <-userChan:
				if user.ID != "" {
					users = append(users, user)
				}
			case err := <-errChan:
				return nil, err
			}
		}
	}

	return users, nil
}

func (us *UserCognitoService) matchesFilter(user User, filter string) bool {
	filter = strings.ToLower(filter)

	return strings.Contains(strings.ToLower(user.Email), filter) ||
		strings.Contains(strings.ToLower(user.Username), filter) ||
		strings.Contains(strings.ToLower(user.ID), filter)
}

func (us *UserCognitoService) listGroupsForUser(username string) ([]string, error) {
	groupResp, err := cognito.GetClient().AdminListGroupsForUser(context.TODO(), &cognitoidentityprovider.AdminListGroupsForUserInput{
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
		Username:   aws.String(username),
	})
	if err != nil {
		return nil, fmt.Errorf("error listing groups for user: %v", err)
	}

	groups := make([]string, len(groupResp.Groups))
	for i, group := range groupResp.Groups {
		groups[i] = aws.ToString(group.GroupName)
	}
	return groups, nil
}

func (us *UserCognitoService) GetUser(idOrEmail string) (*User, error) {
	userOutput, err := cognito.GetClient().AdminGetUser(context.TODO(), &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
		Username:   aws.String(idOrEmail),
	})

	if err != nil {
		return &User{}, fmt.Errorf("error getting user: %v", err)
	}

	userType := types.UserType{
		Username:             userOutput.Username,
		UserStatus:           userOutput.UserStatus,
		Enabled:              userOutput.Enabled,
		UserCreateDate:       userOutput.UserCreateDate,
		UserLastModifiedDate: userOutput.UserLastModifiedDate,
		Attributes:           userOutput.UserAttributes,
	}

	user := MapCognitoUser(userType, nil)

	return &user, nil
}

func (us *UserCognitoService) WhoAmI() (*User, error) {
	// TODO get context of who user is from Auth

	userOutput, err := cognito.GetClient().AdminGetUser(context.TODO(), &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
		Username:   aws.String("cturner@alterian.com"),
	})

	if err != nil {
		return &User{}, fmt.Errorf("error getting user: %v", err)
	}

	userType := types.UserType{
		Username:             userOutput.Username,
		UserStatus:           userOutput.UserStatus,
		Enabled:              userOutput.Enabled,
		UserCreateDate:       userOutput.UserCreateDate,
		UserLastModifiedDate: userOutput.UserLastModifiedDate,
		Attributes:           userOutput.UserAttributes,
	}

	user := MapCognitoUser(userType, nil)

	return &user, nil
}

func (us *UserCognitoService) WhatCanIDo() (*[]string, error) {
	// TODO get context of who user is from Auth

	return &[]string{
		"CreateGroup",
		"CreateUser",
		"DeleteGroup",
		"DeleteUser",
		"LinkPolicyTemplateToPrincipal",
		"ListGroups",
		"ListPolicyTemplates",
		"ListUsers",
		"ListWorkspaces",
		"UnlinkPolicyTemplateFromPrincipal",
		"UpdateGroup",
		"UpdateUser",
		"ViewAdmin",
		"ViewAlpha",
	}, nil
}

func (us *UserCognitoService) CreateUser(email, preferredUsername string) (*User, error) {
	attributes := []types.AttributeType{
		{Name: aws.String("email"), Value: aws.String(email)},
		{Name: aws.String("preferred_username"), Value: &preferredUsername},
		{Name: aws.String("email_verified"), Value: aws.String("True")},
	}

	// attributes = us.addAttribute(attributes, "preferred_username", preferredUsername)
	// attributes = append(attributes, types.AttributeType{Name: aws.String("email_verified"), Value: aws.String("True")})

	request := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:             aws.String(cognito.GetConfig().UserPoolId),
		DesiredDeliveryMediums: []types.DeliveryMediumType{types.DeliveryMediumTypeEmail},
		UserAttributes:         attributes,
		Username:               aws.String(email),
	}

	resp, err := cognito.GetClient().AdminCreateUser(context.TODO(), request)
	if err != nil {
		return &User{}, fmt.Errorf("error creating user: %v", err)
	}

	user := MapCognitoUser(*resp.User, nil)

	return &user, nil
}

func (us *UserCognitoService) UpdateUser(idOrEmail, email, username string) error {
	var attributes []types.AttributeType

	attributes = us.addAttribute(attributes, "email", email)
	attributes = us.addAttribute(attributes, "email_verified", "True")
	attributes = us.addAttribute(attributes, "preferred_username", username)

	if len(attributes) > 0 {
		_, err := cognito.GetClient().AdminUpdateUserAttributes(context.TODO(), &cognitoidentityprovider.AdminUpdateUserAttributesInput{
			UserPoolId:     aws.String(cognito.GetConfig().UserPoolId),
			Username:       aws.String(idOrEmail),
			UserAttributes: attributes,
		})
		if err != nil {
			return fmt.Errorf("failed to update user attributes: %w", err)
		}
	}

	return nil
}

func (us *UserCognitoService) DeleteUser(idOrEmail string) error {
	_, err := cognito.GetClient().AdminDeleteUser(context.TODO(), &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
		Username:   aws.String(idOrEmail),
	})

	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	return nil
}

func (us *UserCognitoService) EnableUser(username string) error {
	input := &cognitoidentityprovider.AdminEnableUserInput{
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
		Username:   aws.String(username),
	}

	if _, err := cognito.GetClient().AdminEnableUser(context.TODO(), input); err != nil {
		return fmt.Errorf("failed to enable user %s: %w", username, err)
	}

	return nil
}

func (us *UserCognitoService) DisableUser(username string) error {
	input := &cognitoidentityprovider.AdminDisableUserInput{
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
		Username:   aws.String(username),
	}

	if _, err := cognito.GetClient().AdminDisableUser(context.TODO(), input); err != nil {
		return fmt.Errorf("failed to disable user %s: %w", username, err)
	}

	return nil
}

func init() {
	userCognitoService = &UserCognitoService{}
}
