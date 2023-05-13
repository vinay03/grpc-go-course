package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/vinay03/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax function was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while fetching stream: %v\n", err)
	}

	watic := make(chan struct{})

	reqs := []*pb.NumberRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data from server: %v\n", err)
				break
			}
			log.Printf("Received data: %v\n", res.Result)
		}
		close(watic)
	}()

	<-watic

}
