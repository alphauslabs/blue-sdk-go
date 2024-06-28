// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: pricing/v1/pricing.proto

package pricing

import (
	pricing "github.com/alphauslabs/blue-sdk-go/api/pricing"
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

// Request message for the Pricing.GetInfo rpc.
type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_v1_pricing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_pricing_v1_pricing_proto_rawDescGZIP(), []int{0}
}

// Response message for the Pricing.GetInfo rpc.
type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_v1_pricing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_pricing_v1_pricing_proto_rawDescGZIP(), []int{1}
}

func (x *GetInfoResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

// Request message for GetPricingInfo
type GetPricingInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cloud   string   `protobuf:"bytes,1,opt,name=cloud,proto3" json:"cloud,omitempty"`
	Service string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Region  string   `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Token   string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Columns []string `protobuf:"bytes,5,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *GetPricingInfoRequest) Reset() {
	*x = GetPricingInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_v1_pricing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPricingInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPricingInfoRequest) ProtoMessage() {}

func (x *GetPricingInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPricingInfoRequest.ProtoReflect.Descriptor instead.
func (*GetPricingInfoRequest) Descriptor() ([]byte, []int) {
	return file_pricing_v1_pricing_proto_rawDescGZIP(), []int{2}
}

func (x *GetPricingInfoRequest) GetCloud() string {
	if x != nil {
		return x.Cloud
	}
	return ""
}

func (x *GetPricingInfoRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *GetPricingInfoRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *GetPricingInfoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPricingInfoRequest) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

// Response message for GetPricingInfo
type GetPricingInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cloud          string                  `protobuf:"bytes,1,opt,name=cloud,proto3" json:"cloud,omitempty"`
	Service        string                  `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	RegionCode     string                  `protobuf:"bytes,3,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	Token          string                  `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Sku            string                  `protobuf:"bytes,5,opt,name=sku,proto3" json:"sku,omitempty"`
	Unit           string                  `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	Priceperunit   float64                 `protobuf:"fixed64,7,opt,name=priceperunit,proto3" json:"priceperunit,omitempty"`
	ServiceDetails *pricing.ServiceDetails `protobuf:"bytes,8,opt,name=serviceDetails,proto3" json:"serviceDetails,omitempty"`
}

func (x *GetPricingInfoResponse) Reset() {
	*x = GetPricingInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_v1_pricing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPricingInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPricingInfoResponse) ProtoMessage() {}

func (x *GetPricingInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPricingInfoResponse.ProtoReflect.Descriptor instead.
func (*GetPricingInfoResponse) Descriptor() ([]byte, []int) {
	return file_pricing_v1_pricing_proto_rawDescGZIP(), []int{3}
}

func (x *GetPricingInfoResponse) GetCloud() string {
	if x != nil {
		return x.Cloud
	}
	return ""
}

func (x *GetPricingInfoResponse) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *GetPricingInfoResponse) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *GetPricingInfoResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPricingInfoResponse) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *GetPricingInfoResponse) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *GetPricingInfoResponse) GetPriceperunit() float64 {
	if x != nil {
		return x.Priceperunit
	}
	return 0
}

func (x *GetPricingInfoResponse) GetServiceDetails() *pricing.ServiceDetails {
	if x != nil {
		return x.ServiceDetails
	}
	return nil
}

var File_pricing_v1_pricing_proto protoreflect.FileDescriptor

var file_pricing_v1_pricing_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x19,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0x96, 0x02, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x6b, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x70, 0x65, 0x72, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x70,
	0x65, 0x72, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x4b, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x32, 0x95, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12,
	0x64, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x30,
	0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x83, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x30, 0x2f,
	0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x9d, 0x01, 0x92, 0x41,
	0x99, 0x01, 0x12, 0x46, 0x28, 0x42, 0x45, 0x54, 0x41, 0x29, 0x20, 0x50, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x20, 0x41, 0x50, 0x49, 0x2e, 0x20, 0x42, 0x61, 0x73, 0x65, 0x20, 0x55, 0x52, 0x4c,
	0x3a, 0x20, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x2f, 0x62, 0x6c,
	0x75, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x1a, 0x4f, 0x0a, 0x12, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x39, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x42, 0x51, 0x0a, 0x19, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x42, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62,
	0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pricing_v1_pricing_proto_rawDescOnce sync.Once
	file_pricing_v1_pricing_proto_rawDescData = file_pricing_v1_pricing_proto_rawDesc
)

func file_pricing_v1_pricing_proto_rawDescGZIP() []byte {
	file_pricing_v1_pricing_proto_rawDescOnce.Do(func() {
		file_pricing_v1_pricing_proto_rawDescData = protoimpl.X.CompressGZIP(file_pricing_v1_pricing_proto_rawDescData)
	})
	return file_pricing_v1_pricing_proto_rawDescData
}

var file_pricing_v1_pricing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pricing_v1_pricing_proto_goTypes = []any{
	(*GetInfoRequest)(nil),         // 0: blueapi.pricing.v1.GetInfoRequest
	(*GetInfoResponse)(nil),        // 1: blueapi.pricing.v1.GetInfoResponse
	(*GetPricingInfoRequest)(nil),  // 2: blueapi.pricing.v1.GetPricingInfoRequest
	(*GetPricingInfoResponse)(nil), // 3: blueapi.pricing.v1.GetPricingInfoResponse
	(*pricing.ServiceDetails)(nil), // 4: blueapi.api.pricing.ServiceDetails
}
var file_pricing_v1_pricing_proto_depIdxs = []int32{
	4, // 0: blueapi.pricing.v1.GetPricingInfoResponse.serviceDetails:type_name -> blueapi.api.pricing.ServiceDetails
	0, // 1: blueapi.pricing.v1.Pricing.GetInfo:input_type -> blueapi.pricing.v1.GetInfoRequest
	2, // 2: blueapi.pricing.v1.Pricing.GetPricingInfo:input_type -> blueapi.pricing.v1.GetPricingInfoRequest
	1, // 3: blueapi.pricing.v1.Pricing.GetInfo:output_type -> blueapi.pricing.v1.GetInfoResponse
	3, // 4: blueapi.pricing.v1.Pricing.GetPricingInfo:output_type -> blueapi.pricing.v1.GetPricingInfoResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pricing_v1_pricing_proto_init() }
func file_pricing_v1_pricing_proto_init() {
	if File_pricing_v1_pricing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pricing_v1_pricing_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetInfoRequest); i {
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
		file_pricing_v1_pricing_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetInfoResponse); i {
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
		file_pricing_v1_pricing_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetPricingInfoRequest); i {
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
		file_pricing_v1_pricing_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetPricingInfoResponse); i {
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
			RawDescriptor: file_pricing_v1_pricing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pricing_v1_pricing_proto_goTypes,
		DependencyIndexes: file_pricing_v1_pricing_proto_depIdxs,
		MessageInfos:      file_pricing_v1_pricing_proto_msgTypes,
	}.Build()
	File_pricing_v1_pricing_proto = out.File
	file_pricing_v1_pricing_proto_rawDesc = nil
	file_pricing_v1_pricing_proto_goTypes = nil
	file_pricing_v1_pricing_proto_depIdxs = nil
}
