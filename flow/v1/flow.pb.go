// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: flow/v1/flow.proto

package flow

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for the Flow.GetInfo rpc.
type GetInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	mi := &file_flow_v1_flow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flow_v1_flow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_flow_v1_flow_proto_rawDescGZIP(), []int{0}
}

// Request message for the Flow.CreateSettings rpc.
type CreateSettingsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The id of the payer.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required. Customization setting for SP.
	Customization string `protobuf:"bytes,2,opt,name=customization,proto3" json:"customization,omitempty"`
	// Required. Term of the SP.
	PlanTerm string `protobuf:"bytes,3,opt,name=planTerm,proto3" json:"planTerm,omitempty"`
	// Required. Payment option for the SP.
	PaymentOption string `protobuf:"bytes,4,opt,name=paymentOption,proto3" json:"paymentOption,omitempty"`
	// Required. Lookback period for recommendation.
	LookBackPeriod string `protobuf:"bytes,5,opt,name=lookBackPeriod,proto3" json:"lookBackPeriod,omitempty"`
	// Optional. If EC2 Instance SP is selected in Customization, request will include list/s of instance family
	InstanceFamily string `protobuf:"bytes,6,opt,name=instanceFamily,proto3" json:"instanceFamily,omitempty"`
	// Required. Annual budget input for SP.
	AnnualBudget string `protobuf:"bytes,7,opt,name=annualBudget,proto3" json:"annualBudget,omitempty"`
	// Required. Monthly budget is automatically generated from the annual budget.
	MonthlyBudget string `protobuf:"bytes,8,opt,name=monthlyBudget,proto3" json:"monthlyBudget,omitempty"`
	// Required. Purchase approval from the payer for the SP.
	Approval      bool `protobuf:"varint,9,opt,name=approval,proto3" json:"approval,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSettingsRequest) Reset() {
	*x = CreateSettingsRequest{}
	mi := &file_flow_v1_flow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSettingsRequest) ProtoMessage() {}

func (x *CreateSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flow_v1_flow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSettingsRequest.ProtoReflect.Descriptor instead.
func (*CreateSettingsRequest) Descriptor() ([]byte, []int) {
	return file_flow_v1_flow_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSettingsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateSettingsRequest) GetCustomization() string {
	if x != nil {
		return x.Customization
	}
	return ""
}

func (x *CreateSettingsRequest) GetPlanTerm() string {
	if x != nil {
		return x.PlanTerm
	}
	return ""
}

func (x *CreateSettingsRequest) GetPaymentOption() string {
	if x != nil {
		return x.PaymentOption
	}
	return ""
}

func (x *CreateSettingsRequest) GetLookBackPeriod() string {
	if x != nil {
		return x.LookBackPeriod
	}
	return ""
}

func (x *CreateSettingsRequest) GetInstanceFamily() string {
	if x != nil {
		return x.InstanceFamily
	}
	return ""
}

func (x *CreateSettingsRequest) GetAnnualBudget() string {
	if x != nil {
		return x.AnnualBudget
	}
	return ""
}

func (x *CreateSettingsRequest) GetMonthlyBudget() string {
	if x != nil {
		return x.MonthlyBudget
	}
	return ""
}

func (x *CreateSettingsRequest) GetApproval() bool {
	if x != nil {
		return x.Approval
	}
	return false
}

// Response message for the Flow.GetInfo rpc.
type GetInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      string                 `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	mi := &file_flow_v1_flow_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flow_v1_flow_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_flow_v1_flow_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfoResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

// Response message for the Flow.CreateSettings rpc.
type CreateSettingsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      string                 `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSettingsResponse) Reset() {
	*x = CreateSettingsResponse{}
	mi := &file_flow_v1_flow_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSettingsResponse) ProtoMessage() {}

func (x *CreateSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flow_v1_flow_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSettingsResponse.ProtoReflect.Descriptor instead.
func (*CreateSettingsResponse) Descriptor() ([]byte, []int) {
	return file_flow_v1_flow_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSettingsResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_flow_v1_flow_proto protoreflect.FileDescriptor

var file_flow_v1_flow_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc5, 0x02, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x72,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x72,
	0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x42,
	0x61, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x42, 0x61, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6e, 0x6e, 0x75, 0x61,
	0x6c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x22, 0x2d, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xfa, 0x02, 0x0a, 0x04, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x5e, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x7a, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x26, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x95, 0x01, 0x92, 0x41, 0x91, 0x01, 0x12, 0x41,
	0x28, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x29, 0x20, 0x46, 0x6c, 0x6f, 0x77, 0x20, 0x41, 0x50, 0x49,
	0x2e, 0x20, 0x42, 0x61, 0x73, 0x65, 0x20, 0x55, 0x52, 0x4c, 0x3a, 0x20, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2f, 0x66, 0x6c, 0x6f,
	0x77, 0x1a, 0x4c, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x42,
	0x48, 0x0a, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x09, 0x46, 0x6c, 0x6f, 0x77, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_flow_v1_flow_proto_rawDescOnce sync.Once
	file_flow_v1_flow_proto_rawDescData []byte
)

func file_flow_v1_flow_proto_rawDescGZIP() []byte {
	file_flow_v1_flow_proto_rawDescOnce.Do(func() {
		file_flow_v1_flow_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_flow_v1_flow_proto_rawDesc), len(file_flow_v1_flow_proto_rawDesc)))
	})
	return file_flow_v1_flow_proto_rawDescData
}

var file_flow_v1_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_flow_v1_flow_proto_goTypes = []any{
	(*GetInfoRequest)(nil),         // 0: blueapi.flow.v1.GetInfoRequest
	(*CreateSettingsRequest)(nil),  // 1: blueapi.flow.v1.CreateSettingsRequest
	(*GetInfoResponse)(nil),        // 2: blueapi.flow.v1.GetInfoResponse
	(*CreateSettingsResponse)(nil), // 3: blueapi.flow.v1.CreateSettingsResponse
}
var file_flow_v1_flow_proto_depIdxs = []int32{
	0, // 0: blueapi.flow.v1.Flow.GetInfo:input_type -> blueapi.flow.v1.GetInfoRequest
	1, // 1: blueapi.flow.v1.Flow.CreateSettings:input_type -> blueapi.flow.v1.CreateSettingsRequest
	2, // 2: blueapi.flow.v1.Flow.GetInfo:output_type -> blueapi.flow.v1.GetInfoResponse
	3, // 3: blueapi.flow.v1.Flow.CreateSettings:output_type -> blueapi.flow.v1.CreateSettingsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flow_v1_flow_proto_init() }
func file_flow_v1_flow_proto_init() {
	if File_flow_v1_flow_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_flow_v1_flow_proto_rawDesc), len(file_flow_v1_flow_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flow_v1_flow_proto_goTypes,
		DependencyIndexes: file_flow_v1_flow_proto_depIdxs,
		MessageInfos:      file_flow_v1_flow_proto_msgTypes,
	}.Build()
	File_flow_v1_flow_proto = out.File
	file_flow_v1_flow_proto_goTypes = nil
	file_flow_v1_flow_proto_depIdxs = nil
}
