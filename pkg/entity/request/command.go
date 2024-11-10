package request

type P2PCommandRequest struct {
	Mode    string     `json:"mode"` //RTC or DTN
	Command P2PCommand `json:"command"`
}

type P2PCommand struct {
	SrcAgentUUID string           `json:"src_agent_uuid"`
	SrcAgentName string           `json:"src_agent_name"`
	DstAgentUUID string           `json:"dst_agent_uuid"`
	DstAgentName string           `json:"dst_agent_name"`
	Body         []P2PCommandBody `json:"body"`
}

type P2PCommandBody struct {
	Environments string `json:"environments"`
	Shell        string `json:"shell"`
	Args         string `json:"args"`
	Timeout      int    `json:"timeout"`
}

type SingleCommandRequest struct {
	Command SingleCommand `json:"command"`
}

type SingleCommand struct {
	AgentUUID string              `json:"agent_uuid"`
	AgentName string              `json:"agent_name"`
	Body      []SingleCommandBody `json:"body"`
}

type SingleCommandBody struct {
	Environments string `json:"environments"`
	Shell        string `json:"shell"`
	Args         string `json:"args"`
	Timeout      int    `json:"timeout"`
}
