// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: api/cover/applyrecommendation.proto

package cover

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

type EC2 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceType  string                 `protobuf:"bytes,1,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EC2) Reset() {
	*x = EC2{}
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EC2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EC2) ProtoMessage() {}

func (x *EC2) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EC2.ProtoReflect.Descriptor instead.
func (*EC2) Descriptor() ([]byte, []int) {
	return file_api_cover_applyrecommendation_proto_rawDescGZIP(), []int{0}
}

func (x *EC2) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

type Lambda struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Memory        int64                  `protobuf:"varint,1,opt,name=memory,proto3" json:"memory,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Lambda) Reset() {
	*x = Lambda{}
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Lambda) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lambda) ProtoMessage() {}

func (x *Lambda) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lambda.ProtoReflect.Descriptor instead.
func (*Lambda) Descriptor() ([]byte, []int) {
	return file_api_cover_applyrecommendation_proto_rawDescGZIP(), []int{1}
}

func (x *Lambda) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

type EBS struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VolumeType    string                 `protobuf:"bytes,1,opt,name=volumeType,proto3" json:"volumeType,omitempty"`
	Iops          int64                  `protobuf:"varint,2,opt,name=iops,proto3" json:"iops,omitempty"`
	SizeInGb      int64                  `protobuf:"varint,3,opt,name=sizeInGb,proto3" json:"sizeInGb,omitempty"`
	Throughput    int64                  `protobuf:"varint,4,opt,name=throughput,proto3" json:"throughput,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EBS) Reset() {
	*x = EBS{}
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EBS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EBS) ProtoMessage() {}

func (x *EBS) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EBS.ProtoReflect.Descriptor instead.
func (*EBS) Descriptor() ([]byte, []int) {
	return file_api_cover_applyrecommendation_proto_rawDescGZIP(), []int{2}
}

func (x *EBS) GetVolumeType() string {
	if x != nil {
		return x.VolumeType
	}
	return ""
}

func (x *EBS) GetIops() int64 {
	if x != nil {
		return x.Iops
	}
	return 0
}

func (x *EBS) GetSizeInGb() int64 {
	if x != nil {
		return x.SizeInGb
	}
	return 0
}

func (x *EBS) GetThroughput() int64 {
	if x != nil {
		return x.Throughput
	}
	return 0
}

type AWS struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EC2           *EC2                   `protobuf:"bytes,1,opt,name=EC2,proto3" json:"EC2,omitempty"`
	Lambda        *Lambda                `protobuf:"bytes,2,opt,name=Lambda,proto3" json:"Lambda,omitempty"`
	EBS           *EBS                   `protobuf:"bytes,3,opt,name=EBS,proto3" json:"EBS,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AWS) Reset() {
	*x = AWS{}
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AWS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWS) ProtoMessage() {}

func (x *AWS) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWS.ProtoReflect.Descriptor instead.
func (*AWS) Descriptor() ([]byte, []int) {
	return file_api_cover_applyrecommendation_proto_rawDescGZIP(), []int{3}
}

func (x *AWS) GetEC2() *EC2 {
	if x != nil {
		return x.EC2
	}
	return nil
}

func (x *AWS) GetLambda() *Lambda {
	if x != nil {
		return x.Lambda
	}
	return nil
}

func (x *AWS) GetEBS() *EBS {
	if x != nil {
		return x.EBS
	}
	return nil
}

type GCP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GCP) Reset() {
	*x = GCP{}
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GCP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCP) ProtoMessage() {}

func (x *GCP) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCP.ProtoReflect.Descriptor instead.
func (*GCP) Descriptor() ([]byte, []int) {
	return file_api_cover_applyrecommendation_proto_rawDescGZIP(), []int{4}
}

type AzureCSP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AzureCSP) Reset() {
	*x = AzureCSP{}
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureCSP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureCSP) ProtoMessage() {}

func (x *AzureCSP) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_applyrecommendation_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureCSP.ProtoReflect.Descriptor instead.
func (*AzureCSP) Descriptor() ([]byte, []int) {
	return file_api_cover_applyrecommendation_proto_rawDescGZIP(), []int{5}
}

var File_api_cover_applyrecommendation_proto protoreflect.FileDescriptor

var file_api_cover_applyrecommendation_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x03, 0x45, 0x43, 0x32, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x75, 0x0a, 0x03, 0x45, 0x42, 0x53, 0x12, 0x1e, 0x0a, 0x0a,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x6f, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x69, 0x6f, 0x70, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x47, 0x62, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x47, 0x62, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x70, 0x75, 0x74, 0x22, 0x8c, 0x01, 0x0a,
	0x03, 0x41, 0x57, 0x53, 0x12, 0x28, 0x0a, 0x03, 0x45, 0x43, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x43, 0x32, 0x52, 0x03, 0x45, 0x43, 0x32, 0x12, 0x31,
	0x0a, 0x06, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x52, 0x06, 0x4c, 0x61, 0x6d, 0x62, 0x64,
	0x61, 0x12, 0x28, 0x0a, 0x03, 0x45, 0x42, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x2e, 0x45, 0x42, 0x53, 0x52, 0x03, 0x45, 0x42, 0x53, 0x22, 0x05, 0x0a, 0x03, 0x47,
	0x43, 0x50, 0x22, 0x0a, 0x0a, 0x08, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x53, 0x50, 0x42, 0x71,
	0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x42, 0x20, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65,
	0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cover_applyrecommendation_proto_rawDescOnce sync.Once
	file_api_cover_applyrecommendation_proto_rawDescData = file_api_cover_applyrecommendation_proto_rawDesc
)

func file_api_cover_applyrecommendation_proto_rawDescGZIP() []byte {
	file_api_cover_applyrecommendation_proto_rawDescOnce.Do(func() {
		file_api_cover_applyrecommendation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_applyrecommendation_proto_rawDescData)
	})
	return file_api_cover_applyrecommendation_proto_rawDescData
}

var file_api_cover_applyrecommendation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_cover_applyrecommendation_proto_goTypes = []any{
	(*EC2)(nil),      // 0: blueapi.api.cover.EC2
	(*Lambda)(nil),   // 1: blueapi.api.cover.Lambda
	(*EBS)(nil),      // 2: blueapi.api.cover.EBS
	(*AWS)(nil),      // 3: blueapi.api.cover.AWS
	(*GCP)(nil),      // 4: blueapi.api.cover.GCP
	(*AzureCSP)(nil), // 5: blueapi.api.cover.AzureCSP
}
var file_api_cover_applyrecommendation_proto_depIdxs = []int32{
	0, // 0: blueapi.api.cover.AWS.EC2:type_name -> blueapi.api.cover.EC2
	1, // 1: blueapi.api.cover.AWS.Lambda:type_name -> blueapi.api.cover.Lambda
	2, // 2: blueapi.api.cover.AWS.EBS:type_name -> blueapi.api.cover.EBS
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_cover_applyrecommendation_proto_init() }
func file_api_cover_applyrecommendation_proto_init() {
	if File_api_cover_applyrecommendation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cover_applyrecommendation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_applyrecommendation_proto_goTypes,
		DependencyIndexes: file_api_cover_applyrecommendation_proto_depIdxs,
		MessageInfos:      file_api_cover_applyrecommendation_proto_msgTypes,
	}.Build()
	File_api_cover_applyrecommendation_proto = out.File
	file_api_cover_applyrecommendation_proto_rawDesc = nil
	file_api_cover_applyrecommendation_proto_goTypes = nil
	file_api_cover_applyrecommendation_proto_depIdxs = nil
}
