package main

import (
	"context"
	tosui "go-micro_playground/proto"
	"fmt"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("tosui.yaml"),
		micro.Version("latest"),
	)

	client := tosui.NewTosuService("tosui.yaml", service.Client())
	rsp, err := client.Hello(context.TODO(), &tosu.Request{
		Name: "John",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Code)
}