package main

import (
	"context"
	"log"

	pb "github.com/flpcastro/grpc-go/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("ReadBlog was invoked")

	req := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
