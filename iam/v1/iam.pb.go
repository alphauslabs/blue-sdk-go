// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: iam/v1/iam.proto

package iam

import (
	api "github.com/alphauslabs/blue-sdk-go/api"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for Iam.WhoAmI rpc.
type WhoAmIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WhoAmIRequest) Reset() {
	*x = WhoAmIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIRequest) ProtoMessage() {}

func (x *WhoAmIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIRequest.ProtoReflect.Descriptor instead.
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{0}
}

// Request message for Iam.ListUsers rpc.
type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{1}
}

// Response message for Iam.ListUsers rpc.
type ListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*api.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{2}
}

func (x *ListUsersResponse) GetUsers() []*api.User {
	if x != nil {
		return x.Users
	}
	return nil
}

// Request message for Iam.GetUser rpc.
type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type IpFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ip filter id. Note that this id is transcient and is not fixed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Can be 'global', 'rootuser', or 'subuser'.
	Scope string `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	// The target of the filter. If global scope, this could be the organization
	// id or billing group id. Rootuser id for rootuser scope, and subuser id
	// for subuser scope. For subuser targets, format is 'rootuserid/subuserid'.
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	// The type of ip filter. It could be 'whitelist' or 'blacklist'.
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// The ip filter value. Should be in CIDR format.
	Value string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IpFilter) Reset() {
	*x = IpFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpFilter) ProtoMessage() {}

func (x *IpFilter) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpFilter.ProtoReflect.Descriptor instead.
func (*IpFilter) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{4}
}

func (x *IpFilter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IpFilter) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *IpFilter) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *IpFilter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IpFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Request message for Iam.ListIpFilters rpc.
type ListIpFiltersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListIpFiltersRequest) Reset() {
	*x = ListIpFiltersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIpFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIpFiltersRequest) ProtoMessage() {}

func (x *ListIpFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIpFiltersRequest.ProtoReflect.Descriptor instead.
func (*ListIpFiltersRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{5}
}

// Response message for Iam.ListIpFilters rpc.
type ListIpFiltersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*IpFilter `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListIpFiltersResponse) Reset() {
	*x = ListIpFiltersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIpFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIpFiltersResponse) ProtoMessage() {}

func (x *ListIpFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIpFiltersResponse.ProtoReflect.Descriptor instead.
func (*ListIpFiltersResponse) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{6}
}

func (x *ListIpFiltersResponse) GetItems() []*IpFilter {
	if x != nil {
		return x.Items
	}
	return nil
}

// Request message for Iam.CreateIpFilter rpc.
type CreateIpFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The IP input to filter, either blacklist or whitelist. Should be in
	// CIDR format, i.e. 1.2.3.4/32
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// Optional. Can be 'whitelist' or 'blacklist'. Defaults to 'blacklist' if empty.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Optional. If valid rootuser and sub_user is empty, filter is for this rootuser.
	// If both root_user and sub_user inputs are empty, filter is at global scope.
	RootUser string `protobuf:"bytes,3,opt,name=root_user,json=rootUser,proto3" json:"root_user,omitempty"`
	// Optional. If valid subuser, filter is for this subuser; root_user value is discarded.
	// If both root_user and sub_user inputs are empty, filter is at global scope.
	SubUser string `protobuf:"bytes,4,opt,name=sub_user,json=subUser,proto3" json:"sub_user,omitempty"`
}

func (x *CreateIpFilterRequest) Reset() {
	*x = CreateIpFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIpFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIpFilterRequest) ProtoMessage() {}

func (x *CreateIpFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIpFilterRequest.ProtoReflect.Descriptor instead.
func (*CreateIpFilterRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{7}
}

func (x *CreateIpFilterRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *CreateIpFilterRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateIpFilterRequest) GetRootUser() string {
	if x != nil {
		return x.RootUser
	}
	return ""
}

func (x *CreateIpFilterRequest) GetSubUser() string {
	if x != nil {
		return x.SubUser
	}
	return ""
}

// Request message for Iam.DeleteIpFilter rpc.
type DeleteIpFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id to delete.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteIpFilterRequest) Reset() {
	*x = DeleteIpFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIpFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIpFilterRequest) ProtoMessage() {}

