package util

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/client/selector/static"
	"github.com/micro/go-plugins/registry/kubernetes"
)

// NewService returns a new go-micro service pre-initialised for k8s
func NewService(opts ...micro.Option) micro.Service {
	// create registry and selector
	r := kubernetes.NewRegistry()
	s := static.NewSelector()

	// set the registry and selector
	options := []micro.Option{
		micro.Registry(r),
		micro.Selector(s),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return micro.NewService(options...)
}