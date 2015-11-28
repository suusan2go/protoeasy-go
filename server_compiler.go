package protoeasy

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"go.pedge.io/pkg/exec"
)

type serverCompiler struct {
	protoSpecProvider ProtoSpecProvider
	options           ServerCompilerOptions
}

func newServerCompiler(
	protoSpecProvider ProtoSpecProvider,
	options ServerCompilerOptions,
) *serverCompiler {
	return &serverCompiler{
		protoSpecProvider,
		options,
	}
}

func (c *serverCompiler) Compile(dirPath string, outDirPath string, directives *Directives) error {
	var err error
	dirPath, err = filepath.Abs(dirPath)
	if err != nil {
		return err
	}
	outDirPath, err = filepath.Abs(outDirPath)
	if err != nil {
		return err
	}
	argsList, err := c.argsList(dirPath, outDirPath, directives)
	if err != nil {
		return err
	}
	if err := makeOutDirs(outDirPath, directives); err != nil {
		return err
	}
	for _, args := range argsList {
		if err := pkgexec.Run(args...); err != nil {
			return err
		}
	}
	return nil
}

func (c *serverCompiler) argsList(dirPath string, outDirPath string, directives *Directives) ([][]string, error) {
	plugins, err := getPlugins(directives)
	if err != nil {
		return nil, err
	}
	protoSpec, err := c.protoSpecProvider.Get(dirPath, directives.ExcludePattern)
	if err != nil {
		return nil, err
	}
	goPath, err := getGoPath()
	if err != nil {
		return nil, err
	}
	var argsList [][]string
	for relDirPath, files := range protoSpec.RelDirPathToFiles {
		for _, plugin := range plugins {
			args := []string{"protoc", fmt.Sprintf("-I%s", dirPath)}
			// TODO(pedge)
			if runtime.GOOS == "darwin" {
				args = append(args, "-I/usr/local/include")
			} else {
				args = append(args, "-I/usr/include")
			}
			args = append(args, fmt.Sprintf("-I%s", filepath.Join(goPath, "src/github.com/gengo/grpc-gateway/third_party/googleapis")))
			iArgs, err := plugin.Args(protoSpec, relDirPath, outDirPath)
			if err != nil {
				return nil, err
			}
			args = append(args, iArgs...)
			for _, file := range files {
				args = append(args, filepath.Join(dirPath, relDirPath, file))
			}
			argsList = append(argsList, args)
		}
	}
	return argsList, nil
}

func getPlugins(directives *Directives) ([]Plugin, error) {
	var plugins []Plugin
	if directives.Cpp {
		plugins = append(
			plugins,
			NewCppPlugin(
				CppPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: directives.CppRelOutDirPath,
						},
						Grpc: directives.Grpc,
					},
				},
			),
		)
	}
	if directives.Csharp {
		plugins = append(
			plugins,
			NewCsharpPlugin(
				CsharpPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: directives.CsharpRelOutDirPath,
						},
						Grpc: directives.Grpc,
					},
				},
			),
		)
	}
	if directives.Objectivec {
		plugins = append(
			plugins,
			NewObjectiveCPlugin(
				ObjectiveCPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: directives.ObjectivecRelOutDirPath,
						},
						Grpc: directives.Grpc,
					},
				},
			),
		)
	}
	if directives.Python {
		plugins = append(
			plugins,
			NewPythonPlugin(
				PythonPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: directives.PythonRelOutDirPath,
						},
						Grpc: directives.Grpc,
					},
				},
			),
		)
	}
	if directives.Ruby {
		plugins = append(
			plugins,
			NewRubyPlugin(
				RubyPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: directives.RubyRelOutDirPath,
						},
						Grpc: directives.Grpc,
					},
				},
			),
		)
	}
	if directives.Go {
		plugins = append(
			plugins,
			NewGoPlugin(
				GoPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: directives.GoRelOutDirPath,
						},
						Grpc: directives.Grpc,
					},
					ImportPath:  directives.GoImportPath,
					GrpcGateway: directives.GrpcGateway,
					Protolog:    directives.Protolog,
				},
			),
		)
	}
	return plugins, nil
}

func makeOutDirs(outDirPath string, directives *Directives) error {
	if err := os.MkdirAll(outDirPath, 0755); err != nil {
		return err
	}
	for _, relDirPath := range []string{
		directives.CppRelOutDirPath,
		directives.CsharpRelOutDirPath,
		directives.ObjectivecRelOutDirPath,
		directives.PythonRelOutDirPath,
		directives.RubyRelOutDirPath,
		directives.GoRelOutDirPath,
	} {
		if relDirPath != "" {
			if err := os.MkdirAll(filepath.Join(outDirPath, relDirPath), 0755); err != nil {
				return err
			}
		}
	}
	return nil
}
