// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: api/cover/view.proto

package cover

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ViewData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsPrivate   bool            `protobuf:"varint,4,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
	IsEditable  bool            `protobuf:"varint,5,opt,name=isEditable,proto3" json:"isEditable,omitempty"`
	Icon        string          `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	CreatedBy   *MemberUserData `protobuf:"bytes,7,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	CreatedAt   string          `protobuf:"bytes,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedBy   *MemberUserData `protobuf:"bytes,9,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	UpdatedAt   string          `protobuf:"bytes,10,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Layout      []*WidgetData   `protobuf:"bytes,11,rep,name=layout,proto3" json:"layout,omitempty"`
	SideMenu    *SideMenu       `protobuf:"bytes,12,opt,name=sideMenu,proto3" json:"sideMenu,omitempty"`
	ReportType  string          `protobuf:"bytes,13,opt,name=reportType,proto3" json:"reportType,omitempty"`
}

func (x *ViewData) Reset() {
	*x = ViewData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewData) ProtoMessage() {}

func (x *ViewData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewData.ProtoReflect.Descriptor instead.
func (*ViewData) Descriptor() ([]byte, []int) {
	return file_api_cover_view_proto_rawDescGZIP(), []int{0}
}

func (x *ViewData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ViewData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ViewData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ViewData) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *ViewData) GetIsEditable() bool {
	if x != nil {
		return x.IsEditable
	}
	return false
}

func (x *ViewData) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ViewData) GetCreatedBy() *MemberUserData {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *ViewData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ViewData) GetUpdatedBy() *MemberUserData {
	if x != nil {
		return x.UpdatedBy
	}
	return nil
}

func (x *ViewData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ViewData) GetLayout() []*WidgetData {
	if x != nil {
		return x.Layout
	}
	return nil
}

func (x *ViewData) GetSideMenu() *SideMenu {
	if x != nil {
		return x.SideMenu
	}
	return nil
}

func (x *ViewData) GetReportType() string {
	if x != nil {
		return x.ReportType
	}
	return ""
}

type ViewList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsPrivate   bool   `protobuf:"varint,4,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
	IsEditable  bool   `protobuf:"varint,5,opt,name=isEditable,proto3" json:"isEditable,omitempty"`
	Icon        string `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ViewList) Reset() {
	*x = ViewList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_view_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewList) ProtoMessage() {}

func (x *ViewList) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_view_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewList.ProtoReflect.Descriptor instead.
func (*ViewList) Descriptor() ([]byte, []int) {
	return file_api_cover_view_proto_rawDescGZIP(), []int{1}
}

func (x *ViewList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ViewList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ViewList) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ViewList) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *ViewList) GetIsEditable() bool {
	if x != nil {
		return x.IsEditable
	}
	return false
}

func (x *ViewList) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ViewList) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ViewList) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ViewLayout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WidgetId   string      `protobuf:"bytes,1,opt,name=widgetId,proto3" json:"widgetId,omitempty"`
	WidgetData *WidgetData `protobuf:"bytes,2,opt,name=widgetData,proto3" json:"widgetData,omitempty"`
}

func (x *ViewLayout) Reset() {
	*x = ViewLayout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_view_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewLayout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewLayout) ProtoMessage() {}

func (x *ViewLayout) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_view_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewLayout.ProtoReflect.Descriptor instead.
func (*ViewLayout) Descriptor() ([]byte, []int) {
	return file_api_cover_view_proto_rawDescGZIP(), []int{2}
}

func (x *ViewLayout) GetWidgetId() string {
	if x != nil {
		return x.WidgetId
	}
	return ""
}

func (x *ViewLayout) GetWidgetData() *WidgetData {
	if x != nil {
		return x.WidgetData
	}
	return nil
}

type WidgetData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	X           float64           `protobuf:"fixed64,2,opt,name=x,proto3" json:"x,omitempty"`
	Y           float64           `protobuf:"fixed64,3,opt,name=y,proto3" json:"y,omitempty"`
	ComponentId string            `protobuf:"bytes,4,opt,name=componentId,proto3" json:"componentId,omitempty"`
	Options     *LayoutOptions    `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	Requests    []*LayoutRequests `protobuf:"bytes,6,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *WidgetData) Reset() {
	*x = WidgetData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_view_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetData) ProtoMessage() {}

func (x *WidgetData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_view_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetData.ProtoReflect.Descriptor instead.
func (*WidgetData) Descriptor() ([]byte, []int) {
	return file_api_cover_view_proto_rawDescGZIP(), []int{3}
}

func (x *WidgetData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WidgetData) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *WidgetData) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *WidgetData) GetComponentId() string {
	if x != nil {
		return x.ComponentId
	}
	return ""
}

func (x *WidgetData) GetOptions() *LayoutOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *WidgetData) GetRequests() []*LayoutRequests {
	if x != nil {
		return x.Requests
	}
	return nil
}

type SideMenu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Favorite           []string `protobuf:"bytes,1,rep,name=favorite,proto3" json:"favorite,omitempty"`
	IsOpenedAdmin      bool     `protobuf:"varint,2,opt,name=isOpenedAdmin,proto3" json:"isOpenedAdmin,omitempty"`
	IsOpenedFeatures   bool     `protobuf:"varint,3,opt,name=isOpenedFeatures,proto3" json:"isOpenedFeatures,omitempty"`
	IsOpenedCostGroups bool     `protobuf:"varint,4,opt,name=isOpenedCostGroups,proto3" json:"isOpenedCostGroups,omitempty"`
}

func (x *SideMenu) Reset() {
	*x = SideMenu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_view_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SideMenu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SideMenu) ProtoMessage() {}

func (x *SideMenu) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_view_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SideMenu.ProtoReflect.Descriptor instead.
func (*SideMenu) Descriptor() ([]byte, []int) {
	return file_api_cover_view_proto_rawDescGZIP(), []int{4}
}

func (x *SideMenu) GetFavorite() []string {
	if x != nil {
		return x.Favorite
	}
	return nil
}

func (x *SideMenu) GetIsOpenedAdmin() bool {
	if x != nil {
		return x.IsOpenedAdmin
	}
	return false
}

func (x *SideMenu) GetIsOpenedFeatures() bool {
	if x != nil {
		return x.IsOpenedFeatures
	}
	return false
}

func (x *SideMenu) GetIsOpenedCostGroups() bool {
	if x != nil {
		return x.IsOpenedCostGroups
	}
	return false
}

type LayoutOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *LayoutOptions) Reset() {
	*x = LayoutOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_view_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LayoutOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LayoutOptions) ProtoMessage() {}

func (x *LayoutOptions) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_view_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LayoutOptions.ProtoReflect.Descriptor instead.
func (*LayoutOptions) Descriptor() ([]byte, []int) {
	return file_api_cover_view_proto_rawDescGZIP(), []int{5}
}

func (x *LayoutOptions) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *LayoutOptions) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type LayoutRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url    string           `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Params *structpb.Struct `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
	Hash   string           `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *LayoutRequests) Reset() {
	*x = LayoutRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_view_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LayoutRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LayoutRequests) ProtoMessage() {}

