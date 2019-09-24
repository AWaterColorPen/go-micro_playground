FROM golang:latest

# micro
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get -u github.com/micro/micro
RUN go get -u github.com/micro/go-micro
RUN go get -u github.com/micro/go-plugins/
RUN go get -u github.com/micro/protobuf/protoc-gen-go
RUN go get -u github.com/micro/protoc-gen-micro

# k8s
RUN go get -u k8s.io/client-go@kubernetes-1.16.0
RUN go get -u k8s.io/api/core/v1
RUN go get -u k8s.io/apimachinery/pkg/apis/meta/v1
RUN go get -u k8s.io/apimachinery/pkg/fields

# linq
RUN go get -u github.com/ahmetb/go-linq

# log
RUN go get -u github.com/rifflock/lfshook
RUN go get -u github.com/lestrrat-go/file-rotatelogs

# elastic log
RUN go get -u github.com/olivere/elastic
RUN go get -u github.com/sohlich/elogrus

# protobuf
RUN go get -u github.com/golang/protobuf/jsonpb
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/golang/protobuf/ptypes

# grpc
RUN go get -u google.golang.org/grpc

# opentracing
RUN go get -u github.com/opentracing/opentracing-go

# prometheus
RUN go get -u github.com/prometheus/client_golang/prometheus

# micro config
RUN go get -u github.com/BurntSushi/toml
RUN go get -u github.com/fsnotify/fsnotify
RUN go get -u github.com/hashicorp/hcl
RUN go get -u github.com/imdario/mergo
RUN go get -u github.com/bitly/go-simplejson
RUN go get -u github.com/spf13/pflag

# support utf8
RUN apt-get update
RUN apt-get install -y locales
RUN sed -ie 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/g' /etc/locale.gen
RUN locale-gen
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV LANGUAGE en_US.UTF-8

WORKDIR $GOPATH