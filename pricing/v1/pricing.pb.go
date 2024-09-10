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

// Request message for Pricing.GetPricing rpc.
type GetPricingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Cloud vendor, only `aws` and `azure` are currently supported.
	Vendor string `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Required. Cloud vendor service.
	// Supported services can be listed using `/{vendor}/services` endpoint. For usage information visit https://labs.alphaus.cloud/blueapidocs/#/Pricing/Pricing_GetSupportedServices.
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// Required. Region code.
	// Supported regions can be listed using `/{vendor}/services` endpoint. For usage information visit https://labs.alphaus.cloud/blueapidocs/#/Pricing/Pricing_GetSupportedServices.
	// View all available AWS services by region at https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/.
	// View all available Azure services by region at https://azure.microsoft.com/en-us/explore/global-infrastructure/products-by-region/.
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	// Optional. Supply token that is included in the latest response to continue fetching the remaining chunks of data. No further data can be retrieved once the token returned is empty.
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	// Optional. Filters to apply to the pricing data.
	// This is a map of column names and values to filter pricing items. Each key-value pair in the map represents a filter condition.
	// Supported filter key-value pairs can be listed using `/{vendor}/services` endpoint. For usage information visit https://labs.alphaus.cloud/blueapidocs/#/Pricing/Pricing_GetSupportedServices.
	//
	// For example, if you want to return AWS EC2 items that has All Upfront purchase option, add `"purchaseOption": "All Upfront"` to the filters.
	//
	// Multiple key-value pairs are supported but keys should not be duplicated.
	// For example, for AWS EC2, the following is valid,
	// ```
	//
	//	"filters": {
	//	  "purchaseOption": "All Upfront",
	//	  "operatingSystem": "Windows"
	//	}
	//
	// ```
	// but not the following,
	// ```
	//
	//	"filters": {
	//	  "purchaseOption": "All Upfront",
	//	  "purchaseOption": "Partial Upfront"
	//	}
	//
	// ```
	Filters map[string]string `protobuf:"bytes,6,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. Only specified columns will be returned, if provided.
	// All columns will be returned if this array is empty.
	// Supported columns can be listed using `/{vendor}/services` endpoint. For usage information visit https://labs.alphaus.cloud/blueapidocs/#/Pricing/Pricing_GetSupportedServices.
	Columns []string `protobuf:"bytes,5,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *GetPricingRequest) Reset() {
	*x = GetPricingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_v1_pricing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPricingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPricingRequest) ProtoMessage() {}

func (x *GetPricingRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPricingRequest.ProtoReflect.Descriptor instead.
func (*GetPricingRequest) Descriptor() ([]byte, []int) {
	return file_pricing_v1_pricing_proto_rawDescGZIP(), []int{2}
}

func (x *GetPricingRequest) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *GetPricingRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *GetPricingRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *GetPricingRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPricingRequest) GetFilters() map[string]string {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *GetPricingRequest) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

// Response message for Pricing.GetPricing rpc.
type GetPricingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use token to retrieve next set of pricing items. An empty string means there are no more items to retrieve.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Array of pricing items details. Maximum number of items returned per call is 1000.
	PricingData []*pricing.PricingData `protobuf:"bytes,2,rep,name=pricingData,proto3" json:"pricingData,omitempty"`
}

func (x *GetPricingResponse) Reset() {
	*x = GetPricingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_v1_pricing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPricingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPricingResponse) ProtoMessage() {}

