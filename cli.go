package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "wsserver/proto"
)

const (
	address = "localhost:51001"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	msg := &pb.StringMessage{Value: "hello world"}
	result, err := client.Echo(context.Background(), msg)

	if err == nil {
		fmt.Print(result.Value)
	} else {
		fmt.Print(err.Error())
	}
}
