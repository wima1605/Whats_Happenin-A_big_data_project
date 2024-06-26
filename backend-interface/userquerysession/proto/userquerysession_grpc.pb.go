// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: userquerysession.proto

package proto

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

const (
	UserQueryService_StartSession_FullMethodName        = "/userquerysession.UserQueryService/StartSession"
	UserQueryService_SendOverallResponse_FullMethodName = "/userquerysession.UserQueryService/SendOverallResponse"
	UserQueryService_SendTopicResponse_FullMethodName   = "/userquerysession.UserQueryService/SendTopicResponse"
)

// UserQueryServiceClient is the client API for UserQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserQueryServiceClient interface {
	StartSession(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*UserID, error)
	SendOverallResponse(ctx context.Context, in *OverallRequest, opts ...grpc.CallOption) (*OverallResponse, error)
	SendTopicResponse(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
}

type userQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserQueryServiceClient(cc grpc.ClientConnInterface) UserQueryServiceClient {
	return &userQueryServiceClient{cc}
}

func (c *userQueryServiceClient) StartSession(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, UserQueryService_StartSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userQueryServiceClient) SendOverallResponse(ctx context.Context, in *OverallRequest, opts ...grpc.CallOption) (*OverallResponse, error) {
	out := new(OverallResponse)
	err := c.cc.Invoke(ctx, UserQueryService_SendOverallResponse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userQueryServiceClient) SendTopicResponse(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, UserQueryService_SendTopicResponse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserQueryServiceServer is the server API for UserQueryService service.
// All implementations must embed UnimplementedUserQueryServiceServer
// for forward compatibility
type UserQueryServiceServer interface {
	StartSession(context.Context, *UserQuery) (*UserID, error)
	SendOverallResponse(context.Context, *OverallRequest) (*OverallResponse, error)
	SendTopicResponse(context.Context, *TopicRequest) (*TopicResponse, error)
	mustEmbedUnimplementedUserQueryServiceServer()
}

// UnimplementedUserQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserQueryServiceServer struct {
}

func (UnimplementedUserQueryServiceServer) StartSession(context.Context, *UserQuery) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSession not implemented")
}
func (UnimplementedUserQueryServiceServer) SendOverallResponse(context.Context, *OverallRequest) (*OverallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOverallResponse not implemented")
}
func (UnimplementedUserQueryServiceServer) SendTopicResponse(context.Context, *TopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTopicResponse not implemented")
}
func (UnimplementedUserQueryServiceServer) mustEmbedUnimplementedUserQueryServiceServer() {}

// UnsafeUserQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserQueryServiceServer will
// result in compilation errors.
type UnsafeUserQueryServiceServer interface {
	mustEmbedUnimplementedUserQueryServiceServer()
}

func RegisterUserQueryServiceServer(s grpc.ServiceRegistrar, srv UserQueryServiceServer) {
	s.RegisterService(&UserQueryService_ServiceDesc, srv)
}

func _UserQueryService_StartSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserQueryServiceServer).StartSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserQueryService_StartSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserQueryServiceServer).StartSession(ctx, req.(*UserQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserQueryService_SendOverallResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OverallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserQueryServiceServer).SendOverallResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserQueryService_SendOverallResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserQueryServiceServer).SendOverallResponse(ctx, req.(*OverallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserQueryService_SendTopicResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserQueryServiceServer).SendTopicResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserQueryService_SendTopicResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserQueryServiceServer).SendTopicResponse(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserQueryService_ServiceDesc is the grpc.ServiceDesc for UserQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userquerysession.UserQueryService",
	HandlerType: (*UserQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartSession",
			Handler:    _UserQueryService_StartSession_Handler,
		},
		{
			MethodName: "SendOverallResponse",
			Handler:    _UserQueryService_SendOverallResponse_Handler,
		},
		{
			MethodName: "SendTopicResponse",
			Handler:    _UserQueryService_SendTopicResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userquerysession.proto",
}