func (x *GetPricingResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPricingResponse.ProtoReflect.Descriptor instead.
func (*GetPricingResponse) Descriptor() ([]byte, []int) {
	return file_pricing_v1_pricing_proto_rawDescGZIP(), []int{3}
}

func (x *GetPricingResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPricingResponse) GetPricingData() []*pricing.PricingData {
	if x != nil {
		return x.PricingData
	}
	return nil
}

// Request message for Pricing.GetSupportedServices rpc.
type GetSupportedServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Cloud vendor, only `aws` and `azure` are currently supported.
	Vendor string `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
}

func (x *GetSupportedServicesRequest) Reset() {
	*x = GetSupportedServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_v1_pricing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupportedServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedServicesRequest) ProtoMessage() {}

func (x *GetSupportedServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportedServicesRequest.ProtoReflect.Descriptor instead.
func (*GetSupportedServicesRequest) Descriptor() ([]byte, []int) {
	return file_pricing_v1_pricing_proto_rawDescGZIP(), []int{4}
}

func (x *GetSupportedServicesRequest) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

// Response message for Pricing.GetSupportedServices rpc.
type GetSupportedServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Suported services, regions, and attributes that can be used to specify which pricing data to retrieve from `/{vendor}/pricing`.
	// For usage information, visit https://labs.alphaus.cloud/blueapidocs/#/Pricing/Pricing_GetPricing.
	SupportedServices []*SupportedService `protobuf:"bytes,1,rep,name=supportedServices,proto3" json:"supportedServices,omitempty"`
}

func (x *GetSupportedServicesResponse) Reset() {
	*x = GetSupportedServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_v1_pricing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupportedServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedServicesResponse) ProtoMessage() {}

func (x *GetSupportedServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportedServicesResponse.ProtoReflect.Descriptor instead.
func (*GetSupportedServicesResponse) Descriptor() ([]byte, []int) {
	return file_pricing_v1_pricing_proto_rawDescGZIP(), []int{5}
}

func (x *GetSupportedServicesResponse) GetSupportedServices() []*SupportedService {
	if x != nil {
		return x.SupportedServices
	}
	return nil
}

type SupportedService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AWS or Azure services only as of now.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// Array of regions supported for the specific service.
	Regions []string `protobuf:"bytes,2,rep,name=regions,proto3" json:"regions,omitempty"`
	// Array of attributes that can be used as key-value pairs for filtering.
	Attributes []*Attribute `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// Array of column names that can be used to specify what columns should `/{vendor}/pricing` return.
	// For usage information, visit https://labs.alphaus.cloud/blueapidocs/#/Pricing/Pricing_GetPricing.
	Columns []string `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *SupportedService) Reset() {
	*x = SupportedService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_v1_pricing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportedService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportedService) ProtoMessage() {}

func (x *SupportedService) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportedService.ProtoReflect.Descriptor instead.
func (*SupportedService) Descriptor() ([]byte, []int) {
	return file_pricing_v1_pricing_proto_rawDescGZIP(), []int{6}
}

func (x *SupportedService) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *SupportedService) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *SupportedService) GetAttributes() []*Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *SupportedService) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filter key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Array of filter values.
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_v1_pricing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_pricing_v1_pricing_proto_rawDescGZIP(), []int{7}
}

func (x *Attribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Attribute) GetValues() []string {
	if x != nil {
		return x.Values
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
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x97, 0x02, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4c,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x6e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x42,
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x35, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x22, 0x72, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x11, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x9f, 0x01,
	0x0a, 0x10, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22,
	0x35, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0xa8, 0x04, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x12, 0x64, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08,
	0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x7c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a,
	0x22, 0x14, 0x2f, 0x76, 0x30, 0x2f, 0x7b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x7d, 0x2f, 0x70,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x98, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x2f, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x30, 0x2f,
	0x7b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x7d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x9d, 0x01, 0x92, 0x41, 0x99, 0x01, 0x12, 0x46, 0x28, 0x42, 0x45, 0x54, 0x41, 0x29,
	0x20, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x20, 0x41, 0x50, 0x49, 0x2e, 0x20, 0x42, 0x61,
	0x73, 0x65, 0x20, 0x55, 0x52, 0x4c, 0x3a, 0x20, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x1a, 0x4f, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x2f, 0x42, 0x51, 0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x42, 0x0c,
	0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_pricing_v1_pricing_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pricing_v1_pricing_proto_goTypes = []any{
	(*GetInfoRequest)(nil),               // 0: blueapi.pricing.v1.GetInfoRequest
	(*GetInfoResponse)(nil),              // 1: blueapi.pricing.v1.GetInfoResponse
	(*GetPricingRequest)(nil),            // 2: blueapi.pricing.v1.GetPricingRequest
	(*GetPricingResponse)(nil),           // 3: blueapi.pricing.v1.GetPricingResponse
	(*GetSupportedServicesRequest)(nil),  // 4: blueapi.pricing.v1.GetSupportedServicesRequest
	(*GetSupportedServicesResponse)(nil), // 5: blueapi.pricing.v1.GetSupportedServicesResponse
	(*SupportedService)(nil),             // 6: blueapi.pricing.v1.SupportedService
	(*Attribute)(nil),                    // 7: blueapi.pricing.v1.Attribute
	nil,                                  // 8: blueapi.pricing.v1.GetPricingRequest.FiltersEntry
	(*pricing.PricingData)(nil),          // 9: blueapi.api.pricing.PricingData
}
var file_pricing_v1_pricing_proto_depIdxs = []int32{
	8, // 0: blueapi.pricing.v1.GetPricingRequest.filters:type_name -> blueapi.pricing.v1.GetPricingRequest.FiltersEntry
	9, // 1: blueapi.pricing.v1.GetPricingResponse.pricingData:type_name -> blueapi.api.pricing.PricingData
	6, // 2: blueapi.pricing.v1.GetSupportedServicesResponse.supportedServices:type_name -> blueapi.pricing.v1.SupportedService
	7, // 3: blueapi.pricing.v1.SupportedService.attributes:type_name -> blueapi.pricing.v1.Attribute
	0, // 4: blueapi.pricing.v1.Pricing.GetInfo:input_type -> blueapi.pricing.v1.GetInfoRequest
	2, // 5: blueapi.pricing.v1.Pricing.GetPricing:input_type -> blueapi.pricing.v1.GetPricingRequest
	4, // 6: blueapi.pricing.v1.Pricing.GetSupportedServices:input_type -> blueapi.pricing.v1.GetSupportedServicesRequest
	1, // 7: blueapi.pricing.v1.Pricing.GetInfo:output_type -> blueapi.pricing.v1.GetInfoResponse
	3, // 8: blueapi.pricing.v1.Pricing.GetPricing:output_type -> blueapi.pricing.v1.GetPricingResponse
	5, // 9: blueapi.pricing.v1.Pricing.GetSupportedServices:output_type -> blueapi.pricing.v1.GetSupportedServicesResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
			switch v := v.(*GetPricingRequest); i {
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
			switch v := v.(*GetPricingResponse); i {
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
		file_pricing_v1_pricing_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetSupportedServicesRequest); i {
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
		file_pricing_v1_pricing_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetSupportedServicesResponse); i {
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
		file_pricing_v1_pricing_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SupportedService); i {
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
		file_pricing_v1_pricing_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Attribute); i {
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
			NumMessages:   9,
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
