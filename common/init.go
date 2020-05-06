package common

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/env"
	"github.com/micro/go-micro/v2/config/source/file"
	log "github.com/sirupsen/logrus"
)

var (
	configFilePath = fmt.Sprintf("/config/%v", file.DefaultPath)
)

func init() {
	_ = config.Load(env.NewSource())
}

func LoadConfig(configFile ...string) func() {
	return func() {
		p := configFilePath
		if len(configFile) > 0 {
			p = configFile[0]
		}

		if err := config.Load(file.NewSource(file.WithPath(p))); err != nil {
			fmt.Println(err)
		}
	}
}

func Config4Customer(kv map[string]interface{}) func() {
	return func() {
		log.Info(string(config.Bytes()))
		for k, v := range kv {
			_ = config.Get(k).Scan(v)
			log.Info(v)
		}
	}
}

func InitOption(opts ...func())  {
	for _, opt := range opts {
		opt()
	}
}