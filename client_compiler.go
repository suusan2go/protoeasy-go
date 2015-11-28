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

func (c *clientCompiler) Compile(dirPath string, outDirPath string, directives *Directives) (retVal [][]string, retErr error) {
	relFilePaths, err := getAllRelProtoFilePaths(dirPath)
	if err != nil {
		return nil, err
	}
	readCloser, err := tar(dirPath, relFilePaths)
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
	compileResponse, err := c.apiClient.Compile(
		context.Background(),
		&CompileRequest{
			Tar:        data,
			Directives: directives,
		},
	)
	if err != nil {
		return nil, err
	}
	argsList := make([][]string, len(compileResponse.Args))
	for i, args := range compileResponse.Args {
		argsList[i] = args.Value
	}
	for _, args := range argsList {
		logArgs(args)
	}
	if err := untar(bytes.NewReader(compileResponse.Tar), outDirPath); err != nil {
		return nil, err
	}
	return argsList, nil
}
