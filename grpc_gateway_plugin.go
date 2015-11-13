package protoeasy

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
