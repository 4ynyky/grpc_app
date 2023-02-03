package grpctr

import (
	"errors"
	"fmt"
	"net"

	"github.com/4ynyky/grpc_app/internal/domains"
	pbv1 "github.com/4ynyky/grpc_app/internal/transport/grpctr/grpcgen/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	ErrConfigInvalid = errors.New("config validation failed")
)

type StorageServicer interface {
	Set(item domains.Item) error
	Get(id string) (domains.Item, error)
	Delete(id string) error
}

type Config struct {
	StorageService StorageServicer
	Port           string
}

type gRPCTransport struct {
	port        string
	storService StorageServicer
	pbv1.UnimplementedStorageServer
}

func (gt *gRPCTransport) Check() error {
	if len(gt.port) == 0 ||
		gt.storService == nil {
		return ErrConfigInvalid
	}
	return nil
}

func NewGrpcTransport(conf Config) *gRPCTransport {
	return &gRPCTransport{storService: conf.StorageService, port: conf.Port}
}

func (gt *gRPCTransport) Start() error {
	if err := gt.Check(); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", gt.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	s := grpc.NewServer()
	pbv1.RegisterStorageServer(s, gt)
	logrus.Infof("gRPC started at port: %v", gt.port)

	if err = s.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}
