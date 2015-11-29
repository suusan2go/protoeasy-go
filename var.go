package protoeasy

import (
	"errors"
	"path/filepath"
)

var (
	errGoPathNotSet = errors.New("protoeasy: GOPATH not set")

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
