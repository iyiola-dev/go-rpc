package main

import (
	"context"
	"fmt"
	"github.com/iyiola-dev/go-rpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
}

func (s Server) Calculator(ctx context.Context, request *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("received from rpc client: %v", request)
	firstNumber := request.GetCalculator().GetFirstNumber()
	secondNumber := request.GetCalculator().GetLastNumber()
	sum := firstNumber + secondNumber
	res := &calculatorpb.CalculatorResponse{
		Sum: sum,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	server := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(server, &Server{})
	lisErr := server.Serve(lis)
	if lisErr != nil {
		log.Fatalf("failed to listen : %v", err)
	}

}