func (x *LayoutRequests) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_view_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LayoutRequests.ProtoReflect.Descriptor instead.
func (*LayoutRequests) Descriptor() ([]byte, []int) {
	return file_api_cover_view_proto_rawDescGZIP(), []int{6}
}

func (x *LayoutRequests) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LayoutRequests) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LayoutRequests) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LayoutRequests) GetParams() *structpb.Struct {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *LayoutRequests) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type Favorites struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsPrivate   bool   `protobuf:"varint,4,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
	Icon        string `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Favorites) Reset() {
	*x = Favorites{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_view_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Favorites) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Favorites) ProtoMessage() {}

func (x *Favorites) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_view_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Favorites.ProtoReflect.Descriptor instead.
func (*Favorites) Descriptor() ([]byte, []int) {
	return file_api_cover_view_proto_rawDescGZIP(), []int{7}
}

func (x *Favorites) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Favorites) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Favorites) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Favorites) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *Favorites) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Favorites) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Favorites) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SideMenuState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value bool   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SideMenuState) Reset() {
	*x = SideMenuState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cover_view_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SideMenuState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SideMenuState) ProtoMessage() {}

func (x *SideMenuState) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_view_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SideMenuState.ProtoReflect.Descriptor instead.
func (*SideMenuState) Descriptor() ([]byte, []int) {
	return file_api_cover_view_proto_rawDescGZIP(), []int{8}
}

func (x *SideMenuState) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SideMenuState) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

var File_api_cover_view_proto protoreflect.FileDescriptor

