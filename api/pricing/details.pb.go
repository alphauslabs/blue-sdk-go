// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: api/pricing/details.proto

package pricing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Pricing details
type PricingData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Cloud vendor.
	Vendor string `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Cloud vendor service. We have a limited number of services supported currently.
	// Refer to https://labs.alphaus.cloud/blueapidocs/#/Pricing/Pricing_GetSupportedServices to list supported services.
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// Region code.
	RegionCode string `protobuf:"bytes,3,opt,name=regionCode,proto3" json:"regionCode,omitempty"`
	// SKU ID.
	Sku string `protobuf:"bytes,4,opt,name=sku,proto3" json:"sku,omitempty"`
	// Unit of measure.
	Unit string `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	// Price per unit.
	PricePerUnit float64 `protobuf:"fixed64,6,opt,name=pricePerUnit,proto3" json:"pricePerUnit,omitempty"`
	// Service details. A protobuf struct which translates to a map for HTTP. Keys are of type string and values can be of type string, float, or bool, depending on the detail.
	//
	// The following is an example of a serviceDetails field of a response,
	// ```
	//
	//	"serviceDetails": {
	//	  "availabilityZone": "NA",
	//	  "capacityStatus": "Used",
	//	  "classicNetworkingSupport": "false",
	//	  "clockSpeed": "3.5 GHz",
	//	  "currency": "USD",
	//	  "currentGeneration": "Yes",
	//	  "dedicatedEbsThroughput": "Up to 10000 Mbps",
	//	  "ebsOptimized": "",
	//	  "ecu": "NA",
	//	  "effectiveDate": "2022-04-01",
	//	  "elasticGraphicsType": "",
	//	  "endingRange": "Inf",
	//	  "enhancedNetworkingSupported": "Yes",
	//	  "fromLocation": "",
	//	  "fromLocationType": "",
	//	  "fromRegionCode": "",
	//	  "gpu": "",
	//	  "gpuMemory": "NA",
	//	  "groupDescription": "",
	//	  "groupings": "",
	//	  "instance": "",
	//	  "instanceCapacity10Xlarge": "",
	//	  "instanceCapacity12Xlarge": "",
	//	  "instanceCapacity16Xlarge": "",
	//	  "instanceCapacity18Xlarge": "",
	//	  "instanceCapacity24Xlarge": "",
	//	  "instanceCapacity2Xlarge": "",
	//	  "instanceCapacity32Xlarge": "",
	//	  "instanceCapacity4Xlarge": "",
	//	  "instanceCapacity8Xlarge": "",
	//	  "instanceCapacity9Xlarge": "",
	//	  "instanceCapacityLarge": "",
	//	  "instanceCapacityMedium": "",
	//	  "instanceCapacityMetal": "",
	//	  "instanceCapacityXlarge": "",
	//	  "instanceFamily": "Compute optimized",
	//	  "instanceSku": "",
	//	  "instanceType": "c6i.large",
	//	  "intelAvx2Available": "Yes",
	//	  "intelAvxAvailable": "Yes",
	//	  "intelTurboAvailable": "Yes",
	//	  "leaseContractLength": "1yr",
	//	  "licenseModel": "No License required",
	//	  "location": "Asia Pacific (Tokyo)",
	//	  "locationType": "AWS Region",
	//	  "marketOption": "OnDemand",
	//	  "maxIopsBurstPerformance": "",
	//	  "maxIopsVolume": "",
	//	  "maxThroughputVolume": "",
	//	  "maxVolumeSize": "",
	//	  "memory": "4 GiB",
	//	  "networkPerformance": "Up to 12500 Megabit",
	//	  "normalizationSizeFactor": "4",
	//	  "offerTermCode": "4NA7Y494T4",
	//	  "offeringClass": "standard",
	//	  "operatingSystem": "Windows",
	//	  "operation": "RunInstances:0202",
	//	  "physicalCores": "",
	//	  "physicalProcessor": "Intel Xeon 8375C (Ice Lake)",
	//	  "preInstalledSw": "SQL Web",
	//	  "priceDescription": "Windows with SQL Server Web (Amazon VPC), c6i.large reserved instance applied",
	//	  "processorArchitecture": "64-bit",
	//	  "processorFeatures": "Intel AVX; Intel AVX2; Intel AVX512; Intel Turbo",
	//	  "productFamily": "Compute Instance",
	//	  "productType": "",
	//	  "provisioned": "",
	//	  "purchaseOption": "No Upfront",
	//	  "rateCode": "2223B6PCG6QAUYY6.4NA7Y494T4.6YS6EN2CT7",
	//	  "relatedTo": "",
	//	  "resourceType": "",
	//	  "serviceCode": "AmazonEC2",
	//	  "serviceName": "Amazon Elastic Compute Cloud",
	//	  "snapShotArchiveFeeType": "",
	//	  "startingRange": "0",
	//	  "storage": "EBS only",
	//	  "storageMedia": "",
	//	  "tenancy": "Dedicated",
	//	  "termType": "Reserved",
	//	  "toLocation": "",
	//	  "toLocationType": "",
	//	  "toRegionCode": "",
	//	  "transferType": "",
	//	  "usageType": "APN1-DedicatedUsage:c6i.large",
	//	  "vcpu": "2",
	//	  "volumeApiName": "",
	//	  "volumeType": "",
	//	  "vpcNetworkingSupport": "true"
	//	}
	//
	// ```
	ServiceDetails *structpb.Struct `protobuf:"bytes,7,opt,name=serviceDetails,proto3" json:"serviceDetails,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PricingData) Reset() {
	*x = PricingData{}
	mi := &file_api_pricing_details_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PricingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricingData) ProtoMessage() {}

func (x *PricingData) ProtoReflect() protoreflect.Message {
	mi := &file_api_pricing_details_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricingData.ProtoReflect.Descriptor instead.
func (*PricingData) Descriptor() ([]byte, []int) {
	return file_api_pricing_details_proto_rawDescGZIP(), []int{0}
}

func (x *PricingData) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *PricingData) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *PricingData) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *PricingData) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *PricingData) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *PricingData) GetPricePerUnit() float64 {
	if x != nil {
		return x.PricePerUnit
	}
	return 0
}

func (x *PricingData) GetServiceDetails() *structpb.Struct {
	if x != nil {
		return x.ServiceDetails
	}
	return nil
}

var File_api_pricing_details_proto protoreflect.FileDescriptor

var file_api_pricing_details_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea,
	0x01, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50,
	0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x6b, 0x0a, 0x21, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x42, 0x16, 0x41, 0x70, 0x69, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_pricing_details_proto_rawDescOnce sync.Once
	file_api_pricing_details_proto_rawDescData = file_api_pricing_details_proto_rawDesc
)

func file_api_pricing_details_proto_rawDescGZIP() []byte {
	file_api_pricing_details_proto_rawDescOnce.Do(func() {
		file_api_pricing_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_pricing_details_proto_rawDescData)
	})
	return file_api_pricing_details_proto_rawDescData
}

var file_api_pricing_details_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_pricing_details_proto_goTypes = []any{
	(*PricingData)(nil),     // 0: blueapi.api.pricing.PricingData
	(*structpb.Struct)(nil), // 1: google.protobuf.Struct
}
var file_api_pricing_details_proto_depIdxs = []int32{
	1, // 0: blueapi.api.pricing.PricingData.serviceDetails:type_name -> google.protobuf.Struct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_pricing_details_proto_init() }
func file_api_pricing_details_proto_init() {
	if File_api_pricing_details_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_pricing_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_pricing_details_proto_goTypes,
		DependencyIndexes: file_api_pricing_details_proto_depIdxs,
		MessageInfos:      file_api_pricing_details_proto_msgTypes,
	}.Build()
	File_api_pricing_details_proto = out.File
	file_api_pricing_details_proto_rawDesc = nil
	file_api_pricing_details_proto_goTypes = nil
	file_api_pricing_details_proto_depIdxs = nil
}
