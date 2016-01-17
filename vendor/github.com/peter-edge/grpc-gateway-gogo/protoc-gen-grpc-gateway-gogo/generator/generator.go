// Package generator provides an abstract interface to code generators.
package generator

import (
	"github.com/peter-edge/grpc-gateway-gogo/protoc-gen-grpc-gateway-gogo/descriptor"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

// Generator is an abstraction of code generators.
type Generator interface {
	// Generate generates output files from input .proto files.
	Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error)
}
