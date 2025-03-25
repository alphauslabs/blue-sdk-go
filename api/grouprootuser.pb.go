// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/grouprootuser.proto

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

// Describes the fields on an Alphaus root user for a billing group or access group.
type GroupRootUser struct {
	state protoimpl.MessageState `protogen:"open.v1"`
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
	WaveStatus    string `protobuf:"bytes,13,opt,name=waveStatus,proto3" json:"waveStatus,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupRootUser) Reset() {
	*x = GroupRootUser{}
	mi := &file_api_grouprootuser_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupRootUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRootUser) ProtoMessage() {}

func (x *GroupRootUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_grouprootuser_proto_msgTypes[0]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
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
	// Control access to the Coverage Ratio pane on Aqua
	AqCoverageRatio bool `protobuf:"varint,27,opt,name=aq_coverage_ratio,json=aqCoverageRatio,proto3" json:"aq_coverage_ratio,omitempty"`
	// Control access to the RI management pane on Aqua
	AqRiManagement bool `protobuf:"varint,28,opt,name=aq_ri_management,json=aqRiManagement,proto3" json:"aq_ri_management,omitempty"`
	// Control access to the savings plan management pane on Aqua
	AqSpManagement bool `protobuf:"varint,29,opt,name=aq_sp_management,json=aqSpManagement,proto3" json:"aq_sp_management,omitempty"`
	// Control access to RI and savings plan recommendations on Aqua
	AqRiSpInstances bool `protobuf:"varint,30,opt,name=aq_ri_sp_instances,json=aqRiSpInstances,proto3" json:"aq_ri_sp_instances,omitempty"`
	// Control access to right-sizing on Aqua
	AqRightSizing bool `protobuf:"varint,31,opt,name=aq_right_sizing,json=aqRightSizing,proto3" json:"aq_right_sizing,omitempty"`
	// Control access to scheduling on Aqua
	AqScheduling bool `protobuf:"varint,32,opt,name=aq_scheduling,json=aqScheduling,proto3" json:"aq_scheduling,omitempty"`
	// Control access to the report filters pane in Wave Pro
	ReportFilters bool `protobuf:"varint,33,opt,name=report_filters,json=reportFilters,proto3" json:"report_filters,omitempty"`
	// Control access to Aqua from Wave Pro
	AquaLink bool `protobuf:"varint,34,opt,name=aqua_link,json=aquaLink,proto3" json:"aqua_link,omitempty"`
	// Control access to budget alerts from Wave Pro
	Budgetalerts  bool `protobuf:"varint,35,opt,name=budgetalerts,proto3" json:"budgetalerts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FeatureFlags) Reset() {
	*x = FeatureFlags{}
	mi := &file_api_grouprootuser_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureFlags) ProtoMessage() {}

