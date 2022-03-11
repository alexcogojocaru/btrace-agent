package main

import (
	"fmt"
	"log"
	"net"

	agent_pb "github.com/alexcogojocaru/btrace-agent/proto-gen/btrace_agent"
	service "github.com/alexcogojocaru/btrace-agent/service"
	"google.golang.org/grpc"
)

const host = "localhost"
const port = 4576

func main() {
	log.SetFlags(log.Lmicroseconds | log.Ldate)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Cannot create a net listener on port %d", port)
	}

	grpcServer := grpc.NewServer()
	agentService := service.NewAgentService()

	log.Printf("Starting Agent on %s:%d", host, port)
	// Register the agent service to the gRPC server and serve the listener
	agent_pb.RegisterAgentServiceServer(grpcServer, agentService)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve the gRPC server")
	}
}