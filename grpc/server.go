package main

import (
	"context"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"time"

	tosui "go-micro_playground/proto"
	"log"
)

type Tosui struct{}

func (g *Tosui) Hello(ctx context.Context, req *tosui.Request, rsp *tosui.Response) error {
	rsp.Code = 200
	log.Print(req.Name)
	return nil
}

func main() {
	service := grpc.NewService(
		micro.Name("tosui"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	service.Init()
	tosui.RegisterTosuiHandler(service.Server(), new(Tosui))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
