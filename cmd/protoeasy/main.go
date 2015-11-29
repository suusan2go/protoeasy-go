package main

import (
	"fmt"
	"os"
	"strings"

	"google.golang.org/grpc"

	"go.pedge.io/env"
	"go.pedge.io/protoeasy"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type appEnv struct {
	Address string `env:"PROTOEASY_ADDRESS"`
}

type options struct {
	GoModifiers    []string
	GoProtocPlugin string
	OutDirPath     string
}

func main() {
	env.Main(do, &appEnv{})
}

func do(appEnvObj interface{}) error {
	appEnv := appEnvObj.(*appEnv)
	directives := &protoeasy.Directives{}
	options := &options{}

	rootCmd := &cobra.Command{
		Use: fmt.Sprintf("%s directory", os.Args[0]),
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("must pass one argument, the directory, but passed %d arguments", len(args))
			}
			if err := optionsToDirectives(options, directives); err != nil {
				return err
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
		},
	}
	bindDirectives(rootCmd.Flags(), directives)
	bindOptions(rootCmd.Flags(), options)

	return rootCmd.Execute()
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
	flagSet.BoolVar(&directives.GoNoDefaultModifiers, "go-no-default-modifiers", false, "Do not set the default Mfile=package modifiers for --go_out.")
	flagSet.BoolVar(&directives.Grpc, "grpc", false, "Output grpc files.")
	flagSet.BoolVar(&directives.GrpcGateway, "grpc-gateway", false, "Output grpc-gateway .gw.go files.")
}

func bindOptions(flagSet *pflag.FlagSet, options *options) {
	flagSet.StringSliceVar(&options.GoModifiers, "go-modifier", []string{}, "Extra Mfile=package modifiers for --go_out, specify just as file=package to this flag.")
	flagSet.StringVar(&options.GoProtocPlugin, "go-protoc-plugin", "go", fmt.Sprintf("The go protoc plugin to use, allowed values are %s, if not go, --go-no-default-modifiers is implied.", strings.Join(protoeasy.AllGoProtocPluginSimpleStrings(), ",")))
	flagSet.StringVar(&options.OutDirPath, "out", "", "Customize out directory path.")
}

func optionsToDirectives(options *options, directives *protoeasy.Directives) error {
	if strings.ToLower(options.GoProtocPlugin) == "none" {
		return fmt.Errorf("invalid value for --go-protoc-plugin: %s", options.GoProtocPlugin)
	}
	goProtocPlugin, err := protoeasy.GoProtocPluginSimpleValueOf(options.GoProtocPlugin)
	if err != nil {
		return err
	}
	directives.GoProtocPlugin = goProtocPlugin
	if goProtocPlugin != protoeasy.GoProtocPlugin_GO_PROTOC_PLUGIN_GO {
		directives.GoNoDefaultModifiers = true
	}
	modifiers := make(map[string]string)
	for _, modifierString := range options.GoModifiers {
		split := strings.SplitN(modifierString, "=", 2)
		if len(split) != 2 {
			return fmt.Errorf("invalid valid for --go-modifier: %s", modifierString)
		}
		modifiers[split[0]] = split[1]
	}
	directives.GoModifier = modifiers
	return nil
}
