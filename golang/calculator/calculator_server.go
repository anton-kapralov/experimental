package main

import (
	"context"
	pb "github.com/anton-kapralov/experimental/proto/calculator"
	"net/http"
)

type CalculatorService struct{}

func (s *CalculatorService) Sum(ctx context.Context, request *pb.SumRequest) (response *pb.SumResponse, err error) {
	var sum int32
	for _, addend := range request.Addends {
		sum += addend
	}

	return &pb.SumResponse{Sum: sum}, nil
}

func main() {
	server := &CalculatorService{} // implements Haberdasher interface
	twirpHandler := pb.NewCalculatorServiceServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
