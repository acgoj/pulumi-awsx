// *** WARNING: this file was generated by pulumi-gen-awsx. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vpc struct {
	pulumi.ResourceState

	// The VPC's subnets.
	Subnets ec2.SubnetArrayOutput `pulumi:"subnets"`
	// The VPC.
	Vpc ec2.VpcOutput `pulumi:"vpc"`
}

// NewVpc registers a new resource with the given unique name, arguments, and options.
func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil {
		args = &VpcArgs{}
	}

	var resource Vpc
	err := ctx.RegisterRemoteComponentResource("awsx:vpc:Vpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type vpcArgs struct {
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6_ipam_pool_id`
	AssignGeneratedIpv6CidrBlock *bool `pulumi:"assignGeneratedIpv6CidrBlock"`
	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4_netmask_length`.
	CidrBlock *string `pulumi:"cidrBlock"`
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink *bool `pulumi:"enableClassiclink"`
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclinkDnsSupport *bool `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames *bool `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport *bool `pulumi:"enableDnsSupport"`
	// A tenancy option for instances launched into the VPC. Default is `default`, which makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
	InstanceTenancy *string `pulumi:"instanceTenancy"`
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	Ipv4IpamPoolId *string `pulumi:"ipv4IpamPoolId"`
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4_ipam_pool_id`.
	Ipv4NetmaskLength *int `pulumi:"ipv4NetmaskLength"`
	// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6_netmask_length`.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// By default when an IPv6 CIDR is assigned to a VPC a default ipv6_cidr_block_network_border_group will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
	Ipv6CidrBlockNetworkBorderGroup *string `pulumi:"ipv6CidrBlockNetworkBorderGroup"`
	// IPAM Pool ID for a IPv6 pool. Conflicts with `assign_generated_ipv6_cidr_block`.
	Ipv6IpamPoolId *string `pulumi:"ipv6IpamPoolId"`
	// Netmask length to request from IPAM Pool. Conflicts with `ipv6_cidr_block`. This can be omitted if IPAM pool as a `allocation_default_netmask_length` set. Valid values: `56`.
	Ipv6NetmaskLength *int `pulumi:"ipv6NetmaskLength"`
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Vpc resource.
type VpcArgs struct {
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6_ipam_pool_id`
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrInput
	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4_netmask_length`.
	CidrBlock pulumi.StringPtrInput
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink pulumi.BoolPtrInput
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclinkDnsSupport pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport pulumi.BoolPtrInput
	// A tenancy option for instances launched into the VPC. Default is `default`, which makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
	InstanceTenancy pulumi.StringPtrInput
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	Ipv4IpamPoolId pulumi.StringPtrInput
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4_ipam_pool_id`.
	Ipv4NetmaskLength pulumi.IntPtrInput
	// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6_netmask_length`.
	Ipv6CidrBlock pulumi.StringPtrInput
	// By default when an IPv6 CIDR is assigned to a VPC a default ipv6_cidr_block_network_border_group will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
	Ipv6CidrBlockNetworkBorderGroup pulumi.StringPtrInput
	// IPAM Pool ID for a IPv6 pool. Conflicts with `assign_generated_ipv6_cidr_block`.
	Ipv6IpamPoolId pulumi.StringPtrInput
	// Netmask length to request from IPAM Pool. Conflicts with `ipv6_cidr_block`. This can be omitted if IPAM pool as a `allocation_default_netmask_length` set. Valid values: `56`.
	Ipv6NetmaskLength pulumi.IntPtrInput
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (VpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}

type VpcInput interface {
	pulumi.Input

	ToVpcOutput() VpcOutput
	ToVpcOutputWithContext(ctx context.Context) VpcOutput
}

func (*Vpc) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil)).Elem()
}

func (i *Vpc) ToVpcOutput() VpcOutput {
	return i.ToVpcOutputWithContext(context.Background())
}

func (i *Vpc) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcOutput)
}

// VpcArrayInput is an input type that accepts VpcArray and VpcArrayOutput values.
// You can construct a concrete instance of `VpcArrayInput` via:
//
//          VpcArray{ VpcArgs{...} }
type VpcArrayInput interface {
	pulumi.Input

	ToVpcArrayOutput() VpcArrayOutput
	ToVpcArrayOutputWithContext(context.Context) VpcArrayOutput
}

type VpcArray []VpcInput

func (VpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (i VpcArray) ToVpcArrayOutput() VpcArrayOutput {
	return i.ToVpcArrayOutputWithContext(context.Background())
}

func (i VpcArray) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcArrayOutput)
}

// VpcMapInput is an input type that accepts VpcMap and VpcMapOutput values.
// You can construct a concrete instance of `VpcMapInput` via:
//
//          VpcMap{ "key": VpcArgs{...} }
type VpcMapInput interface {
	pulumi.Input

	ToVpcMapOutput() VpcMapOutput
	ToVpcMapOutputWithContext(context.Context) VpcMapOutput
}

type VpcMap map[string]VpcInput

func (VpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (i VpcMap) ToVpcMapOutput() VpcMapOutput {
	return i.ToVpcMapOutputWithContext(context.Background())
}

func (i VpcMap) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcMapOutput)
}

type VpcOutput struct{ *pulumi.OutputState }

func (VpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil)).Elem()
}

func (o VpcOutput) ToVpcOutput() VpcOutput {
	return o
}

func (o VpcOutput) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return o
}

type VpcArrayOutput struct{ *pulumi.OutputState }

func (VpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (o VpcArrayOutput) ToVpcArrayOutput() VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) Index(i pulumi.IntInput) VpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vpc {
		return vs[0].([]*Vpc)[vs[1].(int)]
	}).(VpcOutput)
}

type VpcMapOutput struct{ *pulumi.OutputState }

func (VpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (o VpcMapOutput) ToVpcMapOutput() VpcMapOutput {
	return o
}

func (o VpcMapOutput) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return o
}

func (o VpcMapOutput) MapIndex(k pulumi.StringInput) VpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vpc {
		return vs[0].(map[string]*Vpc)[vs[1].(string)]
	}).(VpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcInput)(nil)).Elem(), &Vpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcArrayInput)(nil)).Elem(), VpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcMapInput)(nil)).Elem(), VpcMap{})
	pulumi.RegisterOutputType(VpcOutput{})
	pulumi.RegisterOutputType(VpcArrayOutput{})
	pulumi.RegisterOutputType(VpcMapOutput{})
}
