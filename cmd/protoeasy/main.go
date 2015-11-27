package main

import (
	"fmt"
	"os"

	"go.pedge.io/protoeasy"
	"go.pedge.io/protolog"

	"github.com/spf13/pflag"
)

type options struct {
	OutDirPath string
}

func main() {
	if err := do(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func do() error {
	protolog.SetLevel(protolog.Level_LEVEL_DEBUG)
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

	return protoeasy.DefaultCompiler.Compile(dirPath, outDirPath, directives)
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
