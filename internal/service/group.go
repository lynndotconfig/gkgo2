package service

import (
	"context"

	pb "github.com/lynndotconfig/gkgo2/api/admin"
)

type GroupService struct {
	pb.UnimplementedGroupServer
}

func NewGroupService() *GroupService {
	return &GroupService{}
}

func (s *GroupService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupReply, error) {
	return &pb.CreateGroupReply{}, nil
}
func (s *GroupService) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*pb.UpdateGroupReply, error) {
	return &pb.UpdateGroupReply{}, nil
}
func (s *GroupService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupReply, error) {
	return &pb.DeleteGroupReply{}, nil
}
func (s *GroupService) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GetGroupReply, error) {
	return &pb.GetGroupReply{}, nil
}
func (s *GroupService) ListGroup(ctx context.Context, req *pb.ListGroupRequest) (*pb.ListGroupReply, error) {
	return &pb.ListGroupReply{}, nil
}
