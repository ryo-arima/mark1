package response

type AgentResponse struct {
	Code    string
	Message string
	Agents  []Agent
}

type Agent struct {
	ID        uint   `json:"id"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	IP        string `json:"ip"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
