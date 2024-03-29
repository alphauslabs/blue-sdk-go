// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: api/aws/adjustment.proto

package aws

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

type Adjustment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account        string  `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	GroupId        string  `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Date           string  `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Type           string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	ProductCode    string  `protobuf:"bytes,5,opt,name=productCode,proto3" json:"productCode,omitempty"`
	Description    string  `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Cost           float64 `protobuf:"fixed64,7,opt,name=cost,proto3" json:"cost,omitempty"`
	BaseCurrency   string  `protobuf:"bytes,8,opt,name=baseCurrency,proto3" json:"baseCurrency,omitempty"`
	ExchangeRate   float64 `protobuf:"fixed64,9,opt,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
	TargetCost     float64 `protobuf:"fixed64,10,opt,name=targetCost,proto3" json:"targetCost,omitempty"`
	TargetCurrency string  `protobuf:"bytes,11,opt,name=targetCurrency,proto3" json:"targetCurrency,omitempty"`
}

func (x *Adjustment) Reset() {
	*x = Adjustment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_adjustment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adjustment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adjustment) ProtoMessage() {}

func (x *Adjustment) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_adjustment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adjustment.ProtoReflect.Descriptor instead.
func (*Adjustment) Descriptor() ([]byte, []int) {
	return file_api_aws_adjustment_proto_rawDescGZIP(), []int{0}
}

func (x *Adjustment) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Adjustment) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Adjustment) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Adjustment) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Adjustment) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *Adjustment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Adjustment) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Adjustment) GetBaseCurrency() string {
	if x != nil {
		return x.BaseCurrency
	}
	return ""
}

func (x *Adjustment) GetExchangeRate() float64 {
	if x != nil {
		return x.ExchangeRate
	}
	return 0
}

func (x *Adjustment) GetTargetCost() float64 {
	if x != nil {
		return x.TargetCost
	}
	return 0
}

func (x *Adjustment) GetTargetCurrency() string {
	if x != nil {
		return x.TargetCurrency
	}
	return ""
}

var File_api_aws_adjustment_proto protoreflect.FileDescriptor

var file_api_aws_adjustment_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x22, 0xd0, 0x02, 0x0a, 0x0a,
	0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x43, 0x6f, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x62,
	0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x42,
	0x15, 0x41, 0x70, 0x69, 0x41, 0x77, 0x73, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62,
	0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_aws_adjustment_proto_rawDescOnce sync.Once
	file_api_aws_adjustment_proto_rawDescData = file_api_aws_adjustment_proto_rawDesc
)

func file_api_aws_adjustment_proto_rawDescGZIP() []byte {
	file_api_aws_adjustment_proto_rawDescOnce.Do(func() {
		file_api_aws_adjustment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_aws_adjustment_proto_rawDescData)
	})
	return file_api_aws_adjustment_proto_rawDescData
}

var file_api_aws_adjustment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_aws_adjustment_proto_goTypes = []interface{}{
	(*Adjustment)(nil), // 0: blueapi.api.aws.Adjustment
}
var file_api_aws_adjustment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_aws_adjustment_proto_init() }
func file_api_aws_adjustment_proto_init() {
	if File_api_aws_adjustment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_aws_adjustment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adjustment); i {
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
			RawDescriptor: file_api_aws_adjustment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_aws_adjustment_proto_goTypes,
		DependencyIndexes: file_api_aws_adjustment_proto_depIdxs,
		MessageInfos:      file_api_aws_adjustment_proto_msgTypes,
	}.Build()
	File_api_aws_adjustment_proto = out.File
	file_api_aws_adjustment_proto_rawDesc = nil
	file_api_aws_adjustment_proto_goTypes = nil
	file_api_aws_adjustment_proto_depIdxs = nil
}
