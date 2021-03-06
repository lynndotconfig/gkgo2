// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package admin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TeamClient is the client API for Team service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamClient interface {
	CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamReply, error)
	UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamReply, error)
	DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamReply, error)
	GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamReply, error)
	ListTeam(ctx context.Context, in *ListTeamRequest, opts ...grpc.CallOption) (*ListTeamReply, error)
}

type teamClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamClient(cc grpc.ClientConnInterface) TeamClient {
	return &teamClient{cc}
}

func (c *teamClient) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamReply, error) {
	out := new(CreateTeamReply)
	err := c.cc.Invoke(ctx, "/api.admin.Team/CreateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamReply, error) {
	out := new(UpdateTeamReply)
	err := c.cc.Invoke(ctx, "/api.admin.Team/UpdateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamReply, error) {
	out := new(DeleteTeamReply)
	err := c.cc.Invoke(ctx, "/api.admin.Team/DeleteTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamReply, error) {
	out := new(GetTeamReply)
	err := c.cc.Invoke(ctx, "/api.admin.Team/GetTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) ListTeam(ctx context.Context, in *ListTeamRequest, opts ...grpc.CallOption) (*ListTeamReply, error) {
	out := new(ListTeamReply)
	err := c.cc.Invoke(ctx, "/api.admin.Team/ListTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServer is the server API for Team service.
// All implementations must embed UnimplementedTeamServer
// for forward compatibility
type TeamServer interface {
	CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamReply, error)
	UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamReply, error)
	DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamReply, error)
	GetTeam(context.Context, *GetTeamRequest) (*GetTeamReply, error)
	ListTeam(context.Context, *ListTeamRequest) (*ListTeamReply, error)
	mustEmbedUnimplementedTeamServer()
}

// UnimplementedTeamServer must be embedded to have forward compatible implementations.
type UnimplementedTeamServer struct {
}

func (UnimplementedTeamServer) CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (UnimplementedTeamServer) UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (UnimplementedTeamServer) DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (UnimplementedTeamServer) GetTeam(context.Context, *GetTeamRequest) (*GetTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
}
func (UnimplementedTeamServer) ListTeam(context.Context, *ListTeamRequest) (*ListTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeam not implemented")
}
func (UnimplementedTeamServer) mustEmbedUnimplementedTeamServer() {}

// UnsafeTeamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServer will
// result in compilation errors.
type UnsafeTeamServer interface {
	mustEmbedUnimplementedTeamServer()
}

func RegisterTeamServer(s grpc.ServiceRegistrar, srv TeamServer) {
	s.RegisterService(&Team_ServiceDesc, srv)
}

func _Team_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.Team/CreateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).CreateTeam(ctx, req.(*CreateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.Team/UpdateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).UpdateTeam(ctx, req.(*UpdateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.Team/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).DeleteTeam(ctx, req.(*DeleteTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.Team/GetTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).GetTeam(ctx, req.(*GetTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_ListTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).ListTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.Team/ListTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).ListTeam(ctx, req.(*ListTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Team_ServiceDesc is the grpc.ServiceDesc for Team service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Team_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.Team",
	HandlerType: (*TeamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeam",
			Handler:    _Team_CreateTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _Team_UpdateTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _Team_DeleteTeam_Handler,
		},
		{
			MethodName: "GetTeam",
			Handler:    _Team_GetTeam_Handler,
		},
		{
			MethodName: "ListTeam",
			Handler:    _Team_ListTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/admin/team.proto",
}
