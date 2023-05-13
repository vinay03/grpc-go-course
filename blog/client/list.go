package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vinay03/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("---listBlog was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling listblogs: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving data from server: %v", err)
			break
		}
		log.Println(res)
	}
}
