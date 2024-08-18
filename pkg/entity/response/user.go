package response

type UserResponse struct {
	Code    string
	Message string
	Users   []User
}

type User struct {
	ID        string `json:"id"`
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
