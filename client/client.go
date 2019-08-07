package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	tosui "go-micro_playground/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("tosui.yaml"),
		micro.Version("latest"),
	)

	client := tosui.NewTosuiService("tosui", service.Client())
	rsp, err := client.Hello(context.TODO(), &tosui.Request{
		Name: "John",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Code)
}
