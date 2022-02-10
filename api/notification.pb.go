// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/notification.proto

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

type NotificationSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *NotificationSettings) Reset() {
	*x = NotificationSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationSettings) ProtoMessage() {}

func (x *NotificationSettings) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationSettings.ProtoReflect.Descriptor instead.
func (*NotificationSettings) Descriptor() ([]byte, []int) {
	return file_api_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationSettings) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type NotificationChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Enabled bool            `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Default bool            `protobuf:"varint,3,opt,name=default,proto3" json:"default,omitempty"`
	Name    string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Type    string          `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Email   *EmailChannel   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Slack   *SlackChannel   `protobuf:"bytes,7,opt,name=slack,proto3" json:"slack,omitempty"`
	Msteams *MSTeamsChannel `protobuf:"bytes,8,opt,name=msteams,proto3" json:"msteams,omitempty"`
}

func (x *NotificationChannel) Reset() {
	*x = NotificationChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationChannel) ProtoMessage() {}

func (x *NotificationChannel) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationChannel.ProtoReflect.Descriptor instead.
func (*NotificationChannel) Descriptor() ([]byte, []int) {
	return file_api_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationChannel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NotificationChannel) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *NotificationChannel) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *NotificationChannel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NotificationChannel) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NotificationChannel) GetEmail() *EmailChannel {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *NotificationChannel) GetSlack() *SlackChannel {
	if x != nil {
		return x.Slack
	}
	return nil
}

func (x *NotificationChannel) GetMsteams() *MSTeamsChannel {
	if x != nil {
		return x.Msteams
	}
	return nil
}

type EmailChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format     string   `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty"`
	Recipients []string `protobuf:"bytes,2,rep,name=recipients,proto3" json:"recipients,omitempty"`
}

func (x *EmailChannel) Reset() {
	*x = EmailChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailChannel) ProtoMessage() {}

func (x *EmailChannel) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailChannel.ProtoReflect.Descriptor instead.
func (*EmailChannel) Descriptor() ([]byte, []int) {
	return file_api_notification_proto_rawDescGZIP(), []int{2}
}

func (x *EmailChannel) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *EmailChannel) GetRecipients() []string {
	if x != nil {
		return x.Recipients
	}
	return nil
}

type SlackChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebhookUrl       string `protobuf:"bytes,1,opt,name=webhookUrl,proto3" json:"webhookUrl,omitempty"`
	ChannelId        string `protobuf:"bytes,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Channel          string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	ConfigurationUrl string `protobuf:"bytes,4,opt,name=configurationUrl,proto3" json:"configurationUrl,omitempty"`
}

func (x *SlackChannel) Reset() {
	*x = SlackChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlackChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlackChannel) ProtoMessage() {}

func (x *SlackChannel) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlackChannel.ProtoReflect.Descriptor instead.
func (*SlackChannel) Descriptor() ([]byte, []int) {
	return file_api_notification_proto_rawDescGZIP(), []int{3}
}

func (x *SlackChannel) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

func (x *SlackChannel) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SlackChannel) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *SlackChannel) GetConfigurationUrl() string {
	if x != nil {
		return x.ConfigurationUrl
	}
	return ""
}

type MSTeamsChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebhookUrl string `protobuf:"bytes,1,opt,name=webhookUrl,proto3" json:"webhookUrl,omitempty"`
}

func (x *MSTeamsChannel) Reset() {
	*x = MSTeamsChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MSTeamsChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSTeamsChannel) ProtoMessage() {}

func (x *MSTeamsChannel) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSTeamsChannel.ProtoReflect.Descriptor instead.
func (*MSTeamsChannel) Descriptor() ([]byte, []int) {
	return file_api_notification_proto_rawDescGZIP(), []int{4}
}

func (x *MSTeamsChannel) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

var File_api_notification_proto protoreflect.FileDescriptor

var file_api_notification_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x30, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x9a, 0x02, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2f, 0x0a, 0x05,
	0x73, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x35, 0x0a,
	0x07, 0x6d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x53, 0x54,
	0x65, 0x61, 0x6d, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x6d, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x73, 0x22, 0x46, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x92, 0x01, 0x0a,
	0x0c, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72,
	0x6c, 0x22, 0x30, 0x0a, 0x0e, 0x4d, 0x53, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x55, 0x72, 0x6c, 0x42, 0x59, 0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x42, 0x14, 0x41, 0x70, 0x69, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62,
	0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_notification_proto_rawDescOnce sync.Once
	file_api_notification_proto_rawDescData = file_api_notification_proto_rawDesc
)

func file_api_notification_proto_rawDescGZIP() []byte {
	file_api_notification_proto_rawDescOnce.Do(func() {
		file_api_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_notification_proto_rawDescData)
	})
	return file_api_notification_proto_rawDescData
}

var file_api_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_notification_proto_goTypes = []interface{}{
	(*NotificationSettings)(nil), // 0: blueapi.api.NotificationSettings
	(*NotificationChannel)(nil),  // 1: blueapi.api.NotificationChannel
	(*EmailChannel)(nil),         // 2: blueapi.api.EmailChannel
	(*SlackChannel)(nil),         // 3: blueapi.api.SlackChannel
	(*MSTeamsChannel)(nil),       // 4: blueapi.api.MSTeamsChannel
}
var file_api_notification_proto_depIdxs = []int32{
	2, // 0: blueapi.api.NotificationChannel.email:type_name -> blueapi.api.EmailChannel
	3, // 1: blueapi.api.NotificationChannel.slack:type_name -> blueapi.api.SlackChannel
	4, // 2: blueapi.api.NotificationChannel.msteams:type_name -> blueapi.api.MSTeamsChannel
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_notification_proto_init() }
func file_api_notification_proto_init() {
	if File_api_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationSettings); i {
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
		file_api_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationChannel); i {
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
		file_api_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailChannel); i {
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
		file_api_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlackChannel); i {
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
		file_api_notification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MSTeamsChannel); i {
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
			RawDescriptor: file_api_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_notification_proto_goTypes,
		DependencyIndexes: file_api_notification_proto_depIdxs,
		MessageInfos:      file_api_notification_proto_msgTypes,
	}.Build()
	File_api_notification_proto = out.File
	file_api_notification_proto_rawDesc = nil
	file_api_notification_proto_goTypes = nil
	file_api_notification_proto_depIdxs = nil
}
