package request

type MemberRequest struct {
	Member Member
}

type Member struct {
	ID        uint
	UUID      string
	GroupUUID string
	UserUUID  string
	Role      string
}
