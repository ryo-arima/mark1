package request

type MemberRequest struct {
	Member Member `json:"member"`
}

type Member struct {
	ID        string `json:"id"`
	UUID      string `json:"uuid"`
	GroupUUID string `json:"group_uuid" validate:"required"`
	UserUUID  string `json:"user_uuid" validate:"required"`
	Role      string `json:"role"`
}
