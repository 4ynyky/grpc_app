package server

import (
	"flag"
	"os"

	"github.com/4ynyky/grpc_app/internal/services"
	"github.com/4ynyky/grpc_app/pkg/storage/memcached"
	"github.com/4ynyky/grpc_app/pkg/transport/grpctr"
	"github.com/sirupsen/logrus"
)

var (
	grpcPort     = flag.String("gp", "50051", "gRPC port")
	memcachedURL = flag.String("mu", "0.0.0.0:11211", "Memcached connection URL")
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func Run() {
	flag.Parse()

	storage, err := memcached.NewMemcachedStorage(memcached.Config{Host: *memcachedURL})
	if err != nil {
		logrus.Fatalf("Failed connect to memcached: %v", err)
	}

	storageService := services.NewStorageService(storage)
	err = grpctr.NewGrpcTransport(grpctr.Config{Port: *grpcPort, StorageService: storageService}).Start()
	if err != nil {
		logrus.Fatalf("Failed start gRPC transport: %v", err)
	}

}
