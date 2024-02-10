package main

import (
	"context"
	"log"

	"github.com/kaungmyathan22/golang-grpc/greet/proto"
)

func (*Server) Greet(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("Greet was invoked with %v\n", in)
	return &proto.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
