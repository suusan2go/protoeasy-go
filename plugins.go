package protoeasy

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.pedge.io/protolog"
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
			"google/protobuf/descriptor.proto": "github.com/golang/protobuf/protoc-gen-go/descriptor",
		},
	)

	defaultGoPathRelativeIncludes = []string{
		"src",
		"src/github.com/gengo/grpc-gateway/third_party/googleapis",
	}

	errGoPathNotSet = errors.New("protoeasy: GOPATH not set")
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
	if p.options.GrpcGateway {
		args = append(args, fmt.Sprintf("-I%s", filepath.Join(goPath, "src/github.com/gengo/grpc-gateway/third_party/googleapis")))
		modifiers["google/api/http.proto"] = "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
		modifiers["google/api/annotations.proto"] = "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
	}
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

func getGoPath() (string, error) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		return "", errGoPathNotSet
	}
	split := strings.Split(goPath, ":")
	if len(split) > 1 {
		protolog.Warnf("protoeasy: GOPATH %s has multiple directories, using first directory %s", goPath, split[0])
		return split[0], nil
	}
	return goPath, nil
}

func newGoModifierOptions(dir string, files []string, goPackage string) map[string]string {
	m := make(map[string]string)
	for _, file := range files {
		m[filepath.Join(dir, file)] = goPackage
	}
	return m
}

func mergeStringStringMaps(maps ...map[string]string) map[string]string {
	newMap := make(map[string]string)
	for _, m := range maps {
		for key, value := range m {
			newMap[key] = value
		}
	}
	return newMap
}

func copyStringStringMap(m map[string]string) map[string]string {
	return mergeStringStringMaps(m)
}
