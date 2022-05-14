package main

import (
	"context"
	"fmt"
	"log"

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
	fmt.Printf("created a client %v\n", c)

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
