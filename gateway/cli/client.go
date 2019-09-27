package main

import (
	"context"
	"flag"
	"fmt"
	tosui "gomicro-playground/proto"
	api "gomicro-playground/proto/api"
	"google.golang.org/grpc"
)

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:20001", "grpc address")
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
