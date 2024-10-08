// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/aws/reduction.proto

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

type AwsCostReductions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary          *CostReductionSummary `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	RiCostReductions []*RiCostReduction    `protobuf:"bytes,2,rep,name=riCostReductions,proto3" json:"riCostReductions,omitempty"`
	SpCostReductions []*SpCostReduction    `protobuf:"bytes,3,rep,name=spCostReductions,proto3" json:"spCostReductions,omitempty"`
}

func (x *AwsCostReductions) Reset() {
	*x = AwsCostReductions{}
	mi := &file_api_aws_reduction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsCostReductions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsCostReductions) ProtoMessage() {}

func (x *AwsCostReductions) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_reduction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsCostReductions.ProtoReflect.Descriptor instead.
func (*AwsCostReductions) Descriptor() ([]byte, []int) {
	return file_api_aws_reduction_proto_rawDescGZIP(), []int{0}
}

func (x *AwsCostReductions) GetSummary() *CostReductionSummary {
	if x != nil {
		return x.Summary
	}
	return nil
}

func (x *AwsCostReductions) GetRiCostReductions() []*RiCostReduction {
	if x != nil {
		return x.RiCostReductions
	}
	return nil
}

func (x *AwsCostReductions) GetSpCostReductions() []*SpCostReduction {
	if x != nil {
		return x.SpCostReductions
	}
	return nil
}

type RiCostReduction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Arn                 string  `protobuf:"bytes,2,opt,name=arn,proto3" json:"arn,omitempty"`
	CustomerId          string  `protobuf:"bytes,3,opt,name=customerId,proto3" json:"customerId,omitempty"`
	CustomerName        string  `protobuf:"bytes,4,opt,name=customerName,proto3" json:"customerName,omitempty"`
	BillingInternalId   string  `protobuf:"bytes,5,opt,name=billingInternalId,proto3" json:"billingInternalId,omitempty"`
	BillingGroupId      string  `protobuf:"bytes,6,opt,name=billingGroupId,proto3" json:"billingGroupId,omitempty"`
	BillingGroupName    string  `protobuf:"bytes,7,opt,name=billingGroupName,proto3" json:"billingGroupName,omitempty"`
	DestCustomerId      string  `protobuf:"bytes,8,opt,name=destCustomerId,proto3" json:"destCustomerId,omitempty"`
	Start               string  `protobuf:"bytes,9,opt,name=start,proto3" json:"start,omitempty"`
	End                 string  `protobuf:"bytes,10,opt,name=end,proto3" json:"end,omitempty"`
	Service             string  `protobuf:"bytes,11,opt,name=service,proto3" json:"service,omitempty"`
	InstanceType        string  `protobuf:"bytes,12,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	ModificationStatus  string  `protobuf:"bytes,13,opt,name=modificationStatus,proto3" json:"modificationStatus,omitempty"`
	Number              int64   `protobuf:"varint,14,opt,name=number,proto3" json:"number,omitempty"`
	OfferClass          string  `protobuf:"bytes,15,opt,name=offerClass,proto3" json:"offerClass,omitempty"`
	DeploymentOption    string  `protobuf:"bytes,35,opt,name=deploymentOption,proto3" json:"deploymentOption,omitempty"`
	PaidBy              string  `protobuf:"bytes,16,opt,name=paidBy,proto3" json:"paidBy,omitempty"`
	PaymentOption       string  `protobuf:"bytes,17,opt,name=paymentOption,proto3" json:"paymentOption,omitempty"`
	Platform            string  `protobuf:"bytes,18,opt,name=platform,proto3" json:"platform,omitempty"`
	Region              string  `protobuf:"bytes,19,opt,name=region,proto3" json:"region,omitempty"`
	Remove              bool    `protobuf:"varint,20,opt,name=remove,proto3" json:"remove,omitempty"`
	Scope               string  `protobuf:"bytes,21,opt,name=scope,proto3" json:"scope,omitempty"`
	Tenancy             string  `protobuf:"bytes,22,opt,name=tenancy,proto3" json:"tenancy,omitempty"`
	TermLength          string  `protobuf:"bytes,23,opt,name=termLength,proto3" json:"termLength,omitempty"`
	UsageType           string  `protobuf:"bytes,24,opt,name=usageType,proto3" json:"usageType,omitempty"`
	Vendor              string  `protobuf:"bytes,25,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Zone                string  `protobuf:"bytes,26,opt,name=zone,proto3" json:"zone,omitempty"`
	Disabled            bool    `protobuf:"varint,27,opt,name=disabled,proto3" json:"disabled,omitempty"`
	NormalizationFactor float64 `protobuf:"fixed64,28,opt,name=normalizationFactor,proto3" json:"normalizationFactor,omitempty"`
	UnblendedRate       float64 `protobuf:"fixed64,29,opt,name=unblendedRate,proto3" json:"unblendedRate,omitempty"`
	UpfrontValue        float64 `protobuf:"fixed64,30,opt,name=upfrontValue,proto3" json:"upfrontValue,omitempty"`
	OndemandCost        float64 `protobuf:"fixed64,31,opt,name=ondemandCost,proto3" json:"ondemandCost,omitempty"`
	EffectiveCost       float64 `protobuf:"fixed64,32,opt,name=effectiveCost,proto3" json:"effectiveCost,omitempty"`
	Savings             float64 `protobuf:"fixed64,33,opt,name=savings,proto3" json:"savings,omitempty"`
	// for break even point date
	BreakEven string `protobuf:"bytes,34,opt,name=breakEven,proto3" json:"breakEven,omitempty"`
}

func (x *RiCostReduction) Reset() {
	*x = RiCostReduction{}
	mi := &file_api_aws_reduction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RiCostReduction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiCostReduction) ProtoMessage() {}

func (x *RiCostReduction) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_reduction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiCostReduction.ProtoReflect.Descriptor instead.
func (*RiCostReduction) Descriptor() ([]byte, []int) {
	return file_api_aws_reduction_proto_rawDescGZIP(), []int{1}
}

func (x *RiCostReduction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RiCostReduction) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

func (x *RiCostReduction) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *RiCostReduction) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *RiCostReduction) GetBillingInternalId() string {
	if x != nil {
		return x.BillingInternalId
	}
	return ""
}

func (x *RiCostReduction) GetBillingGroupId() string {
	if x != nil {
		return x.BillingGroupId
	}
	return ""
}

func (x *RiCostReduction) GetBillingGroupName() string {
	if x != nil {
		return x.BillingGroupName
	}
	return ""
}

func (x *RiCostReduction) GetDestCustomerId() string {
	if x != nil {
		return x.DestCustomerId
	}
	return ""
}

func (x *RiCostReduction) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *RiCostReduction) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *RiCostReduction) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *RiCostReduction) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *RiCostReduction) GetModificationStatus() string {
	if x != nil {
		return x.ModificationStatus
	}
	return ""
}

func (x *RiCostReduction) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *RiCostReduction) GetOfferClass() string {
	if x != nil {
		return x.OfferClass
	}
	return ""
}

func (x *RiCostReduction) GetDeploymentOption() string {
	if x != nil {
		return x.DeploymentOption
	}
	return ""
}

func (x *RiCostReduction) GetPaidBy() string {
	if x != nil {
		return x.PaidBy
	}
	return ""
}

func (x *RiCostReduction) GetPaymentOption() string {
	if x != nil {
		return x.PaymentOption
	}
	return ""
}

func (x *RiCostReduction) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *RiCostReduction) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RiCostReduction) GetRemove() bool {
	if x != nil {
		return x.Remove
	}
	return false
}

func (x *RiCostReduction) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *RiCostReduction) GetTenancy() string {
	if x != nil {
		return x.Tenancy
	}
	return ""
}

func (x *RiCostReduction) GetTermLength() string {
	if x != nil {
		return x.TermLength
	}
	return ""
}

func (x *RiCostReduction) GetUsageType() string {
	if x != nil {
		return x.UsageType
	}
	return ""
}

func (x *RiCostReduction) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *RiCostReduction) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *RiCostReduction) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *RiCostReduction) GetNormalizationFactor() float64 {
	if x != nil {
		return x.NormalizationFactor
	}
	return 0
}

func (x *RiCostReduction) GetUnblendedRate() float64 {
	if x != nil {
		return x.UnblendedRate
	}
	return 0
}

func (x *RiCostReduction) GetUpfrontValue() float64 {
	if x != nil {
		return x.UpfrontValue
	}
	return 0
}

func (x *RiCostReduction) GetOndemandCost() float64 {
	if x != nil {
		return x.OndemandCost
	}
	return 0
}

func (x *RiCostReduction) GetEffectiveCost() float64 {
	if x != nil {
		return x.EffectiveCost
	}
	return 0
}

func (x *RiCostReduction) GetSavings() float64 {
	if x != nil {
		return x.Savings
	}
	return 0
}

func (x *RiCostReduction) GetBreakEven() string {
	if x != nil {
		return x.BreakEven
	}
	return ""
}

type SpCostReduction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Arn              string  `protobuf:"bytes,2,opt,name=arn,proto3" json:"arn,omitempty"`
	CustomerId       string  `protobuf:"bytes,3,opt,name=customerId,proto3" json:"customerId,omitempty"`
	CustomerName     string  `protobuf:"bytes,4,opt,name=customerName,proto3" json:"customerName,omitempty"`
	BillingGroupId   string  `protobuf:"bytes,5,opt,name=billingGroupId,proto3" json:"billingGroupId,omitempty"`
	BillingGroupName string  `protobuf:"bytes,6,opt,name=billingGroupName,proto3" json:"billingGroupName,omitempty"`
	Currency         string  `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	Service          string  `protobuf:"bytes,8,opt,name=service,proto3" json:"service,omitempty"`
	SavingsPlan      string  `protobuf:"bytes,9,opt,name=savingsPlan,proto3" json:"savingsPlan,omitempty"`
	Start            string  `protobuf:"bytes,10,opt,name=start,proto3" json:"start,omitempty"`
	End              string  `protobuf:"bytes,11,opt,name=end,proto3" json:"end,omitempty"`
	Region           string  `protobuf:"bytes,12,opt,name=region,proto3" json:"region,omitempty"`
	InstanceFamily   string  `protobuf:"bytes,13,opt,name=instanceFamily,proto3" json:"instanceFamily,omitempty"`
	TermLength       string  `protobuf:"bytes,14,opt,name=termLength,proto3" json:"termLength,omitempty"`
	PaymentOption    string  `protobuf:"bytes,15,opt,name=paymentOption,proto3" json:"paymentOption,omitempty"`
	Commitment       float64 `protobuf:"fixed64,16,opt,name=commitment,proto3" json:"commitment,omitempty"`
	UpfrontFee       float64 `protobuf:"fixed64,17,opt,name=upfrontFee,proto3" json:"upfrontFee,omitempty"`
	RecurringFee     float64 `protobuf:"fixed64,18,opt,name=recurringFee,proto3" json:"recurringFee,omitempty"`
	OndemandCost     float64 `protobuf:"fixed64,19,opt,name=ondemandCost,proto3" json:"ondemandCost,omitempty"`
	EffectiveCost    float64 `protobuf:"fixed64,20,opt,name=effectiveCost,proto3" json:"effectiveCost,omitempty"`
	Savings          float64 `protobuf:"fixed64,21,opt,name=savings,proto3" json:"savings,omitempty"`
	// for break even point date
	BreakEven string `protobuf:"bytes,22,opt,name=breakEven,proto3" json:"breakEven,omitempty"`
}

