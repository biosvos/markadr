package filesystem

import (
	"encoding/json"
	"fmt"
	"github.com/biosvos/markadr/flow/adr"
	"github.com/biosvos/markadr/flow/repository"
	"github.com/pkg/errors"
	"os"
	"strings"
)

var _ repository.Repository = &Repository{}

func NewRepository(workDir string) *Repository {
	return &Repository{workDir: workDir}
}

type Repository struct {
	workDir string
}

func (r *Repository) Update(record *adr.ADR) error {
	_, err := r.get(record.Title)
	if err != nil {
		return errors.WithStack(err)
	}
	marshal, err := json.Marshal(record)
	if err != nil {
		return errors.WithStack(err)
	}
	err = os.WriteFile(fmt.Sprintf("%v/%v.json", r.workDir, record.Title), marshal, 0)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *Repository) Get(title string) (*adr.ADR, error) {
	return r.get(title)
}

func (r *Repository) get(title string) (*adr.ADR, error) {
	contents, err := os.ReadFile(fmt.Sprintf("%v/%v.json", r.workDir, title))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var ret adr.ADR
	err = json.Unmarshal(contents, &ret)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ret.Title = title
	return &ret, nil
}

func (r *Repository) List() ([]*adr.ADR, error) {
	files, err := os.ReadDir(r.workDir)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var ret []*adr.ADR
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		contents, err := os.ReadFile(fmt.Sprintf("%v/%v", r.workDir, file.Name()))
		if err != nil {
			return nil, errors.WithStack(err)
		}

		var item adr.ADR
		err = json.Unmarshal(contents, &item)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		item.Title = strings.TrimSuffix(file.Name(), ".json")
		ret = append(ret, &item)
	}
	return ret, nil
}
