package main

import (
	"log"
	"net"

	pb "github.com/orlandorode97/grpc-golang-course/greet/proto"
	"google.golang.org/grpc"
)

var addr = "0.0.0.0:50051"

func main() {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listing on %s\n", addr)
	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
