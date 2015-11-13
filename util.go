package protoeasy

import (
	"fmt"
	"path/filepath"
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
)

func getGoPathIncludes(goPath string, relativeIncludes []string) []string {
	includes := make([]string, len(relativeIncludes))
	for i, relativeInclude := range relativeIncludes {
		includes[i] = filepath.Join(goPath, relativeInclude)
	}
	return includes
}

func newGoModifierOptions(dir string, files []string, goPackage string) map[string]string {
	m := make(map[string]string)
	for _, file := range files {
		key, value := newGoModifierOption(filepath.Join(dir, file), goPackage)
		m[key] = value
	}
	return m
}

func newGoModifierOption(file string, goPackage string) (string, string) {
	return fmt.Sprintf("M%s", file), goPackage
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
