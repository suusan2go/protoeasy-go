package protoeasy

import (
	"fmt"
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
	defer func(start time.Time) { logCompile(request, response, retErr, time.Since(start)) }(time.Now())
	if request == nil {
		return nil, newErrNil("request")
	}
	if request.Archive == nil {
		return nil, newErrNil("request.Archive")
	}
	if request.Directives == nil {
		return nil, newErrNil("request.Directives")
	}
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
	if err := untar(request.Archive, dirPath); err != nil {
		return nil, err
	}
	commands, err := a.compiler.Compile(dirPath, outDirPath, request.Directives)
	if err != nil {
		return nil, err
	}
	archive, err := tar(outDirPath, []string{"."})
	if err != nil {
		return nil, err
	}
	return &CompileResponse{
		Command: commands,
		Archive: archive,
	}, nil
}

func logCompile(request *CompileRequest, response *CompileResponse, err error, duration time.Duration) {
	// TODO(pedge): this whole thing needs work, this is just to get it logging as of now
	got := "<nil>"
	with := 0
	ran := 0
	returned := 0
	errString := ""
	if request != nil {
		got = fmt.Sprintf("%v", request.Directives)
		if request.Archive != nil {
			with = len(request.Archive.Value)
		}
	}
	if response != nil {
		ran = len(response.Command)
		if response.Archive != nil {
			returned = len(response.Archive.Value)
		}
	}
	if err != nil {
		errString = fmt.Sprintf(", had error %s", err.Error())
	}
	protolog.Infof("protoeasy.API#Compile: took %v, got %v with %d bytes, ran %d commands, returned %d bytes%s\n", duration, got, with, ran, returned, errString)
}
