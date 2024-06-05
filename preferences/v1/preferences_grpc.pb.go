// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: preferences/v1/preferences.proto

package preferences

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Preferences_GetPreferences_FullMethodName = "/blueapi.preferences.v1.Preferences/GetPreferences"
)

// PreferencesClient is the client API for Preferences service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Preferences service definition.
type PreferencesClient interface {
	// WORK-IN-PROGRESS: Gets current preferences.
	GetPreferences(ctx context.Context, in *GetPreferencesRequest, opts ...grpc.CallOption) (*Preference, error)
}

type preferencesClient struct {
	cc grpc.ClientConnInterface
}

func NewPreferencesClient(cc grpc.ClientConnInterface) PreferencesClient {
	return &preferencesClient{cc}
}

func (c *preferencesClient) GetPreferences(ctx context.Context, in *GetPreferencesRequest, opts ...grpc.CallOption) (*Preference, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Preference)
	err := c.cc.Invoke(ctx, Preferences_GetPreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PreferencesServer is the server API for Preferences service.
// All implementations must embed UnimplementedPreferencesServer
// for forward compatibility
//
// Preferences service definition.
type PreferencesServer interface {
	// WORK-IN-PROGRESS: Gets current preferences.
	GetPreferences(context.Context, *GetPreferencesRequest) (*Preference, error)
	mustEmbedUnimplementedPreferencesServer()
}

// UnimplementedPreferencesServer must be embedded to have forward compatible implementations.
type UnimplementedPreferencesServer struct {
}

func (UnimplementedPreferencesServer) GetPreferences(context.Context, *GetPreferencesRequest) (*Preference, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreferences not implemented")
}
func (UnimplementedPreferencesServer) mustEmbedUnimplementedPreferencesServer() {}

// UnsafePreferencesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PreferencesServer will
// result in compilation errors.
type UnsafePreferencesServer interface {
	mustEmbedUnimplementedPreferencesServer()
}

func RegisterPreferencesServer(s grpc.ServiceRegistrar, srv PreferencesServer) {
	s.RegisterService(&Preferences_ServiceDesc, srv)
}

func _Preferences_GetPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreferencesServer).GetPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Preferences_GetPreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreferencesServer).GetPreferences(ctx, req.(*GetPreferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Preferences_ServiceDesc is the grpc.ServiceDesc for Preferences service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Preferences_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blueapi.preferences.v1.Preferences",
	HandlerType: (*PreferencesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPreferences",
			Handler:    _Preferences_GetPreferences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "preferences/v1/preferences.proto",
}
