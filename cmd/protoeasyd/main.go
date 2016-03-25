package main

import (
	"fmt"
	"math"
	"net"

	"google.golang.org/grpc"

	"go.pedge.io/lion/env"
	"go.pedge.io/protoeasy"
)

type appEnv struct {
	Port uint16 `env:"PROTOEASY_PORT,default=6789"`
}

func main() {
	envlion.Main(func(appEnvObj interface{}) error { return do(appEnvObj.(*appEnv)) }, &appEnv{})
}

func do(appEnv *appEnv) error {
	server := grpc.NewServer(grpc.MaxConcurrentStreams(math.MaxUint32))
	protoeasy.RegisterAPIServer(server, protoeasy.DefaultAPIServer)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", appEnv.Port))
	if err != nil {
		return err
	}
	return server.Serve(listener)
}
