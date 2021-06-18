// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package org

import (
	context "context"
	aws "github.com/alphauslabs/blue-sdk-go/types/aws"
	ripple "github.com/alphauslabs/blue-sdk-go/types/ripple"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	CreateOrg(ctx context.Context, in *CreateOrgRequest, opts ...grpc.CallOption) (*CreateOrgResponse, error)
	// Sends (or resends) the verification email. Only valid for unverified
	// organizations. The verification key will be valid for a day.
	SendVerification(ctx context.Context, in *SendVerificationRequest, opts ...grpc.CallOption) (*ripple.Org, error)
	// Verifies an organization using the key received from the verification email.
	// The verification key is only valid for a day.
	VerifyOrg(ctx context.Context, in *VerifyOrgRequest, opts ...grpc.CallOption) (*ripple.Org, error)
	// Gets information about the caller's organization.
	GetOrg(ctx context.Context, in *GetOrgRequest, opts ...grpc.CallOption) (*ripple.Org, error)
	// Updates organization metadata. See https://alphauslabs.github.io/blueapi/
	// for the list of supported attributes.
	UpdateMetadata(ctx context.Context, in *UpdateMetadataRequest, opts ...grpc.CallOption) (*ripple.Org, error)
	// Updates the organization password.
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*ripple.Org, error)
	// Lists master accounts that belongs to the caller's organization.
	ListMasterAccounts(ctx context.Context, in *ListMasterAccountsRequest, opts ...grpc.CallOption) (*ListMasterAccountsResponse, error)
	// Get master account. This call includes all of the account's metadata. See
	// https://alphauslabs.github.io/blueapi/ for the list of supported attributes.
	GetMasterAccount(ctx context.Context, in *GetMasterAccountRequest, opts ...grpc.CallOption) (*aws.Account, error)
	// Deletes the organization.
	DeleteOrg(ctx context.Context, in *DeleteOrgRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orgApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgApiClient(cc grpc.ClientConnInterface) OrgApiClient {
	return &orgApiClient{cc}
}

func (c *orgApiClient) CreateOrg(ctx context.Context, in *CreateOrgRequest, opts ...grpc.CallOption) (*CreateOrgResponse, error) {
	out := new(CreateOrgResponse)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/CreateOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgApiClient) SendVerification(ctx context.Context, in *SendVerificationRequest, opts ...grpc.CallOption) (*ripple.Org, error) {
	out := new(ripple.Org)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/SendVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgApiClient) VerifyOrg(ctx context.Context, in *VerifyOrgRequest, opts ...grpc.CallOption) (*ripple.Org, error) {
	out := new(ripple.Org)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/VerifyOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgApiClient) GetOrg(ctx context.Context, in *GetOrgRequest, opts ...grpc.CallOption) (*ripple.Org, error) {
	out := new(ripple.Org)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/GetOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgApiClient) UpdateMetadata(ctx context.Context, in *UpdateMetadataRequest, opts ...grpc.CallOption) (*ripple.Org, error) {
	out := new(ripple.Org)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/UpdateMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgApiClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*ripple.Org, error) {
	out := new(ripple.Org)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/UpdatePassword", in, out, opts...)
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

func (c *orgApiClient) GetMasterAccount(ctx context.Context, in *GetMasterAccountRequest, opts ...grpc.CallOption) (*aws.Account, error) {
	out := new(aws.Account)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/GetMasterAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgApiClient) DeleteOrg(ctx context.Context, in *DeleteOrgRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.org.v1.OrgApi/DeleteOrg", in, out, opts...)
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
	CreateOrg(context.Context, *CreateOrgRequest) (*CreateOrgResponse, error)
	// Sends (or resends) the verification email. Only valid for unverified
	// organizations. The verification key will be valid for a day.
	SendVerification(context.Context, *SendVerificationRequest) (*ripple.Org, error)
	// Verifies an organization using the key received from the verification email.
	// The verification key is only valid for a day.
	VerifyOrg(context.Context, *VerifyOrgRequest) (*ripple.Org, error)
	// Gets information about the caller's organization.
	GetOrg(context.Context, *GetOrgRequest) (*ripple.Org, error)
	// Updates organization metadata. See https://alphauslabs.github.io/blueapi/
	// for the list of supported attributes.
	UpdateMetadata(context.Context, *UpdateMetadataRequest) (*ripple.Org, error)
	// Updates the organization password.
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*ripple.Org, error)
	// Lists master accounts that belongs to the caller's organization.
	ListMasterAccounts(context.Context, *ListMasterAccountsRequest) (*ListMasterAccountsResponse, error)
	// Get master account. This call includes all of the account's metadata. See
	// https://alphauslabs.github.io/blueapi/ for the list of supported attributes.
	GetMasterAccount(context.Context, *GetMasterAccountRequest) (*aws.Account, error)
	// Deletes the organization.
	DeleteOrg(context.Context, *DeleteOrgRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrgApiServer()
}

// UnimplementedOrgApiServer must be embedded to have forward compatible implementations.
type UnimplementedOrgApiServer struct {
}

func (UnimplementedOrgApiServer) CreateOrg(context.Context, *CreateOrgRequest) (*CreateOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrg not implemented")
}
func (UnimplementedOrgApiServer) SendVerification(context.Context, *SendVerificationRequest) (*ripple.Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerification not implemented")
}
func (UnimplementedOrgApiServer) VerifyOrg(context.Context, *VerifyOrgRequest) (*ripple.Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOrg not implemented")
}
func (UnimplementedOrgApiServer) GetOrg(context.Context, *GetOrgRequest) (*ripple.Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrg not implemented")
}
func (UnimplementedOrgApiServer) UpdateMetadata(context.Context, *UpdateMetadataRequest) (*ripple.Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetadata not implemented")
}
func (UnimplementedOrgApiServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*ripple.Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedOrgApiServer) ListMasterAccounts(context.Context, *ListMasterAccountsRequest) (*ListMasterAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMasterAccounts not implemented")
}
func (UnimplementedOrgApiServer) GetMasterAccount(context.Context, *GetMasterAccountRequest) (*aws.Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMasterAccount not implemented")
}
func (UnimplementedOrgApiServer) DeleteOrg(context.Context, *DeleteOrgRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrg not implemented")
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

func _OrgApi_UpdateMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgApiServer).UpdateMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.org.v1.OrgApi/UpdateMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgApiServer).UpdateMetadata(ctx, req.(*UpdateMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgApi_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgApiServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.org.v1.OrgApi/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgApiServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
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

func _OrgApi_GetMasterAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMasterAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgApiServer).GetMasterAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.org.v1.OrgApi/GetMasterAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgApiServer).GetMasterAccount(ctx, req.(*GetMasterAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgApi_DeleteOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgApiServer).DeleteOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.org.v1.OrgApi/DeleteOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgApiServer).DeleteOrg(ctx, req.(*DeleteOrgRequest))
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
			MethodName: "UpdateMetadata",
			Handler:    _OrgApi_UpdateMetadata_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _OrgApi_UpdatePassword_Handler,
		},
		{
			MethodName: "ListMasterAccounts",
			Handler:    _OrgApi_ListMasterAccounts_Handler,
		},
		{
			MethodName: "GetMasterAccount",
			Handler:    _OrgApi_GetMasterAccount_Handler,
		},
		{
			MethodName: "DeleteOrg",
			Handler:    _OrgApi_DeleteOrg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "org/v1/org.proto",
}
