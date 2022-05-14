package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ajithkumarsekar/grpc_with_go/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(_ context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet method is invoked : %v\n", request)
	name := request.Greeting.GetFirstName() + " " + request.Greeting.GetLastName()
	result := "Hello " + name
	greetResponse := &greetpb.GreetResponse{Result: result}

	return greetResponse, nil
}

func main() {
	fmt.Println("Hello World!")
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
