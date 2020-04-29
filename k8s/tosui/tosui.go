package main

import (
	"context"
	"github.com/AWaterColorPen/go-micro_playground/common"
	"github.com/AWaterColorPen/go-micro_playground/proto"
	"github.com/google/uuid"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	log "github.com/sirupsen/logrus"
)

type ToSui struct{}

func (g *ToSui) Call(ctx context.Context, in *tencho.Request, out *tencho.Response) error {
	id := in.Name
	if id == "" {
		id = uuid.New().String()
	}

	fields := log.Fields{
		"serviceName": "ToSui",
		"Id": id,
	}

	out.Code = 400
	log.WithFields(fields).Info(in)
	out.Code = 200

	return nil

}

func init() {
	common.Init(map[string]interface{}{})
}

func main() {
	log.Info("anst-tosui start")

	srv := micro.NewService(
		micro.Name("anst-tosui"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	srv.Init()
	if err := tencho.RegisterToSuiHandler(srv.Server(), &ToSui{}); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
