// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package admin

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

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// Lists all account groups.
	ListAccountGroups(ctx context.Context, in *ListAccountGroupsRequest, opts ...grpc.CallOption) (Admin_ListAccountGroupsClient, error)
	// Gets an account group.
	GetAccountGroup(ctx context.Context, in *GetAccountGroupRequest, opts ...grpc.CallOption) (*GetAccountGroupResponse, error)
	// WORK-IN-PROGRESS: Gets a CloudFormation launch url for enabling cross account access to your account's billing information.
	GetDefaultBillingInfoTemplateUrl(ctx context.Context, in *GetDefaultBillingInfoTemplateUrlRequest, opts ...grpc.CallOption) (*GetDefaultBillingInfoTemplateUrlResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) ListAccountGroups(ctx context.Context, in *ListAccountGroupsRequest, opts ...grpc.CallOption) (Admin_ListAccountGroupsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Admin_ServiceDesc.Streams[0], "/blueapi.admin.v1.Admin/ListAccountGroups", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminListAccountGroupsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Admin_ListAccountGroupsClient interface {
	Recv() (*ListAccountGroupsResponse, error)
	grpc.ClientStream
}

type adminListAccountGroupsClient struct {
	grpc.ClientStream
}

func (x *adminListAccountGroupsClient) Recv() (*ListAccountGroupsResponse, error) {
	m := new(ListAccountGroupsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminClient) GetAccountGroup(ctx context.Context, in *GetAccountGroupRequest, opts ...grpc.CallOption) (*GetAccountGroupResponse, error) {
	out := new(GetAccountGroupResponse)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/GetAccountGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetDefaultBillingInfoTemplateUrl(ctx context.Context, in *GetDefaultBillingInfoTemplateUrlRequest, opts ...grpc.CallOption) (*GetDefaultBillingInfoTemplateUrlResponse, error) {
	out := new(GetDefaultBillingInfoTemplateUrlResponse)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/GetDefaultBillingInfoTemplateUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// Lists all account groups.
	ListAccountGroups(*ListAccountGroupsRequest, Admin_ListAccountGroupsServer) error
	// Gets an account group.
	GetAccountGroup(context.Context, *GetAccountGroupRequest) (*GetAccountGroupResponse, error)
	// WORK-IN-PROGRESS: Gets a CloudFormation launch url for enabling cross account access to your account's billing information.
	GetDefaultBillingInfoTemplateUrl(context.Context, *GetDefaultBillingInfoTemplateUrlRequest) (*GetDefaultBillingInfoTemplateUrlResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) ListAccountGroups(*ListAccountGroupsRequest, Admin_ListAccountGroupsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAccountGroups not implemented")
}
func (UnimplementedAdminServer) GetAccountGroup(context.Context, *GetAccountGroupRequest) (*GetAccountGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountGroup not implemented")
}
func (UnimplementedAdminServer) GetDefaultBillingInfoTemplateUrl(context.Context, *GetDefaultBillingInfoTemplateUrlRequest) (*GetDefaultBillingInfoTemplateUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultBillingInfoTemplateUrl not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_ListAccountGroups_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAccountGroupsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServer).ListAccountGroups(m, &adminListAccountGroupsServer{stream})
}

type Admin_ListAccountGroupsServer interface {
	Send(*ListAccountGroupsResponse) error
	grpc.ServerStream
}

type adminListAccountGroupsServer struct {
	grpc.ServerStream
}

func (x *adminListAccountGroupsServer) Send(m *ListAccountGroupsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Admin_GetAccountGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetAccountGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/GetAccountGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetAccountGroup(ctx, req.(*GetAccountGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetDefaultBillingInfoTemplateUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultBillingInfoTemplateUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetDefaultBillingInfoTemplateUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/GetDefaultBillingInfoTemplateUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetDefaultBillingInfoTemplateUrl(ctx, req.(*GetDefaultBillingInfoTemplateUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.admin.v1.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountGroup",
			Handler:    _Admin_GetAccountGroup_Handler,
		},
		{
			MethodName: "GetDefaultBillingInfoTemplateUrl",
			Handler:    _Admin_GetDefaultBillingInfoTemplateUrl_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAccountGroups",
			Handler:       _Admin_ListAccountGroups_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "admin/v1/admin.proto",
}
