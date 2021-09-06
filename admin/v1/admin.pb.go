// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: admin/v1/admin.proto

package admin

import (
	api "github.com/alphauslabs/blue-sdk-go/api"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for the Admin.ListAccountGroups rpc.
type ListAccountGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAccountGroupsRequest) Reset() {
	*x = ListAccountGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountGroupsRequest) ProtoMessage() {}

func (x *ListAccountGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListAccountGroupsRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{0}
}

// Response message for the Admin.ListAccountGroups rpc.
type ListAccountGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountGroups []*api.AccountGroup `protobuf:"bytes,1,rep,name=accountGroups,proto3" json:"accountGroups,omitempty"`
}

func (x *ListAccountGroupsResponse) Reset() {
	*x = ListAccountGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountGroupsResponse) ProtoMessage() {}

func (x *ListAccountGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountGroupsResponse.ProtoReflect.Descriptor instead.
func (*ListAccountGroupsResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{1}
}

func (x *ListAccountGroupsResponse) GetAccountGroups() []*api.AccountGroup {
	if x != nil {
		return x.AccountGroups
	}
	return nil
}

// Request message for the Admin.GetAccountGroup rpc.
type GetAccountGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAccountGroupRequest) Reset() {
	*x = GetAccountGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountGroupRequest) ProtoMessage() {}

func (x *GetAccountGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountGroupRequest.ProtoReflect.Descriptor instead.
func (*GetAccountGroupRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Response message for the Admin.GetAccountGroup rpc.
type GetAccountGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AcctGroup *api.AccountGroup `protobuf:"bytes,1,opt,name=acctGroup,proto3" json:"acctGroup,omitempty"`
}

func (x *GetAccountGroupResponse) Reset() {
	*x = GetAccountGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountGroupResponse) ProtoMessage() {}

func (x *GetAccountGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountGroupResponse.ProtoReflect.Descriptor instead.
func (*GetAccountGroupResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_admin_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountGroupResponse) GetAcctGroup() *api.AccountGroup {
	if x != nil {
		return x.AcctGroup
	}
	return nil
}

var File_admin_v1_admin_proto protoreflect.FileDescriptor

var file_admin_v1_admin_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a,
	0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5c, 0x0a, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x52, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x09, 0x61, 0x63, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x32, 0xb6, 0x03, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x8c, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2a, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x30, 0x01, 0x12,
	0x89, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x28, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x12, 0x19, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x74,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x1a, 0x91, 0x01, 0x92, 0x41,
	0x8d, 0x01, 0x12, 0x3c, 0x28, 0x42, 0x45, 0x54, 0x41, 0x29, 0x20, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x20, 0x41, 0x50, 0x49, 0x2e, 0x20, 0x42, 0x61, 0x73, 0x65, 0x20, 0x55, 0x52, 0x4c, 0x3a, 0x20,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x75, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x65,
	0x1a, 0x4d, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x42,
	0x4b, 0x0a, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x0a, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62,
	0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_v1_admin_proto_rawDescOnce sync.Once
	file_admin_v1_admin_proto_rawDescData = file_admin_v1_admin_proto_rawDesc
)

func file_admin_v1_admin_proto_rawDescGZIP() []byte {
	file_admin_v1_admin_proto_rawDescOnce.Do(func() {
		file_admin_v1_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_v1_admin_proto_rawDescData)
	})
	return file_admin_v1_admin_proto_rawDescData
}

var file_admin_v1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_admin_v1_admin_proto_goTypes = []interface{}{
	(*ListAccountGroupsRequest)(nil),  // 0: blueapi.admin.v1.ListAccountGroupsRequest
	(*ListAccountGroupsResponse)(nil), // 1: blueapi.admin.v1.ListAccountGroupsResponse
	(*GetAccountGroupRequest)(nil),    // 2: blueapi.admin.v1.GetAccountGroupRequest
	(*GetAccountGroupResponse)(nil),   // 3: blueapi.admin.v1.GetAccountGroupResponse
	(*api.AccountGroup)(nil),          // 4: blueapi.api.AccountGroup
}
var file_admin_v1_admin_proto_depIdxs = []int32{
	4, // 0: blueapi.admin.v1.ListAccountGroupsResponse.accountGroups:type_name -> blueapi.api.AccountGroup
	4, // 1: blueapi.admin.v1.GetAccountGroupResponse.acctGroup:type_name -> blueapi.api.AccountGroup
	0, // 2: blueapi.admin.v1.Admin.ListAccountGroups:input_type -> blueapi.admin.v1.ListAccountGroupsRequest
	2, // 3: blueapi.admin.v1.Admin.GetAccountGroup:input_type -> blueapi.admin.v1.GetAccountGroupRequest
	1, // 4: blueapi.admin.v1.Admin.ListAccountGroups:output_type -> blueapi.admin.v1.ListAccountGroupsResponse
	3, // 5: blueapi.admin.v1.Admin.GetAccountGroup:output_type -> blueapi.admin.v1.GetAccountGroupResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_admin_v1_admin_proto_init() }
func file_admin_v1_admin_proto_init() {
	if File_admin_v1_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_v1_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountGroupsRequest); i {
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
		file_admin_v1_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountGroupsResponse); i {
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
		file_admin_v1_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountGroupRequest); i {
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
		file_admin_v1_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountGroupResponse); i {
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
			RawDescriptor: file_admin_v1_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_v1_admin_proto_goTypes,
		DependencyIndexes: file_admin_v1_admin_proto_depIdxs,
		MessageInfos:      file_admin_v1_admin_proto_msgTypes,
	}.Build()
	File_admin_v1_admin_proto = out.File
	file_admin_v1_admin_proto_rawDesc = nil
	file_admin_v1_admin_proto_goTypes = nil
	file_admin_v1_admin_proto_depIdxs = nil
}
