package protoeasy

import (
	"fmt"
	"path/filepath"
)

type goPlugin struct {
	options GoPluginOptions
}

func newGoPlugin(options GoPluginOptions) *goPlugin {
	return &goPlugin{options}
}

func (p *goPlugin) Args(protoSpec *ProtoSpec, relDirPath string, outDirPath string) ([]string, error) {
	return nil, nil
}

type plugin struct {
	name    string
	options PluginOptions
}

func newPlugin(name string, options PluginOptions) *plugin {
	return &plugin{name, options}
}

func (p *plugin) Args(protoSpec *ProtoSpec, relDirPath string, outDirPath string) ([]string, error) {
	if p.options.RelOutDirPath != "" {
		outDirPath = filepath.Join(outDirPath, p.options.RelOutDirPath)
	}
	return []string{fmt.Sprintf("--%s_out=%s", p.name, outDirPath)}, nil
}

type grpcPlugin struct {
	*plugin
	options GrpcPluginOptions
}

func newGrpcPlugin(name string, options GrpcPluginOptions) *grpcPlugin {
	return &grpcPlugin{newPlugin(name, options.PluginOptions), options}
}

func (p *grpcPlugin) Args(protoSpec *ProtoSpec, relDirPath string, outDirPath string) ([]string, error) {
	args, err := p.plugin.Args(protoSpec, relDirPath, outDirPath)
	if err != nil {
		return nil, err
	}
	if p.options.Grpc {
		if p.options.RelOutDirPath != "" {
			outDirPath = filepath.Join(outDirPath, p.options.RelOutDirPath)
		}
		return append(args, fmt.Sprintf("--grpc_%s_out=%s", p.name, outDirPath)), nil
	}
	return args, nil
}
