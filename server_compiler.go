package protoeasy

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"go.pedge.io/pkg/exec"
)

type serverCompiler struct {
	options CompilerOptions
}

func newServerCompiler(options CompilerOptions) *serverCompiler {
	return &serverCompiler{options}
}

func (c *serverCompiler) Compile(dirPath string, outDirPath string, compileOptions *CompileOptions) ([]*Command, error) {
	var err error
	dirPath, err = filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	outDirPath, err = filepath.Abs(outDirPath)
	if err != nil {
		return nil, err
	}
	commands, err := c.commands(dirPath, outDirPath, compileOptions)
	if err != nil {
		return nil, err
	}
	relOutDirPaths := getRelOutDirPaths(compileOptions)
	if err := mkdir(outDirPath, relOutDirPaths...); err != nil {
		return nil, err
	}
	for _, command := range commands {
		if err := pkgexec.Run(command.Arg...); err != nil {
			// TODO(pedge): temporary until https://github.com/grpc/grpc-go/issues/606 is fixed
			return nil, errors.New(strings.Replace(err.Error(), "\n", " ", -1))
		}
	}
	return commands, nil
}

func (c *serverCompiler) commands(dirPath string, outDirPath string, compileOptions *CompileOptions) ([]*Command, error) {
	// getPlugins in plugins.go
	plugins := getPlugins(compileOptions)
	protoSpec, err := getProtoSpec(dirPath, compileOptions.ExcludePattern)
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
			if !compileOptions.NoDefaultIncludes {
				for _, goPathRelInclude := range defaultGoPathRelIncludes {
					args = append(args, fmt.Sprintf("-I%s", filepath.Join(goPath, goPathRelInclude)))
				}
			}
			flags, err := plugin.Flags(protoSpec, relDirPath, outDirPath)
			if err != nil {
				return nil, err
			}
			args = append(args, flags...)
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
