// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: cover/v1/cover.proto

package cover

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

// CoverClient is the client API for Cover service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoverClient interface {
	// Invite members to the system
	InviteMember(ctx context.Context, in *InviteMemberRequest, opts ...grpc.CallOption) (*InviteMemberResponse, error)
	// Create a member
	CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*CreateMemberResponse, error)
	// Get all the members/subusers of the company
	GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*GetMembersResponse, error)
	// Get the details of the cover user
	GetMemberDetails(ctx context.Context, in *GetMemberDetailsRequest, opts ...grpc.CallOption) (*GetMemberDetailsResponse, error)
	// Deletes a cover user
	DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*DeleteMemberResponse, error)
	// Trigger reset password from Admin
	AdminResetPassword(ctx context.Context, in *AdminResetPasswordRequest, opts ...grpc.CallOption) (*AdminResetPasswordResponse, error)
	// Reset member's password
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	// Get all the views
	GetViews(ctx context.Context, in *GetViewsRequest, opts ...grpc.CallOption) (*GetViewsResponse, error)
	// Get the details of the current view
	GetCurrentView(ctx context.Context, in *GetCurrentViewRequest, opts ...grpc.CallOption) (*GetCurrentViewResponse, error)
	// Updates the view details
	UpdateView(ctx context.Context, in *UpdateViewRequest, opts ...grpc.CallOption) (*UpdateViewResponse, error)
	// Deletes a view
	DeleteView(ctx context.Context, in *DeleteViewRequest, opts ...grpc.CallOption) (*DeleteViewResponse, error)
}

type coverClient struct {
	cc grpc.ClientConnInterface
}

func NewCoverClient(cc grpc.ClientConnInterface) CoverClient {
	return &coverClient{cc}
}

