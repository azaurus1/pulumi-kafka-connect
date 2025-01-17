package connector

import (
	"context"
	"log"

	"github.com/azaurus1/pulumi-kafka-connect/provider/pkg/kafkaconnect/config"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/ricardo-ch/go-kafka-connect/lib/connectors"
)

type Connector struct{}

type ConnectorArgs struct {
	Config map[string]any `pulumi:"config,omitempty"`
}

type ConnectorState struct {
	ConnectorArgs
	Result string `pulumi:"result"`
}

func (*Connector) Create(ctx context.Context, name string, input ConnectorArgs, preview bool) (string, ConnectorState, error) {
	// get the config from the provider
	config := infer.GetConfig[config.KafkaConnectConfig](ctx)

	state := ConnectorState{}
	if preview {
		return name, state, nil
	}

	client := connectors.NewClient(config.Url)

	request := connectors.CreateConnectorRequest{
		ConnectorRequest: connectors.ConnectorRequest{
			Name: name,
		},
		Config: input.Config,
	}

	if config.User != "" && config.Password != "" {
		client.SetBasicAuth(config.User, config.Password)
	}

	resp, err := client.CreateConnector(request, false)
	if err != nil {
		return name, state, err
	}

	state.Result = resp.Message

	return name, state, nil
}

func (*Connector) Read(ctx context.Context, name string, state ConnectorState) error {
	// get the config from the provider
	config := infer.GetConfig[config.KafkaConnectConfig](ctx)

	client := connectors.NewClient(config.Url)

	request := connectors.ConnectorRequest{
		Name: name,
	}

	resp, err := client.GetConnector(request)
	if err != nil {
		return err
	}

	state.Config = resp.Config

	return nil

}

func (*Connector) Delete(ctx context.Context, name string, state ConnectorState) error {
	// get the config from the provider
	config := infer.GetConfig[config.KafkaConnectConfig](ctx)

	client := connectors.NewClient(config.Url)

	request := connectors.ConnectorRequest{
		Name: name,
	}

	if config.User != "" && config.Password != "" {
		client.SetBasicAuth(config.User, config.Password)
	}

	resp, err := client.DeleteConnector(request, false)
	if err != nil {
		return err
	}

	log.Println(resp)

	return nil
}
