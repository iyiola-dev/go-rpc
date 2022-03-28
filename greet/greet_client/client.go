package main

import (
	"context"
	"github.com/iyiola-dev/go-rpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)

}
func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{Greeting: &greetpb.Greeting{
		FirstName: "Iyiola",
		LastName:  "Akanbi",
	}}
	response, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call rpc function : %v", err)
	}
	log.Printf("response from rpc : %v", response)
}
