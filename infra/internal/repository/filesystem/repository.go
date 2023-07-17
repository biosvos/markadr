package filesystem

import (
	"encoding/json"
	"github.com/biosvos/markadr/flow/domain"
	"github.com/biosvos/markadr/flow/repository"
	"github.com/biosvos/markadr/infra/internal/humble/hfilesystem"
	"github.com/pkg/errors"
)

var _ repository.Repository = &Repository{}

func NewRepository(directory hfilesystem.HDirectory, file hfilesystem.HFile) *Repository {
	return &Repository{
		directory: directory,
		file:      file,
	}
}

type Repository struct {
	directory hfilesystem.HDirectory
	file      hfilesystem.HFile
}

func (r *Repository) Update(record *domain.ADR) error {
	_, err := r.get(record.Title)
	if err != nil {
		return errors.WithStack(err)
	}
	marshal, err := json.MarshalIndent(record, "", "\t")
	if err != nil {
		return errors.WithStack(err)
	}
	err = r.file.WriteFile(record.Title, marshal)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *Repository) Get(title string) (*domain.ADR, error) {
	return r.get(title)
}

func (r *Repository) get(title string) (*domain.ADR, error) {
	contents, err := r.file.ReadFile(title)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var ret domain.ADR
	err = json.Unmarshal(contents, &ret)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ret.Title = title
	return &ret, nil
}

func (r *Repository) List() ([]*domain.ADR, error) {
	files, err := r.directory.ListFiles()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var ret []*domain.ADR
	for _, file := range files {
		contents, err := r.file.ReadFile(file)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		var adr domain.ADR
		err = json.Unmarshal(contents, &adr)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		adr.Title = file
		ret = append(ret, &adr)
	}
	return ret, nil
}
