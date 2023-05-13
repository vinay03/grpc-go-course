package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vinay03/grpc-go-course/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	stream, err := c.Prime(context.Background(), &pb.PrimeRequest{
		Number: 120,
	})
	if err != nil {
		log.Fatalf("Error while calling Prime: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading stream output: %v\n", err)
		}
		log.Printf("Number Received: %v\n", res.Result)
	}
}
