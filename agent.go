package main

import (
	"log"
	"net"

	agent_pb "github.com/alexcogojocaru/btrace-agent/proto-gen/btrace_agent"
	service "github.com/alexcogojocaru/btrace-agent/service"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Ldate)

	lis, err := net.Listen("tcp", ":4576")
	if err != nil {
		log.Fatal("Cannot create a net listener on port 4576")
	}

	grpcServer := grpc.NewServer()
	agentService := service.NewAgentService()

	// Register the agent service to the gRPC server and serve the listener
	agent_pb.RegisterAgentServiceServer(grpcServer, agentService)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve the gRPC server")
	}
}
