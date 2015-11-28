package main

import (
	"fmt"
	"os"

	"google.golang.org/grpc"

	"go.pedge.io/env"
	"go.pedge.io/protoeasy"

	"github.com/spf13/pflag"
)

type appEnv struct {
	Address string `env:"PROTOEASY_ADDRESS"`
}

type options struct {
	OutDirPath string
}

func main() {
	env.Main(do, &appEnv{})
}

func do(appEnvObj interface{}) error {
	appEnv := appEnvObj.(*appEnv)
	directives := &protoeasy.Directives{}
	options := &options{}

	bindDirectives(pflag.CommandLine, directives)
	bindOptions(pflag.CommandLine, options)
	pflag.Parse()

	args := pflag.CommandLine.Args()
	if len(args) != 1 {
		return fmt.Errorf("%s: must pass one argument, the directory, but passed %v", os.Args[0], args)
	}
	dirPath := args[0]
	outDirPath := args[0]
	if options.OutDirPath != "" {
		outDirPath = options.OutDirPath
	}

	compiler := protoeasy.DefaultServerCompiler
	if appEnv.Address != "" {
		clientConn, err := grpc.Dial(appEnv.Address, grpc.WithInsecure())
		if err != nil {
			return err
		}
		compiler = protoeasy.NewClientCompiler(
			protoeasy.NewAPIClient(
				clientConn,
			),
			protoeasy.ClientCompilerOptions{},
		)
	}

	_, err := compiler.Compile(dirPath, outDirPath, directives)
	return err
}

func bindDirectives(flagSet *pflag.FlagSet, directives *protoeasy.Directives) {
	flagSet.StringSliceVar(&directives.ExcludePattern, "exclude", []string{}, "Exclude file patterns.")
	flagSet.BoolVar(&directives.Cpp, "cpp", false, "Output cpp files.")
	flagSet.StringVar(&directives.CppRelOutDirPath, "cpp_rel_out", "", "Output cpp files.")
	flagSet.BoolVar(&directives.Csharp, "csharp", false, "Output csharp files.")
	flagSet.StringVar(&directives.CsharpRelOutDirPath, "csharp_rel_out", "", "Output csharp files.")
	flagSet.BoolVar(&directives.Objectivec, "objectivec", false, "Output objectivec files.")
	flagSet.StringVar(&directives.ObjectivecRelOutDirPath, "objectivec_rel_out", "", "Output objectivec files.")
	flagSet.BoolVar(&directives.Python, "python", false, "Output python files.")
	flagSet.StringVar(&directives.PythonRelOutDirPath, "python_rel_out", "", "Output python files.")
	flagSet.BoolVar(&directives.Ruby, "ruby", false, "Output ruby files.")
	flagSet.StringVar(&directives.RubyRelOutDirPath, "ruby_rel_out", "", "Output ruby files.")
	flagSet.BoolVar(&directives.Go, "go", false, "Output go files.")
	flagSet.StringVar(&directives.GoRelOutDirPath, "go_rel_out", "", "Output go files.")
	flagSet.StringVar(&directives.GoImportPath, "go_import_path", "", "Go package.")
	flagSet.BoolVar(&directives.Grpc, "grpc", false, "Output grpc files.")
	flagSet.BoolVar(&directives.GrpcGateway, "grpc-gateway", false, "Output grpc-gateway files.")
	flagSet.BoolVar(&directives.Protolog, "protolog", false, "Output protolog files.")
}

func bindOptions(flagSet *pflag.FlagSet, options *options) {
	flagSet.StringVar(&options.OutDirPath, "out", "", "Customize out directory path.")
}
