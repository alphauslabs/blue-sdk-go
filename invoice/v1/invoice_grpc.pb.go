// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: invoice/v1/invoice.proto

package invoice

import (
	context "context"
	api "github.com/alphauslabs/blue-sdk-go/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InvoiceClient is the client API for Invoice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoiceClient interface {
	// Gets a invoice.
	GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*api.Invoice, error)
	// Exports a invoice.
	ExportInvoiceFile(ctx context.Context, in *ExportInvoiceFileRequest, opts ...grpc.CallOption) (*ExportInvoiceFileResponse, error)
}

type invoiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceClient(cc grpc.ClientConnInterface) InvoiceClient {
	return &invoiceClient{cc}
}

func (c *invoiceClient) GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*api.Invoice, error) {
	out := new(api.Invoice)
	err := c.cc.Invoke(ctx, "/blueapi.invoice.v1.Invoice/GetInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceClient) ExportInvoiceFile(ctx context.Context, in *ExportInvoiceFileRequest, opts ...grpc.CallOption) (*ExportInvoiceFileResponse, error) {
	out := new(ExportInvoiceFileResponse)
	err := c.cc.Invoke(ctx, "/blueapi.invoice.v1.Invoice/ExportInvoiceFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServer is the server API for Invoice service.
// All implementations must embed UnimplementedInvoiceServer
// for forward compatibility
type InvoiceServer interface {
	// Gets a invoice.
	GetInvoice(context.Context, *GetInvoiceRequest) (*api.Invoice, error)
	// Exports a invoice.
	ExportInvoiceFile(context.Context, *ExportInvoiceFileRequest) (*ExportInvoiceFileResponse, error)
	mustEmbedUnimplementedInvoiceServer()
}

// UnimplementedInvoiceServer must be embedded to have forward compatible implementations.
type UnimplementedInvoiceServer struct {
}

func (UnimplementedInvoiceServer) GetInvoice(context.Context, *GetInvoiceRequest) (*api.Invoice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedInvoiceServer) ExportInvoiceFile(context.Context, *ExportInvoiceFileRequest) (*ExportInvoiceFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportInvoiceFile not implemented")
}
func (UnimplementedInvoiceServer) mustEmbedUnimplementedInvoiceServer() {}

// UnsafeInvoiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoiceServer will
// result in compilation errors.
type UnsafeInvoiceServer interface {
	mustEmbedUnimplementedInvoiceServer()
}

func RegisterInvoiceServer(s grpc.ServiceRegistrar, srv InvoiceServer) {
	s.RegisterService(&Invoice_ServiceDesc, srv)
}

func _Invoice_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.invoice.v1.Invoice/GetInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServer).GetInvoice(ctx, req.(*GetInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoice_ExportInvoiceFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportInvoiceFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServer).ExportInvoiceFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.invoice.v1.Invoice/ExportInvoiceFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServer).ExportInvoiceFile(ctx, req.(*ExportInvoiceFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Invoice_ServiceDesc is the grpc.ServiceDesc for Invoice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invoice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.invoice.v1.Invoice",
	HandlerType: (*InvoiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInvoice",
			Handler:    _Invoice_GetInvoice_Handler,
		},
		{
			MethodName: "ExportInvoiceFile",
			Handler:    _Invoice_ExportInvoiceFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoice/v1/invoice.proto",
}
