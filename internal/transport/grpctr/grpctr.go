package grpctr

import (
	"fmt"
	"net"

	"github.com/4ynyky/grpc_app/internal/services"
	"github.com/4ynyky/grpc_app/internal/transport"
	pb "github.com/4ynyky/grpc_app/internal/transport/grpctr/grpcgen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Config struct {
	StorageService services.IStorageService
	Port           string
}

type storageServer struct {
	port        string
	storService services.IStorageService
	pb.UnimplementedStorageServer
}

func (ss *storageServer) Check() error {
	if len(ss.port) == 0 ||
		ss.storService == nil {
		return transport.ErrConfigInvalid
	}
	return nil
}

func NewGrpcTransport(conf Config) transport.ITransport {
	return &storageServer{storService: conf.StorageService, port: conf.Port}
}

func (ss *storageServer) Start() error {
	if err := ss.Check(); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", ss.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	s := grpc.NewServer()
	pb.RegisterStorageServer(s, ss)
	logrus.Infof("gRPC started at port: %v", ss.port)

	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
