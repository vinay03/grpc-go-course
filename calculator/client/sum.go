package main

import (
	"context"
	"log"

	pb "github.com/vinay03/grpc-go-course/calculator/proto"
)

func doSum(cl pb.CalculatorServiceClient) {
	res, err := cl.GetSum(context.Background(), &pb.SumRequest{
		First:  3,
		Second: 10,
	})
	if err != nil {
		log.Fatalf("Failed to call gRPC")
	}
	log.Println("GetSum result is : ", res.Result)
}
