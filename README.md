[![CircleCI](https://circleci.com/gh/peter-edge/go-protoeasy/tree/master.png)](https://circleci.com/gh/peter-edge/go-protoeasy/tree/master)
[![Docker Repository on Quay](https://quay.io/repository/pedge/protoeasy/status)](https://quay.io/repository/pedge/protoeasy)
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/go.pedge.io/protoeasy)
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/peter-edge/go-protoeasy/blob/master/LICENSE)

Protoeasy is intended to make compiling protocol buffers files easier, and to offload the compilation
to a server for consistency and so that protoc and any associated plugins do not have to be installed locally.

### Motivation

Have you ever had one of the following issues:

* You're the resident protobuf expert on your team, and every time there is a change or you get a new teammate,
you already know what the next hour of your life will be.
* You move to a new development machine, and can't remember how you got protoc installed properly, and now
your old build scripts do not work anymore.
* You upgrade proto3 on your computer, and suddenly protoc is broken.
* One of your teammates upgrades proto3 on their computer, and suddently protoc is broken.
* You did `go get -u google.golang.org/grpc`, or even worse, `google.golang.org` is a dependency
of some package and is updated, and suddenly new, incompatible `.pb.go` files are created.
* You have gRPC working on your mac, but linux is a mess.
* You use [Mosh](https://mosh.mit.edu/) and proto3, but since Mosh is still on proto2, you're in trouble.
* You're sick of resolving all relative paths to absolute paths or vice versa for `-I`.
* You never can figure out a good scheme for a large amount of protocol buffers files and how to do imports properly.
* You think this protobuf thing is really cool, but are sick of maintaining large bash scripts or Makefile directives
to use protoc, and are sick of having to get everything installed on your development machine, and you know there must
be an easier way to do it but haven't found it yet.

Then protoeasy is for you!

### Structure

Protoeasy compiles entire directories of protocol buffer files, as opposed to individual files. To use protoeasy:

* All protocol buffers files in the same sub-directory must have the same protocol buffers package.
* All imports must be relative to the base directory.

Both of these are already good practices. Assuming this structure, all you have to do is pass in a base directory,
and the rest is done for you. All `-I` directives are automatically figured out, and all protocol buffers files
in the same sub-directory are compiled together.

Instead of specifying individual output directives, protoeasy breaks compilation into separate languages, optionally
doing gRPC compilation for each language. Assume we have a single file in our current directory called `foo.proto`.
To compile for ruby, we would do:

```
protoeasy --ruby .
# protoc \
#  -I/tmp/protoeasy-input353662656 \
#  -I/go/src/go.pedge.io/protoeasy/vendor/go.pedge.io/google-protobuf \
#  -I/go/src/go.pedge.io/protoeasy/vendor/github.com/golang/protobuf/protoc-gen-go/descriptor \
#  -I/go/src/go.pedge.io/protoeasy/vendor/github.com/gengo/grpc-gateway/third_party/googleapis \
#  -I/go/src/go.pedge.io/protoeasy/vendor/go.pedge.io/googleapis \
#  --ruby_out=/tmp/protoeasy-output452425503 \
#  /tmp/protoeasy-input353662656/foo.proto
```

To compile for ruby and c++, we would do:

```
protoeasy --ruby --cpp .
#protoc \
#  -I/tmp/protoeasy-input138072562 \
#  -I/go/src/go.pedge.io/protoeasy/vendor/go.pedge.io/google-protobuf \
#  -I/go/src/go.pedge.io/protoeasy/vendor/github.com/golang/protobuf/protoc-gen-go/descriptor \
#  -I/go/src/go.pedge.io/protoeasy/vendor/github.com/gengo/grpc-gateway/third_party/googleapis \
#  -I/go/src/go.pedge.io/protoeasy/vendor/go.pedge.io/googleapis \
#  --cpp_out=/tmp/protoeasy-output084691625 \
#  /tmp/protoeasy-input138072562/foo.proto
#protoc \
#  -I/tmp/protoeasy-input138072562 \
#  -I/go/src/go.pedge.io/protoeasy/vendor/go.pedge.io/google-protobuf \
#  -I/go/src/go.pedge.io/protoeasy/vendor/github.com/golang/protobuf/protoc-gen-go/descriptor \
#  -I/go/src/go.pedge.io/protoeasy/vendor/github.com/gengo/grpc-gateway/third_party/googleapis \
#  -I/go/src/go.pedge.io/protoeasy/vendor/go.pedge.io/googleapis \
#  --ruby_out=/tmp/protoeasy-output084691625 \
#  /tmp/protoeasy-input138072562/foo.proto
```

If we wanted to generate gRPC code for both, we would do:

```
protoeasy --ruby --cpp --grpc .
#protoc \
#  -I/tmp/protoeasy-input924418036 \
#  -I/go/src/go.pedge.io/protoeasy/vendor/go.pedge.io/google-protobuf \
#  -I/go/src/go.pedge.io/protoeasy/vendor/github.com/golang/protobuf/protoc-gen-go/descriptor \
#  -I/go/src/go.pedge.io/protoeasy/vendor/github.com/gengo/grpc-gateway/third_party/googleapis \
#  -I/go/src/go.pedge.io/protoeasy/vendor/go.pedge.io/googleapis \
#  --cpp_out=/tmp/protoeasy-output335846083 \
#  --grpc_out=/tmp/protoeasy-output335846083 \
#  --plugin=protoc-gen-grpc=/usr/local/bin/grpc_cpp_plugin \
#  /tmp/protoeasy-input924418036/foo.proto
#protoc \
#  -I/tmp/protoeasy-input924418036 \
#  -I/go/src/go.pedge.io/protoeasy/vendor/go.pedge.io/google-protobuf \
#  -I/go/src/go.pedge.io/protoeasy/vendor/github.com/golang/protobuf/protoc-gen-go/descriptor \
#  -I/go/src/go.pedge.io/protoeasy/vendor/github.com/gengo/grpc-gateway/third_party/googleapis \
#  -I/go/src/go.pedge.io/protoeasy/vendor/go.pedge.io/googleapis \
#  --ruby_out=/tmp/protoeasy-output335846083 \
#  --grpc_out=/tmp/protoeasy-output335846083 \
#  --plugin=protoc-gen-grpc=/usr/local/bin/grpc_ruby_plugin \
#  /tmp/protoeasy-input924418036/foo.proto
```
