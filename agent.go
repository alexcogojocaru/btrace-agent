package main

import (
	"fmt"
	"log"
	"net"

	"github.com/alexcogojocaru/btrace-agent/config"
	agent "github.com/alexcogojocaru/btrace-agent/proto-gen/btrace_proxy"
	service "github.com/alexcogojocaru/btrace-agent/service"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Ldate)

	conf, err := config.ParseConfig("config/config.yml")
	if err != nil {
		log.Fatal("Error while parsing the config file")
	}

	uri := fmt.Sprintf("%s:%d", conf.Deploy.Hostname, conf.Deploy.Port)
	collectorUri := fmt.Sprintf("%s:%d", conf.Connections.Collector.Hostname, conf.Connections.Collector.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Deploy.Port))
	if err != nil {
		log.Fatalf("Cannot create a net listener on port %d", conf.Deploy.Port)
	}

	grpcServer := grpc.NewServer()
	agentService := service.NewAgentService(collectorUri)

	log.Printf("Agent started on %s", uri)
	agent.RegisterAgentServer(grpcServer, agentService)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve the gRPC server")
	}
}
