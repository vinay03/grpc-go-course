package main

import (
	"io"
	"log"

	pb "github.com/vinay03/grpc-go-course/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {

	log.Println("Max function was invoked")
	var max int32 = -1
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while receiving from client: %v\n", err)
			return nil
		}

		log.Printf("Data received: %v\n", req.Number)
		if max < int32(req.Number) {
			max = int32(req.Number)
		}
		err = stream.Send(&pb.NumberResponse{
			Result: uint32(max),
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
