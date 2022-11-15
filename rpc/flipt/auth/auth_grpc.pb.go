// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: auth/auth.proto

package auth

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

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	GetAuthenticationSelf(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Authentication, error)
	GetAuthentication(ctx context.Context, in *GetAuthenticationRequest, opts ...grpc.CallOption) (*Authentication, error)
	DeleteAuthentication(ctx context.Context, in *DeleteAuthenticationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) GetAuthenticationSelf(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Authentication, error) {
	out := new(Authentication)
	err := c.cc.Invoke(ctx, "/flipt.auth.AuthenticationService/GetAuthenticationSelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) GetAuthentication(ctx context.Context, in *GetAuthenticationRequest, opts ...grpc.CallOption) (*Authentication, error) {
	out := new(Authentication)
	err := c.cc.Invoke(ctx, "/flipt.auth.AuthenticationService/GetAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) DeleteAuthentication(ctx context.Context, in *DeleteAuthenticationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/flipt.auth.AuthenticationService/DeleteAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
// All implementations must embed UnimplementedAuthenticationServiceServer
// for forward compatibility
type AuthenticationServiceServer interface {
	GetAuthenticationSelf(context.Context, *emptypb.Empty) (*Authentication, error)
	GetAuthentication(context.Context, *GetAuthenticationRequest) (*Authentication, error)
	DeleteAuthentication(context.Context, *DeleteAuthenticationRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAuthenticationServiceServer()
}

// UnimplementedAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (UnimplementedAuthenticationServiceServer) GetAuthenticationSelf(context.Context, *emptypb.Empty) (*Authentication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthenticationSelf not implemented")
}
func (UnimplementedAuthenticationServiceServer) GetAuthentication(context.Context, *GetAuthenticationRequest) (*Authentication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthentication not implemented")
}
func (UnimplementedAuthenticationServiceServer) DeleteAuthentication(context.Context, *DeleteAuthenticationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthentication not implemented")
}
func (UnimplementedAuthenticationServiceServer) mustEmbedUnimplementedAuthenticationServiceServer() {}

// UnsafeAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServiceServer will
// result in compilation errors.
type UnsafeAuthenticationServiceServer interface {
	mustEmbedUnimplementedAuthenticationServiceServer()
}

func RegisterAuthenticationServiceServer(s grpc.ServiceRegistrar, srv AuthenticationServiceServer) {
	s.RegisterService(&AuthenticationService_ServiceDesc, srv)
}

func _AuthenticationService_GetAuthenticationSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GetAuthenticationSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flipt.auth.AuthenticationService/GetAuthenticationSelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GetAuthenticationSelf(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_GetAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GetAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flipt.auth.AuthenticationService/GetAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GetAuthentication(ctx, req.(*GetAuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_DeleteAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).DeleteAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flipt.auth.AuthenticationService/DeleteAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).DeleteAuthentication(ctx, req.(*DeleteAuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationService_ServiceDesc is the grpc.ServiceDesc for AuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipt.auth.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthenticationSelf",
			Handler:    _AuthenticationService_GetAuthenticationSelf_Handler,
		},
		{
			MethodName: "GetAuthentication",
			Handler:    _AuthenticationService_GetAuthentication_Handler,
		},
		{
			MethodName: "DeleteAuthentication",
			Handler:    _AuthenticationService_DeleteAuthentication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

// AuthenticationMethodTokenServiceClient is the client API for AuthenticationMethodTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationMethodTokenServiceClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
}

type authenticationMethodTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationMethodTokenServiceClient(cc grpc.ClientConnInterface) AuthenticationMethodTokenServiceClient {
	return &authenticationMethodTokenServiceClient{cc}
}

func (c *authenticationMethodTokenServiceClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, "/flipt.auth.AuthenticationMethodTokenService/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationMethodTokenServiceServer is the server API for AuthenticationMethodTokenService service.
// All implementations must embed UnimplementedAuthenticationMethodTokenServiceServer
// for forward compatibility
type AuthenticationMethodTokenServiceServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	mustEmbedUnimplementedAuthenticationMethodTokenServiceServer()
}

// UnimplementedAuthenticationMethodTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationMethodTokenServiceServer struct {
}

func (UnimplementedAuthenticationMethodTokenServiceServer) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedAuthenticationMethodTokenServiceServer) mustEmbedUnimplementedAuthenticationMethodTokenServiceServer() {
}

// UnsafeAuthenticationMethodTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationMethodTokenServiceServer will
// result in compilation errors.
type UnsafeAuthenticationMethodTokenServiceServer interface {
	mustEmbedUnimplementedAuthenticationMethodTokenServiceServer()
}

func RegisterAuthenticationMethodTokenServiceServer(s grpc.ServiceRegistrar, srv AuthenticationMethodTokenServiceServer) {
	s.RegisterService(&AuthenticationMethodTokenService_ServiceDesc, srv)
}

func _AuthenticationMethodTokenService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationMethodTokenServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flipt.auth.AuthenticationMethodTokenService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationMethodTokenServiceServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationMethodTokenService_ServiceDesc is the grpc.ServiceDesc for AuthenticationMethodTokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationMethodTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipt.auth.AuthenticationMethodTokenService",
	HandlerType: (*AuthenticationMethodTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _AuthenticationMethodTokenService_CreateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
