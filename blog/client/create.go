package main

import (
	"context"
	"log"

	pb "github.com/vinay03/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked")
	blog := &pb.Blog{
		AuthorId: "Vinay",
		Title:    "My first Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("unexpected error: %v\n", err)
	}
	log.Printf("blog has been created: %v\n", res.Id)
	return res.Id
}
