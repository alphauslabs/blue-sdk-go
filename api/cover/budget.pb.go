// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: api/cover/budget.proto

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

type BudgetData struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	Id             string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`               // not required for creating budget
	StartDate      string                  `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"` // use yyyymmdd format
	EndDate        string                  `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`     // use yyyymmdd format
	TotalBudget    float32                 `protobuf:"fixed32,4,opt,name=totalBudget,proto3" json:"totalBudget,omitempty"`
	Period         int64                   `protobuf:"varint,5,opt,name=period,proto3" json:"period,omitempty"`                // 3, 6, 12 months
	GranularBudget []*GranularBudgetData   `protobuf:"bytes,6,rep,name=granularBudget,proto3" json:"granularBudget,omitempty"` // budget per month
	CostGroup      *AlertCostGroup         `protobuf:"bytes,7,opt,name=costGroup,proto3" json:"costGroup,omitempty"`
	Alerts         []*BudgetAlert          `protobuf:"bytes,8,rep,name=alerts,proto3" json:"alerts,omitempty"`                 // threshold(s) and its respective channel(s) to alert
	ForecastedData []*GranularForecastData `protobuf:"bytes,9,rep,name=forecastedData,proto3" json:"forecastedData,omitempty"` // forecast for ongoing months of an active budget; if expired, archived forecast record
	SpendingData   []*GranularSpendingData `protobuf:"bytes,10,rep,name=spendingData,proto3" json:"spendingData,omitempty"`    // spending data
	Expired        bool                    `protobuf:"varint,11,opt,name=expired,proto3" json:"expired,omitempty"`             // true if budget has expired
	Draft          bool                    `protobuf:"varint,12,opt,name=draft,proto3" json:"draft,omitempty"`                 // true if the current budget set is still a draft
	CreatedBy      string                  `protobuf:"bytes,13,opt,name=createdBy,proto3" json:"createdBy,omitempty"`          // not required for creating budget
	CreatedAt      string                  `protobuf:"bytes,14,opt,name=createdAt,proto3" json:"createdAt,omitempty"`          // not required for creating budget
	UpdatedBy      string                  `protobuf:"bytes,15,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`          // not required for creating budget
	UpdatedAt      string                  `protobuf:"bytes,16,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`          // not required for creating budget
	TotalSpend     float32                 `protobuf:"fixed32,17,opt,name=totalSpend,proto3" json:"totalSpend,omitempty"`      // total accumulated spending within the budget period
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BudgetData) Reset() {
	*x = BudgetData{}
	mi := &file_api_cover_budget_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BudgetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetData) ProtoMessage() {}

func (x *BudgetData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_budget_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetData.ProtoReflect.Descriptor instead.
func (*BudgetData) Descriptor() ([]byte, []int) {
	return file_api_cover_budget_proto_rawDescGZIP(), []int{0}
}

func (x *BudgetData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BudgetData) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *BudgetData) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *BudgetData) GetTotalBudget() float32 {
	if x != nil {
		return x.TotalBudget
	}
	return 0
}

func (x *BudgetData) GetPeriod() int64 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *BudgetData) GetGranularBudget() []*GranularBudgetData {
	if x != nil {
		return x.GranularBudget
	}
	return nil
}

func (x *BudgetData) GetCostGroup() *AlertCostGroup {
	if x != nil {
		return x.CostGroup
	}
	return nil
}

func (x *BudgetData) GetAlerts() []*BudgetAlert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

func (x *BudgetData) GetForecastedData() []*GranularForecastData {
	if x != nil {
		return x.ForecastedData
	}
	return nil
}

func (x *BudgetData) GetSpendingData() []*GranularSpendingData {
	if x != nil {
		return x.SpendingData
	}
	return nil
}

func (x *BudgetData) GetExpired() bool {
	if x != nil {
		return x.Expired
	}
	return false
}

func (x *BudgetData) GetDraft() bool {
	if x != nil {
		return x.Draft
	}
	return false
}

