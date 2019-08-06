package main

import (
	"context"
	tosu "demo/proto"
	"fmt"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("tosu.yaml"),
		micro.Version("latest"),
	)

	client := tosu.NewTosuService("tosu.yaml", service.Client())
	rsp, err := client.Hello(context.TODO(), &tosu.Request{
		Name: "John",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Code)
}