package protoeasy

type goPlugin struct {
	options GoPluginOptions
}

func newGoPlugin(options GoPluginOptions) *goPlugin {
	return &goPlugin{options}
}

func (p *goPlugin) Args(protoSpec *ProtoSpec, outDirPath string) ([]string, error) {
	return nil, nil
}
