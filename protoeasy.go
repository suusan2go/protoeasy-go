package protoeasy // import "go.pedge.io/protoeasy"

const (
	DefaultDirectivesFile = ".protoeasy.yml"
	DefaultFileFormat     = FileFormatYML

	FileFormatYML  FileFormat = 0
	FileFormatJSON FileFormat = 1
)

type FileFormat int

type PluginFlags struct {
	Includes []string
	Options  map[string]string
}

type Plugin interface {
	Name() string
	Flags(protoFiles []string) (*PluginFlags, error)
}

type GoPluginOptions struct {
	ImportPrefix       string
	ImportPath         string
	Plugins            []string
	NoDefaultModifiers bool
	Modifiers          []string
}

func NewGoPlugin(options GoPluginOptions) Plugin {
	return newGoPlugin(options)
}

type GrpcGatewayPluginOptions struct {
	NoDefaultModifiers bool
	Modifiers          []string
}

func NewGrpcGatewayPlugin(options GrpcGatewayPluginOptions) Plugin {
	return newGrpcGatewayPlugin(options)
}

type ProtologPluginOptions struct{}

func NewProtologPlugin(options ProtologPluginOptions) Plugin {
	return newProtologPlugin(options)
}

type Directives struct {
	Includes             []string
	StripPackageComments bool
	Go                   *GoPluginOptions
	GrpcGateway          *GrpcGatewayPluginOptions
	Protolog             *ProtologPluginOptions
}

type DirectivesProvider interface {
	Get(dir string) (*Directives, error)
}

type DirectivesProviderOptions struct {
	File       string
	FileFormat FileFormat
}

func NewDirectivesProvider(options DirectivesProviderOptions) DirectivesProvider {
	return newDirectivesProvider(options)
}

type ArgsProvider interface {
	Get(directives *Directives, protoFiles []string) ([]string, error)
}

type ArgsProviderOptions struct{}

func NewArgsProvider(options ArgsProviderOptions) ArgsProvider {
	return newArgsProvider(options)
}

type Compiler interface {
	Compile(directivesDir string, dirs []string) error
}

type CompilerOptions struct{}

func NewCompiler(
	directivesProvider DirectivesProvider,
	argsProvider ArgsProvider,
	options CompilerOptions,
) Compiler {
	return newCompiler(
		directivesProvider,
		argsProvider,
		options,
	)
}
