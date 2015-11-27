package protoeasy

type goPlugin struct {
	goPath  string
	options GoPluginOptions
}

func newGoPlugin(goPath string, options GoPluginOptions) *goPlugin {
	return &goPlugin{goPath, options}
}

func (p *goPlugin) Args(protoSpec *ProtoSpec) ([]string, error) {
	return nil, nil
}
