// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/adjustment.proto

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

// AdjustmentConfig resource definition.
type AdjustmentConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User configuration
	Config []*ConfigFilters `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	// Vendor
	Vendor string `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Management account configuration
	Accounts []*ManagementAccount `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *AdjustmentConfig) Reset() {
	*x = AdjustmentConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_adjustment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdjustmentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustmentConfig) ProtoMessage() {}

func (x *AdjustmentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_adjustment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjustmentConfig.ProtoReflect.Descriptor instead.
func (*AdjustmentConfig) Descriptor() ([]byte, []int) {
	return file_api_adjustment_proto_rawDescGZIP(), []int{0}
}

func (x *AdjustmentConfig) GetConfig() []*ConfigFilters {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *AdjustmentConfig) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *AdjustmentConfig) GetAccounts() []*ManagementAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

// ManagementAccount resource definition.
type ManagementAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Vendor-managed organization account Id
	ManagementAccountId string `protobuf:"bytes,1,opt,name=managementAccountId,proto3" json:"managementAccountId,omitempty"`
	// A list of filtering options. See [ConfigFilters] for more information on each filter item. Multiple filter items will use the logical 'or' operator, e.g. filter1 || filter2 || filter3, etc.
	Config []*ConfigFilters `protobuf:"bytes,2,rep,name=config,proto3" json:"config,omitempty"`
}

func (x *ManagementAccount) Reset() {
	*x = ManagementAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_adjustment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagementAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagementAccount) ProtoMessage() {}

func (x *ManagementAccount) ProtoReflect() protoreflect.Message {
	mi := &file_api_adjustment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagementAccount.ProtoReflect.Descriptor instead.
func (*ManagementAccount) Descriptor() ([]byte, []int) {
	return file_api_adjustment_proto_rawDescGZIP(), []int{1}
}

func (x *ManagementAccount) GetManagementAccountId() string {
	if x != nil {
		return x.ManagementAccountId
	}
	return ""
}

func (x *ManagementAccount) GetConfig() []*ConfigFilters {
	if x != nil {
		return x.Config
	}
	return nil
}

// ConfigFilters resource definition.
// A map of "key:value" config filters. The key indicates the adjustment key while the value is the filter adjustment value which can be prefixed by either "eq:" (equal), "re:" (regular expressions based on https://github.com/google/re2), or "!re:" (reverse "re:"). No prefix is the same as "eq:". Multiple map items will use the logical 'and' operator, e.g. mapfilter1 && mapfilter2 && mapfilter3, etc.
//
// For example, if you want to query lineitems with the adjustment `productCode:AmazonEC2`, set to `{"productCode":"AmazonEC2"}`. You can also use regular expressions for adjustment values, such as `{"description":"re:[A-Za-z0-9]*"}`.
// List of available adjustment keys: productCode, type, description
// For example value on productCode: AmazonEC2, AmazonRDS, AWSLambda, etc.
// For example value on type: Fee, Refund, SppDiscount, etc.
type ConfigFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AndFilters map[string]string `protobuf:"bytes,1,rep,name=andFilters,proto3" json:"andFilters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConfigFilters) Reset() {
	*x = ConfigFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_adjustment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigFilters) ProtoMessage() {}

func (x *ConfigFilters) ProtoReflect() protoreflect.Message {
	mi := &file_api_adjustment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigFilters.ProtoReflect.Descriptor instead.
func (*ConfigFilters) Descriptor() ([]byte, []int) {
	return file_api_adjustment_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigFilters) GetAndFilters() map[string]string {
	if x != nil {
		return x.AndFilters
	}
	return nil
}

var File_api_adjustment_proto protoreflect.FileDescriptor

var file_api_adjustment_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x22, 0x79, 0x0a, 0x11, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x9a, 0x01, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4a, 0x0a,
	0x0a, 0x61, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x6e,
	0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61,
	0x6e, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x6e, 0x64,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x57, 0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x12, 0x41, 0x70, 0x69, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_adjustment_proto_rawDescOnce sync.Once
	file_api_adjustment_proto_rawDescData = file_api_adjustment_proto_rawDesc
)

func file_api_adjustment_proto_rawDescGZIP() []byte {
	file_api_adjustment_proto_rawDescOnce.Do(func() {
		file_api_adjustment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_adjustment_proto_rawDescData)
	})
	return file_api_adjustment_proto_rawDescData
}

var file_api_adjustment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_adjustment_proto_goTypes = []interface{}{
	(*AdjustmentConfig)(nil),  // 0: blueapi.api.AdjustmentConfig
	(*ManagementAccount)(nil), // 1: blueapi.api.ManagementAccount
	(*ConfigFilters)(nil),     // 2: blueapi.api.ConfigFilters
	nil,                       // 3: blueapi.api.ConfigFilters.AndFiltersEntry
}
var file_api_adjustment_proto_depIdxs = []int32{
	2, // 0: blueapi.api.AdjustmentConfig.config:type_name -> blueapi.api.ConfigFilters
	1, // 1: blueapi.api.AdjustmentConfig.accounts:type_name -> blueapi.api.ManagementAccount
	2, // 2: blueapi.api.ManagementAccount.config:type_name -> blueapi.api.ConfigFilters
	3, // 3: blueapi.api.ConfigFilters.andFilters:type_name -> blueapi.api.ConfigFilters.AndFiltersEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_adjustment_proto_init() }
func file_api_adjustment_proto_init() {
	if File_api_adjustment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_adjustment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdjustmentConfig); i {
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
		file_api_adjustment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagementAccount); i {
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
		file_api_adjustment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigFilters); i {
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
			RawDescriptor: file_api_adjustment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_adjustment_proto_goTypes,
		DependencyIndexes: file_api_adjustment_proto_depIdxs,
		MessageInfos:      file_api_adjustment_proto_msgTypes,
	}.Build()
	File_api_adjustment_proto = out.File
	file_api_adjustment_proto_rawDesc = nil
	file_api_adjustment_proto_goTypes = nil
	file_api_adjustment_proto_depIdxs = nil
}
