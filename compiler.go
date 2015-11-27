package protoeasy

type compiler struct {
	protoSpecProvider ProtoSpecProvider
	plugins           []Plugin
	options           CompilerOptions
}

func newCompiler(
	protoSpecProvider ProtoSpecProvider,
	plugins []Plugin,
	options CompilerOptions,
) *compiler {
	return &compiler{
		protoSpecProvider,
		plugins,
		options,
	}
}

func (d *compiler) Args(dirPath string) ([]string, error) {
	return nil, nil
}
