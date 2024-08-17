package response

import "time"

type GroupResponse struct {
	Code    string
	Message string
	Groups  []Group
}

type Group struct {
	ID        uint
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
