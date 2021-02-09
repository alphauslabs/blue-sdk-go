// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: blue/curmgt/v1/curmgt.proto

package curmgt

import (
	common "github.com/alphauslabs/blue-sdk-go/common"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for CurMgt.ListAccounts.
type ListAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAccountsRequest) Reset() {
	*x = ListAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blue_curmgt_v1_curmgt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountsRequest) ProtoMessage() {}

func (x *ListAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blue_curmgt_v1_curmgt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountsRequest.ProtoReflect.Descriptor instead.
func (*ListAccountsRequest) Descriptor() ([]byte, []int) {
	return file_blue_curmgt_v1_curmgt_proto_rawDescGZIP(), []int{0}
}

// Response message for CurMgt.ListAccounts.
type ListAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*common.AwsPayerAccount `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *ListAccountsResponse) Reset() {
	*x = ListAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blue_curmgt_v1_curmgt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountsResponse) ProtoMessage() {}

func (x *ListAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blue_curmgt_v1_curmgt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountsResponse.ProtoReflect.Descriptor instead.
func (*ListAccountsResponse) Descriptor() ([]byte, []int) {
	return file_blue_curmgt_v1_curmgt_proto_rawDescGZIP(), []int{1}
}

func (x *ListAccountsResponse) GetAccounts() []*common.AwsPayerAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_blue_curmgt_v1_curmgt_proto protoreflect.FileDescriptor

var file_blue_curmgt_v1_curmgt_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x6c, 0x75, 0x65, 0x2f, 0x63, 0x75, 0x72, 0x6d, 0x67, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x72, 0x6d, 0x67, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62,
	0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x2e, 0x63, 0x75, 0x72, 0x6d,
	0x67, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x53, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x77, 0x73, 0x50, 0x61,
	0x79, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x32, 0x91, 0x01, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x4d, 0x67, 0x74, 0x12,
	0x86, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x12, 0x2b, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x2e,
	0x63, 0x75, 0x72, 0x6d, 0x67, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x2e, 0x63, 0x75, 0x72,
	0x6d, 0x67, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x6d, 0x67, 0x74, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0xb4, 0x03, 0x0a, 0x1d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x6c, 0x75, 0x65, 0x2e, 0x63, 0x75, 0x72, 0x6d, 0x67, 0x74, 0x42, 0x11, 0x42, 0x6c, 0x75, 0x65,
	0x43, 0x75, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75,
	0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c,
	0x75, 0x65, 0x2f, 0x63, 0x75, 0x72, 0x6d, 0x67, 0x74, 0x92, 0x41, 0xd2, 0x02, 0x12, 0xe4, 0x01,
	0x0a, 0x24, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x20, 0x42, 0x6c, 0x75, 0x65, 0x3a, 0x20,
	0x43, 0x55, 0x52, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5e, 0x0a, 0x24, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x75,
	0x73, 0x20, 0x42, 0x6c, 0x75, 0x65, 0x3a, 0x20, 0x43, 0x55, 0x52, 0x20, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c,
	0x61, 0x62, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2f, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x6d, 0x67,
	0x74, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x2a, 0x57, 0x0a, 0x1b, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x3a, 0x20, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x20, 0x32, 0x2e, 0x30, 0x12, 0x38, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75,
	0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c,
	0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x72, 0x69, 0x0a, 0x2f, 0x4d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f,
	0x75, 0x74, 0x20, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x20, 0x42, 0x6c, 0x75, 0x65, 0x3a,
	0x20, 0x43, 0x55, 0x52, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x6d, 0x67, 0x74, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blue_curmgt_v1_curmgt_proto_rawDescOnce sync.Once
	file_blue_curmgt_v1_curmgt_proto_rawDescData = file_blue_curmgt_v1_curmgt_proto_rawDesc
)

func file_blue_curmgt_v1_curmgt_proto_rawDescGZIP() []byte {
	file_blue_curmgt_v1_curmgt_proto_rawDescOnce.Do(func() {
		file_blue_curmgt_v1_curmgt_proto_rawDescData = protoimpl.X.CompressGZIP(file_blue_curmgt_v1_curmgt_proto_rawDescData)
	})
	return file_blue_curmgt_v1_curmgt_proto_rawDescData
}

var file_blue_curmgt_v1_curmgt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_blue_curmgt_v1_curmgt_proto_goTypes = []interface{}{
	(*ListAccountsRequest)(nil),    // 0: blueapi.blue.curmgt.v1.ListAccountsRequest
	(*ListAccountsResponse)(nil),   // 1: blueapi.blue.curmgt.v1.ListAccountsResponse
	(*common.AwsPayerAccount)(nil), // 2: blueapi.common.AwsPayerAccount
}
var file_blue_curmgt_v1_curmgt_proto_depIdxs = []int32{
	2, // 0: blueapi.blue.curmgt.v1.ListAccountsResponse.accounts:type_name -> blueapi.common.AwsPayerAccount
	0, // 1: blueapi.blue.curmgt.v1.CurMgt.ListAccounts:input_type -> blueapi.blue.curmgt.v1.ListAccountsRequest
	1, // 2: blueapi.blue.curmgt.v1.CurMgt.ListAccounts:output_type -> blueapi.blue.curmgt.v1.ListAccountsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blue_curmgt_v1_curmgt_proto_init() }
func file_blue_curmgt_v1_curmgt_proto_init() {
	if File_blue_curmgt_v1_curmgt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blue_curmgt_v1_curmgt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blue_curmgt_v1_curmgt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blue_curmgt_v1_curmgt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blue_curmgt_v1_curmgt_proto_goTypes,
		DependencyIndexes: file_blue_curmgt_v1_curmgt_proto_depIdxs,
		MessageInfos:      file_blue_curmgt_v1_curmgt_proto_msgTypes,
	}.Build()
	File_blue_curmgt_v1_curmgt_proto = out.File
	file_blue_curmgt_v1_curmgt_proto_rawDesc = nil
	file_blue_curmgt_v1_curmgt_proto_goTypes = nil
	file_blue_curmgt_v1_curmgt_proto_depIdxs = nil
}
