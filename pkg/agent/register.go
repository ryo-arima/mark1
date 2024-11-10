package agent

import (
	auto "ryo-arima/mark1/pkg/agent/auto"

	"github.com/ryo-arima/mark1/pkg/config"
	"google.golang.org/grpc"
)

func NewRegister() {
	grpcServer := grpc.NewServer()
	auto.RegisterAgentServer(grpcServer, &config.AgentServer{})
}
