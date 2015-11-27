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
	compilerDirectives := &protoeasy.CompilerDirectives{}
	options := &options{}

	bindCompilerDirectives(pflag.CommandLine, compilerDirectives)
	bindOptions(pflag.CommandLine, options)
	pflag.Parse()
	args := pflag.CommandLine.Args()
	if len(args) != 1 {
		return fmt.Errorf("%s: must pass only one argument, the directory, but passed %v", os.Args[0], args)
	}

	return protoeasy.DefaultCompiler.Compile(args[0], compilerDirectives)
}

func bindCompilerDirectives(flagSet *pflag.FlagSet, compilerDirectives *protoeasy.CompilerDirectives) {
	flagSet.StringSliceVar(&compilerDirectives.ExcludeFilePatterns, "exclude", []string{}, "Exclude file patterns.")
	flagSet.StringSliceVarP(&compilerDirectives.ProtoPaths, "proto_path", "I", []string{}, "Additional directories for proto imports.")
	flagSet.StringVar(&compilerDirectives.OutDirPath, "out", "", "Customize out directory path.")
	flagSet.BoolVar(&compilerDirectives.Cpp, "cpp", false, "Output cpp files.")
	flagSet.StringVar(&compilerDirectives.CppRelOutDirPath, "cpp_rel_out", "", "Output cpp files.")
	flagSet.BoolVar(&compilerDirectives.Csharp, "csharp", false, "Output csharp files.")
	flagSet.StringVar(&compilerDirectives.CsharpRelOutDirPath, "csharp_rel_out", "", "Output csharp files.")
	flagSet.BoolVar(&compilerDirectives.ObjectiveC, "objectivec", false, "Output objectivec files.")
	flagSet.StringVar(&compilerDirectives.ObjectiveCRelOutDirPath, "objectivec_rel_out", "", "Output objectivec files.")
	flagSet.BoolVar(&compilerDirectives.Python, "python", false, "Output python files.")
	flagSet.StringVar(&compilerDirectives.PythonRelOutDirPath, "python_rel_out", "", "Output python files.")
	flagSet.BoolVar(&compilerDirectives.Ruby, "ruby", false, "Output ruby files.")
	flagSet.StringVar(&compilerDirectives.RubyRelOutDirPath, "ruby_rel_out", "", "Output ruby files.")
	flagSet.BoolVar(&compilerDirectives.Go, "go", false, "Output go files.")
	flagSet.StringVar(&compilerDirectives.GoRelOutDirPath, "go_rel_out", "", "Output go files.")
	flagSet.StringVar(&compilerDirectives.GoImportPath, "go_import_path", "", "Go package.")
	flagSet.BoolVar(&compilerDirectives.Grpc, "grpc", false, "Output grpc files.")
	flagSet.BoolVar(&compilerDirectives.GrpcGateway, "grpc-gateway", false, "Output grpc-gateway files.")
	flagSet.BoolVar(&compilerDirectives.Protolog, "protolog", false, "Output protolog files.")
}

func bindOptions(flagSet *pflag.FlagSet, options *options) {
}
