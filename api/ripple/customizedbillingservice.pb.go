// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/ripple/customizedbillingservice.proto

package ripple

import (
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

// Describes the overall config of a [blueapi.api.ripple.MethodConfig].
type TargetServiceConfig int32

const (
	TargetServiceConfig_ALL TargetServiceConfig = 0
)

// Enum value maps for TargetServiceConfig.
var (
	TargetServiceConfig_name = map[int32]string{
		0: "ALL",
	}
	TargetServiceConfig_value = map[string]int32{
		"ALL": 0,
	}
)

func (x TargetServiceConfig) Enum() *TargetServiceConfig {
	p := new(TargetServiceConfig)
	*p = x
	return p
}

func (x TargetServiceConfig) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetServiceConfig) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ripple_customizedbillingservice_proto_enumTypes[0].Descriptor()
}

func (TargetServiceConfig) Type() protoreflect.EnumType {
	return &file_api_ripple_customizedbillingservice_proto_enumTypes[0]
}

func (x TargetServiceConfig) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetServiceConfig.Descriptor instead.
func (TargetServiceConfig) EnumDescriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{0}
}

// Describes the overall config of a [blueapi.api.ripple.MethodConfig].
type TargetUsageConfig int32

const (
	TargetUsageConfig_USAGE TargetUsageConfig = 0
)

// Enum value maps for TargetUsageConfig.
var (
	TargetUsageConfig_name = map[int32]string{
		0: "USAGE",
	}
	TargetUsageConfig_value = map[string]int32{
		"USAGE": 0,
	}
)

func (x TargetUsageConfig) Enum() *TargetUsageConfig {
	p := new(TargetUsageConfig)
	*p = x
	return p
}

func (x TargetUsageConfig) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetUsageConfig) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ripple_customizedbillingservice_proto_enumTypes[1].Descriptor()
}

func (TargetUsageConfig) Type() protoreflect.EnumType {
	return &file_api_ripple_customizedbillingservice_proto_enumTypes[1]
}

func (x TargetUsageConfig) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetUsageConfig.Descriptor instead.
func (TargetUsageConfig) EnumDescriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{1}
}

// Charging target
// Indicates that `ChargingTarget` in the setting applies to either BillingGroup or Account.
type ChargingTarget int32

const (
	// billing group
	ChargingTarget_BILLINGGROUP ChargingTarget = 0
	// account
	ChargingTarget_ACCOUNT ChargingTarget = 1
)

// Enum value maps for ChargingTarget.
var (
	ChargingTarget_name = map[int32]string{
		0: "BILLINGGROUP",
		1: "ACCOUNT",
	}
	ChargingTarget_value = map[string]int32{
		"BILLINGGROUP": 0,
		"ACCOUNT":      1,
	}
)

func (x ChargingTarget) Enum() *ChargingTarget {
	p := new(ChargingTarget)
	*p = x
	return p
}

func (x ChargingTarget) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChargingTarget) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ripple_customizedbillingservice_proto_enumTypes[2].Descriptor()
}

func (ChargingTarget) Type() protoreflect.EnumType {
	return &file_api_ripple_customizedbillingservice_proto_enumTypes[2]
}

func (x ChargingTarget) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChargingTarget.Descriptor instead.
func (ChargingTarget) EnumDescriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{2}
}

// supported currency.
type MethodConfig_Currency int32

const (
	MethodConfig_JPY MethodConfig_Currency = 0
	MethodConfig_USD MethodConfig_Currency = 1
	MethodConfig_IDR MethodConfig_Currency = 2
	MethodConfig_MYR MethodConfig_Currency = 3
	MethodConfig_SGD MethodConfig_Currency = 4
	MethodConfig_INR MethodConfig_Currency = 5
)

// Enum value maps for MethodConfig_Currency.
var (
	MethodConfig_Currency_name = map[int32]string{
		0: "JPY",
		1: "USD",
		2: "IDR",
		3: "MYR",
		4: "SGD",
		5: "INR",
	}
	MethodConfig_Currency_value = map[string]int32{
		"JPY": 0,
		"USD": 1,
		"IDR": 2,
		"MYR": 3,
		"SGD": 4,
		"INR": 5,
	}
)

