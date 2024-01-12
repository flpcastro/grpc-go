package main

import (
	"context"
	"log"

	pb "github.com/flpcastro/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("UpdateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Felipe",
		Title:    "New Title",
		Content:  "New Content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
