package group

import (
	"insightful/src/user"
	"time"
)

type Group struct {
	GroupName        string    `json:"name"`
	Description      string    `json:"description"`
	CreationDate     time.Time `json:"creationDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
}

type GroupService interface {
	AddUserToGroup(groupName, idOrEmail string) error
	RemoveUserFromGroup(groupName, idOrEmail string) error
	CreateGroup(name, description string) (*Group, error)
	DeleteGroup(name string) error
	ListGroups(filter *string) (*[]Group, error)
	ListUsersInGroup(groupName string) (*[]user.User, error)
}
