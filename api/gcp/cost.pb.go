// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/gcp/cost.proto

package gcp

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

// Cost for Api Response
//
// see gcp billing data schema details:[https://cloud.google.com/billing/docs/how-to/export-data-bigquery-tables]
type Cost struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// account data
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// The group id the account is associated with during the query.
	GroupId string `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	// For daily data, format is `yyyy-mm-dd`; for monthly, `yyyy-mm`.
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	// The invoice.month of the lineitem, if applicable.
	InvoiceMonth string `protobuf:"bytes,12,opt,name=invoiceMonth,proto3" json:"invoiceMonth,omitempty"`
	// The service.Description of the lineitem, if applicable.
	Service string `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
	// The sku.Description of the lineitem, if applicable.
	Sku string `protobuf:"bytes,5,opt,name=sku,proto3" json:"sku,omitempty"`
	// The location.region of the lineitem, if applicable.
	Region string `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	// The location.zone of the lineitem, if applicable.
	Zone string `protobuf:"bytes,7,opt,name=zone,proto3" json:"zone,omitempty"`
	// The credits.type of the lineitem, if applicable.
	CreditsType string `protobuf:"bytes,13,opt,name=creditsType,proto3" json:"creditsType,omitempty"`
	// The credits.name of the lineitem, if applicable.
	CreditsName string `protobuf:"bytes,14,opt,name=creditsName,proto3" json:"creditsName,omitempty"`
	// The usage.pricing_unit of the lineitem, if applicable.
	UsageUnit string `protobuf:"bytes,8,opt,name=usageUnit,proto3" json:"usageUnit,omitempty"`
	// The usage.amount_in_pricing_units of the lineitem, if applicable.
	UsageAmount float64 `protobuf:"fixed64,9,opt,name=usageAmount,proto3" json:"usageAmount,omitempty"`
	// The currency of the lineitem, if applicable.
	BaseCurrency string `protobuf:"bytes,10,opt,name=baseCurrency,proto3" json:"baseCurrency,omitempty"`
	// The cost of the lineitem, if applicable.
	Cost          float64 `protobuf:"fixed64,11,opt,name=cost,proto3" json:"cost,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cost) Reset() {
	*x = Cost{}
	mi := &file_api_gcp_cost_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cost) ProtoMessage() {}

func (x *Cost) ProtoReflect() protoreflect.Message {
	mi := &file_api_gcp_cost_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cost.ProtoReflect.Descriptor instead.
func (*Cost) Descriptor() ([]byte, []int) {
	return file_api_gcp_cost_proto_rawDescGZIP(), []int{0}
}

func (x *Cost) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Cost) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Cost) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Cost) GetInvoiceMonth() string {
	if x != nil {
		return x.InvoiceMonth
	}
	return ""
}

func (x *Cost) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Cost) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *Cost) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Cost) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Cost) GetCreditsType() string {
	if x != nil {
		return x.CreditsType
	}
	return ""
}

func (x *Cost) GetCreditsName() string {
	if x != nil {
		return x.CreditsName
	}
	return ""
}

func (x *Cost) GetUsageUnit() string {
	if x != nil {
		return x.UsageUnit
	}
	return ""
}

func (x *Cost) GetUsageAmount() float64 {
	if x != nil {
		return x.UsageAmount
	}
	return 0
}

func (x *Cost) GetBaseCurrency() string {
	if x != nil {
		return x.BaseCurrency
	}
	return ""
}

func (x *Cost) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

var File_api_gcp_cost_proto protoreflect.FileDescriptor

var file_api_gcp_cost_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x63, 0x70, 0x22, 0x86, 0x03, 0x0a, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f,
	0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61,
	0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x42, 0x5c,
	0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x63, 0x70, 0x42,
	0x0f, 0x41, 0x70, 0x69, 0x47, 0x63, 0x70, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x63, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_gcp_cost_proto_rawDescOnce sync.Once
	file_api_gcp_cost_proto_rawDescData []byte
)

func file_api_gcp_cost_proto_rawDescGZIP() []byte {
	file_api_gcp_cost_proto_rawDescOnce.Do(func() {
		file_api_gcp_cost_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_gcp_cost_proto_rawDesc), len(file_api_gcp_cost_proto_rawDesc)))
	})
	return file_api_gcp_cost_proto_rawDescData
}

var file_api_gcp_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_gcp_cost_proto_goTypes = []any{
	(*Cost)(nil), // 0: blueapi.api.gcp.Cost
}
var file_api_gcp_cost_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_gcp_cost_proto_init() }
func file_api_gcp_cost_proto_init() {
	if File_api_gcp_cost_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_gcp_cost_proto_rawDesc), len(file_api_gcp_cost_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_gcp_cost_proto_goTypes,
		DependencyIndexes: file_api_gcp_cost_proto_depIdxs,
		MessageInfos:      file_api_gcp_cost_proto_msgTypes,
	}.Build()
	File_api_gcp_cost_proto = out.File
	file_api_gcp_cost_proto_goTypes = nil
	file_api_gcp_cost_proto_depIdxs = nil
}
