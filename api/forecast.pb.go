// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: api/forecast.proto

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

type ForecastData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId       string  `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	ProductCategory string  `protobuf:"bytes,2,opt,name=productCategory,proto3" json:"productCategory,omitempty"`
	ProductCode     string  `protobuf:"bytes,3,opt,name=productCode,proto3" json:"productCode,omitempty"`
	Frequency       string  `protobuf:"bytes,4,opt,name=frequency,proto3" json:"frequency,omitempty"` //daily, monthly
	Date            string  `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	HistoricalCost  float64 `protobuf:"fixed64,6,opt,name=historicalCost,proto3" json:"historicalCost,omitempty"`
	ForecastedCost  float64 `protobuf:"fixed64,7,opt,name=forecastedCost,proto3" json:"forecastedCost,omitempty"`
	UpperBound      float64 `protobuf:"fixed64,8,opt,name=upperBound,proto3" json:"upperBound,omitempty"`
	LowerBound      float64 `protobuf:"fixed64,9,opt,name=lowerBound,proto3" json:"lowerBound,omitempty"`
}

func (x *ForecastData) Reset() {
	*x = ForecastData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecastData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecastData) ProtoMessage() {}

func (x *ForecastData) ProtoReflect() protoreflect.Message {
	mi := &file_api_forecast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecastData.ProtoReflect.Descriptor instead.
func (*ForecastData) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{0}
}

func (x *ForecastData) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ForecastData) GetProductCategory() string {
	if x != nil {
		return x.ProductCategory
	}
	return ""
}

func (x *ForecastData) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *ForecastData) GetFrequency() string {
	if x != nil {
		return x.Frequency
	}
	return ""
}

func (x *ForecastData) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ForecastData) GetHistoricalCost() float64 {
	if x != nil {
		return x.HistoricalCost
	}
	return 0
}

func (x *ForecastData) GetForecastedCost() float64 {
	if x != nil {
		return x.ForecastedCost
	}
	return 0
}

func (x *ForecastData) GetUpperBound() float64 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

func (x *ForecastData) GetLowerBound() float64 {
	if x != nil {
		return x.LowerBound
	}
	return 0
}

type AccountGroupForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string          `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Data    []*ForecastData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AccountGroupForecast) Reset() {
	*x = AccountGroupForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountGroupForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountGroupForecast) ProtoMessage() {}

func (x *AccountGroupForecast) ProtoReflect() protoreflect.Message {
	mi := &file_api_forecast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountGroupForecast.ProtoReflect.Descriptor instead.
func (*AccountGroupForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{1}
}

func (x *AccountGroupForecast) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *AccountGroupForecast) GetData() []*ForecastData {
	if x != nil {
		return x.Data
	}
	return nil
}

type BillingGroupForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingInternalId string          `protobuf:"bytes,1,opt,name=billingInternalId,proto3" json:"billingInternalId,omitempty"`
	Data              []*ForecastData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BillingGroupForecast) Reset() {
	*x = BillingGroupForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingGroupForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingGroupForecast) ProtoMessage() {}

func (x *BillingGroupForecast) ProtoReflect() protoreflect.Message {
	mi := &file_api_forecast_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingGroupForecast.ProtoReflect.Descriptor instead.
func (*BillingGroupForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{2}
}

func (x *BillingGroupForecast) GetBillingInternalId() string {
	if x != nil {
		return x.BillingInternalId
	}
	return ""
}

func (x *BillingGroupForecast) GetData() []*ForecastData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrgForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId string                  `protobuf:"bytes,1,opt,name=orgId,proto3" json:"orgId,omitempty"`
	Data  []*BillingGroupForecast `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OrgForecast) Reset() {
	*x = OrgForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgForecast) ProtoMessage() {}

func (x *OrgForecast) ProtoReflect() protoreflect.Message {
	mi := &file_api_forecast_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgForecast.ProtoReflect.Descriptor instead.
func (*OrgForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{3}
}

func (x *OrgForecast) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *OrgForecast) GetData() []*BillingGroupForecast {
	if x != nil {
		return x.Data
	}
	return nil
}

