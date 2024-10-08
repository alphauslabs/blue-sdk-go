// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/cover/asset.proto

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

type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Possible values: `account`, `subscription` or `project`
	Key        string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	mi := &file_api_cover_asset_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_asset_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_api_cover_asset_proto_rawDescGZIP(), []int{0}
}

func (x *Asset) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Asset) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type EC2Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Capacity Reservation.
	CapacityReservationId string `protobuf:"bytes,1,opt,name=capacityReservationId,proto3" json:"capacityReservationId,omitempty"`
	// Indicates whether the instance is optimized for Amazon EBS I/O. This optimization
	// provides dedicated throughput to Amazon EBS and an optimized configuration
	// stack to provide optimal I/O performance. This optimization isn't available
	// with all instance types. Additional usage charges apply when using an EBS
	// Optimized instance.
	EbsOptimized bool `protobuf:"varint,2,opt,name=ebsOptimized,proto3" json:"ebsOptimized,omitempty"`
	// The ID of the AMI used to launch the instance.
	ImageId string `protobuf:"bytes,3,opt,name=imageId,proto3" json:"imageId,omitempty"`
	// The ID of the instance.
	InstanceId string `protobuf:"bytes,4,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	// Indicates whether this is a Spot Instance or a Scheduled Instance.
	InstanceLifecycle string `protobuf:"bytes,5,opt,name=instanceLifecycle,proto3" json:"instanceLifecycle,omitempty"`
	// The instance type.
	InstanceType string `protobuf:"bytes,6,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	// The IPv6 address assigned to the instance.
	Ipv6Address string `protobuf:"bytes,7,opt,name=ipv6Address,proto3" json:"ipv6Address,omitempty"`
	// The kernel associated with this instance, if applicable.
	KernelId string `protobuf:"bytes,8,opt,name=kernelId,proto3" json:"kernelId,omitempty"`
	// The name of the key pair, if this instance was launched with an associated key pair.
	KeyName string `protobuf:"bytes,9,opt,name=keyName,proto3" json:"keyName,omitempty"`
	// The time the instance was launched.
	LaunchTime string `protobuf:"bytes,10,opt,name=launchTime,proto3" json:"launchTime,omitempty"`
	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn string `protobuf:"bytes,11,opt,name=outpostArn,proto3" json:"outpostArn,omitempty"`
	// The value is Windows for Windows instances; otherwise blank.
	Platform string `protobuf:"bytes,12,opt,name=platform,proto3" json:"platform,omitempty"`
	// The platform details value for the instance. For more information, see AMI
	// billing information fields (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/billing-info-fields.html)
	// in the Amazon EC2 User Guide.
	PlatformDetails string `protobuf:"bytes,13,opt,name=platformDetails,proto3" json:"platformDetails,omitempty"`
	// (IPv4 only) The private DNS hostname name assigned to the instance. This
	// DNS hostname can only be used inside the Amazon EC2 network. This name is
	// not available until the instance enters the running state.
	//
	// [EC2-VPC] The Amazon-provided DNS server resolves Amazon-provided private
	// DNS hostnames if you've enabled DNS resolution and DNS hostnames in your
	// VPC. If you are not using the Amazon-provided DNS server in your VPC, your
	// custom domain name servers must resolve the hostname as appropriate.
	PrivateDnsName string `protobuf:"bytes,14,opt,name=privateDnsName,proto3" json:"privateDnsName,omitempty"`
	// The private IPv4 address assigned to the instance.
	PrivateIpAddress string `protobuf:"bytes,15,opt,name=privateIpAddress,proto3" json:"privateIpAddress,omitempty"`
	// (IPv4 only) The public DNS name assigned to the instance. This name is not
	// available until the instance enters the running state. For EC2-VPC, this
	// name is only available if you've enabled DNS hostnames for your VPC.
	PublicDnsName string `protobuf:"bytes,16,opt,name=publicDnsName,proto3" json:"publicDnsName,omitempty"`
	// The public IPv4 address, or the Carrier IP address assigned to the instance,
	// if applicable.
	//
	// A Carrier IP address only applies to an instance launched in a subnet associated
	// with a Wavelength Zone.
	PublicIpAddress string `protobuf:"bytes,17,opt,name=publicIpAddress,proto3" json:"publicIpAddress,omitempty"`
	// The RAM disk associated with this instance, if applicable.
	RamdiskId string `protobuf:"bytes,18,opt,name=ramdiskId,proto3" json:"ramdiskId,omitempty"`
	// The device name of the root device volume (for example, /dev/sda1).
	RootDeviceName string `protobuf:"bytes,19,opt,name=rootDeviceName,proto3" json:"rootDeviceName,omitempty"`
	// The root device type used by the AMI. The AMI can use an EBS volume or an
	// instance store volume.
	RootDeviceType string `protobuf:"bytes,20,opt,name=rootDeviceType,proto3" json:"rootDeviceType,omitempty"`
	// Indicates whether source/destination checking is enabled.
	SourceDestCheck bool `protobuf:"varint,21,opt,name=sourceDestCheck,proto3" json:"sourceDestCheck,omitempty"`
	// If the request is a Spot Instance request, the ID of the request.
	SpotInstanceRequestId string `protobuf:"bytes,22,opt,name=spotInstanceRequestId,proto3" json:"spotInstanceRequestId,omitempty"`
	// Specifies whether enhanced networking with the Intel 82599 Virtual Function
	// interface is enabled.
	SriovNetSupport string `protobuf:"bytes,23,opt,name=sriovNetSupport,proto3" json:"sriovNetSupport,omitempty"`
	// The current state of the instance.
	State string `protobuf:"bytes,24,opt,name=state,proto3" json:"state,omitempty"`
	// [EC2-VPC] The ID of the subnet in which the instance is running.
	SubnetId string `protobuf:"bytes,25,opt,name=subnetId,proto3" json:"subnetId,omitempty"`
	// Any tags assigned to the instance.
	Tags []*Tag `protobuf:"bytes,26,rep,name=tags,proto3" json:"tags,omitempty"`
	// The usage operation value for the instance. For more information, see AMI
	// billing information fields (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/billing-info-fields.html)
	// in the Amazon EC2 User Guide.
	UsageOperation string `protobuf:"bytes,27,opt,name=usageOperation,proto3" json:"usageOperation,omitempty"`
	// The virtualization type of the instance.
	VirtualizationType string `protobuf:"bytes,28,opt,name=virtualizationType,proto3" json:"virtualizationType,omitempty"`
	// [EC2-VPC] The ID of the VPC in which the instance is running.
	VpcId string `protobuf:"bytes,29,opt,name=vpcId,proto3" json:"vpcId,omitempty"`
	// Format: `yyyy-mm`
	Date string `protobuf:"bytes,30,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *EC2Instance) Reset() {
	*x = EC2Instance{}
	mi := &file_api_cover_asset_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EC2Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EC2Instance) ProtoMessage() {}

func (x *EC2Instance) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_asset_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EC2Instance.ProtoReflect.Descriptor instead.
func (*EC2Instance) Descriptor() ([]byte, []int) {
	return file_api_cover_asset_proto_rawDescGZIP(), []int{1}
}

func (x *EC2Instance) GetCapacityReservationId() string {
	if x != nil {
		return x.CapacityReservationId
	}
	return ""
}

func (x *EC2Instance) GetEbsOptimized() bool {
	if x != nil {
		return x.EbsOptimized
	}
	return false
}

func (x *EC2Instance) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *EC2Instance) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *EC2Instance) GetInstanceLifecycle() string {
	if x != nil {
		return x.InstanceLifecycle
	}
	return ""
}

func (x *EC2Instance) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *EC2Instance) GetIpv6Address() string {
	if x != nil {
		return x.Ipv6Address
	}
	return ""
}

func (x *EC2Instance) GetKernelId() string {
	if x != nil {
		return x.KernelId
	}
	return ""
}

func (x *EC2Instance) GetKeyName() string {
	if x != nil {
		return x.KeyName
	}
	return ""
}

func (x *EC2Instance) GetLaunchTime() string {
	if x != nil {
		return x.LaunchTime
	}
	return ""
}

func (x *EC2Instance) GetOutpostArn() string {
	if x != nil {
		return x.OutpostArn
	}
	return ""
}

func (x *EC2Instance) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *EC2Instance) GetPlatformDetails() string {
	if x != nil {
		return x.PlatformDetails
	}
	return ""
}

func (x *EC2Instance) GetPrivateDnsName() string {
	if x != nil {
		return x.PrivateDnsName
	}
	return ""
}

func (x *EC2Instance) GetPrivateIpAddress() string {
	if x != nil {
		return x.PrivateIpAddress
	}
	return ""
}

func (x *EC2Instance) GetPublicDnsName() string {
	if x != nil {
		return x.PublicDnsName
	}
	return ""
}

func (x *EC2Instance) GetPublicIpAddress() string {
	if x != nil {
		return x.PublicIpAddress
	}
	return ""
}

func (x *EC2Instance) GetRamdiskId() string {
	if x != nil {
		return x.RamdiskId
	}
	return ""
}

func (x *EC2Instance) GetRootDeviceName() string {
	if x != nil {
		return x.RootDeviceName
	}
	return ""
}

func (x *EC2Instance) GetRootDeviceType() string {
	if x != nil {
		return x.RootDeviceType
	}
	return ""
}

func (x *EC2Instance) GetSourceDestCheck() bool {
	if x != nil {
		return x.SourceDestCheck
	}
	return false
}

func (x *EC2Instance) GetSpotInstanceRequestId() string {
	if x != nil {
		return x.SpotInstanceRequestId
	}
	return ""
}

func (x *EC2Instance) GetSriovNetSupport() string {
	if x != nil {
		return x.SriovNetSupport
	}
	return ""
}

func (x *EC2Instance) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *EC2Instance) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *EC2Instance) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *EC2Instance) GetUsageOperation() string {
	if x != nil {
		return x.UsageOperation
	}
	return ""
}

func (x *EC2Instance) GetVirtualizationType() string {
	if x != nil {
		return x.VirtualizationType
	}
	return ""
}

func (x *EC2Instance) GetVpcId() string {
	if x != nil {
		return x.VpcId
	}
	return ""
}

func (x *EC2Instance) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	mi := &file_api_cover_asset_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_api_cover_asset_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_api_cover_asset_proto_rawDescGZIP(), []int{2}
}

func (x *Tag) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Tag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_api_cover_asset_proto protoreflect.FileDescriptor

var file_api_cover_asset_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x22, 0xa2, 0x01, 0x0a, 0x05, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x48, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xcd, 0x08, 0x0a, 0x0b, 0x45, 0x43, 0x32, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x15, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x62, 0x73, 0x4f, 0x70, 0x74, 0x69,
	0x6d, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x62, 0x73,
	0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x70, 0x76, 0x36, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x76, 0x36,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6f, 0x75, 0x74, 0x70, 0x6f, 0x73, 0x74, 0x41, 0x72, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x6f, 0x73, 0x74, 0x41, 0x72, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x6e,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x44, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x44, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x44, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x61, 0x6d, 0x64, 0x69,
	0x73, 0x6b, 0x49, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x61, 0x6d, 0x64,
	0x69, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x6f, 0x6f, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x6f, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x72, 0x6f, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x6f, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x65, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x34, 0x0a, 0x15, 0x73, 0x70, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x73, 0x70, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x72, 0x69, 0x6f, 0x76, 0x4e, 0x65,
	0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x73, 0x72, 0x69, 0x6f, 0x76, 0x4e, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49,
	0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x26, 0x0a,
	0x0e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x70, 0x63, 0x49, 0x64, 0x18, 0x1d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x70, 0x63, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x2d, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x63,
	0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e,
	0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x42, 0x12, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c,
	0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cover_asset_proto_rawDescOnce sync.Once
	file_api_cover_asset_proto_rawDescData = file_api_cover_asset_proto_rawDesc
)

func file_api_cover_asset_proto_rawDescGZIP() []byte {
	file_api_cover_asset_proto_rawDescOnce.Do(func() {
		file_api_cover_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cover_asset_proto_rawDescData)
	})
	return file_api_cover_asset_proto_rawDescData
}

var file_api_cover_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_cover_asset_proto_goTypes = []any{
	(*Asset)(nil),       // 0: blueapi.api.cover.Asset
	(*EC2Instance)(nil), // 1: blueapi.api.cover.EC2Instance
	(*Tag)(nil),         // 2: blueapi.api.cover.Tag
	nil,                 // 3: blueapi.api.cover.Asset.AttributesEntry
}
var file_api_cover_asset_proto_depIdxs = []int32{
	3, // 0: blueapi.api.cover.Asset.attributes:type_name -> blueapi.api.cover.Asset.AttributesEntry
	2, // 1: blueapi.api.cover.EC2Instance.tags:type_name -> blueapi.api.cover.Tag
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_cover_asset_proto_init() }
func file_api_cover_asset_proto_init() {
	if File_api_cover_asset_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cover_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cover_asset_proto_goTypes,
		DependencyIndexes: file_api_cover_asset_proto_depIdxs,
		MessageInfos:      file_api_cover_asset_proto_msgTypes,
	}.Build()
	File_api_cover_asset_proto = out.File
	file_api_cover_asset_proto_rawDesc = nil
	file_api_cover_asset_proto_goTypes = nil
	file_api_cover_asset_proto_depIdxs = nil
}
