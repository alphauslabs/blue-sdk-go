// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for the Pricing.GetInfo rpc.
type GetInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	mi := &file_pricing_v1_pricing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[0]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      string                 `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	mi := &file_pricing_v1_pricing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[1]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
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
	Filters map[string]string `protobuf:"bytes,6,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Optional. Only specified columns will be returned, if provided.
	// All columns will be returned if this array is empty.
	// Supported columns can be listed using `/{vendor}/services` endpoint. For usage information visit https://labs.alphaus.cloud/blueapidocs/#/Pricing/Pricing_GetSupportedServices.
	Columns       []string `protobuf:"bytes,5,rep,name=columns,proto3" json:"columns,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPricingRequest) Reset() {
	*x = GetPricingRequest{}
	mi := &file_pricing_v1_pricing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPricingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPricingRequest) ProtoMessage() {}

func (x *GetPricingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[2]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Use token to retrieve next set of pricing items. An empty string means there are no more items to retrieve.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Array of pricing items details. Maximum number of items returned per call is 1000.
	PricingData   []*pricing.PricingData `protobuf:"bytes,2,rep,name=pricingData,proto3" json:"pricingData,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPricingResponse) Reset() {
	*x = GetPricingResponse{}
	mi := &file_pricing_v1_pricing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPricingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPricingResponse) ProtoMessage() {}

func (x *GetPricingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[3]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Cloud vendor, only `aws` and `azure` are currently supported.
	Vendor        string `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSupportedServicesRequest) Reset() {
	*x = GetSupportedServicesRequest{}
	mi := &file_pricing_v1_pricing_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSupportedServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedServicesRequest) ProtoMessage() {}

func (x *GetSupportedServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[4]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Suported services, regions, and attributes that can be used to specify which pricing data to retrieve from `/{vendor}/pricing`.
	// For usage information, visit https://labs.alphaus.cloud/blueapidocs/#/Pricing/Pricing_GetPricing.
	SupportedServices []*SupportedService `protobuf:"bytes,1,rep,name=supportedServices,proto3" json:"supportedServices,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetSupportedServicesResponse) Reset() {
	*x = GetSupportedServicesResponse{}
	mi := &file_pricing_v1_pricing_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSupportedServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedServicesResponse) ProtoMessage() {}

func (x *GetSupportedServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[5]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// AWS or Azure services only as of now.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// Array of regions supported for the specific service.
	Regions []string `protobuf:"bytes,2,rep,name=regions,proto3" json:"regions,omitempty"`
	// Array of attributes that can be used as key-value pairs for filtering.
	Attributes []*Attribute `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// Array of column names that can be used to specify what columns should `/{vendor}/pricing` return.
	// For usage information, visit https://labs.alphaus.cloud/blueapidocs/#/Pricing/Pricing_GetPricing.
	Columns       []string `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SupportedService) Reset() {
	*x = SupportedService{}
	mi := &file_pricing_v1_pricing_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SupportedService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportedService) ProtoMessage() {}

func (x *SupportedService) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[6]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Filter key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Array of filter values.
	Values        []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	mi := &file_pricing_v1_pricing_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_v1_pricing_proto_msgTypes[7]
	if x != nil {
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

const file_pricing_v1_pricing_proto_rawDesc = "" +
	"\n" +
	"\x18pricing/v1/pricing.proto\x12\x12blueapi.pricing.v1\x1a\x19api/pricing/details.proto\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x10\n" +
	"\x0eGetInfoRequest\"-\n" +
	"\x0fGetInfoResponse\x12\x1a\n" +
	"\bresponse\x18\x01 \x01(\tR\bresponse\"\x97\x02\n" +
	"\x11GetPricingRequest\x12\x16\n" +
	"\x06vendor\x18\x01 \x01(\tR\x06vendor\x12\x18\n" +
	"\aservice\x18\x02 \x01(\tR\aservice\x12\x16\n" +
	"\x06region\x18\x03 \x01(\tR\x06region\x12\x14\n" +
	"\x05token\x18\x04 \x01(\tR\x05token\x12L\n" +
	"\afilters\x18\x06 \x03(\v22.blueapi.pricing.v1.GetPricingRequest.FiltersEntryR\afilters\x12\x18\n" +
	"\acolumns\x18\x05 \x03(\tR\acolumns\x1a:\n" +
	"\fFiltersEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"n\n" +
	"\x12GetPricingResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12B\n" +
	"\vpricingData\x18\x02 \x03(\v2 .blueapi.api.pricing.PricingDataR\vpricingData\"5\n" +
	"\x1bGetSupportedServicesRequest\x12\x16\n" +
	"\x06vendor\x18\x01 \x01(\tR\x06vendor\"r\n" +
	"\x1cGetSupportedServicesResponse\x12R\n" +
	"\x11supportedServices\x18\x01 \x03(\v2$.blueapi.pricing.v1.SupportedServiceR\x11supportedServices\"\x9f\x01\n" +
	"\x10SupportedService\x12\x18\n" +
	"\aservice\x18\x01 \x01(\tR\aservice\x12\x18\n" +
	"\aregions\x18\x02 \x03(\tR\aregions\x12=\n" +
	"\n" +
	"attributes\x18\x03 \x03(\v2\x1d.blueapi.pricing.v1.AttributeR\n" +
	"attributes\x12\x18\n" +
	"\acolumns\x18\x04 \x03(\tR\acolumns\"5\n" +
	"\tAttribute\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x16\n" +
	"\x06values\x18\x02 \x03(\tR\x06values2\xa8\x04\n" +
	"\aPricing\x12d\n" +
	"\aGetInfo\x12\".blueapi.pricing.v1.GetInfoRequest\x1a#.blueapi.pricing.v1.GetInfoResponse\"\x10\x82\xd3\xe4\x93\x02\n" +
	"\x12\b/v0/info\x12|\n" +
	"\n" +
	"GetPricing\x12%.blueapi.pricing.v1.GetPricingRequest\x1a&.blueapi.pricing.v1.GetPricingResponse\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/v0/{vendor}/pricing\x12\x98\x01\n" +
	"\x14GetSupportedServices\x12/.blueapi.pricing.v1.GetSupportedServicesRequest\x1a0.blueapi.pricing.v1.GetSupportedServicesResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v0/{vendor}/services\x1a\x9d\x01\x92A\x99\x01\x12F(BETA) Pricing API. Base URL: https://api.alphaus.cloud/m/blue/pricing\x1aO\n" +
	"\x12Service definition\x129https://github.com/alphauslabs/blueapi/tree/main/pricing/BQ\n" +
	"\x19cloud.alphaus.api.pricingB\fPricingProtoZ&github.com/alphauslabs/blueapi/pricingb\x06proto3"

var (
	file_pricing_v1_pricing_proto_rawDescOnce sync.Once
	file_pricing_v1_pricing_proto_rawDescData []byte
)

func file_pricing_v1_pricing_proto_rawDescGZIP() []byte {
	file_pricing_v1_pricing_proto_rawDescOnce.Do(func() {
		file_pricing_v1_pricing_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pricing_v1_pricing_proto_rawDesc), len(file_pricing_v1_pricing_proto_rawDesc)))
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pricing_v1_pricing_proto_rawDesc), len(file_pricing_v1_pricing_proto_rawDesc)),
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
	file_pricing_v1_pricing_proto_goTypes = nil
	file_pricing_v1_pricing_proto_depIdxs = nil
}
