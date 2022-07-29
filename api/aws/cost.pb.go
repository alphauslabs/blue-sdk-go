// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: api/aws/cost.proto

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

type CostAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account        string            `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	GroupId        string            `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	ProductCode    string            `protobuf:"bytes,3,opt,name=productCode,proto3" json:"productCode,omitempty"`
	ServiceCode    string            `protobuf:"bytes,4,opt,name=serviceCode,proto3" json:"serviceCode,omitempty"`
	Region         string            `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	Zone           string            `protobuf:"bytes,6,opt,name=zone,proto3" json:"zone,omitempty"`
	UsageType      string            `protobuf:"bytes,7,opt,name=usageType,proto3" json:"usageType,omitempty"`
	InstanceType   string            `protobuf:"bytes,8,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	Operation      string            `protobuf:"bytes,9,opt,name=operation,proto3" json:"operation,omitempty"`
	InvoiceId      string            `protobuf:"bytes,10,opt,name=invoiceId,proto3" json:"invoiceId,omitempty"`
	Description    string            `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	ResourceId     string            `protobuf:"bytes,12,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	Tags           map[string]string `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CostCategories map[string]string `protobuf:"bytes,14,rep,name=costCategories,proto3" json:"costCategories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CostAttribute) Reset() {
	*x = CostAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_cost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAttribute) ProtoMessage() {}

func (x *CostAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_cost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	Tags           map[string]string `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CostCategories map[string]string `protobuf:"bytes,15,rep,name=costCategories,proto3" json:"costCategories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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
}

func (x *Cost) Reset() {
	*x = Cost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_cost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cost) ProtoMessage() {}

func (x *Cost) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_cost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

var File_api_aws_cost_proto protoreflect.FileDescriptor

var file_api_aws_cost_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x77, 0x73, 0x22, 0x89, 0x05, 0x0a, 0x0d, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x43, 0x6f, 0x73, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x5a, 0x0a, 0x0e, 0x63, 0x6f, 0x73,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x77, 0x73, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41,
	0x0a, 0x13, 0x43, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xee, 0x08, 0x0a, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x51, 0x0a, 0x0e, 0x63,
	0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e,
	0x63, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x6e, 0x62, 0x6c,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x75, 0x6e, 0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x43, 0x6f, 0x73, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x55, 0x6e, 0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x1c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x6e, 0x62, 0x6c, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x24, 0x0a, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73,
	0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6d, 0x6f, 0x72,
	0x74, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x61, 0x6d, 0x6f, 0x72, 0x74, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x72, 0x74, 0x69, 0x7a, 0x65,
	0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x72, 0x74, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x41, 0x0a, 0x13, 0x43, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x5c, 0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x77, 0x73, 0x42, 0x0f, 0x41, 0x70, 0x69, 0x41, 0x77, 0x73, 0x43, 0x6f, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75,
	0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x77, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_aws_cost_proto_rawDescOnce sync.Once
	file_api_aws_cost_proto_rawDescData = file_api_aws_cost_proto_rawDesc
)

func file_api_aws_cost_proto_rawDescGZIP() []byte {
	file_api_aws_cost_proto_rawDescOnce.Do(func() {
		file_api_aws_cost_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_aws_cost_proto_rawDescData)
	})
	return file_api_aws_cost_proto_rawDescData
}

var file_api_aws_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_aws_cost_proto_goTypes = []interface{}{
	(*CostAttribute)(nil), // 0: blueapi.api.aws.CostAttribute
	(*Cost)(nil),          // 1: blueapi.api.aws.Cost
	nil,                   // 2: blueapi.api.aws.CostAttribute.TagsEntry
	nil,                   // 3: blueapi.api.aws.CostAttribute.CostCategoriesEntry
	nil,                   // 4: blueapi.api.aws.Cost.TagsEntry
	nil,                   // 5: blueapi.api.aws.Cost.CostCategoriesEntry
}
var file_api_aws_cost_proto_depIdxs = []int32{
	2, // 0: blueapi.api.aws.CostAttribute.tags:type_name -> blueapi.api.aws.CostAttribute.TagsEntry
	3, // 1: blueapi.api.aws.CostAttribute.costCategories:type_name -> blueapi.api.aws.CostAttribute.CostCategoriesEntry
	4, // 2: blueapi.api.aws.Cost.tags:type_name -> blueapi.api.aws.Cost.TagsEntry
	5, // 3: blueapi.api.aws.Cost.costCategories:type_name -> blueapi.api.aws.Cost.CostCategoriesEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_aws_cost_proto_init() }
func file_api_aws_cost_proto_init() {
	if File_api_aws_cost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_aws_cost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostAttribute); i {
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
		file_api_aws_cost_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cost); i {
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
			RawDescriptor: file_api_aws_cost_proto_rawDesc,
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
	file_api_aws_cost_proto_rawDesc = nil
	file_api_aws_cost_proto_goTypes = nil
	file_api_aws_cost_proto_depIdxs = nil
}
