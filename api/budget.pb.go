// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: api/budget.proto

package api

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

type Budget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Format: yyyy
	Year string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	// key: mm, value: 00.00
	MonthlyBudget map[string]float64 `protobuf:"bytes,3,rep,name=monthlyBudget,proto3" json:"monthlyBudget,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *Budget) Reset() {
	*x = Budget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_budget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Budget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budget) ProtoMessage() {}

func (x *Budget) ProtoReflect() protoreflect.Message {
	mi := &file_api_budget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Budget.ProtoReflect.Descriptor instead.
func (*Budget) Descriptor() ([]byte, []int) {
	return file_api_budget_proto_rawDescGZIP(), []int{0}
}

func (x *Budget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Budget) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Budget) GetMonthlyBudget() map[string]float64 {
	if x != nil {
		return x.MonthlyBudget
	}
	return nil
}

var File_api_budget_proto protoreflect.FileDescriptor

var file_api_budget_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x22,
	0xbc, 0x01, 0x0a, 0x06, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x4c,
	0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x1a, 0x40, 0x0a, 0x12,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x53,
	0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0e, 0x41, 0x70, 0x69,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_budget_proto_rawDescOnce sync.Once
	file_api_budget_proto_rawDescData = file_api_budget_proto_rawDesc
)

func file_api_budget_proto_rawDescGZIP() []byte {
	file_api_budget_proto_rawDescOnce.Do(func() {
		file_api_budget_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_budget_proto_rawDescData)
	})
	return file_api_budget_proto_rawDescData
}

var file_api_budget_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_budget_proto_goTypes = []interface{}{
	(*Budget)(nil), // 0: blueapi.api.Budget
	nil,            // 1: blueapi.api.Budget.MonthlyBudgetEntry
}
var file_api_budget_proto_depIdxs = []int32{
	1, // 0: blueapi.api.Budget.monthlyBudget:type_name -> blueapi.api.Budget.MonthlyBudgetEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_budget_proto_init() }
func file_api_budget_proto_init() {
	if File_api_budget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_budget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Budget); i {
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
			RawDescriptor: file_api_budget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_budget_proto_goTypes,
		DependencyIndexes: file_api_budget_proto_depIdxs,
		MessageInfos:      file_api_budget_proto_msgTypes,
	}.Build()
	File_api_budget_proto = out.File
	file_api_budget_proto_rawDesc = nil
	file_api_budget_proto_goTypes = nil
	file_api_budget_proto_depIdxs = nil
}
