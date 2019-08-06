package main

import (
	"context"
	tosu "demo/proto"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
	"log"
)

type Tosu struct{}

func (g *Tosu) Hello(ctx context.Context, req *tosu.Request, rsp *tosu.Response) error {
	rsp.Code = 0
	log.Print(req.Name)
	return nil
}

func main() {
	service := k8s.NewService(
		micro.Name("tosu.yaml"),
		micro.Version("latest"),
	)

	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}