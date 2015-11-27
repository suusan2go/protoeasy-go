package main

import (
	"fmt"
	"os"
	"strings"

	"go.pedge.io/pkg/cobra"
	"go.pedge.io/protoeasy"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func main() {
	if err := do(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func do() error {
	compilerDirectives := &protoeasy.CompilerDirectives{}

	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "Compile protocol buffers files",
		Long:  "Compile protocol buffers files",
		Run: pkgcobra.RunBoundedArgs(
			pkgcobra.Bounds{
				Min: 1,
				Max: 2,
			},
			func(args []string) error {
				dirPath := args[0]
				outDirPath := args[0]
				if len(args) == 2 {
					outDirPath = args[1]
				}
				return compile(dirPath, outDirPath, compilerDirectives)
			},
		),
	}

	pflag.StringSliceVar(&compilerDirectives.ExcludeFilePatterns, "exclude", []string{}, "Exclude file patterns.")
	pflag.StringSliceVarP(&compilerDirectives.ProtoPaths, "proto_path", "I", []string{}, "Additional directories for proto imports.")
	pflag.BoolVar(&compilerDirectives.Cpp, "cpp", false, "Output cpp files.")
	pflag.BoolVar(&compilerDirectives.Go, "go", false, "Output go files.")
	pflag.BoolVar(&compilerDirectives.Grpc, "grpc", false, "Output grpc files.")
	pflag.BoolVar(&compilerDirectives.GrpcGateway, "grpc-gateway", false, "Output grpc-gateway files.")
	pflag.BoolVar(&compilerDirectives.Protolog, "protolog", false, "Output protolog files.")
	return rootCmd.Execute()
}

func compile(dirPath string, outDirPath string, compilerDirectives *protoeasy.CompilerDirectives) error {
	argsList, err := protoeasy.NewCompiler(
		protoeasy.NewProtoSpecProvider(
			protoeasy.ProtoSpecProviderOptions{},
		),
		protoeasy.CompilerOptions{},
	).ArgsList(
		dirPath,
		outDirPath,
		compilerDirectives,
	)
	if err != nil {
		return err
	}
	for _, args := range argsList {
		fmt.Println(strings.Join(args, " "))
	}
	return nil
}
