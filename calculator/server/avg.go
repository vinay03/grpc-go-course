package main

import (
	"io"
	"log"

	pb "github.com/vinay03/grpc-go-course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	var res float32 = 0
	var cnt float32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: res / cnt,
			})
		}

		if err != nil {
			log.Fatalf("Error receiving stream from client; %v\n", err)
		}
		cnt++
		res += float32(req.Number)

		log.Printf("Recevining: %v\n", req)
	}
}
