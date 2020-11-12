# nitro_playground

## install

```shell
go get github.com/asim/protoc-gen-nitro/v3
```

compile protoc by `compile.sh` file

```shell
protoc --nitro_out=paths=source_relative:. --go_out=paths=source_relative:. your.proto
```

## example

server side

```go
    ## use grpc as base rpc service; grpc / http rpc
    nitro := grpc.NewService(
    ## set service name
        service.Name(common.SERVICE_NAME_AKIN.String()),
    ## use etcd as registry for service discovery; etcd / consul / nats / zookeeper / so on
        service.Registry(etcd.NewRegistry(registry.Addrs())),
    )

    ## just init
    nitro.Init()

    ## register a hander as server
    _ = akin.RegisterAkinHandler(nitro.Server(), &handler{})

    ## run
    if err := nitro.Run(); err != nil {
        log.Fatal(err)
    }
```

client side

```go
    nitro := grpc.NewService(
        service.Name(common.SERVICE_NAME_TOSUI.String()),
        service.Registry(etcd.NewRegistry(registry.Addrs())),
    )
    nitro.Init()

    ## new a client by service name
    srv := akin.NewAkinService(common.SERVICE_NAME_AKIN.String(), nitro.Client())

    ## rpc call
    log.Print(srv.Call(context.Background(), &common.Request{Name: "play ground"}))
```