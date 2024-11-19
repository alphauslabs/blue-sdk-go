// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/cover/costforecast.proto

package cover

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

type AwsCostForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// For daily data, format is `yyyy-mm-dd`; for monthly, `yyyy-mm`.
	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	// The account being queried.
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	// The product code for an AWS service
	ProductCode string `protobuf:"bytes,3,opt,name=productCode,proto3" json:"productCode,omitempty"`
	// Forcasted cost based on true cost (calculated)
	Cost       float64 `protobuf:"fixed64,4,opt,name=cost,proto3" json:"cost,omitempty"`
	LowerBound float64 `protobuf:"fixed64,5,opt,name=lowerBound,proto3" json:"lowerBound,omitempty"`
	UpperBound float64 `protobuf:"fixed64,6,opt,name=upperBound,proto3" json:"upperBound,omitempty"`
	// Forcasted cost based on unblended cost
	UnblendedCost       float64 `protobuf:"fixed64,7,opt,name=unblendedCost,proto3" json:"unblendedCost,omitempty"`
	UnblendedLowerBound float64 `protobuf:"fixed64,8,opt,name=unblendedLowerBound,proto3" json:"unblendedLowerBound,omitempty"`
	UnblendedUpperBound float64 `protobuf:"fixed64,9,opt,name=unblendedUpperBound,proto3" json:"unblendedUpperBound,omitempty"`
}

func (x *AwsCostForecast) Reset() {
	*x = AwsCostForecast{}
	mi := &file_api_cover_costforecast_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsCostForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsCostForecast) ProtoMessage() {}

func (x *AwsCostForecast) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_costforecast_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsCostForecast.ProtoReflect.Descriptor instead.
func (*AwsCostForecast) Descriptor() ([]byte, []int) {
	return file_api_cover_costforecast_proto_rawDescGZIP(), []int{0}
}

func (x *AwsCostForecast) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *AwsCostForecast) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AwsCostForecast) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *AwsCostForecast) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *AwsCostForecast) GetLowerBound() float64 {
	if x != nil {
		return x.LowerBound
	}
	return 0
}

func (x *AwsCostForecast) GetUpperBound() float64 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

func (x *AwsCostForecast) GetUnblendedCost() float64 {
	if x != nil {
		return x.UnblendedCost
	}
	return 0
}

func (x *AwsCostForecast) GetUnblendedLowerBound() float64 {
	if x != nil {
		return x.UnblendedLowerBound
	}
	return 0
}

func (x *AwsCostForecast) GetUnblendedUpperBound() float64 {
	if x != nil {
		return x.UnblendedUpperBound
	}
	return 0
}

var File_api_cover_costforecast_proto protoreflect.FileDescriptor

var file_api_cover_costforecast_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x73, 0x74,
	0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x22, 0xbf, 0x02, 0x0a, 0x0f, 0x41, 0x77, 0x73, 0x43, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x70,
	0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x75,
	0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x6e, 0x62,
	0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x75, 0x6e, 0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x13, 0x75, 0x6e, 0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x4c, 0x6f, 0x77, 0x65,
	0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x75, 0x6e,
	0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x30, 0x0a, 0x13, 0x75, 0x6e, 0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x55, 0x70,
	0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13,
	0x75, 0x6e, 0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x55, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f,
	0x75, 0x6e, 0x64, 0x42, 0x6a, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x19, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73,
	0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cover_costforecast_proto_rawDescOnce sync.Once
	file_api_cover_costforecast_proto_rawDescData = file_api_cover_costforecast_proto_rawDesc
)

func file_api_cover_costforecast_proto_rawDescGZIP() []byte {
	file_api_cover_costforecast_proto_rawDescOnce.Do(func() {
		file_api_cover_costforecast_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_costforecast_proto_rawDescData)
	})
	return file_api_cover_costforecast_proto_rawDescData
}

var file_api_cover_costforecast_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_cover_costforecast_proto_goTypes = []any{
	(*AwsCostForecast)(nil), // 0: blueapi.api.cover.AwsCostForecast
}
var file_api_cover_costforecast_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_cover_costforecast_proto_init() }
func file_api_cover_costforecast_proto_init() {
	if File_api_cover_costforecast_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cover_costforecast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_costforecast_proto_goTypes,
		DependencyIndexes: file_api_cover_costforecast_proto_depIdxs,
		MessageInfos:      file_api_cover_costforecast_proto_msgTypes,
	}.Build()
	File_api_cover_costforecast_proto = out.File
	file_api_cover_costforecast_proto_rawDesc = nil
	file_api_cover_costforecast_proto_goTypes = nil
	file_api_cover_costforecast_proto_depIdxs = nil
}
