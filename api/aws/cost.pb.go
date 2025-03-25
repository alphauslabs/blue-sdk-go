// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/aws/cost.proto

package aws

import (
	api "github.com/alphauslabs/blue-sdk-go/api"
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

type CostAttribute struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Account        string                 `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	GroupId        string                 `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	ProductCode    string                 `protobuf:"bytes,3,opt,name=productCode,proto3" json:"productCode,omitempty"`
	ServiceCode    string                 `protobuf:"bytes,4,opt,name=serviceCode,proto3" json:"serviceCode,omitempty"`
	Region         string                 `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	Zone           string                 `protobuf:"bytes,6,opt,name=zone,proto3" json:"zone,omitempty"`
	UsageType      string                 `protobuf:"bytes,7,opt,name=usageType,proto3" json:"usageType,omitempty"`
	InstanceType   string                 `protobuf:"bytes,8,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	Operation      string                 `protobuf:"bytes,9,opt,name=operation,proto3" json:"operation,omitempty"`
	InvoiceId      string                 `protobuf:"bytes,10,opt,name=invoiceId,proto3" json:"invoiceId,omitempty"`
	Description    string                 `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	ResourceId     string                 `protobuf:"bytes,12,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	Tags           map[string]string      `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CostCategories map[string]string      `protobuf:"bytes,14,rep,name=costCategories,proto3" json:"costCategories,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CostAttribute) Reset() {
	*x = CostAttribute{}
	mi := &file_api_aws_cost_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAttribute) ProtoMessage() {}

func (x *CostAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_cost_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostAttribute.ProtoReflect.Descriptor instead.
func (*CostAttribute) Descriptor() ([]byte, []int) {
	return file_api_aws_cost_proto_rawDescGZIP(), []int{0}
}

func (x *CostAttribute) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CostAttribute) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CostAttribute) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *CostAttribute) GetServiceCode() string {
	if x != nil {
		return x.ServiceCode
	}
	return ""
}

func (x *CostAttribute) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CostAttribute) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *CostAttribute) GetUsageType() string {
	if x != nil {
		return x.UsageType
	}
	return ""
}

func (x *CostAttribute) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *CostAttribute) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *CostAttribute) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *CostAttribute) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CostAttribute) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *CostAttribute) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CostAttribute) GetCostCategories() map[string]string {
	if x != nil {
		return x.CostCategories
	}
	return nil
}

