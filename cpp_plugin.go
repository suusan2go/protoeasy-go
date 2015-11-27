package protoeasy

import "fmt"

type cppPlugin struct {
	options CppPluginOptions
}

func newCppPlugin(options CppPluginOptions) *cppPlugin {
	return &cppPlugin{options}
}

func (p *cppPlugin) Args(protoSpec *ProtoSpec, relDirPath string, outDirPath string) ([]string, error) {
	args := []string{fmt.Sprintf("--cpp_out=%s", outDirPath)}
	if p.options.Grpc {
		args = append(args, fmt.Sprintf("--grpc_cpp_out=%s", outDirPath))
	}
	return args, nil
}
