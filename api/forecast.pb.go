// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
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

type BillingGroupForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId string          `protobuf:"bytes,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Data      []*ForecastData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BillingGroupForecast) Reset() {
	*x = BillingGroupForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingGroupForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingGroupForecast) ProtoMessage() {}

func (x *BillingGroupForecast) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BillingGroupForecast.ProtoReflect.Descriptor instead.
func (*BillingGroupForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{1}
}

func (x *BillingGroupForecast) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
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
		mi := &file_api_forecast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgForecast) ProtoMessage() {}

func (x *OrgForecast) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OrgForecast.ProtoReflect.Descriptor instead.
func (*OrgForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{2}
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
		mi := &file_api_forecast_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonthToDateForecastData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthToDateForecastData) ProtoMessage() {}

func (x *MonthToDateForecastData) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MonthToDateForecastData.ProtoReflect.Descriptor instead.
func (*MonthToDateForecastData) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{3}
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

	CompanyId string                     `protobuf:"bytes,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Data      []*MonthToDateForecastData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BillingGroupMonthToDateForecast) Reset() {
	*x = BillingGroupMonthToDateForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_forecast_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingGroupMonthToDateForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingGroupMonthToDateForecast) ProtoMessage() {}

func (x *BillingGroupMonthToDateForecast) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BillingGroupMonthToDateForecast.ProtoReflect.Descriptor instead.
func (*BillingGroupMonthToDateForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{4}
}

func (x *BillingGroupMonthToDateForecast) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
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
		mi := &file_api_forecast_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgMonthToDateForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgMonthToDateForecast) ProtoMessage() {}

func (x *OrgMonthToDateForecast) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OrgMonthToDateForecast.ProtoReflect.Descriptor instead.
func (*OrgMonthToDateForecast) Descriptor() ([]byte, []int) {
	return file_api_forecast_proto_rawDescGZIP(), []int{5}
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

var File_api_forecast_proto protoreflect.FileDescriptor

var file_api_forecast_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x22, 0xfa, 0x01, 0x0a, 0x0c, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61,
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
	0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x22, 0x63,
	0x0a, 0x14, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x5a, 0x0a, 0x0b, 0x4f, 0x72, 0x67, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0xfd, 0x01, 0x0a, 0x17, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x43,
	0x6f, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x43,
	0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x66, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22,
	0x79, 0x0a, 0x1f, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x70, 0x0a, 0x16, 0x4f, 0x72,
	0x67, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x55, 0x0a, 0x19,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c,
	0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x10, 0x41, 0x70, 0x69, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_api_forecast_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_forecast_proto_goTypes = []interface{}{
	(*ForecastData)(nil),                    // 0: blueapi.api.ForecastData
	(*BillingGroupForecast)(nil),            // 1: blueapi.api.BillingGroupForecast
	(*OrgForecast)(nil),                     // 2: blueapi.api.OrgForecast
	(*MonthToDateForecastData)(nil),         // 3: blueapi.api.MonthToDateForecastData
	(*BillingGroupMonthToDateForecast)(nil), // 4: blueapi.api.BillingGroupMonthToDateForecast
	(*OrgMonthToDateForecast)(nil),          // 5: blueapi.api.OrgMonthToDateForecast
}
var file_api_forecast_proto_depIdxs = []int32{
	0, // 0: blueapi.api.BillingGroupForecast.data:type_name -> blueapi.api.ForecastData
	1, // 1: blueapi.api.OrgForecast.data:type_name -> blueapi.api.BillingGroupForecast
	3, // 2: blueapi.api.BillingGroupMonthToDateForecast.data:type_name -> blueapi.api.MonthToDateForecastData
	4, // 3: blueapi.api.OrgMonthToDateForecast.data:type_name -> blueapi.api.BillingGroupMonthToDateForecast
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
		file_api_forecast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_forecast_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_forecast_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_forecast_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_forecast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
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