package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/service/grpc"
	tosui "gomicro-playground/proto"
)

func main() {
	service := grpc.NewService()
	service.Init()

	client := tosui.NewTosuiService("tosui.grpc", service.Client())
	rsp, err := client.Hello(context.Background(), &tosui.Request{
		Name: "John",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
}
