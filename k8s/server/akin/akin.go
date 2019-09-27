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

type Akin struct{
	client tencho.NatSuService
}

func (g *Akin) A(ctx context.Context, in *tencho.Request, out *tencho.Response) error {
	id := in.Name
	if id == "" {
		id = uuid.New().String()
	}

	fields := log.Fields{
		"serviceName": "Akin",
		"Id": id,
	}

	out.Code = 400
	log.WithFields(fields).Info(in)

	rsp, err := g.client.A(context.Background(), &tencho.Request{
		Name: id,
		Query: in.Query,
	})

	if err != nil {
		log.WithFields(fields).Error(err)
	} else {
		out.Code = rsp.Code
		log.WithFields(fields).Info(rsp)
	}

	return nil
}

func main() {
	util.Initlog()

	service := k8s.NewService(
		micro.Name("akin"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapHandler(
			opentracing.NewHandlerWrapper(_opentracing.GlobalTracer()),
		),
	)

	service.Init()
	akin := new(Akin)
	akin.client = tencho.NewNatSuService("natsu", service.Client())
	if err := tencho.RegisterAkinHandler(service.Server(), akin); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