type Cost struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The account being queried.
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// The group id the account is associated with during the query.
	GroupId string `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Type    string `protobuf:"bytes,29,opt,name=type,proto3" json:"type,omitempty"`
	// For daily data, format is `yyyy-mm-dd`; for monthly, `yyyy-mm`.
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	// The product code for an AWS service, such as `AmazonEC2`, `AmazonRDS`, etc. This can also be an Alphaus-specified custom value.
	ProductCode string `protobuf:"bytes,4,opt,name=productCode,proto3" json:"productCode,omitempty"`
	// The CUR service code of the lineitem, if applicable. Sometimes, this is the same as `productCode`, a subset of `productCode`, an Alphaus-specified custom value, or empty.
	ServiceCode string `protobuf:"bytes,5,opt,name=serviceCode,proto3" json:"serviceCode,omitempty"`
	// The region of the lineitem, if applicable.
	Region string `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	// The zone of the lineitem, if applicable.
	Zone string `protobuf:"bytes,7,opt,name=zone,proto3" json:"zone,omitempty"`
	// The CUR usage type of the lineitem, if applicable.
	UsageType string `protobuf:"bytes,8,opt,name=usageType,proto3" json:"usageType,omitempty"`
	// The CUR instance type of the lineitem, if applicable.
	InstanceType string `protobuf:"bytes,9,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	// The CUR operation of the lineitem, if applicable.
	Operation string `protobuf:"bytes,10,opt,name=operation,proto3" json:"operation,omitempty"`
	// The AWS invoice ID of the lineitem, if applicable.
	InvoiceId string `protobuf:"bytes,11,opt,name=invoiceId,proto3" json:"invoiceId,omitempty"`
	// The description of the lineitem, if applicable.
	Description string `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	// The resource id of the lineitem, if applicable. At the moment, this is not yet fully supported.
	ResourceId     string            `protobuf:"bytes,13,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	Tags           map[string]string `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CostCategories map[string]string `protobuf:"bytes,15,rep,name=costCategories,proto3" json:"costCategories,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Only set when `description` and/or `resourceId` attributes are specified.
	Usage float64 `protobuf:"fixed64,16,opt,name=usage,proto3" json:"usage,omitempty"`
	// The true cost (calculated) for this lineitem.
	Cost float64 `protobuf:"fixed64,17,opt,name=cost,proto3" json:"cost,omitempty"`
	// The unblended cost as reflected in the CUR for this lineitem.
	UnblendedCost float64 `protobuf:"fixed64,27,opt,name=unblendedCost,proto3" json:"unblendedCost,omitempty"`
	// The base currency for `cost`, `unblendedCost`, `effectiveCost`, `amortizedCost`. Always set to `USD`, CUR's default currency.
	BaseCurrency string `protobuf:"bytes,18,opt,name=baseCurrency,proto3" json:"baseCurrency,omitempty"`
	// The exchange rate used to convert `baseCurrency` to `targetCurrency`.
	ExchangeRate float64 `protobuf:"fixed64,19,opt,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
	// Converted `cost`.
	TargetCost float64 `protobuf:"fixed64,20,opt,name=targetCost,proto3" json:"targetCost,omitempty"`
	// Converted `unblendedCost`.
	TargetUnblendedCost float64 `protobuf:"fixed64,28,opt,name=targetUnblendedCost,proto3" json:"targetUnblendedCost,omitempty"`
	// The currency set by `toCurrency`.
	TargetCurrency string  `protobuf:"bytes,21,opt,name=targetCurrency,proto3" json:"targetCurrency,omitempty"`
	EffectiveCost  float64 `protobuf:"fixed64,23,opt,name=effectiveCost,proto3" json:"effectiveCost,omitempty"`
	// Converted `effectiveCost`.
	TargetEffectiveCost float64 `protobuf:"fixed64,24,opt,name=targetEffectiveCost,proto3" json:"targetEffectiveCost,omitempty"`
	AmortizedCost       float64 `protobuf:"fixed64,25,opt,name=amortizedCost,proto3" json:"amortizedCost,omitempty"`
	// Converted `amortizedCost`.
	TargetAmortizedCost float64 `protobuf:"fixed64,26,opt,name=targetAmortizedCost,proto3" json:"targetAmortizedCost,omitempty"`
	TagId               string  `protobuf:"bytes,22,opt,name=tagId,proto3" json:"tagId,omitempty"`
	// Get last update in UNIX time format.
	Timestamp string `protobuf:"bytes,30,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Various metadata associated with this lineitem.
	Metadata      []*api.KeyValue `protobuf:"bytes,31,rep,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cost) Reset() {
	*x = Cost{}
	mi := &file_api_aws_cost_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cost) ProtoMessage() {}

func (x *Cost) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_cost_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cost.ProtoReflect.Descriptor instead.
func (*Cost) Descriptor() ([]byte, []int) {
	return file_api_aws_cost_proto_rawDescGZIP(), []int{1}
}

func (x *Cost) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Cost) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Cost) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Cost) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Cost) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *Cost) GetServiceCode() string {
	if x != nil {
		return x.ServiceCode
	}
	return ""
}

func (x *Cost) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Cost) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Cost) GetUsageType() string {
	if x != nil {
		return x.UsageType
	}
	return ""
}

func (x *Cost) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *Cost) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *Cost) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *Cost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Cost) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Cost) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Cost) GetCostCategories() map[string]string {
	if x != nil {
		return x.CostCategories
	}
	return nil
}

func (x *Cost) GetUsage() float64 {
	if x != nil {
		return x.Usage
	}
	return 0
}

func (x *Cost) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Cost) GetUnblendedCost() float64 {
	if x != nil {
		return x.UnblendedCost
	}
	return 0
}

func (x *Cost) GetBaseCurrency() string {
	if x != nil {
		return x.BaseCurrency
	}
	return ""
}

func (x *Cost) GetExchangeRate() float64 {
	if x != nil {
		return x.ExchangeRate
	}
	return 0
}

func (x *Cost) GetTargetCost() float64 {
	if x != nil {
		return x.TargetCost
	}
	return 0
}

func (x *Cost) GetTargetUnblendedCost() float64 {
	if x != nil {
		return x.TargetUnblendedCost
	}
	return 0
}

