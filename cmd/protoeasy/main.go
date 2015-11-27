package main

import (
	"fmt"
	"os"

	"go.pedge.io/protoeasy"
	"go.pedge.io/protolog"

	"github.com/spf13/pflag"
)

type options struct{}

func main() {
	if err := do(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func do() error {
	protolog.SetLevel(protolog.Level_LEVEL_DEBUG)
	compileOptions := &protoeasy.CompileOptions{}
	options := &options{}

	bindCompileOptions(pflag.CommandLine, compileOptions)
	bindOptions(pflag.CommandLine, options)
	pflag.Parse()
	args := pflag.CommandLine.Args()
	if len(args) != 1 {
		return fmt.Errorf("%s: must pass only one argument, the directory, but passed %v", os.Args[0], args)
	}

	return protoeasy.DefaultCompiler.Compile(args[0], compileOptions)
}

func bindCompileOptions(flagSet *pflag.FlagSet, compileOptions *protoeasy.CompileOptions) {
	flagSet.StringSliceVar(&compileOptions.ExcludeFilePatterns, "exclude", []string{}, "Exclude file patterns.")
	flagSet.StringSliceVarP(&compileOptions.ProtoPaths, "proto_path", "I", []string{}, "Additional directories for proto imports.")
	flagSet.StringVar(&compileOptions.OutDirPath, "out", "", "Customize out directory path.")
	flagSet.BoolVar(&compileOptions.Cpp, "cpp", false, "Output cpp files.")
	flagSet.StringVar(&compileOptions.CppRelOutDirPath, "cpp_rel_out", "", "Output cpp files.")
	flagSet.BoolVar(&compileOptions.Csharp, "csharp", false, "Output csharp files.")
	flagSet.StringVar(&compileOptions.CsharpRelOutDirPath, "csharp_rel_out", "", "Output csharp files.")
	flagSet.BoolVar(&compileOptions.ObjectiveC, "objectivec", false, "Output objectivec files.")
	flagSet.StringVar(&compileOptions.ObjectiveCRelOutDirPath, "objectivec_rel_out", "", "Output objectivec files.")
	flagSet.BoolVar(&compileOptions.Python, "python", false, "Output python files.")
	flagSet.StringVar(&compileOptions.PythonRelOutDirPath, "python_rel_out", "", "Output python files.")
	flagSet.BoolVar(&compileOptions.Ruby, "ruby", false, "Output ruby files.")
	flagSet.StringVar(&compileOptions.RubyRelOutDirPath, "ruby_rel_out", "", "Output ruby files.")
	flagSet.BoolVar(&compileOptions.Go, "go", false, "Output go files.")
	flagSet.StringVar(&compileOptions.GoRelOutDirPath, "go_rel_out", "", "Output go files.")
	flagSet.StringVar(&compileOptions.GoImportPath, "go_import_path", "", "Go package.")
	flagSet.BoolVar(&compileOptions.Grpc, "grpc", false, "Output grpc files.")
	flagSet.BoolVar(&compileOptions.GrpcGateway, "grpc-gateway", false, "Output grpc-gateway files.")
	flagSet.BoolVar(&compileOptions.Protolog, "protolog", false, "Output protolog files.")
}

func bindOptions(flagSet *pflag.FlagSet, options *options) {
}
