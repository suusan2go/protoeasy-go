package protoeasy

import (
	"fmt"
	"path/filepath"
	"runtime"

	"go.pedge.io/pkg/exec"
)

type compiler struct {
	protoSpecProvider ProtoSpecProvider
	options           CompilerOptions
}

func newCompiler(
	protoSpecProvider ProtoSpecProvider,
	options CompilerOptions,
) *compiler {
	return &compiler{
		protoSpecProvider,
		options,
	}
}

func (c *compiler) Compile(dirPath string, options *CompileOptions) error {
	argsList, err := c.argsList(dirPath, options)
	if err != nil {
		return err
	}
	for _, args := range argsList {
		if err := pkgexec.Run(args...); err != nil {
			return err
		}
	}
	return nil
}

func (c *compiler) argsList(dirPath string, options *CompileOptions) ([][]string, error) {
	var err error
	dirPath, err = filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	outDirPath := dirPath
	if options.OutDirPath != "" {
		outDirPath, err = filepath.Abs(options.OutDirPath)
		if err != nil {
			return nil, err
		}
	}
	plugins, err := getPlugins(options)
	if err != nil {
		return nil, err
	}
	protoSpec, err := c.protoSpecProvider.Get(dirPath, options.ExcludeFilePatterns)
	if err != nil {
		return nil, err
	}
	var argsList [][]string
	for relDirPath, files := range protoSpec.RelDirPathToFiles {
		args := []string{"protoc", fmt.Sprintf("-I%s", dirPath)}
		// TODO(pedge)
		if runtime.GOOS == "darwin" {
			args = append(args, "-I/usr/local/include")
		} else {
			args = append(args, "-I/usr/include")
		}
		for _, protoPath := range options.ProtoPaths {
			protoPath, err = filepath.Abs(dirPath)
			if err != nil {
				return nil, err
			}
			args = append(args, fmt.Sprintf("-I%s", protoPath))
		}
		for _, plugin := range plugins {
			iArgs, err := plugin.Args(protoSpec, relDirPath, outDirPath)
			if err != nil {
				return nil, err
			}
			args = append(args, iArgs...)
		}
		for _, file := range files {
			args = append(args, filepath.Join(dirPath, relDirPath, file))
		}
		argsList = append(argsList, args)
	}
	return argsList, nil
}

func getPlugins(options *CompileOptions) ([]Plugin, error) {
	var plugins []Plugin
	if options.Cpp {
		plugins = append(
			plugins,
			NewCppPlugin(
				CppPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: options.CppRelOutDirPath,
						},
						Grpc: options.Grpc,
					},
				},
			),
		)
	}
	if options.Csharp {
		plugins = append(
			plugins,
			NewCsharpPlugin(
				CsharpPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: options.CsharpRelOutDirPath,
						},
						Grpc: options.Grpc,
					},
				},
			),
		)
	}
	if options.ObjectiveC {
		plugins = append(
			plugins,
			NewObjectiveCPlugin(
				ObjectiveCPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: options.ObjectiveCRelOutDirPath,
						},
						Grpc: options.Grpc,
					},
				},
			),
		)
	}
	if options.Python {
		plugins = append(
			plugins,
			NewPythonPlugin(
				PythonPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: options.PythonRelOutDirPath,
						},
						Grpc: options.Grpc,
					},
				},
			),
		)
	}
	if options.Ruby {
		plugins = append(
			plugins,
			NewRubyPlugin(
				RubyPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: options.RubyRelOutDirPath,
						},
						Grpc: options.Grpc,
					},
				},
			),
		)
	}
	if options.Go {
		plugins = append(
			plugins,
			NewGoPlugin(
				GoPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: options.GoRelOutDirPath,
						},
						Grpc: options.Grpc,
					},
					ImportPath:  options.GoImportPath,
					GrpcGateway: options.GrpcGateway,
					Protolog:    options.Protolog,
				},
			),
		)
	}
	return plugins, nil
}
