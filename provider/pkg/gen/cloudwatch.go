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

func generateCloudwatch(awsSpec schema.PackageSpec) schema.PackageSpec {
	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{
			"awsx:awsx:DefaultLogGroup": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: "Log group with default setup unless explicitly skipped.",
					Properties: map[string]schema.PropertySpec{
						"skip": {
							Description: "Skip creation of the log group.",
							TypeSpec: schema.TypeSpec{
								Type:  "boolean",
								Plain: true,
							},
						},
						"existing": {
							Description: "Identity of an existing log group to use. Cannot be used in combination with `args` or `opts`.",
							TypeSpec: schema.TypeSpec{
								Ref:   "#/types/awsx:awsx:ExistingLogGroup",
								Plain: true,
							},
						},
						"args": {
							Description: "Arguments to use instead of the default values during creation.",
							TypeSpec: schema.TypeSpec{
								Ref:   "#/types/awsx:awsx:LogGroup",
								Plain: true,
							},
						},
					},
				},
			},
			"awsx:awsx:OptionalLogGroup": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: "Log group which is only created if enabled.",
					Properties: map[string]schema.PropertySpec{
						"enable": {
							Description: "Enable creation of the log group.",
							TypeSpec: schema.TypeSpec{
								Type:  "boolean",
								Plain: true,
							},
						},
						"existing": {
							Description: "Identity of an existing log group to use. Cannot be used in combination with `args` or `opts`.",
							TypeSpec: schema.TypeSpec{
								Ref:   "#/types/awsx:awsx:ExistingLogGroup",
								Plain: true,
							},
						},
						"args": {
							Description: "Arguments to use instead of the default values during creation.",
							TypeSpec: schema.TypeSpec{
								Ref:   "#/types/awsx:awsx:LogGroup",
								Plain: true,
							},
						},
					},
				},
			},
			"awsx:awsx:RequiredLogGroup": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: "Log group with default setup.",
					Properties: map[string]schema.PropertySpec{
						"existing": {
							Description: "Identity of an existing log group to use. Cannot be used in combination with `args` or `opts`.",
							TypeSpec: schema.TypeSpec{
								Ref:   "#/types/awsx:awsx:ExistingLogGroup",
								Plain: true,
							},
						},
						"args": {
							Description: "Arguments to use instead of the default values during creation.",
							TypeSpec: schema.TypeSpec{
								Ref:   "#/types/awsx:awsx:LogGroup",
								Plain: true,
							},
						},
					},
				},
			},
			"awsx:awsx:ExistingLogGroup": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: "Reference to an existing log group.",
					Properties: map[string]schema.PropertySpec{
						"arn": {
							Description: "Arn of the log group. Only one of [arn] or [name] can be specified.",
							TypeSpec: schema.TypeSpec{
								Type: "string",
							},
						},
						"name": {
							Description: "Name of the log group. Only one of [arn] or [name] can be specified.",
							TypeSpec: schema.TypeSpec{
								Type: "string",
							},
						},
						"region": {
							Description: "Region of the log group. If not specified, the provider region will be used.",
							TypeSpec: schema.TypeSpec{
								Type: "string",
							},
						},
					},
				},
			},
			"awsx:awsx:LogGroup": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: "The set of arguments for constructing a LogGroup resource.",
					Properties:  awsSpec.Resources["aws:cloudwatch/logGroup:LogGroup"].InputProperties,
				},
			},
		},
	}
}
