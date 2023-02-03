package server

import "flag"

type flags struct {
	grpcPort          string
	memcachedURL      string
	isInternalStorage bool
}

func (fl *flags) SetupFlags() {
	flag.StringVar(&fl.grpcPort, "gp", "50051", "gRPC port")
	flag.StringVar(&fl.memcachedURL, "mu", "0.0.0.0:11211", "Memcached connection URL")
	flag.BoolVar(&fl.isInternalStorage, "is", false, "Is use internal storage instead of memcached")

	flag.Parse()
}

func (fl *flags) GrpcPort() string {
	return fl.grpcPort
}

func (fl *flags) MemcachedURL() string {
	return fl.memcachedURL
}

func (fl *flags) IsInternalStorage() bool {
	return fl.isInternalStorage
}
