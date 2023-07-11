package filesystem

import (
	"github.com/biosvos/markadr/flow"
	"github.com/pkg/errors"
	"os"
)

var _ flow.Pager = &Filesystem{}

type Filesystem struct {
	assetPath string
}

func NewFilesystem(assetPath string) *Filesystem {
	return &Filesystem{
		assetPath: assetPath,
	}
}

func (f *Filesystem) List() ([]*flow.Page, error) {
	files, err := os.ReadDir(f.assetPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var ret []*flow.Page
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ret = append(ret, &flow.Page{Name: file.Name()})
	}
	return ret, nil
}
