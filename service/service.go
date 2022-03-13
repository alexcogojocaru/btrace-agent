package service

import (
	"context"
	"log"

	agent_pb "github.com/alexcogojocaru/btrace-agent/proto-gen/btrace_agent"
	"google.golang.org/grpc/metadata"
)

type AgentServiceImpl struct {
	agent_pb.UnimplementedAgentServiceServer
}

func NewAgentService() *AgentServiceImpl {
	return &AgentServiceImpl{}
}

func (agentServiceImpl *AgentServiceImpl) StreamSpan(ctx context.Context, span *agent_pb.Span) (*agent_pb.Response, error) {
	meta, _ := metadata.FromIncomingContext(ctx)
	log.Print(meta)
	return &agent_pb.Response{}, nil
}

func (agentServiceImpl *AgentServiceImpl) StreamBatch(ctx context.Context, batch *agent_pb.BatchSpan) (*agent_pb.Response, error) {
	for _, span := range batch.GetSpans() {
		log.Print(span)
	}

	return &agent_pb.Response{}, nil
}
