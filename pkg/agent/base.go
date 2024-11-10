package agent

import (
	"fmt"
	"log"
	"net"

	auto "ryo-arima/mark1/pkg/agent/auto"

	"google.golang.org/grpc"
)

func Main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	auto.RegisterAgentServer(s, &server{})
	fmt.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
