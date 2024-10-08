package user

type User struct {
	ID                string
	Username          string
	Account           string
	Status            string
	Email             string
	PreferredUsername string
	Enabled           bool
	CreatedAt         string
	UpdatedAt         string
	Attributes        map[string]string
	Groups            []string
}

type UserService interface {
	ListUsers(filter string, includeGroups bool) ([]User, error)
	GetUser(idOrEmail string) (User, error)
	WhoAmI() (User, error)
	WhatCanIDo() (*[]string, error)
	CreateUser(email, username, accountName string) (User, error)
	UpdateUser(idOrEmail, email, username, accountName string) error
	DeleteUser(idOrEmail string) error
	EnableUser(username string) error
	DisableUser(username string) error
}
