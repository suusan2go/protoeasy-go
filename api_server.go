package protoeasy

import (
	"bytes"
	"io/ioutil"
	"os"
	"time"

	"go.pedge.io/proto/rpclog"
	"golang.org/x/net/context"
)

type apiServer struct {
	protorpclog.Logger
	compiler Compiler
}

func newAPIServer(compiler Compiler) *apiServer {
	return &apiServer{protorpclog.NewLogger("protoeasy.API"), compiler}
}

func (a *apiServer) Compile(ctx context.Context, request *CompileRequest) (response *CompileResponse, retErr error) {
	defer func(start time.Time) { a.Log(request.Directives, nil, retErr, time.Since(start)) }(time.Now())
	dirPath, err := ioutil.TempDir("", "")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := os.RemoveAll(dirPath); err != nil && retErr == nil {
			retErr = err
		}
	}()
	outDirPath, err := ioutil.TempDir("", "")
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
	if err := a.compiler.Compile(dirPath, outDirPath, request.Directives); err != nil {
		return nil, err
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
		Tar: data,
	}, nil
}
