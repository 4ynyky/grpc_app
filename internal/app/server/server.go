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

	storage := newStorage(fl)
	storageService := services.NewStorageService(storage)

	grpcConfig := grpctr.Config{Port: fl.GrpcPort(), StorageService: storageService}
	if err := grpctr.NewGrpcTransport(grpcConfig).Start(); err != nil {
		logrus.Fatalf("Failed start gRPC transport: %v", err)
	}
}

func newStorage(fl *flags) services.Storer {
	if fl.IsInternalStorage() {
		logrus.Info("Init inmemory storage")
		return inmemory.NewInMemoryStorage()
	}

	if fl.IsThirdPartyMemcache() {
		logrus.Info("Init third-party memcached connection")
		storage, err := memcached.NewMemcachedStorage(memcached.Config{Host: fl.MemcachedURL()})
		if err != nil {
			logrus.Fatalf("Failed connect to memcached: %v", err)
		}
		return storage
	}

	logrus.Info("Init my simple memcached connection")
	storage, err := memcached.NewMemcachedStorage(memcached.Config{Host: fl.MemcachedURL()})
	if err != nil {
		logrus.Fatalf("Failed connect to memcached: %v", err)
	}
	return storage
}
