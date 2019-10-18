package main

import (
	"context"
	"github.com/google/uuid"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server/grpc"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	_opentracing "github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	tencho "gomicro-playground/k8s/proto"
	"gomicro-playground/k8s/util"
	"sync"
)

type NatSu struct{
	clientS tencho.ShunMuService
	clientT tencho.ToSuiService
}

func (g *NatSu) A(ctx context.Context, in *tencho.Request, out *tencho.Response) error {
	id := in.Name
	if id == "" {
		id = uuid.New().String()
	}

	fields := log.Fields{
		"serviceName": "NatSu",
		"Id": id,
	}

	out.Code = 400
	log.WithFields(fields).Info(in)
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		rsp, err := g.clientS.A(context.Background(), &tencho.Request{
			Name: id,
			Query: in.Query,
		})

		if err != nil {
			log.WithFields(fields).Error(err)
		} else {
			log.WithFields(fields).Info(rsp)
		}

		wg.Done()
	}()

	go func() {
		wg.Add(1)
		rsp, err := g.clientT.A(context.Background(), &tencho.Request{
			Name: id,
			Query: in.Query,
		})

		if err != nil {
			log.WithFields(fields).Error(err)
		} else {
			log.WithFields(fields).Info(rsp)
		}

		wg.Done()
	}()

	wg.Wait()
	out.Code = 200

	return nil

}

func init() {
	grpc.DefaultMaxMsgSize = 50 * 1024 * 1024
}

func main() {
	util.Initlog()
	log.Info("anst-natsu start")

	service := k8s.NewService(
		micro.Name("anst-natsu"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapHandler(
			opentracing.NewHandlerWrapper(_opentracing.GlobalTracer()),
		),
	)

	service.Init()
	natsu := new(NatSu)
	natsu.clientS = tencho.NewShunMuService("anst-shunmu", service.Client())
	natsu.clientT = tencho.NewToSuiService("anst-tosui", service.Client())
	if err := tencho.RegisterNatSuHandler(service.Server(), natsu); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
