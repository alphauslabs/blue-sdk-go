// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/ripple/rounding.proto

package ripple

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// The rounding method.
	Rounding      Rounding_RoundingMethod `protobuf:"varint,1,opt,name=rounding,proto3,enum=blueapi.api.ripple.Rounding_RoundingMethod" json:"rounding,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Rounding) Reset() {
	*x = Rounding{}
	mi := &file_api_ripple_rounding_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rounding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rounding) ProtoMessage() {}

func (x *Rounding) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_rounding_proto_msgTypes[0]
	if x != nil {
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

const file_api_ripple_rounding_proto_rawDesc = "" +
	"\n" +
	"\x19api/ripple/rounding.proto\x12\x12blueapi.api.ripple\"\x8c\x01\n" +
	"\bRounding\x12G\n" +
	"\brounding\x18\x01 \x01(\x0e2+.blueapi.api.ripple.Rounding.RoundingMethodR\brounding\"7\n" +
	"\x0eRoundingMethod\x12\t\n" +
	"\x05ROUND\x10\x00\x12\v\n" +
	"\aROUNDUP\x10\x01\x12\r\n" +
	"\tROUNDDOWN\x10\x02Bi\n" +
	" cloud.alphaus.blueapi.api.rippleB\x16ApiRippleRoundingProtoZ-github.com/alphauslabs/blue-sdk-go/api/rippleb\x06proto3"

var (
	file_api_ripple_rounding_proto_rawDescOnce sync.Once
	file_api_ripple_rounding_proto_rawDescData []byte
)

func file_api_ripple_rounding_proto_rawDescGZIP() []byte {
	file_api_ripple_rounding_proto_rawDescOnce.Do(func() {
		file_api_ripple_rounding_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_ripple_rounding_proto_rawDesc), len(file_api_ripple_rounding_proto_rawDesc)))
	})
	return file_api_ripple_rounding_proto_rawDescData
}

var file_api_ripple_rounding_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_ripple_rounding_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_ripple_rounding_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_ripple_rounding_proto_rawDesc), len(file_api_ripple_rounding_proto_rawDesc)),
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
	file_api_ripple_rounding_proto_goTypes = nil
	file_api_ripple_rounding_proto_depIdxs = nil
}
