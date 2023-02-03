build:
	go build ./cmd/storage_service/main.go
	mv main storage_app

gen:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./pkg/transport/grpctr/proto/storage.proto