package main

import (
	"context"
	"fmt"
	"github.com/ajithkumarsekar/grpc_go_course/calculator/calculator_pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello from calculator client")

	cc, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatalf("Error occured while closing client connection : %v", err)
		}
	}(cc)

	c := calculator_pb.NewCalculatorServiceClient(cc)

	doUnary(c)

	doServerStreaming(c)
}

func doServerStreaming(c calculator_pb.CalculatorServiceClient) {
	number := int64(83276)
	req := &calculator_pb.DecomposeNumberRequest{Num: number}
	stream, err := c.DecomposeToPrimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling DecomposeToPrimes RPC : %v", err)
	}

	for {
		resultStream, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("reached the stream end. Stopping it...")
			break
		}
		if err != nil {
			log.Fatalf("error occured while receiving data from stream : %v", err)
		}
		fmt.Printf("one of the prime factor for %v is %v\n", number, resultStream.PrimeFactor)
	}
}

func doUnary(c calculator_pb.CalculatorServiceClient) {
	req := &calculator_pb.SumRequest{
		SumIt: &calculator_pb.Sum{
			Num1: 932,
			Num2: 3598,
		},
	}

	sumResult, err := c.SumNums(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling sum RPC : %v", err)
	}

	log.Printf("Sum of nums : %v ", sumResult.Result)
}
