package main

import (
	primespb "github.com/orlandorode97/grpc-golang-course/primes/proto"
)

type server struct {
	primespb.PrimeServiceServer
}

func (s *server) DescomposeNumber(r *primespb.PrimeRequest, strem primespb.PrimeService_DescomposeNumberServer) error {
	var divider int32 = 2
	number := r.Integer

	for number > 1 {
		if number%divider != 0 {
			divider++
			continue
		}
		number = number / divider

		strem.Send(&primespb.PrimeRespose{
			Result: divider,
		})
	}

	return nil
}
