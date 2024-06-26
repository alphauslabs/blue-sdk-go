// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: org/v1/org.proto

package org

import (
	context "context"
	ripple "github.com/alphauslabs/blue-sdk-go/api/ripple"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Organization_CreateOrg_FullMethodName        = "/blueapi.org.v1.Organization/CreateOrg"
	Organization_SendVerification_FullMethodName = "/blueapi.org.v1.Organization/SendVerification"
	Organization_VerifyOrg_FullMethodName        = "/blueapi.org.v1.Organization/VerifyOrg"
	Organization_GetOrg_FullMethodName           = "/blueapi.org.v1.Organization/GetOrg"
	Organization_UpdateMetadata_FullMethodName   = "/blueapi.org.v1.Organization/UpdateMetadata"
	Organization_UpdatePassword_FullMethodName   = "/blueapi.org.v1.Organization/UpdatePassword"
	Organization_DeleteOrg_FullMethodName        = "/blueapi.org.v1.Organization/DeleteOrg"
)

// OrganizationClient is the client API for Organization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Organization service definition.
type OrganizationClient interface {
	// Creates the organization account.
	CreateOrg(ctx context.Context, in *CreateOrgRequest, opts ...grpc.CallOption) (*CreateOrgResponse, error)
	// WORK-IN-PROGRESS: Sends (or resends) the verification email. Only valid for unverified
	// organizations. The verification key will be valid for a day.
	SendVerification(ctx context.Context, in *SendVerificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// WORK-IN-PROGRESS: Verifies an organization using the key received from the verification email.
	// The verification key is only valid for a day.
	VerifyOrg(ctx context.Context, in *VerifyOrgRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Gets information about the caller's organization.
	GetOrg(ctx context.Context, in *GetOrgRequest, opts ...grpc.CallOption) (*ripple.Org, error)
	// WORK-IN-PROGRESS: Updates organization metadata. See [https://alphauslabs.github.io/blueapi/]
	// for the list of supported attributes.
	UpdateMetadata(ctx context.Context, in *UpdateMetadataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Updates the organization password.
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// WORK-IN-PROGRESS: Deletes the organization.
	DeleteOrg(ctx context.Context, in *DeleteOrgRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type organizationClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationClient(cc grpc.ClientConnInterface) OrganizationClient {
	return &organizationClient{cc}
}

func (c *organizationClient) CreateOrg(ctx context.Context, in *CreateOrgRequest, opts ...grpc.CallOption) (*CreateOrgResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrgResponse)
	err := c.cc.Invoke(ctx, Organization_CreateOrg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) SendVerification(ctx context.Context, in *SendVerificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Organization_SendVerification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) VerifyOrg(ctx context.Context, in *VerifyOrgRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Organization_VerifyOrg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) GetOrg(ctx context.Context, in *GetOrgRequest, opts ...grpc.CallOption) (*ripple.Org, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ripple.Org)
	err := c.cc.Invoke(ctx, Organization_GetOrg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) UpdateMetadata(ctx context.Context, in *UpdateMetadataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Organization_UpdateMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Organization_UpdatePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) DeleteOrg(ctx context.Context, in *DeleteOrgRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Organization_DeleteOrg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationServer is the server API for Organization service.
// All implementations must embed UnimplementedOrganizationServer
// for forward compatibility
//
// Organization service definition.
type OrganizationServer interface {
	// Creates the organization account.
	CreateOrg(context.Context, *CreateOrgRequest) (*CreateOrgResponse, error)
	// WORK-IN-PROGRESS: Sends (or resends) the verification email. Only valid for unverified
	// organizations. The verification key will be valid for a day.
	SendVerification(context.Context, *SendVerificationRequest) (*emptypb.Empty, error)
	// WORK-IN-PROGRESS: Verifies an organization using the key received from the verification email.
	// The verification key is only valid for a day.
	VerifyOrg(context.Context, *VerifyOrgRequest) (*emptypb.Empty, error)
	// Gets information about the caller's organization.
	GetOrg(context.Context, *GetOrgRequest) (*ripple.Org, error)
	// WORK-IN-PROGRESS: Updates organization metadata. See [https://alphauslabs.github.io/blueapi/]
	// for the list of supported attributes.
	UpdateMetadata(context.Context, *UpdateMetadataRequest) (*emptypb.Empty, error)
	// Updates the organization password.
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error)
	// WORK-IN-PROGRESS: Deletes the organization.
	DeleteOrg(context.Context, *DeleteOrgRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrganizationServer()
}

// UnimplementedOrganizationServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationServer struct {
}

func (UnimplementedOrganizationServer) CreateOrg(context.Context, *CreateOrgRequest) (*CreateOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrg not implemented")
}
func (UnimplementedOrganizationServer) SendVerification(context.Context, *SendVerificationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerification not implemented")
}
func (UnimplementedOrganizationServer) VerifyOrg(context.Context, *VerifyOrgRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOrg not implemented")
}
func (UnimplementedOrganizationServer) GetOrg(context.Context, *GetOrgRequest) (*ripple.Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrg not implemented")
}
func (UnimplementedOrganizationServer) UpdateMetadata(context.Context, *UpdateMetadataRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetadata not implemented")
}
func (UnimplementedOrganizationServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedOrganizationServer) DeleteOrg(context.Context, *DeleteOrgRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrg not implemented")
}
func (UnimplementedOrganizationServer) mustEmbedUnimplementedOrganizationServer() {}

// UnsafeOrganizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationServer will
// result in compilation errors.
type UnsafeOrganizationServer interface {
	mustEmbedUnimplementedOrganizationServer()
}

func RegisterOrganizationServer(s grpc.ServiceRegistrar, srv OrganizationServer) {
	s.RegisterService(&Organization_ServiceDesc, srv)
}

func _Organization_CreateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).CreateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_CreateOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).CreateOrg(ctx, req.(*CreateOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_SendVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).SendVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_SendVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).SendVerification(ctx, req.(*SendVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_VerifyOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).VerifyOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_VerifyOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).VerifyOrg(ctx, req.(*VerifyOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_GetOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).GetOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_GetOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).GetOrg(ctx, req.(*GetOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_UpdateMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).UpdateMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_UpdateMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).UpdateMetadata(ctx, req.(*UpdateMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_DeleteOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).DeleteOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_DeleteOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).DeleteOrg(ctx, req.(*DeleteOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Organization_ServiceDesc is the grpc.ServiceDesc for Organization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Organization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.org.v1.Organization",
	HandlerType: (*OrganizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrg",
			Handler:    _Organization_CreateOrg_Handler,
		},
		{
			MethodName: "SendVerification",
			Handler:    _Organization_SendVerification_Handler,
		},
		{
			MethodName: "VerifyOrg",
			Handler:    _Organization_VerifyOrg_Handler,
		},
		{
			MethodName: "GetOrg",
			Handler:    _Organization_GetOrg_Handler,
		},
		{
			MethodName: "UpdateMetadata",
			Handler:    _Organization_UpdateMetadata_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Organization_UpdatePassword_Handler,
		},
		{
			MethodName: "DeleteOrg",
			Handler:    _Organization_DeleteOrg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "org/v1/org.proto",
}
