// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: pb_invite/invite.proto

package pb_invite

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
	Invite_InitiateChatInvite_FullMethodName = "/pb_invite.Invite/InitiateChatInvite"
	Invite_ChatInviteList_FullMethodName     = "/pb_invite.Invite/ChatInviteList"
	Invite_ChatInviteHandle_FullMethodName   = "/pb_invite.Invite/ChatInviteHandle"
)

// InviteClient is the client API for Invite service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InviteClient interface {
	InitiateChatInvite(ctx context.Context, in *InitiateChatInviteReq, opts ...grpc.CallOption) (*InitiateChatInviteResp, error)
	ChatInviteList(ctx context.Context, in *ChatInviteListReq, opts ...grpc.CallOption) (*ChatInviteListResp, error)
	ChatInviteHandle(ctx context.Context, in *ChatInviteHandleReq, opts ...grpc.CallOption) (*ChatInviteHandleResp, error)
}

type inviteClient struct {
	cc grpc.ClientConnInterface
}

func NewInviteClient(cc grpc.ClientConnInterface) InviteClient {
	return &inviteClient{cc}
}

func (c *inviteClient) InitiateChatInvite(ctx context.Context, in *InitiateChatInviteReq, opts ...grpc.CallOption) (*InitiateChatInviteResp, error) {
	out := new(InitiateChatInviteResp)
	err := c.cc.Invoke(ctx, Invite_InitiateChatInvite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteClient) ChatInviteList(ctx context.Context, in *ChatInviteListReq, opts ...grpc.CallOption) (*ChatInviteListResp, error) {
	out := new(ChatInviteListResp)
	err := c.cc.Invoke(ctx, Invite_ChatInviteList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteClient) ChatInviteHandle(ctx context.Context, in *ChatInviteHandleReq, opts ...grpc.CallOption) (*ChatInviteHandleResp, error) {
	out := new(ChatInviteHandleResp)
	err := c.cc.Invoke(ctx, Invite_ChatInviteHandle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InviteServer is the server API for Invite service.
// All implementations must embed UnimplementedInviteServer
// for forward compatibility
type InviteServer interface {
	InitiateChatInvite(context.Context, *InitiateChatInviteReq) (*InitiateChatInviteResp, error)
	ChatInviteList(context.Context, *ChatInviteListReq) (*ChatInviteListResp, error)
	ChatInviteHandle(context.Context, *ChatInviteHandleReq) (*ChatInviteHandleResp, error)
	mustEmbedUnimplementedInviteServer()
}

// UnimplementedInviteServer must be embedded to have forward compatible implementations.
type UnimplementedInviteServer struct {
}

func (UnimplementedInviteServer) InitiateChatInvite(context.Context, *InitiateChatInviteReq) (*InitiateChatInviteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateChatInvite not implemented")
}
func (UnimplementedInviteServer) ChatInviteList(context.Context, *ChatInviteListReq) (*ChatInviteListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatInviteList not implemented")
}
func (UnimplementedInviteServer) ChatInviteHandle(context.Context, *ChatInviteHandleReq) (*ChatInviteHandleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatInviteHandle not implemented")
}
func (UnimplementedInviteServer) mustEmbedUnimplementedInviteServer() {}

// UnsafeInviteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InviteServer will
// result in compilation errors.
type UnsafeInviteServer interface {
	mustEmbedUnimplementedInviteServer()
}

func RegisterInviteServer(s grpc.ServiceRegistrar, srv InviteServer) {
	s.RegisterService(&Invite_ServiceDesc, srv)
}

func _Invite_InitiateChatInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitiateChatInviteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteServer).InitiateChatInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Invite_InitiateChatInvite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteServer).InitiateChatInvite(ctx, req.(*InitiateChatInviteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invite_ChatInviteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatInviteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteServer).ChatInviteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Invite_ChatInviteList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteServer).ChatInviteList(ctx, req.(*ChatInviteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invite_ChatInviteHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatInviteHandleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteServer).ChatInviteHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Invite_ChatInviteHandle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteServer).ChatInviteHandle(ctx, req.(*ChatInviteHandleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Invite_ServiceDesc is the grpc.ServiceDesc for Invite service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invite_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_invite.Invite",
	HandlerType: (*InviteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitiateChatInvite",
			Handler:    _Invite_InitiateChatInvite_Handler,
		},
		{
			MethodName: "ChatInviteList",
			Handler:    _Invite_ChatInviteList_Handler,
		},
		{
			MethodName: "ChatInviteHandle",
			Handler:    _Invite_ChatInviteHandle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_invite/invite.proto",
}
