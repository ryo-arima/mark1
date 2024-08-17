package request

type GroupRequest struct {
	Group Group
}

type Group struct {
	ID   uint
	UUID string
	Name string
}
