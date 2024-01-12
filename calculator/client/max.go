package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/flpcastro/grpc-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int{4, 7, 2, 19, 4, 6, 32}

		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			stream.Send(&pb.MaxRequest{
				Number: int32(number),
			})
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Problem while reading server stream: %v\n", err)
				break
			}

			log.Printf("Received a new maximum: %d\n", res.Result)
		}
	}()

	<-waitc
}
