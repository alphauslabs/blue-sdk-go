// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: api/aws/cost.proto

package aws

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

	Account        string                 `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	BillingGroupId string                 `protobuf:"bytes,2,opt,name=billingGroupId,proto3" json:"billingGroupId,omitempty"`
	Date           *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	ProductCode    string                 `protobuf:"bytes,4,opt,name=productCode,proto3" json:"productCode,omitempty"`
	ServiceCode    string                 `protobuf:"bytes,5,opt,name=serviceCode,proto3" json:"serviceCode,omitempty"`
	Region         string                 `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	Zone           string                 `protobuf:"bytes,7,opt,name=zone,proto3" json:"zone,omitempty"`
	UsageType      string                 `protobuf:"bytes,8,opt,name=usageType,proto3" json:"usageType,omitempty"`
	InstanceType   string                 `protobuf:"bytes,9,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	Operation      string                 `protobuf:"bytes,10,opt,name=operation,proto3" json:"operation,omitempty"`
	InvoiceId      string                 `protobuf:"bytes,11,opt,name=invoiceId,proto3" json:"invoiceId,omitempty"`
	Description    string                 `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	Tags           map[string]string      `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CostCategories map[string]string      `protobuf:"bytes,14,rep,name=costCategories,proto3" json:"costCategories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Usage          float64                `protobuf:"fixed64,15,opt,name=usage,proto3" json:"usage,omitempty"`
	Cost           float64                `protobuf:"fixed64,16,opt,name=cost,proto3" json:"cost,omitempty"`
	ResourceId     string                 `protobuf:"bytes,17,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	TagId          string                 `protobuf:"bytes,18,opt,name=tagId,proto3" json:"tagId,omitempty"`
}

func (x *Cost) Reset() {
	*x = Cost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_aws_cost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cost) ProtoMessage() {}

func (x *Cost) ProtoReflect() protoreflect.Message {
	mi := &file_api_aws_cost_proto_msgTypes[0]
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
	return file_api_aws_cost_proto_rawDescGZIP(), []int{0}
}

func (x *Cost) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Cost) GetBillingGroupId() string {
	if x != nil {
		return x.BillingGroupId
	}
	return ""
}

func (x *Cost) GetDate() *timestamppb.Timestamp {
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

func (x *Cost) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Cost) GetCostCategories() map[string]string {
	if x != nil {
		return x.CostCategories
	}
	return nil
}

func (x *Cost) GetUsage() float64 {
	if x != nil {
		return x.Usage
	}
	return 0
}

func (x *Cost) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Cost) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Cost) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

var File_api_aws_cost_proto protoreflect.FileDescriptor

var file_api_aws_cost_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x77, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x05, 0x0a, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77,
	0x73, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x51, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x63,
	0x6f, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x41, 0x0a, 0x13, 0x43, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x5c, 0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x77, 0x73, 0x42, 0x0f, 0x41, 0x70, 0x69, 0x41, 0x77, 0x73, 0x43, 0x6f,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_aws_cost_proto_rawDescOnce sync.Once
	file_api_aws_cost_proto_rawDescData = file_api_aws_cost_proto_rawDesc
)

func file_api_aws_cost_proto_rawDescGZIP() []byte {
	file_api_aws_cost_proto_rawDescOnce.Do(func() {
		file_api_aws_cost_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_aws_cost_proto_rawDescData)
	})
	return file_api_aws_cost_proto_rawDescData
}

var file_api_aws_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_aws_cost_proto_goTypes = []interface{}{
	(*Cost)(nil),                  // 0: blueapi.api.aws.Cost
	nil,                           // 1: blueapi.api.aws.Cost.TagsEntry
	nil,                           // 2: blueapi.api.aws.Cost.CostCategoriesEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_api_aws_cost_proto_depIdxs = []int32{
	3, // 0: blueapi.api.aws.Cost.date:type_name -> google.protobuf.Timestamp
	1, // 1: blueapi.api.aws.Cost.tags:type_name -> blueapi.api.aws.Cost.TagsEntry
	2, // 2: blueapi.api.aws.Cost.costCategories:type_name -> blueapi.api.aws.Cost.CostCategoriesEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_aws_cost_proto_init() }
func file_api_aws_cost_proto_init() {
	if File_api_aws_cost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_aws_cost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_api_aws_cost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_aws_cost_proto_goTypes,
		DependencyIndexes: file_api_aws_cost_proto_depIdxs,
		MessageInfos:      file_api_aws_cost_proto_msgTypes,
	}.Build()
	File_api_aws_cost_proto = out.File
	file_api_aws_cost_proto_rawDesc = nil
	file_api_aws_cost_proto_goTypes = nil
	file_api_aws_cost_proto_depIdxs = nil
}
