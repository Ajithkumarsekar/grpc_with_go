package main

import (
	"context"
	"log"
	"time"

	"github.com/ajithkumarsekar/grpc_with_go/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {

	cc, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	//greet(c)
	greetWithDeadline(c)
}

func greetWithDeadline(c greetpb.GreetServiceClient) {
	clientDeadLine := time.Now().Add(time.Second * 1)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadLine)
	defer cancel()

	// GreetWithDeadline is just a function name. it has nothing to
	// do with deadline handling
	resp, err := c.GreetWithDeadline(ctx, &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ajithkumar",
			LastName:  "sekar",
		},
	})
	if err != nil {
		log.Fatalf("error while calling GreetWithDeadline RPC : %v", err)
	}
	log.Printf("GreetWithDeadline response `%v`\n", resp.Result)
}

func greet(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ajithkumar",
			LastName:  "sekar",
		},
	}

	ctx := context.Background()

	// adding some metadata using different methods
	md := metadata.New(map[string]string{"Auth": "Bearer mysecretpassword"})
	// this would replace the existing metadata
	ctx = metadata.NewOutgoingContext(ctx, md)
	// appends to existing metadata
	ctx = metadata.AppendToOutgoingContext(ctx, "key1", "val1")

	// Make RPC using the context with the metadata.
	var header, trailer metadata.MD
	greet, err := c.Greet(ctx, req, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("error while calling greet RPC : %v", err)
	}
	log.Printf("received successfull response")
	log.Printf("checking gRCP response headers ...")
	for key, value := range header {
		log.Printf("header %v:%v", key, value)
	}

	log.Printf("checking gRCP response trailer metadata ...")
	for key, value := range trailer {
		log.Printf("trailer %v:%v", key, value)
	}

	log.Printf("Greet response `%v`\n", greet.Result)
}