func (x *DeleteIpFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIpFilterRequest.ProtoReflect.Descriptor instead.
func (*DeleteIpFilterRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteIpFilterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_iam_v1_iam_proto protoreflect.FileDescriptor

var file_iam_v1_iam_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x1a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a,
	0x0d, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x3c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x72, 0x0a, 0x08, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x47, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x70, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x79, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x75, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x55, 0x73, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0xf7, 0x05, 0x0a, 0x03, 0x49, 0x61, 0x6d, 0x12, 0x52, 0x0a, 0x06, 0x57, 0x68, 0x6f, 0x41, 0x6d,
	0x49, 0x12, 0x1d, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x12, 0x67, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x5a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1e, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x12, 0x77, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x24, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x70, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x70, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x6f, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x70,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x62,
	0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x70, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x1a, 0x7c, 0x92, 0x41, 0x79,
	0x12, 0x2a, 0x42, 0x61, 0x73, 0x65, 0x20, 0x55, 0x52, 0x4c, 0x3a, 0x20, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x1a, 0x4b, 0x0a, 0x12,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x35, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f,
	0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x42, 0x45, 0x0a, 0x15, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x61, 0x6d, 0x42, 0x08, 0x49, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x22, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iam_v1_iam_proto_rawDescOnce sync.Once
	file_iam_v1_iam_proto_rawDescData = file_iam_v1_iam_proto_rawDesc
)

func file_iam_v1_iam_proto_rawDescGZIP() []byte {
	file_iam_v1_iam_proto_rawDescOnce.Do(func() {
		file_iam_v1_iam_proto_rawDescData = protoimpl.X.CompressGZIP(file_iam_v1_iam_proto_rawDescData)
	})
	return file_iam_v1_iam_proto_rawDescData
}

var file_iam_v1_iam_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_iam_v1_iam_proto_goTypes = []interface{}{
	(*WhoAmIRequest)(nil),         // 0: blueapi.iam.v1.WhoAmIRequest
	(*ListUsersRequest)(nil),      // 1: blueapi.iam.v1.ListUsersRequest
	(*ListUsersResponse)(nil),     // 2: blueapi.iam.v1.ListUsersResponse
	(*GetUserRequest)(nil),        // 3: blueapi.iam.v1.GetUserRequest
	(*IpFilter)(nil),              // 4: blueapi.iam.v1.IpFilter
	(*ListIpFiltersRequest)(nil),  // 5: blueapi.iam.v1.ListIpFiltersRequest
	(*ListIpFiltersResponse)(nil), // 6: blueapi.iam.v1.ListIpFiltersResponse
	(*CreateIpFilterRequest)(nil), // 7: blueapi.iam.v1.CreateIpFilterRequest
	(*DeleteIpFilterRequest)(nil), // 8: blueapi.iam.v1.DeleteIpFilterRequest
	(*api.User)(nil),              // 9: blueapi.api.User
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_iam_v1_iam_proto_depIdxs = []int32{
	9,  // 0: blueapi.iam.v1.ListUsersResponse.users:type_name -> blueapi.api.User
	4,  // 1: blueapi.iam.v1.ListIpFiltersResponse.items:type_name -> blueapi.iam.v1.IpFilter
	0,  // 2: blueapi.iam.v1.Iam.WhoAmI:input_type -> blueapi.iam.v1.WhoAmIRequest
	1,  // 3: blueapi.iam.v1.Iam.ListUsers:input_type -> blueapi.iam.v1.ListUsersRequest
	3,  // 4: blueapi.iam.v1.Iam.GetUser:input_type -> blueapi.iam.v1.GetUserRequest
	5,  // 5: blueapi.iam.v1.Iam.ListIpFilters:input_type -> blueapi.iam.v1.ListIpFiltersRequest
	7,  // 6: blueapi.iam.v1.Iam.CreateIpFilter:input_type -> blueapi.iam.v1.CreateIpFilterRequest
	8,  // 7: blueapi.iam.v1.Iam.DeleteIpFilter:input_type -> blueapi.iam.v1.DeleteIpFilterRequest
	9,  // 8: blueapi.iam.v1.Iam.WhoAmI:output_type -> blueapi.api.User
	2,  // 9: blueapi.iam.v1.Iam.ListUsers:output_type -> blueapi.iam.v1.ListUsersResponse
	9,  // 10: blueapi.iam.v1.Iam.GetUser:output_type -> blueapi.api.User
	6,  // 11: blueapi.iam.v1.Iam.ListIpFilters:output_type -> blueapi.iam.v1.ListIpFiltersResponse
	4,  // 12: blueapi.iam.v1.Iam.CreateIpFilter:output_type -> blueapi.iam.v1.IpFilter
	10, // 13: blueapi.iam.v1.Iam.DeleteIpFilter:output_type -> google.protobuf.Empty
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_iam_v1_iam_proto_init() }
func file_iam_v1_iam_proto_init() {
	if File_iam_v1_iam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iam_v1_iam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmIRequest); i {
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
		file_iam_v1_iam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersRequest); i {
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
		file_iam_v1_iam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersResponse); i {
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
		file_iam_v1_iam_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_iam_v1_iam_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpFilter); i {
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
		file_iam_v1_iam_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIpFiltersRequest); i {
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
		file_iam_v1_iam_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIpFiltersResponse); i {
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
		file_iam_v1_iam_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIpFilterRequest); i {
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
		file_iam_v1_iam_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIpFilterRequest); i {
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
			RawDescriptor: file_iam_v1_iam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iam_v1_iam_proto_goTypes,
		DependencyIndexes: file_iam_v1_iam_proto_depIdxs,
		MessageInfos:      file_iam_v1_iam_proto_msgTypes,
	}.Build()
	File_iam_v1_iam_proto = out.File
	file_iam_v1_iam_proto_rawDesc = nil
	file_iam_v1_iam_proto_goTypes = nil
	file_iam_v1_iam_proto_depIdxs = nil
}
