// Copyright 2018 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/api/backend.proto

package serviceconfig

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// `Backend` defines the backend configuration for a service.
type Backend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of API backend rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*BackendRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *Backend) Reset() {
	*x = Backend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Backend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Backend) ProtoMessage() {}

func (x *Backend) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Backend.ProtoReflect.Descriptor instead.
func (*Backend) Descriptor() ([]byte, []int) {
	return file_google_api_backend_proto_rawDescGZIP(), []int{0}
}

func (x *Backend) GetRules() []*BackendRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// A backend rule provides configuration for an individual API element.
type BackendRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The address of the API backend.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// The number of seconds to wait for a response from a request.  The default
	// deadline for gRPC is infinite (no deadline) and HTTP requests is 5 seconds.
	Deadline float64 `protobuf:"fixed64,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Minimum deadline in seconds needed for this method. Calls having deadline
	// value lower than this will be rejected.
	MinDeadline float64 `protobuf:"fixed64,4,opt,name=min_deadline,json=minDeadline,proto3" json:"min_deadline,omitempty"`
}

func (x *BackendRule) Reset() {
	*x = BackendRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendRule) ProtoMessage() {}

func (x *BackendRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendRule.ProtoReflect.Descriptor instead.
func (*BackendRule) Descriptor() ([]byte, []int) {
	return file_google_api_backend_proto_rawDescGZIP(), []int{1}
}

func (x *BackendRule) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *BackendRule) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BackendRule) GetDeadline() float64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

func (x *BackendRule) GetMinDeadline() float64 {
	if x != nil {
		return x.MinDeadline
	}
	return 0
}

var File_google_api_backend_proto protoreflect.FileDescriptor

var file_google_api_backend_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x38, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x12, 0x2d, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x22, 0x82, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x6e, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0c, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xa2, 0x02,
	0x04, 0x47, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_backend_proto_rawDescOnce sync.Once
	file_google_api_backend_proto_rawDescData = file_google_api_backend_proto_rawDesc
)

func file_google_api_backend_proto_rawDescGZIP() []byte {
	file_google_api_backend_proto_rawDescOnce.Do(func() {
		file_google_api_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_backend_proto_rawDescData)
	})
	return file_google_api_backend_proto_rawDescData
}

var file_google_api_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_api_backend_proto_goTypes = []interface{}{
	(*Backend)(nil),     // 0: google.api.Backend
	(*BackendRule)(nil), // 1: google.api.BackendRule
}
var file_google_api_backend_proto_depIdxs = []int32{
	1, // 0: google.api.Backend.rules:type_name -> google.api.BackendRule
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_api_backend_proto_init() }
func file_google_api_backend_proto_init() {
	if File_google_api_backend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_backend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Backend); i {
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
		file_google_api_backend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackendRule); i {
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
			RawDescriptor: file_google_api_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_backend_proto_goTypes,
		DependencyIndexes: file_google_api_backend_proto_depIdxs,
		MessageInfos:      file_google_api_backend_proto_msgTypes,
	}.Build()
	File_google_api_backend_proto = out.File
	file_google_api_backend_proto_rawDesc = nil
	file_google_api_backend_proto_goTypes = nil
	file_google_api_backend_proto_depIdxs = nil
}
