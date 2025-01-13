// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: api/cover/unitcost.proto

package cover

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

type UnitCostData struct {
	state           protoimpl.MessageState        `protogen:"open.v1"`
	Id              string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UnitName        string                        `protobuf:"bytes,2,opt,name=unitName,proto3" json:"unitName,omitempty"`
	Description     string                        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	UnitItems       []*UnitItem                   `protobuf:"bytes,4,rep,name=unitItems,proto3" json:"unitItems,omitempty"`
	SharedResources []*SharedResourcesCombination `protobuf:"bytes,5,rep,name=sharedResources,proto3" json:"sharedResources,omitempty"`
	CreatedBy       string                        `protobuf:"bytes,6,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	CreateTime      string                        `protobuf:"bytes,7,opt,name=createTime,proto3" json:"createTime,omitempty"`
	LastUpdatedAt   string                        `protobuf:"bytes,8,opt,name=lastUpdatedAt,proto3" json:"lastUpdatedAt,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UnitCostData) Reset() {
	*x = UnitCostData{}
	mi := &file_api_cover_unitcost_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnitCostData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitCostData) ProtoMessage() {}

func (x *UnitCostData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_unitcost_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitCostData.ProtoReflect.Descriptor instead.
func (*UnitCostData) Descriptor() ([]byte, []int) {
	return file_api_cover_unitcost_proto_rawDescGZIP(), []int{0}
}

func (x *UnitCostData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UnitCostData) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *UnitCostData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UnitCostData) GetUnitItems() []*UnitItem {
	if x != nil {
		return x.UnitItems
	}
	return nil
}

func (x *UnitCostData) GetSharedResources() []*SharedResourcesCombination {
	if x != nil {
		return x.SharedResources
	}
	return nil
}

func (x *UnitCostData) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *UnitCostData) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *UnitCostData) GetLastUpdatedAt() string {
	if x != nil {
		return x.LastUpdatedAt
	}
	return ""
}

type UnitItem struct {
	state                          protoimpl.MessageState `protogen:"open.v1"`
	ItemName                       string                 `protobuf:"bytes,1,opt,name=itemName,proto3" json:"itemName,omitempty"`
	Distribution                   float64                `protobuf:"fixed64,2,opt,name=distribution,proto3" json:"distribution,omitempty"`
	DedicatedResourcesCombinations *ResourcesCombinations `protobuf:"bytes,3,opt,name=dedicatedResourcesCombinations,proto3" json:"dedicatedResourcesCombinations,omitempty"`
	unknownFields                  protoimpl.UnknownFields
	sizeCache                      protoimpl.SizeCache
}

func (x *UnitItem) Reset() {
	*x = UnitItem{}
	mi := &file_api_cover_unitcost_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnitItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitItem) ProtoMessage() {}

func (x *UnitItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_unitcost_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitItem.ProtoReflect.Descriptor instead.
func (*UnitItem) Descriptor() ([]byte, []int) {
	return file_api_cover_unitcost_proto_rawDescGZIP(), []int{1}
}

func (x *UnitItem) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *UnitItem) GetDistribution() float64 {
	if x != nil {
		return x.Distribution
	}
	return 0
}

func (x *UnitItem) GetDedicatedResourcesCombinations() *ResourcesCombinations {
	if x != nil {
		return x.DedicatedResourcesCombinations
	}
	return nil
}

