FROM golang:latest

WORKDIR /app
ADD . .

RUN mkdir -p /log
RUN go install -v akin/akin.go
RUN go install -v tosui/tosui.go
