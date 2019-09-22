package main

import (
	"context"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	_opentracing "github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	tencho "go-micro_playground/k8s/proto"
	"go-micro_playground/k8s/util"
)

type ToSui struct{}

func (g *ToSui) A(ctx context.Context, req *tencho.Request, rsp *tencho.Response) error {
	rsp.Code = 200
	log.WithFields(log.Fields{
		"serviceName": "ToSui",
	}).Info(req)
	return nil
}

func main() {
	util.Initlog()

	service := k8s.NewService(
		micro.Name("tosui"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapHandler(
			opentracing.NewHandlerWrapper(_opentracing.GlobalTracer()),
		),
	)

	service.Init()
	tencho.RegisterAkinHandler(service.Server(), new(ToSui))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
