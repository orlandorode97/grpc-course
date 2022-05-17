package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/orlandorode97/grpc-golang-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {

    opts:= []grpc.DialOption{}
    isTLS:= true
    if isTLS {
        certFile:= "ssl/ca.rt"
        creds, err:= credentials.NewClientTLSFromFile(certFile, "")

        if err!= nil {
            log.Fatalf("Error while loading CA trust certificates: %v\n", err)
        }

        opts = append(opts, grpc.WithTransportCredentials(creds))
    }


	conn, err := grpc.Dial(addr, opts...)
    if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	doGreetEveryone(c)
}
