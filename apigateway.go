package main

import (
	"flag"
	"fmt"
	"net/http"
	gw "wsserver/proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:51001", "endpoint of YourService")
)

func run() error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)

	if err != nil {
		return err
	}

	return http.ListenAndServe(":9090", mux)
}

func main() {
	if err := run(); err != nil {
		fmt.Print(err.Error())
	}

}