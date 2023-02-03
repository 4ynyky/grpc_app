package server

import (
	"flag"
	"os"

	"github.com/4ynyky/grpc_app/internal/services"
	"github.com/4ynyky/grpc_app/pkg/storage"
	"github.com/4ynyky/grpc_app/pkg/storage/inmemory"
	"github.com/4ynyky/grpc_app/pkg/storage/memcached"
	"github.com/4ynyky/grpc_app/pkg/transport/grpctr"
	"github.com/sirupsen/logrus"
)

var (
	grpcPort          = flag.String("gp", "50051", "gRPC port")
	memcachedURL      = flag.String("mu", "0.0.0.0:11211", "Memcached connection URL")
	isInternalStorage = flag.Bool("is", false, "Is use internal storage instead of memcached")
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func Run() {
	flag.Parse()

	var storage storage.IStorage
	var err error

	if *isInternalStorage {
		logrus.Info("Init inmemory storage")
		storage = inmemory.NewInMemoryStorage()
	} else {
		logrus.Info("Init memcached connection")
		storage, err = memcached.NewMemcachedStorage(memcached.Config{Host: *memcachedURL})
		if err != nil {
			logrus.Fatalf("Failed connect to memcached: %v", err)
		}
	}

	storageService := services.NewStorageService(storage)
	err = grpctr.NewGrpcTransport(grpctr.Config{Port: *grpcPort, StorageService: storageService}).Start()
	if err != nil {
		logrus.Fatalf("Failed start gRPC transport: %v", err)
	}

}
