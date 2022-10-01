// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/server/server.proto

package server

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

// CompanyInfoServiceClient is the client API for CompanyInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyInfoServiceClient interface {
	GetMainInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type companyInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyInfoServiceClient(cc grpc.ClientConnInterface) CompanyInfoServiceClient {
	return &companyInfoServiceClient{cc}
}

func (c *companyInfoServiceClient) GetMainInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rusprofile.CompanyInfoService/GetMainInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyInfoServiceServer is the server API for CompanyInfoService service.
// All implementations should embed UnimplementedCompanyInfoServiceServer
// for forward compatibility
type CompanyInfoServiceServer interface {
	GetMainInfo(context.Context, *Request) (*Response, error)
}

// UnimplementedCompanyInfoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCompanyInfoServiceServer struct {
}

func (UnimplementedCompanyInfoServiceServer) GetMainInfo(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainInfo not implemented")
}

// UnsafeCompanyInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyInfoServiceServer will
// result in compilation errors.
type UnsafeCompanyInfoServiceServer interface {
	mustEmbedUnimplementedCompanyInfoServiceServer()
}

func RegisterCompanyInfoServiceServer(s grpc.ServiceRegistrar, srv CompanyInfoServiceServer) {
	s.RegisterService(&CompanyInfoService_ServiceDesc, srv)
}

func _CompanyInfoService_GetMainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyInfoServiceServer).GetMainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusprofile.CompanyInfoService/GetMainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyInfoServiceServer).GetMainInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyInfoService_ServiceDesc is the grpc.ServiceDesc for CompanyInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rusprofile.CompanyInfoService",
	HandlerType: (*CompanyInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMainInfo",
			Handler:    _CompanyInfoService_GetMainInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/server/server.proto",
}