// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: awscost/v1/awscost.proto

package awscost

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Response message for AwsCost.StreamReadAccountCosts.
type Cost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account      string               `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Date         *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	ProductCode  string               `protobuf:"bytes,3,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	ServiceCode  string               `protobuf:"bytes,4,opt,name=service_code,json=serviceCode,proto3" json:"service_code,omitempty"`
	Region       string               `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	Zone         string               `protobuf:"bytes,6,opt,name=zone,proto3" json:"zone,omitempty"`
	UsageType    string               `protobuf:"bytes,7,opt,name=usage_type,json=usageType,proto3" json:"usage_type,omitempty"`
	InstanceType string               `protobuf:"bytes,8,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	Operation    string               `protobuf:"bytes,9,opt,name=operation,proto3" json:"operation,omitempty"`
	InvoiceId    string               `protobuf:"bytes,10,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	Description  string               `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	UsageAmount  float64              `protobuf:"fixed64,12,opt,name=usage_amount,json=usageAmount,proto3" json:"usage_amount,omitempty"`
	Cost         float64              `protobuf:"fixed64,13,opt,name=cost,proto3" json:"cost,omitempty"`
	BlendedCost  float64              `protobuf:"fixed64,14,opt,name=blended_cost,json=blendedCost,proto3" json:"blended_cost,omitempty"`
}

func (x *Cost) Reset() {
	*x = Cost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awscost_v1_awscost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cost) ProtoMessage() {}

func (x *Cost) ProtoReflect() protoreflect.Message {
	mi := &file_awscost_v1_awscost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_awscost_v1_awscost_proto_rawDescGZIP(), []int{0}
}

func (x *Cost) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Cost) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Cost) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *Cost) GetServiceCode() string {
	if x != nil {
		return x.ServiceCode
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

func (x *Cost) GetUsageType() string {
	if x != nil {
		return x.UsageType
	}
	return ""
}

func (x *Cost) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *Cost) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *Cost) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *Cost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Cost) GetUsageAmount() float64 {
	if x != nil {
		return x.UsageAmount
	}
	return 0
}

func (x *Cost) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Cost) GetBlendedCost() float64 {
	if x != nil {
		return x.BlendedCost
	}
	return 0
}

// Request message for AwsCost.StreamReadAccountCosts.
type StreamReadAccountCostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The AWS account to stream.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Timestamp to start streaming data from. If not set, the first day of the
	// current month will be used.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Optional. Timestamp to end the streaming data. If not set, current date will be used.
	EndTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *StreamReadAccountCostsRequest) Reset() {
	*x = StreamReadAccountCostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awscost_v1_awscost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamReadAccountCostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamReadAccountCostsRequest) ProtoMessage() {}

func (x *StreamReadAccountCostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_awscost_v1_awscost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamReadAccountCostsRequest.ProtoReflect.Descriptor instead.
func (*StreamReadAccountCostsRequest) Descriptor() ([]byte, []int) {
	return file_awscost_v1_awscost_proto_rawDescGZIP(), []int{1}
}

func (x *StreamReadAccountCostsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StreamReadAccountCostsRequest) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *StreamReadAccountCostsRequest) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_awscost_v1_awscost_proto protoreflect.FileDescriptor

var file_awscost_v1_awscost_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x77, 0x73, 0x63, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73,
	0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf,
	0x03, 0x0a, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a,
	0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x62, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74,
	0x22, 0xab, 0x01, 0x0a, 0x1d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xb0,
	0x01, 0x0a, 0x07, 0x41, 0x77, 0x73, 0x43, 0x6f, 0x73, 0x74, 0x12, 0xa4, 0x01, 0x0a, 0x16, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x31, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x77, 0x73, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x61, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x73, 0x74, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x22, 0x30, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f,
	0x63, 0x6f, 0x73, 0x74, 0x73, 0x3a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x30,
	0x01, 0x42, 0x4a, 0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x63, 0x6f, 0x73, 0x74, 0x42, 0x05,
	0x43, 0x6f, 0x73, 0x74, 0x73, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x77, 0x73, 0x63, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_awscost_v1_awscost_proto_rawDescOnce sync.Once
	file_awscost_v1_awscost_proto_rawDescData = file_awscost_v1_awscost_proto_rawDesc
)

func file_awscost_v1_awscost_proto_rawDescGZIP() []byte {
	file_awscost_v1_awscost_proto_rawDescOnce.Do(func() {
		file_awscost_v1_awscost_proto_rawDescData = protoimpl.X.CompressGZIP(file_awscost_v1_awscost_proto_rawDescData)
	})
	return file_awscost_v1_awscost_proto_rawDescData
}

var file_awscost_v1_awscost_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_awscost_v1_awscost_proto_goTypes = []interface{}{
	(*Cost)(nil),                          // 0: blueapi.awscost.v1.Cost
	(*StreamReadAccountCostsRequest)(nil), // 1: blueapi.awscost.v1.StreamReadAccountCostsRequest
	(*timestamp.Timestamp)(nil),           // 2: google.protobuf.Timestamp
}
var file_awscost_v1_awscost_proto_depIdxs = []int32{
	2, // 0: blueapi.awscost.v1.Cost.date:type_name -> google.protobuf.Timestamp
	2, // 1: blueapi.awscost.v1.StreamReadAccountCostsRequest.start_time:type_name -> google.protobuf.Timestamp
	2, // 2: blueapi.awscost.v1.StreamReadAccountCostsRequest.end_time:type_name -> google.protobuf.Timestamp
	1, // 3: blueapi.awscost.v1.AwsCost.StreamReadAccountCosts:input_type -> blueapi.awscost.v1.StreamReadAccountCostsRequest
	0, // 4: blueapi.awscost.v1.AwsCost.StreamReadAccountCosts:output_type -> blueapi.awscost.v1.Cost
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_awscost_v1_awscost_proto_init() }
func file_awscost_v1_awscost_proto_init() {
	if File_awscost_v1_awscost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_awscost_v1_awscost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cost); i {
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
		file_awscost_v1_awscost_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamReadAccountCostsRequest); i {
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
			RawDescriptor: file_awscost_v1_awscost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_awscost_v1_awscost_proto_goTypes,
		DependencyIndexes: file_awscost_v1_awscost_proto_depIdxs,
		MessageInfos:      file_awscost_v1_awscost_proto_msgTypes,
	}.Build()
	File_awscost_v1_awscost_proto = out.File
	file_awscost_v1_awscost_proto_rawDesc = nil
	file_awscost_v1_awscost_proto_goTypes = nil
	file_awscost_v1_awscost_proto_depIdxs = nil
}
