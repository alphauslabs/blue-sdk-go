// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/azureea/cost.proto

package azureea

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

type Cost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName      string  `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	ServiceTier      string  `protobuf:"bytes,2,opt,name=serviceTier,proto3" json:"serviceTier,omitempty"`
	Location         string  `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	PartNumber       string  `protobuf:"bytes,4,opt,name=partNumber,proto3" json:"partNumber,omitempty"`
	OfferId          string  `protobuf:"bytes,5,opt,name=offerId,proto3" json:"offerId,omitempty"`
	Cost             float64 `protobuf:"fixed64,6,opt,name=cost,proto3" json:"cost,omitempty"`
	AccountId        string  `protobuf:"bytes,7,opt,name=accountId,proto3" json:"accountId,omitempty"`
	AccountName      string  `protobuf:"bytes,8,opt,name=accountName,proto3" json:"accountName,omitempty"`
	SubscriptionGuid string  `protobuf:"bytes,9,opt,name=subscriptionGuid,proto3" json:"subscriptionGuid,omitempty"`
	SubscriptionName string  `protobuf:"bytes,10,opt,name=subscriptionName,proto3" json:"subscriptionName,omitempty"`
	Date             string  `protobuf:"bytes,11,opt,name=date,proto3" json:"date,omitempty"`
	MeterCategory    string  `protobuf:"bytes,12,opt,name=meterCategory,proto3" json:"meterCategory,omitempty"`
	MeterSubCategory string  `protobuf:"bytes,13,opt,name=meterSubCategory,proto3" json:"meterSubCategory,omitempty"`
	MeterName        string  `protobuf:"bytes,14,opt,name=meterName,proto3" json:"meterName,omitempty"`
	ConsumedQuantity string  `protobuf:"bytes,15,opt,name=consumedQuantity,proto3" json:"consumedQuantity,omitempty"`
	ConsumedService  string  `protobuf:"bytes,16,opt,name=consumedService,proto3" json:"consumedService,omitempty"`
	AdditionalInfo   string  `protobuf:"bytes,17,opt,name=additionalInfo,proto3" json:"additionalInfo,omitempty"`
	Tags             string  `protobuf:"bytes,18,opt,name=tags,proto3" json:"tags,omitempty"`
	DepartmentName   string  `protobuf:"bytes,19,opt,name=departmentName,proto3" json:"departmentName,omitempty"`
	CostCenter       string  `protobuf:"bytes,20,opt,name=costCenter,proto3" json:"costCenter,omitempty"`
	UnitOfMeasure    string  `protobuf:"bytes,21,opt,name=unitOfMeasure,proto3" json:"unitOfMeasure,omitempty"`
	ResourceGroup    string  `protobuf:"bytes,22,opt,name=resourceGroup,proto3" json:"resourceGroup,omitempty"`
	Enrollment       string  `protobuf:"bytes,23,opt,name=enrollment,proto3" json:"enrollment,omitempty"`
}

func (x *Cost) Reset() {
	*x = Cost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_azureea_cost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cost) ProtoMessage() {}

func (x *Cost) ProtoReflect() protoreflect.Message {
	mi := &file_api_azureea_cost_proto_msgTypes[0]
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
	return file_api_azureea_cost_proto_rawDescGZIP(), []int{0}
}

func (x *Cost) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Cost) GetServiceTier() string {
	if x != nil {
		return x.ServiceTier
	}
	return ""
}

func (x *Cost) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Cost) GetPartNumber() string {
	if x != nil {
		return x.PartNumber
	}
	return ""
}

func (x *Cost) GetOfferId() string {
	if x != nil {
		return x.OfferId
	}
	return ""
}

func (x *Cost) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Cost) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Cost) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *Cost) GetSubscriptionGuid() string {
	if x != nil {
		return x.SubscriptionGuid
	}
	return ""
}

func (x *Cost) GetSubscriptionName() string {
	if x != nil {
		return x.SubscriptionName
	}
	return ""
}

func (x *Cost) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Cost) GetMeterCategory() string {
	if x != nil {
		return x.MeterCategory
	}
	return ""
}

func (x *Cost) GetMeterSubCategory() string {
	if x != nil {
		return x.MeterSubCategory
	}
	return ""
}

func (x *Cost) GetMeterName() string {
	if x != nil {
		return x.MeterName
	}
	return ""
}

func (x *Cost) GetConsumedQuantity() string {
	if x != nil {
		return x.ConsumedQuantity
	}
	return ""
}

func (x *Cost) GetConsumedService() string {
	if x != nil {
		return x.ConsumedService
	}
	return ""
}

func (x *Cost) GetAdditionalInfo() string {
	if x != nil {
		return x.AdditionalInfo
	}
	return ""
}

func (x *Cost) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Cost) GetDepartmentName() string {
	if x != nil {
		return x.DepartmentName
	}
	return ""
}

func (x *Cost) GetCostCenter() string {
	if x != nil {
		return x.CostCenter
	}
	return ""
}

func (x *Cost) GetUnitOfMeasure() string {
	if x != nil {
		return x.UnitOfMeasure
	}
	return ""
}

func (x *Cost) GetResourceGroup() string {
	if x != nil {
		return x.ResourceGroup
	}
	return ""
}

func (x *Cost) GetEnrollment() string {
	if x != nil {
		return x.Enrollment
	}
	return ""
}

var File_api_azureea_cost_proto protoreflect.FileDescriptor

var file_api_azureea_cost_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x65, 0x61, 0x2f, 0x63, 0x6f,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x65, 0x61, 0x22, 0x96, 0x06,
	0x0a, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64,
	0x12, 0x2a, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x53,
	0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x64, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x0f,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x73, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x73, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x6e,
	0x69, 0x74, 0x4f, 0x66, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x4f, 0x66, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x68, 0x0a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x65, 0x61, 0x42, 0x13, 0x41, 0x70, 0x69,
	0x41, 0x7a, 0x75, 0x72, 0x65, 0x45, 0x61, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x65, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_azureea_cost_proto_rawDescOnce sync.Once
	file_api_azureea_cost_proto_rawDescData = file_api_azureea_cost_proto_rawDesc
)

func file_api_azureea_cost_proto_rawDescGZIP() []byte {
	file_api_azureea_cost_proto_rawDescOnce.Do(func() {
		file_api_azureea_cost_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_azureea_cost_proto_rawDescData)
	})
	return file_api_azureea_cost_proto_rawDescData
}

var file_api_azureea_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_azureea_cost_proto_goTypes = []interface{}{
	(*Cost)(nil), // 0: blueapi.api.azureea.Cost
}
var file_api_azureea_cost_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_azureea_cost_proto_init() }
func file_api_azureea_cost_proto_init() {
	if File_api_azureea_cost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_azureea_cost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_azureea_cost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_azureea_cost_proto_goTypes,
		DependencyIndexes: file_api_azureea_cost_proto_depIdxs,
		MessageInfos:      file_api_azureea_cost_proto_msgTypes,
	}.Build()
	File_api_azureea_cost_proto = out.File
	file_api_azureea_cost_proto_rawDesc = nil
	file_api_azureea_cost_proto_goTypes = nil
	file_api_azureea_cost_proto_depIdxs = nil
}
