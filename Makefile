build:
	golangci-lint run -c .golangci.reference.yml
	go build -o ./bin/storage_app ./cmd/storage_service/main.go

test:
	golangci-lint run -c .golangci.reference.yml

gen:
	protoc --go_out=. --go-grpc_out=. ./internal/transport/grpctr/proto.v1/storage.proto

tidy:
	go mod tidy

run:
	golangci-lint run -c .golangci.reference.yml
	go build -o ./bin/storage_app ./cmd/storage_service/main.go
	docker-compose -f docker/memcache/docker-compose.yml up -d
	./bin/storage_app

stop: 
	docker-compose -f docker/memcache/docker-compose.yml down