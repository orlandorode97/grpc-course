package main

import (
	"log"
	"net"

	calculatorpb "github.com/orlandorode97/grpc-golang-course/calculator/proto"
	"google.golang.org/grpc"
)

var address = "0.0.0.0:8082"

func main() {
	lister, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("could not listen on %v\n", address)
	}

	defer lister.Close()

	log.Printf("listening on %v\n", address)

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lister); err != nil {
		log.Fatalf("failed to serve %v", err)
	}

}
