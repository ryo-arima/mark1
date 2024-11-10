package agent

import (
	"context"

	auto "ryo-arima/mark1/pkg/agent/auto"
)

func (agentServer *config.AgentServer) Executer(ctx context.Context, request *auto.CommandRequest) (*auto.CommandReply, error) {
	return &auto.CommandReply{Message: "Hello "}, nil
}
