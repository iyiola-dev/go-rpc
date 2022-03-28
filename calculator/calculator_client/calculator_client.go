package main

import (
	"context"
	"github.com/iyiola-dev/go-rpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}
func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.CalculatorRequest{Calculator: &calculatorpb.Calculator{
		FirstNumber: 3,
		LastNumber:  10,
	}}
	response, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call rpc function : %v", err)
	}
	log.Printf("response from rpc : %v", response)
}
