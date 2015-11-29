package protoeasy

import (
	"fmt"
	"os"
	"path/filepath"

	"go.pedge.io/pkg/exec"
)

type serverCompiler struct {
	options ServerCompilerOptions
}

func newServerCompiler(options ServerCompilerOptions) *serverCompiler {
	return &serverCompiler{options}
}

func (c *serverCompiler) Compile(dirPath string, outDirPath string, directives *Directives) ([][]string, error) {
	var err error
	dirPath, err = filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	outDirPath, err = filepath.Abs(outDirPath)
	if err != nil {
		return nil, err
	}
	argsList, err := c.argsList(dirPath, outDirPath, directives)
	if err != nil {
		return nil, err
	}
	if err := makeOutDirs(outDirPath, directives); err != nil {
		return nil, err
	}
	for _, args := range argsList {
		logArgs(args)
		if err := pkgexec.Run(args...); err != nil {
			return nil, err
		}
	}
	return argsList, nil
}

func (c *serverCompiler) argsList(dirPath string, outDirPath string, directives *Directives) ([][]string, error) {
	plugins, err := getPlugins(directives)
	if err != nil {
		return nil, err
	}
	protoSpec, err := getProtoSpec(dirPath, directives.ExcludePattern)
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
			args = append(args, fmt.Sprintf("-I%s", filepath.Join(goPath, "src/go.pedge.io/protoeasy/vendor/go.pedge.io/google-protobuf")))
			args = append(args, fmt.Sprintf("-I%s", filepath.Join(goPath, "src/go.pedge.io/protoeasy/vendor/github.com/golang/protobuf/protoc-gen-go/descriptor")))
			args = append(args, fmt.Sprintf("-I%s", filepath.Join(goPath, "src/go.pedge.io/protoeasy/vendor/github.com/gengo/grpc-gateway/third_party/googleapis")))
			args = append(args, fmt.Sprintf("-I%s", filepath.Join(goPath, "src/go.pedge.io/protoeasy/vendor/go.pedge.io/googleapis")))
			iArgs, err := plugin.Flags(protoSpec, relDirPath, outDirPath)
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

func getProtoSpec(dirPath string, excludeFilePatterns []string) (*ProtoSpec, error) {
	relFilePaths, err := getAllRelProtoFilePaths(dirPath)
	if err != nil {
		return nil, err
	}
	relFilePaths, err = filterFilePaths(relFilePaths, excludeFilePatterns)
	if err != nil {
		return nil, err
	}
	return &ProtoSpec{
		DirPath:           dirPath,
		RelDirPathToFiles: getRelDirPathToFiles(relFilePaths),
	}, nil
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
	if directives.Objc {
		plugins = append(
			plugins,
			NewObjcPlugin(
				ObjcPluginOptions{
					GrpcPluginOptions: GrpcPluginOptions{
						PluginOptions: PluginOptions{
							RelOutDirPath: directives.ObjcRelOutDirPath,
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
		protocPlugin := directives.GoProtocPlugin
		if protocPlugin == GoProtocPlugin_GO_PROTOC_PLUGIN_NONE {
			protocPlugin = GoProtocPlugin_GO_PROTOC_PLUGIN_GO
		}
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
					ImportPath:         directives.GoImportPath,
					GrpcGateway:        directives.GrpcGateway,
					NoDefaultModifiers: directives.GoNoDefaultModifiers,
					ProtocPlugin:       protocPlugin,
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
		directives.ObjcRelOutDirPath,
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
