package main

import (
	"fmt"
	"os"

	"go.pedge.io/pkg/cobra"
	"go.pedge.io/protoeasy"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	compilerDirectives = &protoeasy.CompilerDirectives{}
)

func init() {
	bindCompilerDirectives(compilerDirectives, pflag.CommandLine)
}

func main() {
	if err := do(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func do() error {
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
	return rootCmd.Execute()
}

func compile(dirPath string, outDirPath string, compilerDirectives *protoeasy.CompilerDirectives) error {
	fmt.Println(dirPath)
	fmt.Println(outDirPath)
	fmt.Printf("%+v\n", compilerDirectives)
	return nil
}

func bindCompilerDirectives(compilerDirectives *protoeasy.CompilerDirectives, flagSet *pflag.FlagSet) {
	flagSet.StringSliceVarP(&compilerDirectives.ProtoPaths, "proto_path", "I", []string{}, "Additional directories for proto imports.")
	flagSet.BoolVar(&compilerDirectives.Cpp, "cpp", false, "Output cpp files.")
	flagSet.BoolVar(&compilerDirectives.Go, "go", false, "Output go files.")
	flagSet.BoolVar(&compilerDirectives.Grpc, "grpc", false, "Output grpc files.")
	flagSet.BoolVar(&compilerDirectives.GrpcGateway, "grpc-gateway", false, "Output grpc-gateway files.")
	flagSet.BoolVar(&compilerDirectives.Protolog, "protolog", false, "Output protolog files.")
}
