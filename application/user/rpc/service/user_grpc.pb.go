// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: user.proto

package service

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
	User_Register_FullMethodName       = "/service.User/Register"
	User_FindByEmail_FullMethodName    = "/service.User/FindByEmail"
	User_FindById_FullMethodName       = "/service.User/FindById"
	User_UpdateUserInfo_FullMethodName = "/service.User/UpdateUserInfo"
	User_ModifyEmail_FullMethodName    = "/service.User/ModifyEmail"
	User_ModifyMobile_FullMethodName   = "/service.User/ModifyMobile"
	User_ModifyPassword_FullMethodName = "/service.User/ModifyPassword"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	FindByEmail(ctx context.Context, in *FindByEmailRequest, opts ...grpc.CallOption) (*FindByEmailResponse, error)
	FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*FindByIdResponse, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoResquest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	ModifyEmail(ctx context.Context, in *ModifyEmailRequest, opts ...grpc.CallOption) (*ModifyEmailResponse, error)
	ModifyMobile(ctx context.Context, in *ModifyMobileRequest, opts ...grpc.CallOption) (*ModifyMobileResponse, error)
	ModifyPassword(ctx context.Context, in *ModifyPasswordRequest, opts ...grpc.CallOption) (*ModifyPasswordResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, User_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FindByEmail(ctx context.Context, in *FindByEmailRequest, opts ...grpc.CallOption) (*FindByEmailResponse, error) {
	out := new(FindByEmailResponse)
	err := c.cc.Invoke(ctx, User_FindByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*FindByIdResponse, error) {
	out := new(FindByIdResponse)
	err := c.cc.Invoke(ctx, User_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoResquest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, User_UpdateUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ModifyEmail(ctx context.Context, in *ModifyEmailRequest, opts ...grpc.CallOption) (*ModifyEmailResponse, error) {
	out := new(ModifyEmailResponse)
	err := c.cc.Invoke(ctx, User_ModifyEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ModifyMobile(ctx context.Context, in *ModifyMobileRequest, opts ...grpc.CallOption) (*ModifyMobileResponse, error) {
	out := new(ModifyMobileResponse)
	err := c.cc.Invoke(ctx, User_ModifyMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ModifyPassword(ctx context.Context, in *ModifyPasswordRequest, opts ...grpc.CallOption) (*ModifyPasswordResponse, error) {
	out := new(ModifyPasswordResponse)
	err := c.cc.Invoke(ctx, User_ModifyPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	FindByEmail(context.Context, *FindByEmailRequest) (*FindByEmailResponse, error)
	FindById(context.Context, *FindByIdRequest) (*FindByIdResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoResquest) (*UpdateUserInfoResponse, error)
	ModifyEmail(context.Context, *ModifyEmailRequest) (*ModifyEmailResponse, error)
	ModifyMobile(context.Context, *ModifyMobileRequest) (*ModifyMobileResponse, error)
	ModifyPassword(context.Context, *ModifyPasswordRequest) (*ModifyPasswordResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) FindByEmail(context.Context, *FindByEmailRequest) (*FindByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEmail not implemented")
}
func (UnimplementedUserServer) FindById(context.Context, *FindByIdRequest) (*FindByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedUserServer) UpdateUserInfo(context.Context, *UpdateUserInfoResquest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserServer) ModifyEmail(context.Context, *ModifyEmailRequest) (*ModifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyEmail not implemented")
}
func (UnimplementedUserServer) ModifyMobile(context.Context, *ModifyMobileRequest) (*ModifyMobileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyMobile not implemented")
}
func (UnimplementedUserServer) ModifyPassword(context.Context, *ModifyPasswordRequest) (*ModifyPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyPassword not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FindByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FindByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_FindByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FindByEmail(ctx, req.(*FindByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FindById(ctx, req.(*FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ModifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ModifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ModifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ModifyEmail(ctx, req.(*ModifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ModifyMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ModifyMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ModifyMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ModifyMobile(ctx, req.(*ModifyMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ModifyPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ModifyPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ModifyPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ModifyPassword(ctx, req.(*ModifyPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "FindByEmail",
			Handler:    _User_FindByEmail_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _User_FindById_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _User_UpdateUserInfo_Handler,
		},
		{
			MethodName: "ModifyEmail",
			Handler:    _User_ModifyEmail_Handler,
		},
		{
			MethodName: "ModifyMobile",
			Handler:    _User_ModifyMobile_Handler,
		},
		{
			MethodName: "ModifyPassword",
			Handler:    _User_ModifyPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
