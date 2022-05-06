package service

import (
	"context"
	"log"

	agent "github.com/alexcogojocaru/btrace-agent/proto-gen/btrace_proxy"
	"google.golang.org/grpc"
)

type AgentServiceImpl struct {
	CollectorClient agent.AgentClient
	agent.UnimplementedAgentServer
}

func NewAgentService() *AgentServiceImpl {
	agentHost := "localhost:4578"
	conn, err := grpc.Dial(agentHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial %s", agentHost)
	}

	return &AgentServiceImpl{
		CollectorClient: agent.NewAgentClient(conn),
	}
}

func (agentServiceImpl *AgentServiceImpl) Send(ctx context.Context, span *agent.Span) (*agent.Response, error) {
	log.Print(span)
	agentServiceImpl.CollectorClient.Send(ctx, span)

	return &agent.Response{}, nil
}
