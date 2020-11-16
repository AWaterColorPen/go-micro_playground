package main

import (
    "context"
    "log"

    gclient "github.com/asim/nitro-plugins/client/grpc/v3"
    "github.com/asim/nitro-plugins/registry/etcd/v3"
    gserver "github.com/asim/nitro-plugins/server/grpc/v3"
    "github.com/asim/nitro/v3/app"
    "github.com/asim/nitro/v3/app/rpc"
    "github.com/asim/nitro/v3/registry"
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
    nitro := rpc.NewApp(
        app.Client(gclient.NewClient()),
        app.Server(gserver.NewServer()),
        app.Registry(etcd.NewRegistry(registry.Addrs())),
    )

    nitro.Name(common.SERVICE_NAME_TOSUI.String())
    var rsp common.Response
    if err := nitro.Call(common.SERVICE_NAME_AKIN.String(), "Handler.Call", &common.Request{Name: "play ground"}, &rsp); err != nil {
        log.Fatal(err)
    }

    log.Print(rsp)

    _ = nitro.Handle(&handler{})
    if err := nitro.Run(); err != nil {
        log.Fatal(err)
    }
}