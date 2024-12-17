// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: api/account.proto

package api

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

type Account struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The vendor
	Vendor string `protobuf:"bytes,6,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// The account id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The account name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The parent organization id.
	OrgId string `protobuf:"bytes,4,opt,name=orgId,proto3" json:"orgId,omitempty"`
	// The parent billing internal id.
	BillingInternalId string `protobuf:"bytes,5,opt,name=billingInternalId,proto3" json:"billingInternalId,omitempty"`
	// The attributes (key/value pair) of the account.
	Metadata      []*KeyValue `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_api_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_api_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Account) GetBillingInternalId() string {
	if x != nil {
		return x.BillingInternalId
	}
	return ""
}

func (x *Account) GetMetadata() []*KeyValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// AccountOriginalResource resource definition.
type AccountOriginalResource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id is associated with CostUsage or BillingReport.
	// Vendor positioning is
	// AWS: `lineItem/UsageAccountId`
	// GoogleCloud: `projectId`
	// Azure: `azure customer id`
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name is associated with CostUsage or BillingReport.
	// Basically the same as id.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The payerId associated with CostUsage or BillingReport.
	// It will be the organization Id associated with the id.
	PayerId string `protobuf:"bytes,3,opt,name=payerId,proto3" json:"payerId,omitempty"`
	// The vendor. supported by `aws`,`gcp`,`azure`
	Vendor string `protobuf:"bytes,4,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// The ProductRegistered indicates whether it is linked to a product(Ripple or Wave(Pro)).
	ProductRegistered bool `protobuf:"varint,5,opt,name=productRegistered,proto3" json:"productRegistered,omitempty"`
	// The attributes (key/value pair) of the account original resource.
	Metadata      []*KeyValueMap `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountOriginalResource) Reset() {
	*x = AccountOriginalResource{}
	mi := &file_api_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountOriginalResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountOriginalResource) ProtoMessage() {}

func (x *AccountOriginalResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountOriginalResource.ProtoReflect.Descriptor instead.
func (*AccountOriginalResource) Descriptor() ([]byte, []int) {
	return file_api_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountOriginalResource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountOriginalResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountOriginalResource) GetPayerId() string {
	if x != nil {
		return x.PayerId
	}
	return ""
}

func (x *AccountOriginalResource) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *AccountOriginalResource) GetProductRegistered() bool {
	if x != nil {
		return x.ProductRegistered
	}
	return false
}

func (x *AccountOriginalResource) GetMetadata() []*KeyValueMap {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_api_account_proto protoreflect.FileDescriptor

var file_api_account_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x31, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xd3, 0x01, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x54, 0x0a, 0x19, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0f, 0x41, 0x70, 0x69, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_account_proto_rawDescOnce sync.Once
	file_api_account_proto_rawDescData = file_api_account_proto_rawDesc
)

func file_api_account_proto_rawDescGZIP() []byte {
	file_api_account_proto_rawDescOnce.Do(func() {
		file_api_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_account_proto_rawDescData)
	})
	return file_api_account_proto_rawDescData
}

var file_api_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_account_proto_goTypes = []any{
	(*Account)(nil),                 // 0: blueapi.api.Account
	(*AccountOriginalResource)(nil), // 1: blueapi.api.AccountOriginalResource
	(*KeyValue)(nil),                // 2: blueapi.api.KeyValue
	(*KeyValueMap)(nil),             // 3: blueapi.api.KeyValueMap
}
var file_api_account_proto_depIdxs = []int32{
	2, // 0: blueapi.api.Account.metadata:type_name -> blueapi.api.KeyValue
	3, // 1: blueapi.api.AccountOriginalResource.metadata:type_name -> blueapi.api.KeyValueMap
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_account_proto_init() }
func file_api_account_proto_init() {
	if File_api_account_proto != nil {
		return
	}
	file_api_keyvalue_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_account_proto_goTypes,
		DependencyIndexes: file_api_account_proto_depIdxs,
		MessageInfos:      file_api_account_proto_msgTypes,
	}.Build()
	File_api_account_proto = out.File
	file_api_account_proto_rawDesc = nil
	file_api_account_proto_goTypes = nil
	file_api_account_proto_depIdxs = nil
}
