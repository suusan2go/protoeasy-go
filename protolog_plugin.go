package protoeasy

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
