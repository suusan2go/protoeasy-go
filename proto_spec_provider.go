package protoeasy

import (
	"os"
	"path/filepath"
	"strings"
)

type protoSpecProvider struct {
	options ProtoSpecProviderOptions
}

func newProtoSpecProvider(options ProtoSpecProviderOptions) *protoSpecProvider {
	return &protoSpecProvider{options}
}

func (d *protoSpecProvider) Get(dirPath string, excludeFilePatterns []string) (*ProtoSpec, error) {
	relFilePaths, err := d.getAllRelProtoFilePaths(dirPath, excludeFilePatterns)
	if err != nil {
		return nil, err
	}
	return &ProtoSpec{
		DirPath:           dirPath,
		RelDirPathToFiles: d.getRelDirPathToFiles(relFilePaths),
	}, nil
}

func (d *protoSpecProvider) getAllRelProtoFilePaths(dirPath string, excludeFilePatterns []string) ([]string, error) {
	var relProtoFilePaths []string
	if err := filepath.Walk(
		dirPath,
		func(filePath string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fileInfo.IsDir() {
				return nil
			}
			if filepath.Ext(filePath) != ".proto" {
				return nil
			}
			relFilePath, err := filepath.Rel(dirPath, filePath)
			if err != nil {
				return err
			}
			// TODO(pedge): handle this logic correctly
			for _, excludeFilePattern := range excludeFilePatterns {
				if strings.HasPrefix(relFilePath, excludeFilePattern) {
					return nil
				}
				matched, err := filepath.Match(excludeFilePattern, relFilePath)
				if err != nil {
					return err
				}
				if matched {
					return nil
				}
			}
			relProtoFilePaths = append(relProtoFilePaths, relFilePath)
			return nil
		},
	); err != nil {
		return nil, err
	}
	return relProtoFilePaths, nil
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
