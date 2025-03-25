// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/cover/user.proto

package cover

import (
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

type UserData struct {
	state                      protoimpl.MessageState `protogen:"open.v1"`
	Id                         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                      string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Avatar                     string                 `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Icon                       string                 `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	ColorTheme                 string                 `protobuf:"bytes,6,opt,name=colorTheme,proto3" json:"colorTheme,omitempty"`
	Initial                    string                 `protobuf:"bytes,7,opt,name=initial,proto3" json:"initial,omitempty"`
	Activated                  bool                   `protobuf:"varint,8,opt,name=activated,proto3" json:"activated,omitempty"`
	IsAdmin                    bool                   `protobuf:"varint,9,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
	Attributes                 []string               `protobuf:"bytes,10,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Locale                     string                 `protobuf:"bytes,11,opt,name=locale,proto3" json:"locale,omitempty"`
	TimeZone                   string                 `protobuf:"bytes,12,opt,name=timeZone,proto3" json:"timeZone,omitempty"`
	Registered                 string                 `protobuf:"bytes,13,opt,name=registered,proto3" json:"registered,omitempty"`
	SsoEnabled                 bool                   `protobuf:"varint,14,opt,name=ssoEnabled,proto3" json:"ssoEnabled,omitempty"`
	MfaEnabled                 bool                   `protobuf:"varint,15,opt,name=mfaEnabled,proto3" json:"mfaEnabled,omitempty"`
	AppTheme                   string                 `protobuf:"bytes,16,opt,name=appTheme,proto3" json:"appTheme,omitempty"`
	MainView                   string                 `protobuf:"bytes,17,opt,name=mainView,proto3" json:"mainView,omitempty"`
	CostGroups                 []*MemberCostGroup     `protobuf:"bytes,18,rep,name=costGroups,proto3" json:"costGroups,omitempty"`
	CreatedBy                  *MemberUserData        `protobuf:"bytes,19,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedAt                  string                 `protobuf:"bytes,20,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	OrgId                      string                 `protobuf:"bytes,21,opt,name=orgId,proto3" json:"orgId,omitempty"`
	IsProfilingDone            bool                   `protobuf:"varint,22,opt,name=isProfilingDone,proto3" json:"isProfilingDone,omitempty"`
	PasswordLastRenewed        string                 `protobuf:"bytes,23,opt,name=passwordLastRenewed,proto3" json:"passwordLastRenewed,omitempty"`
	IsOwner                    bool                   `protobuf:"varint,24,opt,name=isOwner,proto3" json:"isOwner,omitempty"`
	IsAuth0                    bool                   `protobuf:"varint,25,opt,name=isAuth0,proto3" json:"isAuth0,omitempty"`
	ReadCostGroupCreationPopup bool                   `protobuf:"varint,26,opt,name=readCostGroupCreationPopup,proto3" json:"readCostGroupCreationPopup,omitempty"`
	UseNewCostGroupCreationUI  bool                   `protobuf:"varint,27,opt,name=useNewCostGroupCreationUI,proto3" json:"useNewCostGroupCreationUI,omitempty"`
	AutoTimeZone               bool                   `protobuf:"varint,28,opt,name=autoTimeZone,proto3" json:"autoTimeZone,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *UserData) Reset() {
	*x = UserData{}
	mi := &file_api_cover_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_api_cover_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserData) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *UserData) GetColorTheme() string {
	if x != nil {
		return x.ColorTheme
	}
	return ""
}

func (x *UserData) GetInitial() string {
	if x != nil {
		return x.Initial
	}
	return ""
}

func (x *UserData) GetActivated() bool {
	if x != nil {
		return x.Activated
	}
	return false
}

func (x *UserData) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *UserData) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *UserData) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *UserData) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *UserData) GetRegistered() string {
	if x != nil {
		return x.Registered
	}
	return ""
}

func (x *UserData) GetSsoEnabled() bool {
	if x != nil {
		return x.SsoEnabled
	}
	return false
}

func (x *UserData) GetMfaEnabled() bool {
	if x != nil {
		return x.MfaEnabled
	}
	return false
}

func (x *UserData) GetAppTheme() string {
	if x != nil {
		return x.AppTheme
	}
	return ""
}

func (x *UserData) GetMainView() string {
	if x != nil {
		return x.MainView
	}
	return ""
}

func (x *UserData) GetCostGroups() []*MemberCostGroup {
	if x != nil {
		return x.CostGroups
	}
	return nil
}

func (x *UserData) GetCreatedBy() *MemberUserData {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *UserData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *UserData) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *UserData) GetIsProfilingDone() bool {
	if x != nil {
		return x.IsProfilingDone
	}
	return false
}

func (x *UserData) GetPasswordLastRenewed() string {
	if x != nil {
		return x.PasswordLastRenewed
	}
	return ""
}

func (x *UserData) GetIsOwner() bool {
	if x != nil {
		return x.IsOwner
	}
	return false
}

func (x *UserData) GetIsAuth0() bool {
	if x != nil {
		return x.IsAuth0
	}
	return false
}

func (x *UserData) GetReadCostGroupCreationPopup() bool {
	if x != nil {
		return x.ReadCostGroupCreationPopup
	}
	return false
}

func (x *UserData) GetUseNewCostGroupCreationUI() bool {
	if x != nil {
		return x.UseNewCostGroupCreationUI
	}
	return false
}

func (x *UserData) GetAutoTimeZone() bool {
	if x != nil {
		return x.AutoTimeZone
	}
	return false
}

type MemberUserData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Avatar        string                 `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Icon          string                 `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	ColorTheme    string                 `protobuf:"bytes,6,opt,name=colorTheme,proto3" json:"colorTheme,omitempty"`
	Initial       string                 `protobuf:"bytes,7,opt,name=initial,proto3" json:"initial,omitempty"`
	Activated     bool                   `protobuf:"varint,8,opt,name=activated,proto3" json:"activated,omitempty"`
	IsAdmin       bool                   `protobuf:"varint,9,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
	Attributes    []string               `protobuf:"bytes,10,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Locale        string                 `protobuf:"bytes,11,opt,name=locale,proto3" json:"locale,omitempty"`
	TimeZone      string                 `protobuf:"bytes,12,opt,name=timeZone,proto3" json:"timeZone,omitempty"`
	Registered    string                 `protobuf:"bytes,13,opt,name=registered,proto3" json:"registered,omitempty"`
	SsoEnabled    bool                   `protobuf:"varint,14,opt,name=ssoEnabled,proto3" json:"ssoEnabled,omitempty"`
	MfaEnabled    bool                   `protobuf:"varint,15,opt,name=mfaEnabled,proto3" json:"mfaEnabled,omitempty"`
	AppTheme      string                 `protobuf:"bytes,16,opt,name=appTheme,proto3" json:"appTheme,omitempty"`
	MainView      string                 `protobuf:"bytes,17,opt,name=mainView,proto3" json:"mainView,omitempty"`
	CostGroups    []*MemberCostGroup     `protobuf:"bytes,18,rep,name=costGroups,proto3" json:"costGroups,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,19,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemberUserData) Reset() {
	*x = MemberUserData{}
	mi := &file_api_cover_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemberUserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberUserData) ProtoMessage() {}

func (x *MemberUserData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberUserData.ProtoReflect.Descriptor instead.
func (*MemberUserData) Descriptor() ([]byte, []int) {
	return file_api_cover_user_proto_rawDescGZIP(), []int{1}
}

func (x *MemberUserData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MemberUserData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MemberUserData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MemberUserData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *MemberUserData) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *MemberUserData) GetColorTheme() string {
	if x != nil {
		return x.ColorTheme
	}
	return ""
}

func (x *MemberUserData) GetInitial() string {
	if x != nil {
		return x.Initial
	}
	return ""
}

func (x *MemberUserData) GetActivated() bool {
	if x != nil {
		return x.Activated
	}
	return false
}

func (x *MemberUserData) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *MemberUserData) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *MemberUserData) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *MemberUserData) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *MemberUserData) GetRegistered() string {
	if x != nil {
		return x.Registered
	}
	return ""
}

