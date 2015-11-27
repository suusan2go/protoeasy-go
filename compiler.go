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

func (c *compiler) Args(dirPath string, outDirPath string, directives *CompilerDirectives) ([]string, error) {
	var err error
	dirPath, err = filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	outDirPath, err = filepath.Abs(outDirPath)
	if err != nil {
		return nil, err
	}
	protoSpec, err := c.protoSpecProvider.Get(dirPath)
	if err != nil {
		return nil, err
	}
	plugins, err := c.getPlugins(directives)
	if err != nil {
		return nil, err
	}
	args := []string{"protoc"}
	for _, protoPath := range directives.ProtoPaths {
		protoPath, err = filepath.Abs(dirPath)
		if err != nil {
			return nil, err
		}
		args = append(args, "-I", protoPath)
	}
	for _, plugin := range plugins {
		iArgs, err := plugin.Args(protoSpec, outDirPath)
		if err != nil {
			return nil, err
		}
		args = append(args, iArgs...)
	}
	return args, nil
}

func (c *compiler) getPlugins(directives *CompilerDirectives) ([]Plugin, error) {
	return nil, nil
}
