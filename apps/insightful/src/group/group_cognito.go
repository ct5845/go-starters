package group

import (
	"context"
	"insightful/src/cognito"
	"insightful/src/user"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type GroupCognitoService struct{}

var groupCognitoService GroupService
var once sync.Once

func GetGroupCognitoService() *GroupService {
	once.Do(func() {
		groupCognitoService = &GroupCognitoService{}
	})

	return &groupCognitoService
}

// TODO should probably return pointer to the object here
func mapGroup(group types.GroupType) Group {
	return Group{
		GroupName:        aws.ToString(group.GroupName),
		Description:      aws.ToString(group.Description),
		CreationDate:     aws.ToTime(group.CreationDate),
		LastModifiedDate: aws.ToTime(group.LastModifiedDate),
	}
}

func (s *GroupCognitoService) AddUserToGroup(groupName, idOrEmail string) error {
	groupName = strings.ToLower(strings.TrimSpace(groupName))
	_, err := cognito.GetClient().AdminAddUserToGroup(context.TODO(), &cognitoidentityprovider.AdminAddUserToGroupInput{
		GroupName:  &groupName,
		Username:   &idOrEmail,
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
	})
	return err
}

func (s *GroupCognitoService) RemoveUserFromGroup(groupName, idOrEmail string) error {
	groupName = strings.ToLower(strings.TrimSpace(groupName))
	_, err := cognito.GetClient().AdminRemoveUserFromGroup(context.TODO(), &cognitoidentityprovider.AdminRemoveUserFromGroupInput{
		GroupName:  &groupName,
		Username:   &idOrEmail,
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
	})
	return err
}

func (s *GroupCognitoService) CreateGroup(name, description string) (*Group, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	resp, err := cognito.GetClient().CreateGroup(context.TODO(), &cognitoidentityprovider.CreateGroupInput{
		UserPoolId:  aws.String(cognito.GetConfig().UserPoolId),
		GroupName:   &name,
		Description: &description,
	})
	if err != nil {
		return &Group{}, err
	}

	group := mapGroup(*resp.Group)

	return &group, nil
}

func (s *GroupCognitoService) DeleteGroup(name string) error {
	name = strings.ToLower(strings.TrimSpace(name))
	_, err := cognito.GetClient().DeleteGroup(context.TODO(), &cognitoidentityprovider.DeleteGroupInput{
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
		GroupName:  &name,
	})
	return err
}

func (s *GroupCognitoService) ListGroups(filter *string) (*[]Group, error) {
	var groups []Group
	paginator := cognitoidentityprovider.NewListGroupsPaginator(cognito.GetClient(), &cognitoidentityprovider.ListGroupsInput{
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		for _, group := range page.Groups {
			if filter == nil || strings.Contains(strings.ToLower(aws.ToString(group.GroupName)), strings.ToLower(*filter)) ||
				strings.Contains(strings.ToLower(aws.ToString(group.Description)), strings.ToLower(*filter)) {
				groups = append(groups, mapGroup(group))
			}
		}
	}

	return &groups, nil
}

func (s *GroupCognitoService) ListUsersInGroup(groupName string) (*[]user.User, error) {
	var users []user.User
	paginator := cognitoidentityprovider.NewListUsersInGroupPaginator(cognito.GetClient(), &cognitoidentityprovider.ListUsersInGroupInput{
		UserPoolId: aws.String(cognito.GetConfig().UserPoolId),
		GroupName:  &groupName,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		for _, pageUser := range page.Users {
			users = append(users, user.MapCognitoUser(pageUser, nil))
		}
	}

	return &users, nil
}