func (x *MemberUserData) GetSsoEnabled() bool {
	if x != nil {
		return x.SsoEnabled
	}
	return false
}

func (x *MemberUserData) GetMfaEnabled() bool {
	if x != nil {
		return x.MfaEnabled
	}
	return false
}

func (x *MemberUserData) GetAppTheme() string {
	if x != nil {
		return x.AppTheme
	}
	return ""
}

func (x *MemberUserData) GetMainView() string {
	if x != nil {
		return x.MainView
	}
	return ""
}

func (x *MemberUserData) GetCostGroups() []*MemberCostGroup {
	if x != nil {
		return x.CostGroups
	}
	return nil
}

func (x *MemberUserData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type MemberCostGroup struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Image         string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Icon          string                 `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	ColorTheme    string                 `protobuf:"bytes,6,opt,name=colorTheme,proto3" json:"colorTheme,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemberCostGroup) Reset() {
	*x = MemberCostGroup{}
	mi := &file_api_cover_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemberCostGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCostGroup) ProtoMessage() {}

func (x *MemberCostGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCostGroup.ProtoReflect.Descriptor instead.
func (*MemberCostGroup) Descriptor() ([]byte, []int) {
	return file_api_cover_user_proto_rawDescGZIP(), []int{2}
}

func (x *MemberCostGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MemberCostGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MemberCostGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MemberCostGroup) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *MemberCostGroup) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *MemberCostGroup) GetColorTheme() string {
	if x != nil {
		return x.ColorTheme
	}
	return ""
}

