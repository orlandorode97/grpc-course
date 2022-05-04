package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/orlandorode97/grpc-golang-course/greet/proto"
)

type Server struct {
	pb.GreetServiceServer
}

// GRPC Unary type
func (s *Server) Greet(ctx context.Context, r *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", r)
	return &pb.GreetResponse{
		Result: "Hello " + r.FirstName,
	}, nil
}

// GRPC server streaming
func (s *Server) GreetManyTimes(r *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("Greetmany times was invoked with %v\n", r)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", r.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}

// GRPC client streaming
func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Long greet function invoked")
	res := ""
	for {
		// Get every client request
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("error while reading client streaming %v", err)
			return err
		}
		res += fmt.Sprintf("Hello %s\n", req.FirstName)
	}
}

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return nil
		}

		res := "Hello: " + req.FirstName + "!"
		if err := stream.Send(&pb.GreetResponse{
			Result: res,
		}); err != nil {
			log.Printf("error while sending data to client %v", err)
		}
	}
	return nil
}
