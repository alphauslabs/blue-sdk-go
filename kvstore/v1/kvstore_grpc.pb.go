// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kvstore

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KvStoreClient is the client API for KvStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KvStoreClient interface {
	// WORK-IN-PROGRESS. Scans keys from your store.
	Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (KvStore_ScanClient, error)
	// WORK-IN-PROGRESS. Reads a key from your store.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*KeyValue, error)
	// WORK-IN-PROGRESS. Writes a new (or update an existing) key:value data in your store.
	Write(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// WORK-IN-PROGRESS. Deletes a key from your store. Using a `-` (hyphen) as {key} input
	// translates to all keys to be deleted.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type kvStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKvStoreClient(cc grpc.ClientConnInterface) KvStoreClient {
	return &kvStoreClient{cc}
}

func (c *kvStoreClient) Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (KvStore_ScanClient, error) {
	stream, err := c.cc.NewStream(ctx, &KvStore_ServiceDesc.Streams[0], "/blueapi.kvstore.v1.KvStore/Scan", opts...)
	if err != nil {
		return nil, err
	}
	x := &kvStoreScanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KvStore_ScanClient interface {
	Recv() (*ScanResponse, error)
	grpc.ClientStream
}

type kvStoreScanClient struct {
	grpc.ClientStream
}

func (x *kvStoreScanClient) Recv() (*ScanResponse, error) {
	m := new(ScanResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kvStoreClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*KeyValue, error) {
	out := new(KeyValue)
	err := c.cc.Invoke(ctx, "/blueapi.kvstore.v1.KvStore/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvStoreClient) Write(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.kvstore.v1.KvStore/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blueapi.kvstore.v1.KvStore/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KvStoreServer is the server API for KvStore service.
// All implementations must embed UnimplementedKvStoreServer
// for forward compatibility
type KvStoreServer interface {
	// WORK-IN-PROGRESS. Scans keys from your store.
	Scan(*ScanRequest, KvStore_ScanServer) error
	// WORK-IN-PROGRESS. Reads a key from your store.
	Read(context.Context, *ReadRequest) (*KeyValue, error)
	// WORK-IN-PROGRESS. Writes a new (or update an existing) key:value data in your store.
	Write(context.Context, *KeyValue) (*emptypb.Empty, error)
	// WORK-IN-PROGRESS. Deletes a key from your store. Using a `-` (hyphen) as {key} input
	// translates to all keys to be deleted.
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedKvStoreServer()
}

// UnimplementedKvStoreServer must be embedded to have forward compatible implementations.
type UnimplementedKvStoreServer struct {
}

func (UnimplementedKvStoreServer) Scan(*ScanRequest, KvStore_ScanServer) error {
	return status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (UnimplementedKvStoreServer) Read(context.Context, *ReadRequest) (*KeyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedKvStoreServer) Write(context.Context, *KeyValue) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedKvStoreServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKvStoreServer) mustEmbedUnimplementedKvStoreServer() {}

// UnsafeKvStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KvStoreServer will
// result in compilation errors.
type UnsafeKvStoreServer interface {
	mustEmbedUnimplementedKvStoreServer()
}

func RegisterKvStoreServer(s grpc.ServiceRegistrar, srv KvStoreServer) {
	s.RegisterService(&KvStore_ServiceDesc, srv)
}

func _KvStore_Scan_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ScanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KvStoreServer).Scan(m, &kvStoreScanServer{stream})
}

type KvStore_ScanServer interface {
	Send(*ScanResponse) error
	grpc.ServerStream
}

type kvStoreScanServer struct {
	grpc.ServerStream
}

func (x *kvStoreScanServer) Send(m *ScanResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KvStore_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvStoreServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.kvstore.v1.KvStore/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvStoreServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KvStore_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvStoreServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.kvstore.v1.KvStore/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvStoreServer).Write(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _KvStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blueapi.kvstore.v1.KvStore/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvStoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KvStore_ServiceDesc is the grpc.ServiceDesc for KvStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KvStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.kvstore.v1.KvStore",
	HandlerType: (*KvStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _KvStore_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _KvStore_Write_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KvStore_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Scan",
			Handler:       _KvStore_Scan_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kvstore/v1/kvstore.proto",
}
