package main

import (
	"context"
	"github.com/google/uuid"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	_opentracing "github.com/opentracing/opentracing-go"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	tencho "gomicro-playground/k8s/proto"
	"gomicro-playground/k8s/util"
	"github.com/micro/go-micro/client"
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

func qaz(service micro.Service)  {
	c := cron.New()
	c.AddFunc("10 * * * *", func() {
		service := util.NewService(
			micro.Name("anst-akin-client"),
			micro.Version("latest"),
		)

		cli := tencho.NewAkinService("anst-akin", service.Client())
		rsp, err := cli.A(context.Background(), &tencho.Request{
			Name:                 uuid.New().String(),
			Query:                uuid.New().String(),
		})

		fields := log.Fields{
			"caller": "util anst-akin-client",
		}

		if err != nil {
			log.WithFields(fields).Error(err)
		} else {
			log.WithFields(fields).Info(rsp)
		}
	})

	c.AddFunc("10 * * * *", func() {
		cli := tencho.NewAkinService("anst-akin", service.Client())
		rsp, err := cli.A(context.Background(), &tencho.Request{
			Name:                 uuid.New().String(),
			Query:                uuid.New().String(),
		})

		fields := log.Fields{
			"caller": "util anst-akin",
		}

		if err != nil {
			log.WithFields(fields).Error(err)
		} else {
			log.WithFields(fields).Info(rsp)
		}
	})

	c.AddFunc("10 * * * *", func() {
		service := k8s.NewService(
			micro.Name("anst-tosui-client"),
			micro.Version("latest"),
		)

		cli := tencho.NewAkinService("anst-tosui", service.Client())
		rsp, err := cli.A(context.Background(), &tencho.Request{
			Name:                 uuid.New().String(),
			Query:                uuid.New().String(),
		})

		fields := log.Fields{
			"caller": "k8s anst-tosui-client",
		}

		if err != nil {
			log.WithFields(fields).Error(err)
		} else {
			log.WithFields(fields).Info(rsp)
		}
	})

	c.AddFunc("10 * * * *", func() {
		cli := tencho.NewAkinService("anst-tosui", service.Client())
		rsp, err := cli.A(context.Background(), &tencho.Request{
			Name:                 uuid.New().String(),
			Query:                uuid.New().String(),
		})

		fields := log.Fields{
			"caller": "util anst-tosui",
		}

		if err != nil {
			log.WithFields(fields).Error(err)
		} else {
			log.WithFields(fields).Info(rsp)
		}
	})

	c.AddFunc("10 * * * *", func() {
		req := client.NewRequest("anst-akin", "Akin.A", &tencho.Request{
			Name:                 uuid.New().String(),
			Query:                uuid.New().String(),
		})

		rsp := &tencho.Response{}

		err := client.Call(context.Background(), req, rsp)

		fields := log.Fields{
			"caller": "client anst-akin",
		}

		if err != nil {
			log.WithFields(fields).Error(err)
		} else {
			log.WithFields(fields).Info(rsp)
		}
	})

	c.AddFunc("10 * * * *", func() {
		req := client.NewRequest("anst-tosui", "ToSui.A", &tencho.Request{
			Name:                 uuid.New().String(),
			Query:                uuid.New().String(),
		})

		rsp := &tencho.Response{}

		err := client.Call(context.Background(), req, rsp)

		fields := log.Fields{
			"caller": "client anst-tosui",
		}

		if err != nil {
			log.WithFields(fields).Error(err)
		} else {
			log.WithFields(fields).Info(rsp)
		}
	})

	c.AddFunc("10 * * * *", func() {
		req := client.NewRequest("anst-tosui", "Akin.A", &tencho.Request{
			Name:                 uuid.New().String(),
			Query:                uuid.New().String(),
		})

		rsp := &tencho.Response{}

		err := client.Call(context.Background(), req, rsp)

		fields := log.Fields{
			"caller": "client anst-tosui Akin.A ?",
		}

		if err != nil {
			log.WithFields(fields).Error(err)
		} else {
			log.WithFields(fields).Info(rsp)
		}
	})

	c.Start()
}

func main() {
	util.Initlog()
	log.Info("anst-akin start")

	service := util.NewService(
		micro.Name("anst-akin"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapHandler(
			opentracing.NewHandlerWrapper(_opentracing.GlobalTracer()),
		),
	)

	service.Init()
	akin := new(Akin)
	akin.client = tencho.NewNatSuService("anst-natsu", service.Client())
	if err := tencho.RegisterAkinHandler(service.Server(), akin); err != nil {
		log.Fatal(err)
	}

	qaz()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
