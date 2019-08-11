package main

import (
	"github.com/micro/go-micro/web"
	_ "go-micro_playground/proto"
	"log"
)

func main() {
	service := web.NewService(
		web.Name("tosui.web"),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
