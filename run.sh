## micro server
go run ./server/server.go --registry=mdns --server_address=localhost:20001
go run ./server/cli/client.go --registry=consul

# grpc
go run ./grpc/server.go --registry=mdns
go run ./grpc/cli/client.go --registry=mdns

docker build -t gomicrok8s:1.0.8 -f k8s/Dockerfile .