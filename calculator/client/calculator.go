package main

import (
	"context"
	"io"
	"log"
	"time"

	calculatorpb "github.com/orlandorode97/grpc-golang-course/calculator/proto"
)

func doMax(c calculatorpb.CalculatorServiceClient) {
	log.Println("doGreetEveryone has been invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("error while creating strem %v\n", err)
	}

	reqs := []calculatorpb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	wait := make(chan struct{})

	go func() {
		for _, req := range reqs {
			stream.Send(&req)
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
				log.Printf("error while receiving %v", err)
			}

			log.Printf("receiving: %v", res.Max)
		}
		close(wait)
	}()

	<-wait

}
