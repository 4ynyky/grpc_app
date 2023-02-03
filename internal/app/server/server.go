package server

import (
	"os"

	"github.com/4ynyky/grpc_app/internal/services"
	"github.com/4ynyky/grpc_app/internal/storage/inmemory"
	"github.com/4ynyky/grpc_app/internal/storage/memcached"
	"github.com/4ynyky/grpc_app/internal/transport/grpctr"
	"github.com/sirupsen/logrus"
)

func setupLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func Run() {
	setupLogger()
	fl := &flags{}
	fl.SetupFlags()

	var storage services.Storer
	var err error

	if fl.IsInternalStorage() {
		logrus.Info("Init inmemory storage")
		storage = inmemory.NewInMemoryStorage()
	} else {
		logrus.Info("Init memcached connection")
		storage, err = memcached.NewMemcachedStorage(memcached.Config{Host: fl.MemcachedURL()})
		if err != nil {
			logrus.Fatalf("Failed connect to memcached: %v", err)
		}
	}

	storageService := services.NewStorageService(storage)
	err = grpctr.NewGrpcTransport(grpctr.Config{Port: fl.GrpcPort(), StorageService: storageService}).Start()
	if err != nil {
		logrus.Fatalf("Failed start gRPC transport: %v", err)
	}
}
