package protoeasy // import "go.pedge.io/protoeasy"

type ProtoSpec struct {
	DirPath           string
	RelDirPathToFiles map[string][]string
}

type ProtoSpecProvider interface {
	Get(dirPath string) (*ProtoSpec, error)
}

type ProtoSpecProviderOptions struct{}

func NewProtoSpecProvider(options ProtoSpecProviderOptions) ProtoSpecProvider {
	return newProtoSpecProvider(options)
}

type Plugin interface {
	Args(protoSpec *ProtoSpec, outDirPath string) ([]string, error)
}

type GoPluginOptions struct {
	Grpc               bool
	GrpcGateway        bool
	Protolog           bool
	NoDefaultModifiers bool
}

func NewGoPlugin(options GoPluginOptions) Plugin {
	return newGoPlugin(options)
}

type CppPluginOptions struct {
	Grpc bool
}

func NewCppPlugin(options CppPluginOptions) Plugin {
	return newCppPlugin(options)
}

type CompilerDirectives struct {
	ProtoPaths  []string
	Cpp         bool
	Go          bool
	Grpc        bool
	GrpcGateway bool
	Protolog    bool
}

type Compiler interface {
	Args(dirPath string, outDirPath string, directives *CompilerDirectives) ([]string, error)
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