func (x *BudgetData) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *BudgetData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BudgetData) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *BudgetData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BudgetData) GetTotalSpend() float32 {
	if x != nil {
		return x.TotalSpend
	}
	return 0
}

type GranularBudgetData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Date          string                 `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"` // use yyyymm format
	Budget        float32                `protobuf:"fixed32,2,opt,name=budget,proto3" json:"budget,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GranularBudgetData) Reset() {
	*x = GranularBudgetData{}
	mi := &file_api_cover_budget_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GranularBudgetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GranularBudgetData) ProtoMessage() {}

func (x *GranularBudgetData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_budget_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GranularBudgetData.ProtoReflect.Descriptor instead.
func (*GranularBudgetData) Descriptor() ([]byte, []int) {
	return file_api_cover_budget_proto_rawDescGZIP(), []int{1}
}

func (x *GranularBudgetData) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *GranularBudgetData) GetBudget() float32 {
	if x != nil {
		return x.Budget
	}
	return 0
}

type GranularForecastData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Date          string                 `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"` // use yyyymm format
	Mid           float32                `protobuf:"fixed32,2,opt,name=mid,proto3" json:"mid,omitempty"`
	UpperBound    float32                `protobuf:"fixed32,3,opt,name=upperBound,proto3" json:"upperBound,omitempty"`
	LowerBound    float32                `protobuf:"fixed32,4,opt,name=lowerBound,proto3" json:"lowerBound,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GranularForecastData) Reset() {
	*x = GranularForecastData{}
	mi := &file_api_cover_budget_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GranularForecastData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GranularForecastData) ProtoMessage() {}

func (x *GranularForecastData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_budget_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GranularForecastData.ProtoReflect.Descriptor instead.
func (*GranularForecastData) Descriptor() ([]byte, []int) {
	return file_api_cover_budget_proto_rawDescGZIP(), []int{2}
}

func (x *GranularForecastData) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *GranularForecastData) GetMid() float32 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *GranularForecastData) GetUpperBound() float32 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

func (x *GranularForecastData) GetLowerBound() float32 {
	if x != nil {
		return x.LowerBound
	}
	return 0
}

type GranularSpendingData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Date          string                 `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`     // use yyyymm format
	Value         float32                `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"` // actual cost for the month
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GranularSpendingData) Reset() {
	*x = GranularSpendingData{}
	mi := &file_api_cover_budget_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GranularSpendingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GranularSpendingData) ProtoMessage() {}

func (x *GranularSpendingData) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_budget_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GranularSpendingData.ProtoReflect.Descriptor instead.
func (*GranularSpendingData) Descriptor() ([]byte, []int) {
	return file_api_cover_budget_proto_rawDescGZIP(), []int{3}
}

func (x *GranularSpendingData) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *GranularSpendingData) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type BudgetAlert struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Threshold     []*Threshold           `protobuf:"bytes,1,rep,name=threshold,proto3" json:"threshold,omitempty"`
	Channels      *AlertChannels         `protobuf:"bytes,2,opt,name=channels,proto3" json:"channels,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BudgetAlert) Reset() {
	*x = BudgetAlert{}
	mi := &file_api_cover_budget_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BudgetAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetAlert) ProtoMessage() {}

func (x *BudgetAlert) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_budget_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetAlert.ProtoReflect.Descriptor instead.
func (*BudgetAlert) Descriptor() ([]byte, []int) {
	return file_api_cover_budget_proto_rawDescGZIP(), []int{4}
}

func (x *BudgetAlert) GetThreshold() []*Threshold {
	if x != nil {
		return x.Threshold
	}
	return nil
}

func (x *BudgetAlert) GetChannels() *AlertChannels {
	if x != nil {
		return x.Channels
	}
	return nil
}

type Threshold struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`     // exact or percentage
	Value         float32                `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"` // actual value of threshold
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Threshold) Reset() {
	*x = Threshold{}
	mi := &file_api_cover_budget_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Threshold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Threshold) ProtoMessage() {}

