package main

import (
    "fmt"
    "github.com/awatercolorpen/nitro-playground/proto/common"
    "github.com/golang/protobuf/proto"
    "github.com/micro/go-micro/v2/config/source/file"
    "log"
)

var (
    configFilePath = fmt.Sprintf("./%v", file.DefaultPath)
)

type B struct{
    C string `json:"c"`
}

type A struct {
    B *B `json:"b"`
}

func main() {
    a := &common.Request{
        Name: "name",
        Query: "query",
    }

    b, _ := proto.Marshal(a)
    log.Print(string(b))
}
