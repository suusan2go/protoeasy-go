package protoeasy

type compiler struct {
	protoSpecProvider ProtoSpecProvider
	options           CompilerOptions
}

func newCompiler(
	protoSpecProvider ProtoSpecProvider,
	options CompilerOptions,
) *compiler {
	return &compiler{
		protoSpecProvider,
		options,
	}
}

func (c *compiler) Args(dirPath string, outDirPath string, directives *CompilerDirectives) ([]string, error) {
	protoSpec, err := c.protoSpecProvider.Get(dirPath)
	if err != nil {
		return nil, err
	}
	plugins, err := c.getPlugins(directives)
	if err != nil {
		return nil, err
	}
	var args []string
	for _, plugin := range plugins {
		iArgs, err := plugin.Args(protoSpec, outDirPath)
		if err != nil {
			return nil, err
		}
		args = append(args, iArgs...)
	}
	return args, nil
}

func (c *compiler) getPlugins(directives *CompilerDirectives) ([]Plugin, error) {
	return nil, nil
}
