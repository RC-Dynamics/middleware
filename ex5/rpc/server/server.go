package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	pb "../protobuff"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement stringmanipulation.Server.
type server struct{}

// Upper implements stringmanipulation.Server
func (s *server) Upper(ctx context.Context, in *pb.StrRequest) (*pb.StrReply, error) {
	time.Sleep(10 * time.Millisecond)
	return &pb.StrReply{Message: strings.ToUpper(in.Name)}, nil
}

func main() {
	log.Print("Starting Server...")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStringManipulationServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