func (x MethodConfig_Currency) Enum() *MethodConfig_Currency {
	p := new(MethodConfig_Currency)
	*p = x
	return p
}

func (x MethodConfig_Currency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MethodConfig_Currency) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ripple_customizedbillingservice_proto_enumTypes[3].Descriptor()
}

func (MethodConfig_Currency) Type() protoreflect.EnumType {
	return &file_api_ripple_customizedbillingservice_proto_enumTypes[3]
}

func (x MethodConfig_Currency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MethodConfig_Currency.Descriptor instead.
func (MethodConfig_Currency) EnumDescriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{1, 0}
}

type MethodConfig_ChargingMethod int32

const (
	MethodConfig_FIXED_FEE               MethodConfig_ChargingMethod = 0
	MethodConfig_PERCENTAGE              MethodConfig_ChargingMethod = 1
	MethodConfig_FIXED_FEE_OR_PERCENTAGE MethodConfig_ChargingMethod = 2
	MethodConfig_TIERED_PRICE            MethodConfig_ChargingMethod = 3
	MethodConfig_TIERED_PERCENTAGE       MethodConfig_ChargingMethod = 4
)

// Enum value maps for MethodConfig_ChargingMethod.
var (
	MethodConfig_ChargingMethod_name = map[int32]string{
		0: "FIXED_FEE",
		1: "PERCENTAGE",
		2: "FIXED_FEE_OR_PERCENTAGE",
		3: "TIERED_PRICE",
		4: "TIERED_PERCENTAGE",
	}
	MethodConfig_ChargingMethod_value = map[string]int32{
		"FIXED_FEE":               0,
		"PERCENTAGE":              1,
		"FIXED_FEE_OR_PERCENTAGE": 2,
		"TIERED_PRICE":            3,
		"TIERED_PERCENTAGE":       4,
	}
)

func (x MethodConfig_ChargingMethod) Enum() *MethodConfig_ChargingMethod {
	p := new(MethodConfig_ChargingMethod)
	*p = x
	return p
}

func (x MethodConfig_ChargingMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MethodConfig_ChargingMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ripple_customizedbillingservice_proto_enumTypes[4].Descriptor()
}

func (MethodConfig_ChargingMethod) Type() protoreflect.EnumType {
	return &file_api_ripple_customizedbillingservice_proto_enumTypes[4]
}

func (x MethodConfig_ChargingMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MethodConfig_ChargingMethod.Descriptor instead.
func (MethodConfig_ChargingMethod) EnumDescriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{1, 1}
}

// CustomizedBillingService resource definition.
type CustomizedBillingService struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	MethodConfig  *MethodConfig          `protobuf:"bytes,4,opt,name=methodConfig,proto3" json:"methodConfig,omitempty"`
	Created       string                 `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated       string                 `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomizedBillingService) Reset() {
	*x = CustomizedBillingService{}
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomizedBillingService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomizedBillingService) ProtoMessage() {}

func (x *CustomizedBillingService) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomizedBillingService.ProtoReflect.Descriptor instead.
func (*CustomizedBillingService) Descriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{0}
}

func (x *CustomizedBillingService) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CustomizedBillingService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomizedBillingService) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CustomizedBillingService) GetMethodConfig() *MethodConfig {
	if x != nil {
		return x.MethodConfig
	}
	return nil
}

func (x *CustomizedBillingService) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *CustomizedBillingService) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

