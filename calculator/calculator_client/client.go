package main

import (
	"context"
	"fmt"
	"github.com/ajithkumarsekar/grpc_go_course/calculator/calculator_pb"
	"google.golang.org/grpc"
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

	req := &calculator_pb.SumRequest{
		SumIt: &calculator_pb.Sum{
			Num1: 932,
			Num2: 3598,
		},
	}

	sumResult, err := c.SumNums(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC : %v", err)
	}

	log.Printf("Sum of nums : %v ", sumResult.Result)

}
