package protoeasy

type compiler struct {
	directivesProvider DirectivesProvider
	argsProvider       ArgsProvider
	options            CompilerOptions
}

func newCompiler(
	directivesProvider DirectivesProvider,
	argsProvider ArgsProvider,
	options CompilerOptions,
) *compiler {
	return &compiler{
		directivesProvider,
		argsProvider,
		options,
	}
}

func (d *compiler) Compile(directivesDir string, dirs []string) error {
	return nil
}
