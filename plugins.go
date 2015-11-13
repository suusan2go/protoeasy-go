package protoeasy

type goPlugin struct {
	options GoPluginOptions
}

func newGoPlugin(options GoPluginOptions) *goPlugin {
	return &goPlugin{options}
}

func (p *goPlugin) Name() string {
	return "go"
}

func (p *goPlugin) Flags(protoFiles []string) (*PluginFlags, error) {
	return nil, nil
}

type grpcGatewayPlugin struct {
	options GrpcGatewayPluginOptions
}

func newGrpcGatewayPlugin(options GrpcGatewayPluginOptions) *grpcGatewayPlugin {
	return &grpcGatewayPlugin{options}
}

func (p *grpcGatewayPlugin) Name() string {
	return "grpc-gateway"
}

func (p *grpcGatewayPlugin) Flags(protoFiles []string) (*PluginFlags, error) {
	return nil, nil
}

type protologPlugin struct {
	options ProtologPluginOptions
}

func newProtologPlugin(options ProtologPluginOptions) *protologPlugin {
	return &protologPlugin{options}
}

func (p *protologPlugin) Name() string {
	return "protolog"
}

func (p *protologPlugin) Flags(protoFiles []string) (*PluginFlags, error) {
	return nil, nil
}
