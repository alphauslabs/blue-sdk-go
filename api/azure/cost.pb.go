// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/azure/cost.proto

package azure

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

type Cost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account being queried.
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// The group id the account is associated with during the query.
	GroupId string `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	// For daily data, format is `yyyy-mm-dd`; for monthly, `yyyy-mm`.
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	// The service name, such as `Software License`, `Cognosys`, `SendGrid`, `New-Commerce ERP Software License`, etc.
	ServiceName string `protobuf:"bytes,4,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	// The product code for an Azure service, such as `Dsv4 Series Windows VM`, `CentOS 7.6`, etc.
	ProductName string `protobuf:"bytes,5,opt,name=productName,proto3" json:"productName,omitempty"`
	// The region of lineitem, if applicable.
	Region string `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	// The charge type of lineitem, if applicable. Such as `New`, `CycleCharge`, `Prorate fees when cancel`, etc.
	ChargeType string `protobuf:"bytes,7,opt,name=chargeType,proto3" json:"chargeType,omitempty"`
	// The description of lineitem, if applicable.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// The billable quantity of lineitem, if applicable.
	BillableQuantity float64 `protobuf:"fixed64,9,opt,name=billableQuantity,proto3" json:"billableQuantity,omitempty"`
	// The effective unit price of lineitem, if applicable.
	EffectiveUnitPrice float64 `protobuf:"fixed64,10,opt,name=effectiveUnitPrice,proto3" json:"effectiveUnitPrice,omitempty"`
	// The true cost (calculated) for this lineitem.
	Cost float64 `protobuf:"fixed64,11,opt,name=cost,proto3" json:"cost,omitempty"`
	// The base currency for `cost`.
	BaseCurrency string `protobuf:"bytes,12,opt,name=baseCurrency,proto3" json:"baseCurrency,omitempty"`
	// The exchange rate used to convert `baseCurrency` to `targetCurrency`.
	ExchangeRate float64 `protobuf:"fixed64,13,opt,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
	// Converted `cost`.
	TargetCost float64 `protobuf:"fixed64,14,opt,name=targetCost,proto3" json:"targetCost,omitempty"`
	// The currency set by `toCurrency`.
	TargetCurrency string `protobuf:"bytes,15,opt,name=targetCurrency,proto3" json:"targetCurrency,omitempty"`
	// The time interval of lineitem, if applicable. Format is `yyyy-MM-ddThh:MM:ssZ/yyyy-mm-ddTHH:mm:ssZ` (for example 2020-09-16T00:00:00Z/2021-09-24T00:00:00Z).
	TimeInterval string `protobuf:"bytes,16,opt,name=timeInterval,proto3" json:"timeInterval,omitempty"`
	// The billing type of lineitem, if applicable. Such as `MARKETPLACE`, `UPFRONT`, `Refund`, `Credit` and `OTHERS`.
	BillingType string `protobuf:"bytes,17,opt,name=billingType,proto3" json:"billingType,omitempty"`
	// The alternate ID of lineitem, if applicable.
	AlternateId string `protobuf:"bytes,18,opt,name=alternateId,proto3" json:"alternateId,omitempty"`
	// The domain name of lineitem, if applicable.
	DomainName string `protobuf:"bytes,19,opt,name=domainName,proto3" json:"domainName,omitempty"`
	// The operation of lineitem, if applicable. Such as `Cool LRS Write Operations`, `Cool LRS Data Write`, `Standard Data Transfer Out`, etc.
	Operation string `protobuf:"bytes,20,opt,name=operation,proto3" json:"operation,omitempty"`
	// The usage type of lineitem, if applicable. Such as `Standard HDD Managed Disks`, `Tables`, `Blob Storage`, etc.
	UsageType string `protobuf:"bytes,21,opt,name=usageType,proto3" json:"usageType,omitempty"`
	// The instance type of lineitem, if applicable. Such as `Gateway`, `Standard_B2s`, `Standard_D4s_v3`, etc.
	InstanceType string `protobuf:"bytes,22,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	// The category of lineitem, if applicable. Such as `Software License`, `Marketplace`, `RI`, `Other`, etc.
	Category string `protobuf:"bytes,23,opt,name=category,proto3" json:"category,omitempty"`
	// The subscription id.
	SubscriptionId string `protobuf:"bytes,24,opt,name=subscriptionId,proto3" json:"subscriptionId,omitempty"`
	// The entitlement id.
	EntitlementId string `protobuf:"bytes,25,opt,name=entitlementId,proto3" json:"entitlementId,omitempty"`
}

func (x *Cost) Reset() {
	*x = Cost{}
	mi := &file_api_azure_cost_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cost) ProtoMessage() {}

func (x *Cost) ProtoReflect() protoreflect.Message {
	mi := &file_api_azure_cost_proto_msgTypes[0]
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
	return file_api_azure_cost_proto_rawDescGZIP(), []int{0}
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

func (x *Cost) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Cost) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Cost) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Cost) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Cost) GetChargeType() string {
	if x != nil {
		return x.ChargeType
	}
	return ""
}

func (x *Cost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Cost) GetBillableQuantity() float64 {
	if x != nil {
		return x.BillableQuantity
	}
	return 0
}

func (x *Cost) GetEffectiveUnitPrice() float64 {
	if x != nil {
		return x.EffectiveUnitPrice
	}
	return 0
}

func (x *Cost) GetCost() float64 {
	if x != nil {
		return x.Cost
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

func (x *Cost) GetTargetCurrency() string {
	if x != nil {
		return x.TargetCurrency
	}
	return ""
}

func (x *Cost) GetTimeInterval() string {
	if x != nil {
		return x.TimeInterval
	}
	return ""
}

func (x *Cost) GetBillingType() string {
	if x != nil {
		return x.BillingType
	}
	return ""
}

func (x *Cost) GetAlternateId() string {
	if x != nil {
		return x.AlternateId
	}
	return ""
}

func (x *Cost) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *Cost) GetOperation() string {
	if x != nil {
		return x.Operation
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

func (x *Cost) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Cost) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *Cost) GetEntitlementId() string {
	if x != nil {
		return x.EntitlementId
	}
	return ""
}

type CostAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId     string `protobuf:"bytes,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	SubscriptionId string `protobuf:"bytes,2,opt,name=subscriptionId,proto3" json:"subscriptionId,omitempty"`
	EntitlementId  string `protobuf:"bytes,3,opt,name=entitlementId,proto3" json:"entitlementId,omitempty"`
	GroupId        string `protobuf:"bytes,4,opt,name=groupId,proto3" json:"groupId,omitempty"`
	ProductId      string `protobuf:"bytes,5,opt,name=productId,proto3" json:"productId,omitempty"`
	ProductName    string `protobuf:"bytes,6,opt,name=productName,proto3" json:"productName,omitempty"`
	SkuId          string `protobuf:"bytes,7,opt,name=skuId,proto3" json:"skuId,omitempty"`
	SkuName        string `protobuf:"bytes,8,opt,name=skuName,proto3" json:"skuName,omitempty"`
	Description    string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Category       string `protobuf:"bytes,10,opt,name=category,proto3" json:"category,omitempty"`
	DomainName     string `protobuf:"bytes,11,opt,name=domainName,proto3" json:"domainName,omitempty"`
}

