package protoeasy

import (
	"fmt"
	"path/filepath"
	"sort"
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
		newGoModifierOptions(
			"google/protobuf",
			[]string{
				"descriptor.proto",
			},
			"github.com/golang/protobuf/protoc-gen-go/descriptor",
		),
		newGoModifierOptions(
			"google/api",
			[]string{
				"annotations.proto",
				"http.proto",
			},
			"github.com/gengo/grpc-gateway/third_party/googleapis/google/api",
		),
		newGoModifierOptions(
			"google/datastore/v1beta3",
			[]string{
				"datastore.proto",
				"entity.proto",
				"query.proto",
			},
			"go.pedge.io/googleapis/google/datastore/v1beta3",
		),
		newGoModifierOptions(
			"google/devtools/cloudtrace/v1",
			[]string{
				"trace.proto",
			},
			"go.pedge.io/googleapis/google/devtools/cloudtrace/v1",
		),
		newGoModifierOptions(
			"google/example/library/v1",
			[]string{
				"library.proto",
			},
			"go.pedge.io/googleapis/google/example/library/v1",
		),
		newGoModifierOptions(
			"google/iam/v1",
			[]string{
				"iam_policy.proto",
				"policy.proto",
			},
			"go.pedge.io/googleapis/google/iam/v1",
		),
		newGoModifierOptions(
			"google/longrunning",
			[]string{
				"operations.proto",
			},
			"go.pedge.io/googleapis/google/longrunning",
		),
		newGoModifierOptions(
			"google/pubsub/v1",
			[]string{
				"pubsub.proto",
			},
			"go.pedge.io/googleapis/google/pubsub/v1",
		),
		newGoModifierOptions(
			"google/pubsub/v1beta2",
			[]string{
				"pubsub.proto",
			},
			"go.pedge.io/googleapis/google/pubsub/v1beta2",
		),
		newGoModifierOptions(
			"google/rpc",
			[]string{
				"code.proto",
				"error_details.proto",
				"status.proto",
			},
			"go.pedge.io/googleapis/google/rpc",
		),
		newGoModifierOptions(
			"google/type",
			[]string{
				"color.proto",
				"date.proto",
				"dayofweek.proto",
				"latlng.proto",
				"money.proto",
				"timeofday.proto",
			},
			"go.pedge.io/googleapis/google/type",
		),
	)
)

func newGoModifierOptions(dir string, files []string, goPackage string) map[string]string {
	m := make(map[string]string)
	for _, file := range files {
		m[filepath.Join(dir, file)] = goPackage
	}
	return m
}

// plugin is an individual flag provider for a specific language.
type plugin interface {
	// Flags gets the protoc flags for the plugin.
	Flags(protoSpec *protoSpec, relDirPath string, outDirPath string) ([]string, error)
}

type goPlugin struct {
	options *GoPluginOptions
}

func newGoPlugin(options *GoPluginOptions) *goPlugin {
	// since we switch to using generated protos, we have pointers now
	// for options instead of structs, so we do this to be safe, even
	// though we handle it in the rest of the code
	if options == nil {
		options = &GoPluginOptions{}
	}
	if options.GrpcPluginOptions == nil {
		options.GrpcPluginOptions = &GrpcPluginOptions{}
	}
	if options.GrpcPluginOptions.PluginOptions == nil {
		options.GrpcPluginOptions.PluginOptions = &PluginOptions{}
	}
	return &goPlugin{options}
}

func (p *goPlugin) Flags(protoSpec *protoSpec, relDirPath string, outDirPath string) ([]string, error) {
	relOutDirPath := getRelOutDirPath(p.options.GrpcPluginOptions)
	if relOutDirPath != "" {
		outDirPath = filepath.Join(outDirPath, relOutDirPath)
	}
	var args []string
	modifiers := p.getModifiers(protoSpec)
	goOutOpts := modifiers
	if isGrpc(p.options.GrpcPluginOptions) {
		goOutOpts = fmt.Sprintf("%s,plugins=grpc", goOutOpts)
	}
	if len(goOutOpts) > 0 {
		args = append(args, fmt.Sprintf("--%s_out=%s:%s", p.options.PluginType.SimpleString(), goOutOpts, outDirPath))
	} else {
		args = append(args, fmt.Sprintf("--%s_out=%s", p.options.PluginType.SimpleString(), outDirPath))
	}
	if p.options.GrpcGateway {
		if len(modifiers) > 0 {
			args = append(args, fmt.Sprintf("--grpc-gateway_out=%s:%s", modifiers, outDirPath))
		} else {
			args = append(args, fmt.Sprintf("--grpc-gateway_out=%s", outDirPath))
		}
	}
	return args, nil
}

