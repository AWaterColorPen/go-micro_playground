FROM golang:latest

WORKDIR /app
ADD . .

RUN mkdir -p /log
RUN go install -v k8s/akin/akin.go
RUN go install -v k8s/tosui/tosui.go
