package filesystem

import (
	"github.com/biosvos/markadr/infra/internal/humble/hfilesystem"
	"github.com/pkg/errors"
	"os"
)

var _ hfilesystem.HFile = &Filesystem{}

func NewFilesystem() *Filesystem {
	return &Filesystem{}
}

type Filesystem struct {
}

func (f *Filesystem) ReadFile(filename string) ([]byte, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return contents, nil
}

func (f *Filesystem) WriteFile(filename string, contents []byte) error {
	err := os.WriteFile(filename, contents, 0600)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
