package main

import (
	"context"
	"log"

	pb "github.com/vinay03/grpc-go-course/calculator/proto"
)

func (s *Server) GetSum(ctx context.Context, sr *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("GetSum was invoked")
	return &pb.SumResponse{
		Result: sr.First + sr.Second,
	}, nil
}