var file_api_cover_view_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x69, 0x65, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x1a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x03,
	0x0a, 0x08, 0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x73, 0x45, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3f, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x35, 0x0a, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x73, 0x69, 0x64, 0x65, 0x4d,
	0x65, 0x6e, 0x75, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x69,
	0x64, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x08, 0x73, 0x69, 0x64, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xde, 0x01, 0x0a, 0x08, 0x56, 0x69, 0x65, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x45, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x67, 0x0a, 0x0a, 0x56, 0x69, 0x65, 0x77, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0a, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a,
	0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0xd5, 0x01, 0x0a, 0x0a, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x08, 0x53, 0x69, 0x64, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69,
	0x73, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x4f,
	0x70, 0x65, 0x6e, 0x65, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x2e, 0x0a,
	0x12, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x4f, 0x70, 0x65,
	0x6e, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x47, 0x0a,
	0x0d, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x4c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x2f, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x22, 0xbf, 0x01, 0x0a, 0x09, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x37, 0x0a, 0x0d, 0x53, 0x69, 0x64, 0x65, 0x4d, 0x65,
	0x6e, 0x75, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x62, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x42, 0x11, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c,
	0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cover_view_proto_rawDescOnce sync.Once
	file_api_cover_view_proto_rawDescData = file_api_cover_view_proto_rawDesc
)

func file_api_cover_view_proto_rawDescGZIP() []byte {
	file_api_cover_view_proto_rawDescOnce.Do(func() {
		file_api_cover_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_view_proto_rawDescData)
	})
	return file_api_cover_view_proto_rawDescData
}

var file_api_cover_view_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_cover_view_proto_goTypes = []interface{}{
	(*ViewData)(nil),        // 0: blueapi.api.cover.ViewData
	(*ViewList)(nil),        // 1: blueapi.api.cover.ViewList
	(*ViewLayout)(nil),      // 2: blueapi.api.cover.ViewLayout
	(*WidgetData)(nil),      // 3: blueapi.api.cover.WidgetData
	(*SideMenu)(nil),        // 4: blueapi.api.cover.SideMenu
	(*LayoutOptions)(nil),   // 5: blueapi.api.cover.LayoutOptions
	(*LayoutRequests)(nil),  // 6: blueapi.api.cover.LayoutRequests
	(*Favorites)(nil),       // 7: blueapi.api.cover.Favorites
	(*SideMenuState)(nil),   // 8: blueapi.api.cover.SideMenuState
	(*MemberUserData)(nil),  // 9: blueapi.api.cover.MemberUserData
	(*structpb.Struct)(nil), // 10: google.protobuf.Struct
}
var file_api_cover_view_proto_depIdxs = []int32{
	9,  // 0: blueapi.api.cover.ViewData.createdBy:type_name -> blueapi.api.cover.MemberUserData
	9,  // 1: blueapi.api.cover.ViewData.updatedBy:type_name -> blueapi.api.cover.MemberUserData
	3,  // 2: blueapi.api.cover.ViewData.layout:type_name -> blueapi.api.cover.WidgetData
	4,  // 3: blueapi.api.cover.ViewData.sideMenu:type_name -> blueapi.api.cover.SideMenu
	3,  // 4: blueapi.api.cover.ViewLayout.widgetData:type_name -> blueapi.api.cover.WidgetData
	5,  // 5: blueapi.api.cover.WidgetData.options:type_name -> blueapi.api.cover.LayoutOptions
	6,  // 6: blueapi.api.cover.WidgetData.requests:type_name -> blueapi.api.cover.LayoutRequests
	10, // 7: blueapi.api.cover.LayoutRequests.params:type_name -> google.protobuf.Struct
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_cover_view_proto_init() }
func file_api_cover_view_proto_init() {
	if File_api_cover_view_proto != nil {
		return
	}
	file_api_cover_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_cover_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewData); i {
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
		file_api_cover_view_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewList); i {
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
		file_api_cover_view_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewLayout); i {
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
		file_api_cover_view_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetData); i {
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
		file_api_cover_view_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SideMenu); i {
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
		file_api_cover_view_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LayoutOptions); i {
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
		file_api_cover_view_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LayoutRequests); i {
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
		file_api_cover_view_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Favorites); i {
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
		file_api_cover_view_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SideMenuState); i {
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
			RawDescriptor: file_api_cover_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_view_proto_goTypes,
		DependencyIndexes: file_api_cover_view_proto_depIdxs,
		MessageInfos:      file_api_cover_view_proto_msgTypes,
	}.Build()
	File_api_cover_view_proto = out.File
	file_api_cover_view_proto_rawDesc = nil
	file_api_cover_view_proto_goTypes = nil
	file_api_cover_view_proto_depIdxs = nil
}
