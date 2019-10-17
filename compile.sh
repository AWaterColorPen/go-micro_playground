#!/usr/bin bash
#protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=plugins=grpc:. --go_out=. proto/shunmu.proto

#protoc -I/usr/local/include -I. \
#  -I$GOPATH/src \
#  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#  --go_out=plugins=grpc:. \
#  proto/api/api.protodoc
#protoc -I/usr/local/include -I. \
#  -I$GOPATH/src \
#  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#  --grpc-gateway_out=logtostderr=true:. \
#  proto/api/api.proto

protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. \
    k8s/proto/common.proto \
    k8s/proto/akin.proto \
    k8s/proto/natsu.proto \
    k8s/proto/shunmu.proto \
    k8s/proto/tosui.proto
