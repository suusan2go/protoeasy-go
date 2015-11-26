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
	ImportPath           string
	Grpc                 bool
	GrpcGateway          bool
	Protolog             bool
	Modifiers            []string
	StripPackageComments bool
}

func NewGoPlugin(goPath string, options GoPluginOptions) Plugin {
	return newGoPlugin(goPath, options)
}

type Directives struct {
	Includes []string
	Go       *GoPluginOptions
}

type DirectivesProvider interface {
	Get(dir string) (*Directives, error)
}

type DirectivesProviderOptions struct {
	File       string
	FileFormat FileFormat
}

func NewDirectivesProvider(goPath string, options DirectivesProviderOptions) DirectivesProvider {
	return newDirectivesProvider(goPath, options)
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
