// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/cover/recommendation/recommendation.proto

package recommendation

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

type RecommendationData struct {
	state                        protoimpl.MessageState        `protogen:"open.v1"`
	AwsRecommendations           *AWSRecommendations           `protobuf:"bytes,1,opt,name=awsRecommendations,proto3" json:"awsRecommendations,omitempty"`
	GcpRecommendations           *GCPRecommendations           `protobuf:"bytes,2,opt,name=gcpRecommendations,proto3" json:"gcpRecommendations,omitempty"`
	AzureCspRecommendations      *AzureCSPRecommendations      `protobuf:"bytes,3,opt,name=azureCspRecommendations,proto3" json:"azureCspRecommendations,omitempty"`
	OctoGeneratedRecommendations *OCTOGeneratedRecommendations `protobuf:"bytes,23,opt,name=octoGeneratedRecommendations,proto3" json:"octoGeneratedRecommendations,omitempty"`
	Target                       string                        `protobuf:"bytes,18,opt,name=target,proto3" json:"target,omitempty"`
	TargetName                   string                        `protobuf:"bytes,19,opt,name=targetName,proto3" json:"targetName,omitempty"`
	ResourceId                   string                        `protobuf:"bytes,4,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	Service                      string                        `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	CostGroup                    string                        `protobuf:"bytes,6,opt,name=costGroup,proto3" json:"costGroup,omitempty"`
	RecommendationGroup          string                        `protobuf:"bytes,7,opt,name=recommendationGroup,proto3" json:"recommendationGroup,omitempty"`
	Category                     string                        `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	Source                       string                        `protobuf:"bytes,9,opt,name=source,proto3" json:"source,omitempty"`
	Id                           string                        `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	LastUpdatedAt                string                        `protobuf:"bytes,11,opt,name=lastUpdatedAt,proto3" json:"lastUpdatedAt,omitempty"`
	Region                       string                        `protobuf:"bytes,12,opt,name=region,proto3" json:"region,omitempty"`
	Recommendation               string                        `protobuf:"bytes,13,opt,name=recommendation,proto3" json:"recommendation,omitempty"`
	EstimatedMonthlyCost         float64                       `protobuf:"fixed64,14,opt,name=estimatedMonthlyCost,proto3" json:"estimatedMonthlyCost,omitempty"`
	EstimatedMonthlySavings      float64                       `protobuf:"fixed64,15,opt,name=estimatedMonthlySavings,proto3" json:"estimatedMonthlySavings,omitempty"`
	EstimatedSavingsPercentage   float64                       `protobuf:"fixed64,16,opt,name=estimatedSavingsPercentage,proto3" json:"estimatedSavingsPercentage,omitempty"`
	ResourceName                 string                        `protobuf:"bytes,17,opt,name=resourceName,proto3" json:"resourceName,omitempty"`
	RestartNeeded                bool                          `protobuf:"varint,20,opt,name=restartNeeded,proto3" json:"restartNeeded,omitempty"`
	RollbackPossible             bool                          `protobuf:"varint,21,opt,name=rollbackPossible,proto3" json:"rollbackPossible,omitempty"`
	LaunchUrl                    string                        `protobuf:"bytes,22,opt,name=launchUrl,proto3" json:"launchUrl,omitempty"`
	Vendor                       string                        `protobuf:"bytes,24,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Status                       string                        `protobuf:"bytes,25,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *RecommendationData) Reset() {
	*x = RecommendationData{}
	mi := &file_api_cover_recommendation_recommendation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationData) ProtoMessage() {}

func (x *RecommendationData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_recommendation_recommendation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationData.ProtoReflect.Descriptor instead.
func (*RecommendationData) Descriptor() ([]byte, []int) {
	return file_api_cover_recommendation_recommendation_proto_rawDescGZIP(), []int{0}
}

func (x *RecommendationData) GetAwsRecommendations() *AWSRecommendations {
	if x != nil {
		return x.AwsRecommendations
	}
	return nil
}

func (x *RecommendationData) GetGcpRecommendations() *GCPRecommendations {
	if x != nil {
		return x.GcpRecommendations
	}
	return nil
}

func (x *RecommendationData) GetAzureCspRecommendations() *AzureCSPRecommendations {
	if x != nil {
		return x.AzureCspRecommendations
	}
	return nil
}

func (x *RecommendationData) GetOctoGeneratedRecommendations() *OCTOGeneratedRecommendations {
	if x != nil {
		return x.OctoGeneratedRecommendations
	}
	return nil
}

func (x *RecommendationData) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *RecommendationData) GetTargetName() string {
	if x != nil {
		return x.TargetName
	}
	return ""
}

func (x *RecommendationData) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *RecommendationData) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *RecommendationData) GetCostGroup() string {
	if x != nil {
		return x.CostGroup
	}
	return ""
}

func (x *RecommendationData) GetRecommendationGroup() string {
	if x != nil {
		return x.RecommendationGroup
	}
	return ""
}

