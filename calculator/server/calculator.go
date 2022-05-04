package main

import (
	"context"
	"io"
	"log"

	calculatorpb "github.com/orlandorode97/grpc-golang-course/calculator/proto"
)

type Server struct {
	calculatorpb.CalculatorServiceServer
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
