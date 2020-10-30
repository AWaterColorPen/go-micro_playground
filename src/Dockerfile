FROM golang:latest

WORKDIR /app
ADD . .

RUN mkdir -p /log
RUN go install -v src/akin/akin.go
RUN go install -v src/tosui/tosui.go
