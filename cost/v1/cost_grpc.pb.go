// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cost

import (
	context "context"
	aws "github.com/alphauslabs/blue-sdk-go/api/aws"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CostClient is the client API for Cost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CostClient interface {
	// Lists AWS management accounts.
	ListManagementAccounts(ctx context.Context, in *ListManagementAccountsRequest, opts ...grpc.CallOption) (*ListManagementAccountsResponse, error)
	// Gets an AWS management account. This call includes all of the account's metadata.
	// See https://alphauslabs.github.io/blueapi/ for the list of supported attributes.
	GetManagementAccount(ctx context.Context, in *GetManagementAccountRequest, opts ...grpc.CallOption) (*aws.Account, error)
	// Registers an AWS management account. See https://docs.aws.amazon.com/cur/latest/userguide/cur-create.html
	// for more information. Requirements include: Additional report details = 'Include Resource IDS' enabled,
	// Prefix = non-empty (recommendation only), Time granularity = 'Hourly', File format = 'text/csv'.
	CreateManagementAccount(ctx context.Context, in *CreateManagementAccountRequest, opts ...grpc.CallOption) (*aws.Account, error)
	// Deletes an AWS management account.
	DeleteManagementAccount(ctx context.Context, in *DeleteManagementAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Reads the usage-based cost details of an organization (Ripple) or company (Wave).
	// At the moment, the supported {vendor} is 'aws'. If datetime range parameters are
	// not set, month-to-date (current month) will be returned. Date range parameters
	// are 'start_time' and 'end_time, with the 'yyyymmdd' format.
	ReadCosts(ctx context.Context, in *ReadCostsRequest, opts ...grpc.CallOption) (Cost_ReadCostsClient, error)
	// Reads the usage-based cost details of a billing group. At the moment, the supported
	// {vendor} is 'aws'. If datetime range parameters are not set, month-to-date
	// (current month) will be returned.
	ReadBillingGroupCosts(ctx context.Context, in *ReadBillingGroupCostsRequest, opts ...grpc.CallOption) (Cost_ReadBillingGroupCostsClient, error)
	// Reads the usage-based cost details of an account. At the moment, the supported
	// {vendor} is 'aws'. If datetime range parameters are not set, month-to-date
	// (current month) will be returned.
	ReadAccountCosts(ctx context.Context, in *ReadAccountCostsRequest, opts ...grpc.CallOption) (Cost_ReadAccountCostsClient, error)
	// Reads the non-usage-based details of an organization (Ripple) or company (Wave).
	// This API covers non-usage-based adjustments, such as Fees, Credits, Discounts, Tax,
	// Upfront Fees, etc. At the moment, the supported {vendor} is 'aws'. If datetime
	// range parameters are not set, month-to-date (current month) will be returned.
	ReadAdjustments(ctx context.Context, in *ReadAdjustmentsRequest, opts ...grpc.CallOption) (Cost_ReadAdjustmentsClient, error)
	// Reads the non-usage-based details of a billing group. This API covers non-usage-based
	// adjustments, such as Fees, Credits, Discounts, Tax, Upfront Fees, etc. At the moment,
	// the supported {vendor} is 'aws'. If datetime range parameters are not set, month-to-date
	// (current month) will be returned.
	ReadBillingGroupAdjustments(ctx context.Context, in *ReadBillingGroupAdjustmentsRequest, opts ...grpc.CallOption) (Cost_ReadBillingGroupAdjustmentsClient, error)
	// Reads the non-usaged-based details of an AWS account. This API covers non-usage-based
	// adjustments, such as Fees, Credits, Discounts, Tax, Upfront Fees, etc. At the moment,
	// the supported {vendor} is 'aws'. If datetime range parameters are not set, month-to-date
	// (current month) will be returned.
	ReadAccountAdjustments(ctx context.Context, in *ReadAccountAdjustmentsRequest, opts ...grpc.CallOption) (Cost_ReadAccountAdjustmentsClient, error)
	// Reads the usage-based tag costs of a billing group. At the moment, the supported {vendor} is
	// 'aws'. If datetime range parameters are not set, month-to-date (current month) will be returned.
	ReadBillingGroupTagCosts(ctx context.Context, in *ReadBillingGroupTagCostsRequest, opts ...grpc.CallOption) (Cost_ReadBillingGroupTagCostsClient, error)
	// Reads the usage-based tag costs of an AWS account. At the moment, the supported {vendor} is
	// 'aws'. If datetime range parameters are not set, month-to-date (current month) will be returned.
	ReadAccountTagCosts(ctx context.Context, in *ReadAccountTagCostsRequest, opts ...grpc.CallOption) (Cost_ReadAccountTagCostsClient, error)
}

type costClient struct {
	cc grpc.ClientConnInterface
}

func NewCostClient(cc grpc.ClientConnInterface) CostClient {
	return &costClient{cc}
}

func (c *costClient) ListManagementAccounts(ctx context.Context, in *ListManagementAccountsRequest, opts ...grpc.CallOption) (*ListManagementAccountsResponse, error) {
	out := new(ListManagementAccountsResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/ListManagementAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costClient) GetManagementAccount(ctx context.Context, in *GetManagementAccountRequest, opts ...grpc.CallOption) (*aws.Account, error) {
	out := new(aws.Account)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/GetManagementAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costClient) CreateManagementAccount(ctx context.Context, in *CreateManagementAccountRequest, opts ...grpc.CallOption) (*aws.Account, error) {
	out := new(aws.Account)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/CreateManagementAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costClient) DeleteManagementAccount(ctx context.Context, in *DeleteManagementAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/DeleteManagementAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costClient) ReadCosts(ctx context.Context, in *ReadCostsRequest, opts ...grpc.CallOption) (Cost_ReadCostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[0], "/blueapi.cost.v1.Cost/ReadCosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &costReadCostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cost_ReadCostsClient interface {
	Recv() (*CostItem, error)
	grpc.ClientStream
}

type costReadCostsClient struct {
	grpc.ClientStream
}

func (x *costReadCostsClient) Recv() (*CostItem, error) {
	m := new(CostItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *costClient) ReadBillingGroupCosts(ctx context.Context, in *ReadBillingGroupCostsRequest, opts ...grpc.CallOption) (Cost_ReadBillingGroupCostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[1], "/blueapi.cost.v1.Cost/ReadBillingGroupCosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &costReadBillingGroupCostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cost_ReadBillingGroupCostsClient interface {
	Recv() (*CostItem, error)
	grpc.ClientStream
}

type costReadBillingGroupCostsClient struct {
	grpc.ClientStream
}

func (x *costReadBillingGroupCostsClient) Recv() (*CostItem, error) {
	m := new(CostItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *costClient) ReadAccountCosts(ctx context.Context, in *ReadAccountCostsRequest, opts ...grpc.CallOption) (Cost_ReadAccountCostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[2], "/blueapi.cost.v1.Cost/ReadAccountCosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &costReadAccountCostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cost_ReadAccountCostsClient interface {
	Recv() (*CostItem, error)
	grpc.ClientStream
}

type costReadAccountCostsClient struct {
	grpc.ClientStream
}

func (x *costReadAccountCostsClient) Recv() (*CostItem, error) {
	m := new(CostItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *costClient) ReadAdjustments(ctx context.Context, in *ReadAdjustmentsRequest, opts ...grpc.CallOption) (Cost_ReadAdjustmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[3], "/blueapi.cost.v1.Cost/ReadAdjustments", opts...)
	if err != nil {
		return nil, err
	}
	x := &costReadAdjustmentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cost_ReadAdjustmentsClient interface {
	Recv() (*AdjustmentItem, error)
	grpc.ClientStream
}

type costReadAdjustmentsClient struct {
	grpc.ClientStream
}

func (x *costReadAdjustmentsClient) Recv() (*AdjustmentItem, error) {
	m := new(AdjustmentItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *costClient) ReadBillingGroupAdjustments(ctx context.Context, in *ReadBillingGroupAdjustmentsRequest, opts ...grpc.CallOption) (Cost_ReadBillingGroupAdjustmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[4], "/blueapi.cost.v1.Cost/ReadBillingGroupAdjustments", opts...)
	if err != nil {
		return nil, err
	}
	x := &costReadBillingGroupAdjustmentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cost_ReadBillingGroupAdjustmentsClient interface {
	Recv() (*AdjustmentItem, error)
	grpc.ClientStream
}

type costReadBillingGroupAdjustmentsClient struct {
	grpc.ClientStream
}

func (x *costReadBillingGroupAdjustmentsClient) Recv() (*AdjustmentItem, error) {
	m := new(AdjustmentItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *costClient) ReadAccountAdjustments(ctx context.Context, in *ReadAccountAdjustmentsRequest, opts ...grpc.CallOption) (Cost_ReadAccountAdjustmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[5], "/blueapi.cost.v1.Cost/ReadAccountAdjustments", opts...)
	if err != nil {
		return nil, err
	}
	x := &costReadAccountAdjustmentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cost_ReadAccountAdjustmentsClient interface {
	Recv() (*AdjustmentItem, error)
	grpc.ClientStream
}

type costReadAccountAdjustmentsClient struct {
	grpc.ClientStream
}

func (x *costReadAccountAdjustmentsClient) Recv() (*AdjustmentItem, error) {
	m := new(AdjustmentItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *costClient) ReadBillingGroupTagCosts(ctx context.Context, in *ReadBillingGroupTagCostsRequest, opts ...grpc.CallOption) (Cost_ReadBillingGroupTagCostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[6], "/blueapi.cost.v1.Cost/ReadBillingGroupTagCosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &costReadBillingGroupTagCostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cost_ReadBillingGroupTagCostsClient interface {
	Recv() (*CostItem, error)
	grpc.ClientStream
}

type costReadBillingGroupTagCostsClient struct {
	grpc.ClientStream
}

func (x *costReadBillingGroupTagCostsClient) Recv() (*CostItem, error) {
	m := new(CostItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *costClient) ReadAccountTagCosts(ctx context.Context, in *ReadAccountTagCostsRequest, opts ...grpc.CallOption) (Cost_ReadAccountTagCostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[7], "/blueapi.cost.v1.Cost/ReadAccountTagCosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &costReadAccountTagCostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cost_ReadAccountTagCostsClient interface {
	Recv() (*CostItem, error)
	grpc.ClientStream
}

type costReadAccountTagCostsClient struct {
	grpc.ClientStream
}

func (x *costReadAccountTagCostsClient) Recv() (*CostItem, error) {
	m := new(CostItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CostServer is the server API for Cost service.
// All implementations must embed UnimplementedCostServer
// for forward compatibility
type CostServer interface {
	// Lists AWS management accounts.
	ListManagementAccounts(context.Context, *ListManagementAccountsRequest) (*ListManagementAccountsResponse, error)
	// Gets an AWS management account. This call includes all of the account's metadata.
	// See https://alphauslabs.github.io/blueapi/ for the list of supported attributes.
	GetManagementAccount(context.Context, *GetManagementAccountRequest) (*aws.Account, error)
	// Registers an AWS management account. See https://docs.aws.amazon.com/cur/latest/userguide/cur-create.html
	// for more information. Requirements include: Additional report details = 'Include Resource IDS' enabled,
	// Prefix = non-empty (recommendation only), Time granularity = 'Hourly', File format = 'text/csv'.
	CreateManagementAccount(context.Context, *CreateManagementAccountRequest) (*aws.Account, error)
	// Deletes an AWS management account.
	DeleteManagementAccount(context.Context, *DeleteManagementAccountRequest) (*emptypb.Empty, error)
	// Reads the usage-based cost details of an organization (Ripple) or company (Wave).
	// At the moment, the supported {vendor} is 'aws'. If datetime range parameters are
	// not set, month-to-date (current month) will be returned. Date range parameters
	// are 'start_time' and 'end_time, with the 'yyyymmdd' format.
	ReadCosts(*ReadCostsRequest, Cost_ReadCostsServer) error
	// Reads the usage-based cost details of a billing group. At the moment, the supported
	// {vendor} is 'aws'. If datetime range parameters are not set, month-to-date
	// (current month) will be returned.
	ReadBillingGroupCosts(*ReadBillingGroupCostsRequest, Cost_ReadBillingGroupCostsServer) error
	// Reads the usage-based cost details of an account. At the moment, the supported
	// {vendor} is 'aws'. If datetime range parameters are not set, month-to-date
	// (current month) will be returned.
	ReadAccountCosts(*ReadAccountCostsRequest, Cost_ReadAccountCostsServer) error
	// Reads the non-usage-based details of an organization (Ripple) or company (Wave).
	// This API covers non-usage-based adjustments, such as Fees, Credits, Discounts, Tax,
	// Upfront Fees, etc. At the moment, the supported {vendor} is 'aws'. If datetime
	// range parameters are not set, month-to-date (current month) will be returned.
	ReadAdjustments(*ReadAdjustmentsRequest, Cost_ReadAdjustmentsServer) error
	// Reads the non-usage-based details of a billing group. This API covers non-usage-based
	// adjustments, such as Fees, Credits, Discounts, Tax, Upfront Fees, etc. At the moment,
	// the supported {vendor} is 'aws'. If datetime range parameters are not set, month-to-date
	// (current month) will be returned.
	ReadBillingGroupAdjustments(*ReadBillingGroupAdjustmentsRequest, Cost_ReadBillingGroupAdjustmentsServer) error
	// Reads the non-usaged-based details of an AWS account. This API covers non-usage-based
	// adjustments, such as Fees, Credits, Discounts, Tax, Upfront Fees, etc. At the moment,
	// the supported {vendor} is 'aws'. If datetime range parameters are not set, month-to-date
	// (current month) will be returned.
	ReadAccountAdjustments(*ReadAccountAdjustmentsRequest, Cost_ReadAccountAdjustmentsServer) error
	// Reads the usage-based tag costs of a billing group. At the moment, the supported {vendor} is
	// 'aws'. If datetime range parameters are not set, month-to-date (current month) will be returned.
	ReadBillingGroupTagCosts(*ReadBillingGroupTagCostsRequest, Cost_ReadBillingGroupTagCostsServer) error
	// Reads the usage-based tag costs of an AWS account. At the moment, the supported {vendor} is
	// 'aws'. If datetime range parameters are not set, month-to-date (current month) will be returned.
	ReadAccountTagCosts(*ReadAccountTagCostsRequest, Cost_ReadAccountTagCostsServer) error
	mustEmbedUnimplementedCostServer()
}

// UnimplementedCostServer must be embedded to have forward compatible implementations.
type UnimplementedCostServer struct {
}

func (UnimplementedCostServer) ListManagementAccounts(context.Context, *ListManagementAccountsRequest) (*ListManagementAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListManagementAccounts not implemented")
}
func (UnimplementedCostServer) GetManagementAccount(context.Context, *GetManagementAccountRequest) (*aws.Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManagementAccount not implemented")
}
func (UnimplementedCostServer) CreateManagementAccount(context.Context, *CreateManagementAccountRequest) (*aws.Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateManagementAccount not implemented")
}
func (UnimplementedCostServer) DeleteManagementAccount(context.Context, *DeleteManagementAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManagementAccount not implemented")
}
func (UnimplementedCostServer) ReadCosts(*ReadCostsRequest, Cost_ReadCostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadCosts not implemented")
}
func (UnimplementedCostServer) ReadBillingGroupCosts(*ReadBillingGroupCostsRequest, Cost_ReadBillingGroupCostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadBillingGroupCosts not implemented")
}
func (UnimplementedCostServer) ReadAccountCosts(*ReadAccountCostsRequest, Cost_ReadAccountCostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadAccountCosts not implemented")
}
func (UnimplementedCostServer) ReadAdjustments(*ReadAdjustmentsRequest, Cost_ReadAdjustmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadAdjustments not implemented")
}
func (UnimplementedCostServer) ReadBillingGroupAdjustments(*ReadBillingGroupAdjustmentsRequest, Cost_ReadBillingGroupAdjustmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadBillingGroupAdjustments not implemented")
}
func (UnimplementedCostServer) ReadAccountAdjustments(*ReadAccountAdjustmentsRequest, Cost_ReadAccountAdjustmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadAccountAdjustments not implemented")
}
func (UnimplementedCostServer) ReadBillingGroupTagCosts(*ReadBillingGroupTagCostsRequest, Cost_ReadBillingGroupTagCostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadBillingGroupTagCosts not implemented")
}
func (UnimplementedCostServer) ReadAccountTagCosts(*ReadAccountTagCostsRequest, Cost_ReadAccountTagCostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadAccountTagCosts not implemented")
}
func (UnimplementedCostServer) mustEmbedUnimplementedCostServer() {}

// UnsafeCostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CostServer will
// result in compilation errors.
type UnsafeCostServer interface {
	mustEmbedUnimplementedCostServer()
}

func RegisterCostServer(s grpc.ServiceRegistrar, srv CostServer) {
	s.RegisterService(&Cost_ServiceDesc, srv)
}

func _Cost_ListManagementAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListManagementAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).ListManagementAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/ListManagementAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).ListManagementAccounts(ctx, req.(*ListManagementAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cost_GetManagementAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManagementAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).GetManagementAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/GetManagementAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).GetManagementAccount(ctx, req.(*GetManagementAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cost_CreateManagementAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateManagementAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).CreateManagementAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/CreateManagementAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).CreateManagementAccount(ctx, req.(*CreateManagementAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cost_DeleteManagementAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteManagementAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).DeleteManagementAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/DeleteManagementAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).DeleteManagementAccount(ctx, req.(*DeleteManagementAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cost_ReadCosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadCostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CostServer).ReadCosts(m, &costReadCostsServer{stream})
}

type Cost_ReadCostsServer interface {
	Send(*CostItem) error
	grpc.ServerStream
}

type costReadCostsServer struct {
	grpc.ServerStream
}

func (x *costReadCostsServer) Send(m *CostItem) error {
	return x.ServerStream.SendMsg(m)
}

func _Cost_ReadBillingGroupCosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadBillingGroupCostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CostServer).ReadBillingGroupCosts(m, &costReadBillingGroupCostsServer{stream})
}

type Cost_ReadBillingGroupCostsServer interface {
	Send(*CostItem) error
	grpc.ServerStream
}

type costReadBillingGroupCostsServer struct {
	grpc.ServerStream
}

func (x *costReadBillingGroupCostsServer) Send(m *CostItem) error {
	return x.ServerStream.SendMsg(m)
}

func _Cost_ReadAccountCosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadAccountCostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CostServer).ReadAccountCosts(m, &costReadAccountCostsServer{stream})
}

type Cost_ReadAccountCostsServer interface {
	Send(*CostItem) error
	grpc.ServerStream
}

type costReadAccountCostsServer struct {
	grpc.ServerStream
}

func (x *costReadAccountCostsServer) Send(m *CostItem) error {
	return x.ServerStream.SendMsg(m)
}

func _Cost_ReadAdjustments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadAdjustmentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CostServer).ReadAdjustments(m, &costReadAdjustmentsServer{stream})
}

type Cost_ReadAdjustmentsServer interface {
	Send(*AdjustmentItem) error
	grpc.ServerStream
}

type costReadAdjustmentsServer struct {
	grpc.ServerStream
}

func (x *costReadAdjustmentsServer) Send(m *AdjustmentItem) error {
	return x.ServerStream.SendMsg(m)
}

func _Cost_ReadBillingGroupAdjustments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadBillingGroupAdjustmentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CostServer).ReadBillingGroupAdjustments(m, &costReadBillingGroupAdjustmentsServer{stream})
}

type Cost_ReadBillingGroupAdjustmentsServer interface {
	Send(*AdjustmentItem) error
	grpc.ServerStream
}

type costReadBillingGroupAdjustmentsServer struct {
	grpc.ServerStream
}

func (x *costReadBillingGroupAdjustmentsServer) Send(m *AdjustmentItem) error {
	return x.ServerStream.SendMsg(m)
}

func _Cost_ReadAccountAdjustments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadAccountAdjustmentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CostServer).ReadAccountAdjustments(m, &costReadAccountAdjustmentsServer{stream})
}

