// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/budget.proto

package api

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

type Budget struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Format: yyyy
	FiscalYear    string           `protobuf:"bytes,2,opt,name=fiscalYear,proto3" json:"fiscalYear,omitempty"`
	MonthlyBudget []*MonthlyBudget `protobuf:"bytes,3,rep,name=monthlyBudget,proto3" json:"monthlyBudget,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Budget) Reset() {
	*x = Budget{}
	mi := &file_api_budget_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Budget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budget) ProtoMessage() {}

func (x *Budget) ProtoReflect() protoreflect.Message {
	mi := &file_api_budget_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Budget.ProtoReflect.Descriptor instead.
func (*Budget) Descriptor() ([]byte, []int) {
	return file_api_budget_proto_rawDescGZIP(), []int{0}
}

func (x *Budget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Budget) GetFiscalYear() string {
	if x != nil {
		return x.FiscalYear
	}
	return ""
}

func (x *Budget) GetMonthlyBudget() []*MonthlyBudget {
	if x != nil {
		return x.MonthlyBudget
	}
	return nil
}

type MonthlyBudget struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Format: yyyymm
	YearMonth     string  `protobuf:"bytes,1,opt,name=yearMonth,proto3" json:"yearMonth,omitempty"`
	Amount        float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MonthlyBudget) Reset() {
	*x = MonthlyBudget{}
	mi := &file_api_budget_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MonthlyBudget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthlyBudget) ProtoMessage() {}

func (x *MonthlyBudget) ProtoReflect() protoreflect.Message {
	mi := &file_api_budget_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthlyBudget.ProtoReflect.Descriptor instead.
func (*MonthlyBudget) Descriptor() ([]byte, []int) {
	return file_api_budget_proto_rawDescGZIP(), []int{1}
}

func (x *MonthlyBudget) GetYearMonth() string {
	if x != nil {
		return x.YearMonth
	}
	return ""
}

func (x *MonthlyBudget) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// DailyBudgetAlert resource definition.
type DailyBudgetAlert struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. threshold in budget alerts
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// Required. notification enable/disable
	// If disabled, no alert is sent.
	Enabled       bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DailyBudgetAlert) Reset() {
	*x = DailyBudgetAlert{}
	mi := &file_api_budget_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DailyBudgetAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyBudgetAlert) ProtoMessage() {}

func (x *DailyBudgetAlert) ProtoReflect() protoreflect.Message {
	mi := &file_api_budget_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyBudgetAlert.ProtoReflect.Descriptor instead.
func (*DailyBudgetAlert) Descriptor() ([]byte, []int) {
	return file_api_budget_proto_rawDescGZIP(), []int{2}
}

func (x *DailyBudgetAlert) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *DailyBudgetAlert) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// DailyRateIncreaseBudgetAlert resource definition.
type DailyRateIncreaseBudgetAlert struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. threshold in budget alerts
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// Required. notification enable/disable
	// If disabled, no alert is sent.
	Enabled       bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DailyRateIncreaseBudgetAlert) Reset() {
	*x = DailyRateIncreaseBudgetAlert{}
	mi := &file_api_budget_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DailyRateIncreaseBudgetAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyRateIncreaseBudgetAlert) ProtoMessage() {}

func (x *DailyRateIncreaseBudgetAlert) ProtoReflect() protoreflect.Message {
	mi := &file_api_budget_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyRateIncreaseBudgetAlert.ProtoReflect.Descriptor instead.
func (*DailyRateIncreaseBudgetAlert) Descriptor() ([]byte, []int) {
	return file_api_budget_proto_rawDescGZIP(), []int{3}
}

func (x *DailyRateIncreaseBudgetAlert) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *DailyRateIncreaseBudgetAlert) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// MonthlyBudgetAlert resource definition.
type MonthlyBudgetAlert struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. threshold in budget alerts
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// Required. notification enable/disable
	// If disabled, no alert is sent.
	Enabled       bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MonthlyBudgetAlert) Reset() {
	*x = MonthlyBudgetAlert{}
	mi := &file_api_budget_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MonthlyBudgetAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthlyBudgetAlert) ProtoMessage() {}

func (x *MonthlyBudgetAlert) ProtoReflect() protoreflect.Message {
	mi := &file_api_budget_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthlyBudgetAlert.ProtoReflect.Descriptor instead.
func (*MonthlyBudgetAlert) Descriptor() ([]byte, []int) {
	return file_api_budget_proto_rawDescGZIP(), []int{4}
}

func (x *MonthlyBudgetAlert) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *MonthlyBudgetAlert) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// BudgetAlertNotification resource definition.
type BudgetAlertNotification struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. List of channelId. For example, you set to ["channelId1","channelId2","channelId3"].
	Channels []string `protobuf:"bytes,2,rep,name=channels,proto3" json:"channels,omitempty"`
	// Required. notification enable/disable
	// If disabled, no alert is sent.
	Enabled       bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BudgetAlertNotification) Reset() {
	*x = BudgetAlertNotification{}
	mi := &file_api_budget_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BudgetAlertNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetAlertNotification) ProtoMessage() {}

func (x *BudgetAlertNotification) ProtoReflect() protoreflect.Message {
	mi := &file_api_budget_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetAlertNotification.ProtoReflect.Descriptor instead.
func (*BudgetAlertNotification) Descriptor() ([]byte, []int) {
	return file_api_budget_proto_rawDescGZIP(), []int{5}
}

func (x *BudgetAlertNotification) GetChannels() []string {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *BudgetAlertNotification) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// BudgetAlertNotificationDetail resource definition.
type BudgetAlertNotificationDetail struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of channel info.
	Channels []*NotificationChannel `protobuf:"bytes,2,rep,name=channels,proto3" json:"channels,omitempty"`
	// Notification enable/disable
	// If disabled, no alert is sent.
	Enabled       bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BudgetAlertNotificationDetail) Reset() {
	*x = BudgetAlertNotificationDetail{}
	mi := &file_api_budget_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BudgetAlertNotificationDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetAlertNotificationDetail) ProtoMessage() {}

func (x *BudgetAlertNotificationDetail) ProtoReflect() protoreflect.Message {
	mi := &file_api_budget_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetAlertNotificationDetail.ProtoReflect.Descriptor instead.
func (*BudgetAlertNotificationDetail) Descriptor() ([]byte, []int) {
	return file_api_budget_proto_rawDescGZIP(), []int{6}
}

func (x *BudgetAlertNotificationDetail) GetChannels() []*NotificationChannel {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *BudgetAlertNotificationDetail) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_api_budget_proto protoreflect.FileDescriptor

const file_api_budget_proto_rawDesc = "" +
	"\n" +
	"\x10api/budget.proto\x12\vblueapi.api\x1a\x16api/notification.proto\"z\n" +
	"\x06Budget\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1e\n" +
	"\n" +
	"fiscalYear\x18\x02 \x01(\tR\n" +
	"fiscalYear\x12@\n" +
	"\rmonthlyBudget\x18\x03 \x03(\v2\x1a.blueapi.api.MonthlyBudgetR\rmonthlyBudget\"E\n" +
	"\rMonthlyBudget\x12\x1c\n" +
	"\tyearMonth\x18\x01 \x01(\tR\tyearMonth\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\x01R\x06amount\"B\n" +
	"\x10DailyBudgetAlert\x12\x14\n" +
	"\x05value\x18\x01 \x01(\x01R\x05value\x12\x18\n" +
	"\aenabled\x18\x02 \x01(\bR\aenabled\"N\n" +
	"\x1cDailyRateIncreaseBudgetAlert\x12\x14\n" +
	"\x05value\x18\x01 \x01(\x01R\x05value\x12\x18\n" +
	"\aenabled\x18\x02 \x01(\bR\aenabled\"D\n" +
	"\x12MonthlyBudgetAlert\x12\x14\n" +
	"\x05value\x18\x01 \x01(\x01R\x05value\x12\x18\n" +
	"\aenabled\x18\x02 \x01(\bR\aenabled\"O\n" +
	"\x17BudgetAlertNotification\x12\x1a\n" +
	"\bchannels\x18\x02 \x03(\tR\bchannels\x12\x18\n" +
	"\aenabled\x18\x03 \x01(\bR\aenabled\"w\n" +
	"\x1dBudgetAlertNotificationDetail\x12<\n" +
	"\bchannels\x18\x02 \x03(\v2 .blueapi.api.NotificationChannelR\bchannels\x12\x18\n" +
	"\aenabled\x18\x03 \x01(\bR\aenabledBS\n" +
	"\x19cloud.alphaus.blueapi.apiB\x0eApiBudgetProtoZ&github.com/alphauslabs/blue-sdk-go/apib\x06proto3"

var (
	file_api_budget_proto_rawDescOnce sync.Once
	file_api_budget_proto_rawDescData []byte
)

func file_api_budget_proto_rawDescGZIP() []byte {
	file_api_budget_proto_rawDescOnce.Do(func() {
		file_api_budget_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_budget_proto_rawDesc), len(file_api_budget_proto_rawDesc)))
	})
	return file_api_budget_proto_rawDescData
}

var file_api_budget_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_budget_proto_goTypes = []any{
	(*Budget)(nil),                        // 0: blueapi.api.Budget
	(*MonthlyBudget)(nil),                 // 1: blueapi.api.MonthlyBudget
	(*DailyBudgetAlert)(nil),              // 2: blueapi.api.DailyBudgetAlert
	(*DailyRateIncreaseBudgetAlert)(nil),  // 3: blueapi.api.DailyRateIncreaseBudgetAlert
	(*MonthlyBudgetAlert)(nil),            // 4: blueapi.api.MonthlyBudgetAlert
	(*BudgetAlertNotification)(nil),       // 5: blueapi.api.BudgetAlertNotification
	(*BudgetAlertNotificationDetail)(nil), // 6: blueapi.api.BudgetAlertNotificationDetail
	(*NotificationChannel)(nil),           // 7: blueapi.api.NotificationChannel
}
var file_api_budget_proto_depIdxs = []int32{
	1, // 0: blueapi.api.Budget.monthlyBudget:type_name -> blueapi.api.MonthlyBudget
	7, // 1: blueapi.api.BudgetAlertNotificationDetail.channels:type_name -> blueapi.api.NotificationChannel
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_budget_proto_init() }
func file_api_budget_proto_init() {
	if File_api_budget_proto != nil {
		return
	}
	file_api_notification_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_budget_proto_rawDesc), len(file_api_budget_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_budget_proto_goTypes,
		DependencyIndexes: file_api_budget_proto_depIdxs,
		MessageInfos:      file_api_budget_proto_msgTypes,
	}.Build()
	File_api_budget_proto = out.File
	file_api_budget_proto_goTypes = nil
	file_api_budget_proto_depIdxs = nil
}
