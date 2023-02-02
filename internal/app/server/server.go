package server

import (
	"os"

	"github.com/4ynyky/grpc_app/pkg/domains"
	"github.com/4ynyky/grpc_app/pkg/storage"
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
	stor.Set(domains.Item{ID: "ABC", Value: "3"})
	item, err := stor.Get("ABCq")
	if err == storage.ErrNotFound {
		logrus.Warn(err)
	} else if err != nil {
		logrus.Error(err)
	}
	logrus.Debug(item)
}
