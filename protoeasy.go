/*
Package protoeasy is intended to make using protoc simpler.

Protoeasy compiles all protocol buffer files in a directory/subdirectories,
taking care of all include directories, takes care of gRPC compilation,
and take care of package import modifiers for Golang.

Protoeasy also provides a client/server model where compilation is delegated
to a server process meant to be run in a Docker container. This allows you
to not have to install protoc, gRPC, and associated protoc plugins on your
local development machine.

See the README.md file for more details.
*/
package protoeasy // import "go.pedge.io/protoeasy"
import (
	"fmt"
	"strconv"
	"strings"
)

var (
	//DefaultServerCompiler is the default Compiler for a server.
	DefaultServerCompiler = NewServerCompiler(CompilerOptions{})
	// DefaultAPIServer is the default API Server.
	DefaultAPIServer = NewAPIServer(DefaultServerCompiler)
)

// Compiler compiles protocol buffer files.
type Compiler interface {
	// Compile compiles the protocol buffer files in dirPath and outputs the generated
	// files to outDirPath, using the given CompileOptions.
	Compile(dirPath string, outDirPath string, compileOptions *CompileOptions) ([]*Command, error)
}

// CompilerOptions are options for a Compiler.
type CompilerOptions struct{}

// NewServerCompiler returns a new server Compiler.
func NewServerCompiler(options CompilerOptions) Compiler {
	return newServerCompiler(options)
}

// NewAPIServer returns a new APIServer for the given Compiler.
func NewAPIServer(compiler Compiler) APIServer {
	return newAPIServer(compiler)
}

// NewClientCompiler returns a new client Compiler for the given APIClient.
func NewClientCompiler(apiClient APIClient, options CompilerOptions) Compiler {
	return newClientCompiler(apiClient, options)
}

// GoPluginTypeSimpleValueOf returns the GoPluginType for the simple value.
func GoPluginTypeSimpleValueOf(s string) (GoPluginType, error) {
	goPluginTypeObj, ok := GoPluginType_value[fmt.Sprintf("GO_PLUGIN_TYPE_%s", strings.ToUpper(s))]
	if !ok {
		return GoPluginType_GO_PLUGIN_TYPE_NONE, fmt.Errorf("no protoeasy.GoPluginType for %s", s)
	}
	return GoPluginType(goPluginTypeObj), nil
}

// SimpleString returns the simple value for the GoPluginType.
func (x GoPluginType) SimpleString() string {
	s, ok := GoPluginType_name[int32(x)]
	if !ok {
		return strconv.Itoa(int(x))
	}
	return strings.TrimPrefix(strings.ToLower(s), "go_plugin_type_")
}

// AllGoPluginTypeSimpleStrings returns the simple values for all GoPluginTypes.
func AllGoPluginTypeSimpleStrings() []string {
	simpleStrings := make([]string, len(GoPluginType_name)-1)
	for i := range GoPluginType_name {
		if i != 0 {
			simpleStrings[i-1] = ((GoPluginType)(i)).SimpleString()
		}
	}
	return simpleStrings
}
