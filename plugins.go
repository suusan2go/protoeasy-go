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