func (x *SpCostReduction) Reset() {
	*x = SpCostReduction{}
	mi := &file_api_aws_reduction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpCostReduction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpCostReduction) ProtoMessage() {}

func (x *SpCostReduction) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_reduction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpCostReduction.ProtoReflect.Descriptor instead.
func (*SpCostReduction) Descriptor() ([]byte, []int) {
	return file_api_aws_reduction_proto_rawDescGZIP(), []int{2}
}

func (x *SpCostReduction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SpCostReduction) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

func (x *SpCostReduction) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *SpCostReduction) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *SpCostReduction) GetBillingGroupId() string {
	if x != nil {
		return x.BillingGroupId
	}
	return ""
}

func (x *SpCostReduction) GetBillingGroupName() string {
	if x != nil {
		return x.BillingGroupName
	}
	return ""
}

func (x *SpCostReduction) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *SpCostReduction) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *SpCostReduction) GetSavingsPlan() string {
	if x != nil {
		return x.SavingsPlan
	}
	return ""
}

func (x *SpCostReduction) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *SpCostReduction) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *SpCostReduction) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *SpCostReduction) GetInstanceFamily() string {
	if x != nil {
		return x.InstanceFamily
	}
	return ""
}

func (x *SpCostReduction) GetTermLength() string {
	if x != nil {
		return x.TermLength
	}
	return ""
}

