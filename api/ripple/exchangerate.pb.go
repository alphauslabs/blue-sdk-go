// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/ripple/exchangerate.proto

package ripple

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

// ExchangeRate resource definition.
type ExchangeRate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The currency code.
	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	// The rate.
	Rate          float64 `protobuf:"fixed64,2,opt,name=rate,proto3" json:"rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExchangeRate) Reset() {
	*x = ExchangeRate{}
	mi := &file_api_ripple_exchangerate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRate) ProtoMessage() {}

func (x *ExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_exchangerate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRate.ProtoReflect.Descriptor instead.
func (*ExchangeRate) Descriptor() ([]byte, []int) {
	return file_api_ripple_exchangerate_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeRate) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ExchangeRate) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

// MonthlyExchangeRate resource definition.
type MonthlyExchangeRate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The year-month. format: yyyymm
	Month string `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty"`
	// The exchange rate.
	ExchangeRate  []*ExchangeRate `protobuf:"bytes,2,rep,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MonthlyExchangeRate) Reset() {
	*x = MonthlyExchangeRate{}
	mi := &file_api_ripple_exchangerate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MonthlyExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthlyExchangeRate) ProtoMessage() {}

func (x *MonthlyExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_exchangerate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthlyExchangeRate.ProtoReflect.Descriptor instead.
func (*MonthlyExchangeRate) Descriptor() ([]byte, []int) {
	return file_api_ripple_exchangerate_proto_rawDescGZIP(), []int{1}
}

func (x *MonthlyExchangeRate) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *MonthlyExchangeRate) GetExchangeRate() []*ExchangeRate {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

// BillingGroupExchangeRate resource definition.
type BillingGroupExchangeRate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The billing group's internal id.
	BillingInternalId string `protobuf:"bytes,1,opt,name=billingInternalId,proto3" json:"billingInternalId,omitempty"`
	// The billing group's id.
	BillingGroupId string `protobuf:"bytes,2,opt,name=billingGroupId,proto3" json:"billingGroupId,omitempty"`
	// The billing group's name.
	BillingGroupName string `protobuf:"bytes,3,opt,name=billingGroupName,proto3" json:"billingGroupName,omitempty"`
	// The exchange rate.
	ExchangeRate  []*ExchangeRate `protobuf:"bytes,4,rep,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BillingGroupExchangeRate) Reset() {
	*x = BillingGroupExchangeRate{}
	mi := &file_api_ripple_exchangerate_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BillingGroupExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingGroupExchangeRate) ProtoMessage() {}