// MethodConfig resource definition.
type MethodConfig struct {
	state          protoimpl.MessageState      `protogen:"open.v1"`
	Currency       MethodConfig_Currency       `protobuf:"varint,1,opt,name=currency,proto3,enum=blueapi.api.ripple.MethodConfig_Currency" json:"currency,omitempty"`
	ChargingMethod MethodConfig_ChargingMethod `protobuf:"varint,2,opt,name=chargingMethod,proto3,enum=blueapi.api.ripple.MethodConfig_ChargingMethod" json:"chargingMethod,omitempty"`
	// Types that are valid to be assigned to Config:
	//
	//	*MethodConfig_FixedFee
	//	*MethodConfig_Percentage
	//	*MethodConfig_FixedFeeOrPercentage
	//	*MethodConfig_TieredPrice
	//	*MethodConfig_TieredPercentage
	Config        isMethodConfig_Config `protobuf_oneof:"Config"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MethodConfig) Reset() {
	*x = MethodConfig{}
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MethodConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodConfig) ProtoMessage() {}

func (x *MethodConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodConfig.ProtoReflect.Descriptor instead.
func (*MethodConfig) Descriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{1}
}

func (x *MethodConfig) GetCurrency() MethodConfig_Currency {
	if x != nil {
		return x.Currency
	}
	return MethodConfig_JPY
}

func (x *MethodConfig) GetChargingMethod() MethodConfig_ChargingMethod {
	if x != nil {
		return x.ChargingMethod
	}
	return MethodConfig_FIXED_FEE
}

func (x *MethodConfig) GetConfig() isMethodConfig_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *MethodConfig) GetFixedFee() *FixedFee {
	if x != nil {
		if x, ok := x.Config.(*MethodConfig_FixedFee); ok {
			return x.FixedFee
		}
	}
	return nil
}

func (x *MethodConfig) GetPercentage() *Percentage {
	if x != nil {
		if x, ok := x.Config.(*MethodConfig_Percentage); ok {
			return x.Percentage
		}
	}
	return nil
}

func (x *MethodConfig) GetFixedFeeOrPercentage() *FixedFeeOrPercentage {
	if x != nil {
		if x, ok := x.Config.(*MethodConfig_FixedFeeOrPercentage); ok {
			return x.FixedFeeOrPercentage
		}
	}
	return nil
}

func (x *MethodConfig) GetTieredPrice() *TieredPrice {
	if x != nil {
		if x, ok := x.Config.(*MethodConfig_TieredPrice); ok {
			return x.TieredPrice
		}
	}
	return nil
}

func (x *MethodConfig) GetTieredPercentage() *TieredPercentage {
	if x != nil {
		if x, ok := x.Config.(*MethodConfig_TieredPercentage); ok {
			return x.TieredPercentage
		}
	}
	return nil
}

type isMethodConfig_Config interface {
	isMethodConfig_Config()
}

type MethodConfig_FixedFee struct {
	FixedFee *FixedFee `protobuf:"bytes,3,opt,name=fixedFee,proto3,oneof"`
}

type MethodConfig_Percentage struct {
	Percentage *Percentage `protobuf:"bytes,4,opt,name=percentage,proto3,oneof"`
}

type MethodConfig_FixedFeeOrPercentage struct {
	FixedFeeOrPercentage *FixedFeeOrPercentage `protobuf:"bytes,5,opt,name=fixedFeeOrPercentage,proto3,oneof"`
}

type MethodConfig_TieredPrice struct {
	TieredPrice *TieredPrice `protobuf:"bytes,6,opt,name=TieredPrice,proto3,oneof"`
}

type MethodConfig_TieredPercentage struct {
	TieredPercentage *TieredPercentage `protobuf:"bytes,7,opt,name=TieredPercentage,proto3,oneof"`
}

func (*MethodConfig_FixedFee) isMethodConfig_Config() {}

func (*MethodConfig_Percentage) isMethodConfig_Config() {}

func (*MethodConfig_FixedFeeOrPercentage) isMethodConfig_Config() {}

func (*MethodConfig_TieredPrice) isMethodConfig_Config() {}

func (*MethodConfig_TieredPercentage) isMethodConfig_Config() {}

// fixed fee
type FixedFee struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         float64                `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FixedFee) Reset() {
	*x = FixedFee{}
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FixedFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixedFee) ProtoMessage() {}

func (x *FixedFee) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixedFee.ProtoReflect.Descriptor instead.
func (*FixedFee) Descriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{2}
}

