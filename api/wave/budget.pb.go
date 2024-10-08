// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/wave/budget.proto

package wave

import (
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

// Budget resource definition.
type BudgetAlert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// account id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// notification setting.
	Notification []*Notification `protobuf:"bytes,2,rep,name=notification,proto3" json:"notification,omitempty"`
	// budget setting.
	Budget []*Budget `protobuf:"bytes,3,rep,name=budget,proto3" json:"budget,omitempty"`
}

func (x *BudgetAlert) Reset() {
	*x = BudgetAlert{}
	mi := &file_api_wave_budget_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BudgetAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetAlert) ProtoMessage() {}

func (x *BudgetAlert) ProtoReflect() protoreflect.Message {
	mi := &file_api_wave_budget_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetAlert.ProtoReflect.Descriptor instead.
func (*BudgetAlert) Descriptor() ([]byte, []int) {
	return file_api_wave_budget_proto_rawDescGZIP(), []int{0}
}

func (x *BudgetAlert) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BudgetAlert) GetNotification() []*Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

func (x *BudgetAlert) GetBudget() []*Budget {
	if x != nil {
		return x.Budget
	}
	return nil
}

// Notification resource definition.
type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// notification id
	// `email` / `slack`
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// destination email address /slack webhook url
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// notification enable/disable
	Enabled bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_api_wave_budget_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_api_wave_budget_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_api_wave_budget_proto_rawDescGZIP(), []int{1}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Notification) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// Budget resource definition.
type Budget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// budget id
	// `previousDay` / `daily` / `monthly`
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// budget value
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	// budget setting enable/disable
	Enabled bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Budget) Reset() {
	*x = Budget{}
	mi := &file_api_wave_budget_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Budget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budget) ProtoMessage() {}

func (x *Budget) ProtoReflect() protoreflect.Message {
	mi := &file_api_wave_budget_proto_msgTypes[2]
	if x != nil {
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
	return file_api_wave_budget_proto_rawDescGZIP(), []int{2}
}

func (x *Budget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Budget) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Budget) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_api_wave_budget_proto protoreflect.FileDescriptor

var file_api_wave_budget_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x76, 0x65, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x0b, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x42, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x06,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x5a, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x22, 0x48, 0x0a, 0x06, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x61, 0x0a, 0x1e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x42, 0x12,
	0x41, 0x70, 0x69, 0x57, 0x61, 0x76, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d,
	0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x76, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_wave_budget_proto_rawDescOnce sync.Once
	file_api_wave_budget_proto_rawDescData = file_api_wave_budget_proto_rawDesc
)

func file_api_wave_budget_proto_rawDescGZIP() []byte {
	file_api_wave_budget_proto_rawDescOnce.Do(func() {
		file_api_wave_budget_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_wave_budget_proto_rawDescData)
	})
	return file_api_wave_budget_proto_rawDescData
}

var file_api_wave_budget_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_wave_budget_proto_goTypes = []any{
	(*BudgetAlert)(nil),  // 0: blueapi.api.wave.BudgetAlert
	(*Notification)(nil), // 1: blueapi.api.wave.Notification
	(*Budget)(nil),       // 2: blueapi.api.wave.Budget
}
var file_api_wave_budget_proto_depIdxs = []int32{
	1, // 0: blueapi.api.wave.BudgetAlert.notification:type_name -> blueapi.api.wave.Notification
	2, // 1: blueapi.api.wave.BudgetAlert.budget:type_name -> blueapi.api.wave.Budget
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_wave_budget_proto_init() }
func file_api_wave_budget_proto_init() {
	if File_api_wave_budget_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_wave_budget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_wave_budget_proto_goTypes,
		DependencyIndexes: file_api_wave_budget_proto_depIdxs,
		MessageInfos:      file_api_wave_budget_proto_msgTypes,
	}.Build()
	File_api_wave_budget_proto = out.File
	file_api_wave_budget_proto_rawDesc = nil
	file_api_wave_budget_proto_goTypes = nil
	file_api_wave_budget_proto_depIdxs = nil
}
