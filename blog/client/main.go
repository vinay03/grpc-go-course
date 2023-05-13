package main

import (
	"log"

	pb "github.com/vinay03/grpc-go-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start connection: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id) //valid
	// readBlog(c, "anonexistingId") // invalid

	updateBlog(c, id)

	listBlog(c)

	deleteBlog(c, id)

	readBlog(c, id) //valid
}