func (x *CostAttribute) Reset() {
	*x = CostAttribute{}
	mi := &file_api_azure_cost_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAttribute) ProtoMessage() {}

func (x *CostAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_api_azure_cost_proto_msgTypes[1]
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
	return file_api_azure_cost_proto_rawDescGZIP(), []int{1}
}

func (x *CostAttribute) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CostAttribute) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *CostAttribute) GetEntitlementId() string {
	if x != nil {
		return x.EntitlementId
	}
	return ""
}

func (x *CostAttribute) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CostAttribute) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CostAttribute) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *CostAttribute) GetSkuId() string {
	if x != nil {
		return x.SkuId
	}
	return ""
}

func (x *CostAttribute) GetSkuName() string {
	if x != nil {
		return x.SkuName
	}
	return ""
}

func (x *CostAttribute) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CostAttribute) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CostAttribute) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

var File_api_azure_cost_proto protoreflect.FileDescriptor

var file_api_azure_cost_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x22, 0xbe, 0x06, 0x0a, 0x04, 0x43, 0x6f,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x62, 0x69, 0x6c, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x55, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x12, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xe5, 0x02, 0x0a, 0x0d, 0x43,
	0x6f, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b,
	0x75, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x75,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x62, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x42, 0x11, 0x41, 0x70, 0x69, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43,
	0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_azure_cost_proto_rawDescOnce sync.Once
	file_api_azure_cost_proto_rawDescData = file_api_azure_cost_proto_rawDesc
)

func file_api_azure_cost_proto_rawDescGZIP() []byte {
	file_api_azure_cost_proto_rawDescOnce.Do(func() {
		file_api_azure_cost_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_azure_cost_proto_rawDescData)
	})
	return file_api_azure_cost_proto_rawDescData
}

var file_api_azure_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_azure_cost_proto_goTypes = []any{
	(*Cost)(nil),          // 0: blueapi.api.azure.Cost
	(*CostAttribute)(nil), // 1: blueapi.api.azure.CostAttribute
}
var file_api_azure_cost_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_azure_cost_proto_init() }
func file_api_azure_cost_proto_init() {
	if File_api_azure_cost_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_azure_cost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_azure_cost_proto_goTypes,
		DependencyIndexes: file_api_azure_cost_proto_depIdxs,
		MessageInfos:      file_api_azure_cost_proto_msgTypes,
	}.Build()
	File_api_azure_cost_proto = out.File
	file_api_azure_cost_proto_rawDesc = nil
	file_api_azure_cost_proto_goTypes = nil
	file_api_azure_cost_proto_depIdxs = nil
}