func (x *FixedFee) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// percentage
type Percentage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         float64                `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	Service       TargetServiceConfig    `protobuf:"varint,2,opt,name=service,proto3,enum=blueapi.api.ripple.TargetServiceConfig" json:"service,omitempty"`
	Usage         TargetUsageConfig      `protobuf:"varint,3,opt,name=usage,proto3,enum=blueapi.api.ripple.TargetUsageConfig" json:"usage,omitempty"`
	Discounted    bool                   `protobuf:"varint,4,opt,name=discounted,proto3" json:"discounted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Percentage) Reset() {
	*x = Percentage{}
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Percentage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Percentage) ProtoMessage() {}

func (x *Percentage) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Percentage.ProtoReflect.Descriptor instead.
func (*Percentage) Descriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{3}
}

func (x *Percentage) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Percentage) GetService() TargetServiceConfig {
	if x != nil {
		return x.Service
	}
	return TargetServiceConfig_ALL
}

func (x *Percentage) GetUsage() TargetUsageConfig {
	if x != nil {
		return x.Usage
	}
	return TargetUsageConfig_USAGE
}

func (x *Percentage) GetDiscounted() bool {
	if x != nil {
		return x.Discounted
	}
	return false
}

// fixed fee or percentage
type FixedFeeOrPercentage struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	FixedFeeValue   float64                `protobuf:"fixed64,1,opt,name=fixedFeeValue,proto3" json:"fixedFeeValue,omitempty"`
	PercentageValue float64                `protobuf:"fixed64,2,opt,name=percentageValue,proto3" json:"percentageValue,omitempty"`
	Service         TargetServiceConfig    `protobuf:"varint,3,opt,name=service,proto3,enum=blueapi.api.ripple.TargetServiceConfig" json:"service,omitempty"`
	Usage           TargetUsageConfig      `protobuf:"varint,4,opt,name=usage,proto3,enum=blueapi.api.ripple.TargetUsageConfig" json:"usage,omitempty"`
	Discounted      bool                   `protobuf:"varint,5,opt,name=discounted,proto3" json:"discounted,omitempty"`
	UpperLimit      float64                `protobuf:"fixed64,6,opt,name=upperLimit,proto3" json:"upperLimit,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *FixedFeeOrPercentage) Reset() {
	*x = FixedFeeOrPercentage{}
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FixedFeeOrPercentage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixedFeeOrPercentage) ProtoMessage() {}

func (x *FixedFeeOrPercentage) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixedFeeOrPercentage.ProtoReflect.Descriptor instead.
func (*FixedFeeOrPercentage) Descriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{4}
}

func (x *FixedFeeOrPercentage) GetFixedFeeValue() float64 {
	if x != nil {
		return x.FixedFeeValue
	}
	return 0
}

func (x *FixedFeeOrPercentage) GetPercentageValue() float64 {
	if x != nil {
		return x.PercentageValue
	}
	return 0
}

func (x *FixedFeeOrPercentage) GetService() TargetServiceConfig {
	if x != nil {
		return x.Service
	}
	return TargetServiceConfig_ALL
}

func (x *FixedFeeOrPercentage) GetUsage() TargetUsageConfig {
	if x != nil {
		return x.Usage
	}
	return TargetUsageConfig_USAGE
}

func (x *FixedFeeOrPercentage) GetDiscounted() bool {
	if x != nil {
		return x.Discounted
	}
	return false
}

func (x *FixedFeeOrPercentage) GetUpperLimit() float64 {
	if x != nil {
		return x.UpperLimit
	}
	return 0
}

// tired price
type TieredPrice struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TiredConfig   []*TierdConfig         `protobuf:"bytes,1,rep,name=tiredConfig,proto3" json:"tiredConfig,omitempty"`
	Service       TargetServiceConfig    `protobuf:"varint,2,opt,name=service,proto3,enum=blueapi.api.ripple.TargetServiceConfig" json:"service,omitempty"`
	Usage         TargetUsageConfig      `protobuf:"varint,3,opt,name=usage,proto3,enum=blueapi.api.ripple.TargetUsageConfig" json:"usage,omitempty"`
	Discounted    bool                   `protobuf:"varint,4,opt,name=discounted,proto3" json:"discounted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TieredPrice) Reset() {
	*x = TieredPrice{}
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TieredPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TieredPrice) ProtoMessage() {}

