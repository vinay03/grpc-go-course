package main

import (
	"log"

	pb "github.com/vinay03/grpc-go-course/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Println("Prime function was invoked")
	var k uint32 = 2
	N := in.Number
	for N > 1 {
		if N%k == 0 { // if k evenly divides into N
			stream.Send(&pb.PrimeResponse{
				Result: uint32(k),
			}) // this is a factor
			N = N / k // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}
	}
	return nil
}
