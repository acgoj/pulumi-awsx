// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gen

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateVpc(awsSpec schema.PackageSpec) schema.PackageSpec {
	packageSpec := schema.PackageSpec{
		Resources: map[string]schema.ResourceSpec{
			"awsx:vpc:Vpc": vpcResource(awsSpec),
		},
		// If we have an input to the schema that is not a primitive type, we must define it below:
		Types: map[string]schema.ComplexTypeSpec{
			//"awsx:ecs:FargateServiceTaskDefinition": {
			//	ObjectTypeSpec: schema.ObjectTypeSpec{
			//		Type:        "object",
			//		Description: fargateTaskDefinitionResource.Description,
			//		Properties:  fargateTaskDefinitionResource.InputProperties,
			//	},
			//},
		},
	}

	return packageSpec
}

func vpcResource(awsSpec schema.PackageSpec) schema.ResourceSpec {
	awsVpcResource := awsSpec.Resources["aws:ec2/vpc:Vpc"]
	inputProperties := map[string]schema.PropertySpec{}
	for k, v := range awsVpcResource.InputProperties {
		inputProperties[k] = renamePropertyRefs(v, "#/types/aws:", awsRef("#/types/aws:"))
	}

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"vpc": {
					Description: "The VPC.",
					TypeSpec: schema.TypeSpec{
						Ref: awsRef("#/resources/aws:ec2%2fvpc:Vpc"),
					},
					Language: map[string]schema.RawMessage{
						"csharp": schema.RawMessage(`{
									"name": "AwsVpc"
								}`),
					},
				},
				"subnets": {
					Description: "The VPC's subnets.",
					TypeSpec: schema.TypeSpec{
						Type: "array",
						Items: &schema.TypeSpec{
							Ref: awsRef("#/resources/aws:ec2%2fsubnet:Subnet"),
						},
					},
				},
			},
		},
		InputProperties: inputProperties,
	}
}
