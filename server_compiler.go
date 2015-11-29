package protoeasy

import (
	"fmt"
	"os"
	"path/filepath"

	"go.pedge.io/pkg/exec"
)

type serverCompiler struct {
	options CompilerOptions
}

func newServerCompiler(options CompilerOptions) *serverCompiler {
	return &serverCompiler{options}
}

func (c *serverCompiler) Compile(dirPath string, outDirPath string, directives *Directives) ([]*Command, error) {
	var err error
	dirPath, err = filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	outDirPath, err = filepath.Abs(outDirPath)
	if err != nil {
		return nil, err
	}
	commands, err := c.commands(dirPath, outDirPath, directives)
	if err != nil {
		return nil, err
	}
	if err := makeOutDirs(outDirPath, directives); err != nil {
		return nil, err
	}
	for _, command := range commands {
		logCommand(command)
		if err := pkgexec.Run(command.Arg...); err != nil {
			return nil, err
		}
	}
	return commands, nil
}

func (c *serverCompiler) commands(dirPath string, outDirPath string, directives *Directives) ([]*Command, error) {
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
	var commands []*Command
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
			commands = append(commands, &Command{Arg: args})
		}
	}
	return commands, nil
}

func getProtoSpec(dirPath string, excludeFilePatterns []string) (*protoSpec, error) {
	relFilePaths, err := getAllRelProtoFilePaths(dirPath)
	if err != nil {
		return nil, err
	}
	relFilePaths, err = filterFilePaths(relFilePaths, excludeFilePatterns)
	if err != nil {
		return nil, err
	}
	return &protoSpec{
		DirPath:           dirPath,
		RelDirPathToFiles: getRelDirPathToFiles(relFilePaths),
	}, nil
}

func getPlugins(directives *Directives) ([]plugin, error) {
	var plugins []plugin
	for _, pluginType := range directives.PluginType {
		switch pluginType {
		case PluginType_PLUGIN_TYPE_CPP:
			plugins = append(plugins, newCppPlugin(directives.CppPluginOptions))
		case PluginType_PLUGIN_TYPE_CSHARP:
			plugins = append(plugins, newCsharpPlugin(directives.CsharpPluginOptions))
		case PluginType_PLUGIN_TYPE_GO:
			plugins = append(plugins, newGoPlugin(directives.GoPluginOptions))
		case PluginType_PLUGIN_TYPE_OBJC:
			plugins = append(plugins, newObjcPlugin(directives.ObjcPluginOptions))
		case PluginType_PLUGIN_TYPE_PYTHON:
			plugins = append(plugins, newPythonPlugin(directives.PythonPluginOptions))
		case PluginType_PLUGIN_TYPE_RUBY:
			plugins = append(plugins, newRubyPlugin(directives.RubyPluginOptions))
		default:
			return nil, fmt.Errorf("proteasy: invalid PluginType: %v", pluginType)
		}
	}
	return plugins, nil
}

func makeOutDirs(outDirPath string, directives *Directives) error {
	if err := os.MkdirAll(outDirPath, 0755); err != nil {
		return err
	}
	if directives.CppPluginOptions != nil {
		if err := makeRelOutDirForGrpcPluginOptions(outDirPath, directives.CppPluginOptions.GrpcPluginOptions); err != nil {
			return err
		}
	}
	if directives.CsharpPluginOptions != nil {
		if err := makeRelOutDirForGrpcPluginOptions(outDirPath, directives.CsharpPluginOptions.GrpcPluginOptions); err != nil {
			return err
		}
	}
	if directives.GoPluginOptions != nil {
		if err := makeRelOutDirForGrpcPluginOptions(outDirPath, directives.GoPluginOptions.GrpcPluginOptions); err != nil {
			return err
		}
	}
	if directives.ObjcPluginOptions != nil {
		if err := makeRelOutDirForGrpcPluginOptions(outDirPath, directives.ObjcPluginOptions.GrpcPluginOptions); err != nil {
			return err
		}
	}
	if directives.PythonPluginOptions != nil {
		if err := makeRelOutDirForGrpcPluginOptions(outDirPath, directives.PythonPluginOptions.GrpcPluginOptions); err != nil {
			return err
		}
	}
	if directives.RubyPluginOptions != nil {
		if err := makeRelOutDirForGrpcPluginOptions(outDirPath, directives.RubyPluginOptions.GrpcPluginOptions); err != nil {
			return err
		}
	}
	return nil
}

func makeRelOutDirForGrpcPluginOptions(outDirPath string, grpcPluginOptions *GrpcPluginOptions) error {
	relOutDirPath := getRelOutDirPath(grpcPluginOptions)
	if relOutDirPath != "" {
		return os.MkdirAll(filepath.Join(outDirPath, relOutDirPath), 0755)
	}
	return nil
}
