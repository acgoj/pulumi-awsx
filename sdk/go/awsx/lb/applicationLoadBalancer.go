// *** WARNING: this file was generated by pulumi-gen-awsx. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationLoadBalancer struct {
	pulumi.ResourceState

	// Underlying Load Balancer resource
	LoadBalancer lb.LoadBalancerOutput `pulumi:"loadBalancer"`
}

// NewApplicationLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewApplicationLoadBalancer(ctx *pulumi.Context,
	name string, args *ApplicationLoadBalancerArgs, opts ...pulumi.ResourceOption) (*ApplicationLoadBalancer, error) {
	if args == nil {
		args = &ApplicationLoadBalancerArgs{}
	}

	var resource ApplicationLoadBalancer
	err := ctx.RegisterRemoteComponentResource("awsx:lb:ApplicationLoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type applicationLoadBalancerArgs struct {
	// An Access Logs block. Access Logs documented below.
	AccessLogs *lb.LoadBalancerAccessLogs `pulumi:"accessLogs"`
	// The ID of the customer owned ipv4 pool to use for this load balancer.
	CustomerOwnedIpv4Pool *string `pulumi:"customerOwnedIpv4Pool"`
	// Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are `monitor`, `defensive` (default), `strictest`.
	DesyncMitigationMode *string `pulumi:"desyncMitigationMode"`
	// Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
	DropInvalidHeaderFields *bool `pulumi:"dropInvalidHeaderFields"`
	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a `network` load balancer feature. Defaults to `false`.
	EnableCrossZoneLoadBalancing *bool `pulumi:"enableCrossZoneLoadBalancing"`
	// If true, deletion of the load balancer will be disabled via
	// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection *bool `pulumi:"enableDeletionProtection"`
	// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 *bool `pulumi:"enableHttp2"`
	// Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. Defaults to `false`.
	EnableWafFailOpen *bool `pulumi:"enableWafFailOpen"`
	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout *int `pulumi:"idleTimeout"`
	// If true, the LB will be internal.
	Internal *bool `pulumi:"internal"`
	// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
	IpAddressType *string `pulumi:"ipAddressType"`
	// The type of load balancer to create. Possible values are `application`, `gateway`, or `network`. The default value is `application`.
	LoadBalancerType *string `pulumi:"loadBalancerType"`
	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
	// this provider will autogenerate a name beginning with `tf-lb`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
	SecurityGroups []string `pulumi:"securityGroups"`
	// A subnet mapping block as documented below.
	SubnetMappings []lb.LoadBalancerSubnetMapping `pulumi:"subnetMappings"`
	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type `network`. Changing this value
	// for load balancers of type `network` will force a recreation of the resource.
	Subnets []string `pulumi:"subnets"`
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ApplicationLoadBalancer resource.
type ApplicationLoadBalancerArgs struct {
	// An Access Logs block. Access Logs documented below.
	AccessLogs lb.LoadBalancerAccessLogsPtrInput
	// The ID of the customer owned ipv4 pool to use for this load balancer.
	CustomerOwnedIpv4Pool pulumi.StringPtrInput
	// Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are `monitor`, `defensive` (default), `strictest`.
	DesyncMitigationMode pulumi.StringPtrInput
	// Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
	DropInvalidHeaderFields pulumi.BoolPtrInput
	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a `network` load balancer feature. Defaults to `false`.
	EnableCrossZoneLoadBalancing pulumi.BoolPtrInput
	// If true, deletion of the load balancer will be disabled via
	// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection pulumi.BoolPtrInput
	// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 pulumi.BoolPtrInput
	// Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. Defaults to `false`.
	EnableWafFailOpen pulumi.BoolPtrInput
	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout pulumi.IntPtrInput
	// If true, the LB will be internal.
	Internal pulumi.BoolPtrInput
	// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
	IpAddressType pulumi.StringPtrInput
	// The type of load balancer to create. Possible values are `application`, `gateway`, or `network`. The default value is `application`.
	LoadBalancerType pulumi.StringPtrInput
	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
	// this provider will autogenerate a name beginning with `tf-lb`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
	SecurityGroups pulumi.StringArrayInput
	// A subnet mapping block as documented below.
	SubnetMappings lb.LoadBalancerSubnetMappingArrayInput
	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type `network`. Changing this value
	// for load balancers of type `network` will force a recreation of the resource.
	Subnets pulumi.StringArrayInput
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ApplicationLoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationLoadBalancerArgs)(nil)).Elem()
}

type ApplicationLoadBalancerInput interface {
	pulumi.Input

	ToApplicationLoadBalancerOutput() ApplicationLoadBalancerOutput
	ToApplicationLoadBalancerOutputWithContext(ctx context.Context) ApplicationLoadBalancerOutput
}

func (*ApplicationLoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLoadBalancer)(nil)).Elem()
}

func (i *ApplicationLoadBalancer) ToApplicationLoadBalancerOutput() ApplicationLoadBalancerOutput {
	return i.ToApplicationLoadBalancerOutputWithContext(context.Background())
}

func (i *ApplicationLoadBalancer) ToApplicationLoadBalancerOutputWithContext(ctx context.Context) ApplicationLoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLoadBalancerOutput)
}

