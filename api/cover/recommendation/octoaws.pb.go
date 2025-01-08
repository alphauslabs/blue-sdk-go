// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: api/cover/recommendation/octoaws.proto

package recommendation

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

type OctoGeneratedAWSRecommendations struct {
	state                 protoimpl.MessageState                   `protogen:"open.v1"`
	CurrentDetails        *OctoGeneratedAWSRecommendations_Details `protobuf:"bytes,1,opt,name=currentDetails,proto3" json:"currentDetails,omitempty"`
	RecommendationDetails *OctoGeneratedAWSRecommendations_Details `protobuf:"bytes,2,opt,name=recommendationDetails,proto3" json:"recommendationDetails,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *OctoGeneratedAWSRecommendations) Reset() {
	*x = OctoGeneratedAWSRecommendations{}
	mi := &file_api_cover_recommendation_octoaws_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OctoGeneratedAWSRecommendations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OctoGeneratedAWSRecommendations) ProtoMessage() {}

func (x *OctoGeneratedAWSRecommendations) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_recommendation_octoaws_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OctoGeneratedAWSRecommendations.ProtoReflect.Descriptor instead.
func (*OctoGeneratedAWSRecommendations) Descriptor() ([]byte, []int) {
	return file_api_cover_recommendation_octoaws_proto_rawDescGZIP(), []int{0}
}

func (x *OctoGeneratedAWSRecommendations) GetCurrentDetails() *OctoGeneratedAWSRecommendations_Details {
	if x != nil {
		return x.CurrentDetails
	}
	return nil
}

func (x *OctoGeneratedAWSRecommendations) GetRecommendationDetails() *OctoGeneratedAWSRecommendations_Details {
	if x != nil {
		return x.RecommendationDetails
	}
	return nil
}

type OctoGeneratedAWSRecommendations_Details struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Ec2Details    *AWSResourceDetails_EC2Details `protobuf:"bytes,1,opt,name=ec2Details,proto3" json:"ec2Details,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OctoGeneratedAWSRecommendations_Details) Reset() {
	*x = OctoGeneratedAWSRecommendations_Details{}
	mi := &file_api_cover_recommendation_octoaws_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OctoGeneratedAWSRecommendations_Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OctoGeneratedAWSRecommendations_Details) ProtoMessage() {}

func (x *OctoGeneratedAWSRecommendations_Details) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_recommendation_octoaws_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OctoGeneratedAWSRecommendations_Details.ProtoReflect.Descriptor instead.
func (*OctoGeneratedAWSRecommendations_Details) Descriptor() ([]byte, []int) {
	return file_api_cover_recommendation_octoaws_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OctoGeneratedAWSRecommendations_Details) GetEc2Details() *AWSResourceDetails_EC2Details {
	if x != nil {
		return x.Ec2Details
	}
	return nil
}

var File_api_cover_recommendation_octoaws_proto protoreflect.FileDescriptor

var file_api_cover_recommendation_octoaws_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x63, 0x74, 0x6f, 0x61,
	0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x61,
	0x77, 0x73, 0x1a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x77, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x1f, 0x4f, 0x63, 0x74, 0x6f, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x57, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x79, 0x0a, 0x0e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x51, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x61, 0x77, 0x73, 0x2e, 0x4f, 0x63,
	0x74, 0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x57, 0x53, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x87, 0x01, 0x0a, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x61, 0x77, 0x73,
	0x2e, 0x4f, 0x63, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x57,
	0x53, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a,
	0x6e, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x63, 0x0a, 0x0a, 0x65, 0x63,
	0x32, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x45, 0x43, 0x32, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x0a, 0x65, 0x63, 0x32, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42,
	0x91, 0x01, 0x0a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75,
	0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x22, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x4f, 0x63, 0x74, 0x6f,
	0x41, 0x77, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62,
	0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cover_recommendation_octoaws_proto_rawDescOnce sync.Once
	file_api_cover_recommendation_octoaws_proto_rawDescData = file_api_cover_recommendation_octoaws_proto_rawDesc
)

func file_api_cover_recommendation_octoaws_proto_rawDescGZIP() []byte {
	file_api_cover_recommendation_octoaws_proto_rawDescOnce.Do(func() {
		file_api_cover_recommendation_octoaws_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_recommendation_octoaws_proto_rawDescData)
	})
	return file_api_cover_recommendation_octoaws_proto_rawDescData
}

var file_api_cover_recommendation_octoaws_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_cover_recommendation_octoaws_proto_goTypes = []any{
	(*OctoGeneratedAWSRecommendations)(nil),         // 0: blueapi.api.cover.recommendation.octoaws.OctoGeneratedAWSRecommendations
	(*OctoGeneratedAWSRecommendations_Details)(nil), // 1: blueapi.api.cover.recommendation.octoaws.OctoGeneratedAWSRecommendations.Details
	(*AWSResourceDetails_EC2Details)(nil),           // 2: blueapi.api.cover.recommendation.aws.AWSResourceDetails.EC2Details
}
var file_api_cover_recommendation_octoaws_proto_depIdxs = []int32{
	1, // 0: blueapi.api.cover.recommendation.octoaws.OctoGeneratedAWSRecommendations.currentDetails:type_name -> blueapi.api.cover.recommendation.octoaws.OctoGeneratedAWSRecommendations.Details
	1, // 1: blueapi.api.cover.recommendation.octoaws.OctoGeneratedAWSRecommendations.recommendationDetails:type_name -> blueapi.api.cover.recommendation.octoaws.OctoGeneratedAWSRecommendations.Details
	2, // 2: blueapi.api.cover.recommendation.octoaws.OctoGeneratedAWSRecommendations.Details.ec2Details:type_name -> blueapi.api.cover.recommendation.aws.AWSResourceDetails.EC2Details
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_cover_recommendation_octoaws_proto_init() }
func file_api_cover_recommendation_octoaws_proto_init() {
	if File_api_cover_recommendation_octoaws_proto != nil {
		return
	}
	file_api_cover_recommendation_aws_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cover_recommendation_octoaws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_recommendation_octoaws_proto_goTypes,
		DependencyIndexes: file_api_cover_recommendation_octoaws_proto_depIdxs,
		MessageInfos:      file_api_cover_recommendation_octoaws_proto_msgTypes,
	}.Build()
	File_api_cover_recommendation_octoaws_proto = out.File
	file_api_cover_recommendation_octoaws_proto_rawDesc = nil
	file_api_cover_recommendation_octoaws_proto_goTypes = nil
	file_api_cover_recommendation_octoaws_proto_depIdxs = nil
}
