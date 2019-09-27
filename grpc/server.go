package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server/grpc"
	tosui "gomicro-playground/proto"
	"log"
	"time"
)

type Tosui struct{}

func (g *Tosui) Hello(ctx context.Context, req *tosui.Request, rsp *tosui.Response) error {
	rsp.Code = 200
	log.Print(req.Name)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name("tosui.grpc"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()
	if err := tosui.RegisterTosuiHandler(service.Server(), new(Tosui)); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
