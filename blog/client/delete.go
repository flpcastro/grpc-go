package main

import (
	"context"
	"log"

	pb "github.com/flpcastro/grpc-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("DeleteBlog was invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}
