package filesystem

import (
	"fmt"
	"github.com/biosvos/markadr/infra/internal/humble/hfilesystem"
	"github.com/pkg/errors"
	"os"
)

var (
	_ hfilesystem.HFile      = &Workspace{}
	_ hfilesystem.HDirectory = &Workspace{}
)

func NewWorkspace(core hfilesystem.HFile, workspace string) *Workspace {
	return &Workspace{workspace: workspace, core: core}
}

type Workspace struct {
	workspace string
	core      hfilesystem.HFile
}

func (w *Workspace) ListFiles() ([]string, error) {
	files, err := os.ReadDir(w.workspace)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var ret []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ret = append(ret, file.Name())
	}
	return ret, nil
}

func (w *Workspace) ReadFile(filename string) ([]byte, error) {
	filename = fmt.Sprintf("%v/%v", w.workspace, filename)
	return w.core.ReadFile(filename)
}

func (w *Workspace) WriteFile(filename string, contents []byte) error {
	filename = fmt.Sprintf("%v/%v", w.workspace, filename)
	return w.core.WriteFile(filename, contents)
}