func (x *MemberCostGroup) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MemberCostGroup) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_api_cover_user_proto protoreflect.FileDescriptor

var file_api_cover_user_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x22, 0xb9, 0x07, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f,
	0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f,
	0x6e, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x73, 0x6f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x73, 0x6f, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x66, 0x61, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6d, 0x66, 0x61, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x6f,
	0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x3f,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69,
	0x6e, 0x67, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x30, 0x0a,
	0x13, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6e,
	0x65, 0x77, 0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41,
	0x75, 0x74, 0x68, 0x30, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x75,
	0x74, 0x68, 0x30, 0x12, 0x3e, 0x0a, 0x1a, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x70, 0x75,
	0x70, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x73,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x70, 0x75, 0x70, 0x12, 0x3c, 0x0a, 0x19, 0x75, 0x73, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x73,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x49,
	0x18, 0x1b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x75, 0x73, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x6f,
	0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x49, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e,
	0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x54, 0x69, 0x6d,
	0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0xb6, 0x04, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x73, 0x6f, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x73, 0x6f, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x66, 0x61, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6d, 0x66, 0x61, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x54, 0x68, 0x65, 0x6d, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x54, 0x68, 0x65, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x12, 0x42, 0x0a, 0x0a,
	0x63, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdd,
	0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x68, 0x65, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x68, 0x65,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x62,
	0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x42, 0x11, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75,
	0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_cover_user_proto_rawDescOnce sync.Once
	file_api_cover_user_proto_rawDescData []byte
)

func file_api_cover_user_proto_rawDescGZIP() []byte {
	file_api_cover_user_proto_rawDescOnce.Do(func() {
		file_api_cover_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_cover_user_proto_rawDesc), len(file_api_cover_user_proto_rawDesc)))
	})
	return file_api_cover_user_proto_rawDescData
}

var file_api_cover_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_cover_user_proto_goTypes = []any{
	(*UserData)(nil),        // 0: blueapi.api.cover.UserData
	(*MemberUserData)(nil),  // 1: blueapi.api.cover.MemberUserData
	(*MemberCostGroup)(nil), // 2: blueapi.api.cover.MemberCostGroup
}
var file_api_cover_user_proto_depIdxs = []int32{
	2, // 0: blueapi.api.cover.UserData.costGroups:type_name -> blueapi.api.cover.MemberCostGroup
	1, // 1: blueapi.api.cover.UserData.createdBy:type_name -> blueapi.api.cover.MemberUserData
	2, // 2: blueapi.api.cover.MemberUserData.costGroups:type_name -> blueapi.api.cover.MemberCostGroup
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_cover_user_proto_init() }
func file_api_cover_user_proto_init() {
	if File_api_cover_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_cover_user_proto_rawDesc), len(file_api_cover_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_user_proto_goTypes,
		DependencyIndexes: file_api_cover_user_proto_depIdxs,
		MessageInfos:      file_api_cover_user_proto_msgTypes,
	}.Build()
	File_api_cover_user_proto = out.File
	file_api_cover_user_proto_goTypes = nil
	file_api_cover_user_proto_depIdxs = nil
}
