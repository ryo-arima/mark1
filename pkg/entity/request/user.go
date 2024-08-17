package request

type UserRequest struct {
	User User `json:"user"`
}

type User struct {
	ID       string `json:"id"`
	UUID     string `json:"uuid"`
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
