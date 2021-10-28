// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/grouprootuser.proto

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

// Describes the fields on an Alphaus root user for a billing group or access group.
type GroupRootUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The email address associated with the user. If there is no email
	// address provided then this field will contain "Not Set".
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// The password associated with the user. This field will only be
	// populated when the reseller is first created.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// The ID of the group to which the user belongs. This will either
	// be a billing group ID or an access group ID.
	GroupId string `protobuf:"bytes,3,opt,name=groupId,proto3" json:"groupId,omitempty"`
	// The name of the group to which the user belongs. This will either
	// be the name of a billing group or an access group.
	GroupName string `protobuf:"bytes,4,opt,name=groupName,proto3" json:"groupName,omitempty"`
	// This field describes what type of group to which the user belongs.
	// This field will contain either a value of "billing_group" or "access_group".
	GroupType string `protobuf:"bytes,5,opt,name=groupType,proto3" json:"groupType,omitempty"`
	// A collection of feature flags and whether or not they should be enabled.
	// For a full list of such flags, see https://alphauslabs.github.io/blueapi/apis/groups.html.
	Meta *FeatureFlags `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
	// The time when the password was last updated. This value will be null if
	// the password has never been updated.
	PasswordUpdateTime string `protobuf:"bytes,7,opt,name=passwordUpdateTime,proto3" json:"passwordUpdateTime,omitempty"`
	// The time when the user was last updated. This value will be null if the
	// user has never been updated.
	UpdateTime string `protobuf:"bytes,8,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	// An ID that uniquely identifies the user in the context of its access group.
	// If the user does not belong to an access group then this value field will
	// contain the same value as the user ID.
	UserAccessId string `protobuf:"bytes,9,opt,name=userAccessId,proto3" json:"userAccessId,omitempty"`
	// An ID that unqiuely identifies the user.
	UserId string `protobuf:"bytes,10,opt,name=userId,proto3" json:"userId,omitempty"`
	// How many days of Wave use remain. This value is only used for trial accounts.
	// This field cannot be updated.
	WaveAvailabilityDays int32 `protobuf:"varint,11,opt,name=waveAvailabilityDays,proto3" json:"waveAvailabilityDays,omitempty"`
	// When the user registered on Wave.
	WaveRegistered string `protobuf:"bytes,12,opt,name=waveRegistered,proto3" json:"waveRegistered,omitempty"`
	// The plan associated with the user. Possible values include "trial",
	// "limited3" and "limited4".
	WaveStatus string `protobuf:"bytes,13,opt,name=waveStatus,proto3" json:"waveStatus,omitempty"`
}

func (x *GroupRootUser) Reset() {
	*x = GroupRootUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grouprootuser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRootUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRootUser) ProtoMessage() {}

func (x *GroupRootUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_grouprootuser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRootUser.ProtoReflect.Descriptor instead.
func (*GroupRootUser) Descriptor() ([]byte, []int) {
	return file_api_grouprootuser_proto_rawDescGZIP(), []int{0}
}

func (x *GroupRootUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GroupRootUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GroupRootUser) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupRootUser) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GroupRootUser) GetGroupType() string {
	if x != nil {
		return x.GroupType
	}
	return ""
}

func (x *GroupRootUser) GetMeta() *FeatureFlags {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GroupRootUser) GetPasswordUpdateTime() string {
	if x != nil {
		return x.PasswordUpdateTime
	}
	return ""
}

func (x *GroupRootUser) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *GroupRootUser) GetUserAccessId() string {
	if x != nil {
		return x.UserAccessId
	}
	return ""
}

func (x *GroupRootUser) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GroupRootUser) GetWaveAvailabilityDays() int32 {
	if x != nil {
		return x.WaveAvailabilityDays
	}
	return 0
}

func (x *GroupRootUser) GetWaveRegistered() string {
	if x != nil {
		return x.WaveRegistered
	}
	return ""
}

func (x *GroupRootUser) GetWaveStatus() string {
	if x != nil {
		return x.WaveStatus
	}
	return ""
}

