// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: auth.proto

package auth_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthV1Client is the client API for AuthV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthV1Client interface {
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	ConfirmEmail(ctx context.Context, in *ConfirmEmailRequest, opts ...grpc.CallOption) (*ConfirmEmailResponse, error)
	User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserIdFromToken(ctx context.Context, in *GetUserIdFromTokenRequest, opts ...grpc.CallOption) (*GetUserIdFromTokenResponse, error)
}

type authV1Client struct {
	cc grpc.ClientConnInterface
}

func NewAuthV1Client(cc grpc.ClientConnInterface) AuthV1Client {
	return &authV1Client{cc}
}

func (c *authV1Client) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthV1/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authV1Client) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthV1/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authV1Client) Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthV1/Registration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authV1Client) ConfirmEmail(ctx context.Context, in *ConfirmEmailRequest, opts ...grpc.CallOption) (*ConfirmEmailResponse, error) {
	out := new(ConfirmEmailResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthV1/ConfirmEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authV1Client) User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthV1/User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authV1Client) GetUserIdFromToken(ctx context.Context, in *GetUserIdFromTokenRequest, opts ...grpc.CallOption) (*GetUserIdFromTokenResponse, error) {
	out := new(GetUserIdFromTokenResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthV1/GetUserIdFromToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthV1Server is the server API for AuthV1 service.
// All implementations must embed UnimplementedAuthV1Server
// for forward compatibility
type AuthV1Server interface {
	Status(context.Context, *emptypb.Empty) (*StatusResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	ConfirmEmail(context.Context, *ConfirmEmailRequest) (*ConfirmEmailResponse, error)
	User(context.Context, *UserRequest) (*UserResponse, error)
	GetUserIdFromToken(context.Context, *GetUserIdFromTokenRequest) (*GetUserIdFromTokenResponse, error)
	mustEmbedUnimplementedAuthV1Server()
}

// UnimplementedAuthV1Server must be embedded to have forward compatible implementations.
type UnimplementedAuthV1Server struct {
}

func (UnimplementedAuthV1Server) Status(context.Context, *emptypb.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedAuthV1Server) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthV1Server) Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}
func (UnimplementedAuthV1Server) ConfirmEmail(context.Context, *ConfirmEmailRequest) (*ConfirmEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmEmail not implemented")
}
func (UnimplementedAuthV1Server) User(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedAuthV1Server) GetUserIdFromToken(context.Context, *GetUserIdFromTokenRequest) (*GetUserIdFromTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserIdFromToken not implemented")
}
func (UnimplementedAuthV1Server) mustEmbedUnimplementedAuthV1Server() {}

// UnsafeAuthV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthV1Server will
// result in compilation errors.
type UnsafeAuthV1Server interface {
	mustEmbedUnimplementedAuthV1Server()
}

func RegisterAuthV1Server(s grpc.ServiceRegistrar, srv AuthV1Server) {
	s.RegisterService(&AuthV1_ServiceDesc, srv)
}

func _AuthV1_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthV1/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthV1_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthV1/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthV1_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthV1/Registration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).Registration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthV1_ConfirmEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).ConfirmEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthV1/ConfirmEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).ConfirmEmail(ctx, req.(*ConfirmEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthV1_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthV1/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).User(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthV1_GetUserIdFromToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIdFromTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).GetUserIdFromToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthV1/GetUserIdFromToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).GetUserIdFromToken(ctx, req.(*GetUserIdFromTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthV1_ServiceDesc is the grpc.ServiceDesc for AuthV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_v1.AuthV1",
	HandlerType: (*AuthV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AuthV1_Status_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthV1_Login_Handler,
		},
		{
			MethodName: "Registration",
			Handler:    _AuthV1_Registration_Handler,
		},
		{
			MethodName: "ConfirmEmail",
			Handler:    _AuthV1_ConfirmEmail_Handler,
		},
		{
			MethodName: "User",
			Handler:    _AuthV1_User_Handler,
		},
		{
			MethodName: "GetUserIdFromToken",
			Handler:    _AuthV1_GetUserIdFromToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
