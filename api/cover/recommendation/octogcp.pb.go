// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/cover/recommendation/octogcp.proto

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

type OctoGeneratedGCPRecommendations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OctoGeneratedGCPRecommendations) Reset() {
	*x = OctoGeneratedGCPRecommendations{}
	mi := &file_api_cover_recommendation_octogcp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OctoGeneratedGCPRecommendations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OctoGeneratedGCPRecommendations) ProtoMessage() {}

func (x *OctoGeneratedGCPRecommendations) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_recommendation_octogcp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OctoGeneratedGCPRecommendations.ProtoReflect.Descriptor instead.
func (*OctoGeneratedGCPRecommendations) Descriptor() ([]byte, []int) {
	return file_api_cover_recommendation_octogcp_proto_rawDescGZIP(), []int{0}
}

var File_api_cover_recommendation_octogcp_proto protoreflect.FileDescriptor

var file_api_cover_recommendation_octogcp_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x63, 0x74, 0x6f, 0x67,
	0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x67,
	0x63, 0x70, 0x22, 0x21, 0x0a, 0x1f, 0x4f, 0x63, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x47, 0x43, 0x50, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x91, 0x01, 0x0a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x22, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x4f, 0x63, 0x74, 0x6f, 0x47, 0x63, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_cover_recommendation_octogcp_proto_rawDescOnce sync.Once
	file_api_cover_recommendation_octogcp_proto_rawDescData = file_api_cover_recommendation_octogcp_proto_rawDesc
)

func file_api_cover_recommendation_octogcp_proto_rawDescGZIP() []byte {
	file_api_cover_recommendation_octogcp_proto_rawDescOnce.Do(func() {
		file_api_cover_recommendation_octogcp_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_recommendation_octogcp_proto_rawDescData)
	})
	return file_api_cover_recommendation_octogcp_proto_rawDescData
}

var file_api_cover_recommendation_octogcp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_cover_recommendation_octogcp_proto_goTypes = []any{
	(*OctoGeneratedGCPRecommendations)(nil), // 0: blueapi.api.cover.recommendation.octogcp.OctoGeneratedGCPRecommendations
}
var file_api_cover_recommendation_octogcp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_cover_recommendation_octogcp_proto_init() }
func file_api_cover_recommendation_octogcp_proto_init() {
	if File_api_cover_recommendation_octogcp_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cover_recommendation_octogcp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_recommendation_octogcp_proto_goTypes,
		DependencyIndexes: file_api_cover_recommendation_octogcp_proto_depIdxs,
		MessageInfos:      file_api_cover_recommendation_octogcp_proto_msgTypes,
	}.Build()
	File_api_cover_recommendation_octogcp_proto = out.File
	file_api_cover_recommendation_octogcp_proto_rawDesc = nil
	file_api_cover_recommendation_octogcp_proto_goTypes = nil
	file_api_cover_recommendation_octogcp_proto_depIdxs = nil
}
