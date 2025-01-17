// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/azaurus1/pulumi-kafka-connect/sdk/go/kafkaconnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafkaconnect:password")
}

// The url for the kafka connect cluster
func GetUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafkaconnect:url")
}
func GetUser(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafkaconnect:user")
}
