package response

import "time"

type MemberResponse struct {
	Code    string
	Message string
	Members []Member
}

type Member struct {
	ID        uint
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
