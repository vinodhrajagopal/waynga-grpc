// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package entityservice

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

// EntityServiceClient is the client API for EntityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntityServiceClient interface {
	// Obtains the Record with a given id.
	GetRecord(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Record, error)
	// Runs a query and returns an array of Records
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Inserts an array of Records
	Insert(ctx context.Context, in *IRecordArr, opts ...grpc.CallOption) (*SaveResponse, error)
	// Updates an array of Records
	Update(ctx context.Context, in *IRecordArr, opts ...grpc.CallOption) (*SaveResponse, error)
	// Deletes an array of ids
	Delete(ctx context.Context, in *IdArr, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type entityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityServiceClient(cc grpc.ClientConnInterface) EntityServiceClient {
	return &entityServiceClient{cc}
}

func (c *entityServiceClient) GetRecord(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, "/entityservice.EntityService/GetRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/entityservice.EntityService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) Insert(ctx context.Context, in *IRecordArr, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/entityservice.EntityService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) Update(ctx context.Context, in *IRecordArr, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/entityservice.EntityService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) Delete(ctx context.Context, in *IdArr, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/entityservice.EntityService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityServiceServer is the server API for EntityService service.
// All implementations must embed UnimplementedEntityServiceServer
// for forward compatibility
type EntityServiceServer interface {
	// Obtains the Record with a given id.
	GetRecord(context.Context, *Id) (*Record, error)
	// Runs a query and returns an array of Records
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	// Inserts an array of Records
	Insert(context.Context, *IRecordArr) (*SaveResponse, error)
	// Updates an array of Records
	Update(context.Context, *IRecordArr) (*SaveResponse, error)
	// Deletes an array of ids
	Delete(context.Context, *IdArr) (*DeleteResponse, error)
	mustEmbedUnimplementedEntityServiceServer()
}

// UnimplementedEntityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEntityServiceServer struct {
}

func (UnimplementedEntityServiceServer) GetRecord(context.Context, *Id) (*Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}
func (UnimplementedEntityServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedEntityServiceServer) Insert(context.Context, *IRecordArr) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedEntityServiceServer) Update(context.Context, *IRecordArr) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEntityServiceServer) Delete(context.Context, *IdArr) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEntityServiceServer) mustEmbedUnimplementedEntityServiceServer() {}

// UnsafeEntityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntityServiceServer will
// result in compilation errors.
type UnsafeEntityServiceServer interface {
	mustEmbedUnimplementedEntityServiceServer()
}

func RegisterEntityServiceServer(s grpc.ServiceRegistrar, srv EntityServiceServer) {
	s.RegisterService(&EntityService_ServiceDesc, srv)
}

func _EntityService_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityservice.EntityService/GetRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).GetRecord(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityservice.EntityService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IRecordArr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityservice.EntityService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Insert(ctx, req.(*IRecordArr))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IRecordArr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityservice.EntityService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Update(ctx, req.(*IRecordArr))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdArr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityservice.EntityService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Delete(ctx, req.(*IdArr))
	}
	return interceptor(ctx, in, info, handler)
}

// EntityService_ServiceDesc is the grpc.ServiceDesc for EntityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entityservice.EntityService",
	HandlerType: (*EntityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecord",
			Handler:    _EntityService_GetRecord_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _EntityService_Query_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _EntityService_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EntityService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EntityService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entityservice/entity.proto",
}