func (p *goPlugin) getModifiers(protoSpec *protoSpec) string {
	var modifiers map[string]string
	if p.options.NoDefaultModifiers {
		modifiers = make(map[string]string)
	} else {
		modifiers = copyStringStringMap(defaultGoModifierOptions)
	}
	for key, value := range p.options.Modifiers {
		modifiers[key] = value
	}
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
	// this is just so the command line is easier to read and deterministic
	modifierKeys := make([]string, len(modifiers))
	i := 0
	for key := range modifiers {
		modifierKeys[i] = key
		i++
	}
	sort.Strings(modifierKeys)
	modifierStrings := make([]string, len(modifierKeys))
	for i, modifierKey := range modifierKeys {
		modifierStrings[i] = fmt.Sprintf("M%s=%s", modifierKey, modifiers[modifierKey])
	}
	return strings.Join(modifierStrings, ",")
}

type basePlugin struct {
	name    string
	options *PluginOptions
}

func newBasePlugin(name string, options *PluginOptions) *basePlugin {
	if options == nil {
		options = &PluginOptions{}
	}
	return &basePlugin{name, options}
}

func (p *basePlugin) Flags(protoSpec *protoSpec, relDirPath string, outDirPath string) ([]string, error) {
	if p.options.RelOutDirPath != "" {
		outDirPath = filepath.Join(outDirPath, p.options.RelOutDirPath)
	}
	return []string{fmt.Sprintf("--%s_out=%s", p.name, outDirPath)}, nil
}

type grpcPlugin struct {
	*basePlugin
	grpcName string
	options  *GrpcPluginOptions
}

func newGrpcPlugin(name string, grpcName string, options *GrpcPluginOptions) *grpcPlugin {
	if options == nil {
		options = &GrpcPluginOptions{}
	}
	if options.PluginOptions == nil {
		options.PluginOptions = &PluginOptions{}
	}
	return &grpcPlugin{newBasePlugin(name, options.PluginOptions), grpcName, options}
}

func (p *grpcPlugin) Flags(protoSpec *protoSpec, relDirPath string, outDirPath string) ([]string, error) {
	args, err := p.basePlugin.Flags(protoSpec, relDirPath, outDirPath)
	if err != nil {
		return nil, err
	}
	if isGrpc(p.options) {
		relOutDirPath := getRelOutDirPath(p.options)
		if relOutDirPath != "" {
			outDirPath = filepath.Join(outDirPath, relOutDirPath)
		}
		whichGrpcPlugin, err := which(fmt.Sprintf("grpc_%s_plugin", p.grpcName))
		if err != nil {
			return nil, err
		}
		return append(args, fmt.Sprintf("--grpc_out=%s", outDirPath), fmt.Sprintf("--plugin=protoc-gen-grpc=%s", whichGrpcPlugin)), nil
	}
	return args, nil
}

func newCppPlugin(options *CppPluginOptions) plugin {
	if options == nil {
		options = &CppPluginOptions{}
	}
	return newGrpcPlugin("cpp", "cpp", options.GrpcPluginOptions)
}

func newCsharpPlugin(options *CsharpPluginOptions) plugin {
	if options == nil {
		options = &CsharpPluginOptions{}
	}
	return newGrpcPlugin("csharp", "csharp", options.GrpcPluginOptions)
}

func newObjcPlugin(options *ObjcPluginOptions) plugin {
	if options == nil {
		options = &ObjcPluginOptions{}
	}
	return newGrpcPlugin("objc", "objective_c", options.GrpcPluginOptions)
}

func newPythonPlugin(options *PythonPluginOptions) plugin {
	if options == nil {
		options = &PythonPluginOptions{}
	}
	return newGrpcPlugin("python", "python", options.GrpcPluginOptions)
}

func newRubyPlugin(options *RubyPluginOptions) plugin {
	if options == nil {
		options = &RubyPluginOptions{}
	}
	return newGrpcPlugin("ruby", "ruby", options.GrpcPluginOptions)
}
