package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rakesh1194/grpc-go/calculator/sumpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("I am a client....")

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not able to connect: %v", err)
	}

	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)

	doUnary(c)
}

func doUnary(c sumpb.SumServiceClient) {
	fmt.Println("Starting to do Unary RPC...")
	req := &sumpb.SumRequest{
		Sum: &sumpb.Sum{
			A: 10,
			B: 3,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greet RPC: %v", err)
	}
	log.Printf("Response from greet %v", res.Result)
}
