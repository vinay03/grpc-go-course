package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/vinay03/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*empty.Empty, error) {
	log.Printf("UpdateBlog was invoked with: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot pass ID",
		)
	}

	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{
			"$set": data,
		},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}
	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find Blog with ID",
		)
	}

	return &empty.Empty{}, nil
}
