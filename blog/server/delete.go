package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/vinay03/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*empty.Empty, error) {
	log.Printf("deleteBlog was invoked with: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot parse ID",
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog was not found",
		)
	}
	return &empty.Empty{}, nil
}
