package protoeasy

import (
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
	defer func(start time.Time) {
		var compileOptions *CompileOptions
		if request != nil {
			compileOptions = request.CompileOptions
		}
		var compileInfo *CompileInfo
		if response != nil {
			compileInfo = response.CompileInfo
		}
		a.Log(compileOptions, compileInfo, retErr, time.Since(start))
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
	if err := untar(request.Tar, dirPath); err != nil {
		return nil, err
	}
	commands, err := a.compiler.Compile(dirPath, outDirPath, request.CompileOptions)
	if err != nil {
		return nil, err
	}
	tar, err := tar(outDirPath, []string{"."})
	if err != nil {
		return nil, err
	}
	return &CompileResponse{
		Tar: tar,
		CompileInfo: &CompileInfo{
			Command: commands,
		},
	}, nil
}
