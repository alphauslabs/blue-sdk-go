// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/aws/recommendation.proto

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

type AwsRecommendations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary           *ReservationRecommendationSummary `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	RiRecommendations []*RiRecommendation               `protobuf:"bytes,2,rep,name=riRecommendations,proto3" json:"riRecommendations,omitempty"`
	SpRecommendations []*SpRecommendation               `protobuf:"bytes,3,rep,name=spRecommendations,proto3" json:"spRecommendations,omitempty"`
	EstimatedCoverage []*ReservationEstimatedCoverage   `protobuf:"bytes,4,rep,name=estimatedCoverage,proto3" json:"estimatedCoverage,omitempty"`
}

func (x *AwsRecommendations) Reset() {
	*x = AwsRecommendations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_recommendation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsRecommendations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsRecommendations) ProtoMessage() {}

func (x *AwsRecommendations) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_recommendation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsRecommendations.ProtoReflect.Descriptor instead.
func (*AwsRecommendations) Descriptor() ([]byte, []int) {
	return file_api_aws_recommendation_proto_rawDescGZIP(), []int{0}
}

func (x *AwsRecommendations) GetSummary() *ReservationRecommendationSummary {
	if x != nil {
		return x.Summary
	}
	return nil
}

func (x *AwsRecommendations) GetRiRecommendations() []*RiRecommendation {
	if x != nil {
		return x.RiRecommendations
	}
	return nil
}

func (x *AwsRecommendations) GetSpRecommendations() []*SpRecommendation {
	if x != nil {
		return x.SpRecommendations
	}
	return nil
}

func (x *AwsRecommendations) GetEstimatedCoverage() []*ReservationEstimatedCoverage {
	if x != nil {
		return x.EstimatedCoverage
	}
	return nil
}

type RiRecommendation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId               string  `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	ProductCode             string  `protobuf:"bytes,2,opt,name=productCode,proto3" json:"productCode,omitempty"`
	Location                string  `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	InstanceType            string  `protobuf:"bytes,4,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	Quantity                int64   `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	NormalizedUnit          float64 `protobuf:"fixed64,6,opt,name=normalizedUnit,proto3" json:"normalizedUnit,omitempty"`
	OperatingSystem         string  `protobuf:"bytes,7,opt,name=operatingSystem,proto3" json:"operatingSystem,omitempty"`
	PreInstalledSW          string  `protobuf:"bytes,8,opt,name=preInstalledSW,proto3" json:"preInstalledSW,omitempty"`
	Tenancy                 string  `protobuf:"bytes,9,opt,name=tenancy,proto3" json:"tenancy,omitempty"`
	OndemandCost            float64 `protobuf:"fixed64,10,opt,name=ondemandCost,proto3" json:"ondemandCost,omitempty"`
	Ondemandrate            float64 `protobuf:"fixed64,11,opt,name=ondemandrate,proto3" json:"ondemandrate,omitempty"`
	RiRate                  float64 `protobuf:"fixed64,12,opt,name=riRate,proto3" json:"riRate,omitempty"`
	UpfrontCost             float64 `protobuf:"fixed64,13,opt,name=upfrontCost,proto3" json:"upfrontCost,omitempty"`
	DiscountedCost          float64 `protobuf:"fixed64,14,opt,name=discountedCost,proto3" json:"discountedCost,omitempty"`
	MonthlyAmortizedCost    float64 `protobuf:"fixed64,15,opt,name=monthlyAmortizedCost,proto3" json:"monthlyAmortizedCost,omitempty"`
	MonthlyRecurringCost    float64 `protobuf:"fixed64,16,opt,name=monthlyRecurringCost,proto3" json:"monthlyRecurringCost,omitempty"`
	YearlyDiscountedCost    float64 `protobuf:"fixed64,17,opt,name=yearlyDiscountedCost,proto3" json:"yearlyDiscountedCost,omitempty"`
	ReductionRate           float64 `protobuf:"fixed64,18,opt,name=reductionRate,proto3" json:"reductionRate,omitempty"`
	EstimatedMonthlySavings float64 `protobuf:"fixed64,19,opt,name=estimatedMonthlySavings,proto3" json:"estimatedMonthlySavings,omitempty"`
}

