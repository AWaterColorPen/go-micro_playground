package main

import (
	"context"
	"flag"
	"fmt"

	api "go-micro_playground/proto/api"
	tosui "go-micro_playground/proto"
	"google.golang.org/grpc"
)

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:20001", "grpc grpc address")
)


func main() {
	conn, err := grpc.Dial(*endpoint, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	client := api.NewTosuiClient(conn)
	rsp, err := client.Hello(context.Background(), &tosui.Request{
		Name: "John",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
}
