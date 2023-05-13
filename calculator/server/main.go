package main

import (
	"log"
	"net"

	pb "github.com/vinay03/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
