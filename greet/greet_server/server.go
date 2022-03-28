package main

import (
	"context"
	"github.com/iyiola-dev/go-rpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
}

func (s *Server) GreetManyTimes(request *greetpb.GreetManyTimesRequest, server greetpb.GreetService_GreetManyTimesServer) error {

}

func (s *Server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := request.GetGreeting().GetFirstName()
	result := "hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
