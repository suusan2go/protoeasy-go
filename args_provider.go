package protoeasy

type argsProvider struct {
	options ArgsProviderOptions
}

func newArgsProvider(options ArgsProviderOptions) *argsProvider {
	return &argsProvider{options}
}

func (d *argsProvider) Get(directives *Directives, protoFiles []string) ([]string, error) {
	return nil, nil
}
