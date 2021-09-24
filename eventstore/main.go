package eventstore

import (
	"context"

	"github.com/ldenholm/porus/data"
	"github.com/ldenholm/porus/pb"
)

type Server struct {
}

func (s *Server) CreateEvent(ctx context.Context, in *pb.Event) (*pb.Response, error) {
	// Persist data into EventStore database
	command := data.EventStore{}
	// Persist events as immutable logs into CockroachDB
	err := command.CreateEvent()
	if err != nil {
		return nil, err
	}
	// Publish event on NATS Streaming Server
	go publishEvent(s.StreamingComponent, in)
	return &pb.Response{IsSuccess: true}, nil
}
