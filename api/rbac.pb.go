// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: api/rbac.proto

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

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Effect string   `protobuf:"bytes,2,opt,name=effect,proto3" json:"effect,omitempty"`
	Key    string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,4,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_api_rbac_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Policy) GetEffect() string {
	if x != nil {
		return x.Effect
	}
	return ""
}

func (x *Policy) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Policy) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace   string    `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Permissions []string  `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Policies    []*Policy `protobuf:"bytes,3,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_api_rbac_proto_rawDescGZIP(), []int{1}
}

func (x *Permission) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Permission) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Permission) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   string    `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Permissions []string  `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Policies    []*Policy `protobuf:"bytes,4,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_api_rbac_proto_rawDescGZIP(), []int{2}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Role) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Role) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

type UserRoleMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RootUser  string `protobuf:"bytes,1,opt,name=rootUser,proto3" json:"rootUser,omitempty"`
	SubUser   string `protobuf:"bytes,2,opt,name=subUser,proto3" json:"subUser,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Role      string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	Filter    string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *UserRoleMapping) Reset() {
	*x = UserRoleMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRoleMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoleMapping) ProtoMessage() {}

func (x *UserRoleMapping) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoleMapping.ProtoReflect.Descriptor instead.
func (*UserRoleMapping) Descriptor() ([]byte, []int) {
	return file_api_rbac_proto_rawDescGZIP(), []int{3}
}

func (x *UserRoleMapping) GetRootUser() string {
	if x != nil {
		return x.RootUser
	}
	return ""
}

func (x *UserRoleMapping) GetSubUser() string {
	if x != nil {
		return x.SubUser
	}
	return ""
}

func (x *UserRoleMapping) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *UserRoleMapping) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserRoleMapping) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

var File_api_rbac_proto protoreflect.FileDescriptor

var file_api_rbac_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x62, 0x0a,
	0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x22, 0x7d, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2f, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x22, 0x8b, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a,
	0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x22, 0x91,
	0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x42, 0x51, 0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x42,
	0x0c, 0x41, 0x70, 0x69, 0x52, 0x62, 0x61, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75,
	0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rbac_proto_rawDescOnce sync.Once
	file_api_rbac_proto_rawDescData = file_api_rbac_proto_rawDesc
)

func file_api_rbac_proto_rawDescGZIP() []byte {
	file_api_rbac_proto_rawDescOnce.Do(func() {
		file_api_rbac_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rbac_proto_rawDescData)
	})
	return file_api_rbac_proto_rawDescData
}

var file_api_rbac_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_rbac_proto_goTypes = []interface{}{
	(*Policy)(nil),          // 0: blueapi.api.Policy
	(*Permission)(nil),      // 1: blueapi.api.Permission
	(*Role)(nil),            // 2: blueapi.api.Role
	(*UserRoleMapping)(nil), // 3: blueapi.api.UserRoleMapping
}
var file_api_rbac_proto_depIdxs = []int32{
	0, // 0: blueapi.api.Permission.policies:type_name -> blueapi.api.Policy
	0, // 1: blueapi.api.Role.policies:type_name -> blueapi.api.Policy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_rbac_proto_init() }
func file_api_rbac_proto_init() {
	if File_api_rbac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rbac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_api_rbac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_api_rbac_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_api_rbac_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRoleMapping); i {
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
			RawDescriptor: file_api_rbac_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_rbac_proto_goTypes,
		DependencyIndexes: file_api_rbac_proto_depIdxs,
		MessageInfos:      file_api_rbac_proto_msgTypes,
	}.Build()
	File_api_rbac_proto = out.File
	file_api_rbac_proto_rawDesc = nil
	file_api_rbac_proto_goTypes = nil
	file_api_rbac_proto_depIdxs = nil
}
