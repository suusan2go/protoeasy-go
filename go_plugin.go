package protoeasy

type goPlugin struct {
	goPath  string
	options GoPluginOptions
}

func newGoPlugin(goPath string, options GoPluginOptions) *goPlugin {
	return &goPlugin{goPath, options}
}

func (p *goPlugin) Name() string {
	return "go"
}

func (p *goPlugin) Flags(protoFiles []string) (*PluginFlags, error) {
	return nil, nil
}
