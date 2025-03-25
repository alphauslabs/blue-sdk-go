// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/cover/recommendation/azurecsp.proto

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

const file_api_cover_recommendation_azurecsp_proto_rawDesc = "" +
	"\n" +
	"'api/cover/recommendation/azurecsp.proto\x12)blueapi.api.cover.recommendation.azurecsp\"\x19\n" +
	"\x17AzureCSPRecommendationsB\x92\x01\n" +
	".cloud.alphaus.blueapi.api.cover.recommendationB#ApiCoverAzureCspRecommendationProtoZ;github.com/alphauslabs/blue-sdk-go/api/cover/recommendationb\x06proto3"

var (
	file_api_cover_recommendation_azurecsp_proto_rawDescOnce sync.Once
	file_api_cover_recommendation_azurecsp_proto_rawDescData []byte
)

func file_api_cover_recommendation_azurecsp_proto_rawDescGZIP() []byte {
	file_api_cover_recommendation_azurecsp_proto_rawDescOnce.Do(func() {
		file_api_cover_recommendation_azurecsp_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_cover_recommendation_azurecsp_proto_rawDesc), len(file_api_cover_recommendation_azurecsp_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_cover_recommendation_azurecsp_proto_rawDesc), len(file_api_cover_recommendation_azurecsp_proto_rawDesc)),
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
	file_api_cover_recommendation_azurecsp_proto_goTypes = nil
	file_api_cover_recommendation_azurecsp_proto_depIdxs = nil
}
