package grpctr

import (
	"context"

	"github.com/4ynyky/grpc_app/internal/domains"
	pb "github.com/4ynyky/grpc_app/internal/transport/grpctr/grpcgen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ss *storageServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetReply, error) {
	err := ss.storService.Set(domains.Item{ID: in.GetId(), Value: in.GetValue()})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.SetReply{}, nil
}
func (ss *storageServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	item, err := ss.storService.Get(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.GetReply{Value: item.Value}, nil
}
func (ss *storageServer) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteReply, error) {
	err := ss.storService.Delete(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.DeleteReply{}, nil
}
