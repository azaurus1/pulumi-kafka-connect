package main

import (
	"github.com/azaurus1/pulumi-kafka-connect/sdk/go/kafkaconnect"
	"github.com/azaurus1/pulumi-kafka-connect/sdk/go/kafkaconnect/connector"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultProvider, err := kafkaconnect.NewProvider(ctx, "defaultProvider", &kafkaconnect.ProviderArgs{
			Url: pulumi.String("http://localhost:8083"),
		})
		if err != nil {
			return err
		}
		mirrorHeartbeatConnectorYxdw, err := connector.NewConnector(ctx, "mirror-heartbeat-connector-yxdw", &connector.ConnectorArgs{
			Config: pulumi.Map{
				"connector.class":                     pulumi.Any("org.apache.kafka.connect.mirror.MirrorHeartbeatConnector"),
				"source.cluster.alias":                pulumi.Any("source"),
				"heartbeats.topic.replication.factor": pulumi.Any("-1"),
				"name":                                pulumi.Any("mirror-heartbeat-connector-yxdw"),
			},
		}, pulumi.Provider(defaultProvider))
		if err != nil {
			return err
		}
		ctx.Export("output", pulumi.StringMap{
			"value": mirrorHeartbeatConnectorYxdw.Result,
		})
		return nil
	})
}
