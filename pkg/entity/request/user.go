package request

type UserRequest struct {
	User User `json:"user"`
}

type User struct {
	ID       string `json:"id"`
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   string `json:"status"`
}
