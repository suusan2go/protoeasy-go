package protoeasy

type protoSpecProvider struct {
	options ProtoSpecProviderOptions
}

func newProtoSpecProvider(options ProtoSpecProviderOptions) *protoSpecProvider {
	return &protoSpecProvider{options}
}

func (d *protoSpecProvider) Get(dirPath string) (*ProtoSpec, error) {
	return nil, nil
}
