// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: api/cover/alert.proto

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

type AlertData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Granularity   string                 `protobuf:"bytes,3,opt,name=granularity,proto3" json:"granularity,omitempty"`
	CostGroups    []*AlertCostGroup      `protobuf:"bytes,4,rep,name=costGroups,proto3" json:"costGroups,omitempty"`
	Channels      *AlertChannels         `protobuf:"bytes,5,opt,name=channels,proto3" json:"channels,omitempty"`
	CreatedBy     string                 `protobuf:"bytes,6,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	FixedAmount   float32                `protobuf:"fixed32,9,opt,name=fixedAmount,proto3" json:"fixedAmount,omitempty"`
	Percentage    float32                `protobuf:"fixed32,10,opt,name=percentage,proto3" json:"percentage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertData) Reset() {
	*x = AlertData{}
	mi := &file_api_cover_alert_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertData) ProtoMessage() {}

func (x *AlertData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_alert_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertData.ProtoReflect.Descriptor instead.
func (*AlertData) Descriptor() ([]byte, []int) {
	return file_api_cover_alert_proto_rawDescGZIP(), []int{0}
}

func (x *AlertData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AlertData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertData) GetGranularity() string {
	if x != nil {
		return x.Granularity
	}
	return ""
}

func (x *AlertData) GetCostGroups() []*AlertCostGroup {
	if x != nil {
		return x.CostGroups
	}
	return nil
}

func (x *AlertData) GetChannels() *AlertChannels {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *AlertData) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *AlertData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AlertData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AlertData) GetFixedAmount() float32 {
	if x != nil {
		return x.FixedAmount
	}
	return 0
}

func (x *AlertData) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

type ChannelData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // actual email value of channel name in slack or ms teams
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"` // email, slack, msteams
	WebhookUrl    string                 `protobuf:"bytes,4,opt,name=webhookUrl,proto3" json:"webhookUrl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChannelData) Reset() {
	*x = ChannelData{}
	mi := &file_api_cover_alert_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChannelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelData) ProtoMessage() {}

func (x *ChannelData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_alert_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelData.ProtoReflect.Descriptor instead.
func (*ChannelData) Descriptor() ([]byte, []int) {
	return file_api_cover_alert_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChannelData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ChannelData) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

type AlertChannels struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         []*AlertChannel        `protobuf:"bytes,1,rep,name=email,proto3" json:"email,omitempty"`
	Slack         []*AlertChannel        `protobuf:"bytes,2,rep,name=slack,proto3" json:"slack,omitempty"`
	Msteams       []*AlertChannel        `protobuf:"bytes,3,rep,name=msteams,proto3" json:"msteams,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertChannels) Reset() {
	*x = AlertChannels{}
	mi := &file_api_cover_alert_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertChannels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertChannels) ProtoMessage() {}

func (x *AlertChannels) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_alert_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertChannels.ProtoReflect.Descriptor instead.
func (*AlertChannels) Descriptor() ([]byte, []int) {
	return file_api_cover_alert_proto_rawDescGZIP(), []int{2}
}

func (x *AlertChannels) GetEmail() []*AlertChannel {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *AlertChannels) GetSlack() []*AlertChannel {
	if x != nil {
		return x.Slack
	}
	return nil
}

func (x *AlertChannels) GetMsteams() []*AlertChannel {
	if x != nil {
		return x.Msteams
	}
	return nil
}

type AlertChannel struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertChannel) Reset() {
	*x = AlertChannel{}
	mi := &file_api_cover_alert_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertChannel) ProtoMessage() {}

func (x *AlertChannel) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_alert_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertChannel.ProtoReflect.Descriptor instead.
func (*AlertChannel) Descriptor() ([]byte, []int) {
	return file_api_cover_alert_proto_rawDescGZIP(), []int{3}
}

func (x *AlertChannel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AlertChannel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AlertCostGroup struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertCostGroup) Reset() {
	*x = AlertCostGroup{}
	mi := &file_api_cover_alert_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertCostGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertCostGroup) ProtoMessage() {}

func (x *AlertCostGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_alert_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertCostGroup.ProtoReflect.Descriptor instead.
func (*AlertCostGroup) Descriptor() ([]byte, []int) {
	return file_api_cover_alert_proto_rawDescGZIP(), []int{4}
}

func (x *AlertCostGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AlertCostGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_cover_alert_proto protoreflect.FileDescriptor

var file_api_cover_alert_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x22, 0xee, 0x02, 0x0a, 0x09, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x41,
	0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x73, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x22, 0x65, 0x0a, 0x0b, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55,
	0x72, 0x6c, 0x22, 0xb8, 0x01, 0x0a, 0x0d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x12, 0x35, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x35, 0x0a, 0x05, 0x73,
	0x6c, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x05, 0x73, 0x6c, 0x61,
	0x63, 0x6b, 0x12, 0x39, 0x0a, 0x07, 0x6d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x6d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x32, 0x0a,
	0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x34, 0x0a, 0x0e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x63, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x12, 0x41, 0x70, 0x69, 0x43,
	0x6f, 0x76, 0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cover_alert_proto_rawDescOnce sync.Once
	file_api_cover_alert_proto_rawDescData = file_api_cover_alert_proto_rawDesc
)

func file_api_cover_alert_proto_rawDescGZIP() []byte {
	file_api_cover_alert_proto_rawDescOnce.Do(func() {
		file_api_cover_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_alert_proto_rawDescData)
	})
	return file_api_cover_alert_proto_rawDescData
}

var file_api_cover_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_cover_alert_proto_goTypes = []any{
	(*AlertData)(nil),      // 0: blueapi.api.cover.AlertData
	(*ChannelData)(nil),    // 1: blueapi.api.cover.ChannelData
	(*AlertChannels)(nil),  // 2: blueapi.api.cover.AlertChannels
	(*AlertChannel)(nil),   // 3: blueapi.api.cover.AlertChannel
	(*AlertCostGroup)(nil), // 4: blueapi.api.cover.AlertCostGroup
}
var file_api_cover_alert_proto_depIdxs = []int32{
	4, // 0: blueapi.api.cover.AlertData.costGroups:type_name -> blueapi.api.cover.AlertCostGroup
	2, // 1: blueapi.api.cover.AlertData.channels:type_name -> blueapi.api.cover.AlertChannels
	3, // 2: blueapi.api.cover.AlertChannels.email:type_name -> blueapi.api.cover.AlertChannel
	3, // 3: blueapi.api.cover.AlertChannels.slack:type_name -> blueapi.api.cover.AlertChannel
	3, // 4: blueapi.api.cover.AlertChannels.msteams:type_name -> blueapi.api.cover.AlertChannel
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_cover_alert_proto_init() }
func file_api_cover_alert_proto_init() {
	if File_api_cover_alert_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cover_alert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_alert_proto_goTypes,
		DependencyIndexes: file_api_cover_alert_proto_depIdxs,
		MessageInfos:      file_api_cover_alert_proto_msgTypes,
	}.Build()
	File_api_cover_alert_proto = out.File
	file_api_cover_alert_proto_rawDesc = nil
	file_api_cover_alert_proto_goTypes = nil
	file_api_cover_alert_proto_depIdxs = nil
}
