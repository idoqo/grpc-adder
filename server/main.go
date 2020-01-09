package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "github.com/idoqo/grpc-adder"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Compute(ctx context.Context, r *proto.AddRequest) (*proto.AddResponse, error) {
	result := &proto.AddResponse{}
	result.Result = r.A + r.B

	msg := fmt.Sprintf("A: %d B: %d sum: %d", r.A, r.B, result.Result)
	log.Println(msg)

	return result, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterAddServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
