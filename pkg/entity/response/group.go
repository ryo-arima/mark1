package response

type GroupResponse struct {
	Code    string
	Message string
	Groups  []Group
}

type Group struct {
	ID        uint   `json:"id"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
