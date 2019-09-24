FROM playground_base_images:1.0.0

WORKDIR $GOPATH/src

ADD go-micro_playground go-micro_playground

RUN go install -v go-micro_playground/k8s/server/akin/akin.go



