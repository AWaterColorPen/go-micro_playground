package main

import (
	"context"
	"github.com/micro/go-micro"
	tosui "gomicro-playground/proto"
	"log"
)

type Tosui struct{}

func (g *Tosui) Hello(ctx context.Context, req *tosui.Request, rsp *tosui.Response) error {
	rsp.Code = 200
	log.Print(req.Name)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("tosui"),
		micro.Version("latest"),
	)

	service.Init()
	tosui.RegisterTosuiHandler(service.Server(), new(Tosui))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