func (x *BillingGroupExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_exchangerate_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingGroupExchangeRate.ProtoReflect.Descriptor instead.
func (*BillingGroupExchangeRate) Descriptor() ([]byte, []int) {
	return file_api_ripple_exchangerate_proto_rawDescGZIP(), []int{2}
}

func (x *BillingGroupExchangeRate) GetBillingInternalId() string {
	if x != nil {
		return x.BillingInternalId
	}
	return ""
}

func (x *BillingGroupExchangeRate) GetBillingGroupId() string {
	if x != nil {
		return x.BillingGroupId
	}
	return ""
}

func (x *BillingGroupExchangeRate) GetBillingGroupName() string {
	if x != nil {
		return x.BillingGroupName
	}
	return ""
}

func (x *BillingGroupExchangeRate) GetExchangeRate() []*ExchangeRate {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

// PayerExchangeRate resource definition.
type PayerExchangeRate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The payer account Id.
	PayerAccountId string `protobuf:"bytes,1,opt,name=payerAccountId,proto3" json:"payerAccountId,omitempty"`
	// The payer account Name.
	PayerAccountName string `protobuf:"bytes,2,opt,name=payerAccountName,proto3" json:"payerAccountName,omitempty"`
	// The exchange rate.
	ExchangeRate  []*ExchangeRate `protobuf:"bytes,4,rep,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PayerExchangeRate) Reset() {
	*x = PayerExchangeRate{}
	mi := &file_api_ripple_exchangerate_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PayerExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayerExchangeRate) ProtoMessage() {}

func (x *PayerExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_exchangerate_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayerExchangeRate.ProtoReflect.Descriptor instead.
func (*PayerExchangeRate) Descriptor() ([]byte, []int) {
	return file_api_ripple_exchangerate_proto_rawDescGZIP(), []int{3}
}

func (x *PayerExchangeRate) GetPayerAccountId() string {
	if x != nil {
		return x.PayerAccountId
	}
	return ""
}

func (x *PayerExchangeRate) GetPayerAccountName() string {
	if x != nil {
		return x.PayerAccountName
	}
	return ""
}

func (x *PayerExchangeRate) GetExchangeRate() []*ExchangeRate {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

// CommonExchangeRate resource definition.
type CommonExchangeRate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Aws           []*ExchangeRate        `protobuf:"bytes,1,rep,name=aws,proto3" json:"aws,omitempty"`
	Gcp           []*ExchangeRate        `protobuf:"bytes,2,rep,name=gcp,proto3" json:"gcp,omitempty"`
	Azure         []*ExchangeRate        `protobuf:"bytes,3,rep,name=azure,proto3" json:"azure,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommonExchangeRate) Reset() {
	*x = CommonExchangeRate{}
	mi := &file_api_ripple_exchangerate_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommonExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonExchangeRate) ProtoMessage() {}

func (x *CommonExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_exchangerate_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonExchangeRate.ProtoReflect.Descriptor instead.
func (*CommonExchangeRate) Descriptor() ([]byte, []int) {
	return file_api_ripple_exchangerate_proto_rawDescGZIP(), []int{4}
}

func (x *CommonExchangeRate) GetAws() []*ExchangeRate {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *CommonExchangeRate) GetGcp() []*ExchangeRate {
	if x != nil {
		return x.Gcp
	}
	return nil
}

func (x *CommonExchangeRate) GetAzure() []*ExchangeRate {
	if x != nil {
		return x.Azure
	}
	return nil
}

// VendorPayerExchangeRate resource definition.
type VendorPayerExchangeRate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Aws           []*PayerExchangeRate   `protobuf:"bytes,1,rep,name=aws,proto3" json:"aws,omitempty"`
	Gcp           []*PayerExchangeRate   `protobuf:"bytes,2,rep,name=gcp,proto3" json:"gcp,omitempty"`
	Azure         []*PayerExchangeRate   `protobuf:"bytes,3,rep,name=azure,proto3" json:"azure,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VendorPayerExchangeRate) Reset() {
	*x = VendorPayerExchangeRate{}
	mi := &file_api_ripple_exchangerate_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VendorPayerExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorPayerExchangeRate) ProtoMessage() {}

func (x *VendorPayerExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_exchangerate_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorPayerExchangeRate.ProtoReflect.Descriptor instead.
func (*VendorPayerExchangeRate) Descriptor() ([]byte, []int) {
	return file_api_ripple_exchangerate_proto_rawDescGZIP(), []int{5}
}

func (x *VendorPayerExchangeRate) GetAws() []*PayerExchangeRate {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *VendorPayerExchangeRate) GetGcp() []*PayerExchangeRate {
	if x != nil {
		return x.Gcp
	}
	return nil
}

func (x *VendorPayerExchangeRate) GetAzure() []*PayerExchangeRate {
	if x != nil {
		return x.Azure
	}
	return nil
}

var File_api_ripple_exchangerate_proto protoreflect.FileDescriptor

var file_api_ripple_exchangerate_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70,
	0x70, 0x6c, 0x65, 0x22, 0x3e, 0x0a, 0x0c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72,
	0x61, 0x74, 0x65, 0x22, 0x71, 0x0a, 0x13, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x12, 0x44, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x22, 0xe2, 0x01, 0x0a, 0x18, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65,
	0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x11,
	0x50, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x61, 0x79,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65,
	0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x12,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69,
	0x70, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x32, 0x0a, 0x03, 0x67, 0x63, 0x70, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x03, 0x67, 0x63, 0x70, 0x12, 0x36, 0x0a, 0x05, 0x61, 0x7a,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x05, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x22, 0xc8, 0x01, 0x0a, 0x17, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x50, 0x61, 0x79,
	0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x37,
	0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65,
	0x2e, 0x50, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x37, 0x0a, 0x03, 0x67, 0x63, 0x70, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x65, 0x72, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x03, 0x67, 0x63, 0x70,
	0x12, 0x3b, 0x0a, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69,
	0x70, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x42, 0x6d, 0x0a,
	0x20, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62,
	0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c,
	0x65, 0x42, 0x1a, 0x41, 0x70, 0x69, 0x52, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75,
	0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_ripple_exchangerate_proto_rawDescOnce sync.Once
	file_api_ripple_exchangerate_proto_rawDescData []byte
)

func file_api_ripple_exchangerate_proto_rawDescGZIP() []byte {
	file_api_ripple_exchangerate_proto_rawDescOnce.Do(func() {
		file_api_ripple_exchangerate_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_ripple_exchangerate_proto_rawDesc), len(file_api_ripple_exchangerate_proto_rawDesc)))
	})
	return file_api_ripple_exchangerate_proto_rawDescData
}

var file_api_ripple_exchangerate_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_ripple_exchangerate_proto_goTypes = []any{
	(*ExchangeRate)(nil),             // 0: blueapi.api.ripple.ExchangeRate
	(*MonthlyExchangeRate)(nil),      // 1: blueapi.api.ripple.MonthlyExchangeRate
	(*BillingGroupExchangeRate)(nil), // 2: blueapi.api.ripple.BillingGroupExchangeRate
	(*PayerExchangeRate)(nil),        // 3: blueapi.api.ripple.PayerExchangeRate
	(*CommonExchangeRate)(nil),       // 4: blueapi.api.ripple.CommonExchangeRate
	(*VendorPayerExchangeRate)(nil),  // 5: blueapi.api.ripple.VendorPayerExchangeRate
}
var file_api_ripple_exchangerate_proto_depIdxs = []int32{
	0, // 0: blueapi.api.ripple.MonthlyExchangeRate.exchangeRate:type_name -> blueapi.api.ripple.ExchangeRate
	0, // 1: blueapi.api.ripple.BillingGroupExchangeRate.exchangeRate:type_name -> blueapi.api.ripple.ExchangeRate
	0, // 2: blueapi.api.ripple.PayerExchangeRate.exchangeRate:type_name -> blueapi.api.ripple.ExchangeRate
	0, // 3: blueapi.api.ripple.CommonExchangeRate.aws:type_name -> blueapi.api.ripple.ExchangeRate
	0, // 4: blueapi.api.ripple.CommonExchangeRate.gcp:type_name -> blueapi.api.ripple.ExchangeRate
	0, // 5: blueapi.api.ripple.CommonExchangeRate.azure:type_name -> blueapi.api.ripple.ExchangeRate
	3, // 6: blueapi.api.ripple.VendorPayerExchangeRate.aws:type_name -> blueapi.api.ripple.PayerExchangeRate
	3, // 7: blueapi.api.ripple.VendorPayerExchangeRate.gcp:type_name -> blueapi.api.ripple.PayerExchangeRate
	3, // 8: blueapi.api.ripple.VendorPayerExchangeRate.azure:type_name -> blueapi.api.ripple.PayerExchangeRate
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_api_ripple_exchangerate_proto_init() }
func file_api_ripple_exchangerate_proto_init() {
	if File_api_ripple_exchangerate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_ripple_exchangerate_proto_rawDesc), len(file_api_ripple_exchangerate_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_ripple_exchangerate_proto_goTypes,
		DependencyIndexes: file_api_ripple_exchangerate_proto_depIdxs,
		MessageInfos:      file_api_ripple_exchangerate_proto_msgTypes,
	}.Build()
	File_api_ripple_exchangerate_proto = out.File
	file_api_ripple_exchangerate_proto_goTypes = nil
	file_api_ripple_exchangerate_proto_depIdxs = nil
}
