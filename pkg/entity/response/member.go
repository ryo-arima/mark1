package response

type MemberResponse struct {
	Code    string
	Message string
	Members []Member
}

type Member struct {
	ID        string `json:"id"`
	UUID      string `json:"uuid"`
	GroupUUID string `json:"group_uuid"`
	UserUUID  string `json:"user_uuid"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
