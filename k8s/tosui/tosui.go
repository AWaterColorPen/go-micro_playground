package main

import (
	"context"
	"github.com/AWaterColorPen/go-micro_playground/common"
	tencho "github.com/AWaterColorPen/go-micro_playground/proto"
	"github.com/google/uuid"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server/grpc"
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
	grpc.DefaultMaxMsgSize = 50 * 1024 * 1024
	common.Init(map[string]interface{}{})
}

func main() {
	log.Info("anst-tosui start")

	service := k8s.NewService(
		micro.Name("anst-tosui"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	service.Init()
	if err := tencho.RegisterToSuiHandler(service.Server(), &ToSui{}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
