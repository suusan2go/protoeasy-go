package protoeasy

import "path/filepath"

type protoSpecProvider struct {
	options ProtoSpecProviderOptions
}

func newProtoSpecProvider(options ProtoSpecProviderOptions) *protoSpecProvider {
	return &protoSpecProvider{options}
}

func (d *protoSpecProvider) Get(dirPath string, excludeFilePatterns []string) (*ProtoSpec, error) {
	relFilePaths, err := getAllRelProtoFilePaths(dirPath)
	if err != nil {
		return nil, err
	}
	relFilePaths, err = filterFilePaths(relFilePaths, excludeFilePatterns)
	if err != nil {
		return nil, err
	}
	return &ProtoSpec{
		DirPath:           dirPath,
		RelDirPathToFiles: d.getRelDirPathToFiles(relFilePaths),
	}, nil
}

func (d *protoSpecProvider) getRelDirPathToFiles(relFilePaths []string) map[string][]string {
	relDirPathToFiles := make(map[string][]string)
	for _, relFilePath := range relFilePaths {
		relDirPath := filepath.Dir(relFilePath)
		if _, ok := relDirPathToFiles[relDirPath]; !ok {
			relDirPathToFiles[relDirPath] = make([]string, 0)
		}
		relDirPathToFiles[relDirPath] = append(relDirPathToFiles[relDirPath], filepath.Base(relFilePath))
	}
	return relDirPathToFiles
}
