package protoeasy // import "go.pedge.io/protoeasy"

const (
	DefaultCompilerOptionsFile = ".protoeasy.yml"
	DefaultFileFormat          = FileFormatYML

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

type CompilerOptions struct {
	Includes             []string
	Dirs                 []string
	StripPackageComments bool
	Go                   *GoPluginOptions
	GrpcGateway          *GrpcGatewayPluginOptions
	Protolog             *ProtologPluginOptions
}

type CompilerOptionsProvider interface {
	Get() (*CompilerOptions, error)
}

type CompilerOptionsProviderOptions struct {
	File       string
	FileFormat FileFormat
}

func NewCompilerOptionsProvider(options CompilerOptionsProviderOptions) CompilerOptionsProvider {
	return newCompilerOptionsProvider(options)
}

type Compiler interface {
	Compile(dir string) error
}

func NewCompiler(options CompilerOptions) Compiler {
	return newCompiler(options)
}