// Describes the features and whether or not they are enabled for a particular user.
type FeatureFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Control view-access to the Wave dashboard graph
	DashboardGraph bool `protobuf:"varint,1,opt,name=dashboard_graph,json=dashboardGraph,proto3" json:"dashboard_graph,omitempty"`
	// Control access to the Accounts pane under usage reports on Wave
	UsageAccount bool `protobuf:"varint,2,opt,name=usage_account,json=usageAccount,proto3" json:"usage_account,omitempty"`
	// Control view-access to the Account-level graph under usage reports on Wave
	UsageAccountGraph bool `protobuf:"varint,3,opt,name=usage_account_graph,json=usageAccountGraph,proto3" json:"usage_account_graph,omitempty"`
	// Control edit-access to Accounts under usage reports on Wave
	UsageAccountMenuAccountEdit bool `protobuf:"varint,4,opt,name=usage_account_menu_account_edit,json=usageAccountMenuAccountEdit,proto3" json:"usage_account_menu_account_edit,omitempty"`
	// Control view-access to budget settings for an Account under usage reports on Wave
	UsageAccountMenuBudget bool `protobuf:"varint,5,opt,name=usage_account_menu_budget,json=usageAccountMenuBudget,proto3" json:"usage_account_menu_budget,omitempty"`
	// Control edit-access to budget settings for an Account under usage reports on Wave
	UsageAccountMenuBudgetEdit bool `protobuf:"varint,6,opt,name=usage_account_menu_budget_edit,json=usageAccountMenuBudgetEdit,proto3" json:"usage_account_menu_budget_edit,omitempty"`
	// Control view-access to the fees menu for an Account under usage reports on Wave
	UsageAccountMenuFeesFee bool `protobuf:"varint,7,opt,name=usage_account_menu_fees_fee,json=usageAccountMenuFeesFee,proto3" json:"usage_account_menu_fees_fee,omitempty"`
	// Control view-access to credits for an Account under usage reports on Wave
	UsageAccountMenuFeesCredit bool `protobuf:"varint,8,opt,name=usage_account_menu_fees_credit,json=usageAccountMenuFeesCredit,proto3" json:"usage_account_menu_fees_credit,omitempty"`
	// Control view-access to refunds for an Account under usage reports on Wave
	UsageAccountMenuFeesRefund bool `protobuf:"varint,9,opt,name=usage_account_menu_fees_refund,json=usageAccountMenuFeesRefund,proto3" json:"usage_account_menu_fees_refund,omitempty"`
	// Control view-access to other fees for an Account under usage reports on Wave
	UsageAccountMenuFeesOtherFees bool `protobuf:"varint,10,opt,name=usage_account_menu_fees_other_fees,json=usageAccountMenuFeesOtherFees,proto3" json:"usage_account_menu_fees_other_fees,omitempty"`
	// Control download-access to usage report on Wave
	UsageReportDownload bool `protobuf:"varint,11,opt,name=usage_report_download,json=usageReportDownload,proto3" json:"usage_report_download,omitempty"`
	// Control access to the Groups pane under usage reports on Wave
	UsageGroup bool `protobuf:"varint,12,opt,name=usage_group,json=usageGroup,proto3" json:"usage_group,omitempty"`
	// Control view-access to the Group-level graph under usage reports on Wave
	UsageGroupGraph bool `protobuf:"varint,13,opt,name=usage_group_graph,json=usageGroupGraph,proto3" json:"usage_group_graph,omitempty"`
	// Control view-access to the Tags pane under usage reports on Wave
	UsageTag bool `protobuf:"varint,14,opt,name=usage_tag,json=usageTag,proto3" json:"usage_tag,omitempty"`
	// Control view-access to the Tag-level graph under usage reports on Wave
	UsageTagGraph bool `protobuf:"varint,15,opt,name=usage_tag_graph,json=usageTagGraph,proto3" json:"usage_tag_graph,omitempty"`
	// Control view-access to the Tags pane for crosstag groups under usage reports on Wave
	UsageCrosstag bool `protobuf:"varint,16,opt,name=usage_crosstag,json=usageCrosstag,proto3" json:"usage_crosstag,omitempty"`
	// Control view-access to the Tag-level graph for crosstag groups under usage reports on Wave
	UsageCrosstagGraph bool `protobuf:"varint,17,opt,name=usage_crosstag_graph,json=usageCrosstagGraph,proto3" json:"usage_crosstag_graph,omitempty"`
	// Control view-access to purchased reservations under reserved instances on Wave
	RiPurchased bool `protobuf:"varint,18,opt,name=ri_purchased,json=riPurchased,proto3" json:"ri_purchased,omitempty"`
	// Control access to RI utilization under reserved instances on Wave
	RiUtilization bool `protobuf:"varint,19,opt,name=ri_utilization,json=riUtilization,proto3" json:"ri_utilization,omitempty"`
	// Control access to RI recommendations on Wave
	RiRecommendation bool `protobuf:"varint,20,opt,name=ri_recommendation,json=riRecommendation,proto3" json:"ri_recommendation,omitempty"`
	// Control access to purchased savings plans on Wave
	SpPurchased bool `protobuf:"varint,21,opt,name=sp_purchased,json=spPurchased,proto3" json:"sp_purchased,omitempty"`
	// Control access to the invoice on Wave
	Invoice bool `protobuf:"varint,22,opt,name=invoice,proto3" json:"invoice,omitempty"`
	// Control download-access to the discounted invoice CSV on Wave
	InvoiceDownloadCsvDiscount bool `protobuf:"varint,23,opt,name=invoice_download_csv_discount,json=invoiceDownloadCsvDiscount,proto3" json:"invoice_download_csv_discount,omitempty"`
	// Control download-access to the merged invoice CSV on Wave
	InvoiceDownloadCsvMerged bool `protobuf:"varint,24,opt,name=invoice_download_csv_merged,json=invoiceDownloadCsvMerged,proto3" json:"invoice_download_csv_merged,omitempty"`
	// Control access to API tokens on Wave
	OpenApi bool `protobuf:"varint,25,opt,name=open_api,json=openApi,proto3" json:"open_api,omitempty"`
	// Control access to sub-user management on Wave
	UsersManagement bool `protobuf:"varint,26,opt,name=users_management,json=usersManagement,proto3" json:"users_management,omitempty"`
	// Control access to the dashboard on Aqua
	AqDashboard bool `protobuf:"varint,27,opt,name=aq_dashboard,json=aqDashboard,proto3" json:"aq_dashboard,omitempty"`
	// Control access to the Coverage Ratio pane on Aqua
	AqCoverageRatio bool `protobuf:"varint,28,opt,name=aq_coverage_ratio,json=aqCoverageRatio,proto3" json:"aq_coverage_ratio,omitempty"`
	// Control access to the RI management pane on Aqua
	AqRiManagement bool `protobuf:"varint,29,opt,name=aq_ri_management,json=aqRiManagement,proto3" json:"aq_ri_management,omitempty"`
	// Control access to the savings plan management pane on Aqua
	AqSpManagement bool `protobuf:"varint,30,opt,name=aq_sp_management,json=aqSpManagement,proto3" json:"aq_sp_management,omitempty"`
	// Control access to RI and savings plan recommendations on Aqua
	AqRiSpInstances bool `protobuf:"varint,31,opt,name=aq_ri_sp_instances,json=aqRiSpInstances,proto3" json:"aq_ri_sp_instances,omitempty"`
	// Control access to right-sizing on Aqua
	AqRightSizing bool `protobuf:"varint,32,opt,name=aq_right_sizing,json=aqRightSizing,proto3" json:"aq_right_sizing,omitempty"`
	// Control access to scheduling on Aqua
	AqScheduling bool `protobuf:"varint,33,opt,name=aq_scheduling,json=aqScheduling,proto3" json:"aq_scheduling,omitempty"`
}

