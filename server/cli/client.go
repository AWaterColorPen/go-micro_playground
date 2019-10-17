package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	tosui "gomicro-playground/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("tosui.client"),
		micro.Version("latest"),
	)

	service.Init()
	client := tosui.NewTosuiService("tosui", service.Client())
	rsp, err := client.Hello(context.Background(), &tosui.Request{
		Name: "John",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
}
