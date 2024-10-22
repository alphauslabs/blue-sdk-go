// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/cover/accountaccess.proto

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

type RegistrationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiAccess           bool `protobuf:"varint,1,opt,name=apiAccess,proto3" json:"apiAccess,omitempty"`
	CloudwatchStreaming bool `protobuf:"varint,2,opt,name=cloudwatchStreaming,proto3" json:"cloudwatchStreaming,omitempty"`
	CurExport           bool `protobuf:"varint,3,opt,name=curExport,proto3" json:"curExport,omitempty"`
	Payer               bool `protobuf:"varint,4,opt,name=payer,proto3" json:"payer,omitempty"`
	StackSet            bool `protobuf:"varint,5,opt,name=stackSet,proto3" json:"stackSet,omitempty"`
	TransferAccount     bool `protobuf:"varint,6,opt,name=transferAccount,proto3" json:"transferAccount,omitempty"`
	IsDataAvailable     bool `protobuf:"varint,7,opt,name=isDataAvailable,proto3" json:"isDataAvailable,omitempty"`
}

func (x *RegistrationStatus) Reset() {
	*x = RegistrationStatus{}
	mi := &file_api_cover_accountaccess_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistrationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationStatus) ProtoMessage() {}

func (x *RegistrationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_accountaccess_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationStatus.ProtoReflect.Descriptor instead.
func (*RegistrationStatus) Descriptor() ([]byte, []int) {
	return file_api_cover_accountaccess_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrationStatus) GetApiAccess() bool {
	if x != nil {
		return x.ApiAccess
	}
	return false
}

func (x *RegistrationStatus) GetCloudwatchStreaming() bool {
	if x != nil {
		return x.CloudwatchStreaming
	}
	return false
}

func (x *RegistrationStatus) GetCurExport() bool {
	if x != nil {
		return x.CurExport
	}
	return false
}

func (x *RegistrationStatus) GetPayer() bool {
	if x != nil {
		return x.Payer
	}
	return false
}

func (x *RegistrationStatus) GetStackSet() bool {
	if x != nil {
		return x.StackSet
	}
	return false
}

func (x *RegistrationStatus) GetTransferAccount() bool {
	if x != nil {
		return x.TransferAccount
	}
	return false
}

func (x *RegistrationStatus) GetIsDataAvailable() bool {
	if x != nil {
		return x.IsDataAvailable
	}
	return false
}

type TagData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagKey   string   `protobuf:"bytes,1,opt,name=tagKey,proto3" json:"tagKey,omitempty"`
	TagValue []string `protobuf:"bytes,2,rep,name=tagValue,proto3" json:"tagValue,omitempty"`
}

func (x *TagData) Reset() {
	*x = TagData{}
	mi := &file_api_cover_accountaccess_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TagData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagData) ProtoMessage() {}

func (x *TagData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_accountaccess_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagData.ProtoReflect.Descriptor instead.
func (*TagData) Descriptor() ([]byte, []int) {
	return file_api_cover_accountaccess_proto_rawDescGZIP(), []int{1}
}

func (x *TagData) GetTagKey() string {
	if x != nil {
		return x.TagKey
	}
	return ""
}

func (x *TagData) GetTagValue() []string {
	if x != nil {
		return x.TagValue
	}
	return nil
}

type GcpOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName string `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	// where dataset is stored
	ProjectId     string `protobuf:"bytes,2,opt,name=projectId,proto3" json:"projectId,omitempty"`
	DatasetId     string `protobuf:"bytes,3,opt,name=datasetId,proto3" json:"datasetId,omitempty"`
	DatasetRegion string `protobuf:"bytes,4,opt,name=datasetRegion,proto3" json:"datasetRegion,omitempty"`
}

func (x *GcpOptions) Reset() {
	*x = GcpOptions{}
	mi := &file_api_cover_accountaccess_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GcpOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpOptions) ProtoMessage() {}

func (x *GcpOptions) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_accountaccess_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpOptions.ProtoReflect.Descriptor instead.
func (*GcpOptions) Descriptor() ([]byte, []int) {
	return file_api_cover_accountaccess_proto_rawDescGZIP(), []int{2}
}

func (x *GcpOptions) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GcpOptions) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GcpOptions) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *GcpOptions) GetDatasetRegion() string {
	if x != nil {
		return x.DatasetRegion
	}
	return ""
}

type AzureOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName     string `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	AzureCustomerId string `protobuf:"bytes,2,opt,name=azureCustomerId,proto3" json:"azureCustomerId,omitempty"`
	AzurePlanId     string `protobuf:"bytes,3,opt,name=azurePlanId,proto3" json:"azurePlanId,omitempty"`
	ServiceAcct     string `protobuf:"bytes,4,opt,name=serviceAcct,proto3" json:"serviceAcct,omitempty"`
	PartnerAcct     string `protobuf:"bytes,5,opt,name=partnerAcct,proto3" json:"partnerAcct,omitempty"`
	CompanyId       string `protobuf:"bytes,6,opt,name=companyId,proto3" json:"companyId,omitempty"`
	PayerId         string `protobuf:"bytes,7,opt,name=payerId,proto3" json:"payerId,omitempty"`
}

