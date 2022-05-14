package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ajithkumarsekar/grpc_with_go/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet method is invoked : %v\n", request)

	log.Printf("Checking for metadata")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		// by default this metadata will always be present
		//:authority:[localhost:9000]
		//content-type:[application/grpc]
		//user-agent:[grpc-go/1.46.2]
		log.Printf("No metadata present")
		return nil, fmt.Errorf("no metadata present")
	}

	for key, value := range md {
		log.Printf("metadata - %v:%v", key, value)
	}

	name := request.Greeting.GetFirstName() + " " + request.Greeting.GetLastName()
	result := "Hello " + name + "!"
	greetResponse := &greetpb.GreetResponse{Result: result}

	// create and send header
	header := metadata.Pairs("header-key", "val")
	// All the metadata will be sent out when one of the following happens:
	//- grpc.SendHeader() is called;
	//- The first response is sent out;
	//- An RPC status is sent out (error or success).
	grpc.SetHeader(ctx, md)      // sending the same metadata back
	grpc.SendHeader(ctx, header) // Sends header to client. this can be called at most once.

	// create and set trailer
	trailer := metadata.Pairs("trailer-key", "val")
	grpc.SetTrailer(ctx, trailer)

	return greetResponse, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	//reflection helps generic gRPC client(evans) to discover the Interface
	// definition language (proto)
	reflection.Register(s)

	greetpb.RegisterGreetServiceServer(s, &server{})

	log.Printf("starting grpc server on port 9000")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
