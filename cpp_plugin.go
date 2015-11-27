package protoeasy

type cppPlugin struct {
	options CppPluginOptions
}

func newCppPlugin(options CppPluginOptions) *cppPlugin {
	return &cppPlugin{options}
}

func (p *cppPlugin) Args(protoSpec *ProtoSpec, outDirPath string) ([]string, error) {
	return nil, nil
}
