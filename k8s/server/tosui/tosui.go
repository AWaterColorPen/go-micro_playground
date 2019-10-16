package main

import (
	"context"
	"github.com/google/uuid"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	_opentracing "github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	tencho "gomicro-playground/k8s/proto"
	"gomicro-playground/k8s/util"
)

type ToSui struct{}

func (g *ToSui) A(ctx context.Context, in *tencho.Request, out *tencho.Response) error {
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

func main() {
	util.Initlog()
	log.Info("anst-tosui start")

	service := k8s.NewService(
		micro.Name("anst-tosui"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapHandler(
			opentracing.NewHandlerWrapper(_opentracing.GlobalTracer()),
		),
	)

	service.Init()
	if err := tencho.RegisterToSuiHandler(service.Server(), new(ToSui)); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
