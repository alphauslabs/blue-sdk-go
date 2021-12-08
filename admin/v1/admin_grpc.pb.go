// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package admin

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

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// Lists all account groups.
	ListAccountGroups(ctx context.Context, in *ListAccountGroupsRequest, opts ...grpc.CallOption) (Admin_ListAccountGroupsClient, error)
	// Gets an account group.
	GetAccountGroup(ctx context.Context, in *GetAccountGroupRequest, opts ...grpc.CallOption) (*GetAccountGroupResponse, error)
	// Gets a CloudFormation launch url for enabling cross-account access to your account's billing information.
	// Upon successful deployment, you need to validate the access by calling 'POST /admin/v1/aws/xacct/default'.
	GetDefaultBillingInfoTemplateUrl(ctx context.Context, in *GetDefaultBillingInfoTemplateUrlRequest, opts ...grpc.CallOption) (*GetDefaultBillingInfoTemplateUrlResponse, error)
	// WORK-IN-PROGRESS: Lists the current role attached to accounts under caller.
	ListDefaultBillingInfo(ctx context.Context, in *ListDefaultBillingInfoRequest, opts ...grpc.CallOption) (Admin_ListDefaultBillingInfoClient, error)
	// Gets the current role attached to input account.
	GetDefaultBillingInfo(ctx context.Context, in *GetDefaultBillingInfoRequest, opts ...grpc.CallOption) (*DefaultBillingInfo, error)
	// Starts validation of a cross-account access stack deployment. If successful,
	// the new IAM role will be registered to the target account.
	CreateDefaultBillingInfoRole(ctx context.Context, in *CreateDefaultBillingInfoRoleRequest, opts ...grpc.CallOption) (*DefaultBillingInfo, error)
	// Starts an update to an existing CloudFormation for template changes, if any.
	UpdateDefaultBillingInfoRole(ctx context.Context, in *UpdateDefaultBillingInfoRoleRequest, opts ...grpc.CallOption) (*api.Operation, error)
	// Deletes the current role attached to this target account. This does not delete the
	// CloudFormation deployment in your account.
	DeleteDefaultBillingInfoRole(ctx context.Context, in *DeleteDefaultBillingInfoRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
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

func (c *adminClient) ListDefaultBillingInfo(ctx context.Context, in *ListDefaultBillingInfoRequest, opts ...grpc.CallOption) (Admin_ListDefaultBillingInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Admin_ServiceDesc.Streams[1], "/blueapi.admin.v1.Admin/ListDefaultBillingInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminListDefaultBillingInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Admin_ListDefaultBillingInfoClient interface {
	Recv() (*DefaultBillingInfo, error)
	grpc.ClientStream
}

type adminListDefaultBillingInfoClient struct {
	grpc.ClientStream
}

func (x *adminListDefaultBillingInfoClient) Recv() (*DefaultBillingInfo, error) {
	m := new(DefaultBillingInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminClient) GetDefaultBillingInfo(ctx context.Context, in *GetDefaultBillingInfoRequest, opts ...grpc.CallOption) (*DefaultBillingInfo, error) {
	out := new(DefaultBillingInfo)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/GetDefaultBillingInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreateDefaultBillingInfoRole(ctx context.Context, in *CreateDefaultBillingInfoRoleRequest, opts ...grpc.CallOption) (*DefaultBillingInfo, error) {
	out := new(DefaultBillingInfo)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/CreateDefaultBillingInfoRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UpdateDefaultBillingInfoRole(ctx context.Context, in *UpdateDefaultBillingInfoRoleRequest, opts ...grpc.CallOption) (*api.Operation, error) {
	out := new(api.Operation)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/UpdateDefaultBillingInfoRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeleteDefaultBillingInfoRole(ctx context.Context, in *DeleteDefaultBillingInfoRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/DeleteDefaultBillingInfoRole", in, out, opts...)
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
	// Gets a CloudFormation launch url for enabling cross-account access to your account's billing information.
	// Upon successful deployment, you need to validate the access by calling 'POST /admin/v1/aws/xacct/default'.
	GetDefaultBillingInfoTemplateUrl(context.Context, *GetDefaultBillingInfoTemplateUrlRequest) (*GetDefaultBillingInfoTemplateUrlResponse, error)
	// WORK-IN-PROGRESS: Lists the current role attached to accounts under caller.
	ListDefaultBillingInfo(*ListDefaultBillingInfoRequest, Admin_ListDefaultBillingInfoServer) error
	// Gets the current role attached to input account.
	GetDefaultBillingInfo(context.Context, *GetDefaultBillingInfoRequest) (*DefaultBillingInfo, error)
	// Starts validation of a cross-account access stack deployment. If successful,
	// the new IAM role will be registered to the target account.
	CreateDefaultBillingInfoRole(context.Context, *CreateDefaultBillingInfoRoleRequest) (*DefaultBillingInfo, error)
	// Starts an update to an existing CloudFormation for template changes, if any.
	UpdateDefaultBillingInfoRole(context.Context, *UpdateDefaultBillingInfoRoleRequest) (*api.Operation, error)
	// Deletes the current role attached to this target account. This does not delete the
	// CloudFormation deployment in your account.
	DeleteDefaultBillingInfoRole(context.Context, *DeleteDefaultBillingInfoRoleRequest) (*emptypb.Empty, error)
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
func (UnimplementedAdminServer) ListDefaultBillingInfo(*ListDefaultBillingInfoRequest, Admin_ListDefaultBillingInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method ListDefaultBillingInfo not implemented")
}
func (UnimplementedAdminServer) GetDefaultBillingInfo(context.Context, *GetDefaultBillingInfoRequest) (*DefaultBillingInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultBillingInfo not implemented")
}
func (UnimplementedAdminServer) CreateDefaultBillingInfoRole(context.Context, *CreateDefaultBillingInfoRoleRequest) (*DefaultBillingInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDefaultBillingInfoRole not implemented")
}
func (UnimplementedAdminServer) UpdateDefaultBillingInfoRole(context.Context, *UpdateDefaultBillingInfoRoleRequest) (*api.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDefaultBillingInfoRole not implemented")
}
func (UnimplementedAdminServer) DeleteDefaultBillingInfoRole(context.Context, *DeleteDefaultBillingInfoRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDefaultBillingInfoRole not implemented")
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

func _Admin_ListDefaultBillingInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListDefaultBillingInfoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServer).ListDefaultBillingInfo(m, &adminListDefaultBillingInfoServer{stream})
}

type Admin_ListDefaultBillingInfoServer interface {
	Send(*DefaultBillingInfo) error
	grpc.ServerStream
}

type adminListDefaultBillingInfoServer struct {
	grpc.ServerStream
}

func (x *adminListDefaultBillingInfoServer) Send(m *DefaultBillingInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Admin_GetDefaultBillingInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultBillingInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetDefaultBillingInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/GetDefaultBillingInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetDefaultBillingInfo(ctx, req.(*GetDefaultBillingInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreateDefaultBillingInfoRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDefaultBillingInfoRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateDefaultBillingInfoRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/CreateDefaultBillingInfoRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateDefaultBillingInfoRole(ctx, req.(*CreateDefaultBillingInfoRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UpdateDefaultBillingInfoRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDefaultBillingInfoRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UpdateDefaultBillingInfoRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/UpdateDefaultBillingInfoRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UpdateDefaultBillingInfoRole(ctx, req.(*UpdateDefaultBillingInfoRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeleteDefaultBillingInfoRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDefaultBillingInfoRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeleteDefaultBillingInfoRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/DeleteDefaultBillingInfoRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeleteDefaultBillingInfoRole(ctx, req.(*DeleteDefaultBillingInfoRoleRequest))
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
		{
			MethodName: "GetDefaultBillingInfo",
			Handler:    _Admin_GetDefaultBillingInfo_Handler,
		},
		{
			MethodName: "CreateDefaultBillingInfoRole",
			Handler:    _Admin_CreateDefaultBillingInfoRole_Handler,
		},
		{
			MethodName: "UpdateDefaultBillingInfoRole",
			Handler:    _Admin_UpdateDefaultBillingInfoRole_Handler,
		},
		{
			MethodName: "DeleteDefaultBillingInfoRole",
			Handler:    _Admin_DeleteDefaultBillingInfoRole_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAccountGroups",
			Handler:       _Admin_ListAccountGroups_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListDefaultBillingInfo",
			Handler:       _Admin_ListDefaultBillingInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "admin/v1/admin.proto",
}
