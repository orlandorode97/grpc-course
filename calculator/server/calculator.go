package main

import (
	"context"
	"io"
	"log"
	"math"

	calculatorpb "github.com/orlandorode97/grpc-golang-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (s *Server) Sum(ctx context.Context, r *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	num1, num2 := r.GetNumber_1(), r.GetNumber_2()
	return &calculatorpb.SumResponse{
		Sum: num1 + num2,
	}, nil
}

func (s *Server) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	sum, counter := 0, 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Average: float64(sum) / float64(counter),
			})
		}
		if err != nil {
			log.Fatalf("error while reading client streaming %v ", err)
		}
		counter++
		sum += int(req.Number)
	}
}

func (s *Server) Max(stream calculatorpb.CalculatorService_MaxServer) error {
	var max int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if max < req.Number {
			max = req.Number

			if err := stream.Send(&calculatorpb.MaxResponse{
				Max: max,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Server) Sqrt(ctx context.Context, r *calculatorpb.SqrtRequest) (*calculatorpb.SqrtResponse, error) {
	number := r.GetNumber()
	if number < 0 {
		return nil, status.Error(codes.InvalidArgument, "number cannot be less than zero")
	}
	return &calculatorpb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
