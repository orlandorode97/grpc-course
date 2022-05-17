package main

import (
	"log"
	"net"

	pb "github.com/orlandorode97/grpc-golang-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr = "0.0.0.0:50051"

func main() {
	listener, err := net.Listen("tcp", addr)

    opts:= []grpc.ServerOption{}
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listing on %s\n", addr)

    isTLS:= true

    if isTLS {
        certFile:= "ssl/server.crt"
        keyFile:= "ssl/server.perm"
        creds, err:= credentials.NewServerTLSFromFile(certFile, keyFile)
        if err!= nil {
            log.Fatalf("Failed to create a server TLS from file %v\n", err)
        }

        opts = append(opts, grpc.Creds(creds))
    }

	s := grpc.NewServer(opts...)

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
