// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: api/cover/recommendation/azurecsp.proto

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

type AzureCSPRecommendations struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AzureCSPRecommendations) Reset() {
	*x = AzureCSPRecommendations{}
	mi := &file_api_cover_recommendation_azurecsp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureCSPRecommendations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureCSPRecommendations) ProtoMessage() {}

func (x *AzureCSPRecommendations) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_recommendation_azurecsp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureCSPRecommendations.ProtoReflect.Descriptor instead.
func (*AzureCSPRecommendations) Descriptor() ([]byte, []int) {
	return file_api_cover_recommendation_azurecsp_proto_rawDescGZIP(), []int{0}
}

var File_api_cover_recommendation_azurecsp_proto protoreflect.FileDescriptor

var file_api_cover_recommendation_azurecsp_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65,
	0x63, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x7a, 0x75, 0x72,
	0x65, 0x63, 0x73, 0x70, 0x22, 0x19, 0x0a, 0x17, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x53, 0x50,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x92, 0x01, 0x0a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75,
	0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x23, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x7a, 0x75, 0x72,
	0x65, 0x43, 0x73, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cover_recommendation_azurecsp_proto_rawDescOnce sync.Once
	file_api_cover_recommendation_azurecsp_proto_rawDescData = file_api_cover_recommendation_azurecsp_proto_rawDesc
)

func file_api_cover_recommendation_azurecsp_proto_rawDescGZIP() []byte {
	file_api_cover_recommendation_azurecsp_proto_rawDescOnce.Do(func() {
		file_api_cover_recommendation_azurecsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_recommendation_azurecsp_proto_rawDescData)
	})
	return file_api_cover_recommendation_azurecsp_proto_rawDescData
}

var file_api_cover_recommendation_azurecsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_cover_recommendation_azurecsp_proto_goTypes = []any{
	(*AzureCSPRecommendations)(nil), // 0: blueapi.api.cover.recommendation.azurecsp.AzureCSPRecommendations
}
var file_api_cover_recommendation_azurecsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_cover_recommendation_azurecsp_proto_init() }
func file_api_cover_recommendation_azurecsp_proto_init() {
	if File_api_cover_recommendation_azurecsp_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cover_recommendation_azurecsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_recommendation_azurecsp_proto_goTypes,
		DependencyIndexes: file_api_cover_recommendation_azurecsp_proto_depIdxs,
		MessageInfos:      file_api_cover_recommendation_azurecsp_proto_msgTypes,
	}.Build()
	File_api_cover_recommendation_azurecsp_proto = out.File
	file_api_cover_recommendation_azurecsp_proto_rawDesc = nil
	file_api_cover_recommendation_azurecsp_proto_goTypes = nil
	file_api_cover_recommendation_azurecsp_proto_depIdxs = nil
}
