FROM golang:latest

WORKDIR /go/app

ADD . go-micro_playground

WORKDIR /go/app/go-micro_playground
RUN go install -v k8s/server/tosui/tosui.go
