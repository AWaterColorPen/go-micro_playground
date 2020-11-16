package main

import (
    "context"
    "github.com/asim/nitro-plugins/registry/etcd/v3"
    "github.com/asim/nitro/v3/registry"
    "log"

    gclient "github.com/asim/nitro-plugins/client/grpc/v3"
    gserver "github.com/asim/nitro-plugins/server/grpc/v3"
    "github.com/asim/nitro/v3/app"
    "github.com/asim/nitro/v3/app/rpc"
    "github.com/awatercolorpen/nitro-playground/proto/common"
)

type Handler struct {
}

func (h *Handler) Call(ctx context.Context, in *common.Request, out *common.Response) error {
    log.Print(in)
    out.Code = 200
    return nil
}

func main() {
    nitro := rpc.NewApp(
        app.Client(gclient.NewClient()),
        app.Server(gserver.NewServer()),
        app.Registry(etcd.NewRegistry(registry.Addrs())),
        )

    nitro.Name(common.SERVICE_NAME_AKIN.String())
    _ = nitro.Handle(&Handler{})
    if err := nitro.Run(); err != nil {
        log.Fatal(err)
    }
}