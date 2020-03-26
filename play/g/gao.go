package main

import (
    "context"
    "github.com/AWaterColorPen/go_micro_playground/proto"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
    pr "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/push"
    "log"
)

type Say struct{}

func (s *Say) Call(ctx context.Context, req *tencho.Request, rsp *tencho.Response) error {
    log.Print(req)
    rsp.Code = 200
    return nil
}

func main() {
    service := micro.NewService(
        micro.Name("greeter"),
        micro.Version("fuck"),
        micro.WrapHandler(
            prometheus.NewHandlerWrapper(),),
    )

    push.New("http://pushgateway.domain.com/metrics", "Gao").Gatherer(pr.DefaultGatherer)

    service.Init()

    _ = tencho.RegisterAkinHandler(service.Server(), new(Say))
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}