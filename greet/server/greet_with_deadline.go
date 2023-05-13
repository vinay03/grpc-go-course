package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vinay03/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked: %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client cancelled the request")
			return nil, status.Error(codes.Canceled, "the client cancelled the request")
		}

		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
