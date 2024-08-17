package response

import "time"

type UserResponse struct {
	Code    string
	Message string
	Users   []User
}

type User struct {
	ID        uint
	UUID      string
	Email     string
	Name      string
	Password  string
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
