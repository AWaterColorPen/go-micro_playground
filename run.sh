#!/usr/bin/env bash

go run ./grpc/server.go --registry=consul --server_address=localhost:20001
go run ./gateway/server.go
go run ./client/client.go  --registry=consul

go run ./grpc/server.go --registry=mdns --server_address=localhost:20001
go run ./gateway/server.go
go run ./client/client.go --registry=mdns

## micro server
go run ./server/server.go --registry=mdns --server_address=localhost:20001

# gateway
go run ./gateway/cli/client.go

# grpc
go run ./grpc/server.go --registry=mdns
go run ./grpc/cli/client.go --registry=mdns

# web
go run ./web/server.go