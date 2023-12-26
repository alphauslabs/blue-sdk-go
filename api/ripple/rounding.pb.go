// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/ripple/rounding.proto

package ripple

import (
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

// RoundingMethod
type Rounding_RoundingMethod int32

const (
	// round
	Rounding_ROUND Rounding_RoundingMethod = 0
	// round up
	Rounding_ROUNDUP Rounding_RoundingMethod = 1
	// round down
	Rounding_ROUNDDOWN Rounding_RoundingMethod = 2
)

// Enum value maps for Rounding_RoundingMethod.
var (
	Rounding_RoundingMethod_name = map[int32]string{
		0: "ROUND",
		1: "ROUNDUP",
		2: "ROUNDDOWN",
	}
	Rounding_RoundingMethod_value = map[string]int32{
		"ROUND":     0,
		"ROUNDUP":   1,
		"ROUNDDOWN": 2,
	}
)

func (x Rounding_RoundingMethod) Enum() *Rounding_RoundingMethod {
	p := new(Rounding_RoundingMethod)
	*p = x
	return p
}

func (x Rounding_RoundingMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Rounding_RoundingMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ripple_rounding_proto_enumTypes[0].Descriptor()
}

func (Rounding_RoundingMethod) Type() protoreflect.EnumType {
	return &file_api_ripple_rounding_proto_enumTypes[0]
}

func (x Rounding_RoundingMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Rounding_RoundingMethod.Descriptor instead.
func (Rounding_RoundingMethod) EnumDescriptor() ([]byte, []int) {
	return file_api_ripple_rounding_proto_rawDescGZIP(), []int{0, 0}
}

// Rounding resource definition.
type Rounding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The rounding method.
	Rounding Rounding_RoundingMethod `protobuf:"varint,1,opt,name=rounding,proto3,enum=blueapi.api.ripple.Rounding_RoundingMethod" json:"rounding,omitempty"`
}

func (x *Rounding) Reset() {
	*x = Rounding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ripple_rounding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rounding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rounding) ProtoMessage() {}

func (x *Rounding) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_rounding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rounding.ProtoReflect.Descriptor instead.
func (*Rounding) Descriptor() ([]byte, []int) {
	return file_api_ripple_rounding_proto_rawDescGZIP(), []int{0}
}

func (x *Rounding) GetRounding() Rounding_RoundingMethod {
	if x != nil {
		return x.Rounding
	}
	return Rounding_ROUND
}

var File_api_ripple_rounding_proto protoreflect.FileDescriptor

var file_api_ripple_rounding_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x22,
	0x8c, 0x01, 0x0a, 0x08, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x47, 0x0a, 0x08,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70,
	0x70, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x08, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x37, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x55, 0x50, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x42, 0x69,
	0x0a, 0x20, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70,
	0x6c, 0x65, 0x42, 0x16, 0x41, 0x70, 0x69, 0x52, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_ripple_rounding_proto_rawDescOnce sync.Once
	file_api_ripple_rounding_proto_rawDescData = file_api_ripple_rounding_proto_rawDesc
)

func file_api_ripple_rounding_proto_rawDescGZIP() []byte {
	file_api_ripple_rounding_proto_rawDescOnce.Do(func() {
		file_api_ripple_rounding_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ripple_rounding_proto_rawDescData)
	})
	return file_api_ripple_rounding_proto_rawDescData
}

var file_api_ripple_rounding_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_ripple_rounding_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_ripple_rounding_proto_goTypes = []interface{}{
	(Rounding_RoundingMethod)(0), // 0: blueapi.api.ripple.Rounding.RoundingMethod
	(*Rounding)(nil),             // 1: blueapi.api.ripple.Rounding
}
var file_api_ripple_rounding_proto_depIdxs = []int32{
	0, // 0: blueapi.api.ripple.Rounding.rounding:type_name -> blueapi.api.ripple.Rounding.RoundingMethod
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_ripple_rounding_proto_init() }
func file_api_ripple_rounding_proto_init() {
	if File_api_ripple_rounding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ripple_rounding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rounding); i {
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
			RawDescriptor: file_api_ripple_rounding_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_ripple_rounding_proto_goTypes,
		DependencyIndexes: file_api_ripple_rounding_proto_depIdxs,
		EnumInfos:         file_api_ripple_rounding_proto_enumTypes,
		MessageInfos:      file_api_ripple_rounding_proto_msgTypes,
	}.Build()
	File_api_ripple_rounding_proto = out.File
	file_api_ripple_rounding_proto_rawDesc = nil
	file_api_ripple_rounding_proto_goTypes = nil
	file_api_ripple_rounding_proto_depIdxs = nil
}
