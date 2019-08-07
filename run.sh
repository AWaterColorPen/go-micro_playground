#!/usr/bin/env bash
go run ./grpc/server.go --registry=mdns --server_address=localhost:20001
go run ./gateway/server.go