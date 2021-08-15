// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package orgsignup

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

// SignupServiceClient is the client API for SignupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignupServiceClient interface {
	// Signups an org with the provided details
	OrgSignup(ctx context.Context, in *OrgSignupRequest, opts ...grpc.CallOption) (*OrgSignupResponse, error)
	UserSignup(ctx context.Context, in *UserSignupRequest, opts ...grpc.CallOption) (*OrgSignupResponse, error)
}

type signupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignupServiceClient(cc grpc.ClientConnInterface) SignupServiceClient {
	return &signupServiceClient{cc}
}

func (c *signupServiceClient) OrgSignup(ctx context.Context, in *OrgSignupRequest, opts ...grpc.CallOption) (*OrgSignupResponse, error) {
	out := new(OrgSignupResponse)
	err := c.cc.Invoke(ctx, "/signup.SignupService/OrgSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) UserSignup(ctx context.Context, in *UserSignupRequest, opts ...grpc.CallOption) (*OrgSignupResponse, error) {
	out := new(OrgSignupResponse)
	err := c.cc.Invoke(ctx, "/signup.SignupService/UserSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignupServiceServer is the server API for SignupService service.
// All implementations must embed UnimplementedSignupServiceServer
// for forward compatibility
type SignupServiceServer interface {
	// Signups an org with the provided details
	OrgSignup(context.Context, *OrgSignupRequest) (*OrgSignupResponse, error)
	UserSignup(context.Context, *UserSignupRequest) (*OrgSignupResponse, error)
	mustEmbedUnimplementedSignupServiceServer()
}

// UnimplementedSignupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSignupServiceServer struct {
}

func (UnimplementedSignupServiceServer) OrgSignup(context.Context, *OrgSignupRequest) (*OrgSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrgSignup not implemented")
}
func (UnimplementedSignupServiceServer) UserSignup(context.Context, *UserSignupRequest) (*OrgSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignup not implemented")
}
func (UnimplementedSignupServiceServer) mustEmbedUnimplementedSignupServiceServer() {}

// UnsafeSignupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignupServiceServer will
// result in compilation errors.
type UnsafeSignupServiceServer interface {
	mustEmbedUnimplementedSignupServiceServer()
}

func RegisterSignupServiceServer(s grpc.ServiceRegistrar, srv SignupServiceServer) {
	s.RegisterService(&SignupService_ServiceDesc, srv)
}

func _SignupService_OrgSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).OrgSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup.SignupService/OrgSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).OrgSignup(ctx, req.(*OrgSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_UserSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).UserSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup.SignupService/UserSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).UserSignup(ctx, req.(*UserSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignupService_ServiceDesc is the grpc.ServiceDesc for SignupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "signup.SignupService",
	HandlerType: (*SignupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrgSignup",
			Handler:    _SignupService_OrgSignup_Handler,
		},
		{
			MethodName: "UserSignup",
			Handler:    _SignupService_UserSignup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signup/signup.proto",
}