func (x *RecommendationData) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *RecommendationData) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *RecommendationData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RecommendationData) GetLastUpdatedAt() string {
	if x != nil {
		return x.LastUpdatedAt
	}
	return ""
}

func (x *RecommendationData) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RecommendationData) GetRecommendation() string {
	if x != nil {
		return x.Recommendation
	}
	return ""
}

func (x *RecommendationData) GetEstimatedMonthlyCost() float64 {
	if x != nil {
		return x.EstimatedMonthlyCost
	}
	return 0
}

func (x *RecommendationData) GetEstimatedMonthlySavings() float64 {
	if x != nil {
		return x.EstimatedMonthlySavings
	}
	return 0
}

func (x *RecommendationData) GetEstimatedSavingsPercentage() float64 {
	if x != nil {
		return x.EstimatedSavingsPercentage
	}
	return 0
}

func (x *RecommendationData) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *RecommendationData) GetRestartNeeded() bool {
	if x != nil {
		return x.RestartNeeded
	}
	return false
}

func (x *RecommendationData) GetRollbackPossible() bool {
	if x != nil {
		return x.RollbackPossible
	}
	return false
}

func (x *RecommendationData) GetLaunchUrl() string {
	if x != nil {
		return x.LaunchUrl
	}
	return ""
}

