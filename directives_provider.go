package protoeasy

type directivesProvider struct {
	options DirectivesProviderOptions
}

func newDirectivesProvider(options DirectivesProviderOptions) *directivesProvider {
	return &directivesProvider{options}
}

func (d *directivesProvider) Get(dir string) (*Directives, error) {
	return nil, nil
}
