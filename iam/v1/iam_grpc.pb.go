// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package iam

import (
	context "context"
	api "github.com/alphauslabs/blue-sdk-go/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IamClient is the client API for Iam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamClient interface {
	// Gets user information about the caller. This call includes all of the user metadata.
	// See https://alphauslabs.github.io/blueapi/ for the list of supported attributes.
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*api.User, error)
	// Lists all subusers.
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	// Gets subuser information. This call includes all of the subuser metadata. See
	// https://alphauslabs.github.io/blueapi/ for the list of supported attributes.
	// If the {name} parameter is 'me' or '-', returns the caller information, which
	// is equivalent to `WhoAmI()` or `GET:/iam/v*/whoami`.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*api.User, error)
	// Lists all IP filters. At the moment, this API is only available for root users.
	ListIpFilters(ctx context.Context, in *ListIpFiltersRequest, opts ...grpc.CallOption) (*ListIpFiltersResponse, error)
	// Deletes an IP filter item. At the moment, this API is only available for root users.
	DeleteIpFilter(ctx context.Context, in *DeleteIpFilterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type iamClient struct {
	cc grpc.ClientConnInterface
}

func NewIamClient(cc grpc.ClientConnInterface) IamClient {
	return &iamClient{cc}
}

func (c *iamClient) WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*api.User, error) {
	out := new(api.User)
	err := c.cc.Invoke(ctx, "/blueapi.iam.v1.Iam/WhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/blueapi.iam.v1.Iam/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*api.User, error) {
	out := new(api.User)
	err := c.cc.Invoke(ctx, "/blueapi.iam.v1.Iam/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamClient) ListIpFilters(ctx context.Context, in *ListIpFiltersRequest, opts ...grpc.CallOption) (*ListIpFiltersResponse, error) {
	out := new(ListIpFiltersResponse)
	err := c.cc.Invoke(ctx, "/blueapi.iam.v1.Iam/ListIpFilters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamClient) DeleteIpFilter(ctx context.Context, in *DeleteIpFilterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.iam.v1.Iam/DeleteIpFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamServer is the server API for Iam service.
// All implementations must embed UnimplementedIamServer
// for forward compatibility
type IamServer interface {
	// Gets user information about the caller. This call includes all of the user metadata.
	// See https://alphauslabs.github.io/blueapi/ for the list of supported attributes.
	WhoAmI(context.Context, *WhoAmIRequest) (*api.User, error)
	// Lists all subusers.
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	// Gets subuser information. This call includes all of the subuser metadata. See
	// https://alphauslabs.github.io/blueapi/ for the list of supported attributes.
	// If the {name} parameter is 'me' or '-', returns the caller information, which
	// is equivalent to `WhoAmI()` or `GET:/iam/v*/whoami`.
	GetUser(context.Context, *GetUserRequest) (*api.User, error)
	// Lists all IP filters. At the moment, this API is only available for root users.
	ListIpFilters(context.Context, *ListIpFiltersRequest) (*ListIpFiltersResponse, error)
	// Deletes an IP filter item. At the moment, this API is only available for root users.
	DeleteIpFilter(context.Context, *DeleteIpFilterRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedIamServer()
}

// UnimplementedIamServer must be embedded to have forward compatible implementations.
type UnimplementedIamServer struct {
}

func (UnimplementedIamServer) WhoAmI(context.Context, *WhoAmIRequest) (*api.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
}
func (UnimplementedIamServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedIamServer) GetUser(context.Context, *GetUserRequest) (*api.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedIamServer) ListIpFilters(context.Context, *ListIpFiltersRequest) (*ListIpFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIpFilters not implemented")
}
func (UnimplementedIamServer) DeleteIpFilter(context.Context, *DeleteIpFilterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIpFilter not implemented")
}
func (UnimplementedIamServer) mustEmbedUnimplementedIamServer() {}

// UnsafeIamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IamServer will
// result in compilation errors.
type UnsafeIamServer interface {
	mustEmbedUnimplementedIamServer()
}

func RegisterIamServer(s grpc.ServiceRegistrar, srv IamServer) {
	s.RegisterService(&Iam_ServiceDesc, srv)
}

func _Iam_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.iam.v1.Iam/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServer).WhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iam_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.iam.v1.Iam/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iam_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.iam.v1.Iam/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iam_ListIpFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIpFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServer).ListIpFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.iam.v1.Iam/ListIpFilters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServer).ListIpFilters(ctx, req.(*ListIpFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iam_DeleteIpFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIpFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServer).DeleteIpFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.iam.v1.Iam/DeleteIpFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServer).DeleteIpFilter(ctx, req.(*DeleteIpFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Iam_ServiceDesc is the grpc.ServiceDesc for Iam service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Iam_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.iam.v1.Iam",
	HandlerType: (*IamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhoAmI",
			Handler:    _Iam_WhoAmI_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Iam_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Iam_GetUser_Handler,
		},
		{
			MethodName: "ListIpFilters",
			Handler:    _Iam_ListIpFilters_Handler,
		},
		{
			MethodName: "DeleteIpFilter",
			Handler:    _Iam_DeleteIpFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam/v1/iam.proto",
}