type MonthToDateForecastData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId       string  `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	ProductCategory string  `protobuf:"bytes,2,opt,name=productCategory,proto3" json:"productCategory,omitempty"`
	ProductCode     string  `protobuf:"bytes,3,opt,name=productCode,proto3" json:"productCode,omitempty"`
	Date            string  `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	AccumulatedCost float64 `protobuf:"fixed64,5,opt,name=accumulatedCost,proto3" json:"accumulatedCost,omitempty"`
	ForecastCost    float64 `protobuf:"fixed64,6,opt,name=forecastCost,proto3" json:"forecastCost,omitempty"`
	Budget          float64 `protobuf:"fixed64,7,opt,name=budget,proto3" json:"budget,omitempty"`
}

func (x *MonthToDateForecastData) Reset() {
	*x = MonthToDateForecastData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonthToDateForecastData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthToDateForecastData) ProtoMessage() {}

func (x *MonthToDateForecastData) ProtoReflect() protoreflect.Message {
	mi := &file_api_forecast_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthToDateForecastData.ProtoReflect.Descriptor instead.
func (*MonthToDateForecastData) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{4}
}

func (x *MonthToDateForecastData) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *MonthToDateForecastData) GetProductCategory() string {
	if x != nil {
		return x.ProductCategory
	}
	return ""
}

func (x *MonthToDateForecastData) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *MonthToDateForecastData) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *MonthToDateForecastData) GetAccumulatedCost() float64 {
	if x != nil {
		return x.AccumulatedCost
	}
	return 0
}

func (x *MonthToDateForecastData) GetForecastCost() float64 {
	if x != nil {
		return x.ForecastCost
	}
	return 0
}

func (x *MonthToDateForecastData) GetBudget() float64 {
	if x != nil {
		return x.Budget
	}
	return 0
}

type BillingGroupMonthToDateForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingInternalId string                     `protobuf:"bytes,1,opt,name=billingInternalId,proto3" json:"billingInternalId,omitempty"`
	Data              []*MonthToDateForecastData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BillingGroupMonthToDateForecast) Reset() {
	*x = BillingGroupMonthToDateForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingGroupMonthToDateForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingGroupMonthToDateForecast) ProtoMessage() {}

func (x *BillingGroupMonthToDateForecast) ProtoReflect() protoreflect.Message {
	mi := &file_api_forecast_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingGroupMonthToDateForecast.ProtoReflect.Descriptor instead.
func (*BillingGroupMonthToDateForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{5}
}

func (x *BillingGroupMonthToDateForecast) GetBillingInternalId() string {
	if x != nil {
		return x.BillingInternalId
	}
	return ""
}

func (x *BillingGroupMonthToDateForecast) GetData() []*MonthToDateForecastData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrgMonthToDateForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId string                             `protobuf:"bytes,1,opt,name=orgId,proto3" json:"orgId,omitempty"`
	Data  []*BillingGroupMonthToDateForecast `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OrgMonthToDateForecast) Reset() {
	*x = OrgMonthToDateForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgMonthToDateForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgMonthToDateForecast) ProtoMessage() {}

func (x *OrgMonthToDateForecast) ProtoReflect() protoreflect.Message {
	mi := &file_api_forecast_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgMonthToDateForecast.ProtoReflect.Descriptor instead.
func (*OrgMonthToDateForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{6}
}

func (x *OrgMonthToDateForecast) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *OrgMonthToDateForecast) GetData() []*BillingGroupMonthToDateForecast {
	if x != nil {
		return x.Data
	}
	return nil
}

type MonthlyCostForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date         string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	CostActual   float64 `protobuf:"fixed64,2,opt,name=costActual,proto3" json:"costActual,omitempty"`
	CostForecast float64 `protobuf:"fixed64,3,opt,name=costForecast,proto3" json:"costForecast,omitempty"`
	Budget       float64 `protobuf:"fixed64,4,opt,name=budget,proto3" json:"budget,omitempty"`
	UpperBound   float64 `protobuf:"fixed64,5,opt,name=upperBound,proto3" json:"upperBound,omitempty"`
	LowerBound   float64 `protobuf:"fixed64,6,opt,name=lowerBound,proto3" json:"lowerBound,omitempty"`
}

func (x *MonthlyCostForecast) Reset() {
	*x = MonthlyCostForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonthlyCostForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthlyCostForecast) ProtoMessage() {}

func (x *MonthlyCostForecast) ProtoReflect() protoreflect.Message {
	mi := &file_api_forecast_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthlyCostForecast.ProtoReflect.Descriptor instead.
func (*MonthlyCostForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{7}
}

func (x *MonthlyCostForecast) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *MonthlyCostForecast) GetCostActual() float64 {
	if x != nil {
		return x.CostActual
	}
	return 0
}

func (x *MonthlyCostForecast) GetCostForecast() float64 {
	if x != nil {
		return x.CostForecast
	}
	return 0
}

func (x *MonthlyCostForecast) GetBudget() float64 {
	if x != nil {
		return x.Budget
	}
	return 0
}

func (x *MonthlyCostForecast) GetUpperBound() float64 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

func (x *MonthlyCostForecast) GetLowerBound() float64 {
	if x != nil {
		return x.LowerBound
	}
	return 0
}

type MonthOnMonthCostForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category     string  `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	CostCurrent  float64 `protobuf:"fixed64,2,opt,name=costCurrent,proto3" json:"costCurrent,omitempty"`
	CostPrev     float64 `protobuf:"fixed64,3,opt,name=costPrev,proto3" json:"costPrev,omitempty"`
	CostForecast float64 `protobuf:"fixed64,4,opt,name=costForecast,proto3" json:"costForecast,omitempty"`
}

