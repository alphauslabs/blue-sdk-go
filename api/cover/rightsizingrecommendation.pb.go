// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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

const file_api_cover_rightsizingrecommendation_proto_rawDesc = "" +
	"\n" +
	")api/cover/rightsizingrecommendation.proto\x12\x11blueapi.api.cover\"\xc7\x04\n" +
	"\fResourceData\x12\x16\n" +
	"\x06vendor\x18\x01 \x01(\tR\x06vendor\x12\x1c\n" +
	"\taccountId\x18\x02 \x01(\tR\taccountId\x12\x1e\n" +
	"\n" +
	"resourceId\x18\x03 \x01(\tR\n" +
	"resourceId\x12\x16\n" +
	"\x06region\x18\x04 \x01(\tR\x06region\x12\"\n" +
	"\fresourceType\x18\x05 \x01(\tR\fresourceType\x12[\n" +
	"\x14recommendationDetail\x18\x06 \x03(\v2'.blueapi.api.cover.RecommendationDetailR\x14recommendationDetail\x12 \n" +
	"\vcurrentCost\x18\a \x01(\x01R\vcurrentCost\x12\"\n" +
	"\fresourceName\x18\b \x01(\tR\fresourceName\x12(\n" +
	"\x0fconsumedService\x18\t \x01(\tR\x0fconsumedService\x12,\n" +
	"\x11maxCpuUtilization\x18\n" +
	" \x01(\x01R\x11maxCpuUtilization\x124\n" +
	"\x15maxStorageUtilization\x18\v \x01(\x01R\x15maxStorageUtilization\x122\n" +
	"\x14maxMemoryUtilization\x18\f \x01(\x01R\x14maxMemoryUtilization\x12(\n" +
	"\x0fnetworkCapacity\x18\r \x01(\tR\x0fnetworkCapacity\x12\x16\n" +
	"\x06status\x18\x0e \x01(\tR\x06status\"\xe2\x01\n" +
	"\x14RecommendationDetail\x12&\n" +
	"\x0erecommendation\x18\x01 \x01(\tR\x0erecommendation\x128\n" +
	"\x17recommendedResourceType\x18\x02 \x01(\tR\x17recommendedResourceType\x12$\n" +
	"\restimatedCost\x18\x03 \x01(\x01R\restimatedCost\x12*\n" +
	"\x10estimatedSavings\x18\x04 \x01(\x01R\x10estimatedSavings\x12\x16\n" +
	"\x06region\x18\x05 \x01(\tR\x06regionBw\n" +
	"\x1fcloud.alphaus.blueapi.api.coverB&ApiCoverRightSizingRecommendationProtoZ,github.com/alphauslabs/blue-sdk-go/api/coverb\x06proto3"

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