func (x *SpCostReduction) GetPaymentOption() string {
	if x != nil {
		return x.PaymentOption
	}
	return ""
}

func (x *SpCostReduction) GetCommitment() float64 {
	if x != nil {
		return x.Commitment
	}
	return 0
}

func (x *SpCostReduction) GetUpfrontFee() float64 {
	if x != nil {
		return x.UpfrontFee
	}
	return 0
}

func (x *SpCostReduction) GetRecurringFee() float64 {
	if x != nil {
		return x.RecurringFee
	}
	return 0
}

func (x *SpCostReduction) GetOndemandCost() float64 {
	if x != nil {
		return x.OndemandCost
	}
	return 0
}

func (x *SpCostReduction) GetEffectiveCost() float64 {
	if x != nil {
		return x.EffectiveCost
	}
	return 0
}

func (x *SpCostReduction) GetSavings() float64 {
	if x != nil {
		return x.Savings
	}
	return 0
}

func (x *SpCostReduction) GetBreakEven() string {
	if x != nil {
		return x.BreakEven
	}
	return ""
}

type CostReductionSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalOnDemandCost    float64 `protobuf:"fixed64,1,opt,name=totalOnDemandCost,proto3" json:"totalOnDemandCost,omitempty"`
	TotalEffectiveCost   float64 `protobuf:"fixed64,2,opt,name=totalEffectiveCost,proto3" json:"totalEffectiveCost,omitempty"`
	TotalSavings         float64 `protobuf:"fixed64,3,opt,name=totalSavings,proto3" json:"totalSavings,omitempty"`
	TotalRiOnDemandCost  float64 `protobuf:"fixed64,4,opt,name=totalRiOnDemandCost,proto3" json:"totalRiOnDemandCost,omitempty"`
	TotalRiEffectiveCost float64 `protobuf:"fixed64,5,opt,name=totalRiEffectiveCost,proto3" json:"totalRiEffectiveCost,omitempty"`
	TotalRiSavings       float64 `protobuf:"fixed64,6,opt,name=totalRiSavings,proto3" json:"totalRiSavings,omitempty"`
	TotalSpOnDemandCost  float64 `protobuf:"fixed64,7,opt,name=totalSpOnDemandCost,proto3" json:"totalSpOnDemandCost,omitempty"`
	TotalSpEffectiveCost float64 `protobuf:"fixed64,8,opt,name=totalSpEffectiveCost,proto3" json:"totalSpEffectiveCost,omitempty"`
	TotalSpSavings       float64 `protobuf:"fixed64,9,opt,name=totalSpSavings,proto3" json:"totalSpSavings,omitempty"`
}

