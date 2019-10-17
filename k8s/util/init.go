package util

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-plugins/config/source/configmap"
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

func init() {
	_ = config.Load(
		env.NewSource(env.WithStrippedPrefix("MICRO")),
		configmap.NewSource(
			configmap.WithNamespace(getNamespace()),
			configmap.WithName("micro")))

	_ = config.Get("file_log").Scan(GetRotateLogOption())
	_ = config.Get("elastic").Scan(GetElasticLogOption())
}
