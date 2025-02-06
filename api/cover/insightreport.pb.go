// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/cover/insightreport.proto

package cover

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

type ExecutiveSummary struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The total cost usage of the current period
	CostUsage float64 `protobuf:"fixed64,1,opt,name=costUsage,proto3" json:"costUsage,omitempty"`
	// The total cost usage of the previous period
	PreviousPeriodCostUsage float64 `protobuf:"fixed64,2,opt,name=previousPeriodCostUsage,proto3" json:"previousPeriodCostUsage,omitempty"`
	// The status of the cost usage compared to the previous period. "increased", "decreased"
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	// The percentage changed compared to the previous period
	PercentageChanged float64 `protobuf:"fixed64,4,opt,name=percentageChanged,proto3" json:"percentageChanged,omitempty"`
	// The average monthly changed
	AverageMonthlyChanged        float64 `protobuf:"fixed64,5,opt,name=averageMonthlyChanged,proto3" json:"averageMonthlyChanged,omitempty"`
	DifferenceFromPreviousPeriod float64 `protobuf:"fixed64,6,opt,name=differenceFromPreviousPeriod,proto3" json:"differenceFromPreviousPeriod,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *ExecutiveSummary) Reset() {
	*x = ExecutiveSummary{}
	mi := &file_api_cover_insightreport_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutiveSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutiveSummary) ProtoMessage() {}

func (x *ExecutiveSummary) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_insightreport_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutiveSummary.ProtoReflect.Descriptor instead.
func (*ExecutiveSummary) Descriptor() ([]byte, []int) {
	return file_api_cover_insightreport_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutiveSummary) GetCostUsage() float64 {
	if x != nil {
		return x.CostUsage
	}
	return 0
}

func (x *ExecutiveSummary) GetPreviousPeriodCostUsage() float64 {
	if x != nil {
		return x.PreviousPeriodCostUsage
	}
	return 0
}

func (x *ExecutiveSummary) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ExecutiveSummary) GetPercentageChanged() float64 {
	if x != nil {
		return x.PercentageChanged
	}
	return 0
}

func (x *ExecutiveSummary) GetAverageMonthlyChanged() float64 {
	if x != nil {
		return x.AverageMonthlyChanged
	}
	return 0
}

func (x *ExecutiveSummary) GetDifferenceFromPreviousPeriod() float64 {
	if x != nil {
		return x.DifferenceFromPreviousPeriod
	}
	return 0
}

type OptimizationRecommendationSummary struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Potential savings per month
	PotentialSavings []float64 `protobuf:"fixed64,1,rep,packed,name=potentialSavings,proto3" json:"potentialSavings,omitempty"`
	// The recommended action to perform to achieve potential savings
	Action        []string `protobuf:"bytes,2,rep,name=action,proto3" json:"action,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OptimizationRecommendationSummary) Reset() {
	*x = OptimizationRecommendationSummary{}
	mi := &file_api_cover_insightreport_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptimizationRecommendationSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptimizationRecommendationSummary) ProtoMessage() {}

func (x *OptimizationRecommendationSummary) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_insightreport_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptimizationRecommendationSummary.ProtoReflect.Descriptor instead.
func (*OptimizationRecommendationSummary) Descriptor() ([]byte, []int) {
	return file_api_cover_insightreport_proto_rawDescGZIP(), []int{1}
}

func (x *OptimizationRecommendationSummary) GetPotentialSavings() []float64 {
	if x != nil {
		return x.PotentialSavings
	}
	return nil
}

func (x *OptimizationRecommendationSummary) GetAction() []string {
	if x != nil {
		return x.Action
	}
	return nil
}

type SavingsSummary struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Total number of detected recommendation for a given period
	TotalRecommendations int64 `protobuf:"varint,1,opt,name=totalRecommendations,proto3" json:"totalRecommendations,omitempty"`
	// Total number of executed recommendations for a given period
	TotalExecutedRecommendations int64 `protobuf:"varint,2,opt,name=totalExecutedRecommendations,proto3" json:"totalExecutedRecommendations,omitempty"`
	// Total of estimated savings for a given period, as a result of executed recommendations
	TotalEstimatedSavings float64 `protobuf:"fixed64,3,opt,name=totalEstimatedSavings,proto3" json:"totalEstimatedSavings,omitempty"`
	// Total of estimated cost for a given period, as a result of executed recommendations
	TotalEstimatedCost float64 `protobuf:"fixed64,4,opt,name=totalEstimatedCost,proto3" json:"totalEstimatedCost,omitempty"`
	// Total of estimated percentage savings for a given period, as a result of executed recommendations
	PercentageSavings float64 `protobuf:"fixed64,5,opt,name=percentageSavings,proto3" json:"percentageSavings,omitempty"`
	// The most effective recommendation
	MostEffectiveRecommendation string `protobuf:"bytes,6,opt,name=mostEffectiveRecommendation,proto3" json:"mostEffectiveRecommendation,omitempty"`
	// The most effective recommendation savings
	MostEffectiveRecommendationEstimatedSavings float64 `protobuf:"fixed64,7,opt,name=mostEffectiveRecommendationEstimatedSavings,proto3" json:"mostEffectiveRecommendationEstimatedSavings,omitempty"`
	// The most effective recommendation cost
	MostEffectiveRecommendationEstimatedCost float64 `protobuf:"fixed64,8,opt,name=mostEffectiveRecommendationEstimatedCost,proto3" json:"mostEffectiveRecommendationEstimatedCost,omitempty"`
	// The most effective recommendation percentage savings
	MostEffectiveRecommendationPercentageSavings float64 `protobuf:"fixed64,9,opt,name=mostEffectiveRecommendationPercentageSavings,proto3" json:"mostEffectiveRecommendationPercentageSavings,omitempty"`
	unknownFields                                protoimpl.UnknownFields
	sizeCache                                    protoimpl.SizeCache
}

