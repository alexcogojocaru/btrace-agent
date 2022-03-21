package service

import (
	"context"
	"log"

	agent_pb "github.com/alexcogojocaru/btrace-agent/proto-gen/btrace_agent"
	collector_pb "github.com/alexcogojocaru/btrace-agent/proto-gen/btrace_collector"
	"google.golang.org/grpc"
)

type AgentServiceImpl struct {
	agent_pb.UnimplementedAgentServiceServer
	CollectorClient collector_pb.CollectorServiceClient
}

func NormalizeSpan(span *agent_pb.Span) collector_pb.SpanC {
	return collector_pb.SpanC{
		Name: span.Name,
		CurrentContext: &collector_pb.ContextC{
			TraceID: span.CurrentContext.TraceID,
			SpanID:  span.CurrentContext.SpanID,
		},
		ParentContext: &collector_pb.ContextC{
			TraceID: span.ParentContext.TraceID,
			SpanID:  span.ParentContext.SpanID,
		},
		Timestamp: &collector_pb.TimestampC{
			Started:  span.Timestamp.Started,
			Ended:    span.Timestamp.Ended,
			Duration: span.Timestamp.Duration,
		},
	}
}

func NewAgentService() *AgentServiceImpl {
	conn, err := grpc.Dial("localhost:4578", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot dial localhost:4578")
	}

	client := collector_pb.NewCollectorServiceClient(conn)

	return &AgentServiceImpl{
		CollectorClient: client,
	}
}

func (agentServiceImpl *AgentServiceImpl) StreamSpan(ctx context.Context, span *agent_pb.Span) (*agent_pb.Response, error) {
	// meta, _ := metadata.FromIncomingContext(ctx)
	// log.Print(meta)
	log.Print(span)

	collectorSpan := NormalizeSpan(span)
	agentServiceImpl.CollectorClient.StreamSpan(ctx, &collectorSpan)

	return &agent_pb.Response{}, nil
}

func (agentServiceImpl *AgentServiceImpl) StreamBatch(ctx context.Context, batch *agent_pb.BatchSpan) (*agent_pb.Response, error) {
	for _, span := range batch.GetSpans() {
		log.Print(span)
	}

	return &agent_pb.Response{}, nil
}
