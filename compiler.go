package protoeasy

import "path/filepath"

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

func (c *compiler) ArgsList(dirPath string, outDirPath string, directives *CompilerDirectives) ([][]string, error) {
	var err error
	dirPath, err = filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	outDirPath, err = filepath.Abs(outDirPath)
	if err != nil {
		return nil, err
	}
	plugins, err := c.getPlugins(directives)
	if err != nil {
		return nil, err
	}
	protoSpec, err := c.protoSpecProvider.Get(dirPath, directives.ExcludeFilePatterns)
	if err != nil {
		return nil, err
	}
	var argsList [][]string
	for relDirPath, files := range protoSpec.RelDirPathToFiles {
		args := []string{"protoc"}
		for _, protoPath := range directives.ProtoPaths {
			protoPath, err = filepath.Abs(dirPath)
			if err != nil {
				return nil, err
			}
			args = append(args, "-I", protoPath)
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

func (c *compiler) getPlugins(directives *CompilerDirectives) ([]Plugin, error) {
	var plugins []Plugin
	if directives.Cpp {
		plugins = append(
			plugins,
			NewCppPlugin(
				CppPluginOptions{
					Grpc: directives.Grpc,
				},
			),
		)
	}
	if directives.Go {
		plugins = append(
			plugins,
			NewGoPlugin(
				GoPluginOptions{
					Grpc:        directives.Grpc,
					GrpcGateway: directives.GrpcGateway,
					Protolog:    directives.Protolog,
				},
			),
		)
	}
	return plugins, nil
}
