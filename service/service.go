package service

import (
	"context"
	"log"

	agent_pb "github.com/alexcogojocaru/btrace-agent/proto-gen/btrace_agent"
)

type AgentServiceImpl struct {
	agent_pb.UnimplementedAgentServiceServer
}

func NewAgentService() *AgentServiceImpl {
	return &AgentServiceImpl{}
}

func (agentServiceImpl *AgentServiceImpl) StreamSpan(ctx context.Context, span *agent_pb.Span) (*agent_pb.Response, error) {
	log.Print("Received a StreamSpan request")
	log.Print(span)
	return &agent_pb.Response{}, nil
}

func (agentServiceImpl *AgentServiceImpl) StreamBatch(ctx context.Context, batch *agent_pb.BatchSpan) (*agent_pb.Response, error) {
	log.Print("Received a StreamBatch request")

	for _, span := range batch.GetSpans() {
		log.Print(span)
	}

	return &agent_pb.Response{}, nil
}
