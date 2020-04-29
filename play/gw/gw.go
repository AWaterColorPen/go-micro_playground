package main

import (
    "context"
    "flag"
    "log"
    "net/http"

    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "google.golang.org/grpc"

    "github.com/AWaterColorPen/go-micro_playground/proto"
)

var (
    echoEndpoint = flag.String("echo_endpoint", "localhost:57278", "endpoint of YourService")
)

func run() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := tencho.RegisterAkinHandlerFromEndpoint(ctx, mux, "localhost:62699", opts)
    if err != nil {
        return err
    }

    return http.ListenAndServe(":8080", mux)
}

func main() {
    flag.Parse()
    if err := run(); err != nil {
        log.Fatal(err)
    }
}