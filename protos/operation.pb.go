// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: protos/operation.proto

package protos

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The server-assigned name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata. Any method that returns a
	// long-running operation should document the metadata type, if any.
	Metadata *anypb.Any `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// If the value is `false`, it means the operation is still in progress.
	// If `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	// The operation result, which can be either an `error` or a valid `response`.
	// If `done` == `false`, neither `error` nor `response` is set.
	// If `done` == `true`, exactly one of `error` or `response` is set.
	//
	// Types that are assignable to Result:
	//
	//	*Operation_Error
	//	*Operation_Response
	Result isOperation_Result `protobuf_oneof:"result"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_protos_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_protos_operation_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Operation) GetMetadata() *anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Operation) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (m *Operation) GetResult() isOperation_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *Operation) GetError() *status.Status {
	if x, ok := x.GetResult().(*Operation_Error); ok {
		return x.Error
	}
	return nil
}

func (x *Operation) GetResponse() *anypb.Any {
	if x, ok := x.GetResult().(*Operation_Response); ok {
		return x.Response
	}
	return nil
}

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	// The error result of the operation in case of failure or cancellation.
	Error *status.Status `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

type Operation_Response struct {
	// The normal response of the operation in case of success. If the original method returns
	// no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the
	// original method is standard `Get`/`Create`/`Update`, the response should be the
	// resource. For other methods, the response should have the type `XxxResponse`, where
	// `Xxx` is the original method name. For example, if the original method name is
	// `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
	Response *anypb.Any `protobuf:"bytes,5,opt,name=response,proto3,oneof"`
}

func (*Operation_Error) isOperation_Result() {}

func (*Operation_Response) isOperation_Result() {}

type OperationImportCurMetadataV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Month    string   `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty"`
	Accounts []string `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Created  string   `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated  string   `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *OperationImportCurMetadataV1) Reset() {
	*x = OperationImportCurMetadataV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_operation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationImportCurMetadataV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationImportCurMetadataV1) ProtoMessage() {}

func (x *OperationImportCurMetadataV1) ProtoReflect() protoreflect.Message {
	mi := &file_protos_operation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationImportCurMetadataV1.ProtoReflect.Descriptor instead.
func (*OperationImportCurMetadataV1) Descriptor() ([]byte, []int) {
	return file_protos_operation_proto_rawDescGZIP(), []int{1}
}

func (x *OperationImportCurMetadataV1) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *OperationImportCurMetadataV1) GetAccounts() []string {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *OperationImportCurMetadataV1) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *OperationImportCurMetadataV1) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

type OperationAwsCalculateCostsMetadataV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Organization id.
	OrgId string `protobuf:"bytes,1,opt,name=orgId,proto3" json:"orgId,omitempty"`
	// The month being calculated.
	Month string `protobuf:"bytes,2,opt,name=month,proto3" json:"month,omitempty"`
	// If the request is for a specific group(s) (billing groups at the moment),
	// their equivalent ids are listed here. Otherwise, empty.
	GroupIds []string `protobuf:"bytes,3,rep,name=groupIds,proto3" json:"groupIds,omitempty"`
	// Latest status information.
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// Timestamp when operation was created/started in RFC3339 format.
	Created string `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	// Timestamp of the operation's last update in RFC3339 format.
	Updated string `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *OperationAwsCalculateCostsMetadataV1) Reset() {
	*x = OperationAwsCalculateCostsMetadataV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_operation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationAwsCalculateCostsMetadataV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationAwsCalculateCostsMetadataV1) ProtoMessage() {}

func (x *OperationAwsCalculateCostsMetadataV1) ProtoReflect() protoreflect.Message {
	mi := &file_protos_operation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationAwsCalculateCostsMetadataV1.ProtoReflect.Descriptor instead.
func (*OperationAwsCalculateCostsMetadataV1) Descriptor() ([]byte, []int) {
	return file_protos_operation_proto_rawDescGZIP(), []int{2}
}

func (x *OperationAwsCalculateCostsMetadataV1) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *OperationAwsCalculateCostsMetadataV1) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *OperationAwsCalculateCostsMetadataV1) GetGroupIds() []string {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

func (x *OperationAwsCalculateCostsMetadataV1) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OperationAwsCalculateCostsMetadataV1) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *OperationAwsCalculateCostsMetadataV1) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

type OperationAwsDiscoverAssetsMetadataV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Services []string `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *OperationAwsDiscoverAssetsMetadataV1) Reset() {
	*x = OperationAwsDiscoverAssetsMetadataV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_operation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationAwsDiscoverAssetsMetadataV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationAwsDiscoverAssetsMetadataV1) ProtoMessage() {}

func (x *OperationAwsDiscoverAssetsMetadataV1) ProtoReflect() protoreflect.Message {
	mi := &file_protos_operation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationAwsDiscoverAssetsMetadataV1.ProtoReflect.Descriptor instead.
func (*OperationAwsDiscoverAssetsMetadataV1) Descriptor() ([]byte, []int) {
	return file_protos_operation_proto_rawDescGZIP(), []int{3}
}

func (x *OperationAwsDiscoverAssetsMetadataV1) GetAccounts() []string {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *OperationAwsDiscoverAssetsMetadataV1) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_protos_operation_proto protoreflect.FileDescriptor

var file_protos_operation_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x1c, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x75, 0x72, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0xba, 0x01,
	0x0a, 0x24, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x77, 0x73, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x56, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x5e, 0x0a, 0x24, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x77, 0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x56, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x45, 0x0a, 0x14, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x42, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_operation_proto_rawDescOnce sync.Once
	file_protos_operation_proto_rawDescData = file_protos_operation_proto_rawDesc
)

func file_protos_operation_proto_rawDescGZIP() []byte {
	file_protos_operation_proto_rawDescOnce.Do(func() {
		file_protos_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_operation_proto_rawDescData)
	})
	return file_protos_operation_proto_rawDescData
}

var file_protos_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_operation_proto_goTypes = []interface{}{
	(*Operation)(nil),                            // 0: protos.Operation
	(*OperationImportCurMetadataV1)(nil),         // 1: protos.OperationImportCurMetadataV1
	(*OperationAwsCalculateCostsMetadataV1)(nil), // 2: protos.OperationAwsCalculateCostsMetadataV1
	(*OperationAwsDiscoverAssetsMetadataV1)(nil), // 3: protos.OperationAwsDiscoverAssetsMetadataV1
	(*anypb.Any)(nil),                            // 4: google.protobuf.Any
	(*status.Status)(nil),                        // 5: google.rpc.Status
}
var file_protos_operation_proto_depIdxs = []int32{
	4, // 0: protos.Operation.metadata:type_name -> google.protobuf.Any
	5, // 1: protos.Operation.error:type_name -> google.rpc.Status
	4, // 2: protos.Operation.response:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_operation_proto_init() }
func file_protos_operation_proto_init() {
	if File_protos_operation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_protos_operation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationImportCurMetadataV1); i {
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
		file_protos_operation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationAwsCalculateCostsMetadataV1); i {
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
		file_protos_operation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationAwsDiscoverAssetsMetadataV1); i {
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
	file_protos_operation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_operation_proto_goTypes,
		DependencyIndexes: file_protos_operation_proto_depIdxs,
		MessageInfos:      file_protos_operation_proto_msgTypes,
	}.Build()
	File_protos_operation_proto = out.File
	file_protos_operation_proto_rawDesc = nil
	file_protos_operation_proto_goTypes = nil
	file_protos_operation_proto_depIdxs = nil
}
