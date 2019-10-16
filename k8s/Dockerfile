FROM hub.docker.com/_/golang:latest

WORKDIR /app
ADD . .

RUN go install -v k8s/server/akin/akin.go
RUN go install -v k8s/server/natsu/natsu.go
RUN go install -v k8s/server/shunmu/shunmu.go
RUN go install -v k8s/server/tosui/tosui.go
