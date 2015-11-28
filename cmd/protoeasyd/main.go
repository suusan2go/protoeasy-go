package main

import (
	"google.golang.org/grpc"

	"go.pedge.io/env"
	"go.pedge.io/proto/server"
	"go.pedge.io/protoeasy"
)

type appEnv struct {
	Port uint16 `env:"PROTOEASY_PORT,default=6789"`
}

func main() {
	env.Main(do, &appEnv{})
}

func do(appEnvObj interface{}) error {
	appEnv := appEnvObj.(*appEnv)
	return protoserver.Serve(
		appEnv.Port,
		func(s *grpc.Server) {
			protoeasy.RegisterAPIServer(
				s,
				protoeasy.NewAPIServer(
					protoeasy.DefaultServerCompiler,
				),
			)
		},
		protoserver.ServeOptions{},
	)
}
