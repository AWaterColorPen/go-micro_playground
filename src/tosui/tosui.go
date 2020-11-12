package main

import (
    "context"
    "log"

    "github.com/asim/nitro-plugins/registry/etcd/v3"
    "github.com/asim/nitro-plugins/service/grpc/v3"
    "github.com/asim/nitro/v3/registry"
    "github.com/asim/nitro/v3/service"
    "github.com/awatercolorpen/nitro-playground/proto/akin"
    "github.com/awatercolorpen/nitro-playground/proto/common"
    "github.com/awatercolorpen/nitro-playground/proto/tosui"
)

type handler struct {
}

func (h *handler) Call(ctx context.Context, in *common.Request, out *common.Response) error {
    log.Print(in)
    out.Code = 200
    return nil
}

func main() {
    nitro := grpc.NewService(
        service.Name(common.SERVICE_NAME_TOSUI.String()),
        service.Registry(etcd.NewRegistry(registry.Addrs())),
    )
    nitro.Init()

    srv := akin.NewAkinService(common.SERVICE_NAME_AKIN.String(), nitro.Client())
    log.Print(srv.Call(context.Background(), &common.Request{Name: "play ground"}))

    _ = tosui.RegisterToSuiHandler(nitro.Server(), &handler{})
    if err := nitro.Run(); err != nil {
        log.Fatal(err)
    }
}