type SharedResourcesCombination struct {
	state                       protoimpl.MessageState `protogen:"open.v1"`
	CombinationName             string                 `protobuf:"bytes,1,opt,name=combinationName,proto3" json:"combinationName,omitempty"`
	SharedResourcesCombinations *ResourcesCombinations `protobuf:"bytes,2,opt,name=sharedResourcesCombinations,proto3" json:"sharedResourcesCombinations,omitempty"`
	// List the unit items and their corresponding percentages.
	Distribution  map[string]float64 `protobuf:"bytes,3,rep,name=distribution,proto3" json:"distribution,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SharedResourcesCombination) Reset() {
	*x = SharedResourcesCombination{}
	mi := &file_api_cover_unitcost_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SharedResourcesCombination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedResourcesCombination) ProtoMessage() {}

func (x *SharedResourcesCombination) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_unitcost_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedResourcesCombination.ProtoReflect.Descriptor instead.
func (*SharedResourcesCombination) Descriptor() ([]byte, []int) {
	return file_api_cover_unitcost_proto_rawDescGZIP(), []int{2}
}

func (x *SharedResourcesCombination) GetCombinationName() string {
	if x != nil {
		return x.CombinationName
	}
	return ""
}

func (x *SharedResourcesCombination) GetSharedResourcesCombinations() *ResourcesCombinations {
	if x != nil {
		return x.SharedResourcesCombinations
	}
	return nil
}

func (x *SharedResourcesCombination) GetDistribution() map[string]float64 {
	if x != nil {
		return x.Distribution
	}
	return nil
}

type ResourcesCombinations struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. Valid only for the `aws` vendor. AWS-specific options.
	AwsOptions *CostGroupAwsOptions `protobuf:"bytes,1,opt,name=awsOptions,proto3" json:"awsOptions,omitempty"`
	// Optional. Valid only for the `azure` vendor. Azure-specific options.
	AzureOptions *CostGroupAzureOptions `protobuf:"bytes,2,opt,name=azureOptions,proto3" json:"azureOptions,omitempty"`
	// Optional. Valid only for the `gcp` vendor. GCP-specific options.
	GcpOptions *CostGroupGcpOptions `protobuf:"bytes,3,opt,name=gcpOptions,proto3" json:"gcpOptions,omitempty"`
	// Optional. Valid only for the `azurecsp` vendor. AzureCSP-specific options.
	AzurecspOptions *CostGroupAzureCspOptions `protobuf:"bytes,4,opt,name=azurecspOptions,proto3" json:"azurecspOptions,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ResourcesCombinations) Reset() {
	*x = ResourcesCombinations{}
	mi := &file_api_cover_unitcost_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourcesCombinations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcesCombinations) ProtoMessage() {}

func (x *ResourcesCombinations) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_unitcost_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcesCombinations.ProtoReflect.Descriptor instead.
func (*ResourcesCombinations) Descriptor() ([]byte, []int) {
	return file_api_cover_unitcost_proto_rawDescGZIP(), []int{3}
}

func (x *ResourcesCombinations) GetAwsOptions() *CostGroupAwsOptions {
	if x != nil {
		return x.AwsOptions
	}
	return nil
}

func (x *ResourcesCombinations) GetAzureOptions() *CostGroupAzureOptions {
	if x != nil {
		return x.AzureOptions
	}
	return nil
}

func (x *ResourcesCombinations) GetGcpOptions() *CostGroupGcpOptions {
	if x != nil {
		return x.GcpOptions
	}
	return nil
}

func (x *ResourcesCombinations) GetAzurecspOptions() *CostGroupAzureCspOptions {
	if x != nil {
		return x.AzurecspOptions
	}
	return nil
}

var File_api_cover_unitcost_proto protoreflect.FileDescriptor

var file_api_cover_unitcost_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x75, 0x6e, 0x69, 0x74,
	0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x1a, 0x19, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x02, 0x0a, 0x0c, 0x55, 0x6e, 0x69,
	0x74, 0x43, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x55,
	0x6e, 0x69, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x57, 0x0a, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x43,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xbc, 0x01, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x70, 0x0a, 0x1e,
	0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x1e,
	0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xd8,
	0x02, 0x0a, 0x1a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x6a, 0x0a, 0x1b, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62,
	0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x1b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x63, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3f, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcc, 0x02, 0x0a, 0x15, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x0a, 0x61, 0x77, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x73, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x77, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0a, 0x61, 0x77, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c, 0x0a, 0x0c, 0x61,
	0x7a, 0x75, 0x72, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x7a, 0x75, 0x72, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x0a, 0x67, 0x63, 0x70,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x63, 0x70, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x67, 0x63, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x55, 0x0a, 0x0f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x63, 0x73, 0x70, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x73, 0x70,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x63, 0x73,
	0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x61, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x10, 0x41, 0x70, 0x69,
	0x55, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75,
	0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_cover_unitcost_proto_rawDescOnce sync.Once
	file_api_cover_unitcost_proto_rawDescData = file_api_cover_unitcost_proto_rawDesc
)

