package user

import "time"

type User struct {
	Sub          string
	Status       string
	Email        string
	Enabled      bool
	Created      *time.Time
	LastModified *time.Time
	UpdatedAt    string
	Groups       []string
}