type Cost_ReadAccountAdjustmentsServer interface {
	Send(*AdjustmentItem) error
	grpc.ServerStream
}

type costReadAccountAdjustmentsServer struct {
	grpc.ServerStream
}

func (x *costReadAccountAdjustmentsServer) Send(m *AdjustmentItem) error {
	return x.ServerStream.SendMsg(m)
}

func _Cost_ReadBillingGroupTagCosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadBillingGroupTagCostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CostServer).ReadBillingGroupTagCosts(m, &costReadBillingGroupTagCostsServer{stream})
}

type Cost_ReadBillingGroupTagCostsServer interface {
	Send(*CostItem) error
	grpc.ServerStream
}

type costReadBillingGroupTagCostsServer struct {
	grpc.ServerStream
}

func (x *costReadBillingGroupTagCostsServer) Send(m *CostItem) error {
	return x.ServerStream.SendMsg(m)
}

func _Cost_ReadAccountTagCosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadAccountTagCostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CostServer).ReadAccountTagCosts(m, &costReadAccountTagCostsServer{stream})
}

type Cost_ReadAccountTagCostsServer interface {
	Send(*CostItem) error
	grpc.ServerStream
}

type costReadAccountTagCostsServer struct {
	grpc.ServerStream
}

func (x *costReadAccountTagCostsServer) Send(m *CostItem) error {
	return x.ServerStream.SendMsg(m)
}