func (x *FeatureFlags) Reset() {
	*x = FeatureFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grouprootuser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureFlags) ProtoMessage() {}

func (x *FeatureFlags) ProtoReflect() protoreflect.Message {
	mi := &file_api_grouprootuser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureFlags.ProtoReflect.Descriptor instead.
func (*FeatureFlags) Descriptor() ([]byte, []int) {
	return file_api_grouprootuser_proto_rawDescGZIP(), []int{1}
}

func (x *FeatureFlags) GetDashboardGraph() bool {
	if x != nil {
		return x.DashboardGraph
	}
	return false
}

func (x *FeatureFlags) GetUsageAccount() bool {
	if x != nil {
		return x.UsageAccount
	}
	return false
}

func (x *FeatureFlags) GetUsageAccountGraph() bool {
	if x != nil {
		return x.UsageAccountGraph
	}
	return false
}

func (x *FeatureFlags) GetUsageAccountMenuAccountEdit() bool {
	if x != nil {
		return x.UsageAccountMenuAccountEdit
	}
	return false
}

func (x *FeatureFlags) GetUsageAccountMenuBudget() bool {
	if x != nil {
		return x.UsageAccountMenuBudget
	}
	return false
}

func (x *FeatureFlags) GetUsageAccountMenuBudgetEdit() bool {
	if x != nil {
		return x.UsageAccountMenuBudgetEdit
	}
	return false
}

func (x *FeatureFlags) GetUsageAccountMenuFeesFee() bool {
	if x != nil {
		return x.UsageAccountMenuFeesFee
	}
	return false
}

func (x *FeatureFlags) GetUsageAccountMenuFeesCredit() bool {
	if x != nil {
		return x.UsageAccountMenuFeesCredit
	}
	return false
}

func (x *FeatureFlags) GetUsageAccountMenuFeesRefund() bool {
	if x != nil {
		return x.UsageAccountMenuFeesRefund
	}
	return false
}

func (x *FeatureFlags) GetUsageAccountMenuFeesOtherFees() bool {
	if x != nil {
		return x.UsageAccountMenuFeesOtherFees
	}
	return false
}

func (x *FeatureFlags) GetUsageReportDownload() bool {
	if x != nil {
		return x.UsageReportDownload
	}
	return false
}

func (x *FeatureFlags) GetUsageGroup() bool {
	if x != nil {
		return x.UsageGroup
	}
	return false
}

func (x *FeatureFlags) GetUsageGroupGraph() bool {
	if x != nil {
		return x.UsageGroupGraph
	}
	return false
}

func (x *FeatureFlags) GetUsageTag() bool {
	if x != nil {
		return x.UsageTag
	}
	return false
}

func (x *FeatureFlags) GetUsageTagGraph() bool {
	if x != nil {
		return x.UsageTagGraph
	}
	return false
}

func (x *FeatureFlags) GetUsageCrosstag() bool {
	if x != nil {
		return x.UsageCrosstag
	}
	return false
}

func (x *FeatureFlags) GetUsageCrosstagGraph() bool {
	if x != nil {
		return x.UsageCrosstagGraph
	}
	return false
}

func (x *FeatureFlags) GetRiPurchased() bool {
	if x != nil {
		return x.RiPurchased
	}
	return false
}

func (x *FeatureFlags) GetRiUtilization() bool {
	if x != nil {
		return x.RiUtilization
	}
	return false
}

func (x *FeatureFlags) GetRiRecommendation() bool {
	if x != nil {
		return x.RiRecommendation
	}
	return false
}

func (x *FeatureFlags) GetSpPurchased() bool {
	if x != nil {
		return x.SpPurchased
	}
	return false
}

func (x *FeatureFlags) GetInvoice() bool {
	if x != nil {
		return x.Invoice
	}
	return false
}

func (x *FeatureFlags) GetInvoiceDownloadCsvDiscount() bool {
	if x != nil {
		return x.InvoiceDownloadCsvDiscount
	}
	return false
}

func (x *FeatureFlags) GetInvoiceDownloadCsvMerged() bool {
	if x != nil {
		return x.InvoiceDownloadCsvMerged
	}
	return false
}

func (x *FeatureFlags) GetOpenApi() bool {
	if x != nil {
		return x.OpenApi
	}
	return false
}

func (x *FeatureFlags) GetUsersManagement() bool {
	if x != nil {
		return x.UsersManagement
	}
	return false
}

func (x *FeatureFlags) GetAqDashboard() bool {
	if x != nil {
		return x.AqDashboard
	}
	return false
}

func (x *FeatureFlags) GetAqCoverageRatio() bool {
	if x != nil {
		return x.AqCoverageRatio
	}
	return false
}

func (x *FeatureFlags) GetAqRiManagement() bool {
	if x != nil {
		return x.AqRiManagement
	}
	return false
}

func (x *FeatureFlags) GetAqSpManagement() bool {
	if x != nil {
		return x.AqSpManagement
	}
	return false
}

func (x *FeatureFlags) GetAqRiSpInstances() bool {
	if x != nil {
		return x.AqRiSpInstances
	}
	return false
}

func (x *FeatureFlags) GetAqRightSizing() bool {
	if x != nil {
		return x.AqRightSizing
	}
	return false
}

func (x *FeatureFlags) GetAqScheduling() bool {
	if x != nil {
		return x.AqScheduling
	}
	return false
}

var File_api_grouprootuser_proto protoreflect.FileDescriptor

var file_api_grouprootuser_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x72, 0x6f, 0x6f, 0x74, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x22, 0xce, 0x03, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x6f, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x2e, 0x0a, 0x12, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x77,
	0x61, 0x76, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44,
	0x61, 0x79, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x77, 0x61, 0x76, 0x65, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x61, 0x79, 0x73, 0x12,
	0x26, 0x0a, 0x0e, 0x77, 0x61, 0x76, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x61, 0x76, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x76, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x61, 0x76,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9a, 0x0c, 0x0a, 0x0c, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x73, 0x61, 0x67, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x75, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x44, 0x0a, 0x1f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x6e,
	0x75, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x64, 0x69, 0x74, 0x12, 0x39, 0x0a, 0x19,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x6e, 0x75, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x16, 0x75, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x6e,
	0x75, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x42, 0x0a, 0x1e, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x62, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x6e,
	0x75, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x45, 0x64, 0x69, 0x74, 0x12, 0x3c, 0x0a, 0x1b, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6e,
	0x75, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x17, 0x75, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x46, 0x65, 0x65, 0x73, 0x46, 0x65, 0x65, 0x12, 0x42, 0x0a, 0x1e, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f,
	0x66, 0x65, 0x65, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x1a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d,
	0x65, 0x6e, 0x75, 0x46, 0x65, 0x65, 0x73, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x42, 0x0a,
	0x1e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x6e, 0x75, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x46, 0x65, 0x65, 0x73, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x12, 0x49, 0x0a, 0x22, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x5f, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x46,
	0x65, 0x65, 0x73, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x65, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x15,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x67, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73,
	0x73, 0x74, 0x61, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x74, 0x61, 0x67, 0x12, 0x30, 0x0a, 0x14, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x74, 0x61, 0x67, 0x5f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x75, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72,
	0x6f, 0x73, 0x73, 0x74, 0x61, 0x67, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x69, 0x5f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x72, 0x69, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x69, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x69, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x69, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x10, 0x72, 0x69, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70, 0x5f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x70, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12,
	0x41, 0x0a, 0x1d, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x63, 0x73, 0x76, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x73, 0x76, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x73, 0x76, 0x5f, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x73, 0x76, 0x4d, 0x65, 0x72, 0x67, 0x65,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x41, 0x70, 0x69, 0x12, 0x29, 0x0a, 0x10,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x71, 0x5f, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61,
	0x71, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x71,
	0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18,
	0x1c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x71, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x71, 0x5f, 0x72, 0x69, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x61, 0x71, 0x52, 0x69, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x28, 0x0a, 0x10, 0x61, 0x71, 0x5f, 0x73, 0x70, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x71, 0x53, 0x70,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x12, 0x61, 0x71,
	0x5f, 0x72, 0x69, 0x5f, 0x73, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x71, 0x52, 0x69, 0x53, 0x70, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x71, 0x5f, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x61, 0x71, 0x52, 0x69, 0x67, 0x68, 0x74, 0x53, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x71, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67,
	0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x71, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x69, 0x6e, 0x67, 0x42, 0x5a, 0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x42, 0x15, 0x41, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x6f, 0x6f, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grouprootuser_proto_rawDescOnce sync.Once
	file_api_grouprootuser_proto_rawDescData = file_api_grouprootuser_proto_rawDesc
)