func (x *MonthOnMonthCostForecast) Reset() {
	*x = MonthOnMonthCostForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonthOnMonthCostForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthOnMonthCostForecast) ProtoMessage() {}

func (x *MonthOnMonthCostForecast) ProtoReflect() protoreflect.Message {
	mi := &file_api_forecast_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthOnMonthCostForecast.ProtoReflect.Descriptor instead.
func (*MonthOnMonthCostForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{8}
}

func (x *MonthOnMonthCostForecast) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *MonthOnMonthCostForecast) GetCostCurrent() float64 {
	if x != nil {
		return x.CostCurrent
	}
	return 0
}

func (x *MonthOnMonthCostForecast) GetCostPrev() float64 {
	if x != nil {
		return x.CostPrev
	}
	return 0
}

func (x *MonthOnMonthCostForecast) GetCostForecast() float64 {
	if x != nil {
		return x.CostForecast
	}
	return 0
}

type MonthToDateCostForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date            string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	CostPrev        float64 `protobuf:"fixed64,2,opt,name=costPrev,proto3" json:"costPrev,omitempty"`
	CostAccumulated float64 `protobuf:"fixed64,3,opt,name=costAccumulated,proto3" json:"costAccumulated,omitempty"`
	CostForecast    float64 `protobuf:"fixed64,4,opt,name=costForecast,proto3" json:"costForecast,omitempty"`
	UpperBound      float64 `protobuf:"fixed64,5,opt,name=upperBound,proto3" json:"upperBound,omitempty"`
	LowerBound      float64 `protobuf:"fixed64,6,opt,name=lowerBound,proto3" json:"lowerBound,omitempty"`
}

func (x *MonthToDateCostForecast) Reset() {
	*x = MonthToDateCostForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonthToDateCostForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthToDateCostForecast) ProtoMessage() {}

func (x *MonthToDateCostForecast) ProtoReflect() protoreflect.Message {
	mi := &file_api_forecast_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthToDateCostForecast.ProtoReflect.Descriptor instead.
func (*MonthToDateCostForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{9}
}

func (x *MonthToDateCostForecast) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *MonthToDateCostForecast) GetCostPrev() float64 {
	if x != nil {
		return x.CostPrev
	}
	return 0
}

func (x *MonthToDateCostForecast) GetCostAccumulated() float64 {
	if x != nil {
		return x.CostAccumulated
	}
	return 0
}

func (x *MonthToDateCostForecast) GetCostForecast() float64 {
	if x != nil {
		return x.CostForecast
	}
	return 0
}

func (x *MonthToDateCostForecast) GetUpperBound() float64 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

func (x *MonthToDateCostForecast) GetLowerBound() float64 {
	if x != nil {
		return x.LowerBound
	}
	return 0
}

var File_api_forecast_proto protoreflect.FileDescriptor

