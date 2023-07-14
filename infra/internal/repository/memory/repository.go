package memory

import (
	"github.com/biosvos/markadr/flow/adr"
	"github.com/biosvos/markadr/flow/repository"
	"github.com/pkg/errors"
)

var _ repository.Repository = &Repository{}

type Repository struct {
	items map[string]*adr.ADR
}

func (r *Repository) Get(title string) (*adr.ADR, error) {
	ret, ok := r.items[title]
	if !ok {
		return nil, errors.New("not found")
	}
	return ret, nil
}

func (r *Repository) List() ([]*adr.ADR, error) {
	var ret []*adr.ADR
	for _, item := range r.items {
		ret = append(ret, item)
	}
	return ret, nil
}
