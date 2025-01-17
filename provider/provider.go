// Copyright 2016-2023, Pulumi Corporation.
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

package provider

import (
	"github.com/azaurus1/pulumi-kafka-connect/provider/pkg/kafkaconnect/config"
	"github.com/azaurus1/pulumi-kafka-connect/provider/pkg/kafkaconnect/connector"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "kafkaconnect"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single resource and component
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[*connector.Connector, connector.ConnectorArgs, connector.ConnectorState](),
		},
		Components: []infer.InferredComponent{},
		Config:     infer.Config[*config.KafkaConnectConfig](),
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
		Metadata: schema.Metadata{
			DisplayName: "Kafka Connect",
			Description: "A Pulumi native provider for Kafka Connect",
			Keywords: []string{
				"pulumi",
				"kafka",
				"kafkaconnect",
				"category/utility",
				"kind/native",
			},
			Publisher:         "azaurus1",
			License:           "Apache-2.0",
			PluginDownloadURL: "github://api.github.com/azaurus1/pulumi-kafka-connect",
			Repository:        "https://github.com/azaurus1/pulumi-kafka-connect",
			LanguageMap: map[string]any{
				"go": map[string]any{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/azaurus1/pulumi-kafka-connect/sdk/go/pinecone",
				},
				"csharp": map[string]any{
					"packageReferences": map[string]string{
						"Pulumi": "3.*",
					},
					"rootNamespace": "Azaurus",
				},
				"nodejs": map[string]any{
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
					"packageName": "@azaurus/pulumi",
				},
				"python": map[string]any{
					"requires": map[string]string{
						"pulumi": ">=3.0.0,<4.0.0",
					},
					"packageName": "pulumi-kafkaconnect",
				},
			},
		},
	})
}
