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
	flagSet.StringVar(&directives.CppRelOutDirPath, "cpp-rel-out", "", "Output cpp files.")
	flagSet.BoolVar(&directives.Csharp, "csharp", false, "Output csharp files.")
	flagSet.StringVar(&directives.CsharpRelOutDirPath, "csharp-rel-out", "", "Output csharp files.")
	flagSet.BoolVar(&directives.Objc, "objc", false, "Output objc files.")
	flagSet.StringVar(&directives.ObjcRelOutDirPath, "objc-rel-out", "", "Output objc files.")
	flagSet.BoolVar(&directives.Python, "python", false, "Output python files.")
	flagSet.StringVar(&directives.PythonRelOutDirPath, "python-rel-out", "", "Output python files.")
	flagSet.BoolVar(&directives.Ruby, "ruby", false, "Output ruby files.")
	flagSet.StringVar(&directives.RubyRelOutDirPath, "ruby-rel-out", "", "Output ruby files.")
	flagSet.BoolVar(&directives.Go, "go", false, "Output go files.")
	flagSet.StringVar(&directives.GoRelOutDirPath, "go-rel-out", "", "Output go files.")
	flagSet.StringVar(&directives.GoImportPath, "go-import-path", "", "Go package.")
	flagSet.BoolVar(&directives.Grpc, "grpc", false, "Output grpc files.")
	flagSet.BoolVar(&directives.GrpcGateway, "grpc-gateway", false, "Output grpc-gateway .gw.go files.")
	flagSet.BoolVar(&directives.NoDefaultModifiers, "no-default-modifiers", false, "Do not set the default Mfile=package modifiers for --go_out.")
}

func bindOptions(flagSet *pflag.FlagSet, options *options) {
	flagSet.StringVar(&options.OutDirPath, "out", "", "Customize out directory path.")
}
