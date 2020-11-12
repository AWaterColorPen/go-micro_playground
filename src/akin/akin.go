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
        service.Name("akin"),
        service.Registry(etcd.NewRegistry(registry.Addrs())),
    )
    nitro.Init()

    _ = akin.RegisterAkinHandler(nitro.Server(), &handler{})
    if err := nitro.Run(); err != nil {
        log.Fatal(err)
    }
}