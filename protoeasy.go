/*
Package protoeasy is intended to make using protoc simpler.

Protoeasy compiles all protocol buffer files in a directory/subdirectories,
taking care of all include directories, takes care of gRPC compilation,
and take care of package import modifiers for Golang.

Protoeasy also provides a client/server model where compilation is delegated
to a server process meant to be run in a Docker container. This allows you
to not have to install protoc, gRPC, and associated protoc plugins on your
local development machine.

See the README.md file for more details.
*/
package protoeasy // import "go.pedge.io/protoeasy"

var (
	//DefaultServerCompiler is the default Compiler for a server.
	DefaultServerCompiler = NewServerCompiler(ServerCompilerOptions{})
)

// ProtoSpec specifies the absolute directory path being used as a base
// for the current compilation, as well as the relative (to DirPath)
// directory path to all protocol buffer files in each directory within DirPath.
type ProtoSpec struct {
	DirPath           string
	RelDirPathToFiles map[string][]string
}

// Plugin is an individual flag provider for a specific language.
type Plugin interface {
	// Flags gets the protoc flags for the plugin.
	Flags(protoSpec *ProtoSpec, relDirPath string, outDirPath string) ([]string, error)
}

// PluginOptions are options for all Plugins.
type PluginOptions struct {
	RelOutDirPath string
}

// GrpcPluginOptions are options for Plugins that are associated with gRPC.
type GrpcPluginOptions struct {
	PluginOptions
	Grpc bool
}

// CppPluginOptions are options for the C++ Plugin.
type CppPluginOptions struct {
	GrpcPluginOptions
}

// CsharpPluginOptions are options for the C# Plugin.
type CsharpPluginOptions struct {
	GrpcPluginOptions
}

// ObjcPluginOptions are options for the Objective-C Plugin.
type ObjcPluginOptions struct {
	GrpcPluginOptions
}

// PythonPluginOptions are options for the Python plugin.
type PythonPluginOptions struct {
	GrpcPluginOptions
}

// RubyPluginOptions are options for the Ruby Plugin.
type RubyPluginOptions struct {
	GrpcPluginOptions
}

// GoPluginOptions are options for the Go Plugin.
type GoPluginOptions struct {
	GrpcPluginOptions
	ImportPath         string
	GrpcGateway        bool
	NoDefaultModifiers bool
}

// NewCppPlugin creates a new C++ Plugin.
func NewCppPlugin(options CppPluginOptions) Plugin {
	return newGrpcPlugin("cpp", "cpp", options.GrpcPluginOptions)
}

// NewCsharpPlugin creates a new C# Plugin.
func NewCsharpPlugin(options CsharpPluginOptions) Plugin {
	return newGrpcPlugin("csharp", "csharp", options.GrpcPluginOptions)
}

// NewObjcPlugin creates a new Objective-C Plugin.
func NewObjcPlugin(options ObjcPluginOptions) Plugin {
	return newGrpcPlugin("objc", "objective_c", options.GrpcPluginOptions)
}

// NewPythonPlugin creates a new Python Plugin.
func NewPythonPlugin(options PythonPluginOptions) Plugin {
	return newGrpcPlugin("python", "python", options.GrpcPluginOptions)
}

// NewRubyPlugin creates a new Ruby Plugin.
func NewRubyPlugin(options RubyPluginOptions) Plugin {
	return newGrpcPlugin("ruby", "ruby", options.GrpcPluginOptions)
}

// NewGoPlugin creates a new Go Plugin.
func NewGoPlugin(options GoPluginOptions) Plugin {
	return newGoPlugin(options)
}

// Compiler compiles protocol buffer files.
type Compiler interface {
	// Compile compiles the protocol buffer files in dirPath and outputs the generated
	// files to outDirPath, using the given Directives. It returns a list of the commands run.
	Compile(dirPath string, outDirPath string, directives *Directives) ([][]string, error)
}

// ServerCompilerOptions are options for a server Compiler.
type ServerCompilerOptions struct{}

// NewServerCompiler returns a new server Compiler.
func NewServerCompiler(options ServerCompilerOptions) Compiler {
	return newServerCompiler(options)
}

// NewAPIServer returns a new APIServer for the given Compiler.
func NewAPIServer(compiler Compiler) APIServer {
	return newAPIServer(compiler)
}

// ClientCompilerOptions are options for a client Compiler.
type ClientCompilerOptions struct{}

// NewClientCompiler returns a new client Compiler for the given APIClient.
func NewClientCompiler(
	apiClient APIClient,
	options ClientCompilerOptions,
) Compiler {
	return newClientCompiler(
		apiClient,
		options,
	)
}
