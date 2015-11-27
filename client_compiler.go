package protoeasy

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/net/context"
)

type clientCompiler struct {
	apiClient APIClient
	options   ClientCompilerOptions
}

func newClientCompiler(
	apiClient APIClient,
	options ClientCompilerOptions,
) *clientCompiler {
	return &clientCompiler{
		apiClient,
		options,
	}
}

func (c *clientCompiler) Compile(dirPath string, outDirPath string, directives *Directives) (retErr error) {
	relFilePaths, err := getAllRelProtoFilePaths(dirPath)
	if err != nil {
		return err
	}
	readCloser, err := tar(dirPath, relFilePaths)
	if err != nil {
		return err
	}
	defer func() {
		if err := readCloser.Close(); err != nil && retErr == nil {
			retErr = err
		}
	}()
	data, err := ioutil.ReadAll(readCloser)
	if err != nil {
		return err
	}
	compileResponse, err := c.apiClient.Compile(
		context.Background(),
		&CompileRequest{
			Tar:        data,
			Directives: directives,
		},
	)
	if err != nil {
		return err
	}
	return untar(bytes.NewReader(compileResponse.Tar), outDirPath)
}
