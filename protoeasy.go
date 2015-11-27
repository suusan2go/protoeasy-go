package protoeasy // import "go.pedge.io/protoeasy"

type ProtoSpec struct {
	DirPath           string
	RelDirPathToFiles map[string][]string
	OutDirPath        string
}

type ProtoSpecProvider interface {
	Get(dirPath string) (*ProtoSpec, error)
}

type ProtoSpecProviderOptions struct{}

func NewProtoSpecProvider(options ProtoSpecProviderOptions) ProtoSpecProvider {
	return newProtoSpecProvider(options)
}

type Plugin interface {
	Args(protoSpec *ProtoSpec) ([]string, error)
}

type GoPluginOptions struct {
	Grpc               bool
	GrpcGateway        bool
	Protolog           bool
	NoDefaultModifiers bool
}

func NewGoPlugin(goPath string, options GoPluginOptions) Plugin {
	return newGoPlugin(goPath, options)
}

type CppPluginOptions struct {
	Grpc bool
}

func NewCppPlugin(options CppPluginOptions) Plugin {
	return newCppPlugin(options)
}

type Compiler interface {
	Args(dirPath string) ([]string, error)
}

type CompilerOptions struct{}

func NewCompiler(
	protoSpecProvider ProtoSpecProvider,
	plugins []Plugin,
	options CompilerOptions,
) Compiler {
	return newCompiler(
		protoSpecProvider,
		plugins,
		options,
	)
}
