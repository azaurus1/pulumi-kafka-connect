package connector

import (
	"context"

	"github.com/azaurus1/pulumi-kafka-connect/provider/pkg/kafkaconnect/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/ricardo-ch/go-kafka-connect/lib/connectors"
)

type Connector struct{}

type ConnectorArgs struct {
	Config map[string]any `pulumi:"config,omitempty"`
}

type ConnectorState struct {
	Config map[string]any `pulumi:"config"`
	Result string         `pulumi:"result"`
}

func getClient(ctx context.Context) connectors.HighLevelClient {
	config := infer.GetConfig[config.KafkaConnectConfig](ctx)

	client := connectors.NewClient(config.Url)

	if config.User != "" && config.Password != "" {
		client.SetBasicAuth(config.User, config.Password)
	}

	return client
}

func (*Connector) Create(ctx context.Context, req infer.CreateRequest[ConnectorArgs]) (infer.CreateResponse[ConnectorState], error) {
	client := getClient(ctx)

	resp, err := client.CreateConnector(connectors.CreateConnectorRequest{
		ConnectorRequest: connectors.ConnectorRequest{
			Name: req.Name,
		},
		Config: req.Inputs.Config,
	}, false)
	if err != nil {
		return infer.CreateResponse[ConnectorState]{}, err
	}
	p.GetLogger(ctx).Infof("Create connector %s, got config %v", req.Name, resp.Config)

	return infer.CreateResponse[ConnectorState]{
		ID: req.Name,
		Output: ConnectorState{
			Config: resp.Config,
			Result: resp.Message,
		},
	}, nil
}

func (c *Connector) Update(ctx context.Context, req infer.UpdateRequest[ConnectorArgs, ConnectorState]) (infer.UpdateResponse[ConnectorState], error) {
	client := getClient(ctx)
	resp, err := client.UpdateConnector(connectors.CreateConnectorRequest{
		ConnectorRequest: connectors.ConnectorRequest{
			Name: req.ID,
		},
		Config: req.Inputs.Config,
	}, false)
	if err != nil {
		return infer.UpdateResponse[ConnectorState]{}, err
	}

	return infer.UpdateResponse[ConnectorState]{
		Output: ConnectorState{
			Config: resp.Config,
			Result: resp.Message,
		},
	}, nil
}

func (c *Connector) Read(ctx context.Context, req infer.ReadRequest[ConnectorArgs, ConnectorState]) (infer.ReadResponse[ConnectorArgs, ConnectorState], error) {
	client := getClient(ctx)

	resp, err := client.GetConnector(connectors.ConnectorRequest{Name: req.ID})
	if err != nil {
		return infer.ReadResponse[ConnectorArgs, ConnectorState]{}, err
	}
	p.GetLogger(ctx).Infof("Read connector %s, got config %v", req.ID, resp.Config)

	return infer.ReadResponse[ConnectorArgs, ConnectorState]{
		ID: req.ID,
		Inputs: ConnectorArgs{
			Config: resp.Config,
		},
		State: ConnectorState{
			Config: resp.Config,
			Result: resp.Message,
		},
	}, nil
}

func (*Connector) Delete(ctx context.Context, req infer.DeleteRequest[ConnectorState]) (infer.DeleteResponse, error) {
	client := getClient(ctx)
	_, err := client.DeleteConnector(connectors.ConnectorRequest{Name: req.ID}, false)
	return infer.DeleteResponse{}, err
}

func (c *Connector) WireDependencies(f infer.FieldSelector, args *ConnectorArgs, state *ConnectorState) {
	f.OutputField(&state.Config).DependsOn(f.InputField(&args.Config))
}
