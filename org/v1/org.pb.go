// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: org/v1/org.proto

package org

import (
	ripple "github.com/alphauslabs/blue-sdk-go/api/ripple"
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

// Request message for the Organization.CreateOrg rpc.
type CreateOrgRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Organization's email
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// Required. Organization's description or company name
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. User can decide to let the system generates the password.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// To be used for onboarding profile
	Name          string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Role          string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	Usage         string `protobuf:"bytes,6,opt,name=usage,proto3" json:"usage,omitempty"`
	Intention     string `protobuf:"bytes,7,opt,name=intention,proto3" json:"intention,omitempty"`
	IsAuth0Signup bool   `protobuf:"varint,8,opt,name=isAuth0Signup,proto3" json:"isAuth0Signup,omitempty"`
	// Optional. Customer Id from AWS marketplace
	CustomerId string `protobuf:"bytes,9,opt,name=customerId,proto3" json:"customerId,omitempty"`
	// Optional. User locale. Default value is en if not supplied.
	Locale string `protobuf:"bytes,10,opt,name=locale,proto3" json:"locale,omitempty"`
	// Optional. Auth0 user id for users registering using social login.
	Auth0UserId   string `protobuf:"bytes,11,opt,name=auth0UserId,proto3" json:"auth0UserId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrgRequest) Reset() {
	*x = CreateOrgRequest{}
	mi := &file_org_v1_org_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrgRequest) ProtoMessage() {}

func (x *CreateOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[0]
	if x != nil {
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
	return file_org_v1_org_proto_rawDescGZIP(), []int{0}
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

func (x *CreateOrgRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateOrgRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateOrgRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *CreateOrgRequest) GetUsage() string {
	if x != nil {
		return x.Usage
	}
	return ""
}

func (x *CreateOrgRequest) GetIntention() string {
	if x != nil {
		return x.Intention
	}
	return ""
}

func (x *CreateOrgRequest) GetIsAuth0Signup() bool {
	if x != nil {
		return x.IsAuth0Signup
	}
	return false
}

func (x *CreateOrgRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateOrgRequest) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *CreateOrgRequest) GetAuth0UserId() string {
	if x != nil {
		return x.Auth0UserId
	}
	return ""
}

// Response message for the Organization.CreateOrg rpc.
type CreateOrgResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Org           *ripple.Org            `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrgResponse) Reset() {
	*x = CreateOrgResponse{}
	mi := &file_org_v1_org_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrgResponse) ProtoMessage() {}

