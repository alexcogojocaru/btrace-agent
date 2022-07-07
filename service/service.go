package service

import (
	"context"
	"log"
	"time"

	agent "github.com/alexcogojocaru/btrace-agent/proto-gen/btrace_proxy"
	"google.golang.org/grpc"
)

const POLLING_INTERVAL = 1
const BUFFER_SIZE = 50

type AgentServiceImpl struct {
	agent.UnimplementedExporterServer
	CollectorClient agent.ExporterClient
	Buffer          chan *agent.Span
}

func NewAgentService(collectorUri string) *AgentServiceImpl {
	conn, err := grpc.Dial(collectorUri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial %s", collectorUri)
	}

	service := &AgentServiceImpl{
		CollectorClient: agent.NewExporterClient(conn),
		Buffer:          make(chan *agent.Span, BUFFER_SIZE),
	}

	go func() {
		service.Poll(context.Background())
	}()

	return service
}

func (ag *AgentServiceImpl) Poll(ctx context.Context) {
	for {
		bufferSize := len(ag.Buffer)
		if bufferSize > 0 || bufferSize == BUFFER_SIZE {
			log.Printf("batch size %d", bufferSize)
			stream, err := ag.CollectorClient.Stream(ctx)
			if err != nil {
				log.Fatalf("Stream error %v", err)
			}

			for idx := 0; idx < bufferSize; idx++ {
				span := <-ag.Buffer
				log.Print(span)
				if err := stream.Send(span); err != nil {
					log.Fatalf("Send error %v", err)
				}
			}

			_, err = stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("CloseAndRecv err %v", err)
			}
		}

		time.Sleep(POLLING_INTERVAL)
	}
}

func (agentServiceImpl *AgentServiceImpl) Send(ctx context.Context, span *agent.Span) (*agent.Response, error) {
	// agentServiceImpl.CollectorClient.Send(ctx, span)
	agentServiceImpl.Buffer <- span
	return &agent.Response{}, nil
}
