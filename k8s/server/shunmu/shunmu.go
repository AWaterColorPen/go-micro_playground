package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	_opentracing "github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	tencho "gomicro-playground/k8s/proto"
	"gomicro-playground/k8s/util"
)

type ShunMu struct{}

func (g *ShunMu) A(ctx context.Context, in *tencho.Request, out *tencho.Response) error {
	id := in.Name
	if id == "" {
		id = uuid.New().String()
	}

	fields := log.Fields{
		"serviceName": "ShunMu",
		"Id": id,
	}

	out.Code = 400
	log.WithFields(fields).Info(in)
	out.Code = 200

	return nil

}

func main() {
	util.Initlog()
	log.Info("anst-shunmu start")

	service := util.NewService(
		micro.Name("anst-shunmu"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapHandler(
			opentracing.NewHandlerWrapper(_opentracing.GlobalTracer()),
		),
	)

	service.Init()
	if err := tencho.RegisterAkinHandler(service.Server(), new(ShunMu)); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
