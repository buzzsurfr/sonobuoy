package server

import (
	"context"
	"log"
	"net"

	pb "github.com/buzzsurfr/sonobuoy/proto"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	pb.UnimplementedEchoServer
}

func (s *GrpcServer) Signal(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
	log.Print("Received ping")
	return &pb.Pong{}, nil
}

func (s *GrpcServer) Serve(lis net.Listener) error {
	srv := grpc.NewServer()
	pb.RegisterEchoServer(srv, s)
	return srv.Serve(lis)
}
