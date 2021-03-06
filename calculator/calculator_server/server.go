package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/ajithkumarsekar/grpc_with_go/calculator/calculator_pb"
	"google.golang.org/grpc"
)

type server struct{}

func (s server) FindMaximum(streamRequest calculator_pb.CalculatorService_FindMaximumServer) error {
	fmt.Println("Received FindMaximum RPC")
	maxNum := int64(0)

	for {
		recv, err := streamRequest.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		fmt.Printf("Received a new num : %v\n", recv.GetNum())
		if recv.GetNum() > maxNum {
			maxNum = recv.GetNum()
			err := streamRequest.SendMsg(&calculator_pb.FindMaximumResponse{
				MaxNum: maxNum,
			})
			if err != nil {
				log.Fatalf("Error occured while sending maxNum : %v", err)
				return err
			}
		}
	}

}

func (s server) ComputeAverage(streamRequest calculator_pb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("Received ComputeAverage RPC\n")

	sum := int64(0)
	n := 0

	for {
		num, err := streamRequest.Recv()
		if err == io.EOF {
			result := float64(sum) / float64(n)
			return streamRequest.SendAndClose(
				&calculator_pb.ComputeAverageResponse{
					Result: result,
				},
			)
		}
		if err != nil {
			log.Fatalf("Error occured while receiving nums %v", err)
		}

		sum += num.GetNum()
		n++

	}

	return nil
}

func (s server) DecomposeToPrimes(request *calculator_pb.DecomposeNumberRequest, stream calculator_pb.CalculatorService_DecomposeToPrimesServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v\n", request)

	n := request.Num
	k := int64(2)
	for n > 1 {
		if n%k == 0 {
			err := stream.Send(
				&calculator_pb.DecomposedNumbersResponse{
					PrimeFactor: k,
				},
			)
			if err != nil {
				log.Fatalf("Error occured while sending prime factors to client : %v", err)
			}
			n = n / k
		} else {
			k++
			fmt.Printf("Divisor has increased to %v\n", k)
		}
	}

	return nil

}

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
