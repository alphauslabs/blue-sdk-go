// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/aws/calculator.proto

package aws

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

type AccountMonthlyCostCalculated struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// AWS account id.
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Expected format: yyyy-mm
	Month string `protobuf:"bytes,2,opt,name=month,proto3" json:"month,omitempty"`
	// Expected format: RFC3339
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Usually set when ondemand/invoice request, otherwise, empty.
	RunId         string            `protobuf:"bytes,4,opt,name=runId,proto3" json:"runId,omitempty"`
	Meta          map[string]string `protobuf:"bytes,5,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountMonthlyCostCalculated) Reset() {
	*x = AccountMonthlyCostCalculated{}
	mi := &file_api_aws_calculator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountMonthlyCostCalculated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountMonthlyCostCalculated) ProtoMessage() {}

func (x *AccountMonthlyCostCalculated) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_calculator_proto_msgTypes[0]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Types         []string               `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdjustmentsTypeList) Reset() {
	*x = AdjustmentsTypeList{}
	mi := &file_api_aws_calculator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdjustmentsTypeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustmentsTypeList) ProtoMessage() {}

func (x *AdjustmentsTypeList) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_calculator_proto_msgTypes[1]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// The management account (formerly master account) the `productCode` belongs to.
	ManagementAccount string `protobuf:"bytes,1,opt,name=managementAccount,proto3" json:"managementAccount,omitempty"`
	// The excluded product code.
	ProductCode string `protobuf:"bytes,2,opt,name=productCode,proto3" json:"productCode,omitempty"`
	// If the product code is converted to Adjustments or not.
	ConvertedToAdjustments bool `protobuf:"varint,3,opt,name=convertedToAdjustments,proto3" json:"convertedToAdjustments,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *ExcludedServiceFromUsage) Reset() {
	*x = ExcludedServiceFromUsage{}
	mi := &file_api_aws_calculator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExcludedServiceFromUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExcludedServiceFromUsage) ProtoMessage() {}

func (x *ExcludedServiceFromUsage) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_calculator_proto_msgTypes[2]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// CUR lineitem types included in Ripple's 'Adjusting Entries'.
	AdjustmentsTypeList *AdjustmentsTypeList `protobuf:"bytes,1,opt,name=adjustmentsTypeList,proto3" json:"adjustmentsTypeList,omitempty"`
	// List of services that are of type 'Usage' in CUR that are excluded,
	// optionally converted to Adjustments.
	ExcludedServicesFromUsage []*ExcludedServiceFromUsage `protobuf:"bytes,2,rep,name=excludedServicesFromUsage,proto3" json:"excludedServicesFromUsage,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *CalculatorConfig) Reset() {
	*x = CalculatorConfig{}
	mi := &file_api_aws_calculator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalculatorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorConfig) ProtoMessage() {}

func (x *CalculatorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_calculator_proto_msgTypes[3]
	if x != nil {
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

const file_api_aws_calculator_proto_rawDesc = "" +
	"\n" +
	"\x18api/aws/calculator.proto\x12\x0fblueapi.api.aws\"\x88\x02\n" +
	"\x1cAccountMonthlyCostCalculated\x12\x18\n" +
	"\aaccount\x18\x01 \x01(\tR\aaccount\x12\x14\n" +
	"\x05month\x18\x02 \x01(\tR\x05month\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\tR\ttimestamp\x12\x14\n" +
	"\x05runId\x18\x04 \x01(\tR\x05runId\x12K\n" +
	"\x04meta\x18\x05 \x03(\v27.blueapi.api.aws.AccountMonthlyCostCalculated.MetaEntryR\x04meta\x1a7\n" +
	"\tMetaEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"+\n" +
	"\x13AdjustmentsTypeList\x12\x14\n" +
	"\x05types\x18\x01 \x03(\tR\x05types\"\xa2\x01\n" +
	"\x18ExcludedServiceFromUsage\x12,\n" +
	"\x11managementAccount\x18\x01 \x01(\tR\x11managementAccount\x12 \n" +
	"\vproductCode\x18\x02 \x01(\tR\vproductCode\x126\n" +
	"\x16convertedToAdjustments\x18\x03 \x01(\bR\x16convertedToAdjustments\"\xd3\x01\n" +
	"\x10CalculatorConfig\x12V\n" +
	"\x13adjustmentsTypeList\x18\x01 \x01(\v2$.blueapi.api.aws.AdjustmentsTypeListR\x13adjustmentsTypeList\x12g\n" +
	"\x19excludedServicesFromUsage\x18\x02 \x03(\v2).blueapi.api.aws.ExcludedServiceFromUsageR\x19excludedServicesFromUsageBb\n" +
	"\x1dcloud.alphaus.blueapi.api.awsB\x15ApiAwsCalculatorProtoZ*github.com/alphauslabs/blue-sdk-go/api/awsb\x06proto3"

var (
	file_api_aws_calculator_proto_rawDescOnce sync.Once
	file_api_aws_calculator_proto_rawDescData []byte
)

func file_api_aws_calculator_proto_rawDescGZIP() []byte {
	file_api_aws_calculator_proto_rawDescOnce.Do(func() {
		file_api_aws_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_aws_calculator_proto_rawDesc), len(file_api_aws_calculator_proto_rawDesc)))
	})
	return file_api_aws_calculator_proto_rawDescData
}

var file_api_aws_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_aws_calculator_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_aws_calculator_proto_rawDesc), len(file_api_aws_calculator_proto_rawDesc)),
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
	file_api_aws_calculator_proto_goTypes = nil
	file_api_aws_calculator_proto_depIdxs = nil
}
