package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ajithkumarsekar/grpc_with_go/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")

	cc, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
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

	greet, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC : %v", err)
	}

	log.Printf("Greet response `%v`\n", greet.Result)

}