var file_api_forecast_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x22, 0xba, 0x02, 0x0a, 0x0c, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63,
	0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e,
	0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x5f,
	0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x73, 0x0a, 0x14, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x5a, 0x0a, 0x0b, 0x4f, 0x72, 0x67, 0x46, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xfd, 0x01, 0x0a, 0x17, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65,
	0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63,
	0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74,
	0x43, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x66, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x22, 0x89, 0x01, 0x0a, 0x1f, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x70, 0x0a, 0x16,
	0x4f, 0x72, 0x67, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65,
	0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc5,
	0x01, 0x0a, 0x13, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x73, 0x74, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x63, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f,
	0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x70, 0x65,
	0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x18, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x4f, 0x6e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x43, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x65, 0x76, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x65, 0x76, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73,
	0x74, 0x22, 0xd7, 0x01, 0x0a, 0x17, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x65, 0x76, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x65, 0x76, 0x12, 0x28, 0x0a,
	0x0f, 0x63, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x63, 0x75,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x63,
	0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x55, 0x0a, 0x19, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x10, 0x41, 0x70, 0x69, 0x46, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_forecast_proto_rawDescOnce sync.Once
	file_api_forecast_proto_rawDescData = file_api_forecast_proto_rawDesc
)

func file_api_forecast_proto_rawDescGZIP() []byte {
	file_api_forecast_proto_rawDescOnce.Do(func() {
		file_api_forecast_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_forecast_proto_rawDescData)
	})
	return file_api_forecast_proto_rawDescData
}

var file_api_forecast_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_forecast_proto_goTypes = []interface{}{
	(*ForecastData)(nil),                    // 0: blueapi.api.ForecastData
	(*AccountGroupForecast)(nil),            // 1: blueapi.api.AccountGroupForecast
	(*BillingGroupForecast)(nil),            // 2: blueapi.api.BillingGroupForecast
	(*OrgForecast)(nil),                     // 3: blueapi.api.OrgForecast
	(*MonthToDateForecastData)(nil),         // 4: blueapi.api.MonthToDateForecastData
	(*BillingGroupMonthToDateForecast)(nil), // 5: blueapi.api.BillingGroupMonthToDateForecast
	(*OrgMonthToDateForecast)(nil),          // 6: blueapi.api.OrgMonthToDateForecast
	(*MonthlyCostForecast)(nil),             // 7: blueapi.api.MonthlyCostForecast
	(*MonthOnMonthCostForecast)(nil),        // 8: blueapi.api.MonthOnMonthCostForecast
	(*MonthToDateCostForecast)(nil),         // 9: blueapi.api.MonthToDateCostForecast
}
var file_api_forecast_proto_depIdxs = []int32{
	0, // 0: blueapi.api.AccountGroupForecast.data:type_name -> blueapi.api.ForecastData
	0, // 1: blueapi.api.BillingGroupForecast.data:type_name -> blueapi.api.ForecastData
	2, // 2: blueapi.api.OrgForecast.data:type_name -> blueapi.api.BillingGroupForecast
	4, // 3: blueapi.api.BillingGroupMonthToDateForecast.data:type_name -> blueapi.api.MonthToDateForecastData
	5, // 4: blueapi.api.OrgMonthToDateForecast.data:type_name -> blueapi.api.BillingGroupMonthToDateForecast
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_forecast_proto_init() }
func file_api_forecast_proto_init() {
	if File_api_forecast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_forecast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecastData); i {
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
		file_api_forecast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountGroupForecast); i {
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
		file_api_forecast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingGroupForecast); i {
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
		file_api_forecast_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgForecast); i {
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
		file_api_forecast_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonthToDateForecastData); i {
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
		file_api_forecast_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingGroupMonthToDateForecast); i {
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
		file_api_forecast_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgMonthToDateForecast); i {
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
		file_api_forecast_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonthlyCostForecast); i {
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
		file_api_forecast_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonthOnMonthCostForecast); i {
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
		file_api_forecast_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonthToDateCostForecast); i {
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
			RawDescriptor: file_api_forecast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_forecast_proto_goTypes,
		DependencyIndexes: file_api_forecast_proto_depIdxs,
		MessageInfos:      file_api_forecast_proto_msgTypes,
	}.Build()
	File_api_forecast_proto = out.File
	file_api_forecast_proto_rawDesc = nil
	file_api_forecast_proto_goTypes = nil
	file_api_forecast_proto_depIdxs = nil
}
