package server

import (
	"github.com/yemingfeng/sdb/internal/cluster"
	"github.com/yemingfeng/sdb/internal/pb"
	"github.com/yemingfeng/sdb/internal/service"
	"golang.org/x/net/context"
)

type MapServer struct {
	pb.UnimplementedSDBServer
}

func (server *MapServer) MPush(_ context.Context, request *pb.MPushRequest) (*pb.MPushResponse, error) {
	err := cluster.Apply("MPush", request)
	return &pb.MPushResponse{Success: err == nil}, err
}

func (server *MapServer) MPop(_ context.Context, request *pb.MPopRequest) (*pb.MPopResponse, error) {
	err := cluster.Apply("MPop", request)
	return &pb.MPopResponse{Success: err == nil}, err
}

func (server *MapServer) MExist(_ context.Context, request *pb.MExistRequest) (*pb.MExistResponse, error) {
	res, err := service.MExist(request.Key, request.Keys)
	return &pb.MExistResponse{Exists: res}, err
}

func (server *MapServer) MDel(_ context.Context, request *pb.MDelRequest) (*pb.MDelResponse, error) {
	err := cluster.Apply("MDel", request)
	return &pb.MDelResponse{Success: err == nil}, err
}

func (server *MapServer) MCount(_ context.Context, request *pb.MCountRequest) (*pb.MCountResponse, error) {
	res, err := service.MCount(request.Key)
	return &pb.MCountResponse{Count: res}, err
}

func (server *MapServer) MMembers(_ context.Context, request *pb.MMembersRequest) (*pb.MMembersResponse, error) {
	res, err := service.MMembers(request.Key)
	return &pb.MMembersResponse{Pairs: res}, err
}
