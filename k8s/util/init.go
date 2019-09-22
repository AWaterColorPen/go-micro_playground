package util

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/consul"
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

func init() {
	//configMapSource := configmap.NewSource(
	//	configmap.WithNamespace("go-micro"),
	//	configmap.WithName("micro"),
	//)

	consulSource := consul.NewSource(
		consul.WithAddress("consul.local"),
		configmap.WithName("/micro/config"),
	)

	_ = config.Load(consulSource)
	_ = config.Get("micro", "config", "file_log").Scan(GetRotateLogOption())
	_ = config.Get("micro", "config", "elastic").Scan(GetElasticLogOption())
}
