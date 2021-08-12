// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The physical capacity of the Amazon Web Services Snow Family device.
type Capacity struct {

	// The amount of capacity available for use on the device.
	Available *int64

	// The name of the type of capacity, such as memory.
	Name *string

	// The total capacity on the device.
	Total *int64

	// The unit of measure for the type of capacity.
	Unit *string

	// The amount of capacity used on the device.
	Used *int64

	noSmithyDocumentSerde
}

// The command given to the device to execute.
//
// The following types satisfy this interface:
//  CommandMemberUnlock
//  CommandMemberReboot
type Command interface {
	isCommand()
}

// Unlocks the device.
type CommandMemberUnlock struct {
	Value Unlock

	noSmithyDocumentSerde
}

func (*CommandMemberUnlock) isCommand() {}

// Reboots the device.
type CommandMemberReboot struct {
	Value Reboot

	noSmithyDocumentSerde
}

func (*CommandMemberReboot) isCommand() {}

// The options for how a device's CPU is configured.
type CpuOptions struct {

	// The number of cores that the CPU can use.
	CoreCount *int32

	// The number of threads per core in the CPU.
	ThreadsPerCore *int32

	noSmithyDocumentSerde
}

// Identifying information about the device.
type DeviceSummary struct {

	// The ID of the job used to order the device.
	AssociatedWithJob *string

	// The Amazon Resource Name (ARN) of the device.
	ManagedDeviceArn *string

	// The ID of the device.
	ManagedDeviceId *string

	// Optional metadata that you assign to a resource. You can use tags to categorize
	// a resource in different ways, such as by purpose, owner, or environment.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Describes a parameter used to set up an Amazon Elastic Block Store (Amazon EBS)
// volume in a block device mapping.
type EbsInstanceBlockDevice struct {

	// When the attachment was initiated.
	AttachTime *time.Time

	// A value that indicates whether the volume is deleted on instance termination.
	DeleteOnTermination *bool

	// The attachment state.
	Status AttachmentStatus

	// The ID of the Amazon EBS volume.
	VolumeId *string

	noSmithyDocumentSerde
}

// The summary of a task execution on a specified device.
type ExecutionSummary struct {

	// The ID of the execution.
	ExecutionId *string

	// The ID of the managed device that the task is being executed on.
	ManagedDeviceId *string

	// The state of the execution.
	State ExecutionState

	// The ID of the task.
	TaskId *string

	noSmithyDocumentSerde
}

// The description of an instance. Currently, Amazon EC2 instances are the only
// supported instance type.
type Instance struct {

	// The Amazon Machine Image (AMI) launch index, which you can use to find this
	// instance in the launch group.
	AmiLaunchIndex *int32

	// Any block device mapping entries for the instance.
	BlockDeviceMappings []InstanceBlockDeviceMapping

	// The CPU options for the instance.
	CpuOptions *CpuOptions

	// When the instance was created.
	CreatedAt *time.Time

	// The ID of the AMI used to launch the instance.
	ImageId *string

	// The ID of the instance.
	InstanceId *string

	// The instance type.
	InstanceType *string

	// The private IPv4 address assigned to the instance.
	PrivateIpAddress *string

	// The public IPv4 address assigned to the instance.
	PublicIpAddress *string

	// The device name of the root device volume (for example, /dev/sda1).
	RootDeviceName *string

	// The security groups for the instance.
	SecurityGroups []SecurityGroupIdentifier

	// The description of the current state of an instance.
	State *InstanceState

	// When the instance was last updated.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// The description of a block device mapping.
type InstanceBlockDeviceMapping struct {

	// The block device name.
	DeviceName *string

	// The parameters used to automatically set up Amazon Elastic Block Store (Amazon
	// EBS) volumes when the instance is launched.
	Ebs *EbsInstanceBlockDevice

	noSmithyDocumentSerde
}

// The description of the current state of an instance.
type InstanceState struct {

	// The state of the instance as a 16-bit unsigned integer. The high byte is all of
	// the bits between 2^8 and (2^16)-1, which equals decimal values between 256 and
	// 65,535. These numerical values are used for internal purposes and should be
	// ignored. The low byte is all of the bits between 2^0 and (2^8)-1, which equals
	// decimal values between 0 and 255. The valid values for the instance state code
	// are all in the range of the low byte. These values are:
	//
	// * 0 : pending
	//
	// * 16 :
	// running
	//
	// * 32 : shutting-down
	//
	// * 48 : terminated
	//
	// * 64 : stopping
	//
	// * 80 :
	// stopped
	//
	// You can ignore the high byte value by zeroing out all of the bits above
	// 2^8 or 256 in decimal.
	Code *int32

	// The current state of the instance.
	Name InstanceStateName

	noSmithyDocumentSerde
}

// The details about the instance.
type InstanceSummary struct {

	// A structure containing details about the instance.
	Instance *Instance

	// When the instance summary was last updated.
	LastUpdatedAt *time.Time

	noSmithyDocumentSerde
}

// The details about the physical network interface for the device.
type PhysicalNetworkInterface struct {

	// The default gateway of the device.
	DefaultGateway *string

	// The IP address of the device.
	IpAddress *string

	// A value that describes whether the IP address is dynamic or persistent.
	IpAddressAssignment IpAddressAssignment

	// The MAC address of the device.
	MacAddress *string

	// The netmask used to divide the IP address into subnets.
	Netmask *string

	// The physical connector type.
	PhysicalConnectorType PhysicalConnectorType

	// The physical network interface ID.
	PhysicalNetworkInterfaceId *string

	noSmithyDocumentSerde
}

// A structure used to reboot the device.
type Reboot struct {
	noSmithyDocumentSerde
}

// A summary of a resource available on the device.
type ResourceSummary struct {

	// The resource type.
	//
	// This member is required.
	ResourceType *string

	// The Amazon Resource Name (ARN) of the resource.
	Arn *string

	// The ID of the resource.
	Id *string

	noSmithyDocumentSerde
}

// Information about the device's security group.
type SecurityGroupIdentifier struct {

	// The security group ID.
	GroupId *string

	// The security group name.
	GroupName *string

	noSmithyDocumentSerde
}

// Information about the software on the device.
type SoftwareInformation struct {

	// The state of the software that is installed or that is being installed on the
	// device.
	InstallState *string

	// The version of the software currently installed on the device.
	InstalledVersion *string

	// The version of the software being installed on the device.
	InstallingVersion *string

	noSmithyDocumentSerde
}

// Information about the task assigned to one or many devices.
type TaskSummary struct {

	// The task ID.
	//
	// This member is required.
	TaskId *string

	// The state of the task assigned to one or many devices.
	State TaskState

	// Optional metadata that you assign to a resource. You can use tags to categorize
	// a resource in different ways, such as by purpose, owner, or environment.
	Tags map[string]string

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string

	noSmithyDocumentSerde
}

// A structure used to unlock a device.
type Unlock struct {
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isCommand() {}
