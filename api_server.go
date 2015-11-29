package protoeasy

import (
	"bytes"
	"io/ioutil"
	"os"
	"time"

	"go.pedge.io/protolog"
	"golang.org/x/net/context"
)

type apiServer struct {
	compiler Compiler
}

func newAPIServer(compiler Compiler) *apiServer {
	return &apiServer{compiler}
}

func (a *apiServer) Compile(ctx context.Context, request *CompileRequest) (response *CompileResponse, retErr error) {
	defer func(start time.Time) {
		protolog.Infof("compile: took %v, got %v with %d bytes, ran %d commands, returned %d bytes\n", time.Since(start), request.Directives, len(request.Tar), len(response.Args), len(response.Tar))
	}(time.Now())
	dirPath, err := ioutil.TempDir("", "protoeasy-input")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := os.RemoveAll(dirPath); err != nil && retErr == nil {
			retErr = err
		}
	}()
	outDirPath, err := ioutil.TempDir("", "protoeasy-output")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := os.RemoveAll(outDirPath); err != nil && retErr == nil {
			retErr = err
		}
	}()
	if err := untar(bytes.NewReader(request.Tar), dirPath); err != nil {
		return nil, err
	}
	argsList, err := a.compiler.Compile(dirPath, outDirPath, request.Directives)
	if err != nil {
		return nil, err
	}
	protoArgs := make([]*Args, len(argsList))
	for i, args := range argsList {
		protoArgs[i] = &Args{
			Value: args,
		}
	}
	readCloser, err := tar(outDirPath, []string{"."})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := readCloser.Close(); err != nil && retErr == nil {
			retErr = err
		}
	}()
	data, err := ioutil.ReadAll(readCloser)
	if err != nil {
		return nil, err
	}
	return &CompileResponse{
		Args: protoArgs,
		Tar:  data,
	}, nil
}