func (x *CostReductionSummary) Reset() {
	*x = CostReductionSummary{}
	mi := &file_api_aws_reduction_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostReductionSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostReductionSummary) ProtoMessage() {}

func (x *CostReductionSummary) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_reduction_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostReductionSummary.ProtoReflect.Descriptor instead.
func (*CostReductionSummary) Descriptor() ([]byte, []int) {
	return file_api_aws_reduction_proto_rawDescGZIP(), []int{3}
}

func (x *CostReductionSummary) GetTotalOnDemandCost() float64 {
	if x != nil {
		return x.TotalOnDemandCost
	}
	return 0
}

func (x *CostReductionSummary) GetTotalEffectiveCost() float64 {
	if x != nil {
		return x.TotalEffectiveCost
	}
	return 0
}

func (x *CostReductionSummary) GetTotalSavings() float64 {
	if x != nil {
		return x.TotalSavings
	}
	return 0
}

func (x *CostReductionSummary) GetTotalRiOnDemandCost() float64 {
	if x != nil {
		return x.TotalRiOnDemandCost
	}
	return 0
}

func (x *CostReductionSummary) GetTotalRiEffectiveCost() float64 {
	if x != nil {
		return x.TotalRiEffectiveCost
	}
	return 0
}

func (x *CostReductionSummary) GetTotalRiSavings() float64 {
	if x != nil {
		return x.TotalRiSavings
	}
	return 0
}

func (x *CostReductionSummary) GetTotalSpOnDemandCost() float64 {
	if x != nil {
		return x.TotalSpOnDemandCost
	}
	return 0
}

