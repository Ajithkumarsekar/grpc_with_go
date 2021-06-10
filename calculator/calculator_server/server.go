package main

import (
	"context"
	"fmt"
	"github.com/ajithkumarsekar/grpc_go_course/calculator/calculator_pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s server) SumNums(_ context.Context, request *calculator_pb.SumRequest) (*calculator_pb.SumResponse, error) {
	fmt.Printf("SumNums method is invoked : %v\n", request)
	sumOf2Nums := request.SumIt.GetNum1() + request.SumIt.GetNum2()

	response := &calculator_pb.SumResponse{
		Result: sumOf2Nums,
	}

	return response, nil
}

func main() {
	fmt.Println("Hello from calculator server")
	listen, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	calculator_pb.RegisterCalculatorServiceServer(s, &server{})

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
