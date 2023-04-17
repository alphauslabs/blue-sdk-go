// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: api/ripple/exchangerate.proto

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

// ExchangeRate resource definition.
type ExchangeRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The currency code.
	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	// The rate.
	Rate float64 `protobuf:"fixed64,2,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *ExchangeRate) Reset() {
	*x = ExchangeRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ripple_exchangerate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRate) ProtoMessage() {}

func (x *ExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_exchangerate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRate.ProtoReflect.Descriptor instead.
func (*ExchangeRate) Descriptor() ([]byte, []int) {
	return file_api_ripple_exchangerate_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeRate) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ExchangeRate) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

// MonthlyExchangeRate resource definition.
type MonthlyExchangeRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The year-month. format: yyyymm
	Month string `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty"`
	// The exchange rate.
	ExchangeRate []*ExchangeRate `protobuf:"bytes,2,rep,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
}

func (x *MonthlyExchangeRate) Reset() {
	*x = MonthlyExchangeRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ripple_exchangerate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonthlyExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthlyExchangeRate) ProtoMessage() {}

func (x *MonthlyExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_exchangerate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthlyExchangeRate.ProtoReflect.Descriptor instead.
func (*MonthlyExchangeRate) Descriptor() ([]byte, []int) {
	return file_api_ripple_exchangerate_proto_rawDescGZIP(), []int{1}
}

func (x *MonthlyExchangeRate) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *MonthlyExchangeRate) GetExchangeRate() []*ExchangeRate {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

var File_api_ripple_exchangerate_proto protoreflect.FileDescriptor

var file_api_ripple_exchangerate_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70,
	0x70, 0x6c, 0x65, 0x22, 0x3e, 0x0a, 0x0c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72,
	0x61, 0x74, 0x65, 0x22, 0x71, 0x0a, 0x13, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x12, 0x44, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x42, 0x6d, 0x0a, 0x20, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x42, 0x1a, 0x41, 0x70, 0x69, 0x52,
	0x69, 0x70, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62,
	0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x69, 0x70, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ripple_exchangerate_proto_rawDescOnce sync.Once
	file_api_ripple_exchangerate_proto_rawDescData = file_api_ripple_exchangerate_proto_rawDesc
)

func file_api_ripple_exchangerate_proto_rawDescGZIP() []byte {
	file_api_ripple_exchangerate_proto_rawDescOnce.Do(func() {
		file_api_ripple_exchangerate_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ripple_exchangerate_proto_rawDescData)
	})
	return file_api_ripple_exchangerate_proto_rawDescData
}

var file_api_ripple_exchangerate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_ripple_exchangerate_proto_goTypes = []interface{}{
	(*ExchangeRate)(nil),        // 0: blueapi.api.ripple.ExchangeRate
	(*MonthlyExchangeRate)(nil), // 1: blueapi.api.ripple.MonthlyExchangeRate
}
var file_api_ripple_exchangerate_proto_depIdxs = []int32{
	0, // 0: blueapi.api.ripple.MonthlyExchangeRate.exchangeRate:type_name -> blueapi.api.ripple.ExchangeRate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_ripple_exchangerate_proto_init() }
func file_api_ripple_exchangerate_proto_init() {
	if File_api_ripple_exchangerate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ripple_exchangerate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRate); i {
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
		file_api_ripple_exchangerate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonthlyExchangeRate); i {
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
			RawDescriptor: file_api_ripple_exchangerate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_ripple_exchangerate_proto_goTypes,
		DependencyIndexes: file_api_ripple_exchangerate_proto_depIdxs,
		MessageInfos:      file_api_ripple_exchangerate_proto_msgTypes,
	}.Build()
	File_api_ripple_exchangerate_proto = out.File
	file_api_ripple_exchangerate_proto_rawDesc = nil
	file_api_ripple_exchangerate_proto_goTypes = nil
	file_api_ripple_exchangerate_proto_depIdxs = nil
}