func file_api_grouprootuser_proto_rawDescGZIP() []byte {
	file_api_grouprootuser_proto_rawDescOnce.Do(func() {
		file_api_grouprootuser_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grouprootuser_proto_rawDescData)
	})
	return file_api_grouprootuser_proto_rawDescData
}

var file_api_grouprootuser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_grouprootuser_proto_goTypes = []interface{}{
	(*GroupRootUser)(nil), // 0: blueapi.api.GroupRootUser
	(*FeatureFlags)(nil),  // 1: blueapi.api.FeatureFlags
}
var file_api_grouprootuser_proto_depIdxs = []int32{
	1, // 0: blueapi.api.GroupRootUser.meta:type_name -> blueapi.api.FeatureFlags
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_grouprootuser_proto_init() }
func file_api_grouprootuser_proto_init() {
	if File_api_grouprootuser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grouprootuser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRootUser); i {
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
		file_api_grouprootuser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureFlags); i {
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
			RawDescriptor: file_api_grouprootuser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_grouprootuser_proto_goTypes,
		DependencyIndexes: file_api_grouprootuser_proto_depIdxs,
		MessageInfos:      file_api_grouprootuser_proto_msgTypes,
	}.Build()
	File_api_grouprootuser_proto = out.File
	file_api_grouprootuser_proto_rawDesc = nil
	file_api_grouprootuser_proto_goTypes = nil
	file_api_grouprootuser_proto_depIdxs = nil
}
