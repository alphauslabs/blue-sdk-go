// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: billing/v1/billing.proto

package billing

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
	// Reads the daily calculation history of all accounts in your billing groups. Only available in Ripple.
	ListAwsDailyRunHistory(ctx context.Context, in *ListAwsDailyRunHistoryRequest, opts ...grpc.CallOption) (Billing_ListAwsDailyRunHistoryClient, error)
	// Returns a list of accounts that have been updated after invoice along with the differences in costs, if any. Only available in Ripple.
	ListUsageCostsDrift(ctx context.Context, in *ListUsageCostsDriftRequest, opts ...grpc.CallOption) (Billing_ListUsageCostsDriftClient, error)
	// WORK-IN-PROGRESS: Creates an invoice. Only available in Ripple.
	CreateInvoice(ctx context.Context, in *CreateInvoiceRequest, opts ...grpc.CallOption) (*api.InvoiceMessage, error)
	// WORK-IN-PROGRESS: Gets an invoice. Only available in Ripple.
	GetInvoiceStatus(ctx context.Context, in *GetInvoiceStatusRequest, opts ...grpc.CallOption) (*api.InvoiceMessage, error)
	// Gets an invoice.
	GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*api.Invoice, error)
	// Updates an invoice preview. Only available in Ripple.
	UpdateInvoicePreviews(ctx context.Context, in *UpdateInvoicePreviewsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Exports an invoice.
	ExportInvoiceFile(ctx context.Context, in *ExportInvoiceFileRequest, opts ...grpc.CallOption) (*ExportInvoiceFileResponse, error)
	// Reads the invoice service discounts. Only available in Ripple.
	ListInvoiceServiceDiscounts(ctx context.Context, in *ListInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (Billing_ListInvoiceServiceDiscountsClient, error)
	// Reads the account invoice service discounts. Only available in Ripple.
	ListAccountInvoiceServiceDiscounts(ctx context.Context, in *ListAccountInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (Billing_ListAccountInvoiceServiceDiscountsClient, error)
	// Registers the account invoice service discounts. Only available in Ripple.
	CreateAccountInvoiceServiceDiscounts(ctx context.Context, in *CreateAccountInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (*CreateAccountInvoiceServiceDiscountsResponse, error)
	// Updates the account invoice service discounts. Only available in Ripple.
	UpdateAccountInvoiceServiceDiscounts(ctx context.Context, in *UpdateAccountInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (*UpdateAccountInvoiceServiceDiscountsResponse, error)
	// Removes the account invoice service discounts. Only available in Ripple.
	RemoveAccountInvoiceServiceDiscounts(ctx context.Context, in *RemoveAccountInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deletes the account invoice service discounts. Only available in Ripple.
	DeleteAccountInvoiceServiceDiscounts(ctx context.Context, in *DeleteAccountInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
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

func (c *billingClient) ListAwsDailyRunHistory(ctx context.Context, in *ListAwsDailyRunHistoryRequest, opts ...grpc.CallOption) (Billing_ListAwsDailyRunHistoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Billing_ServiceDesc.Streams[1], "/blueapi.billing.v1.Billing/ListAwsDailyRunHistory", opts...)
	if err != nil {
		return nil, err
	}
	x := &billingListAwsDailyRunHistoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Billing_ListAwsDailyRunHistoryClient interface {
	Recv() (*AwsDailyRunHistory, error)
	grpc.ClientStream
}

type billingListAwsDailyRunHistoryClient struct {
	grpc.ClientStream
}

func (x *billingListAwsDailyRunHistoryClient) Recv() (*AwsDailyRunHistory, error) {
	m := new(AwsDailyRunHistory)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *billingClient) ListUsageCostsDrift(ctx context.Context, in *ListUsageCostsDriftRequest, opts ...grpc.CallOption) (Billing_ListUsageCostsDriftClient, error) {
	stream, err := c.cc.NewStream(ctx, &Billing_ServiceDesc.Streams[2], "/blueapi.billing.v1.Billing/ListUsageCostsDrift", opts...)
	if err != nil {
		return nil, err
	}
	x := &billingListUsageCostsDriftClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Billing_ListUsageCostsDriftClient interface {
	Recv() (*UsageCostsDrift, error)
	grpc.ClientStream
}

type billingListUsageCostsDriftClient struct {
	grpc.ClientStream
}

func (x *billingListUsageCostsDriftClient) Recv() (*UsageCostsDrift, error) {
	m := new(UsageCostsDrift)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *billingClient) CreateInvoice(ctx context.Context, in *CreateInvoiceRequest, opts ...grpc.CallOption) (*api.InvoiceMessage, error) {
	out := new(api.InvoiceMessage)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/CreateInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) GetInvoiceStatus(ctx context.Context, in *GetInvoiceStatusRequest, opts ...grpc.CallOption) (*api.InvoiceMessage, error) {
	out := new(api.InvoiceMessage)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/GetInvoiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*api.Invoice, error) {
	out := new(api.Invoice)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/GetInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) UpdateInvoicePreviews(ctx context.Context, in *UpdateInvoicePreviewsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/UpdateInvoicePreviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) ExportInvoiceFile(ctx context.Context, in *ExportInvoiceFileRequest, opts ...grpc.CallOption) (*ExportInvoiceFileResponse, error) {
	out := new(ExportInvoiceFileResponse)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/ExportInvoiceFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) ListInvoiceServiceDiscounts(ctx context.Context, in *ListInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (Billing_ListInvoiceServiceDiscountsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Billing_ServiceDesc.Streams[3], "/blueapi.billing.v1.Billing/ListInvoiceServiceDiscounts", opts...)
	if err != nil {
		return nil, err
	}
	x := &billingListInvoiceServiceDiscountsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Billing_ListInvoiceServiceDiscountsClient interface {
	Recv() (*InvoiceServiceDiscounts, error)
	grpc.ClientStream
}

type billingListInvoiceServiceDiscountsClient struct {
	grpc.ClientStream
}

func (x *billingListInvoiceServiceDiscountsClient) Recv() (*InvoiceServiceDiscounts, error) {
	m := new(InvoiceServiceDiscounts)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *billingClient) ListAccountInvoiceServiceDiscounts(ctx context.Context, in *ListAccountInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (Billing_ListAccountInvoiceServiceDiscountsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Billing_ServiceDesc.Streams[4], "/blueapi.billing.v1.Billing/ListAccountInvoiceServiceDiscounts", opts...)
	if err != nil {
		return nil, err
	}
	x := &billingListAccountInvoiceServiceDiscountsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Billing_ListAccountInvoiceServiceDiscountsClient interface {
	Recv() (*AccountInvoiceServiceDiscounts, error)
	grpc.ClientStream
}

type billingListAccountInvoiceServiceDiscountsClient struct {
	grpc.ClientStream
}

func (x *billingListAccountInvoiceServiceDiscountsClient) Recv() (*AccountInvoiceServiceDiscounts, error) {
	m := new(AccountInvoiceServiceDiscounts)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *billingClient) CreateAccountInvoiceServiceDiscounts(ctx context.Context, in *CreateAccountInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (*CreateAccountInvoiceServiceDiscountsResponse, error) {
	out := new(CreateAccountInvoiceServiceDiscountsResponse)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/CreateAccountInvoiceServiceDiscounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) UpdateAccountInvoiceServiceDiscounts(ctx context.Context, in *UpdateAccountInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (*UpdateAccountInvoiceServiceDiscountsResponse, error) {
	out := new(UpdateAccountInvoiceServiceDiscountsResponse)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/UpdateAccountInvoiceServiceDiscounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) RemoveAccountInvoiceServiceDiscounts(ctx context.Context, in *RemoveAccountInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/RemoveAccountInvoiceServiceDiscounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) DeleteAccountInvoiceServiceDiscounts(ctx context.Context, in *DeleteAccountInvoiceServiceDiscountsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.billing.v1.Billing/DeleteAccountInvoiceServiceDiscounts", in, out, opts...)
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
	// Reads the daily calculation history of all accounts in your billing groups. Only available in Ripple.
	ListAwsDailyRunHistory(*ListAwsDailyRunHistoryRequest, Billing_ListAwsDailyRunHistoryServer) error
	// Returns a list of accounts that have been updated after invoice along with the differences in costs, if any. Only available in Ripple.
	ListUsageCostsDrift(*ListUsageCostsDriftRequest, Billing_ListUsageCostsDriftServer) error
	// WORK-IN-PROGRESS: Creates an invoice. Only available in Ripple.
	CreateInvoice(context.Context, *CreateInvoiceRequest) (*api.InvoiceMessage, error)
	// WORK-IN-PROGRESS: Gets an invoice. Only available in Ripple.
	GetInvoiceStatus(context.Context, *GetInvoiceStatusRequest) (*api.InvoiceMessage, error)
	// Gets an invoice.
	GetInvoice(context.Context, *GetInvoiceRequest) (*api.Invoice, error)
	// Updates an invoice preview. Only available in Ripple.
	UpdateInvoicePreviews(context.Context, *UpdateInvoicePreviewsRequest) (*emptypb.Empty, error)
	// Exports an invoice.
	ExportInvoiceFile(context.Context, *ExportInvoiceFileRequest) (*ExportInvoiceFileResponse, error)
	// Reads the invoice service discounts. Only available in Ripple.
	ListInvoiceServiceDiscounts(*ListInvoiceServiceDiscountsRequest, Billing_ListInvoiceServiceDiscountsServer) error
	// Reads the account invoice service discounts. Only available in Ripple.
	ListAccountInvoiceServiceDiscounts(*ListAccountInvoiceServiceDiscountsRequest, Billing_ListAccountInvoiceServiceDiscountsServer) error
	// Registers the account invoice service discounts. Only available in Ripple.
	CreateAccountInvoiceServiceDiscounts(context.Context, *CreateAccountInvoiceServiceDiscountsRequest) (*CreateAccountInvoiceServiceDiscountsResponse, error)
	// Updates the account invoice service discounts. Only available in Ripple.
	UpdateAccountInvoiceServiceDiscounts(context.Context, *UpdateAccountInvoiceServiceDiscountsRequest) (*UpdateAccountInvoiceServiceDiscountsResponse, error)
	// Removes the account invoice service discounts. Only available in Ripple.
	RemoveAccountInvoiceServiceDiscounts(context.Context, *RemoveAccountInvoiceServiceDiscountsRequest) (*emptypb.Empty, error)
	// Deletes the account invoice service discounts. Only available in Ripple.
	DeleteAccountInvoiceServiceDiscounts(context.Context, *DeleteAccountInvoiceServiceDiscountsRequest) (*emptypb.Empty, error)
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
func (UnimplementedBillingServer) ListAwsDailyRunHistory(*ListAwsDailyRunHistoryRequest, Billing_ListAwsDailyRunHistoryServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAwsDailyRunHistory not implemented")
}
func (UnimplementedBillingServer) ListUsageCostsDrift(*ListUsageCostsDriftRequest, Billing_ListUsageCostsDriftServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUsageCostsDrift not implemented")
}
func (UnimplementedBillingServer) CreateInvoice(context.Context, *CreateInvoiceRequest) (*api.InvoiceMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoice not implemented")
}
func (UnimplementedBillingServer) GetInvoiceStatus(context.Context, *GetInvoiceStatusRequest) (*api.InvoiceMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoiceStatus not implemented")
}
func (UnimplementedBillingServer) GetInvoice(context.Context, *GetInvoiceRequest) (*api.Invoice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedBillingServer) UpdateInvoicePreviews(context.Context, *UpdateInvoicePreviewsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInvoicePreviews not implemented")
}
func (UnimplementedBillingServer) ExportInvoiceFile(context.Context, *ExportInvoiceFileRequest) (*ExportInvoiceFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportInvoiceFile not implemented")
}
func (UnimplementedBillingServer) ListInvoiceServiceDiscounts(*ListInvoiceServiceDiscountsRequest, Billing_ListInvoiceServiceDiscountsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListInvoiceServiceDiscounts not implemented")
}
func (UnimplementedBillingServer) ListAccountInvoiceServiceDiscounts(*ListAccountInvoiceServiceDiscountsRequest, Billing_ListAccountInvoiceServiceDiscountsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAccountInvoiceServiceDiscounts not implemented")
}
func (UnimplementedBillingServer) CreateAccountInvoiceServiceDiscounts(context.Context, *CreateAccountInvoiceServiceDiscountsRequest) (*CreateAccountInvoiceServiceDiscountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountInvoiceServiceDiscounts not implemented")
}
func (UnimplementedBillingServer) UpdateAccountInvoiceServiceDiscounts(context.Context, *UpdateAccountInvoiceServiceDiscountsRequest) (*UpdateAccountInvoiceServiceDiscountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountInvoiceServiceDiscounts not implemented")
}
func (UnimplementedBillingServer) RemoveAccountInvoiceServiceDiscounts(context.Context, *RemoveAccountInvoiceServiceDiscountsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAccountInvoiceServiceDiscounts not implemented")
}
func (UnimplementedBillingServer) DeleteAccountInvoiceServiceDiscounts(context.Context, *DeleteAccountInvoiceServiceDiscountsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccountInvoiceServiceDiscounts not implemented")
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

func _Billing_ListAwsDailyRunHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAwsDailyRunHistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BillingServer).ListAwsDailyRunHistory(m, &billingListAwsDailyRunHistoryServer{stream})
}

type Billing_ListAwsDailyRunHistoryServer interface {
	Send(*AwsDailyRunHistory) error
	grpc.ServerStream
}

type billingListAwsDailyRunHistoryServer struct {
	grpc.ServerStream
}

func (x *billingListAwsDailyRunHistoryServer) Send(m *AwsDailyRunHistory) error {
	return x.ServerStream.SendMsg(m)
}

func _Billing_ListUsageCostsDrift_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUsageCostsDriftRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BillingServer).ListUsageCostsDrift(m, &billingListUsageCostsDriftServer{stream})
}

type Billing_ListUsageCostsDriftServer interface {
	Send(*UsageCostsDrift) error
	grpc.ServerStream
}

type billingListUsageCostsDriftServer struct {
	grpc.ServerStream
}

func (x *billingListUsageCostsDriftServer) Send(m *UsageCostsDrift) error {
	return x.ServerStream.SendMsg(m)
}

func _Billing_CreateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).CreateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/CreateInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).CreateInvoice(ctx, req.(*CreateInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_GetInvoiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GetInvoiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/GetInvoiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GetInvoiceStatus(ctx, req.(*GetInvoiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/GetInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GetInvoice(ctx, req.(*GetInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_UpdateInvoicePreviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInvoicePreviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).UpdateInvoicePreviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/UpdateInvoicePreviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).UpdateInvoicePreviews(ctx, req.(*UpdateInvoicePreviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_ExportInvoiceFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportInvoiceFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).ExportInvoiceFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/ExportInvoiceFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).ExportInvoiceFile(ctx, req.(*ExportInvoiceFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_ListInvoiceServiceDiscounts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListInvoiceServiceDiscountsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BillingServer).ListInvoiceServiceDiscounts(m, &billingListInvoiceServiceDiscountsServer{stream})
}

type Billing_ListInvoiceServiceDiscountsServer interface {
	Send(*InvoiceServiceDiscounts) error
	grpc.ServerStream
}

type billingListInvoiceServiceDiscountsServer struct {
	grpc.ServerStream
}

func (x *billingListInvoiceServiceDiscountsServer) Send(m *InvoiceServiceDiscounts) error {
	return x.ServerStream.SendMsg(m)
}

func _Billing_ListAccountInvoiceServiceDiscounts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAccountInvoiceServiceDiscountsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BillingServer).ListAccountInvoiceServiceDiscounts(m, &billingListAccountInvoiceServiceDiscountsServer{stream})
}

type Billing_ListAccountInvoiceServiceDiscountsServer interface {
	Send(*AccountInvoiceServiceDiscounts) error
	grpc.ServerStream
}

type billingListAccountInvoiceServiceDiscountsServer struct {
	grpc.ServerStream
}

func (x *billingListAccountInvoiceServiceDiscountsServer) Send(m *AccountInvoiceServiceDiscounts) error {
	return x.ServerStream.SendMsg(m)
}

func _Billing_CreateAccountInvoiceServiceDiscounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountInvoiceServiceDiscountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).CreateAccountInvoiceServiceDiscounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/CreateAccountInvoiceServiceDiscounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).CreateAccountInvoiceServiceDiscounts(ctx, req.(*CreateAccountInvoiceServiceDiscountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_UpdateAccountInvoiceServiceDiscounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountInvoiceServiceDiscountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).UpdateAccountInvoiceServiceDiscounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/UpdateAccountInvoiceServiceDiscounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).UpdateAccountInvoiceServiceDiscounts(ctx, req.(*UpdateAccountInvoiceServiceDiscountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_RemoveAccountInvoiceServiceDiscounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAccountInvoiceServiceDiscountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).RemoveAccountInvoiceServiceDiscounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/RemoveAccountInvoiceServiceDiscounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).RemoveAccountInvoiceServiceDiscounts(ctx, req.(*RemoveAccountInvoiceServiceDiscountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_DeleteAccountInvoiceServiceDiscounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountInvoiceServiceDiscountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).DeleteAccountInvoiceServiceDiscounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.billing.v1.Billing/DeleteAccountInvoiceServiceDiscounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).DeleteAccountInvoiceServiceDiscounts(ctx, req.(*DeleteAccountInvoiceServiceDiscountsRequest))
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
		{
			MethodName: "CreateInvoice",
			Handler:    _Billing_CreateInvoice_Handler,
		},
		{
			MethodName: "GetInvoiceStatus",
			Handler:    _Billing_GetInvoiceStatus_Handler,
		},
		{
			MethodName: "GetInvoice",
			Handler:    _Billing_GetInvoice_Handler,
		},
		{
			MethodName: "UpdateInvoicePreviews",
			Handler:    _Billing_UpdateInvoicePreviews_Handler,
		},
		{
			MethodName: "ExportInvoiceFile",
			Handler:    _Billing_ExportInvoiceFile_Handler,
		},
		{
			MethodName: "CreateAccountInvoiceServiceDiscounts",
			Handler:    _Billing_CreateAccountInvoiceServiceDiscounts_Handler,
		},
		{
			MethodName: "UpdateAccountInvoiceServiceDiscounts",
			Handler:    _Billing_UpdateAccountInvoiceServiceDiscounts_Handler,
		},
		{
			MethodName: "RemoveAccountInvoiceServiceDiscounts",
			Handler:    _Billing_RemoveAccountInvoiceServiceDiscounts_Handler,
		},
		{
			MethodName: "DeleteAccountInvoiceServiceDiscounts",
			Handler:    _Billing_DeleteAccountInvoiceServiceDiscounts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBillingGroups",
			Handler:       _Billing_ListBillingGroups_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListAwsDailyRunHistory",
			Handler:       _Billing_ListAwsDailyRunHistory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListUsageCostsDrift",
			Handler:       _Billing_ListUsageCostsDrift_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListInvoiceServiceDiscounts",
			Handler:       _Billing_ListInvoiceServiceDiscounts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListAccountInvoiceServiceDiscounts",
			Handler:       _Billing_ListAccountInvoiceServiceDiscounts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "billing/v1/billing.proto",
}