func (x *Cost) GetTargetCurrency() string {
	if x != nil {
		return x.TargetCurrency
	}
	return ""
}

func (x *Cost) GetEffectiveCost() float64 {
	if x != nil {
		return x.EffectiveCost
	}
	return 0
}

func (x *Cost) GetTargetEffectiveCost() float64 {
	if x != nil {
		return x.TargetEffectiveCost
	}
	return 0
}

func (x *Cost) GetAmortizedCost() float64 {
	if x != nil {
		return x.AmortizedCost
	}
	return 0
}

func (x *Cost) GetTargetAmortizedCost() float64 {
	if x != nil {
		return x.TargetAmortizedCost
	}
	return 0
}

func (x *Cost) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *Cost) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Cost) GetMetadata() []*api.KeyValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_api_aws_cost_proto protoreflect.FileDescriptor

const file_api_aws_cost_proto_rawDesc = "" +
	"\n" +
	"\x12api/aws/cost.proto\x12\x0fblueapi.api.aws\x1a\x12api/keyvalue.proto\"\x89\x05\n" +
	"\rCostAttribute\x12\x18\n" +
	"\aaccount\x18\x01 \x01(\tR\aaccount\x12\x18\n" +
	"\agroupId\x18\x02 \x01(\tR\agroupId\x12 \n" +
	"\vproductCode\x18\x03 \x01(\tR\vproductCode\x12 \n" +
	"\vserviceCode\x18\x04 \x01(\tR\vserviceCode\x12\x16\n" +
	"\x06region\x18\x05 \x01(\tR\x06region\x12\x12\n" +
	"\x04zone\x18\x06 \x01(\tR\x04zone\x12\x1c\n" +
	"\tusageType\x18\a \x01(\tR\tusageType\x12\"\n" +
	"\finstanceType\x18\b \x01(\tR\finstanceType\x12\x1c\n" +
	"\toperation\x18\t \x01(\tR\toperation\x12\x1c\n" +
	"\tinvoiceId\x18\n" +
	" \x01(\tR\tinvoiceId\x12 \n" +
	"\vdescription\x18\v \x01(\tR\vdescription\x12\x1e\n" +
	"\n" +
	"resourceId\x18\f \x01(\tR\n" +
	"resourceId\x12<\n" +
	"\x04tags\x18\r \x03(\v2(.blueapi.api.aws.CostAttribute.TagsEntryR\x04tags\x12Z\n" +
	"\x0ecostCategories\x18\x0e \x03(\v22.blueapi.api.aws.CostAttribute.CostCategoriesEntryR\x0ecostCategories\x1a7\n" +
	"\tTagsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1aA\n" +
	"\x13CostCategoriesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xbf\t\n" +
	"\x04Cost\x12\x18\n" +
	"\aaccount\x18\x01 \x01(\tR\aaccount\x12\x18\n" +
	"\agroupId\x18\x02 \x01(\tR\agroupId\x12\x12\n" +
	"\x04type\x18\x1d \x01(\tR\x04type\x12\x12\n" +
	"\x04date\x18\x03 \x01(\tR\x04date\x12 \n" +
	"\vproductCode\x18\x04 \x01(\tR\vproductCode\x12 \n" +
	"\vserviceCode\x18\x05 \x01(\tR\vserviceCode\x12\x16\n" +
	"\x06region\x18\x06 \x01(\tR\x06region\x12\x12\n" +
	"\x04zone\x18\a \x01(\tR\x04zone\x12\x1c\n" +
	"\tusageType\x18\b \x01(\tR\tusageType\x12\"\n" +
	"\finstanceType\x18\t \x01(\tR\finstanceType\x12\x1c\n" +
	"\toperation\x18\n" +
	" \x01(\tR\toperation\x12\x1c\n" +
	"\tinvoiceId\x18\v \x01(\tR\tinvoiceId\x12 \n" +
	"\vdescription\x18\f \x01(\tR\vdescription\x12\x1e\n" +
	"\n" +
	"resourceId\x18\r \x01(\tR\n" +
	"resourceId\x123\n" +
	"\x04tags\x18\x0e \x03(\v2\x1f.blueapi.api.aws.Cost.TagsEntryR\x04tags\x12Q\n" +
	"\x0ecostCategories\x18\x0f \x03(\v2).blueapi.api.aws.Cost.CostCategoriesEntryR\x0ecostCategories\x12\x14\n" +
	"\x05usage\x18\x10 \x01(\x01R\x05usage\x12\x12\n" +
	"\x04cost\x18\x11 \x01(\x01R\x04cost\x12$\n" +
	"\runblendedCost\x18\x1b \x01(\x01R\runblendedCost\x12\"\n" +
	"\fbaseCurrency\x18\x12 \x01(\tR\fbaseCurrency\x12\"\n" +
	"\fexchangeRate\x18\x13 \x01(\x01R\fexchangeRate\x12\x1e\n" +
	"\n" +
	"targetCost\x18\x14 \x01(\x01R\n" +
	"targetCost\x120\n" +
	"\x13targetUnblendedCost\x18\x1c \x01(\x01R\x13targetUnblendedCost\x12&\n" +
	"\x0etargetCurrency\x18\x15 \x01(\tR\x0etargetCurrency\x12$\n" +
	"\reffectiveCost\x18\x17 \x01(\x01R\reffectiveCost\x120\n" +
	"\x13targetEffectiveCost\x18\x18 \x01(\x01R\x13targetEffectiveCost\x12$\n" +
	"\ramortizedCost\x18\x19 \x01(\x01R\ramortizedCost\x120\n" +
	"\x13targetAmortizedCost\x18\x1a \x01(\x01R\x13targetAmortizedCost\x12\x14\n" +
	"\x05tagId\x18\x16 \x01(\tR\x05tagId\x12\x1c\n" +
	"\ttimestamp\x18\x1e \x01(\tR\ttimestamp\x121\n" +
	"\bmetadata\x18\x1f \x03(\v2\x15.blueapi.api.KeyValueR\bmetadata\x1a7\n" +
	"\tTagsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1aA\n" +
	"\x13CostCategoriesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\\\n" +
	"\x1dcloud.alphaus.blueapi.api.awsB\x0fApiAwsCostProtoZ*github.com/alphauslabs/blue-sdk-go/api/awsb\x06proto3"

