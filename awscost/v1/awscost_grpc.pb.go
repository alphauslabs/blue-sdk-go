// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package awscost

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

// AwsCostClient is the client API for AwsCost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AwsCostClient interface {
	// Streams back the cost details of an account. If date range parameters are not set,
	// month-to-date (current month) will be returned.
	StreamReadAccountCosts(ctx context.Context, in *StreamReadAccountCostsRequest, opts ...grpc.CallOption) (AwsCost_StreamReadAccountCostsClient, error)
}

type awsCostClient struct {
	cc grpc.ClientConnInterface
}

func NewAwsCostClient(cc grpc.ClientConnInterface) AwsCostClient {
	return &awsCostClient{cc}
}

func (c *awsCostClient) StreamReadAccountCosts(ctx context.Context, in *StreamReadAccountCostsRequest, opts ...grpc.CallOption) (AwsCost_StreamReadAccountCostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AwsCost_ServiceDesc.Streams[0], "/blueapi.awscost.v1.AwsCost/StreamReadAccountCosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &awsCostStreamReadAccountCostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AwsCost_StreamReadAccountCostsClient interface {
	Recv() (*Cost, error)
	grpc.ClientStream
}

type awsCostStreamReadAccountCostsClient struct {
	grpc.ClientStream
}

func (x *awsCostStreamReadAccountCostsClient) Recv() (*Cost, error) {
	m := new(Cost)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AwsCostServer is the server API for AwsCost service.
// All implementations must embed UnimplementedAwsCostServer
// for forward compatibility
type AwsCostServer interface {
	// Streams back the cost details of an account. If date range parameters are not set,
	// month-to-date (current month) will be returned.
	StreamReadAccountCosts(*StreamReadAccountCostsRequest, AwsCost_StreamReadAccountCostsServer) error
	mustEmbedUnimplementedAwsCostServer()
}

// UnimplementedAwsCostServer must be embedded to have forward compatible implementations.
type UnimplementedAwsCostServer struct {
}

func (UnimplementedAwsCostServer) StreamReadAccountCosts(*StreamReadAccountCostsRequest, AwsCost_StreamReadAccountCostsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamReadAccountCosts not implemented")
}
func (UnimplementedAwsCostServer) mustEmbedUnimplementedAwsCostServer() {}

// UnsafeAwsCostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AwsCostServer will
// result in compilation errors.
type UnsafeAwsCostServer interface {
	mustEmbedUnimplementedAwsCostServer()
}

func RegisterAwsCostServer(s grpc.ServiceRegistrar, srv AwsCostServer) {
	s.RegisterService(&AwsCost_ServiceDesc, srv)
}

func _AwsCost_StreamReadAccountCosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamReadAccountCostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AwsCostServer).StreamReadAccountCosts(m, &awsCostStreamReadAccountCostsServer{stream})
}

type AwsCost_StreamReadAccountCostsServer interface {
	Send(*Cost) error
	grpc.ServerStream
}

type awsCostStreamReadAccountCostsServer struct {
	grpc.ServerStream
}

func (x *awsCostStreamReadAccountCostsServer) Send(m *Cost) error {
	return x.ServerStream.SendMsg(m)
}

// AwsCost_ServiceDesc is the grpc.ServiceDesc for AwsCost service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AwsCost_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.awscost.v1.AwsCost",
	HandlerType: (*AwsCostServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamReadAccountCosts",
			Handler:       _AwsCost_StreamReadAccountCosts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "awscost/v1/awscost.proto",
}
