build:
	go build ./cmd/storage_service/main.go
	mv main storage_app

gen:
	protoc --go_out=. --go-grpc_out=. ./pkg/transport/grpctr/proto/storage.proto