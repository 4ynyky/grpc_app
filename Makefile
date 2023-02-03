build:
	golangci-lint run -c .golangci.reference.yml
	go build -o ./bin/storage_app ./cmd/storage_service/main.go

gen:
	protoc --go_out=. --go-grpc_out=. ./internal/transport/grpctr/proto.v1/storage.proto

tidy:
	go mod tidy