package service

import (
	"context"

	pb "github.com/lynndotconfig/gkgo2/api/admin"
)

type TeamService struct {
	pb.UnimplementedTeamServer
}

func NewTeamService() *TeamService {
	return &TeamService{}
}

func (s *TeamService) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.CreateTeamReply, error) {
	return &pb.CreateTeamReply{}, nil
}
func (s *TeamService) UpdateTeam(ctx context.Context, req *pb.UpdateTeamRequest) (*pb.UpdateTeamReply, error) {
	return &pb.UpdateTeamReply{}, nil
}
func (s *TeamService) DeleteTeam(ctx context.Context, req *pb.DeleteTeamRequest) (*pb.DeleteTeamReply, error) {
	return &pb.DeleteTeamReply{}, nil
}
func (s *TeamService) GetTeam(ctx context.Context, req *pb.GetTeamRequest) (*pb.GetTeamReply, error) {
	return &pb.GetTeamReply{}, nil
}
func (s *TeamService) ListTeam(ctx context.Context, req *pb.ListTeamRequest) (*pb.ListTeamReply, error) {
	return &pb.ListTeamReply{}, nil
}
