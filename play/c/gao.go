package main

import (
    "context"
    "github.com/AWaterColorPen/go_micro_playground/proto"
    "google.golang.org/grpc"
    "log"
)

func main() {
    // service := micro.NewService(
    //     micro.Name("greeter"),
    //     micro.Version("fuck"),
    // )
    // service.Init()
    //
    // c := tencho.NewAkinService("greeter", service.Client())
    // r , err := c.Call(context.Background(), &tencho.Request{Name:"1"})
    // if err != nil {
    //     // log.Fatal(err)
    // }
    //
    // log.Print(r)

    conn, err := grpc.Dial("localhost:59406", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    cli := tencho.NewAkinClient(conn)
    cr , err := cli.Call(context.Background(), &tencho.Request{Name:"1"})
    if err != nil {
        log.Fatal(err)
    }
    log.Print(cr)
}