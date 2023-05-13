package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vinay03/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error calling Avg function: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v\n", err)
	}

	log.Printf("Avg Response: %v\n", res.Result)

}
