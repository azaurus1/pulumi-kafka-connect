package main

import (
	"example.com/pulumi-kafkaconnect/sdk/go/kafkaconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := kafkaconnect.NewRandom(ctx, "myRandomResource", &kafkaconnect.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		_, err = kafkaconnect.NewRandomComponent(ctx, "myRandomComponent", &kafkaconnect.RandomComponentArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", pulumi.StringMap{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
