package connector

import (
	"context"

	"github.com/azaurus1/pulumi-kafka-connect/provider/pkg/kafkaconnect/config"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/ricardo-ch/go-kafka-connect/lib/connectors"
)

type Connector struct{}

type ConnectorArgs struct {
	Name   string         `pulumi:"name"`
	Config map[string]any `pulumi:"config"`
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
			Name: input.Name,
		},
		Config: input.Config,
	}

	resp, err := client.CreateConnector(request, false)
	if err != nil {
		return name, state, err
	}

	state.Result = resp.Message

	return name, state, nil
}
