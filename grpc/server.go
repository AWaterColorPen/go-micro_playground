package main

import (
	"context"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	tosui "go-micro_playground/proto"
	"log"
)

type Tosui struct{}

func (g *Tosui) Hello(ctx context.Context, req *tosui.Request, rsp *tosui.Response) error {
	rsp.Code = 0
	log.Print(req.Name)
	return nil
}

func main() {
	service := k8s.NewService(
		micro.Name("tosui"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
