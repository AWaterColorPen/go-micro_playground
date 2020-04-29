package main

import "log"

var (
    b a
)

type a struct {
    c string
}

func main() {

    b.c = "s"
    log.Print(b)
}