func (x *SavingsSummary) Reset() {
	*x = SavingsSummary{}
	mi := &file_api_cover_insightreport_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SavingsSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavingsSummary) ProtoMessage() {}

func (x *SavingsSummary) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_insightreport_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavingsSummary.ProtoReflect.Descriptor instead.
func (*SavingsSummary) Descriptor() ([]byte, []int) {
	return file_api_cover_insightreport_proto_rawDescGZIP(), []int{2}
}

func (x *SavingsSummary) GetTotalRecommendations() int64 {
	if x != nil {
		return x.TotalRecommendations
	}
	return 0
}

func (x *SavingsSummary) GetTotalExecutedRecommendations() int64 {
	if x != nil {
		return x.TotalExecutedRecommendations
	}
	return 0
}

func (x *SavingsSummary) GetTotalEstimatedSavings() float64 {
	if x != nil {
		return x.TotalEstimatedSavings
	}
	return 0
}

func (x *SavingsSummary) GetTotalEstimatedCost() float64 {
	if x != nil {
		return x.TotalEstimatedCost
	}
	return 0
}

func (x *SavingsSummary) GetPercentageSavings() float64 {
	if x != nil {
		return x.PercentageSavings
	}
	return 0
}

func (x *SavingsSummary) GetMostEffectiveRecommendation() string {
	if x != nil {
		return x.MostEffectiveRecommendation
	}
	return ""
}

func (x *SavingsSummary) GetMostEffectiveRecommendationEstimatedSavings() float64 {
	if x != nil {
		return x.MostEffectiveRecommendationEstimatedSavings
	}
	return 0
}

func (x *SavingsSummary) GetMostEffectiveRecommendationEstimatedCost() float64 {
	if x != nil {
		return x.MostEffectiveRecommendationEstimatedCost
	}
	return 0
}

func (x *SavingsSummary) GetMostEffectiveRecommendationPercentageSavings() float64 {
	if x != nil {
		return x.MostEffectiveRecommendationPercentageSavings
	}
	return 0
}

var File_api_cover_insightreport_proto protoreflect.FileDescriptor

var file_api_cover_insightreport_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x22, 0xaa, 0x02, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x17, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x11, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x1c, 0x64,
	0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x1c, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22,
	0x67, 0x0a, 0x21, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x10,
	0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x05, 0x0a, 0x0e, 0x53, 0x61, 0x76,
	0x69, 0x6e, 0x67, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x14, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x42, 0x0a, 0x1c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x64, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x40, 0x0a, 0x1b, 0x6d, 0x6f, 0x73, 0x74, 0x45,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x6d, 0x6f,
	0x73, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x60, 0x0a, 0x2b, 0x6d, 0x6f, 0x73,
	0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x2b,
	0x6d, 0x6f, 0x73, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x73, 0x74, 0x69, 0x6d,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5a, 0x0a, 0x28, 0x6d,
	0x6f, 0x73, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61,
	0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x28, 0x6d,
	0x6f, 0x73, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61,
	0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x62, 0x0a, 0x2c, 0x6d, 0x6f, 0x73, 0x74, 0x45,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x2c, 0x6d,
	0x6f, 0x73, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x6b, 0x0a, 0x1f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x1a,
	0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_cover_insightreport_proto_rawDescOnce sync.Once
	file_api_cover_insightreport_proto_rawDescData []byte
)

func file_api_cover_insightreport_proto_rawDescGZIP() []byte {
	file_api_cover_insightreport_proto_rawDescOnce.Do(func() {
		file_api_cover_insightreport_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_cover_insightreport_proto_rawDesc), len(file_api_cover_insightreport_proto_rawDesc)))
	})
	return file_api_cover_insightreport_proto_rawDescData
}

var file_api_cover_insightreport_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_cover_insightreport_proto_goTypes = []any{
	(*ExecutiveSummary)(nil),                  // 0: blueapi.api.cover.ExecutiveSummary
	(*OptimizationRecommendationSummary)(nil), // 1: blueapi.api.cover.OptimizationRecommendationSummary
	(*SavingsSummary)(nil),                    // 2: blueapi.api.cover.SavingsSummary
}
var file_api_cover_insightreport_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_cover_insightreport_proto_init() }
func file_api_cover_insightreport_proto_init() {
	if File_api_cover_insightreport_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_cover_insightreport_proto_rawDesc), len(file_api_cover_insightreport_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_insightreport_proto_goTypes,
		DependencyIndexes: file_api_cover_insightreport_proto_depIdxs,
		MessageInfos:      file_api_cover_insightreport_proto_msgTypes,
	}.Build()
	File_api_cover_insightreport_proto = out.File
	file_api_cover_insightreport_proto_goTypes = nil
	file_api_cover_insightreport_proto_depIdxs = nil
}