func (x *FeatureFlags) ProtoReflect() protoreflect.Message {
	mi := &file_api_grouprootuser_proto_msgTypes[1]
	if x != nil {
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

func (x *FeatureFlags) GetReportFilters() bool {
	if x != nil {
		return x.ReportFilters
	}
	return false
}

func (x *FeatureFlags) GetAquaLink() bool {
	if x != nil {
		return x.AquaLink
	}
	return false
}

func (x *FeatureFlags) GetBudgetalerts() bool {
	if x != nil {
		return x.Budgetalerts
	}
	return false
}

var File_api_grouprootuser_proto protoreflect.FileDescriptor

const file_api_grouprootuser_proto_rawDesc = "" +
	"\n" +
	"\x17api/grouprootuser.proto\x12\vblueapi.api\"\xce\x03\n" +
	"\rGroupRootUser\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x18\n" +
	"\agroupId\x18\x03 \x01(\tR\agroupId\x12\x1c\n" +
	"\tgroupName\x18\x04 \x01(\tR\tgroupName\x12\x1c\n" +
	"\tgroupType\x18\x05 \x01(\tR\tgroupType\x12-\n" +
	"\x04meta\x18\x06 \x01(\v2\x19.blueapi.api.FeatureFlagsR\x04meta\x12.\n" +
	"\x12passwordUpdateTime\x18\a \x01(\tR\x12passwordUpdateTime\x12\x1e\n" +
	"\n" +
	"updateTime\x18\b \x01(\tR\n" +
	"updateTime\x12\"\n" +
	"\fuserAccessId\x18\t \x01(\tR\fuserAccessId\x12\x16\n" +
	"\x06userId\x18\n" +
	" \x01(\tR\x06userId\x122\n" +
	"\x14waveAvailabilityDays\x18\v \x01(\x05R\x14waveAvailabilityDays\x12&\n" +
	"\x0ewaveRegistered\x18\f \x01(\tR\x0ewaveRegistered\x12\x1e\n" +
	"\n" +
	"waveStatus\x18\r \x01(\tR\n" +
	"waveStatus\"\xdf\f\n" +
	"\fFeatureFlags\x12'\n" +
	"\x0fdashboard_graph\x18\x01 \x01(\bR\x0edashboardGraph\x12#\n" +
	"\rusage_account\x18\x02 \x01(\bR\fusageAccount\x12.\n" +
	"\x13usage_account_graph\x18\x03 \x01(\bR\x11usageAccountGraph\x12D\n" +
	"\x1fusage_account_menu_account_edit\x18\x04 \x01(\bR\x1busageAccountMenuAccountEdit\x129\n" +
	"\x19usage_account_menu_budget\x18\x05 \x01(\bR\x16usageAccountMenuBudget\x12B\n" +
	"\x1eusage_account_menu_budget_edit\x18\x06 \x01(\bR\x1ausageAccountMenuBudgetEdit\x12<\n" +
	"\x1busage_account_menu_fees_fee\x18\a \x01(\bR\x17usageAccountMenuFeesFee\x12B\n" +
	"\x1eusage_account_menu_fees_credit\x18\b \x01(\bR\x1ausageAccountMenuFeesCredit\x12B\n" +
	"\x1eusage_account_menu_fees_refund\x18\t \x01(\bR\x1ausageAccountMenuFeesRefund\x12I\n" +
	"\"usage_account_menu_fees_other_fees\x18\n" +
	" \x01(\bR\x1dusageAccountMenuFeesOtherFees\x122\n" +
	"\x15usage_report_download\x18\v \x01(\bR\x13usageReportDownload\x12\x1f\n" +
	"\vusage_group\x18\f \x01(\bR\n" +
	"usageGroup\x12*\n" +
	"\x11usage_group_graph\x18\r \x01(\bR\x0fusageGroupGraph\x12\x1b\n" +
	"\tusage_tag\x18\x0e \x01(\bR\busageTag\x12&\n" +
	"\x0fusage_tag_graph\x18\x0f \x01(\bR\rusageTagGraph\x12%\n" +
	"\x0eusage_crosstag\x18\x10 \x01(\bR\rusageCrosstag\x120\n" +
	"\x14usage_crosstag_graph\x18\x11 \x01(\bR\x12usageCrosstagGraph\x12!\n" +
	"\fri_purchased\x18\x12 \x01(\bR\vriPurchased\x12%\n" +
	"\x0eri_utilization\x18\x13 \x01(\bR\rriUtilization\x12+\n" +
	"\x11ri_recommendation\x18\x14 \x01(\bR\x10riRecommendation\x12!\n" +
	"\fsp_purchased\x18\x15 \x01(\bR\vspPurchased\x12\x18\n" +
	"\ainvoice\x18\x16 \x01(\bR\ainvoice\x12A\n" +
	"\x1dinvoice_download_csv_discount\x18\x17 \x01(\bR\x1ainvoiceDownloadCsvDiscount\x12=\n" +
	"\x1binvoice_download_csv_merged\x18\x18 \x01(\bR\x18invoiceDownloadCsvMerged\x12\x19\n" +
	"\bopen_api\x18\x19 \x01(\bR\aopenApi\x12)\n" +
	"\x10users_management\x18\x1a \x01(\bR\x0fusersManagement\x12*\n" +
	"\x11aq_coverage_ratio\x18\x1b \x01(\bR\x0faqCoverageRatio\x12(\n" +
	"\x10aq_ri_management\x18\x1c \x01(\bR\x0eaqRiManagement\x12(\n" +
	"\x10aq_sp_management\x18\x1d \x01(\bR\x0eaqSpManagement\x12+\n" +
	"\x12aq_ri_sp_instances\x18\x1e \x01(\bR\x0faqRiSpInstances\x12&\n" +
	"\x0faq_right_sizing\x18\x1f \x01(\bR\raqRightSizing\x12#\n" +
	"\raq_scheduling\x18  \x01(\bR\faqScheduling\x12%\n" +
	"\x0ereport_filters\x18! \x01(\bR\rreportFilters\x12\x1b\n" +
	"\taqua_link\x18\" \x01(\bR\baquaLink\x12\"\n" +
	"\fbudgetalerts\x18# \x01(\bR\fbudgetalertsBZ\n" +
	"\x19cloud.alphaus.blueapi.apiB\x15ApiGroupRootUserProtoZ&github.com/alphauslabs/blue-sdk-go/apib\x06proto3"

var (
	file_api_grouprootuser_proto_rawDescOnce sync.Once
	file_api_grouprootuser_proto_rawDescData []byte
)

func file_api_grouprootuser_proto_rawDescGZIP() []byte {
	file_api_grouprootuser_proto_rawDescOnce.Do(func() {
		file_api_grouprootuser_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_grouprootuser_proto_rawDesc), len(file_api_grouprootuser_proto_rawDesc)))
	})
	return file_api_grouprootuser_proto_rawDescData
}

var file_api_grouprootuser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_grouprootuser_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_grouprootuser_proto_rawDesc), len(file_api_grouprootuser_proto_rawDesc)),
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
	file_api_grouprootuser_proto_goTypes = nil
	file_api_grouprootuser_proto_depIdxs = nil
}
