package protoeasy // import "go.pedge.io/protoeasy"

type ProtoSpec struct {
	DirPath           string
	RelDirPathToFiles map[string][]string
}

type ProtoSpecProvider interface {
	Get(dirPath string, excludeFilePatterns []string) (*ProtoSpec, error)
}

type ProtoSpecProviderOptions struct{}

func NewProtoSpecProvider(options ProtoSpecProviderOptions) ProtoSpecProvider {
	return newProtoSpecProvider(options)
}

type Plugin interface {
	Args(protoSpec *ProtoSpec, relDirPath string, outDirPath string) ([]string, error)
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
	Protolog           bool
	NoDefaultModifiers bool
}

func NewCppPlugin(options CppPluginOptions) Plugin {
	return newGrpcPlugin("cpp", options.GrpcPluginOptions)
}

func NewCsharpPlugin(options CsharpPluginOptions) Plugin {
	return newGrpcPlugin("csharp", options.GrpcPluginOptions)
}

func NewObjectiveCPlugin(options ObjectiveCPluginOptions) Plugin {
	return newGrpcPlugin("objectivec", options.GrpcPluginOptions)
}

func NewPythonPlugin(options PythonPluginOptions) Plugin {
	return newGrpcPlugin("python", options.GrpcPluginOptions)
}

func NewRubyPlugin(options RubyPluginOptions) Plugin {
	return newGrpcPlugin("ruby", options.GrpcPluginOptions)
}

func NewGoPlugin(options GoPluginOptions) Plugin {
	return newGoPlugin(options)
}

type CompilerDirectives struct {
	ExcludeFilePatterns     []string
	ProtoPaths              []string
	OutDirPath              string
	Cpp                     bool
	CppRelOutDirPath        string
	Csharp                  bool
	CsharpRelOutDirPath     string
	ObjectiveC              bool
	ObjectiveCRelOutDirPath string
	Python                  bool
	PythonRelOutDirPath     string
	Ruby                    bool
	RubyRelOutDirPath       string
	Go                      bool
	GoRelOutDirPath         string
	GoImportPath            string
	Grpc                    bool
	GrpcGateway             bool
	Protolog                bool
}

type Compiler interface {
	ArgsList(dirPath string, directives *CompilerDirectives) ([][]string, error)
}

type CompilerOptions struct{}

func NewCompiler(
	protoSpecProvider ProtoSpecProvider,
	options CompilerOptions,
) Compiler {
	return newCompiler(
		protoSpecProvider,
		options,
	)
}
