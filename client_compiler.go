package protoeasy

import "golang.org/x/net/context"

type clientCompiler struct {
	apiClient APIClient
	options   CompilerOptions
}

func newClientCompiler(apiClient APIClient, options CompilerOptions) *clientCompiler {
	return &clientCompiler{apiClient, options}
}

func (c *clientCompiler) Compile(dirPath string, outDirPath string, directives *Directives) (retVal []*Command, retErr error) {
	if directives == nil {
		return nil, newErrNil("directives")
	}
	relFilePaths, err := getAllRelProtoFilePaths(dirPath)
	if err != nil {
		return nil, err
	}
	archive, err := tar(dirPath, relFilePaths)
	if err != nil {
		return nil, err
	}
	compileResponse, err := c.apiClient.Compile(
		context.Background(),
		&CompileRequest{
			Archive:    archive,
			Directives: directives,
		},
	)
	if err != nil {
		return nil, err
	}
	if compileResponse == nil {
		return nil, newErrNil("compileResponse")
	}
	if compileResponse.Archive == nil {
		return nil, newErrNil("compileResponse.Archive")
	}
	for _, command := range compileResponse.Command {
		logCommand(command)
	}
	if err := untar(compileResponse.Archive, outDirPath); err != nil {
		return nil, err
	}
	return compileResponse.Command, nil
}
