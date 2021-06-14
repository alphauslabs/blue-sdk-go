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

// OrgApiClient is the client API for OrgApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrgApiClient interface {
	// Creates the organization account.
	CreateOrg(ctx context.Context, in *CreateOrgRequest, opts ...grpc.CallOption) (*Org, error)
	// Sends (or resends) the verification email. Only valid for unverified
	// organizations.
	SendVerification(ctx context.Context, in *SendVerificationRequest, opts ...grpc.CallOption) (*Org, error)
	// Verifies an organization using the key received from the verification email.
	VerifyOrg(ctx context.Context, in *VerifyOrgRequest, opts ...grpc.CallOption) (*Org, error)
	// Gets information about the caller's organization.
	GetOrg(ctx context.Context, in *GetOrgRequest, opts ...grpc.CallOption) (*Org, error)
	// Lists master accounts that belongs to the caller's organization.
	ListMasterAccounts(ctx context.Context, in *ListMasterAccountsRequest, opts ...grpc.CallOption) (*ListMasterAccountsResponse, error)
}

type orgApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgApiClient(cc grpc.ClientConnInterface) OrgApiClient {
	return &orgApiClient{cc}
}

func (c *orgApiClient) CreateOrg(ctx context.Context, in *CreateOrgRequest, opts ...grpc.CallOption) (*Org, error) {
	out := new(Org)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/CreateOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgApiClient) SendVerification(ctx context.Context, in *SendVerificationRequest, opts ...grpc.CallOption) (*Org, error) {
	out := new(Org)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/SendVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgApiClient) VerifyOrg(ctx context.Context, in *VerifyOrgRequest, opts ...grpc.CallOption) (*Org, error) {
	out := new(Org)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/VerifyOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgApiClient) GetOrg(ctx context.Context, in *GetOrgRequest, opts ...grpc.CallOption) (*Org, error) {
	out := new(Org)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/GetOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgApiClient) ListMasterAccounts(ctx context.Context, in *ListMasterAccountsRequest, opts ...grpc.CallOption) (*ListMasterAccountsResponse, error) {
	out := new(ListMasterAccountsResponse)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/ListMasterAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgApiServer is the server API for OrgApi service.
// All implementations must embed UnimplementedOrgApiServer
// for forward compatibility
type OrgApiServer interface {
	// Creates the organization account.
	CreateOrg(context.Context, *CreateOrgRequest) (*Org, error)
	// Sends (or resends) the verification email. Only valid for unverified
	// organizations.
	SendVerification(context.Context, *SendVerificationRequest) (*Org, error)
	// Verifies an organization using the key received from the verification email.
	VerifyOrg(context.Context, *VerifyOrgRequest) (*Org, error)
	// Gets information about the caller's organization.
	GetOrg(context.Context, *GetOrgRequest) (*Org, error)
	// Lists master accounts that belongs to the caller's organization.
	ListMasterAccounts(context.Context, *ListMasterAccountsRequest) (*ListMasterAccountsResponse, error)
	mustEmbedUnimplementedOrgApiServer()
}

// UnimplementedOrgApiServer must be embedded to have forward compatible implementations.
type UnimplementedOrgApiServer struct {
}

func (UnimplementedOrgApiServer) CreateOrg(context.Context, *CreateOrgRequest) (*Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrg not implemented")
}
func (UnimplementedOrgApiServer) SendVerification(context.Context, *SendVerificationRequest) (*Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerification not implemented")
}
func (UnimplementedOrgApiServer) VerifyOrg(context.Context, *VerifyOrgRequest) (*Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOrg not implemented")
}
func (UnimplementedOrgApiServer) GetOrg(context.Context, *GetOrgRequest) (*Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrg not implemented")
}
func (UnimplementedOrgApiServer) ListMasterAccounts(context.Context, *ListMasterAccountsRequest) (*ListMasterAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMasterAccounts not implemented")
}
func (UnimplementedOrgApiServer) mustEmbedUnimplementedOrgApiServer() {}

// UnsafeOrgApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrgApiServer will
// result in compilation errors.
type UnsafeOrgApiServer interface {
	mustEmbedUnimplementedOrgApiServer()
}

func RegisterOrgApiServer(s grpc.ServiceRegistrar, srv OrgApiServer) {
	s.RegisterService(&OrgApi_ServiceDesc, srv)
}

func _OrgApi_CreateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgApiServer).CreateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.org.v1.OrgApi/CreateOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgApiServer).CreateOrg(ctx, req.(*CreateOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgApi_SendVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgApiServer).SendVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.org.v1.OrgApi/SendVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgApiServer).SendVerification(ctx, req.(*SendVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgApi_VerifyOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgApiServer).VerifyOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.org.v1.OrgApi/VerifyOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgApiServer).VerifyOrg(ctx, req.(*VerifyOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgApi_GetOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgApiServer).GetOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.org.v1.OrgApi/GetOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgApiServer).GetOrg(ctx, req.(*GetOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgApi_ListMasterAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMasterAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgApiServer).ListMasterAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.org.v1.OrgApi/ListMasterAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgApiServer).ListMasterAccounts(ctx, req.(*ListMasterAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrgApi_ServiceDesc is the grpc.ServiceDesc for OrgApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrgApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.org.v1.OrgApi",
	HandlerType: (*OrgApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrg",
			Handler:    _OrgApi_CreateOrg_Handler,
		},
		{
			MethodName: "SendVerification",
			Handler:    _OrgApi_SendVerification_Handler,
		},
		{
			MethodName: "VerifyOrg",
			Handler:    _OrgApi_VerifyOrg_Handler,
		},
		{
			MethodName: "GetOrg",
			Handler:    _OrgApi_GetOrg_Handler,
		},
		{
			MethodName: "ListMasterAccounts",
			Handler:    _OrgApi_ListMasterAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "org/v1/org.proto",
}
