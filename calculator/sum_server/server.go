package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rakesh1194/grpc-go/calculator/sumpb"
	"google.golang.org/grpc"
)

type server struct {
	sumpb.UnimplementedSumServiceServer
}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked %v", req)
	a := req.GetSum().GetA()
	b := req.GetSum().GetB()
	result := a + b

	res := &sumpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Staring Server....")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