func (x *TieredPrice) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TieredPrice.ProtoReflect.Descriptor instead.
func (*TieredPrice) Descriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{5}
}

func (x *TieredPrice) GetTiredConfig() []*TierdConfig {
	if x != nil {
		return x.TiredConfig
	}
	return nil
}

func (x *TieredPrice) GetService() TargetServiceConfig {
	if x != nil {
		return x.Service
	}
	return TargetServiceConfig_ALL
}

func (x *TieredPrice) GetUsage() TargetUsageConfig {
	if x != nil {
		return x.Usage
	}
	return TargetUsageConfig_USAGE
}

func (x *TieredPrice) GetDiscounted() bool {
	if x != nil {
		return x.Discounted
	}
	return false
}

// tired percentage
type TieredPercentage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TiredConfig   []*TierdConfig         `protobuf:"bytes,1,rep,name=tiredConfig,proto3" json:"tiredConfig,omitempty"`
	Service       TargetServiceConfig    `protobuf:"varint,2,opt,name=service,proto3,enum=blueapi.api.ripple.TargetServiceConfig" json:"service,omitempty"`
	Usage         TargetUsageConfig      `protobuf:"varint,3,opt,name=usage,proto3,enum=blueapi.api.ripple.TargetUsageConfig" json:"usage,omitempty"`
	Discounted    bool                   `protobuf:"varint,4,opt,name=discounted,proto3" json:"discounted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TieredPercentage) Reset() {
	*x = TieredPercentage{}
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TieredPercentage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TieredPercentage) ProtoMessage() {}

func (x *TieredPercentage) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TieredPercentage.ProtoReflect.Descriptor instead.
func (*TieredPercentage) Descriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{6}
}

func (x *TieredPercentage) GetTiredConfig() []*TierdConfig {
	if x != nil {
		return x.TiredConfig
	}
	return nil
}

func (x *TieredPercentage) GetService() TargetServiceConfig {
	if x != nil {
		return x.Service
	}
	return TargetServiceConfig_ALL
}

func (x *TieredPercentage) GetUsage() TargetUsageConfig {
	if x != nil {
		return x.Usage
	}
	return TargetUsageConfig_USAGE
}

func (x *TieredPercentage) GetDiscounted() bool {
	if x != nil {
		return x.Discounted
	}
	return false
}

// tired config
type TierdConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Min           float64                `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	Max           float64                `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	Value         float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TierdConfig) Reset() {
	*x = TierdConfig{}
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TierdConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TierdConfig) ProtoMessage() {}

func (x *TierdConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_customizedbillingservice_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TierdConfig.ProtoReflect.Descriptor instead.
func (*TierdConfig) Descriptor() ([]byte, []int) {
	return file_api_ripple_customizedbillingservice_proto_rawDescGZIP(), []int{7}
}

func (x *TierdConfig) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *TierdConfig) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *TierdConfig) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_api_ripple_customizedbillingservice_proto protoreflect.FileDescriptor

const file_api_ripple_customizedbillingservice_proto_rawDesc = "" +
	"\n" +
	")api/ripple/customizedbillingservice.proto\x12\x12blueapi.api.ripple\x1a\x1fgoogle/api/field_behavior.proto\"\xec\x01\n" +
	"\x18CustomizedBillingService\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12&\n" +
	"\vdescription\x18\x03 \x01(\tB\x04\xe2A\x01\x01R\vdescription\x12D\n" +
	"\fmethodConfig\x18\x04 \x01(\v2 .blueapi.api.ripple.MethodConfigR\fmethodConfig\x12\x1e\n" +
	"\acreated\x18\x05 \x01(\tB\x04\xe2A\x01\x03R\acreated\x12\x1e\n" +
	"\aupdated\x18\x06 \x01(\tB\x04\xe2A\x01\x03R\aupdated\"\xe8\x05\n" +
	"\fMethodConfig\x12E\n" +
	"\bcurrency\x18\x01 \x01(\x0e2).blueapi.api.ripple.MethodConfig.CurrencyR\bcurrency\x12W\n" +
	"\x0echargingMethod\x18\x02 \x01(\x0e2/.blueapi.api.ripple.MethodConfig.ChargingMethodR\x0echargingMethod\x12:\n" +
	"\bfixedFee\x18\x03 \x01(\v2\x1c.blueapi.api.ripple.FixedFeeH\x00R\bfixedFee\x12@\n" +
	"\n" +
	"percentage\x18\x04 \x01(\v2\x1e.blueapi.api.ripple.PercentageH\x00R\n" +
	"percentage\x12^\n" +
	"\x14fixedFeeOrPercentage\x18\x05 \x01(\v2(.blueapi.api.ripple.FixedFeeOrPercentageH\x00R\x14fixedFeeOrPercentage\x12C\n" +
	"\vTieredPrice\x18\x06 \x01(\v2\x1f.blueapi.api.ripple.TieredPriceH\x00R\vTieredPrice\x12R\n" +
	"\x10TieredPercentage\x18\a \x01(\v2$.blueapi.api.ripple.TieredPercentageH\x00R\x10TieredPercentage\"@\n" +
	"\bCurrency\x12\a\n" +
	"\x03JPY\x10\x00\x12\a\n" +
	"\x03USD\x10\x01\x12\a\n" +
	"\x03IDR\x10\x02\x12\a\n" +
	"\x03MYR\x10\x03\x12\a\n" +
	"\x03SGD\x10\x04\x12\a\n" +
	"\x03INR\x10\x05\"u\n" +
	"\x0eChargingMethod\x12\r\n" +
	"\tFIXED_FEE\x10\x00\x12\x0e\n" +
	"\n" +
	"PERCENTAGE\x10\x01\x12\x1b\n" +
	"\x17FIXED_FEE_OR_PERCENTAGE\x10\x02\x12\x10\n" +
	"\fTIERED_PRICE\x10\x03\x12\x15\n" +
	"\x11TIERED_PERCENTAGE\x10\x04B\b\n" +
	"\x06Config\" \n" +
	"\bFixedFee\x12\x14\n" +
	"\x05value\x18\x01 \x01(\x01R\x05value\"\xc2\x01\n" +
	"\n" +
	"Percentage\x12\x14\n" +
	"\x05value\x18\x01 \x01(\x01R\x05value\x12A\n" +
	"\aservice\x18\x02 \x01(\x0e2'.blueapi.api.ripple.TargetServiceConfigR\aservice\x12;\n" +
	"\x05usage\x18\x03 \x01(\x0e2%.blueapi.api.ripple.TargetUsageConfigR\x05usage\x12\x1e\n" +
	"\n" +
	"discounted\x18\x04 \x01(\bR\n" +
	"discounted\"\xa6\x02\n" +
	"\x14FixedFeeOrPercentage\x12$\n" +
	"\rfixedFeeValue\x18\x01 \x01(\x01R\rfixedFeeValue\x12(\n" +
	"\x0fpercentageValue\x18\x02 \x01(\x01R\x0fpercentageValue\x12A\n" +
	"\aservice\x18\x03 \x01(\x0e2'.blueapi.api.ripple.TargetServiceConfigR\aservice\x12;\n" +
	"\x05usage\x18\x04 \x01(\x0e2%.blueapi.api.ripple.TargetUsageConfigR\x05usage\x12\x1e\n" +
	"\n" +
	"discounted\x18\x05 \x01(\bR\n" +
	"discounted\x12\x1e\n" +
	"\n" +
	"upperLimit\x18\x06 \x01(\x01R\n" +
	"upperLimit\"\xf0\x01\n" +
	"\vTieredPrice\x12A\n" +
	"\vtiredConfig\x18\x01 \x03(\v2\x1f.blueapi.api.ripple.TierdConfigR\vtiredConfig\x12A\n" +
	"\aservice\x18\x02 \x01(\x0e2'.blueapi.api.ripple.TargetServiceConfigR\aservice\x12;\n" +
	"\x05usage\x18\x03 \x01(\x0e2%.blueapi.api.ripple.TargetUsageConfigR\x05usage\x12\x1e\n" +
	"\n" +
	"discounted\x18\x04 \x01(\bR\n" +
	"discounted\"\xf5\x01\n" +
	"\x10TieredPercentage\x12A\n" +
	"\vtiredConfig\x18\x01 \x03(\v2\x1f.blueapi.api.ripple.TierdConfigR\vtiredConfig\x12A\n" +
	"\aservice\x18\x02 \x01(\x0e2'.blueapi.api.ripple.TargetServiceConfigR\aservice\x12;\n" +
	"\x05usage\x18\x03 \x01(\x0e2%.blueapi.api.ripple.TargetUsageConfigR\x05usage\x12\x1e\n" +
	"\n" +
	"discounted\x18\x04 \x01(\bR\n" +
	"discounted\"G\n" +
	"\vTierdConfig\x12\x10\n" +
	"\x03min\x18\x01 \x01(\x01R\x03min\x12\x10\n" +
	"\x03max\x18\x02 \x01(\x01R\x03max\x12\x14\n" +
	"\x05value\x18\x03 \x01(\x01R\x05value*\x1e\n" +
	"\x13TargetServiceConfig\x12\a\n" +
	"\x03ALL\x10\x00*\x1e\n" +
	"\x11TargetUsageConfig\x12\t\n" +
	"\x05USAGE\x10\x00*/\n" +
	"\x0eChargingTarget\x12\x10\n" +
	"\fBILLINGGROUP\x10\x00\x12\v\n" +
	"\aACCOUNT\x10\x01By\n" +
	" cloud.alphaus.blueapi.api.rippleB&ApiRippleCustomizedBillingServiceProtoZ-github.com/alphauslabs/blue-sdk-go/api/rippleb\x06proto3"

var (
	file_api_ripple_customizedbillingservice_proto_rawDescOnce sync.Once
	file_api_ripple_customizedbillingservice_proto_rawDescData []byte
)

func file_api_ripple_customizedbillingservice_proto_rawDescGZIP() []byte {
	file_api_ripple_customizedbillingservice_proto_rawDescOnce.Do(func() {
		file_api_ripple_customizedbillingservice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_ripple_customizedbillingservice_proto_rawDesc), len(file_api_ripple_customizedbillingservice_proto_rawDesc)))
	})
	return file_api_ripple_customizedbillingservice_proto_rawDescData
}

var file_api_ripple_customizedbillingservice_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_api_ripple_customizedbillingservice_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_ripple_customizedbillingservice_proto_goTypes = []any{
	(TargetServiceConfig)(0),         // 0: blueapi.api.ripple.TargetServiceConfig
	(TargetUsageConfig)(0),           // 1: blueapi.api.ripple.TargetUsageConfig
	(ChargingTarget)(0),              // 2: blueapi.api.ripple.ChargingTarget
	(MethodConfig_Currency)(0),       // 3: blueapi.api.ripple.MethodConfig.Currency
	(MethodConfig_ChargingMethod)(0), // 4: blueapi.api.ripple.MethodConfig.ChargingMethod
	(*CustomizedBillingService)(nil), // 5: blueapi.api.ripple.CustomizedBillingService
	(*MethodConfig)(nil),             // 6: blueapi.api.ripple.MethodConfig
	(*FixedFee)(nil),                 // 7: blueapi.api.ripple.FixedFee
	(*Percentage)(nil),               // 8: blueapi.api.ripple.Percentage
	(*FixedFeeOrPercentage)(nil),     // 9: blueapi.api.ripple.FixedFeeOrPercentage
	(*TieredPrice)(nil),              // 10: blueapi.api.ripple.TieredPrice
	(*TieredPercentage)(nil),         // 11: blueapi.api.ripple.TieredPercentage
	(*TierdConfig)(nil),              // 12: blueapi.api.ripple.TierdConfig
}
var file_api_ripple_customizedbillingservice_proto_depIdxs = []int32{
	6,  // 0: blueapi.api.ripple.CustomizedBillingService.methodConfig:type_name -> blueapi.api.ripple.MethodConfig
	3,  // 1: blueapi.api.ripple.MethodConfig.currency:type_name -> blueapi.api.ripple.MethodConfig.Currency
	4,  // 2: blueapi.api.ripple.MethodConfig.chargingMethod:type_name -> blueapi.api.ripple.MethodConfig.ChargingMethod
	7,  // 3: blueapi.api.ripple.MethodConfig.fixedFee:type_name -> blueapi.api.ripple.FixedFee
	8,  // 4: blueapi.api.ripple.MethodConfig.percentage:type_name -> blueapi.api.ripple.Percentage
	9,  // 5: blueapi.api.ripple.MethodConfig.fixedFeeOrPercentage:type_name -> blueapi.api.ripple.FixedFeeOrPercentage
	10, // 6: blueapi.api.ripple.MethodConfig.TieredPrice:type_name -> blueapi.api.ripple.TieredPrice
	11, // 7: blueapi.api.ripple.MethodConfig.TieredPercentage:type_name -> blueapi.api.ripple.TieredPercentage
	0,  // 8: blueapi.api.ripple.Percentage.service:type_name -> blueapi.api.ripple.TargetServiceConfig
	1,  // 9: blueapi.api.ripple.Percentage.usage:type_name -> blueapi.api.ripple.TargetUsageConfig
	0,  // 10: blueapi.api.ripple.FixedFeeOrPercentage.service:type_name -> blueapi.api.ripple.TargetServiceConfig
	1,  // 11: blueapi.api.ripple.FixedFeeOrPercentage.usage:type_name -> blueapi.api.ripple.TargetUsageConfig
	12, // 12: blueapi.api.ripple.TieredPrice.tiredConfig:type_name -> blueapi.api.ripple.TierdConfig
	0,  // 13: blueapi.api.ripple.TieredPrice.service:type_name -> blueapi.api.ripple.TargetServiceConfig
	1,  // 14: blueapi.api.ripple.TieredPrice.usage:type_name -> blueapi.api.ripple.TargetUsageConfig
	12, // 15: blueapi.api.ripple.TieredPercentage.tiredConfig:type_name -> blueapi.api.ripple.TierdConfig
	0,  // 16: blueapi.api.ripple.TieredPercentage.service:type_name -> blueapi.api.ripple.TargetServiceConfig
	1,  // 17: blueapi.api.ripple.TieredPercentage.usage:type_name -> blueapi.api.ripple.TargetUsageConfig
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_api_ripple_customizedbillingservice_proto_init() }
func file_api_ripple_customizedbillingservice_proto_init() {
	if File_api_ripple_customizedbillingservice_proto != nil {
		return
	}
	file_api_ripple_customizedbillingservice_proto_msgTypes[1].OneofWrappers = []any{
		(*MethodConfig_FixedFee)(nil),
		(*MethodConfig_Percentage)(nil),
		(*MethodConfig_FixedFeeOrPercentage)(nil),
		(*MethodConfig_TieredPrice)(nil),
		(*MethodConfig_TieredPercentage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_ripple_customizedbillingservice_proto_rawDesc), len(file_api_ripple_customizedbillingservice_proto_rawDesc)),
			NumEnums:      5,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_ripple_customizedbillingservice_proto_goTypes,
		DependencyIndexes: file_api_ripple_customizedbillingservice_proto_depIdxs,
		EnumInfos:         file_api_ripple_customizedbillingservice_proto_enumTypes,
		MessageInfos:      file_api_ripple_customizedbillingservice_proto_msgTypes,
	}.Build()
	File_api_ripple_customizedbillingservice_proto = out.File
	file_api_ripple_customizedbillingservice_proto_goTypes = nil
	file_api_ripple_customizedbillingservice_proto_depIdxs = nil
}
