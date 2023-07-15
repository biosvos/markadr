package filesystem

import (
	"fmt"
	"github.com/biosvos/markadr/infra/internal/humble/hfilesystem"
	"github.com/pkg/errors"
	"strings"
)

var (
	_ hfilesystem.HFile      = &Extension{}
	_ hfilesystem.HDirectory = &Extension{}
)

func NewExtension(directory hfilesystem.HDirectory, file hfilesystem.HFile, extension string) *Extension {
	return &Extension{
		directory: directory,
		file:      file,
		extension: extension,
	}
}

type Extension struct {
	directory hfilesystem.HDirectory
	file      hfilesystem.HFile
	extension string
}

func (e *Extension) ListFiles() ([]string, error) {
	files, err := e.directory.ListFiles()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var ret []string
	for _, file := range files {
		if !strings.HasSuffix(file, fmt.Sprintf(".%v", e.extension)) {
			continue
		}
		file = strings.TrimSuffix(file, fmt.Sprintf(".%v", e.extension))
		ret = append(ret, file)
	}
	return ret, nil
}

func (e *Extension) ReadFile(filename string) ([]byte, error) {
	filename = fmt.Sprintf("%v.%v", filename, e.extension)
	return e.file.ReadFile(filename)
}

func (e *Extension) WriteFile(filename string, contents []byte) error {
	filename = fmt.Sprintf("%v.%v", filename, e.extension)
	return e.file.WriteFile(filename, contents)
}
