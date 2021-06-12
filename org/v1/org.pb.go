// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: org/v1/org.proto

package org

import (
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

// Org resource definition.
type Org struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email       string            `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Metadata    map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Org) Reset() {
	*x = Org{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_v1_org_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Org) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Org) ProtoMessage() {}

func (x *Org) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Org.ProtoReflect.Descriptor instead.
func (*Org) Descriptor() ([]byte, []int) {
	return file_org_v1_org_proto_rawDescGZIP(), []int{0}
}

func (x *Org) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Org) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Org) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Org) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type CreateOrgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateOrgRequest) Reset() {
	*x = CreateOrgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_v1_org_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrgRequest) ProtoMessage() {}

func (x *CreateOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrgRequest.ProtoReflect.Descriptor instead.
func (*CreateOrgRequest) Descriptor() ([]byte, []int) {
	return file_org_v1_org_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrgRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateOrgRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type SendVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SendVerificationRequest) Reset() {
	*x = SendVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_v1_org_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationRequest) ProtoMessage() {}

func (x *SendVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerificationRequest.ProtoReflect.Descriptor instead.
func (*SendVerificationRequest) Descriptor() ([]byte, []int) {
	return file_org_v1_org_proto_rawDescGZIP(), []int{2}
}

func (x *SendVerificationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type VerifyOrgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *VerifyOrgRequest) Reset() {
	*x = VerifyOrgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_v1_org_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOrgRequest) ProtoMessage() {}

func (x *VerifyOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOrgRequest.ProtoReflect.Descriptor instead.
func (*VerifyOrgRequest) Descriptor() ([]byte, []int) {
	return file_org_v1_org_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyOrgRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VerifyOrgRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Request message for OrgApi.GetOrg rpc.
type GetOrgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrgRequest) Reset() {
	*x = GetOrgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_v1_org_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrgRequest) ProtoMessage() {}

func (x *GetOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrgRequest.ProtoReflect.Descriptor instead.
func (*GetOrgRequest) Descriptor() ([]byte, []int) {
	return file_org_v1_org_proto_rawDescGZIP(), []int{4}
}

var File_org_v1_org_proto protoreflect.FileDescriptor

var file_org_v1_org_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xcd, 0x01, 0x0a, 0x03, 0x4f, 0x72, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x4a, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x17,
	0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x10, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xab, 0x03, 0x0a, 0x06, 0x4f, 0x72, 0x67, 0x41, 0x70,
	0x69, 0x12, 0x5b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x12, 0x20,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x72, 0x67, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f,
	0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x81,
	0x01, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x62,
	0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x67, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x24, 0x2f, 0x6f, 0x72, 0x67, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x73,
	0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x67, 0x12,
	0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d,
	0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x73, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x3a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x67, 0x3a, 0x01, 0x2a,
	0x12, 0x52, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x12, 0x1d, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x67, 0x73, 0x42, 0xf4, 0x02, 0x0a, 0x15, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x42, 0x08,
	0x4f, 0x72, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x67, 0x92, 0x41, 0xab, 0x02,
	0x12, 0xcb, 0x01, 0x0a, 0x19, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x20, 0x42, 0x6c, 0x75,
	0x65, 0x3a, 0x20, 0x4f, 0x72, 0x67, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x50,
	0x0a, 0x19, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x20, 0x42, 0x6c, 0x75, 0x65, 0x3a, 0x20,
	0x4f, 0x72, 0x67, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x74, 0x6d, 0x6c,
	0x2a, 0x57, 0x0a, 0x1b, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x3a, 0x20, 0x41, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x32, 0x2e, 0x30, 0x12,
	0x38, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x72, 0x5b,
	0x0a, 0x24, 0x4d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0x75, 0x73, 0x20, 0x42, 0x6c, 0x75, 0x65, 0x3a, 0x20, 0x4f, 0x72, 0x67, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_org_v1_org_proto_rawDescOnce sync.Once
	file_org_v1_org_proto_rawDescData = file_org_v1_org_proto_rawDesc
)

func file_org_v1_org_proto_rawDescGZIP() []byte {
	file_org_v1_org_proto_rawDescOnce.Do(func() {
		file_org_v1_org_proto_rawDescData = protoimpl.X.CompressGZIP(file_org_v1_org_proto_rawDescData)
	})
	return file_org_v1_org_proto_rawDescData
}

var file_org_v1_org_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_org_v1_org_proto_goTypes = []interface{}{
	(*Org)(nil),                     // 0: blueapi.org.v1.Org
	(*CreateOrgRequest)(nil),        // 1: blueapi.org.v1.CreateOrgRequest
	(*SendVerificationRequest)(nil), // 2: blueapi.org.v1.SendVerificationRequest
	(*VerifyOrgRequest)(nil),        // 3: blueapi.org.v1.VerifyOrgRequest
	(*GetOrgRequest)(nil),           // 4: blueapi.org.v1.GetOrgRequest
	nil,                             // 5: blueapi.org.v1.Org.MetadataEntry
}
var file_org_v1_org_proto_depIdxs = []int32{
	5, // 0: blueapi.org.v1.Org.metadata:type_name -> blueapi.org.v1.Org.MetadataEntry
	1, // 1: blueapi.org.v1.OrgApi.CreateOrg:input_type -> blueapi.org.v1.CreateOrgRequest
	2, // 2: blueapi.org.v1.OrgApi.SendVerification:input_type -> blueapi.org.v1.SendVerificationRequest
	3, // 3: blueapi.org.v1.OrgApi.VerifyOrg:input_type -> blueapi.org.v1.VerifyOrgRequest
	4, // 4: blueapi.org.v1.OrgApi.GetOrg:input_type -> blueapi.org.v1.GetOrgRequest
	0, // 5: blueapi.org.v1.OrgApi.CreateOrg:output_type -> blueapi.org.v1.Org
	0, // 6: blueapi.org.v1.OrgApi.SendVerification:output_type -> blueapi.org.v1.Org
	0, // 7: blueapi.org.v1.OrgApi.VerifyOrg:output_type -> blueapi.org.v1.Org
	0, // 8: blueapi.org.v1.OrgApi.GetOrg:output_type -> blueapi.org.v1.Org
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_org_v1_org_proto_init() }
func file_org_v1_org_proto_init() {
	if File_org_v1_org_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_org_v1_org_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Org); i {
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
		file_org_v1_org_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrgRequest); i {
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
		file_org_v1_org_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerificationRequest); i {
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
		file_org_v1_org_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOrgRequest); i {
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
		file_org_v1_org_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrgRequest); i {
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
			RawDescriptor: file_org_v1_org_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_org_v1_org_proto_goTypes,
		DependencyIndexes: file_org_v1_org_proto_depIdxs,
		MessageInfos:      file_org_v1_org_proto_msgTypes,
	}.Build()
	File_org_v1_org_proto = out.File
	file_org_v1_org_proto_rawDesc = nil
	file_org_v1_org_proto_goTypes = nil
	file_org_v1_org_proto_depIdxs = nil
}
