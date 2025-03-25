// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/account.proto

package api

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

const file_api_account_proto_rawDesc = "" +
	"\n" +
	"\x11api/account.proto\x12\vblueapi.api\x1a\x12api/keyvalue.proto\"\xbc\x01\n" +
	"\aAccount\x12\x16\n" +
	"\x06vendor\x18\x06 \x01(\tR\x06vendor\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05orgId\x18\x04 \x01(\tR\x05orgId\x12,\n" +
	"\x11billingInternalId\x18\x05 \x01(\tR\x11billingInternalId\x121\n" +
	"\bmetadata\x18\x03 \x03(\v2\x15.blueapi.api.KeyValueR\bmetadata\"\xd3\x01\n" +
	"\x17AccountOriginalResource\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x18\n" +
	"\apayerId\x18\x03 \x01(\tR\apayerId\x12\x16\n" +
	"\x06vendor\x18\x04 \x01(\tR\x06vendor\x12,\n" +
	"\x11productRegistered\x18\x05 \x01(\bR\x11productRegistered\x124\n" +
	"\bmetadata\x18\x06 \x03(\v2\x18.blueapi.api.KeyValueMapR\bmetadataBT\n" +
	"\x19cloud.alphaus.blueapi.apiB\x0fApiAccountProtoZ&github.com/alphauslabs/blue-sdk-go/apib\x06proto3"

var (
	file_api_account_proto_rawDescOnce sync.Once
	file_api_account_proto_rawDescData []byte
)

func file_api_account_proto_rawDescGZIP() []byte {
	file_api_account_proto_rawDescOnce.Do(func() {
		file_api_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_account_proto_rawDesc), len(file_api_account_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_account_proto_rawDesc), len(file_api_account_proto_rawDesc)),
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
	file_api_account_proto_goTypes = nil
	file_api_account_proto_depIdxs = nil
}
