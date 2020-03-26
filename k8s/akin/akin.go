package main

import (
	"context"
	"github.com/AWaterColorPen/go-micro_playground/common"
	tencho "github.com/AWaterColorPen/go-micro_playground/proto"
	"github.com/google/uuid"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server/grpc"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

type Akin struct{
	client tencho.ToSuiService
}

func (g *Akin) Call(ctx context.Context, in *tencho.Request, out *tencho.Response) error {
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

	rsp, err := g.client.Call(context.Background(), &tencho.Request{
		Name: id,
		Query: in.Query,
	}, )

	if err != nil {
		log.WithFields(fields).Error(err)
	} else {
		out.Code = rsp.Code
		log.WithFields(fields).Info(rsp)
	}

	return nil
}

func qaz(service micro.Service)  {
	c := cron.New()
	_, _ = c.AddFunc("0/10 * * * *", func() {
		cli := tencho.NewAkinService("anst-akin", service.Client())
		rsp, err := cli.Call(context.Background(), &tencho.Request{
			Name:  uuid.New().String(),
			Query: uuid.New().String(),
		})

		fields := log.Fields{
			"caller": "k8s anst-akin",
		}

		if err != nil {
			log.WithFields(fields).Error(err)
		} else {
			log.WithFields(fields).Info(rsp)
		}
	})

	c.Start()
}

func init() {
	grpc.DefaultMaxMsgSize = 50 * 1024 * 1024
	common.Init(map[string]interface{}{})
}

func main() {
	log.Info("anst-akin start")

	service := micro.NewService(
		micro.Name("anst-akin"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	service.Init()
	akin := &Akin{client: tencho.NewToSuiService("anst-tosui", service.Client())}
	if err := tencho.RegisterAkinHandler(service.Server(), akin); err != nil {
		log.Fatal(err)
	}

	qaz(service)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