func (x *RecommendationData) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *RecommendationData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type OCTOGeneratedRecommendations struct {
	state         protoimpl.MessageState                `protogen:"open.v1"`
	Aws           *OctoGeneratedAWSRecommendations      `protobuf:"bytes,1,opt,name=aws,proto3" json:"aws,omitempty"`
	Gcp           *OctoGeneratedGCPRecommendations      `protobuf:"bytes,2,opt,name=gcp,proto3" json:"gcp,omitempty"`
	Azurecsp      *OctoGeneratedAzureCSPRecommendations `protobuf:"bytes,3,opt,name=azurecsp,proto3" json:"azurecsp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OCTOGeneratedRecommendations) Reset() {
	*x = OCTOGeneratedRecommendations{}
	mi := &file_api_cover_recommendation_recommendation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OCTOGeneratedRecommendations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCTOGeneratedRecommendations) ProtoMessage() {}

func (x *OCTOGeneratedRecommendations) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_recommendation_recommendation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCTOGeneratedRecommendations.ProtoReflect.Descriptor instead.
func (*OCTOGeneratedRecommendations) Descriptor() ([]byte, []int) {
	return file_api_cover_recommendation_recommendation_proto_rawDescGZIP(), []int{1}
}

func (x *OCTOGeneratedRecommendations) GetAws() *OctoGeneratedAWSRecommendations {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *OCTOGeneratedRecommendations) GetGcp() *OctoGeneratedGCPRecommendations {
	if x != nil {
		return x.Gcp
	}
	return nil
}

func (x *OCTOGeneratedRecommendations) GetAzurecsp() *OctoGeneratedAzureCSPRecommendations {
	if x != nil {
		return x.Azurecsp
	}
	return nil
}

var File_api_cover_recommendation_recommendation_proto protoreflect.FileDescriptor

var file_api_cover_recommendation_recommendation_proto_rawDesc = string([]byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x63, 0x74, 0x6f,
	0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x63, 0x74, 0x6f, 0x67, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x63, 0x74, 0x6f,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x63, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x63, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc9, 0x09, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x68, 0x0a, 0x12, 0x61, 0x77, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x41, 0x57, 0x53, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x12, 0x61, 0x77,
	0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x68, 0x0a, 0x12, 0x67, 0x63, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x62,
	0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x67, 0x63, 0x70, 0x2e, 0x47, 0x43, 0x50, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x12, 0x67, 0x63, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x7c, 0x0a, 0x17, 0x61, 0x7a,
	0x75, 0x72, 0x65, 0x43, 0x73, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61,
	0x7a, 0x75, 0x72, 0x65, 0x63, 0x73, 0x70, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x53, 0x50,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x17, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x73, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x1c, 0x6f, 0x63, 0x74,
	0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3e, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4f, 0x43, 0x54, 0x4f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x1c, 0x6f, 0x63, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a,
	0x13, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x14, 0x65, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f, 0x73,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x38, 0x0a,
	0x17, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c,
	0x79, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x17,
	0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79,
	0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3e, 0x0a, 0x1a, 0x65, 0x73, 0x74, 0x69, 0x6d,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1a, 0x65, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x65, 0x65, 0x64, 0x65,
	0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x73,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x72, 0x6f, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x55, 0x72, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc9, 0x02, 0x0a, 0x1c,
	0x4f, 0x43, 0x54, 0x4f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5b, 0x0a, 0x03,
	0x61, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x63, 0x74,
	0x6f, 0x61, 0x77, 0x73, 0x2e, 0x4f, 0x63, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x57, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x5b, 0x0a, 0x03, 0x67, 0x63, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x67, 0x63,
	0x70, 0x2e, 0x4f, 0x63, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x47,
	0x43, 0x50, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x03, 0x67, 0x63, 0x70, 0x12, 0x6f, 0x0a, 0x08, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x63,
	0x73, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x63, 0x74, 0x6f,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x63, 0x73, 0x70, 0x2e, 0x4f, 0x63, 0x74, 0x6f, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x53, 0x50, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x08, 0x61,
	0x7a, 0x75, 0x72, 0x65, 0x63, 0x73, 0x70, 0x42, 0x8a, 0x01, 0x0a, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1b, 0x41, 0x70, 0x69, 0x43,
	0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_cover_recommendation_recommendation_proto_rawDescOnce sync.Once
	file_api_cover_recommendation_recommendation_proto_rawDescData []byte
)

func file_api_cover_recommendation_recommendation_proto_rawDescGZIP() []byte {
	file_api_cover_recommendation_recommendation_proto_rawDescOnce.Do(func() {
		file_api_cover_recommendation_recommendation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_cover_recommendation_recommendation_proto_rawDesc), len(file_api_cover_recommendation_recommendation_proto_rawDesc)))
	})
	return file_api_cover_recommendation_recommendation_proto_rawDescData
}

var file_api_cover_recommendation_recommendation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_cover_recommendation_recommendation_proto_goTypes = []any{
	(*RecommendationData)(nil),                   // 0: blueapi.api.cover.recommendation.RecommendationData
	(*OCTOGeneratedRecommendations)(nil),         // 1: blueapi.api.cover.recommendation.OCTOGeneratedRecommendations
	(*AWSRecommendations)(nil),                   // 2: blueapi.api.cover.recommendation.aws.AWSRecommendations
	(*GCPRecommendations)(nil),                   // 3: blueapi.api.cover.recommendation.gcp.GCPRecommendations
	(*AzureCSPRecommendations)(nil),              // 4: blueapi.api.cover.recommendation.azurecsp.AzureCSPRecommendations
	(*OctoGeneratedAWSRecommendations)(nil),      // 5: blueapi.api.cover.recommendation.octoaws.OctoGeneratedAWSRecommendations
	(*OctoGeneratedGCPRecommendations)(nil),      // 6: blueapi.api.cover.recommendation.octogcp.OctoGeneratedGCPRecommendations
	(*OctoGeneratedAzureCSPRecommendations)(nil), // 7: blueapi.api.cover.recommendation.octoazurecsp.OctoGeneratedAzureCSPRecommendations
}
var file_api_cover_recommendation_recommendation_proto_depIdxs = []int32{
	2, // 0: blueapi.api.cover.recommendation.RecommendationData.awsRecommendations:type_name -> blueapi.api.cover.recommendation.aws.AWSRecommendations
	3, // 1: blueapi.api.cover.recommendation.RecommendationData.gcpRecommendations:type_name -> blueapi.api.cover.recommendation.gcp.GCPRecommendations
	4, // 2: blueapi.api.cover.recommendation.RecommendationData.azureCspRecommendations:type_name -> blueapi.api.cover.recommendation.azurecsp.AzureCSPRecommendations
	1, // 3: blueapi.api.cover.recommendation.RecommendationData.octoGeneratedRecommendations:type_name -> blueapi.api.cover.recommendation.OCTOGeneratedRecommendations
	5, // 4: blueapi.api.cover.recommendation.OCTOGeneratedRecommendations.aws:type_name -> blueapi.api.cover.recommendation.octoaws.OctoGeneratedAWSRecommendations
	6, // 5: blueapi.api.cover.recommendation.OCTOGeneratedRecommendations.gcp:type_name -> blueapi.api.cover.recommendation.octogcp.OctoGeneratedGCPRecommendations
	7, // 6: blueapi.api.cover.recommendation.OCTOGeneratedRecommendations.azurecsp:type_name -> blueapi.api.cover.recommendation.octoazurecsp.OctoGeneratedAzureCSPRecommendations
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_cover_recommendation_recommendation_proto_init() }
func file_api_cover_recommendation_recommendation_proto_init() {
	if File_api_cover_recommendation_recommendation_proto != nil {
		return
	}
	file_api_cover_recommendation_octoaws_proto_init()
	file_api_cover_recommendation_octogcp_proto_init()
	file_api_cover_recommendation_octoazurecsp_proto_init()
	file_api_cover_recommendation_gcp_proto_init()
	file_api_cover_recommendation_azurecsp_proto_init()
	file_api_cover_recommendation_aws_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_cover_recommendation_recommendation_proto_rawDesc), len(file_api_cover_recommendation_recommendation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_recommendation_recommendation_proto_goTypes,
		DependencyIndexes: file_api_cover_recommendation_recommendation_proto_depIdxs,
		MessageInfos:      file_api_cover_recommendation_recommendation_proto_msgTypes,
	}.Build()
	File_api_cover_recommendation_recommendation_proto = out.File
	file_api_cover_recommendation_recommendation_proto_goTypes = nil
	file_api_cover_recommendation_recommendation_proto_depIdxs = nil
}