func (x *CreateOrgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrgResponse.ProtoReflect.Descriptor instead.
func (*CreateOrgResponse) Descriptor() ([]byte, []int) {
	return file_org_v1_org_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrgResponse) GetOrg() *ripple.Org {
	if x != nil {
		return x.Org
	}
	return nil
}

func (x *CreateOrgResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Request message for the Organization.SendVerification rpc.
type SendVerificationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendVerificationRequest) Reset() {
	*x = SendVerificationRequest{}
	mi := &file_org_v1_org_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationRequest) ProtoMessage() {}

func (x *SendVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[2]
	if x != nil {
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

// Request message for the Organization.VerifyOrg rpc.
type VerifyOrgRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyOrgRequest) Reset() {
	*x = VerifyOrgRequest{}
	mi := &file_org_v1_org_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOrgRequest) ProtoMessage() {}

func (x *VerifyOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[3]
	if x != nil {
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

func (x *VerifyOrgRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Request message for the Organization.GetOrg rpc.
type GetOrgRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrgRequest) Reset() {
	*x = GetOrgRequest{}
	mi := &file_org_v1_org_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrgRequest) ProtoMessage() {}

func (x *GetOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[4]
	if x != nil {
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

// Request message for the Organization.UpdateMetadata rpc.
type UpdateMetadataRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateMetadataRequest) Reset() {
	*x = UpdateMetadataRequest{}
	mi := &file_org_v1_org_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMetadataRequest) ProtoMessage() {}

func (x *UpdateMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMetadataRequest.ProtoReflect.Descriptor instead.
func (*UpdateMetadataRequest) Descriptor() ([]byte, []int) {
	return file_org_v1_org_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateMetadataRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateMetadataRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Request message for the Organization.UpdatePassword rpc.
type UpdatePasswordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OldPassword   string                 `protobuf:"bytes,1,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword   string                 `protobuf:"bytes,2,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePasswordRequest) Reset() {
	*x = UpdatePasswordRequest{}
	mi := &file_org_v1_org_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordRequest) ProtoMessage() {}

func (x *UpdatePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return file_org_v1_org_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *UpdatePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

// Request message for the Organization.DeleteOrg rpc.
type DeleteOrgRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOrgRequest) Reset() {
	*x = DeleteOrgRequest{}
	mi := &file_org_v1_org_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrgRequest) ProtoMessage() {}

func (x *DeleteOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_v1_org_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrgRequest.ProtoReflect.Descriptor instead.
func (*DeleteOrgRequest) Descriptor() ([]byte, []int) {
	return file_org_v1_org_proto_rawDescGZIP(), []int{7}
}

var File_org_v1_org_proto protoreflect.FileDescriptor

var file_org_v1_org_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x02, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x73, 0x41, 0x75,
	0x74, 0x68, 0x30, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x30, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x30, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x74,
	0x68, 0x30, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x03, 0x6f, 0x72, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e,
	0x4f, 0x72, 0x67, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x24, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5b, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xe8, 0x06, 0x0a, 0x0c, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x64, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x12, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x12,
	0x75, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x18, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x76, 0x31, 0x3a, 0x73, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5d, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x4f, 0x72, 0x67, 0x12, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0e, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x3a, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x51, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x12,
	0x1d, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70,
	0x70, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x67, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12,
	0x07, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x12, 0x6c, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x3a, 0x01, 0x2a, 0x1a, 0x10, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x6a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a,
	0x01, 0x2a, 0x1a, 0x0e, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x64, 0x12, 0x56, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x12,
	0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x09, 0x2a, 0x07, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x1a, 0x96, 0x01, 0x92, 0x41, 0x92,
	0x01, 0x12, 0x43, 0x28, 0x42, 0x45, 0x54, 0x41, 0x29, 0x20, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x41, 0x50, 0x49, 0x2e, 0x20, 0x42, 0x61, 0x73, 0x65,
	0x20, 0x55, 0x52, 0x4c, 0x3a, 0x20, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x1a, 0x4b, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x42, 0xe5, 0x09, 0x92, 0x41, 0x9c, 0x09, 0x12, 0xfa, 0x08, 0x0a, 0x12, 0x42,
	0x6c, 0x75, 0x65, 0x20, 0x41, 0x50, 0x49, 0x20, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0xdf, 0x08, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x20, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x69, 0x74, 0x68,
	0x20, 0x69, 0x74, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x20, 0x42,
	0x6c, 0x75, 0x65, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x52, 0x45, 0x53,
	0x54, 0x66, 0x75, 0x6c, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x63, 0x61,
	0x6e, 0x20, 0x62, 0x65, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x20, 0x62, 0x79,
	0x20, 0x61, 0x6e, 0x20, 0x48, 0x54, 0x54, 0x50, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20,
	0x73, 0x75, 0x63, 0x68, 0x20, 0x61, 0x73, 0x20, 0x60, 0x63, 0x75, 0x72, 0x6c, 0x60, 0x2c, 0x20,
	0x6f, 0x72, 0x20, 0x61, 0x6e, 0x79, 0x20, 0x48, 0x54, 0x54, 0x50, 0x20, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x69, 0x73, 0x20, 0x70, 0x61, 0x72,
	0x74, 0x20, 0x6f, 0x66, 0x20, 0x6d, 0x6f, 0x73, 0x74, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x6e,
	0x20, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x20, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x20, 0x69, 0x73, 0x20, 0x61, 0x75,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x20, 0x66, 0x72, 0x6f, 0x6d,
	0x20, 0x5b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x20, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x29, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x5b, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5d,
	0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x29, 0x2c, 0x20, 0x74, 0x6f, 0x67, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x6f, 0x75, 0x72, 0x20, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x5b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x73, 0x2f,
	0x29, 0x2e, 0x20, 0x53, 0x65, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6f, 0x66, 0x66, 0x69, 0x63,
	0x69, 0x61, 0x6c, 0x20, 0x5b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69,
	0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x6f,
	0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6d, 0x6f,
	0x72, 0x65, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x0a,
	0x0a, 0x59, 0x6f, 0x75, 0x20, 0x6d, 0x61, 0x79, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x20, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x20, 0x6d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x20, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x3a, 0x0a, 0x2d, 0x20,
	0x2a, 0x2a, 0x28, 0x57, 0x4f, 0x52, 0x4b, 0x2d, 0x49, 0x4e, 0x2d, 0x50, 0x52, 0x4f, 0x47, 0x52,
	0x45, 0x53, 0x53, 0x29, 0x2a, 0x2a, 0x20, 0x2d, 0x20, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x20, 0x69, 0x73, 0x20, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x2c,
	0x20, 0x64, 0x6f, 0x6e, 0x27, 0x74, 0x20, 0x75, 0x73, 0x65, 0x20, 0x79, 0x65, 0x74, 0x3b, 0x0a,
	0x2d, 0x20, 0x2a, 0x2a, 0x28, 0x42, 0x45, 0x54, 0x41, 0x29, 0x2a, 0x2a, 0x20, 0x2d, 0x20, 0x4e,
	0x65, 0x77, 0x20, 0x6f, 0x72, 0x20, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x20, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2c, 0x20, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x3b,
	0x20, 0x61, 0x6e, 0x64, 0x0a, 0x2d, 0x20, 0x2a, 0x2a, 0x28, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43,
	0x41, 0x54, 0x45, 0x44, 0x29, 0x2a, 0x2a, 0x20, 0x2d, 0x20, 0x4f, 0x75, 0x74, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x20, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x0a, 0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x20, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2c, 0x20, 0x65, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x61, 0x6c, 0x6c, 0x79, 0x20, 0x74, 0x68, 0x6f, 0x73, 0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x20, 0x6f, 0x66, 0x20,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2c, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x73, 0x3b, 0x20, 0x6e, 0x65, 0x77, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x73, 0x65, 0x70,
	0x61, 0x72, 0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x20, 0x6f, 0x66,
	0x20, 0xe2, 0x80, 0x9c, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0xe2, 0x80, 0x9d, 0x2e, 0x20, 0x45,
	0x61, 0x63, 0x68, 0x20, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x20, 0x69, 0x73, 0x20, 0x61, 0x6e, 0x20,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x63, 0x61,
	0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x65, 0x69, 0x74, 0x68, 0x65, 0x72,
	0x20, 0x61, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x20, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x20, 0x6f, 0x72, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x20, 0x4f, 0x6e, 0x6c, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6c, 0x61, 0x73, 0x74, 0x20, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2c, 0x20, 0x69, 0x66, 0x20, 0x61,
	0x6e, 0x79, 0x2e, 0x32, 0x02, 0x76, 0x31, 0x1a, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x75, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x22, 0x07, 0x2f, 0x6d, 0x2f, 0x62,
	0x6c, 0x75, 0x65, 0x2a, 0x01, 0x02, 0x0a, 0x15, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x42, 0x08, 0x4f,
	0x72, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
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

var file_org_v1_org_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_org_v1_org_proto_goTypes = []any{
	(*CreateOrgRequest)(nil),        // 0: blueapi.org.v1.CreateOrgRequest
	(*CreateOrgResponse)(nil),       // 1: blueapi.org.v1.CreateOrgResponse
	(*SendVerificationRequest)(nil), // 2: blueapi.org.v1.SendVerificationRequest
	(*VerifyOrgRequest)(nil),        // 3: blueapi.org.v1.VerifyOrgRequest
	(*GetOrgRequest)(nil),           // 4: blueapi.org.v1.GetOrgRequest
	(*UpdateMetadataRequest)(nil),   // 5: blueapi.org.v1.UpdateMetadataRequest
	(*UpdatePasswordRequest)(nil),   // 6: blueapi.org.v1.UpdatePasswordRequest
	(*DeleteOrgRequest)(nil),        // 7: blueapi.org.v1.DeleteOrgRequest
	(*ripple.Org)(nil),              // 8: blueapi.api.ripple.Org
	(*emptypb.Empty)(nil),           // 9: google.protobuf.Empty
}
var file_org_v1_org_proto_depIdxs = []int32{
	8, // 0: blueapi.org.v1.CreateOrgResponse.org:type_name -> blueapi.api.ripple.Org
	0, // 1: blueapi.org.v1.Organization.CreateOrg:input_type -> blueapi.org.v1.CreateOrgRequest
	2, // 2: blueapi.org.v1.Organization.SendVerification:input_type -> blueapi.org.v1.SendVerificationRequest
	3, // 3: blueapi.org.v1.Organization.VerifyOrg:input_type -> blueapi.org.v1.VerifyOrgRequest
	4, // 4: blueapi.org.v1.Organization.GetOrg:input_type -> blueapi.org.v1.GetOrgRequest
	5, // 5: blueapi.org.v1.Organization.UpdateMetadata:input_type -> blueapi.org.v1.UpdateMetadataRequest
	6, // 6: blueapi.org.v1.Organization.UpdatePassword:input_type -> blueapi.org.v1.UpdatePasswordRequest
	7, // 7: blueapi.org.v1.Organization.DeleteOrg:input_type -> blueapi.org.v1.DeleteOrgRequest
	1, // 8: blueapi.org.v1.Organization.CreateOrg:output_type -> blueapi.org.v1.CreateOrgResponse
	9, // 9: blueapi.org.v1.Organization.SendVerification:output_type -> google.protobuf.Empty
	9, // 10: blueapi.org.v1.Organization.VerifyOrg:output_type -> google.protobuf.Empty
	8, // 11: blueapi.org.v1.Organization.GetOrg:output_type -> blueapi.api.ripple.Org
	9, // 12: blueapi.org.v1.Organization.UpdateMetadata:output_type -> google.protobuf.Empty
	9, // 13: blueapi.org.v1.Organization.UpdatePassword:output_type -> google.protobuf.Empty
	9, // 14: blueapi.org.v1.Organization.DeleteOrg:output_type -> google.protobuf.Empty
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_org_v1_org_proto_init() }
func file_org_v1_org_proto_init() {
	if File_org_v1_org_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_org_v1_org_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
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
