package common

import (
	"fmt"
	"github.com/AWaterColorPen/go-micro_playground/common/logger"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-plugins/config/source/configmap"
	log "github.com/sirupsen/logrus"
	"net"
)

func currentIpAddress() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addr {
		if ip, ok := address.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}

	return ""
}

func getNamespace() string {
	conf := config.NewConfig()
	_ = conf.Load(env.NewSource())
	return conf.Get("NAMESPACE").String("default")
}

func initConfig()  {
	if err := config.Load(
		env.NewSource(env.WithStrippedPrefix("MICRO")),
		file.NewSource(file.WithPath(fmt.Sprintf("/config/%v", file.DefaultPath))),
			configmap.NewSource(
				configmap.WithNamespace(getNamespace()),
				configmap.WithName("micro")),
	); err != nil {
		fmt.Println(err)
	}

	_ = config.Get("file_log").Scan(&logger.RotateLog)
	_ = config.Get("elastic").Scan(&logger.ElasticLog)
}

func init() {
	initConfig()
	logger.Init(logger.Metadata{Ip:currentIpAddress()})
}

func Init(kv map[string]interface{}) {
	log.Info(config.Map())

	for k, v := range kv {
		_ = config.Get(k).Scan(v)
		log.Info(v)
	}
}