var (
	file_api_aws_cost_proto_rawDescOnce sync.Once
	file_api_aws_cost_proto_rawDescData []byte
)

func file_api_aws_cost_proto_rawDescGZIP() []byte {
	file_api_aws_cost_proto_rawDescOnce.Do(func() {
		file_api_aws_cost_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_aws_cost_proto_rawDesc), len(file_api_aws_cost_proto_rawDesc)))
	})
	return file_api_aws_cost_proto_rawDescData
}

var file_api_aws_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_aws_cost_proto_goTypes = []any{
	(*CostAttribute)(nil), // 0: blueapi.api.aws.CostAttribute
	(*Cost)(nil),          // 1: blueapi.api.aws.Cost
	nil,                   // 2: blueapi.api.aws.CostAttribute.TagsEntry
	nil,                   // 3: blueapi.api.aws.CostAttribute.CostCategoriesEntry
	nil,                   // 4: blueapi.api.aws.Cost.TagsEntry
	nil,                   // 5: blueapi.api.aws.Cost.CostCategoriesEntry
	(*api.KeyValue)(nil),  // 6: blueapi.api.KeyValue
}
var file_api_aws_cost_proto_depIdxs = []int32{
	2, // 0: blueapi.api.aws.CostAttribute.tags:type_name -> blueapi.api.aws.CostAttribute.TagsEntry
	3, // 1: blueapi.api.aws.CostAttribute.costCategories:type_name -> blueapi.api.aws.CostAttribute.CostCategoriesEntry
	4, // 2: blueapi.api.aws.Cost.tags:type_name -> blueapi.api.aws.Cost.TagsEntry
	5, // 3: blueapi.api.aws.Cost.costCategories:type_name -> blueapi.api.aws.Cost.CostCategoriesEntry
	6, // 4: blueapi.api.aws.Cost.metadata:type_name -> blueapi.api.KeyValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_aws_cost_proto_init() }
func file_api_aws_cost_proto_init() {
	if File_api_aws_cost_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_aws_cost_proto_rawDesc), len(file_api_aws_cost_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_aws_cost_proto_goTypes,
		DependencyIndexes: file_api_aws_cost_proto_depIdxs,
		MessageInfos:      file_api_aws_cost_proto_msgTypes,
	}.Build()
	File_api_aws_cost_proto = out.File
	file_api_aws_cost_proto_goTypes = nil
	file_api_aws_cost_proto_depIdxs = nil
}