func (c *coverClient) InviteMember(ctx context.Context, in *InviteMemberRequest, opts ...grpc.CallOption) (*InviteMemberResponse, error) {
	out := new(InviteMemberResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/InviteMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverClient) CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*CreateMemberResponse, error) {
	out := new(CreateMemberResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/CreateMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverClient) GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*GetMembersResponse, error) {
	out := new(GetMembersResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/GetMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverClient) GetMemberDetails(ctx context.Context, in *GetMemberDetailsRequest, opts ...grpc.CallOption) (*GetMemberDetailsResponse, error) {
	out := new(GetMemberDetailsResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/GetMemberDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverClient) DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*DeleteMemberResponse, error) {
	out := new(DeleteMemberResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/DeleteMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverClient) AdminResetPassword(ctx context.Context, in *AdminResetPasswordRequest, opts ...grpc.CallOption) (*AdminResetPasswordResponse, error) {
	out := new(AdminResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/AdminResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverClient) GetViews(ctx context.Context, in *GetViewsRequest, opts ...grpc.CallOption) (*GetViewsResponse, error) {
	out := new(GetViewsResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/GetViews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverClient) GetCurrentView(ctx context.Context, in *GetCurrentViewRequest, opts ...grpc.CallOption) (*GetCurrentViewResponse, error) {
	out := new(GetCurrentViewResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/GetCurrentView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverClient) UpdateView(ctx context.Context, in *UpdateViewRequest, opts ...grpc.CallOption) (*UpdateViewResponse, error) {
	out := new(UpdateViewResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/UpdateView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverClient) DeleteView(ctx context.Context, in *DeleteViewRequest, opts ...grpc.CallOption) (*DeleteViewResponse, error) {
	out := new(DeleteViewResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cover.v1.Cover/DeleteView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoverServer is the server API for Cover service.
// All implementations must embed UnimplementedCoverServer
// for forward compatibility
type CoverServer interface {
	// Invite members to the system
	InviteMember(context.Context, *InviteMemberRequest) (*InviteMemberResponse, error)
	// Create a member
	CreateMember(context.Context, *CreateMemberRequest) (*CreateMemberResponse, error)
	// Get all the members/subusers of the company
	GetMembers(context.Context, *GetMembersRequest) (*GetMembersResponse, error)
	// Get the details of the cover user
	GetMemberDetails(context.Context, *GetMemberDetailsRequest) (*GetMemberDetailsResponse, error)
	// Deletes a cover user
	DeleteMember(context.Context, *DeleteMemberRequest) (*DeleteMemberResponse, error)
	// Trigger reset password from Admin
	AdminResetPassword(context.Context, *AdminResetPasswordRequest) (*AdminResetPasswordResponse, error)
	// Reset member's password
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	// Get all the views
	GetViews(context.Context, *GetViewsRequest) (*GetViewsResponse, error)
	// Get the details of the current view
	GetCurrentView(context.Context, *GetCurrentViewRequest) (*GetCurrentViewResponse, error)
	// Updates the view details
	UpdateView(context.Context, *UpdateViewRequest) (*UpdateViewResponse, error)
	// Deletes a view
	DeleteView(context.Context, *DeleteViewRequest) (*DeleteViewResponse, error)
	mustEmbedUnimplementedCoverServer()
}

// UnimplementedCoverServer must be embedded to have forward compatible implementations.
type UnimplementedCoverServer struct {
}

func (UnimplementedCoverServer) InviteMember(context.Context, *InviteMemberRequest) (*InviteMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteMember not implemented")
}
func (UnimplementedCoverServer) CreateMember(context.Context, *CreateMemberRequest) (*CreateMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMember not implemented")
}
func (UnimplementedCoverServer) GetMembers(context.Context, *GetMembersRequest) (*GetMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembers not implemented")
}
func (UnimplementedCoverServer) GetMemberDetails(context.Context, *GetMemberDetailsRequest) (*GetMemberDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberDetails not implemented")
}
func (UnimplementedCoverServer) DeleteMember(context.Context, *DeleteMemberRequest) (*DeleteMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMember not implemented")
}
func (UnimplementedCoverServer) AdminResetPassword(context.Context, *AdminResetPasswordRequest) (*AdminResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminResetPassword not implemented")
}
func (UnimplementedCoverServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedCoverServer) GetViews(context.Context, *GetViewsRequest) (*GetViewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetViews not implemented")
}
func (UnimplementedCoverServer) GetCurrentView(context.Context, *GetCurrentViewRequest) (*GetCurrentViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentView not implemented")
}
func (UnimplementedCoverServer) UpdateView(context.Context, *UpdateViewRequest) (*UpdateViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateView not implemented")
}
func (UnimplementedCoverServer) DeleteView(context.Context, *DeleteViewRequest) (*DeleteViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteView not implemented")
}
func (UnimplementedCoverServer) mustEmbedUnimplementedCoverServer() {}

// UnsafeCoverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoverServer will
// result in compilation errors.
type UnsafeCoverServer interface {
	mustEmbedUnimplementedCoverServer()
}

func RegisterCoverServer(s grpc.ServiceRegistrar, srv CoverServer) {
	s.RegisterService(&Cover_ServiceDesc, srv)
}

func _Cover_InviteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).InviteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/InviteMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).InviteMember(ctx, req.(*InviteMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cover_CreateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).CreateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/CreateMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).CreateMember(ctx, req.(*CreateMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cover_GetMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).GetMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/GetMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).GetMembers(ctx, req.(*GetMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cover_GetMemberDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).GetMemberDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/GetMemberDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).GetMemberDetails(ctx, req.(*GetMemberDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cover_DeleteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).DeleteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/DeleteMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).DeleteMember(ctx, req.(*DeleteMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cover_AdminResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).AdminResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/AdminResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).AdminResetPassword(ctx, req.(*AdminResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cover_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cover_GetViews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetViewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).GetViews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/GetViews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).GetViews(ctx, req.(*GetViewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cover_GetCurrentView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).GetCurrentView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/GetCurrentView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).GetCurrentView(ctx, req.(*GetCurrentViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cover_UpdateView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).UpdateView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/UpdateView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).UpdateView(ctx, req.(*UpdateViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cover_DeleteView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServer).DeleteView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cover.v1.Cover/DeleteView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServer).DeleteView(ctx, req.(*DeleteViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cover_ServiceDesc is the grpc.ServiceDesc for Cover service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cover_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.cover.v1.Cover",
	HandlerType: (*CoverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InviteMember",
			Handler:    _Cover_InviteMember_Handler,
		},
		{
			MethodName: "CreateMember",
			Handler:    _Cover_CreateMember_Handler,
		},
		{
			MethodName: "GetMembers",
			Handler:    _Cover_GetMembers_Handler,
		},
		{
			MethodName: "GetMemberDetails",
			Handler:    _Cover_GetMemberDetails_Handler,
		},
		{
			MethodName: "DeleteMember",
			Handler:    _Cover_DeleteMember_Handler,
		},
		{
			MethodName: "AdminResetPassword",
			Handler:    _Cover_AdminResetPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _Cover_ResetPassword_Handler,
		},
		{
			MethodName: "GetViews",
			Handler:    _Cover_GetViews_Handler,
		},
		{
			MethodName: "GetCurrentView",
			Handler:    _Cover_GetCurrentView_Handler,
		},
		{
			MethodName: "UpdateView",
			Handler:    _Cover_UpdateView_Handler,
		},
		{
			MethodName: "DeleteView",
			Handler:    _Cover_DeleteView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cover/v1/cover.proto",
}
