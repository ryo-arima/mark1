package request

type AgentRequest struct {
	Agent Agent `json:"agent"`
}

type Agent struct {
	ID     string `json:"id"`
	UUID   string `json:"uuid"`
	Name   string `json:"name" validate:"required"`
	IP     string `json:"ip" validate:"required"`
	Status string `json:"status"`
}
