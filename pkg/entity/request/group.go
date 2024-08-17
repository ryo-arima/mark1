package request

type GroupRequest struct {
	Group Group `json:"group"`
}

type Group struct {
	ID   string `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name" validate:"required"`
}
