package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/orlandorode97/grpc-golang-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet has invoked")
	resp, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Orlando Romo Delgado",
	})
	if err != nil {
	}

	log.Printf("Greeting: %v\n", resp.Result)

}

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone has been invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while creating strem %v\n", err)
	}

	reqs := []pb.GreetRequest{
		{FirstName: "Orlando"},
		{FirstName: "Jpsd"},
		{FirstName: "pRRO"},
		{FirstName: "cuLO SUCIO"},
		{FirstName: "pEDRITO"},
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

			log.Printf("receiving: %v", res.Result)
		}
		close(wait)
	}()

	<-wait

}
