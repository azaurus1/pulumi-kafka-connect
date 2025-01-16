package config

import (
	"context"
	"fmt"
	"os"

	goprovider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type KafkaConnectConfig struct {
	Url      string `pulumi:"url"`
	User     string `pulumi:"user,optional" provider:"secret"`
	Password string `pulumi:"password,optional" provider:"secret"`
}

func (c *KafkaConnectConfig) Annotate(a infer.Annotator) {
	a.Describe(&c.Url, "The url for the kafka connect cluster")
}

func (c *KafkaConnectConfig) Configure(ctx context.Context) error {
	goprovider.GetLogger(ctx).Debugf("Configuring Kafka Connect provider")
	if c.Url == "" {
		url, exists := os.LookupEnv("KAFKA_CONNECT_URL")
		if exists {
			c.Url = url
		}
		return fmt.Errorf("URL is required")
	}

	if c.User == "" {
		user, exists := os.LookupEnv("KAFKA_CONNECT_USER")
		if exists {
			c.User = user
		}
	}

	if c.Password == "" {
		password, exists := os.LookupEnv("KAFKA_CONNECT_PASSWORD")
		if exists {
			c.Password = password
		}
	}
	return nil
}
