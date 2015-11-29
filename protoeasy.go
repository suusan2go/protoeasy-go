/*
Package protoeasy is intended to make using protoc simpler, if certain conventions are followed.
*/
package protoeasy // import "go.pedge.io/protoeasy"

var (
	DefaultServerCompiler = NewServerCompiler(ServerCompilerOptions{})
)

type ProtoSpec struct {
	DirPath           string
	RelDirPathToFiles map[string][]string
}

type Plugin interface {
	Flags(protoSpec *ProtoSpec, relDirPath string, outDirPath string) ([]string, error)
}

type PluginOptions struct {
	RelOutDirPath string
}

type GrpcPluginOptions struct {
	PluginOptions
	Grpc bool
}

type CppPluginOptions struct {
	GrpcPluginOptions
}

type CsharpPluginOptions struct {
	GrpcPluginOptions
}

type ObjectiveCPluginOptions struct {
	GrpcPluginOptions
}

type PythonPluginOptions struct {
	GrpcPluginOptions
}

type RubyPluginOptions struct {
	GrpcPluginOptions
}

type GoPluginOptions struct {
	GrpcPluginOptions
	ImportPath         string
	GrpcGateway        bool
	NoDefaultModifiers bool
}

func NewCppPlugin(options CppPluginOptions) Plugin {
	return newGrpcPlugin("cpp", "cpp", options.GrpcPluginOptions)
}

func NewCsharpPlugin(options CsharpPluginOptions) Plugin {
	return newGrpcPlugin("csharp", "csharp", options.GrpcPluginOptions)
}

func NewObjectiveCPlugin(options ObjectiveCPluginOptions) Plugin {
	return newGrpcPlugin("objc", "objective_c", options.GrpcPluginOptions)
}

func NewPythonPlugin(options PythonPluginOptions) Plugin {
	return newGrpcPlugin("python", "python", options.GrpcPluginOptions)
}

func NewRubyPlugin(options RubyPluginOptions) Plugin {
	return newGrpcPlugin("ruby", "ruby", options.GrpcPluginOptions)
}

func NewGoPlugin(options GoPluginOptions) Plugin {
	return newGoPlugin(options)
}

type Compiler interface {
	// Returns a list of the commands run
	Compile(dirPath string, outDirPath string, directives *Directives) ([][]string, error)
}

type ServerCompilerOptions struct{}

func NewServerCompiler(options ServerCompilerOptions) Compiler {
	return newServerCompiler(options)
}

func NewAPIServer(compiler Compiler) APIServer {
	return newAPIServer(compiler)
}

type ClientCompilerOptions struct{}

func NewClientCompiler(
	apiClient APIClient,
	options ClientCompilerOptions,
) Compiler {
	return newClientCompiler(
		apiClient,
		options,
	)
}
