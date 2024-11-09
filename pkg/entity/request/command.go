package request

type CommandRequest struct {
	Command Command `json:"command"`
}

type Command struct {
	AgentUUID string        `json:"agent_uuid"`
	AgentName string        `json:"agent_name"`
	Body      []CommandBody `json:"body"`
}

type CommandBody struct {
	Environments string `json:"environments"`
	Shell        string `json:"shell"`
	Args         string `json:"args"`
	Timeout      int    `json:"timeout"`
}
