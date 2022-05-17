package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	primespb "github.com/orlandorode97/grpc-golang-course/primes/proto"
)

var addr = "0.0.0.0:50051"

func main() {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listing on %s\n", addr)
	s := grpc.NewServer()

	primespb.RegisterPrimeServiceServer(s, &server{})

	fmt.Println("Oke que la verga")

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
