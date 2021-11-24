// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package billing

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

// BillingClient is the client API for Billing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillingClient interface {
	// Lists all billing groups.
	ListBillingGroups(ctx context.Context, in *ListBillingGroupsRequest, opts ...grpc.CallOption) (Billing_ListBillingGroupsClient, error)
	// Registers a billing group.
	CreateBillingGroup(ctx context.Context, in *CreateBillingGroupRequest, opts ...grpc.CallOption) (*BillingGroup, error)
	// Gets a billing group.
	GetBillingGroup(ctx context.Context, in *GetBillingGroupRequest, opts ...grpc.CallOption) (*GetBillingGroupResponse, error)
	// WORK-IN-PROGRESS: Gets an access group.
	GetAccessGroup(ctx context.Context, in *GetAccessGroupRequest, opts ...grpc.CallOption) (*GetAccessGroupResponse, error)
}

type billingClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingClient(cc grpc.ClientConnInterface) BillingClient {
	return &billingClient{cc}
}

func (c *billingClient) ListBillingGroups(ctx context.Context, in *ListBillingGroupsRequest, opts ...grpc.CallOption) (Billing_ListBillingGroupsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Billing_ServiceDesc.Streams[0], "/blueapi.billing.v1.Billing/ListBillingGroups", opts...)
	if err != nil {
		return nil, err
	}
	x := &billingListBillingGroupsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Billing_ListBillingGroupsClient interface {
	Recv() (*BillingGroup, error)
	grpc.ClientStream
}

type billingListBillingGroupsClient struct {
	grpc.ClientStream
}

func (x *billingListBillingGroupsClient) Recv() (*BillingGroup, error) {
	m := new(BillingGroup)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *billingClient) CreateBillingGroup(ctx context.Context, in *CreateBillingGroupRequest, opts ...grpc.CallOption) (*BillingGroup, error) {
	out := new(BillingGroup)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/CreateBillingGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) GetBillingGroup(ctx context.Context, in *GetBillingGroupRequest, opts ...grpc.CallOption) (*GetBillingGroupResponse, error) {
	out := new(GetBillingGroupResponse)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/GetBillingGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) GetAccessGroup(ctx context.Context, in *GetAccessGroupRequest, opts ...grpc.CallOption) (*GetAccessGroupResponse, error) {
	out := new(GetAccessGroupResponse)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/GetAccessGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingServer is the server API for Billing service.
// All implementations must embed UnimplementedBillingServer
// for forward compatibility
type BillingServer interface {
	// Lists all billing groups.
	ListBillingGroups(*ListBillingGroupsRequest, Billing_ListBillingGroupsServer) error
	// Registers a billing group.
	CreateBillingGroup(context.Context, *CreateBillingGroupRequest) (*BillingGroup, error)
	// Gets a billing group.
	GetBillingGroup(context.Context, *GetBillingGroupRequest) (*GetBillingGroupResponse, error)
	// WORK-IN-PROGRESS: Gets an access group.
	GetAccessGroup(context.Context, *GetAccessGroupRequest) (*GetAccessGroupResponse, error)
	mustEmbedUnimplementedBillingServer()
}

// UnimplementedBillingServer must be embedded to have forward compatible implementations.
type UnimplementedBillingServer struct {
}

func (UnimplementedBillingServer) ListBillingGroups(*ListBillingGroupsRequest, Billing_ListBillingGroupsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBillingGroups not implemented")
}
func (UnimplementedBillingServer) CreateBillingGroup(context.Context, *CreateBillingGroupRequest) (*BillingGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBillingGroup not implemented")
}
func (UnimplementedBillingServer) GetBillingGroup(context.Context, *GetBillingGroupRequest) (*GetBillingGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBillingGroup not implemented")
}
func (UnimplementedBillingServer) GetAccessGroup(context.Context, *GetAccessGroupRequest) (*GetAccessGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessGroup not implemented")
}
func (UnimplementedBillingServer) mustEmbedUnimplementedBillingServer() {}

// UnsafeBillingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingServer will
// result in compilation errors.
type UnsafeBillingServer interface {
	mustEmbedUnimplementedBillingServer()
}

func RegisterBillingServer(s grpc.ServiceRegistrar, srv BillingServer) {
	s.RegisterService(&Billing_ServiceDesc, srv)
}

func _Billing_ListBillingGroups_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBillingGroupsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BillingServer).ListBillingGroups(m, &billingListBillingGroupsServer{stream})
}

type Billing_ListBillingGroupsServer interface {
	Send(*BillingGroup) error
	grpc.ServerStream
}

type billingListBillingGroupsServer struct {
	grpc.ServerStream
}

func (x *billingListBillingGroupsServer) Send(m *BillingGroup) error {
	return x.ServerStream.SendMsg(m)
}

func _Billing_CreateBillingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBillingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).CreateBillingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/CreateBillingGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).CreateBillingGroup(ctx, req.(*CreateBillingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_GetBillingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GetBillingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/GetBillingGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GetBillingGroup(ctx, req.(*GetBillingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_GetAccessGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GetAccessGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/GetAccessGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GetAccessGroup(ctx, req.(*GetAccessGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Billing_ServiceDesc is the grpc.ServiceDesc for Billing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Billing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.billing.v1.Billing",
	HandlerType: (*BillingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBillingGroup",
			Handler:    _Billing_CreateBillingGroup_Handler,
		},
		{
			MethodName: "GetBillingGroup",
			Handler:    _Billing_GetBillingGroup_Handler,
		},
		{
			MethodName: "GetAccessGroup",
			Handler:    _Billing_GetAccessGroup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBillingGroups",
			Handler:       _Billing_ListBillingGroups_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "billing/v1/billing.proto",
}
