// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/ripple/billinggroup.proto

package ripple

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

// BillingGroupInfo resource definition. Only available in Ripple.
type BillingGroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The title of the billing
	BillingTitle string `protobuf:"bytes,1,opt,name=billingTitle,proto3" json:"billingTitle,omitempty"`
	// Optional. Company’s phone number
	PhoneNumber string `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	// Optional. Company’s postal code
	PostalCode string `protobuf:"bytes,3,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	// Optional. Company’s address
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// Optional. Addressee
	Personal string `protobuf:"bytes,5,opt,name=personal,proto3" json:"personal,omitempty"`
	// Optional. Any remarks about the billing group
	Remarks string `protobuf:"bytes,6,opt,name=remarks,proto3" json:"remarks,omitempty"`
	// Optional. Project code
	ProjectId string `protobuf:"bytes,7,opt,name=projectId,proto3" json:"projectId,omitempty"`
	// Optional. Invoice language
	Language string `protobuf:"bytes,8,opt,name=language,proto3" json:"language,omitempty"`
	// Optional. Calculation type, true unblended or unblended
	DisplayCost string `protobuf:"bytes,9,opt,name=displayCost,proto3" json:"displayCost,omitempty"`
	// Optional. Exchange rate type, payer or billing group
	ExchangeRateType string `protobuf:"bytes,10,opt,name=exchangeRateType,proto3" json:"exchangeRateType,omitempty"`
	// Optional. qrCode
	QrCode bool `protobuf:"varint,11,opt,name=qrCode,proto3" json:"qrCode,omitempty"`
	// Optional. Invoice template Id
	InvoiceTemplateId string `protobuf:"bytes,12,opt,name=invoiceTemplateId,proto3" json:"invoiceTemplateId,omitempty"`
}

func (x *BillingGroupInfo) Reset() {
	*x = BillingGroupInfo{}
	mi := &file_api_ripple_billinggroup_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BillingGroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingGroupInfo) ProtoMessage() {}

func (x *BillingGroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_billinggroup_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingGroupInfo.ProtoReflect.Descriptor instead.
func (*BillingGroupInfo) Descriptor() ([]byte, []int) {
	return file_api_ripple_billinggroup_proto_rawDescGZIP(), []int{0}
}

func (x *BillingGroupInfo) GetBillingTitle() string {
	if x != nil {
		return x.BillingTitle
	}
	return ""
}

func (x *BillingGroupInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *BillingGroupInfo) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *BillingGroupInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BillingGroupInfo) GetPersonal() string {
	if x != nil {
		return x.Personal
	}
	return ""
}

func (x *BillingGroupInfo) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *BillingGroupInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *BillingGroupInfo) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *BillingGroupInfo) GetDisplayCost() string {
	if x != nil {
		return x.DisplayCost
	}
	return ""
}

func (x *BillingGroupInfo) GetExchangeRateType() string {
	if x != nil {
		return x.ExchangeRateType
	}
	return ""
}

func (x *BillingGroupInfo) GetQrCode() bool {
	if x != nil {
		return x.QrCode
	}
	return false
}

func (x *BillingGroupInfo) GetInvoiceTemplateId() string {
	if x != nil {
		return x.InvoiceTemplateId
	}
	return ""
}

var File_api_ripple_billinggroup_proto protoreflect.FileDescriptor

var file_api_ripple_billinggroup_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70,
	0x70, 0x6c, 0x65, 0x22, 0x96, 0x03, 0x0a, 0x10, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x71, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x71, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2c,
	0x0a, 0x11, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x42, 0x71, 0x0a, 0x20,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65,
	0x42, 0x1e, 0x41, 0x70, 0x69, 0x52, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ripple_billinggroup_proto_rawDescOnce sync.Once
	file_api_ripple_billinggroup_proto_rawDescData = file_api_ripple_billinggroup_proto_rawDesc
)

func file_api_ripple_billinggroup_proto_rawDescGZIP() []byte {
	file_api_ripple_billinggroup_proto_rawDescOnce.Do(func() {
		file_api_ripple_billinggroup_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ripple_billinggroup_proto_rawDescData)
	})
	return file_api_ripple_billinggroup_proto_rawDescData
}

var file_api_ripple_billinggroup_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_ripple_billinggroup_proto_goTypes = []any{
	(*BillingGroupInfo)(nil), // 0: blueapi.api.ripple.BillingGroupInfo
}
var file_api_ripple_billinggroup_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_ripple_billinggroup_proto_init() }
func file_api_ripple_billinggroup_proto_init() {
	if File_api_ripple_billinggroup_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_ripple_billinggroup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_ripple_billinggroup_proto_goTypes,
		DependencyIndexes: file_api_ripple_billinggroup_proto_depIdxs,
		MessageInfos:      file_api_ripple_billinggroup_proto_msgTypes,
	}.Build()
	File_api_ripple_billinggroup_proto = out.File
	file_api_ripple_billinggroup_proto_rawDesc = nil
	file_api_ripple_billinggroup_proto_goTypes = nil
	file_api_ripple_billinggroup_proto_depIdxs = nil
}
