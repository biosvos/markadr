package filesystem

import (
	"fmt"
	"github.com/biosvos/markadr/flow"
	"github.com/pkg/errors"
	"os"
	"strings"
)

var _ flow.Page = &File{}

type File struct {
	title string
	path  string
}

func NewFile(title string, path string) *File {
	return &File{
		title: title,
		path:  path,
	}
}

func (f *File) Title() string {
	return f.title
}

func (f *File) Get() ([]byte, error) {
	ret, err := os.ReadFile(f.path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ret, nil
}

var _ flow.Pager = &Filesystem{}

type Filesystem struct {
	assetPath string
}

func NewFilesystem(assetPath string) *Filesystem {
	return &Filesystem{
		assetPath: assetPath,
	}
}

func (f *Filesystem) List() ([]flow.Page, error) {
	files, err := os.ReadDir(f.assetPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var ret []flow.Page
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".md") {
			continue
		}
		title := strings.TrimSuffix(file.Name(), ".md")
		ret = append(ret, NewFile(title, fmt.Sprintf("%v/%v", f.assetPath, file.Name())))
	}
	return ret, nil
}

func (f *Filesystem) Get(title string) (flow.Page, error) {
	file := NewFile(title, fmt.Sprintf("%v/%v.md", f.assetPath, title))
	return file, nil
}