func (x *RiRecommendation) Reset() {
	*x = RiRecommendation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_recommendation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiRecommendation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiRecommendation) ProtoMessage() {}

func (x *RiRecommendation) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_recommendation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiRecommendation.ProtoReflect.Descriptor instead.
func (*RiRecommendation) Descriptor() ([]byte, []int) {
	return file_api_aws_recommendation_proto_rawDescGZIP(), []int{1}
}

func (x *RiRecommendation) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *RiRecommendation) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *RiRecommendation) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *RiRecommendation) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *RiRecommendation) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *RiRecommendation) GetNormalizedUnit() float64 {
	if x != nil {
		return x.NormalizedUnit
	}
	return 0
}

func (x *RiRecommendation) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *RiRecommendation) GetPreInstalledSW() string {
	if x != nil {
		return x.PreInstalledSW
	}
	return ""
}

func (x *RiRecommendation) GetTenancy() string {
	if x != nil {
		return x.Tenancy
	}
	return ""
}

func (x *RiRecommendation) GetOndemandCost() float64 {
	if x != nil {
		return x.OndemandCost
	}
	return 0
}

func (x *RiRecommendation) GetOndemandrate() float64 {
	if x != nil {
		return x.Ondemandrate
	}
	return 0
}

func (x *RiRecommendation) GetRiRate() float64 {
	if x != nil {
		return x.RiRate
	}
	return 0
}

func (x *RiRecommendation) GetUpfrontCost() float64 {
	if x != nil {
		return x.UpfrontCost
	}
	return 0
}

func (x *RiRecommendation) GetDiscountedCost() float64 {
	if x != nil {
		return x.DiscountedCost
	}
	return 0
}

func (x *RiRecommendation) GetMonthlyAmortizedCost() float64 {
	if x != nil {
		return x.MonthlyAmortizedCost
	}
	return 0
}

func (x *RiRecommendation) GetMonthlyRecurringCost() float64 {
	if x != nil {
		return x.MonthlyRecurringCost
	}
	return 0
}

func (x *RiRecommendation) GetYearlyDiscountedCost() float64 {
	if x != nil {
		return x.YearlyDiscountedCost
	}
	return 0
}

func (x *RiRecommendation) GetReductionRate() float64 {
	if x != nil {
		return x.ReductionRate
	}
	return 0
}

func (x *RiRecommendation) GetEstimatedMonthlySavings() float64 {
	if x != nil {
		return x.EstimatedMonthlySavings
	}
	return 0
}

type SpRecommendation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId               string  `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	SpProductFamily         string  `protobuf:"bytes,2,opt,name=spProductFamily,proto3" json:"spProductFamily,omitempty"`
	Location                string  `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	InstanceFamily          string  `protobuf:"bytes,4,opt,name=instanceFamily,proto3" json:"instanceFamily,omitempty"`
	OndemandCost            float64 `protobuf:"fixed64,5,opt,name=ondemandCost,proto3" json:"ondemandCost,omitempty"`
	Commitment              float64 `protobuf:"fixed64,6,opt,name=commitment,proto3" json:"commitment,omitempty"`
	NormalizedUnit          float64 `protobuf:"fixed64,7,opt,name=normalizedUnit,proto3" json:"normalizedUnit,omitempty"`
	DiscountedCost          float64 `protobuf:"fixed64,8,opt,name=discountedCost,proto3" json:"discountedCost,omitempty"`
	MonthlyDiscountedCost   float64 `protobuf:"fixed64,9,opt,name=monthlyDiscountedCost,proto3" json:"monthlyDiscountedCost,omitempty"`
	YearlyDiscountedCost    float64 `protobuf:"fixed64,10,opt,name=yearlyDiscountedCost,proto3" json:"yearlyDiscountedCost,omitempty"`
	EstimatedMonthlySavings float64 `protobuf:"fixed64,11,opt,name=estimatedMonthlySavings,proto3" json:"estimatedMonthlySavings,omitempty"`
	ReductionRate           float64 `protobuf:"fixed64,12,opt,name=reductionRate,proto3" json:"reductionRate,omitempty"`
}

