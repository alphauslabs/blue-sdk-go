// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/cover/rightsizingrecommendation.proto

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

type ResourceData struct {
	state                 protoimpl.MessageState  `protogen:"open.v1"`
	Vendor                string                  `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
	AccountId             string                  `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	ResourceId            string                  `protobuf:"bytes,3,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	Region                string                  `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	ResourceType          string                  `protobuf:"bytes,5,opt,name=resourceType,proto3" json:"resourceType,omitempty"`
	RecommendationDetail  []*RecommendationDetail `protobuf:"bytes,6,rep,name=recommendationDetail,proto3" json:"recommendationDetail,omitempty"`
	CurrentCost           float64                 `protobuf:"fixed64,7,opt,name=currentCost,proto3" json:"currentCost,omitempty"`
	ResourceName          string                  `protobuf:"bytes,8,opt,name=resourceName,proto3" json:"resourceName,omitempty"`
	ConsumedService       string                  `protobuf:"bytes,9,opt,name=consumedService,proto3" json:"consumedService,omitempty"`
	MaxCpuUtilization     float64                 `protobuf:"fixed64,10,opt,name=maxCpuUtilization,proto3" json:"maxCpuUtilization,omitempty"`
	MaxStorageUtilization float64                 `protobuf:"fixed64,11,opt,name=maxStorageUtilization,proto3" json:"maxStorageUtilization,omitempty"`
	MaxMemoryUtilization  float64                 `protobuf:"fixed64,12,opt,name=maxMemoryUtilization,proto3" json:"maxMemoryUtilization,omitempty"`
	NetworkCapacity       string                  `protobuf:"bytes,13,opt,name=networkCapacity,proto3" json:"networkCapacity,omitempty"`
	Status                string                  `protobuf:"bytes,14,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ResourceData) Reset() {
	*x = ResourceData{}
	mi := &file_api_cover_rightsizingrecommendation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceData) ProtoMessage() {}

func (x *ResourceData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_rightsizingrecommendation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceData.ProtoReflect.Descriptor instead.
func (*ResourceData) Descriptor() ([]byte, []int) {
	return file_api_cover_rightsizingrecommendation_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceData) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *ResourceData) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ResourceData) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *ResourceData) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ResourceData) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *ResourceData) GetRecommendationDetail() []*RecommendationDetail {
	if x != nil {
		return x.RecommendationDetail
	}
	return nil
}

func (x *ResourceData) GetCurrentCost() float64 {
	if x != nil {
		return x.CurrentCost
	}
	return 0
}

func (x *ResourceData) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ResourceData) GetConsumedService() string {
	if x != nil {
		return x.ConsumedService
	}
	return ""
}

func (x *ResourceData) GetMaxCpuUtilization() float64 {
	if x != nil {
		return x.MaxCpuUtilization
	}
	return 0
}

func (x *ResourceData) GetMaxStorageUtilization() float64 {
	if x != nil {
		return x.MaxStorageUtilization
	}
	return 0
}

func (x *ResourceData) GetMaxMemoryUtilization() float64 {
	if x != nil {
		return x.MaxMemoryUtilization
	}
	return 0
}

func (x *ResourceData) GetNetworkCapacity() string {
	if x != nil {
		return x.NetworkCapacity
	}
	return ""
}

func (x *ResourceData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type RecommendationDetail struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	Recommendation          string                 `protobuf:"bytes,1,opt,name=recommendation,proto3" json:"recommendation,omitempty"`
	RecommendedResourceType string                 `protobuf:"bytes,2,opt,name=recommendedResourceType,proto3" json:"recommendedResourceType,omitempty"`
	EstimatedCost           float64                `protobuf:"fixed64,3,opt,name=estimatedCost,proto3" json:"estimatedCost,omitempty"`
	EstimatedSavings        float64                `protobuf:"fixed64,4,opt,name=estimatedSavings,proto3" json:"estimatedSavings,omitempty"`
	Region                  string                 `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *RecommendationDetail) Reset() {
	*x = RecommendationDetail{}
	mi := &file_api_cover_rightsizingrecommendation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendationDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationDetail) ProtoMessage() {}

func (x *RecommendationDetail) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_rightsizingrecommendation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationDetail.ProtoReflect.Descriptor instead.
func (*RecommendationDetail) Descriptor() ([]byte, []int) {
	return file_api_cover_rightsizingrecommendation_proto_rawDescGZIP(), []int{1}
}

func (x *RecommendationDetail) GetRecommendation() string {
	if x != nil {
		return x.Recommendation
	}
	return ""
}

func (x *RecommendationDetail) GetRecommendedResourceType() string {
	if x != nil {
		return x.RecommendedResourceType
	}
	return ""
}

func (x *RecommendationDetail) GetEstimatedCost() float64 {
	if x != nil {
		return x.EstimatedCost
	}
	return 0
}

func (x *RecommendationDetail) GetEstimatedSavings() float64 {
	if x != nil {
		return x.EstimatedSavings
	}
	return 0
}

func (x *RecommendationDetail) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

var File_api_cover_rightsizingrecommendation_proto protoreflect.FileDescriptor

var file_api_cover_rightsizingrecommendation_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x22, 0xc7,
	0x04, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x5b, 0x0a, 0x14, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x14, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x11, 0x6d, 0x61, 0x78, 0x43, 0x70, 0x75, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x43, 0x70,
	0x75, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x15,
	0x6d, 0x61, 0x78, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x6d, 0x61, 0x78,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55,
	0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x14, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x74, 0x69, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xe2, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x17, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x65, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x10, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x53, 0x61,
	0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x77, 0x0a,
	0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62,
	0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x42, 0x26, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x69, 0x67, 0x68, 0x74, 0x53,
	0x69, 0x7a, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_cover_rightsizingrecommendation_proto_rawDescOnce sync.Once
	file_api_cover_rightsizingrecommendation_proto_rawDescData []byte
)

func file_api_cover_rightsizingrecommendation_proto_rawDescGZIP() []byte {
	file_api_cover_rightsizingrecommendation_proto_rawDescOnce.Do(func() {
		file_api_cover_rightsizingrecommendation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_cover_rightsizingrecommendation_proto_rawDesc), len(file_api_cover_rightsizingrecommendation_proto_rawDesc)))
	})
	return file_api_cover_rightsizingrecommendation_proto_rawDescData
}

var file_api_cover_rightsizingrecommendation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_cover_rightsizingrecommendation_proto_goTypes = []any{
	(*ResourceData)(nil),         // 0: blueapi.api.cover.ResourceData
	(*RecommendationDetail)(nil), // 1: blueapi.api.cover.RecommendationDetail
}
var file_api_cover_rightsizingrecommendation_proto_depIdxs = []int32{
	1, // 0: blueapi.api.cover.ResourceData.recommendationDetail:type_name -> blueapi.api.cover.RecommendationDetail
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_cover_rightsizingrecommendation_proto_init() }
func file_api_cover_rightsizingrecommendation_proto_init() {
	if File_api_cover_rightsizingrecommendation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_cover_rightsizingrecommendation_proto_rawDesc), len(file_api_cover_rightsizingrecommendation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_rightsizingrecommendation_proto_goTypes,
		DependencyIndexes: file_api_cover_rightsizingrecommendation_proto_depIdxs,
		MessageInfos:      file_api_cover_rightsizingrecommendation_proto_msgTypes,
	}.Build()
	File_api_cover_rightsizingrecommendation_proto = out.File
	file_api_cover_rightsizingrecommendation_proto_goTypes = nil
	file_api_cover_rightsizingrecommendation_proto_depIdxs = nil
}
