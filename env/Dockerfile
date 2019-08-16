FROM golang:latest

# micro
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get -u github.com/micro/micro
RUN go get -u github.com/micro/go-micro
RUN go get -u github.com/micro/go-plugins/
RUN go get -u github.com/micro/protobuf/protoc-gen-go
RUN go get -u github.com/micro/protoc-gen-micro

# linq
RUN go get -u github.com/ahmetb/go-linq

# log
RUN go get -u github.com/rifflock/lfshook
RUN go get -u github.com/lestrrat-go/file-rotatelogs

# elastic log
RUN go get -u github.com/olivere/elastic
RUN go get -u gopkg.in/sohlich/elogrus.v3

# protobuf
RUN go get -u github.com/golang/protobuf/jsonpb
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/golang/protobuf/ptypes

# opentracing
RUN go get -u github.com/opentracing/opentracing-go

# grpc
RUN go get -u google.golang.org/grpc

WORKDIR $GOPATH