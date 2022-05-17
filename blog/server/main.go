package main

import (
	"context"
	"log"
	"net"
	"time"

	blogpb "github.com/orlandorode97/grpc-golang-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	collection         *mongo.Collection
	address            = "0.0.0.0:8082"
	defaultGRPCTimeout = 30 * time.Second
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:iLV31gIaSz7GRgcV@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")
	lister, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("could not listen on %v\n", address)
	}

	defer lister.Close()

	log.Printf("listening on %v\n", address)

	s := grpc.NewServer()

	blogpb.RegisterBlogServiceServer(s, &GRPCServer{
		collection: collection,
	})

	if err = s.Serve(lister); err != nil {

		log.Fatalf("failed to serve %v", err)
	}

}
