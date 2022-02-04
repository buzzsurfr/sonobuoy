package server

import (
	"context"
	"log"

	pb "github.com/buzzsurfr/sonobuoy/proto"
)

type Server struct {
	pb.UnimplementedEchoServer
}

func (s *Server) Signal(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
	log.Print("Received ping")
	return &pb.Pong{}, nil
}