func (x *SpRecommendation) Reset() {
	*x = SpRecommendation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_recommendation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpRecommendation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpRecommendation) ProtoMessage() {}

func (x *SpRecommendation) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_recommendation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpRecommendation.ProtoReflect.Descriptor instead.
func (*SpRecommendation) Descriptor() ([]byte, []int) {
	return file_api_aws_recommendation_proto_rawDescGZIP(), []int{2}
}

func (x *SpRecommendation) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *SpRecommendation) GetSpProductFamily() string {
	if x != nil {
		return x.SpProductFamily
	}
	return ""
}

func (x *SpRecommendation) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *SpRecommendation) GetInstanceFamily() string {
	if x != nil {
		return x.InstanceFamily
	}
	return ""
}

func (x *SpRecommendation) GetOndemandCost() float64 {
	if x != nil {
		return x.OndemandCost
	}
	return 0
}

func (x *SpRecommendation) GetCommitment() float64 {
	if x != nil {
		return x.Commitment
	}
	return 0
}

func (x *SpRecommendation) GetNormalizedUnit() float64 {
	if x != nil {
		return x.NormalizedUnit
	}
	return 0
}

func (x *SpRecommendation) GetDiscountedCost() float64 {
	if x != nil {
		return x.DiscountedCost
	}
	return 0
}

func (x *SpRecommendation) GetMonthlyDiscountedCost() float64 {
	if x != nil {
		return x.MonthlyDiscountedCost
	}
	return 0
}

func (x *SpRecommendation) GetYearlyDiscountedCost() float64 {
	if x != nil {
		return x.YearlyDiscountedCost
	}
	return 0
}

func (x *SpRecommendation) GetEstimatedMonthlySavings() float64 {
	if x != nil {
		return x.EstimatedMonthlySavings
	}
	return 0
}

func (x *SpRecommendation) GetReductionRate() float64 {
	if x != nil {
		return x.ReductionRate
	}
	return 0
}

type ReservationRecommendationSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalOnDemandCost            float64 `protobuf:"fixed64,1,opt,name=totalOnDemandCost,proto3" json:"totalOnDemandCost,omitempty"`
	TotalDiscountedCost          float64 `protobuf:"fixed64,2,opt,name=totalDiscountedCost,proto3" json:"totalDiscountedCost,omitempty"`
	TotalDiscountedMonthlyCost   float64 `protobuf:"fixed64,3,opt,name=totalDiscountedMonthlyCost,proto3" json:"totalDiscountedMonthlyCost,omitempty"`
	TotalDiscountedYearlyCost    float64 `protobuf:"fixed64,4,opt,name=totalDiscountedYearlyCost,proto3" json:"totalDiscountedYearlyCost,omitempty"`
	ReductionRate                float64 `protobuf:"fixed64,5,opt,name=reductionRate,proto3" json:"reductionRate,omitempty"`
	TotalEstimatedMonthlySavings float64 `protobuf:"fixed64,6,opt,name=totalEstimatedMonthlySavings,proto3" json:"totalEstimatedMonthlySavings,omitempty"`
}

func (x *ReservationRecommendationSummary) Reset() {
	*x = ReservationRecommendationSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_recommendation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationRecommendationSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationRecommendationSummary) ProtoMessage() {}

func (x *ReservationRecommendationSummary) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_recommendation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationRecommendationSummary.ProtoReflect.Descriptor instead.
func (*ReservationRecommendationSummary) Descriptor() ([]byte, []int) {
	return file_api_aws_recommendation_proto_rawDescGZIP(), []int{3}
}

