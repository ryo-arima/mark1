package response

type CommandResponse struct {
	Command Command `json:"command"`
}

type Command struct {
	AgentUUID string          `json:"agent_uuid"`
	AgentName string          `json:"agent_name"`
	Results   []CommandResult `json:"results"`
}

type CommandResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
