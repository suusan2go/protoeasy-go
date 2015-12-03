package protoeasy

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// plugin is an individual flag provider for a specific language.
type plugin interface {
	// Flags gets the protoc flags for the plugin.
	Flags(protoSpec *protoSpec, relDirPath string, outDirPath string) ([]string, error)
}

func newCppPlugin(options *CompileOptions) plugin {
	return newGrpcPlugin("cpp", "cpp", options, options.CppRelOut)
}

func newCsharpPlugin(options *CompileOptions) plugin {
	return newGrpcPlugin("csharp", "csharp", options, options.CsharpRelOut)
}

func newObjcPlugin(options *CompileOptions) plugin {
	return newGrpcPlugin("objc", "objective_c", options, options.ObjcRelOut)
}

func newPythonPlugin(options *CompileOptions) plugin {
	return newGrpcPlugin("python", "python", options, options.PythonRelOut)
}

func newRubyPlugin(options *CompileOptions) plugin {
	return newGrpcPlugin("ruby", "ruby", options, options.RubyRelOut)
}

type goPlugin struct {
	options *CompileOptions
}

func newGoPlugin(options *CompileOptions) *goPlugin {
	if options == nil {
		options = &CompileOptions{}
	}
	return &goPlugin{options}
}

func (p *goPlugin) Flags(protoSpec *protoSpec, relDirPath string, outDirPath string) ([]string, error) {
	if p.options.GoRelOut != "" {
		outDirPath = filepath.Join(outDirPath, p.options.GoRelOut)
	}
	goPluginType := p.options.GoPluginType
	if goPluginType == GoPluginType_GO_PLUGIN_TYPE_NONE {
		goPluginType = GoPluginType_GO_PLUGIN_TYPE_GO
	}
	noDefaultModifiers := p.options.GoNoDefaultModifiers
	if goPluginType != GoPluginType_GO_PLUGIN_TYPE_GO {
		noDefaultModifiers = true
	}
	if p.options.NoDefaultIncludes {
		noDefaultModifiers = true
	}
	modifiers := p.getModifiers(protoSpec, relDirPath, noDefaultModifiers)
	goOutOpts := modifiers
	if p.options.Grpc {
		goOutOpts = fmt.Sprintf("%s,plugins=grpc", goOutOpts)
	}
	var flags []string
	if len(goOutOpts) > 0 {
		flags = append(flags, fmt.Sprintf("--%s_out=%s:%s", goPluginType.SimpleString(), goOutOpts, outDirPath))
	} else {
		flags = append(flags, fmt.Sprintf("--%s_out=%s", goPluginType.SimpleString(), outDirPath))
	}
	if p.options.GrpcGateway {
		if len(modifiers) > 0 {
			flags = append(flags, fmt.Sprintf("--grpc-gateway_out=%s:%s", modifiers, outDirPath))
		} else {
			flags = append(flags, fmt.Sprintf("--grpc-gateway_out=%s", outDirPath))
		}
	}
	return flags, nil
}

func (p *goPlugin) getModifiers(protoSpec *protoSpec, curRelDirPath string, noDefaultModifiers bool) string {
	var modifiers map[string]string
	if noDefaultModifiers {
		modifiers = make(map[string]string)
	} else {
		modifiers = copyStringStringMap(defaultGoModifierOptions)
	}
	for key, value := range p.options.GoModifiers {
		modifiers[key] = value
	}
	for relDirPath, files := range protoSpec.RelDirPathToFiles {
		if relDirPath != curRelDirPath {
			importPath := relDirPath
			if p.options.GoImportPath != "" {
				importPath = filepath.Join(p.options.GoImportPath, relDirPath)
			}
			// TODO(pedge)
			if importPath != "" && importPath != "." {
				for _, file := range files {
					modifiers[filepath.Join(relDirPath, file)] = importPath
				}
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
	name          string
	options       *CompileOptions
	relOutDirPath string
}

func newBasePlugin(name string, options *CompileOptions, relOutDirPath string) *basePlugin {
	if options == nil {
		options = &CompileOptions{}
	}
	return &basePlugin{name, options, relOutDirPath}
}

func (p *basePlugin) Flags(protoSpec *protoSpec, relDirPath string, outDirPath string) ([]string, error) {
	if p.relOutDirPath != "" {
		outDirPath = filepath.Join(outDirPath, p.relOutDirPath)
	}
	return []string{fmt.Sprintf("--%s_out=%s", p.name, outDirPath)}, nil
}

type grpcPlugin struct {
	*basePlugin
	grpcName string
}

func newGrpcPlugin(name string, grpcName string, options *CompileOptions, relOutDirPath string) *grpcPlugin {
	if options == nil {
		options = &CompileOptions{}
	}
	return &grpcPlugin{newBasePlugin(name, options, relOutDirPath), grpcName}
}

func (p *grpcPlugin) Flags(protoSpec *protoSpec, relDirPath string, outDirPath string) ([]string, error) {
	args, err := p.basePlugin.Flags(protoSpec, relDirPath, outDirPath)
	if err != nil {
		return nil, err
	}
	if p.options.Grpc {
		if p.relOutDirPath != "" {
			outDirPath = filepath.Join(outDirPath, p.relOutDirPath)
		}
		whichGrpcPlugin, err := which(fmt.Sprintf("grpc_%s_plugin", p.grpcName))
		if err != nil {
			return nil, err
		}
		return append(args, fmt.Sprintf("--grpc_out=%s", outDirPath), fmt.Sprintf("--plugin=protoc-gen-grpc=%s", whichGrpcPlugin)), nil
	}
	return args, nil
}