// ApplicationLoadBalancerArrayInput is an input type that accepts ApplicationLoadBalancerArray and ApplicationLoadBalancerArrayOutput values.
// You can construct a concrete instance of `ApplicationLoadBalancerArrayInput` via:
//
//          ApplicationLoadBalancerArray{ ApplicationLoadBalancerArgs{...} }
type ApplicationLoadBalancerArrayInput interface {
	pulumi.Input

	ToApplicationLoadBalancerArrayOutput() ApplicationLoadBalancerArrayOutput
	ToApplicationLoadBalancerArrayOutputWithContext(context.Context) ApplicationLoadBalancerArrayOutput
}

type ApplicationLoadBalancerArray []ApplicationLoadBalancerInput

func (ApplicationLoadBalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationLoadBalancer)(nil)).Elem()
}

func (i ApplicationLoadBalancerArray) ToApplicationLoadBalancerArrayOutput() ApplicationLoadBalancerArrayOutput {
	return i.ToApplicationLoadBalancerArrayOutputWithContext(context.Background())
}

func (i ApplicationLoadBalancerArray) ToApplicationLoadBalancerArrayOutputWithContext(ctx context.Context) ApplicationLoadBalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLoadBalancerArrayOutput)
}

// ApplicationLoadBalancerMapInput is an input type that accepts ApplicationLoadBalancerMap and ApplicationLoadBalancerMapOutput values.
// You can construct a concrete instance of `ApplicationLoadBalancerMapInput` via:
//
//          ApplicationLoadBalancerMap{ "key": ApplicationLoadBalancerArgs{...} }
type ApplicationLoadBalancerMapInput interface {
	pulumi.Input

	ToApplicationLoadBalancerMapOutput() ApplicationLoadBalancerMapOutput
	ToApplicationLoadBalancerMapOutputWithContext(context.Context) ApplicationLoadBalancerMapOutput
}

type ApplicationLoadBalancerMap map[string]ApplicationLoadBalancerInput

func (ApplicationLoadBalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationLoadBalancer)(nil)).Elem()
}

func (i ApplicationLoadBalancerMap) ToApplicationLoadBalancerMapOutput() ApplicationLoadBalancerMapOutput {
	return i.ToApplicationLoadBalancerMapOutputWithContext(context.Background())
}

func (i ApplicationLoadBalancerMap) ToApplicationLoadBalancerMapOutputWithContext(ctx context.Context) ApplicationLoadBalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLoadBalancerMapOutput)
}

type ApplicationLoadBalancerOutput struct{ *pulumi.OutputState }

func (ApplicationLoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLoadBalancer)(nil)).Elem()
}

func (o ApplicationLoadBalancerOutput) ToApplicationLoadBalancerOutput() ApplicationLoadBalancerOutput {
	return o
}

func (o ApplicationLoadBalancerOutput) ToApplicationLoadBalancerOutputWithContext(ctx context.Context) ApplicationLoadBalancerOutput {
	return o
}

type ApplicationLoadBalancerArrayOutput struct{ *pulumi.OutputState }

func (ApplicationLoadBalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationLoadBalancer)(nil)).Elem()
}

func (o ApplicationLoadBalancerArrayOutput) ToApplicationLoadBalancerArrayOutput() ApplicationLoadBalancerArrayOutput {
	return o
}

func (o ApplicationLoadBalancerArrayOutput) ToApplicationLoadBalancerArrayOutputWithContext(ctx context.Context) ApplicationLoadBalancerArrayOutput {
	return o
}

func (o ApplicationLoadBalancerArrayOutput) Index(i pulumi.IntInput) ApplicationLoadBalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationLoadBalancer {
		return vs[0].([]*ApplicationLoadBalancer)[vs[1].(int)]
	}).(ApplicationLoadBalancerOutput)
}

type ApplicationLoadBalancerMapOutput struct{ *pulumi.OutputState }

func (ApplicationLoadBalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationLoadBalancer)(nil)).Elem()
}

func (o ApplicationLoadBalancerMapOutput) ToApplicationLoadBalancerMapOutput() ApplicationLoadBalancerMapOutput {
	return o
}

func (o ApplicationLoadBalancerMapOutput) ToApplicationLoadBalancerMapOutputWithContext(ctx context.Context) ApplicationLoadBalancerMapOutput {
	return o
}

func (o ApplicationLoadBalancerMapOutput) MapIndex(k pulumi.StringInput) ApplicationLoadBalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationLoadBalancer {
		return vs[0].(map[string]*ApplicationLoadBalancer)[vs[1].(string)]
	}).(ApplicationLoadBalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationLoadBalancerInput)(nil)).Elem(), &ApplicationLoadBalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationLoadBalancerArrayInput)(nil)).Elem(), ApplicationLoadBalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationLoadBalancerMapInput)(nil)).Elem(), ApplicationLoadBalancerMap{})
	pulumi.RegisterOutputType(ApplicationLoadBalancerOutput{})
	pulumi.RegisterOutputType(ApplicationLoadBalancerArrayOutput{})
	pulumi.RegisterOutputType(ApplicationLoadBalancerMapOutput{})
}
