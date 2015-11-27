package protoeasy

type cppPlugin struct {
	options CppPluginOptions
}

func newCppPlugin(options CppPluginOptions) *cppPlugin {
	return &cppPlugin{options}
}

func (p *cppPlugin) Args(protoSpec *ProtoSpec) ([]string, error) {
	return nil, nil
}