func (x *CostReductionSummary) GetTotalSpEffectiveCost() float64 {
	if x != nil {
		return x.TotalSpEffectiveCost
	}
	return 0
}

func (x *CostReductionSummary) GetTotalSpSavings() float64 {
	if x != nil {
		return x.TotalSpSavings
	}
	return 0
}

var File_api_aws_reduction_proto protoreflect.FileDescriptor

var file_api_aws_reduction_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x11, 0x41,
	0x77, 0x73, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x3f, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x77, 0x73, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x4c, 0x0a, 0x10, 0x72, 0x69, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x52, 0x69,
	0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x72,
	0x69, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x4c, 0x0a, 0x10, 0x73, 0x70, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x53, 0x70, 0x43, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x70, 0x43,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xd9, 0x08,
	0x0a, 0x0f, 0x52, 0x69, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x72, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x73,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x2a, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x69, 0x64, 0x42, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x69,
	0x64, 0x42, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x6e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x24, 0x0a,
	0x0d, 0x75, 0x6e, 0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x18, 0x1d,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x75, 0x6e, 0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x70, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x75, 0x70, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x6e, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6f,
	0x6e, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x20, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x21, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x73, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x72, 0x65, 0x61, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x72, 0x65, 0x61, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x22, 0xb7, 0x05, 0x0a, 0x0f, 0x53, 0x70,
	0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x72, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x6c, 0x61, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x6c, 0x61, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x72,
	0x6d, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x75, 0x70, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x65,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x6e, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73,
	0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6f, 0x6e, 0x64, 0x65, 0x6d, 0x61, 0x6e,
	0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x73, 0x61,
	0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x45, 0x76,
	0x65, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x45,
	0x76, 0x65, 0x6e, 0x22, 0xb4, 0x03, 0x0a, 0x14, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x11,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x6e,
	0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x30,
	0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x69, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e,
	0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x52, 0x69, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x69, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x69, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x69, 0x53,
	0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x52, 0x69, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x13,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43,
	0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x70, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x70, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x53, 0x61, 0x76,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x53, 0x70, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x61, 0x0a, 0x1d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x42, 0x14, 0x41, 0x70, 0x69,
	0x41, 0x77, 0x73, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73,
	0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x77, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_aws_reduction_proto_rawDescOnce sync.Once
	file_api_aws_reduction_proto_rawDescData = file_api_aws_reduction_proto_rawDesc
)

func file_api_aws_reduction_proto_rawDescGZIP() []byte {
	file_api_aws_reduction_proto_rawDescOnce.Do(func() {
		file_api_aws_reduction_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_aws_reduction_proto_rawDescData)
	})
	return file_api_aws_reduction_proto_rawDescData
}

var file_api_aws_reduction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_aws_reduction_proto_goTypes = []any{
	(*AwsCostReductions)(nil),    // 0: blueapi.api.aws.AwsCostReductions
	(*RiCostReduction)(nil),      // 1: blueapi.api.aws.RiCostReduction
	(*SpCostReduction)(nil),      // 2: blueapi.api.aws.SpCostReduction
	(*CostReductionSummary)(nil), // 3: blueapi.api.aws.CostReductionSummary
}
var file_api_aws_reduction_proto_depIdxs = []int32{
	3, // 0: blueapi.api.aws.AwsCostReductions.summary:type_name -> blueapi.api.aws.CostReductionSummary
	1, // 1: blueapi.api.aws.AwsCostReductions.riCostReductions:type_name -> blueapi.api.aws.RiCostReduction
	2, // 2: blueapi.api.aws.AwsCostReductions.spCostReductions:type_name -> blueapi.api.aws.SpCostReduction
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_aws_reduction_proto_init() }
func file_api_aws_reduction_proto_init() {
	if File_api_aws_reduction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_aws_reduction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_aws_reduction_proto_goTypes,
		DependencyIndexes: file_api_aws_reduction_proto_depIdxs,
		MessageInfos:      file_api_aws_reduction_proto_msgTypes,
	}.Build()
	File_api_aws_reduction_proto = out.File
	file_api_aws_reduction_proto_rawDesc = nil
	file_api_aws_reduction_proto_goTypes = nil
	file_api_aws_reduction_proto_depIdxs = nil
}
