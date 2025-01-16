package main

import (
	"example.com/pulumi-kafkaconnect/sdk/go/kafkaconnect"
	"example.com/pulumi-kafkaconnect/sdk/go/kafkaconnect/connector"
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
		myConnector, err := connector.NewConnector(ctx, "myConnector", nil, pulumi.Provider(defaultProvider))
		if err != nil {
			return err
		}
		ctx.Export("output", pulumi.StringMap{
			"value": myConnector.Result,
		})
		return nil
	})
}
