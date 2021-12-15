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
	// Gets a CloudFormation launch URL for enabling the default cross-account access to your account's cost information based on type.
	// See comments on the type for more information on what each template does.
	GetDefaultCostAccessTemplateUrl(ctx context.Context, in *GetDefaultCostAccessTemplateUrlRequest, opts ...grpc.CallOption) (*GetDefaultCostAccessTemplateUrlResponse, error)
	// Lists the default cross-account access role(s) attached to accounts under caller.
	ListDefaultCostAccess(ctx context.Context, in *ListDefaultCostAccessRequest, opts ...grpc.CallOption) (Admin_ListDefaultCostAccessClient, error)
	// Gets the current default cross-account role attached to the input target.
	GetDefaultCostAccess(ctx context.Context, in *GetDefaultCostAccessRequest, opts ...grpc.CallOption) (*DefaultCostAccess, error)
	// Starts validation of a default cross-account access stack deployment. If successful, the
	// IAM role created from the CloudFormation stack will be registered to the target.
	CreateDefaultCostAccess(ctx context.Context, in *CreateDefaultCostAccessRequest, opts ...grpc.CallOption) (*DefaultCostAccess, error)
	// Starts an update to an existing default cross-account access CloudFormation stack for template changes, if any.
	// Only call this API if the status of your default cross-account access is 'outdated'.
	UpdateDefaultCostAccess(ctx context.Context, in *UpdateDefaultCostAccessRequest, opts ...grpc.CallOption) (*api.Operation, error)
	// Deletes the current default cross-account access role attached to this target account.
	// This does not delete the CloudFormation deployment in your account.
	DeleteDefaultCostAccess(ctx context.Context, in *DeleteDefaultCostAccessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// WORK-IN-PROGRESS: Get notification settings for login user's organization or group.
	GetNotificationSettings(ctx context.Context, in *GetNotificationSettingsRequest, opts ...grpc.CallOption) (*api.NotificationSettings, error)
	// WORK-IN-PROGRESS: Creates or updates notification settings for login user's organization or group.
	SaveNotificationSettings(ctx context.Context, in *SaveNotificationSettingsRequest, opts ...grpc.CallOption) (*api.NotificationSettings, error)
	// WORK-IN-PROGRESS: Lists all notification channels for login user's organization or group.
	ListNotificationChannels(ctx context.Context, in *ListNotificationChannelsRequest, opts ...grpc.CallOption) (*ListNotificationChannelsResponse, error)
	// WORK-IN-PROGRESS: Gets notification channel for login user's organization or group.
	GetNotificationChannel(ctx context.Context, in *GetNotificationChannelRequest, opts ...grpc.CallOption) (*api.NotificationChannel, error)
	// WORK-IN-PROGRESS: Creates notification settings for login user's organization or group.
	CreateNotificationChannel(ctx context.Context, in *CreateNotificationChannelRequest, opts ...grpc.CallOption) (*api.NotificationChannel, error)
	// WORK-IN-PROGRESS: Updates notification settings for login user's organization or group.
	UpdateNotificationChannel(ctx context.Context, in *UpdateNotificationChannelRequest, opts ...grpc.CallOption) (*api.NotificationChannel, error)
	// WORK-IN-PROGRESS: Deletes notification settings for login user's organization or group.
	DeleteNotificationChannel(ctx context.Context, in *DeleteNotificationChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
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

func (c *adminClient) GetDefaultCostAccessTemplateUrl(ctx context.Context, in *GetDefaultCostAccessTemplateUrlRequest, opts ...grpc.CallOption) (*GetDefaultCostAccessTemplateUrlResponse, error) {
	out := new(GetDefaultCostAccessTemplateUrlResponse)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/GetDefaultCostAccessTemplateUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListDefaultCostAccess(ctx context.Context, in *ListDefaultCostAccessRequest, opts ...grpc.CallOption) (Admin_ListDefaultCostAccessClient, error) {
	stream, err := c.cc.NewStream(ctx, &Admin_ServiceDesc.Streams[1], "/blueapi.admin.v1.Admin/ListDefaultCostAccess", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminListDefaultCostAccessClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Admin_ListDefaultCostAccessClient interface {
	Recv() (*DefaultCostAccess, error)
	grpc.ClientStream
}

type adminListDefaultCostAccessClient struct {
	grpc.ClientStream
}

func (x *adminListDefaultCostAccessClient) Recv() (*DefaultCostAccess, error) {
	m := new(DefaultCostAccess)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminClient) GetDefaultCostAccess(ctx context.Context, in *GetDefaultCostAccessRequest, opts ...grpc.CallOption) (*DefaultCostAccess, error) {
	out := new(DefaultCostAccess)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/GetDefaultCostAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreateDefaultCostAccess(ctx context.Context, in *CreateDefaultCostAccessRequest, opts ...grpc.CallOption) (*DefaultCostAccess, error) {
	out := new(DefaultCostAccess)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/CreateDefaultCostAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UpdateDefaultCostAccess(ctx context.Context, in *UpdateDefaultCostAccessRequest, opts ...grpc.CallOption) (*api.Operation, error) {
	out := new(api.Operation)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/UpdateDefaultCostAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeleteDefaultCostAccess(ctx context.Context, in *DeleteDefaultCostAccessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/DeleteDefaultCostAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetNotificationSettings(ctx context.Context, in *GetNotificationSettingsRequest, opts ...grpc.CallOption) (*api.NotificationSettings, error) {
	out := new(api.NotificationSettings)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/GetNotificationSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SaveNotificationSettings(ctx context.Context, in *SaveNotificationSettingsRequest, opts ...grpc.CallOption) (*api.NotificationSettings, error) {
	out := new(api.NotificationSettings)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/SaveNotificationSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListNotificationChannels(ctx context.Context, in *ListNotificationChannelsRequest, opts ...grpc.CallOption) (*ListNotificationChannelsResponse, error) {
	out := new(ListNotificationChannelsResponse)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/ListNotificationChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetNotificationChannel(ctx context.Context, in *GetNotificationChannelRequest, opts ...grpc.CallOption) (*api.NotificationChannel, error) {
	out := new(api.NotificationChannel)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/GetNotificationChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreateNotificationChannel(ctx context.Context, in *CreateNotificationChannelRequest, opts ...grpc.CallOption) (*api.NotificationChannel, error) {
	out := new(api.NotificationChannel)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/CreateNotificationChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UpdateNotificationChannel(ctx context.Context, in *UpdateNotificationChannelRequest, opts ...grpc.CallOption) (*api.NotificationChannel, error) {
	out := new(api.NotificationChannel)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/UpdateNotificationChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeleteNotificationChannel(ctx context.Context, in *DeleteNotificationChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.admin.v1.Admin/DeleteNotificationChannel", in, out, opts...)
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
	// Gets a CloudFormation launch URL for enabling the default cross-account access to your account's cost information based on type.
	// See comments on the type for more information on what each template does.
	GetDefaultCostAccessTemplateUrl(context.Context, *GetDefaultCostAccessTemplateUrlRequest) (*GetDefaultCostAccessTemplateUrlResponse, error)
	// Lists the default cross-account access role(s) attached to accounts under caller.
	ListDefaultCostAccess(*ListDefaultCostAccessRequest, Admin_ListDefaultCostAccessServer) error
	// Gets the current default cross-account role attached to the input target.
	GetDefaultCostAccess(context.Context, *GetDefaultCostAccessRequest) (*DefaultCostAccess, error)
	// Starts validation of a default cross-account access stack deployment. If successful, the
	// IAM role created from the CloudFormation stack will be registered to the target.
	CreateDefaultCostAccess(context.Context, *CreateDefaultCostAccessRequest) (*DefaultCostAccess, error)
	// Starts an update to an existing default cross-account access CloudFormation stack for template changes, if any.
	// Only call this API if the status of your default cross-account access is 'outdated'.
	UpdateDefaultCostAccess(context.Context, *UpdateDefaultCostAccessRequest) (*api.Operation, error)
	// Deletes the current default cross-account access role attached to this target account.
	// This does not delete the CloudFormation deployment in your account.
	DeleteDefaultCostAccess(context.Context, *DeleteDefaultCostAccessRequest) (*emptypb.Empty, error)
	// WORK-IN-PROGRESS: Get notification settings for login user's organization or group.
	GetNotificationSettings(context.Context, *GetNotificationSettingsRequest) (*api.NotificationSettings, error)
	// WORK-IN-PROGRESS: Creates or updates notification settings for login user's organization or group.
	SaveNotificationSettings(context.Context, *SaveNotificationSettingsRequest) (*api.NotificationSettings, error)
	// WORK-IN-PROGRESS: Lists all notification channels for login user's organization or group.
	ListNotificationChannels(context.Context, *ListNotificationChannelsRequest) (*ListNotificationChannelsResponse, error)
	// WORK-IN-PROGRESS: Gets notification channel for login user's organization or group.
	GetNotificationChannel(context.Context, *GetNotificationChannelRequest) (*api.NotificationChannel, error)
	// WORK-IN-PROGRESS: Creates notification settings for login user's organization or group.
	CreateNotificationChannel(context.Context, *CreateNotificationChannelRequest) (*api.NotificationChannel, error)
	// WORK-IN-PROGRESS: Updates notification settings for login user's organization or group.
	UpdateNotificationChannel(context.Context, *UpdateNotificationChannelRequest) (*api.NotificationChannel, error)
	// WORK-IN-PROGRESS: Deletes notification settings for login user's organization or group.
	DeleteNotificationChannel(context.Context, *DeleteNotificationChannelRequest) (*emptypb.Empty, error)
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
func (UnimplementedAdminServer) GetDefaultCostAccessTemplateUrl(context.Context, *GetDefaultCostAccessTemplateUrlRequest) (*GetDefaultCostAccessTemplateUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultCostAccessTemplateUrl not implemented")
}
func (UnimplementedAdminServer) ListDefaultCostAccess(*ListDefaultCostAccessRequest, Admin_ListDefaultCostAccessServer) error {
	return status.Errorf(codes.Unimplemented, "method ListDefaultCostAccess not implemented")
}
func (UnimplementedAdminServer) GetDefaultCostAccess(context.Context, *GetDefaultCostAccessRequest) (*DefaultCostAccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultCostAccess not implemented")
}
func (UnimplementedAdminServer) CreateDefaultCostAccess(context.Context, *CreateDefaultCostAccessRequest) (*DefaultCostAccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDefaultCostAccess not implemented")
}
func (UnimplementedAdminServer) UpdateDefaultCostAccess(context.Context, *UpdateDefaultCostAccessRequest) (*api.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDefaultCostAccess not implemented")
}
func (UnimplementedAdminServer) DeleteDefaultCostAccess(context.Context, *DeleteDefaultCostAccessRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDefaultCostAccess not implemented")
}
func (UnimplementedAdminServer) GetNotificationSettings(context.Context, *GetNotificationSettingsRequest) (*api.NotificationSettings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationSettings not implemented")
}
func (UnimplementedAdminServer) SaveNotificationSettings(context.Context, *SaveNotificationSettingsRequest) (*api.NotificationSettings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveNotificationSettings not implemented")
}
func (UnimplementedAdminServer) ListNotificationChannels(context.Context, *ListNotificationChannelsRequest) (*ListNotificationChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotificationChannels not implemented")
}
func (UnimplementedAdminServer) GetNotificationChannel(context.Context, *GetNotificationChannelRequest) (*api.NotificationChannel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationChannel not implemented")
}
func (UnimplementedAdminServer) CreateNotificationChannel(context.Context, *CreateNotificationChannelRequest) (*api.NotificationChannel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotificationChannel not implemented")
}
func (UnimplementedAdminServer) UpdateNotificationChannel(context.Context, *UpdateNotificationChannelRequest) (*api.NotificationChannel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotificationChannel not implemented")
}
func (UnimplementedAdminServer) DeleteNotificationChannel(context.Context, *DeleteNotificationChannelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotificationChannel not implemented")
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

func _Admin_GetDefaultCostAccessTemplateUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultCostAccessTemplateUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetDefaultCostAccessTemplateUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/GetDefaultCostAccessTemplateUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetDefaultCostAccessTemplateUrl(ctx, req.(*GetDefaultCostAccessTemplateUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListDefaultCostAccess_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListDefaultCostAccessRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServer).ListDefaultCostAccess(m, &adminListDefaultCostAccessServer{stream})
}

type Admin_ListDefaultCostAccessServer interface {
	Send(*DefaultCostAccess) error
	grpc.ServerStream
}

type adminListDefaultCostAccessServer struct {
	grpc.ServerStream
}

func (x *adminListDefaultCostAccessServer) Send(m *DefaultCostAccess) error {
	return x.ServerStream.SendMsg(m)
}

func _Admin_GetDefaultCostAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultCostAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetDefaultCostAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/GetDefaultCostAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetDefaultCostAccess(ctx, req.(*GetDefaultCostAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreateDefaultCostAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDefaultCostAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateDefaultCostAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/CreateDefaultCostAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateDefaultCostAccess(ctx, req.(*CreateDefaultCostAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UpdateDefaultCostAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDefaultCostAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UpdateDefaultCostAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/UpdateDefaultCostAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UpdateDefaultCostAccess(ctx, req.(*UpdateDefaultCostAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeleteDefaultCostAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDefaultCostAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeleteDefaultCostAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/DeleteDefaultCostAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeleteDefaultCostAccess(ctx, req.(*DeleteDefaultCostAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetNotificationSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetNotificationSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/GetNotificationSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetNotificationSettings(ctx, req.(*GetNotificationSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SaveNotificationSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveNotificationSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SaveNotificationSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/SaveNotificationSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SaveNotificationSettings(ctx, req.(*SaveNotificationSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListNotificationChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotificationChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListNotificationChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/ListNotificationChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListNotificationChannels(ctx, req.(*ListNotificationChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetNotificationChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetNotificationChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/GetNotificationChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetNotificationChannel(ctx, req.(*GetNotificationChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreateNotificationChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateNotificationChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/CreateNotificationChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateNotificationChannel(ctx, req.(*CreateNotificationChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UpdateNotificationChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotificationChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UpdateNotificationChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/UpdateNotificationChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UpdateNotificationChannel(ctx, req.(*UpdateNotificationChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeleteNotificationChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotificationChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeleteNotificationChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.admin.v1.Admin/DeleteNotificationChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeleteNotificationChannel(ctx, req.(*DeleteNotificationChannelRequest))
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
			MethodName: "GetDefaultCostAccessTemplateUrl",
			Handler:    _Admin_GetDefaultCostAccessTemplateUrl_Handler,
		},
		{
			MethodName: "GetDefaultCostAccess",
			Handler:    _Admin_GetDefaultCostAccess_Handler,
		},
		{
			MethodName: "CreateDefaultCostAccess",
			Handler:    _Admin_CreateDefaultCostAccess_Handler,
		},
		{
			MethodName: "UpdateDefaultCostAccess",
			Handler:    _Admin_UpdateDefaultCostAccess_Handler,
		},
		{
			MethodName: "DeleteDefaultCostAccess",
			Handler:    _Admin_DeleteDefaultCostAccess_Handler,
		},
		{
			MethodName: "GetNotificationSettings",
			Handler:    _Admin_GetNotificationSettings_Handler,
		},
		{
			MethodName: "SaveNotificationSettings",
			Handler:    _Admin_SaveNotificationSettings_Handler,
		},
		{
			MethodName: "ListNotificationChannels",
			Handler:    _Admin_ListNotificationChannels_Handler,
		},
		{
			MethodName: "GetNotificationChannel",
			Handler:    _Admin_GetNotificationChannel_Handler,
		},
		{
			MethodName: "CreateNotificationChannel",
			Handler:    _Admin_CreateNotificationChannel_Handler,
		},
		{
			MethodName: "UpdateNotificationChannel",
			Handler:    _Admin_UpdateNotificationChannel_Handler,
		},
		{
			MethodName: "DeleteNotificationChannel",
			Handler:    _Admin_DeleteNotificationChannel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAccountGroups",
			Handler:       _Admin_ListAccountGroups_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListDefaultCostAccess",
			Handler:       _Admin_ListDefaultCostAccess_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "admin/v1/admin.proto",
}
