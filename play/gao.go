package main

import (
    "encoding/json"
    "log"
)

type A struct {
    V string `json:"v"`
    D string `json:"d"`
}

type B struct {
    A
    C string `json:"c"`
}

func main() {
    b := &B{
        C: "as",
        A: A{
            V: "c",
        },
    }

    s , _ := json.Marshal(b)
    log.Print(string(s))
}
