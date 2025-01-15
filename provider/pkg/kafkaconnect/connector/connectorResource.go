package connector

import "context"

type Connector struct{}

type ConnectorArgs struct {
}

type ConnectorState struct {
	ConnectorArgs

	Result string `pulumi:"result"`
}

func (Connector) Create(ctx context.Context, name string, input ConnectorArgs, preview bool) (string, ConnectorState, error) {
	state := ConnectorState{}
	if preview {
		return name, state, nil
	}

	state.Result = "test"
	return name, state, nil
}