// Cost_ServiceDesc is the grpc.ServiceDesc for Cost service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cost_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.cost.v1.Cost",
	HandlerType: (*CostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListManagementAccounts",
			Handler:    _Cost_ListManagementAccounts_Handler,
		},
		{
			MethodName: "GetManagementAccount",
			Handler:    _Cost_GetManagementAccount_Handler,
		},
		{
			MethodName: "CreateManagementAccount",
			Handler:    _Cost_CreateManagementAccount_Handler,
		},
		{
			MethodName: "DeleteManagementAccount",
			Handler:    _Cost_DeleteManagementAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadCosts",
			Handler:       _Cost_ReadCosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadBillingGroupCosts",
			Handler:       _Cost_ReadBillingGroupCosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadAccountCosts",
			Handler:       _Cost_ReadAccountCosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadAdjustments",
			Handler:       _Cost_ReadAdjustments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadBillingGroupAdjustments",
			Handler:       _Cost_ReadBillingGroupAdjustments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadAccountAdjustments",
			Handler:       _Cost_ReadAccountAdjustments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadBillingGroupTagCosts",
			Handler:       _Cost_ReadBillingGroupTagCosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadAccountTagCosts",
			Handler:       _Cost_ReadAccountTagCosts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cost/v1/cost.proto",
}