func file_api_cover_unitcost_proto_rawDescGZIP() []byte {
	file_api_cover_unitcost_proto_rawDescOnce.Do(func() {
		file_api_cover_unitcost_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_unitcost_proto_rawDescData)
	})
	return file_api_cover_unitcost_proto_rawDescData
}

var file_api_cover_unitcost_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_cover_unitcost_proto_goTypes = []any{
	(*UnitCostData)(nil),               // 0: blueapi.api.cover.UnitCostData
	(*UnitItem)(nil),                   // 1: blueapi.api.cover.UnitItem
	(*SharedResourcesCombination)(nil), // 2: blueapi.api.cover.SharedResourcesCombination
	(*ResourcesCombinations)(nil),      // 3: blueapi.api.cover.ResourcesCombinations
	nil,                                // 4: blueapi.api.cover.SharedResourcesCombination.DistributionEntry
	(*CostGroupAwsOptions)(nil),        // 5: blueapi.api.cover.CostGroupAwsOptions
	(*CostGroupAzureOptions)(nil),      // 6: blueapi.api.cover.CostGroupAzureOptions
	(*CostGroupGcpOptions)(nil),        // 7: blueapi.api.cover.CostGroupGcpOptions
	(*CostGroupAzureCspOptions)(nil),   // 8: blueapi.api.cover.CostGroupAzureCspOptions
}
var file_api_cover_unitcost_proto_depIdxs = []int32{
	1, // 0: blueapi.api.cover.UnitCostData.unitItems:type_name -> blueapi.api.cover.UnitItem
	2, // 1: blueapi.api.cover.UnitCostData.sharedResources:type_name -> blueapi.api.cover.SharedResourcesCombination
	3, // 2: blueapi.api.cover.UnitItem.dedicatedResourcesCombinations:type_name -> blueapi.api.cover.ResourcesCombinations
	3, // 3: blueapi.api.cover.SharedResourcesCombination.sharedResourcesCombinations:type_name -> blueapi.api.cover.ResourcesCombinations
	4, // 4: blueapi.api.cover.SharedResourcesCombination.distribution:type_name -> blueapi.api.cover.SharedResourcesCombination.DistributionEntry
	5, // 5: blueapi.api.cover.ResourcesCombinations.awsOptions:type_name -> blueapi.api.cover.CostGroupAwsOptions
	6, // 6: blueapi.api.cover.ResourcesCombinations.azureOptions:type_name -> blueapi.api.cover.CostGroupAzureOptions
	7, // 7: blueapi.api.cover.ResourcesCombinations.gcpOptions:type_name -> blueapi.api.cover.CostGroupGcpOptions
	8, // 8: blueapi.api.cover.ResourcesCombinations.azurecspOptions:type_name -> blueapi.api.cover.CostGroupAzureCspOptions
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_api_cover_unitcost_proto_init() }
func file_api_cover_unitcost_proto_init() {
	if File_api_cover_unitcost_proto != nil {
		return
	}
	file_api_cover_costgroup_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cover_unitcost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_unitcost_proto_goTypes,
		DependencyIndexes: file_api_cover_unitcost_proto_depIdxs,
		MessageInfos:      file_api_cover_unitcost_proto_msgTypes,
	}.Build()
	File_api_cover_unitcost_proto = out.File
	file_api_cover_unitcost_proto_rawDesc = nil
	file_api_cover_unitcost_proto_goTypes = nil
	file_api_cover_unitcost_proto_depIdxs = nil
}
