module github.com/awatercolorpen/nitro-playground

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/asim/nitro-plugins/client/grpc/v3 v3.4.0
	github.com/asim/nitro-plugins/registry/etcd/v3 v3.4.0
	github.com/asim/nitro-plugins/server/grpc/v3 v3.4.0
	github.com/asim/nitro/v3 v3.4.1
	github.com/golang/protobuf v1.4.3
)
