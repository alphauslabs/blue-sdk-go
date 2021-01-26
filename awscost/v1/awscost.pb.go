// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: awscost/v1/awscost.proto

package awscost

import (
	proto "github.com/golang/protobuf/proto"
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

type GetAccountCostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAccountCostsRequest) Reset() {
	*x = GetAccountCostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awscost_v1_awscost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountCostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountCostsRequest) ProtoMessage() {}

func (x *GetAccountCostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_awscost_v1_awscost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountCostsRequest.ProtoReflect.Descriptor instead.
func (*GetAccountCostsRequest) Descriptor() ([]byte, []int) {
	return file_awscost_v1_awscost_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountCostsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAccountCostsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetAccountCostsReply) Reset() {
	*x = GetAccountCostsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awscost_v1_awscost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountCostsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountCostsReply) ProtoMessage() {}

func (x *GetAccountCostsReply) ProtoReflect() protoreflect.Message {
	mi := &file_awscost_v1_awscost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountCostsReply.ProtoReflect.Descriptor instead.
func (*GetAccountCostsReply) Descriptor() ([]byte, []int) {
	return file_awscost_v1_awscost_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountCostsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_awscost_v1_awscost_proto protoreflect.FileDescriptor

var file_awscost_v1_awscost_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x77, 0x73, 0x63, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73,
	0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x98, 0x01, 0x0a,
	0x07, 0x41, 0x77, 0x73, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x8c, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x62,
	0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f,
	0x63, 0x6f, 0x73, 0x74, 0x73, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x77, 0x73, 0x63, 0x6f,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_awscost_v1_awscost_proto_rawDescOnce sync.Once
	file_awscost_v1_awscost_proto_rawDescData = file_awscost_v1_awscost_proto_rawDesc
)

func file_awscost_v1_awscost_proto_rawDescGZIP() []byte {
	file_awscost_v1_awscost_proto_rawDescOnce.Do(func() {
		file_awscost_v1_awscost_proto_rawDescData = protoimpl.X.CompressGZIP(file_awscost_v1_awscost_proto_rawDescData)
	})
	return file_awscost_v1_awscost_proto_rawDescData
}

var file_awscost_v1_awscost_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_awscost_v1_awscost_proto_goTypes = []interface{}{
	(*GetAccountCostsRequest)(nil), // 0: blueapi.awscost.v1.GetAccountCostsRequest
	(*GetAccountCostsReply)(nil),   // 1: blueapi.awscost.v1.GetAccountCostsReply
}
var file_awscost_v1_awscost_proto_depIdxs = []int32{
	0, // 0: blueapi.awscost.v1.AwsCost.GetAccountCosts:input_type -> blueapi.awscost.v1.GetAccountCostsRequest
	1, // 1: blueapi.awscost.v1.AwsCost.GetAccountCosts:output_type -> blueapi.awscost.v1.GetAccountCostsReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_awscost_v1_awscost_proto_init() }
func file_awscost_v1_awscost_proto_init() {
	if File_awscost_v1_awscost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_awscost_v1_awscost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountCostsRequest); i {
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
		file_awscost_v1_awscost_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountCostsReply); i {
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
			RawDescriptor: file_awscost_v1_awscost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_awscost_v1_awscost_proto_goTypes,
		DependencyIndexes: file_awscost_v1_awscost_proto_depIdxs,
		MessageInfos:      file_awscost_v1_awscost_proto_msgTypes,
	}.Build()
	File_awscost_v1_awscost_proto = out.File
	file_awscost_v1_awscost_proto_rawDesc = nil
	file_awscost_v1_awscost_proto_goTypes = nil
	file_awscost_v1_awscost_proto_depIdxs = nil
}
