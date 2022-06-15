// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: api/cover/costgroup.proto

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

type CostGroupData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Image        string          `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Icon         string          `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	ColorTheme   string          `protobuf:"bytes,6,opt,name=colorTheme,proto3" json:"colorTheme,omitempty"`
	CreatedAt    string          `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    string          `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Members      []*UserData     `protobuf:"bytes,9,rep,name=members,proto3" json:"members,omitempty"`
	Combinations []*Combinations `protobuf:"bytes,10,rep,name=combinations,proto3" json:"combinations,omitempty"`
	CreatedBy    *UserData       `protobuf:"bytes,11,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
}

func (x *CostGroupData) Reset() {
	*x = CostGroupData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_costgroup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostGroupData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostGroupData) ProtoMessage() {}

func (x *CostGroupData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_costgroup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostGroupData.ProtoReflect.Descriptor instead.
func (*CostGroupData) Descriptor() ([]byte, []int) {
	return file_api_cover_costgroup_proto_rawDescGZIP(), []int{0}
}

func (x *CostGroupData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CostGroupData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CostGroupData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CostGroupData) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CostGroupData) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *CostGroupData) GetColorTheme() string {
	if x != nil {
		return x.ColorTheme
	}
	return ""
}

func (x *CostGroupData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CostGroupData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *CostGroupData) GetMembers() []*UserData {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *CostGroupData) GetCombinations() []*Combinations {
	if x != nil {
		return x.Combinations
	}
	return nil
}

func (x *CostGroupData) GetCreatedBy() *UserData {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

type Combinations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Criteria   string `protobuf:"bytes,1,opt,name=criteria,proto3" json:"criteria,omitempty"`
	Comparator string `protobuf:"bytes,2,opt,name=comparator,proto3" json:"comparator,omitempty"`
	Value      string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	TagKey     string `protobuf:"bytes,4,opt,name=tagKey,proto3" json:"tagKey,omitempty"`
	TagValue   string `protobuf:"bytes,5,opt,name=tagValue,proto3" json:"tagValue,omitempty"`
}

func (x *Combinations) Reset() {
	*x = Combinations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_costgroup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Combinations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Combinations) ProtoMessage() {}

func (x *Combinations) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_costgroup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Combinations.ProtoReflect.Descriptor instead.
func (*Combinations) Descriptor() ([]byte, []int) {
	return file_api_cover_costgroup_proto_rawDescGZIP(), []int{1}
}

func (x *Combinations) GetCriteria() string {
	if x != nil {
		return x.Criteria
	}
	return ""
}

func (x *Combinations) GetComparator() string {
	if x != nil {
		return x.Comparator
	}
	return ""
}

func (x *Combinations) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Combinations) GetTagKey() string {
	if x != nil {
		return x.TagKey
	}
	return ""
}

func (x *Combinations) GetTagValue() string {
	if x != nil {
		return x.TagValue
	}
	return ""
}

var File_api_cover_costgroup_proto protoreflect.FileDescriptor

var file_api_cover_costgroup_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x1a, 0x14,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x54,
	0x68, 0x65, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x43, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x67, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x67, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x67, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75,
	0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x42, 0x16, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x73,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_cover_costgroup_proto_rawDescOnce sync.Once
	file_api_cover_costgroup_proto_rawDescData = file_api_cover_costgroup_proto_rawDesc
)

func file_api_cover_costgroup_proto_rawDescGZIP() []byte {
	file_api_cover_costgroup_proto_rawDescOnce.Do(func() {
		file_api_cover_costgroup_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_costgroup_proto_rawDescData)
	})
	return file_api_cover_costgroup_proto_rawDescData
}

var file_api_cover_costgroup_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_cover_costgroup_proto_goTypes = []interface{}{
	(*CostGroupData)(nil), // 0: blueapi.api.cover.CostGroupData
	(*Combinations)(nil),  // 1: blueapi.api.cover.Combinations
	(*UserData)(nil),      // 2: blueapi.api.cover.UserData
}
var file_api_cover_costgroup_proto_depIdxs = []int32{
	2, // 0: blueapi.api.cover.CostGroupData.members:type_name -> blueapi.api.cover.UserData
	1, // 1: blueapi.api.cover.CostGroupData.combinations:type_name -> blueapi.api.cover.Combinations
	2, // 2: blueapi.api.cover.CostGroupData.createdBy:type_name -> blueapi.api.cover.UserData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_cover_costgroup_proto_init() }
func file_api_cover_costgroup_proto_init() {
	if File_api_cover_costgroup_proto != nil {
		return
	}
	file_api_cover_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_cover_costgroup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostGroupData); i {
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
		file_api_cover_costgroup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Combinations); i {
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
			RawDescriptor: file_api_cover_costgroup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_costgroup_proto_goTypes,
		DependencyIndexes: file_api_cover_costgroup_proto_depIdxs,
		MessageInfos:      file_api_cover_costgroup_proto_msgTypes,
	}.Build()
	File_api_cover_costgroup_proto = out.File
	file_api_cover_costgroup_proto_rawDesc = nil
	file_api_cover_costgroup_proto_goTypes = nil
	file_api_cover_costgroup_proto_depIdxs = nil
}
