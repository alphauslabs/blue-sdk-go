// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.17.3
// source: api/aws/calculator.proto

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

type AccountMonthlyCostCalculated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AWS account id.
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Expected format: yyyy-mm
	Month string `protobuf:"bytes,2,opt,name=month,proto3" json:"month,omitempty"`
	// Expected format: RFC3339
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Usually set when ondemand/invoice request, otherwise, empty.
	RunId string            `protobuf:"bytes,4,opt,name=runId,proto3" json:"runId,omitempty"`
	Meta  map[string]string `protobuf:"bytes,5,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AccountMonthlyCostCalculated) Reset() {
	*x = AccountMonthlyCostCalculated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountMonthlyCostCalculated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountMonthlyCostCalculated) ProtoMessage() {}

func (x *AccountMonthlyCostCalculated) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountMonthlyCostCalculated.ProtoReflect.Descriptor instead.
func (*AccountMonthlyCostCalculated) Descriptor() ([]byte, []int) {
	return file_api_aws_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *AccountMonthlyCostCalculated) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AccountMonthlyCostCalculated) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *AccountMonthlyCostCalculated) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *AccountMonthlyCostCalculated) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *AccountMonthlyCostCalculated) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type AdjustmentsTypeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []string `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
}

func (x *AdjustmentsTypeList) Reset() {
	*x = AdjustmentsTypeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdjustmentsTypeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustmentsTypeList) ProtoMessage() {}

func (x *AdjustmentsTypeList) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjustmentsTypeList.ProtoReflect.Descriptor instead.
func (*AdjustmentsTypeList) Descriptor() ([]byte, []int) {
	return file_api_aws_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *AdjustmentsTypeList) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

type ExcludedServiceFromUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The management account (formerly master account) the `productCode` belongs to.
	ManagementAccount string `protobuf:"bytes,1,opt,name=managementAccount,proto3" json:"managementAccount,omitempty"`
	// The excluded product code.
	ProductCode string `protobuf:"bytes,2,opt,name=productCode,proto3" json:"productCode,omitempty"`
	// If the product code is converted to Adjustments or not.
	ConvertedToAdjustments bool `protobuf:"varint,3,opt,name=convertedToAdjustments,proto3" json:"convertedToAdjustments,omitempty"`
}

func (x *ExcludedServiceFromUsage) Reset() {
	*x = ExcludedServiceFromUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExcludedServiceFromUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExcludedServiceFromUsage) ProtoMessage() {}

func (x *ExcludedServiceFromUsage) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExcludedServiceFromUsage.ProtoReflect.Descriptor instead.
func (*ExcludedServiceFromUsage) Descriptor() ([]byte, []int) {
	return file_api_aws_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *ExcludedServiceFromUsage) GetManagementAccount() string {
	if x != nil {
		return x.ManagementAccount
	}
	return ""
}

func (x *ExcludedServiceFromUsage) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *ExcludedServiceFromUsage) GetConvertedToAdjustments() bool {
	if x != nil {
		return x.ConvertedToAdjustments
	}
	return false
}

// AWS's calculation engine configuration(s).
type CalculatorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CUR lineitem types included in Ripple's 'Adjusting Entries'.
	AdjustmentsTypeList *AdjustmentsTypeList `protobuf:"bytes,1,opt,name=adjustmentsTypeList,proto3" json:"adjustmentsTypeList,omitempty"`
	// List of services that are of type 'Usage' in CUR that are excluded,
	// optionally converted to Adjustments.
	ExcludedServicesFromUsage []*ExcludedServiceFromUsage `protobuf:"bytes,2,rep,name=excludedServicesFromUsage,proto3" json:"excludedServicesFromUsage,omitempty"`
}

func (x *CalculatorConfig) Reset() {
	*x = CalculatorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorConfig) ProtoMessage() {}

func (x *CalculatorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorConfig.ProtoReflect.Descriptor instead.
func (*CalculatorConfig) Descriptor() ([]byte, []int) {
	return file_api_aws_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *CalculatorConfig) GetAdjustmentsTypeList() *AdjustmentsTypeList {
	if x != nil {
		return x.AdjustmentsTypeList
	}
	return nil
}

func (x *CalculatorConfig) GetExcludedServicesFromUsage() []*ExcludedServiceFromUsage {
	if x != nil {
		return x.ExcludedServicesFromUsage
	}
	return nil
}

var File_api_aws_calculator_proto protoreflect.FileDescriptor

var file_api_aws_calculator_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x22, 0x88, 0x02, 0x0a, 0x1c,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f,
	0x73, 0x74, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x75,
	0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64,
	0x12, 0x4b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43,
	0x6f, 0x73, 0x74, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a,
	0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x13, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x18, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x36, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x16, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x41, 0x64, 0x6a,
	0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xd3, 0x01, 0x0a, 0x10, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x56, 0x0a,
	0x13, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x41, 0x64, 0x6a,
	0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x13, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x67, 0x0a, 0x19, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x45, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x19, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x42, 0x62,
	0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x42,
	0x15, 0x41, 0x70, 0x69, 0x41, 0x77, 0x73, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62,
	0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_aws_calculator_proto_rawDescOnce sync.Once
	file_api_aws_calculator_proto_rawDescData = file_api_aws_calculator_proto_rawDesc
)

func file_api_aws_calculator_proto_rawDescGZIP() []byte {
	file_api_aws_calculator_proto_rawDescOnce.Do(func() {
		file_api_aws_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_aws_calculator_proto_rawDescData)
	})
	return file_api_aws_calculator_proto_rawDescData
}

var file_api_aws_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_aws_calculator_proto_goTypes = []interface{}{
	(*AccountMonthlyCostCalculated)(nil), // 0: blueapi.api.aws.AccountMonthlyCostCalculated
	(*AdjustmentsTypeList)(nil),          // 1: blueapi.api.aws.AdjustmentsTypeList
	(*ExcludedServiceFromUsage)(nil),     // 2: blueapi.api.aws.ExcludedServiceFromUsage
	(*CalculatorConfig)(nil),             // 3: blueapi.api.aws.CalculatorConfig
	nil,                                  // 4: blueapi.api.aws.AccountMonthlyCostCalculated.MetaEntry
}
var file_api_aws_calculator_proto_depIdxs = []int32{
	4, // 0: blueapi.api.aws.AccountMonthlyCostCalculated.meta:type_name -> blueapi.api.aws.AccountMonthlyCostCalculated.MetaEntry
	1, // 1: blueapi.api.aws.CalculatorConfig.adjustmentsTypeList:type_name -> blueapi.api.aws.AdjustmentsTypeList
	2, // 2: blueapi.api.aws.CalculatorConfig.excludedServicesFromUsage:type_name -> blueapi.api.aws.ExcludedServiceFromUsage
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_aws_calculator_proto_init() }
func file_api_aws_calculator_proto_init() {
	if File_api_aws_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_aws_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountMonthlyCostCalculated); i {
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
		file_api_aws_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdjustmentsTypeList); i {
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
		file_api_aws_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExcludedServiceFromUsage); i {
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
		file_api_aws_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorConfig); i {
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
			RawDescriptor: file_api_aws_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_aws_calculator_proto_goTypes,
		DependencyIndexes: file_api_aws_calculator_proto_depIdxs,
		MessageInfos:      file_api_aws_calculator_proto_msgTypes,
	}.Build()
	File_api_aws_calculator_proto = out.File
	file_api_aws_calculator_proto_rawDesc = nil
	file_api_aws_calculator_proto_goTypes = nil
	file_api_aws_calculator_proto_depIdxs = nil
}
