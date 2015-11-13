package protoeasy

type directivesProvider struct {
	goPath  string
	options DirectivesProviderOptions
}

func newDirectivesProvider(goPath string, options DirectivesProviderOptions) *directivesProvider {
	return &directivesProvider{goPath, options}
}

func (d *directivesProvider) Get(dir string) (*Directives, error) {
	return nil, nil
}
