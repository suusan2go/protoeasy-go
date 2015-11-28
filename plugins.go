package protoeasy

import (
	"fmt"
	"path/filepath"
	"strings"
)

var (
	defaultGoModifierOptions = mergeStringStringMaps(
		newGoModifierOptions(
			"google/protobuf",
			[]string{
				"any.proto",
				"api.proto",
				"duration.proto",
				"empty.proto",
				"field_mask.proto",
				"source_context.proto",
				"struct.proto",
				"timestamp.proto",
				"type.proto",
				"wrappers.proto",
			},
			"go.pedge.io/google-protobuf",
		),
		map[string]string{
			"google/api/annotations.proto":     "github.com/gengo/grpc-gateway/third_party/googleapis/google/api",
			"google/api/http.proto":            "github.com/gengo/grpc-gateway/third_party/googleapis/google/api",
			"google/protobuf/descriptor.proto": "github.com/golang/protobuf/protoc-gen-go/descriptor",
		},
	)
)

type goPlugin struct {
	options GoPluginOptions
}

func newGoPlugin(options GoPluginOptions) *goPlugin {
	return &goPlugin{options}
}

func (p *goPlugin) Args(protoSpec *ProtoSpec, relDirPath string, outDirPath string) ([]string, error) {
	if p.options.RelOutDirPath != "" {
		outDirPath = filepath.Join(outDirPath, p.options.RelOutDirPath)
	}
	goPath, err := getGoPath()
	if err != nil {
		return nil, err
	}
	args := []string{fmt.Sprintf("-I%s", filepath.Join(goPath, "src"))}
	modifiers := p.getModifiers(protoSpec)
	var goOutOpts []string
	for key, value := range modifiers {
		goOutOpts = append(goOutOpts, fmt.Sprintf("M%s=%s", key, value))
	}
	if p.options.Grpc {
		goOutOpts = append(goOutOpts, "plugins=grpc")
	}
	if len(goOutOpts) > 0 {
		args = append(args, fmt.Sprintf("--go_out=%s:%s", strings.Join(goOutOpts, ","), outDirPath))
	} else {
		args = append(args, fmt.Sprintf("--go_out=%s", outDirPath))
	}
	if p.options.GrpcGateway {
		var grpcGatewayOutOpts []string
		for key, value := range modifiers {
			grpcGatewayOutOpts = append(grpcGatewayOutOpts, fmt.Sprintf("M%s=%s", key, value))
		}
		if len(grpcGatewayOutOpts) > 0 {
			args = append(args, fmt.Sprintf("--grpc-gateway_out=%s:%s", strings.Join(grpcGatewayOutOpts, ","), outDirPath))
		} else {
			args = append(args, fmt.Sprintf("--grpc-gateway_out=%s", outDirPath))
		}
	}
	if p.options.Protolog {
		args = append(args, fmt.Sprintf("--protolog_out=%s", outDirPath))
	}
	return args, nil
}

func (p *goPlugin) getModifiers(protoSpec *ProtoSpec) map[string]string {
	modifiers := copyStringStringMap(defaultGoModifierOptions)
	for relDirPath, files := range protoSpec.RelDirPathToFiles {
		importPath := relDirPath
		if p.options.ImportPath != "" {
			importPath = filepath.Join(p.options.ImportPath, relDirPath)
		}
		// TODO(pedge)
		if importPath != "" && importPath != "." {
			for _, file := range files {
				modifiers[filepath.Join(relDirPath, file)] = importPath
			}
		}
	}
	return modifiers
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
	grpcName string
	options  GrpcPluginOptions
}

func newGrpcPlugin(name string, grpcName string, options GrpcPluginOptions) *grpcPlugin {
	return &grpcPlugin{newPlugin(name, options.PluginOptions), grpcName, options}
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
		whichGrpcPlugin, err := which(fmt.Sprintf("grpc_%s_plugin", p.grpcName))
		if err != nil {
			return nil, err
		}
		return append(args, fmt.Sprintf("--grpc_out=%s", outDirPath), fmt.Sprintf("--plugin=protoc-gen-grpc=%s", whichGrpcPlugin)), nil
	}
	return args, nil
}

func newGoModifierOptions(dir string, files []string, goPackage string) map[string]string {
	m := make(map[string]string)
	for _, file := range files {
		m[filepath.Join(dir, file)] = goPackage
	}
	return m
}
