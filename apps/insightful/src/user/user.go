package user

type User struct {
	ID                string `json:"sub"`
	Username          string
	Account           string
	Status            string `json:"status"`
	Email             string `json:"email"`
	PreferredUsername string
	Enabled           bool              `json:"enabled"`
	CreatedAt         string            `json:"userCreateDate"`
	UpdatedAt         string            `json:"userLastModifiedDate"`
	Attributes        map[string]string `json:"attributes"`
	Groups            []string          `json:"groups"`
}

type UserService interface {
	ListUsers(filter string, includeGroups bool) ([]User, error)
	GetUser(idOrEmail string) (*User, error)
	WhoAmI() (*User, error)
	WhatCanIDo() (*[]string, error)
	CreateUser(email, username string) (*User, error)
	UpdateUser(idOrEmail, email, username string) error
	DeleteUser(idOrEmail string) error
	EnableUser(username string) error
	DisableUser(username string) error
}
