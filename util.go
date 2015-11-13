package protoeasy

import (
	"fmt"
	"path/filepath"
)

var (
	googleProtobufFiles = []string{
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
	}

	googleProtobufGoModifierOptions = newGoModifierOptions(
		"google/protobuf",
		googleProtobufFiles,
		"go.pedge.io/google-protobuf",
	)
)

type keyValue struct {
	key   string
	value string
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
