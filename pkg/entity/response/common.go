package response

type TokenResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
