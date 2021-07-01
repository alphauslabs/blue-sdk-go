// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cost

import (
	context "context"
	api "github.com/alphauslabs/blue-sdk-go/api"
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
	// Gets an AWS management account's update history, which is a list of timestamps our system tracks when the account's CUR files are
	// exported to your S3 by AWS, which in turn, triggers the import from your S3 to our system for processing.
	GetManagementAccountUpdateHistory(ctx context.Context, in *GetManagementAccountUpdateHistoryRequest, opts ...grpc.CallOption) (*GetManagementAccountUpdateHistoryResponse, error)
	// Registers an AWS management account. See [https://docs.aws.amazon.com/cur/latest/userguide/cur-create.html]
	// for more information. Requirements include: Additional report details = 'Include Resource IDS' enabled,
	// Prefix = non-empty (recommendation only), Time granularity = 'Hourly', File format = 'text/csv'.
	CreateManagementAccount(ctx context.Context, in *CreateManagementAccountRequest, opts ...grpc.CallOption) (*aws.Account, error)
	// Deletes an AWS management account.
	DeleteManagementAccount(ctx context.Context, in *DeleteManagementAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Initiates an ondemand import of all registered CUR files.
	ImportCurFiles(ctx context.Context, in *ImportCurFilesRequest, opts ...grpc.CallOption) (*api.Operation, error)
	// Triggers monthly calculations for costs and invoices at either organization or billing group level.
	CalculateCosts(ctx context.Context, in *CalculateCostsRequest, opts ...grpc.CallOption) (*api.Operation, error)
	// Reads the usage-based cost details of an organization (Ripple) or billing group (Wave).
	// At the moment, the supported {vendor} is 'aws'. If datetime range parameters are
	// not set, month-to-date (current month) will be returned.
	ReadCosts(ctx context.Context, in *ReadCostsRequest, opts ...grpc.CallOption) (Cost_ReadCostsClient, error)
	// Reads the non-usage-based details of an organization (Ripple) or billing group (Wave).
	// This API covers non-usage-based adjustments, such as Fees, Credits, Discounts, Tax,
	// Upfront Fees, etc. At the moment, the supported {vendor} is 'aws'. If datetime
	// range parameters are not set, month-to-date (current month) will be returned.
	ReadAdjustments(ctx context.Context, in *ReadAdjustmentsRequest, opts ...grpc.CallOption) (Cost_ReadAdjustmentsClient, error)
	// Reads the usage-based tag costs of a billing group. At the moment, the supported {vendor} is
	// 'aws'. If datetime range parameters are not set, month-to-date (current month) will be returned.
	ReadBillingGroupTagCosts(ctx context.Context, in *ReadBillingGroupTagCostsRequest, opts ...grpc.CallOption) (Cost_ReadBillingGroupTagCostsClient, error)
	// Reads the usage-based tag costs of an AWS account. At the moment, the supported {vendor} is
	// 'aws'. If datetime range parameters are not set, month-to-date (current month) will be returned.
	ReadAccountTagCosts(ctx context.Context, in *ReadAccountTagCostsRequest, opts ...grpc.CallOption) (Cost_ReadAccountTagCostsClient, error)
	// Creates a budget configuration.
	CreateBudgetConfig(ctx context.Context, in *CreateBudgetConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Updates an existing budget configuration.
	UpdateBudgetConfig(ctx context.Context, in *UpdateBudgetConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Fetches budget configurations for all accounts under the specified billing group.
	// Set accountId to fetch budget configuration for specific account only.
	GetBudgetConfig(ctx context.Context, in *GetBudgetConfigRequest, opts ...grpc.CallOption) (*GetBudgetConfigResponse, error)
	// Fetches cost forecasts for the specified billing group.
	// Includes historical cost (up to previous month) and forecasted cost (up to three months for now).
	GetForecasts(ctx context.Context, in *GetForecastsRequest, opts ...grpc.CallOption) (*GetForecastsResponse, error)
	// Fetches month-to-date accumulated costs vs forecasted cost vs budget for the specified billing group.
	GetMonthToDateForecast(ctx context.Context, in *GetMonthToDateForecastRequest, opts ...grpc.CallOption) (*GetMonthToDateForecastResponse, error)
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

func (c *costClient) GetManagementAccountUpdateHistory(ctx context.Context, in *GetManagementAccountUpdateHistoryRequest, opts ...grpc.CallOption) (*GetManagementAccountUpdateHistoryResponse, error) {
	out := new(GetManagementAccountUpdateHistoryResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/GetManagementAccountUpdateHistory", in, out, opts...)
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

func (c *costClient) ImportCurFiles(ctx context.Context, in *ImportCurFilesRequest, opts ...grpc.CallOption) (*api.Operation, error) {
	out := new(api.Operation)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/ImportCurFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costClient) CalculateCosts(ctx context.Context, in *CalculateCostsRequest, opts ...grpc.CallOption) (*api.Operation, error) {
	out := new(api.Operation)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/CalculateCosts", in, out, opts...)
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

func (c *costClient) ReadAdjustments(ctx context.Context, in *ReadAdjustmentsRequest, opts ...grpc.CallOption) (Cost_ReadAdjustmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[1], "/blueapi.cost.v1.Cost/ReadAdjustments", opts...)
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

func (c *costClient) ReadBillingGroupTagCosts(ctx context.Context, in *ReadBillingGroupTagCostsRequest, opts ...grpc.CallOption) (Cost_ReadBillingGroupTagCostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[2], "/blueapi.cost.v1.Cost/ReadBillingGroupTagCosts", opts...)
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
	stream, err := c.cc.NewStream(ctx, &Cost_ServiceDesc.Streams[3], "/blueapi.cost.v1.Cost/ReadAccountTagCosts", opts...)
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

func (c *costClient) CreateBudgetConfig(ctx context.Context, in *CreateBudgetConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/CreateBudgetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costClient) UpdateBudgetConfig(ctx context.Context, in *UpdateBudgetConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/UpdateBudgetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costClient) GetBudgetConfig(ctx context.Context, in *GetBudgetConfigRequest, opts ...grpc.CallOption) (*GetBudgetConfigResponse, error) {
	out := new(GetBudgetConfigResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/GetBudgetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costClient) GetForecasts(ctx context.Context, in *GetForecastsRequest, opts ...grpc.CallOption) (*GetForecastsResponse, error) {
	out := new(GetForecastsResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/GetForecasts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costClient) GetMonthToDateForecast(ctx context.Context, in *GetMonthToDateForecastRequest, opts ...grpc.CallOption) (*GetMonthToDateForecastResponse, error) {
	out := new(GetMonthToDateForecastResponse)
	err := c.cc.Invoke(ctx, "/blueapi.cost.v1.Cost/GetMonthToDateForecast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	// Gets an AWS management account's update history, which is a list of timestamps our system tracks when the account's CUR files are
	// exported to your S3 by AWS, which in turn, triggers the import from your S3 to our system for processing.
	GetManagementAccountUpdateHistory(context.Context, *GetManagementAccountUpdateHistoryRequest) (*GetManagementAccountUpdateHistoryResponse, error)
	// Registers an AWS management account. See [https://docs.aws.amazon.com/cur/latest/userguide/cur-create.html]
	// for more information. Requirements include: Additional report details = 'Include Resource IDS' enabled,
	// Prefix = non-empty (recommendation only), Time granularity = 'Hourly', File format = 'text/csv'.
	CreateManagementAccount(context.Context, *CreateManagementAccountRequest) (*aws.Account, error)
	// Deletes an AWS management account.
	DeleteManagementAccount(context.Context, *DeleteManagementAccountRequest) (*emptypb.Empty, error)
	// Initiates an ondemand import of all registered CUR files.
	ImportCurFiles(context.Context, *ImportCurFilesRequest) (*api.Operation, error)
	// Triggers monthly calculations for costs and invoices at either organization or billing group level.
	CalculateCosts(context.Context, *CalculateCostsRequest) (*api.Operation, error)
	// Reads the usage-based cost details of an organization (Ripple) or billing group (Wave).
	// At the moment, the supported {vendor} is 'aws'. If datetime range parameters are
	// not set, month-to-date (current month) will be returned.
	ReadCosts(*ReadCostsRequest, Cost_ReadCostsServer) error
	// Reads the non-usage-based details of an organization (Ripple) or billing group (Wave).
	// This API covers non-usage-based adjustments, such as Fees, Credits, Discounts, Tax,
	// Upfront Fees, etc. At the moment, the supported {vendor} is 'aws'. If datetime
	// range parameters are not set, month-to-date (current month) will be returned.
	ReadAdjustments(*ReadAdjustmentsRequest, Cost_ReadAdjustmentsServer) error
	// Reads the usage-based tag costs of a billing group. At the moment, the supported {vendor} is
	// 'aws'. If datetime range parameters are not set, month-to-date (current month) will be returned.
	ReadBillingGroupTagCosts(*ReadBillingGroupTagCostsRequest, Cost_ReadBillingGroupTagCostsServer) error
	// Reads the usage-based tag costs of an AWS account. At the moment, the supported {vendor} is
	// 'aws'. If datetime range parameters are not set, month-to-date (current month) will be returned.
	ReadAccountTagCosts(*ReadAccountTagCostsRequest, Cost_ReadAccountTagCostsServer) error
	// Creates a budget configuration.
	CreateBudgetConfig(context.Context, *CreateBudgetConfigRequest) (*emptypb.Empty, error)
	// Updates an existing budget configuration.
	UpdateBudgetConfig(context.Context, *UpdateBudgetConfigRequest) (*emptypb.Empty, error)
	// Fetches budget configurations for all accounts under the specified billing group.
	// Set accountId to fetch budget configuration for specific account only.
	GetBudgetConfig(context.Context, *GetBudgetConfigRequest) (*GetBudgetConfigResponse, error)
	// Fetches cost forecasts for the specified billing group.
	// Includes historical cost (up to previous month) and forecasted cost (up to three months for now).
	GetForecasts(context.Context, *GetForecastsRequest) (*GetForecastsResponse, error)
	// Fetches month-to-date accumulated costs vs forecasted cost vs budget for the specified billing group.
	GetMonthToDateForecast(context.Context, *GetMonthToDateForecastRequest) (*GetMonthToDateForecastResponse, error)
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
func (UnimplementedCostServer) GetManagementAccountUpdateHistory(context.Context, *GetManagementAccountUpdateHistoryRequest) (*GetManagementAccountUpdateHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManagementAccountUpdateHistory not implemented")
}
func (UnimplementedCostServer) CreateManagementAccount(context.Context, *CreateManagementAccountRequest) (*aws.Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateManagementAccount not implemented")
}
func (UnimplementedCostServer) DeleteManagementAccount(context.Context, *DeleteManagementAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManagementAccount not implemented")
}
func (UnimplementedCostServer) ImportCurFiles(context.Context, *ImportCurFilesRequest) (*api.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportCurFiles not implemented")
}
func (UnimplementedCostServer) CalculateCosts(context.Context, *CalculateCostsRequest) (*api.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateCosts not implemented")
}
func (UnimplementedCostServer) ReadCosts(*ReadCostsRequest, Cost_ReadCostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadCosts not implemented")
}
func (UnimplementedCostServer) ReadAdjustments(*ReadAdjustmentsRequest, Cost_ReadAdjustmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadAdjustments not implemented")
}
func (UnimplementedCostServer) ReadBillingGroupTagCosts(*ReadBillingGroupTagCostsRequest, Cost_ReadBillingGroupTagCostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadBillingGroupTagCosts not implemented")
}
func (UnimplementedCostServer) ReadAccountTagCosts(*ReadAccountTagCostsRequest, Cost_ReadAccountTagCostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadAccountTagCosts not implemented")
}
func (UnimplementedCostServer) CreateBudgetConfig(context.Context, *CreateBudgetConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBudgetConfig not implemented")
}
func (UnimplementedCostServer) UpdateBudgetConfig(context.Context, *UpdateBudgetConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBudgetConfig not implemented")
}
func (UnimplementedCostServer) GetBudgetConfig(context.Context, *GetBudgetConfigRequest) (*GetBudgetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBudgetConfig not implemented")
}
func (UnimplementedCostServer) GetForecasts(context.Context, *GetForecastsRequest) (*GetForecastsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForecasts not implemented")
}
func (UnimplementedCostServer) GetMonthToDateForecast(context.Context, *GetMonthToDateForecastRequest) (*GetMonthToDateForecastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthToDateForecast not implemented")
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

func _Cost_GetManagementAccountUpdateHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManagementAccountUpdateHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).GetManagementAccountUpdateHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/GetManagementAccountUpdateHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).GetManagementAccountUpdateHistory(ctx, req.(*GetManagementAccountUpdateHistoryRequest))
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

func _Cost_ImportCurFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportCurFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).ImportCurFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/ImportCurFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).ImportCurFiles(ctx, req.(*ImportCurFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cost_CalculateCosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateCostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).CalculateCosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/CalculateCosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).CalculateCosts(ctx, req.(*CalculateCostsRequest))
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

func _Cost_CreateBudgetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBudgetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).CreateBudgetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/CreateBudgetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).CreateBudgetConfig(ctx, req.(*CreateBudgetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cost_UpdateBudgetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBudgetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).UpdateBudgetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/UpdateBudgetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).UpdateBudgetConfig(ctx, req.(*UpdateBudgetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cost_GetBudgetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBudgetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).GetBudgetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/GetBudgetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).GetBudgetConfig(ctx, req.(*GetBudgetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cost_GetForecasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetForecastsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).GetForecasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/GetForecasts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).GetForecasts(ctx, req.(*GetForecastsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cost_GetMonthToDateForecast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMonthToDateForecastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostServer).GetMonthToDateForecast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.cost.v1.Cost/GetMonthToDateForecast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostServer).GetMonthToDateForecast(ctx, req.(*GetMonthToDateForecastRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "GetManagementAccountUpdateHistory",
			Handler:    _Cost_GetManagementAccountUpdateHistory_Handler,
		},
		{
			MethodName: "CreateManagementAccount",
			Handler:    _Cost_CreateManagementAccount_Handler,
		},
		{
			MethodName: "DeleteManagementAccount",
			Handler:    _Cost_DeleteManagementAccount_Handler,
		},
		{
			MethodName: "ImportCurFiles",
			Handler:    _Cost_ImportCurFiles_Handler,
		},
		{
			MethodName: "CalculateCosts",
			Handler:    _Cost_CalculateCosts_Handler,
		},
		{
			MethodName: "CreateBudgetConfig",
			Handler:    _Cost_CreateBudgetConfig_Handler,
		},
		{
			MethodName: "UpdateBudgetConfig",
			Handler:    _Cost_UpdateBudgetConfig_Handler,
		},
		{
			MethodName: "GetBudgetConfig",
			Handler:    _Cost_GetBudgetConfig_Handler,
		},
		{
			MethodName: "GetForecasts",
			Handler:    _Cost_GetForecasts_Handler,
		},
		{
			MethodName: "GetMonthToDateForecast",
			Handler:    _Cost_GetMonthToDateForecast_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadCosts",
			Handler:       _Cost_ReadCosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadAdjustments",
			Handler:       _Cost_ReadAdjustments_Handler,
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
