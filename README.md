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
to use protoc, and are sick of having to get everything installed on your development machine, and you want an easier
way to do it.

Then protoeasy is for you!
