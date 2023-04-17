// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: api/cover/fee.proto

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

type Accounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId  string `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Allocation string `protobuf:"bytes,2,opt,name=allocation,proto3" json:"allocation,omitempty"`
	Duration   string `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Accounts) Reset() {
	*x = Accounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_fee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accounts) ProtoMessage() {}

func (x *Accounts) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_fee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accounts.ProtoReflect.Descriptor instead.
func (*Accounts) Descriptor() ([]byte, []int) {
	return file_api_cover_fee_proto_rawDescGZIP(), []int{0}
}

func (x *Accounts) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Accounts) GetAllocation() string {
	if x != nil {
		return x.Allocation
	}
	return ""
}

func (x *Accounts) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

type CostGroups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId  string `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Allocation string `protobuf:"bytes,2,opt,name=allocation,proto3" json:"allocation,omitempty"`
	Duration   string `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *CostGroups) Reset() {
	*x = CostGroups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_fee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostGroups) ProtoMessage() {}

func (x *CostGroups) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_fee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostGroups.ProtoReflect.Descriptor instead.
func (*CostGroups) Descriptor() ([]byte, []int) {
	return file_api_cover_fee_proto_rawDescGZIP(), []int{1}
}

func (x *CostGroups) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CostGroups) GetAllocation() string {
	if x != nil {
		return x.Allocation
	}
	return ""
}

func (x *CostGroups) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

var File_api_cover_fee_proto protoreflect.FileDescriptor

var file_api_cover_fee_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x22, 0x64, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x66,
	0x0a, 0x0a, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x10, 0x41, 0x70, 0x69, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x46, 0x65, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_cover_fee_proto_rawDescOnce sync.Once
	file_api_cover_fee_proto_rawDescData = file_api_cover_fee_proto_rawDesc
)

func file_api_cover_fee_proto_rawDescGZIP() []byte {
	file_api_cover_fee_proto_rawDescOnce.Do(func() {
		file_api_cover_fee_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_fee_proto_rawDescData)
	})
	return file_api_cover_fee_proto_rawDescData
}

var file_api_cover_fee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_cover_fee_proto_goTypes = []interface{}{
	(*Accounts)(nil),   // 0: blueapi.api.cover.Accounts
	(*CostGroups)(nil), // 1: blueapi.api.cover.CostGroups
}
var file_api_cover_fee_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_cover_fee_proto_init() }
func file_api_cover_fee_proto_init() {
	if File_api_cover_fee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_cover_fee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accounts); i {
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
		file_api_cover_fee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostGroups); i {
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
			RawDescriptor: file_api_cover_fee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_fee_proto_goTypes,
		DependencyIndexes: file_api_cover_fee_proto_depIdxs,
		MessageInfos:      file_api_cover_fee_proto_msgTypes,
	}.Build()
	File_api_cover_fee_proto = out.File
	file_api_cover_fee_proto_rawDesc = nil
	file_api_cover_fee_proto_goTypes = nil
	file_api_cover_fee_proto_depIdxs = nil
}