func (x *ReservationRecommendationSummary) GetTotalOnDemandCost() float64 {
	if x != nil {
		return x.TotalOnDemandCost
	}
	return 0
}

func (x *ReservationRecommendationSummary) GetTotalDiscountedCost() float64 {
	if x != nil {
		return x.TotalDiscountedCost
	}
	return 0
}

func (x *ReservationRecommendationSummary) GetTotalDiscountedMonthlyCost() float64 {
	if x != nil {
		return x.TotalDiscountedMonthlyCost
	}
	return 0
}

func (x *ReservationRecommendationSummary) GetTotalDiscountedYearlyCost() float64 {
	if x != nil {
		return x.TotalDiscountedYearlyCost
	}
	return 0
}

func (x *ReservationRecommendationSummary) GetReductionRate() float64 {
	if x != nil {
		return x.ReductionRate
	}
	return 0
}

func (x *ReservationRecommendationSummary) GetTotalEstimatedMonthlySavings() float64 {
	if x != nil {
		return x.TotalEstimatedMonthlySavings
	}
	return 0
}

type ReservationEstimatedCoverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId         string  `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	ProductCode       string  `protobuf:"bytes,2,opt,name=productCode,proto3" json:"productCode,omitempty"`
	RiCoverage        float64 `protobuf:"fixed64,3,opt,name=riCoverage,proto3" json:"riCoverage,omitempty"`
	Ec2SpCoverage     float64 `protobuf:"fixed64,4,opt,name=ec2SpCoverage,proto3" json:"ec2SpCoverage,omitempty"`
	ComputeSpCoverage float64 `protobuf:"fixed64,5,opt,name=computeSpCoverage,proto3" json:"computeSpCoverage,omitempty"`
}

func (x *ReservationEstimatedCoverage) Reset() {
	*x = ReservationEstimatedCoverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_recommendation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationEstimatedCoverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationEstimatedCoverage) ProtoMessage() {}

func (x *ReservationEstimatedCoverage) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_recommendation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationEstimatedCoverage.ProtoReflect.Descriptor instead.
func (*ReservationEstimatedCoverage) Descriptor() ([]byte, []int) {
	return file_api_aws_recommendation_proto_rawDescGZIP(), []int{4}
}

func (x *ReservationEstimatedCoverage) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ReservationEstimatedCoverage) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *ReservationEstimatedCoverage) GetRiCoverage() float64 {
	if x != nil {
		return x.RiCoverage
	}
	return 0
}

func (x *ReservationEstimatedCoverage) GetEc2SpCoverage() float64 {
	if x != nil {
		return x.Ec2SpCoverage
	}
	return 0
}

func (x *ReservationEstimatedCoverage) GetComputeSpCoverage() float64 {
	if x != nil {
		return x.ComputeSpCoverage
	}
	return 0
}

var File_api_aws_recommendation_proto protoreflect.FileDescriptor

var file_api_aws_recommendation_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x22,
	0xe0, 0x02, 0x0a, 0x12, 0x41, 0x77, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4b, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x4f, 0x0a, 0x11, 0x72, 0x69, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x52, 0x69, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x11, 0x72, 0x69, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4f, 0x0a, 0x11, 0x73, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77,
	0x73, 0x2e, 0x53, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x11, 0x73, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5b, 0x0a, 0x11, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x77, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x11, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x22, 0xe8, 0x05, 0x0a, 0x10, 0x52, 0x69, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x57, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x72, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x57, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x6e, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6f,
	0x6e, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f,
	0x6e, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x6f, 0x6e, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x69, 0x52, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x72, 0x69, 0x52, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x75, 0x70,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x14, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x41, 0x6d, 0x6f, 0x72,
	0x74, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x14, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x41, 0x6d, 0x6f, 0x72, 0x74, 0x69, 0x7a, 0x65,
	0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79,
	0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x14, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x75,
	0x72, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x79, 0x65, 0x61,
	0x72, 0x6c, 0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73,
	0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x79, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x17, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xfc, 0x03,
	0x0a, 0x10, 0x53, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x0f, 0x73, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x70, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x6f, 0x6e, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6f, 0x6e, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x55, 0x6e, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x15, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x79, 0x65, 0x61, 0x72,
	0x6c, 0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x79, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x17,
	0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79,
	0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x17, 0x65,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53,
	0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x72,
	0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x22, 0xea, 0x02, 0x0a,
	0x20, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61,
	0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73,
	0x74, 0x12, 0x3e, 0x0a, 0x1a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f, 0x73,
	0x74, 0x12, 0x3c, 0x0a, 0x19, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x64, 0x59, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x19, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x64, 0x59, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x1c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x61,
	0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1c, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x1c, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x69,
	0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x72, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x63,
	0x32, 0x53, 0x70, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0d, 0x65, 0x63, 0x32, 0x53, 0x70, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x53, 0x70, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x53, 0x70, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x42, 0x66,
	0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x42,
	0x19, 0x41, 0x70, 0x69, 0x41, 0x77, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_aws_recommendation_proto_rawDescOnce sync.Once
	file_api_aws_recommendation_proto_rawDescData = file_api_aws_recommendation_proto_rawDesc
)

func file_api_aws_recommendation_proto_rawDescGZIP() []byte {
	file_api_aws_recommendation_proto_rawDescOnce.Do(func() {
		file_api_aws_recommendation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_aws_recommendation_proto_rawDescData)
	})
	return file_api_aws_recommendation_proto_rawDescData
}

var file_api_aws_recommendation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_aws_recommendation_proto_goTypes = []interface{}{
	(*AwsRecommendations)(nil),               // 0: blueapi.api.aws.AwsRecommendations
	(*RiRecommendation)(nil),                 // 1: blueapi.api.aws.RiRecommendation
	(*SpRecommendation)(nil),                 // 2: blueapi.api.aws.SpRecommendation
	(*ReservationRecommendationSummary)(nil), // 3: blueapi.api.aws.ReservationRecommendationSummary
	(*ReservationEstimatedCoverage)(nil),     // 4: blueapi.api.aws.ReservationEstimatedCoverage
}
var file_api_aws_recommendation_proto_depIdxs = []int32{
	3, // 0: blueapi.api.aws.AwsRecommendations.summary:type_name -> blueapi.api.aws.ReservationRecommendationSummary
	1, // 1: blueapi.api.aws.AwsRecommendations.riRecommendations:type_name -> blueapi.api.aws.RiRecommendation
	2, // 2: blueapi.api.aws.AwsRecommendations.spRecommendations:type_name -> blueapi.api.aws.SpRecommendation
	4, // 3: blueapi.api.aws.AwsRecommendations.estimatedCoverage:type_name -> blueapi.api.aws.ReservationEstimatedCoverage
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_aws_recommendation_proto_init() }
func file_api_aws_recommendation_proto_init() {
	if File_api_aws_recommendation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_aws_recommendation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsRecommendations); i {
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
		file_api_aws_recommendation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiRecommendation); i {
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
		file_api_aws_recommendation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpRecommendation); i {
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
		file_api_aws_recommendation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationRecommendationSummary); i {
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
		file_api_aws_recommendation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationEstimatedCoverage); i {
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
			RawDescriptor: file_api_aws_recommendation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_aws_recommendation_proto_goTypes,
		DependencyIndexes: file_api_aws_recommendation_proto_depIdxs,
		MessageInfos:      file_api_aws_recommendation_proto_msgTypes,
	}.Build()
	File_api_aws_recommendation_proto = out.File
	file_api_aws_recommendation_proto_rawDesc = nil
	file_api_aws_recommendation_proto_goTypes = nil
	file_api_aws_recommendation_proto_depIdxs = nil
}
