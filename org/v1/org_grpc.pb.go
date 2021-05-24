// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package org

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

// OrgClient is the client API for Org service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrgClient interface {
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error)
}

type orgClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgClient(cc grpc.ClientConnInterface) OrgClient {
	return &orgClient{cc}
}

func (c *orgClient) WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error) {
	out := new(WhoAmIResponse)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.Org/WhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgServer is the server API for Org service.
// All implementations must embed UnimplementedOrgServer
// for forward compatibility
type OrgServer interface {
	WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResponse, error)
	mustEmbedUnimplementedOrgServer()
}

// UnimplementedOrgServer must be embedded to have forward compatible implementations.
type UnimplementedOrgServer struct {
}

func (UnimplementedOrgServer) WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
}
func (UnimplementedOrgServer) mustEmbedUnimplementedOrgServer() {}

// UnsafeOrgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrgServer will
// result in compilation errors.
type UnsafeOrgServer interface {
	mustEmbedUnimplementedOrgServer()
}

func RegisterOrgServer(s grpc.ServiceRegistrar, srv OrgServer) {
	s.RegisterService(&Org_ServiceDesc, srv)
}

func _Org_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.org.v1.Org/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).WhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Org_ServiceDesc is the grpc.ServiceDesc for Org service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Org_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.org.v1.Org",
	HandlerType: (*OrgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhoAmI",
			Handler:    _Org_WhoAmI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "org/v1/org.proto",
}