func (x *Threshold) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_budget_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Threshold.ProtoReflect.Descriptor instead.
func (*Threshold) Descriptor() ([]byte, []int) {
	return file_api_cover_budget_proto_rawDescGZIP(), []int{5}
}

func (x *Threshold) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Threshold) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_api_cover_budget_proto protoreflect.FileDescriptor

var file_api_cover_budget_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x1a, 0x15, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbc, 0x05, 0x0a, 0x0a, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x4d, 0x0a, 0x0e, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e,
	0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0e, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43,
	0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x36, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x0e, 0x66,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72,
	0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x66, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4b, 0x0a, 0x0c,
	0x73, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x53,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x73, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x6e,
	0x64, 0x22, 0x40, 0x0a, 0x12, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x22, 0x7c, 0x0a, 0x14, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x22, 0x40, 0x0a, 0x14, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x53, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12,
	0x3c, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x73, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x22, 0x35, 0x0a,
	0x09, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x64, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x13, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_cover_budget_proto_rawDescOnce sync.Once
	file_api_cover_budget_proto_rawDescData = file_api_cover_budget_proto_rawDesc
)

func file_api_cover_budget_proto_rawDescGZIP() []byte {
	file_api_cover_budget_proto_rawDescOnce.Do(func() {
		file_api_cover_budget_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_budget_proto_rawDescData)
	})
	return file_api_cover_budget_proto_rawDescData
}

var file_api_cover_budget_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_cover_budget_proto_goTypes = []any{
	(*BudgetData)(nil),           // 0: blueapi.api.cover.BudgetData
	(*GranularBudgetData)(nil),   // 1: blueapi.api.cover.GranularBudgetData
	(*GranularForecastData)(nil), // 2: blueapi.api.cover.GranularForecastData
	(*GranularSpendingData)(nil), // 3: blueapi.api.cover.GranularSpendingData
	(*BudgetAlert)(nil),          // 4: blueapi.api.cover.BudgetAlert
	(*Threshold)(nil),            // 5: blueapi.api.cover.Threshold
	(*AlertCostGroup)(nil),       // 6: blueapi.api.cover.AlertCostGroup
	(*AlertChannels)(nil),        // 7: blueapi.api.cover.AlertChannels
}
var file_api_cover_budget_proto_depIdxs = []int32{
	1, // 0: blueapi.api.cover.BudgetData.granularBudget:type_name -> blueapi.api.cover.GranularBudgetData
	6, // 1: blueapi.api.cover.BudgetData.costGroup:type_name -> blueapi.api.cover.AlertCostGroup
	4, // 2: blueapi.api.cover.BudgetData.alerts:type_name -> blueapi.api.cover.BudgetAlert
	2, // 3: blueapi.api.cover.BudgetData.forecastedData:type_name -> blueapi.api.cover.GranularForecastData
	3, // 4: blueapi.api.cover.BudgetData.spendingData:type_name -> blueapi.api.cover.GranularSpendingData
	5, // 5: blueapi.api.cover.BudgetAlert.threshold:type_name -> blueapi.api.cover.Threshold
	7, // 6: blueapi.api.cover.BudgetAlert.channels:type_name -> blueapi.api.cover.AlertChannels
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_cover_budget_proto_init() }
func file_api_cover_budget_proto_init() {
	if File_api_cover_budget_proto != nil {
		return
	}
	file_api_cover_alert_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cover_budget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_budget_proto_goTypes,
		DependencyIndexes: file_api_cover_budget_proto_depIdxs,
		MessageInfos:      file_api_cover_budget_proto_msgTypes,
	}.Build()
	File_api_cover_budget_proto = out.File
	file_api_cover_budget_proto_rawDesc = nil
	file_api_cover_budget_proto_goTypes = nil
	file_api_cover_budget_proto_depIdxs = nil
}
