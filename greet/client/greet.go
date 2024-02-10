package main

import (
	"context"
	"log"

	"github.com/kaungmyathan22/golang-grpc/greet/proto"
)

func doGreet(c proto.GreetServiceClient) {
	log.Println("doGreet was invoked")
	r, err := c.Greet(context.Background(), &proto.GreetRequest{FirstName: "Kaung Myat Han"})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", r.Result)
}
