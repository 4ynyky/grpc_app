package grpctr

import (
	"context"

	"github.com/4ynyky/grpc_app/internal/domains"
	pb "github.com/4ynyky/grpc_app/internal/transport/grpctr/grpcgen/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (gt *gRPCTransport) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetReply, error) {
	err := gt.storService.Set(domains.Item{ID: in.GetId(), Value: in.GetValue()})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.SetReply{}, nil
}
func (gt *gRPCTransport) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	item, err := gt.storService.Get(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.GetReply{Value: item.Value}, nil
}
func (gt *gRPCTransport) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteReply, error) {
	err := gt.storService.Delete(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.DeleteReply{}, nil
}
