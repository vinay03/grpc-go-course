package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/vinay03/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlogs(in *empty.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("ListBlogs was invoked")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongodB: %v\n", err),
			)
		}
		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return nil
}
