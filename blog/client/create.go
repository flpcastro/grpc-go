package main

import (
	"context"
	"log"

	pb "github.com/flpcastro/grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("CreateBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Felipe",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", res.Id)
	return res.Id
}
