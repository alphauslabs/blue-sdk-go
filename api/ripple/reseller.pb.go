// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: api/ripple/reseller.proto

package ripple

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

type Reseller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reseller id. Generated automatically.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// email
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// billing internal id
	BillingInternalId string `protobuf:"bytes,3,opt,name=billingInternalId,proto3" json:"billingInternalId,omitempty"`
	// billing group id
	BillingGroupId string `protobuf:"bytes,4,opt,name=billingGroupId,proto3" json:"billingGroupId,omitempty"`
	// billing group name
	BillingGroupName string `protobuf:"bytes,5,opt,name=billingGroupName,proto3" json:"billingGroupName,omitempty"`
	// group type
	// Refer to the following for available values
	// billing_group: Billing Group.
	// access_group: Access Group.
	GroupType string `protobuf:"bytes,10,opt,name=groupType,proto3" json:"groupType,omitempty"`
	// wave status
	WaveStatus string `protobuf:"bytes,6,opt,name=waveStatus,proto3" json:"waveStatus,omitempty"`
	// wave feature config
	WaveConfig []*ResellerConfig `protobuf:"bytes,7,rep,name=waveConfig,proto3" json:"waveConfig,omitempty"`
	// wave pro feature config
	WaveProConfig []*ResellerConfig `protobuf:"bytes,8,rep,name=waveProConfig,proto3" json:"waveProConfig,omitempty"`
	// aqua feature config
	AquaConfig []*ResellerConfig `protobuf:"bytes,9,rep,name=aquaConfig,proto3" json:"aquaConfig,omitempty"`
}

func (x *Reseller) Reset() {
	*x = Reseller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ripple_reseller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reseller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reseller) ProtoMessage() {}

func (x *Reseller) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_reseller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reseller.ProtoReflect.Descriptor instead.
func (*Reseller) Descriptor() ([]byte, []int) {
	return file_api_ripple_reseller_proto_rawDescGZIP(), []int{0}
}

func (x *Reseller) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reseller) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Reseller) GetBillingInternalId() string {
	if x != nil {
		return x.BillingInternalId
	}
	return ""
}

func (x *Reseller) GetBillingGroupId() string {
	if x != nil {
		return x.BillingGroupId
	}
	return ""
}

func (x *Reseller) GetBillingGroupName() string {
	if x != nil {
		return x.BillingGroupName
	}
	return ""
}

func (x *Reseller) GetGroupType() string {
	if x != nil {
		return x.GroupType
	}
	return ""
}

func (x *Reseller) GetWaveStatus() string {
	if x != nil {
		return x.WaveStatus
	}
	return ""
}

func (x *Reseller) GetWaveConfig() []*ResellerConfig {
	if x != nil {
		return x.WaveConfig
	}
	return nil
}

func (x *Reseller) GetWaveProConfig() []*ResellerConfig {
	if x != nil {
		return x.WaveProConfig
	}
	return nil
}

func (x *Reseller) GetAquaConfig() []*ResellerConfig {
	if x != nil {
		return x.AquaConfig
	}
	return nil
}

type ResellerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value
	Value bool `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ResellerConfig) Reset() {
	*x = ResellerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ripple_reseller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResellerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResellerConfig) ProtoMessage() {}

func (x *ResellerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_reseller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResellerConfig.ProtoReflect.Descriptor instead.
func (*ResellerConfig) Descriptor() ([]byte, []int) {
	return file_api_ripple_reseller_proto_rawDescGZIP(), []int{1}
}

func (x *ResellerConfig) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ResellerConfig) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

var File_api_ripple_reseller_proto protoreflect.FileDescriptor

var file_api_ripple_reseller_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x73,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x22,
	0xc2, 0x03, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x42, 0x0a, 0x0a, 0x77, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x77, 0x61, 0x76, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x48, 0x0a, 0x0d, 0x77, 0x61, 0x76, 0x65, 0x50, 0x72,
	0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70,
	0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0d, 0x77, 0x61, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x42, 0x0a, 0x0a, 0x61, 0x71, 0x75, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x61, 0x71, 0x75, 0x61, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x38, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x69,
	0x0a, 0x20, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70,
	0x6c, 0x65, 0x42, 0x16, 0x41, 0x70, 0x69, 0x52, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_ripple_reseller_proto_rawDescOnce sync.Once
	file_api_ripple_reseller_proto_rawDescData = file_api_ripple_reseller_proto_rawDesc
)

func file_api_ripple_reseller_proto_rawDescGZIP() []byte {
	file_api_ripple_reseller_proto_rawDescOnce.Do(func() {
		file_api_ripple_reseller_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ripple_reseller_proto_rawDescData)
	})
	return file_api_ripple_reseller_proto_rawDescData
}

var file_api_ripple_reseller_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_ripple_reseller_proto_goTypes = []interface{}{
	(*Reseller)(nil),       // 0: blueapi.api.ripple.Reseller
	(*ResellerConfig)(nil), // 1: blueapi.api.ripple.ResellerConfig
}
var file_api_ripple_reseller_proto_depIdxs = []int32{
	1, // 0: blueapi.api.ripple.Reseller.waveConfig:type_name -> blueapi.api.ripple.ResellerConfig
	1, // 1: blueapi.api.ripple.Reseller.waveProConfig:type_name -> blueapi.api.ripple.ResellerConfig
	1, // 2: blueapi.api.ripple.Reseller.aquaConfig:type_name -> blueapi.api.ripple.ResellerConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_ripple_reseller_proto_init() }
func file_api_ripple_reseller_proto_init() {
	if File_api_ripple_reseller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ripple_reseller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reseller); i {
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
		file_api_ripple_reseller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResellerConfig); i {
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
			RawDescriptor: file_api_ripple_reseller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_ripple_reseller_proto_goTypes,
		DependencyIndexes: file_api_ripple_reseller_proto_depIdxs,
		MessageInfos:      file_api_ripple_reseller_proto_msgTypes,
	}.Build()
	File_api_ripple_reseller_proto = out.File
	file_api_ripple_reseller_proto_rawDesc = nil
	file_api_ripple_reseller_proto_goTypes = nil
	file_api_ripple_reseller_proto_depIdxs = nil
}
