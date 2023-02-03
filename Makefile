build:
	go build ./cmd/storage_service/main.go
	mv main storage_app

gen:
	protoc --go_out=. --go-grpc_out=. ./internal/transport/grpctr/proto.v1/storage.proto

tidy:
	go mod tidy