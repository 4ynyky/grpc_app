package server

import (
	"os"

	"github.com/4ynyky/grpc_app/internal/services"
	"github.com/4ynyky/grpc_app/pkg/domains"
	"github.com/4ynyky/grpc_app/pkg/storage/memcached"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func Run() {
	stor, err := memcached.NewMemcachedStorage(memcached.Config{Host: "0.0.0.0:11211"})
	if err != nil {
		logrus.Fatalf("Failed connect to memcached: %w", err)
	}

	ss := services.NewStorageService(stor)
	ss.Set(domains.Item{ID: "ABC", Value: "3"})
	item, err := ss.Get("ABC")
	if err == nil {
		logrus.Debug(item)
	}
}