func (x *AzureOptions) Reset() {
	*x = AzureOptions{}
	mi := &file_api_cover_accountaccess_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureOptions) ProtoMessage() {}

func (x *AzureOptions) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_accountaccess_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureOptions.ProtoReflect.Descriptor instead.
func (*AzureOptions) Descriptor() ([]byte, []int) {
	return file_api_cover_accountaccess_proto_rawDescGZIP(), []int{3}
}

func (x *AzureOptions) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *AzureOptions) GetAzureCustomerId() string {
	if x != nil {
		return x.AzureCustomerId
	}
	return ""
}

func (x *AzureOptions) GetAzurePlanId() string {
	if x != nil {
		return x.AzurePlanId
	}
	return ""
}

func (x *AzureOptions) GetServiceAcct() string {
	if x != nil {
		return x.ServiceAcct
	}
	return ""
}

func (x *AzureOptions) GetPartnerAcct() string {
	if x != nil {
		return x.PartnerAcct
	}
	return ""
}

func (x *AzureOptions) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *AzureOptions) GetPayerId() string {
	if x != nil {
		return x.PayerId
	}
	return ""
}

type AwsOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName        string              `protobuf:"bytes,1,opt,name=AccountName,proto3" json:"AccountName,omitempty"`
	PayerId            string              `protobuf:"bytes,2,opt,name=PayerId,proto3" json:"PayerId,omitempty"`
	RoleArn            string              `protobuf:"bytes,3,opt,name=RoleArn,proto3" json:"RoleArn,omitempty"`
	ExternalId         string              `protobuf:"bytes,4,opt,name=ExternalId,proto3" json:"ExternalId,omitempty"`
	StackId            string              `protobuf:"bytes,5,opt,name=StackId,proto3" json:"StackId,omitempty"`
	StackRegion        string              `protobuf:"bytes,6,opt,name=StackRegion,proto3" json:"StackRegion,omitempty"`
	TemplateUrl        string              `protobuf:"bytes,7,opt,name=TemplateUrl,proto3" json:"TemplateUrl,omitempty"`
	BucketName         string              `protobuf:"bytes,8,opt,name=BucketName,proto3" json:"BucketName,omitempty"`
	Prefix             string              `protobuf:"bytes,9,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	ReportName         string              `protobuf:"bytes,10,opt,name=ReportName,proto3" json:"ReportName,omitempty"`
	RegistrationStatus *RegistrationStatus `protobuf:"bytes,11,opt,name=registrationStatus,proto3" json:"registrationStatus,omitempty"`
	Status             string              `protobuf:"bytes,12,opt,name=Status,proto3" json:"Status,omitempty"`
	// Valid values for now are: 'console', 'terraform'. default would be 'console'
	RegistrationMethod string `protobuf:"bytes,13,opt,name=RegistrationMethod,proto3" json:"RegistrationMethod,omitempty"`
}

func (x *AwsOptions) Reset() {
	*x = AwsOptions{}
	mi := &file_api_cover_accountaccess_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsOptions) ProtoMessage() {}

func (x *AwsOptions) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_accountaccess_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsOptions.ProtoReflect.Descriptor instead.
func (*AwsOptions) Descriptor() ([]byte, []int) {
	return file_api_cover_accountaccess_proto_rawDescGZIP(), []int{4}
}

func (x *AwsOptions) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *AwsOptions) GetPayerId() string {
	if x != nil {
		return x.PayerId
	}
	return ""
}

func (x *AwsOptions) GetRoleArn() string {
	if x != nil {
		return x.RoleArn
	}
	return ""
}

func (x *AwsOptions) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *AwsOptions) GetStackId() string {
	if x != nil {
		return x.StackId
	}
	return ""
}

func (x *AwsOptions) GetStackRegion() string {
	if x != nil {
		return x.StackRegion
	}
	return ""
}

func (x *AwsOptions) GetTemplateUrl() string {
	if x != nil {
		return x.TemplateUrl
	}
	return ""
}

func (x *AwsOptions) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *AwsOptions) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *AwsOptions) GetReportName() string {
	if x != nil {
		return x.ReportName
	}
	return ""
}

func (x *AwsOptions) GetRegistrationStatus() *RegistrationStatus {
	if x != nil {
		return x.RegistrationStatus
	}
	return nil
}

func (x *AwsOptions) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AwsOptions) GetRegistrationMethod() string {
	if x != nil {
		return x.RegistrationMethod
	}
	return ""
}

var File_api_cover_accountaccess_proto protoreflect.FileDescriptor

var file_api_cover_accountaccess_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x22, 0x88, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x69,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x70,
	0x69, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x75, 0x72,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x75,
	0x72, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x73, 0x44, 0x61, 0x74, 0x61, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x3d, 0x0a,
	0x07, 0x54, 0x61, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x4b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x67, 0x4b, 0x65, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x90, 0x01, 0x0a,
	0x0a, 0x47, 0x63, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22,
	0xf8, 0x01, 0x0a, 0x0c, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0xd7, 0x03, 0x0a, 0x0a, 0x41,
	0x77, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x42, 0x6b, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x1a, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65,
	0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cover_accountaccess_proto_rawDescOnce sync.Once
	file_api_cover_accountaccess_proto_rawDescData = file_api_cover_accountaccess_proto_rawDesc
)

func file_api_cover_accountaccess_proto_rawDescGZIP() []byte {
	file_api_cover_accountaccess_proto_rawDescOnce.Do(func() {
		file_api_cover_accountaccess_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_accountaccess_proto_rawDescData)
	})
	return file_api_cover_accountaccess_proto_rawDescData
}

var file_api_cover_accountaccess_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_cover_accountaccess_proto_goTypes = []any{
	(*RegistrationStatus)(nil), // 0: blueapi.api.cover.RegistrationStatus
	(*TagData)(nil),            // 1: blueapi.api.cover.TagData
	(*GcpOptions)(nil),         // 2: blueapi.api.cover.GcpOptions
	(*AzureOptions)(nil),       // 3: blueapi.api.cover.AzureOptions
	(*AwsOptions)(nil),         // 4: blueapi.api.cover.AwsOptions
}
var file_api_cover_accountaccess_proto_depIdxs = []int32{
	0, // 0: blueapi.api.cover.AwsOptions.registrationStatus:type_name -> blueapi.api.cover.RegistrationStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_cover_accountaccess_proto_init() }
func file_api_cover_accountaccess_proto_init() {
	if File_api_cover_accountaccess_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cover_accountaccess_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_accountaccess_proto_goTypes,
		DependencyIndexes: file_api_cover_accountaccess_proto_depIdxs,
		MessageInfos:      file_api_cover_accountaccess_proto_msgTypes,
	}.Build()
	File_api_cover_accountaccess_proto = out.File
	file_api_cover_accountaccess_proto_rawDesc = nil
	file_api_cover_accountaccess_proto_goTypes = nil
	file_api_cover_accountaccess_proto_depIdxs = nil
}
