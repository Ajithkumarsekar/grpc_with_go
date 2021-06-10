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

	//doUnary(c)

	//doServerStreaming(c)

	doClientStreaming(c)
}

func doClientStreaming(c calculator_pb.CalculatorServiceClient) {
	fmt.Println("Starting to do a ComputeAverage Client Streaming RPC...")
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	allNums := []int64{2, 4, 6, 9, 12, 2014}

	for _, num := range allNums {
		fmt.Printf("Sending number: %v\n", num)
		err := stream.Send(&calculator_pb.ComputeAverageRequest{
			Num: num,
		})
		if err != nil {
			log.Fatalf("Error while sending nums %v", err)
		}
	}

	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	fmt.Printf("The Average is: %v\n", result.GetResult())

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
			fmt.Printf("reached the stream end. Stopping it...\n")
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
