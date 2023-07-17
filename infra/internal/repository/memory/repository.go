package memory

import (
	"github.com/biosvos/markadr/flow/domain"
	"github.com/biosvos/markadr/flow/repository"
	"github.com/pkg/errors"
)

var _ repository.Repository = &Repository{}

func NewRepository(items map[string]*domain.ADR) *Repository {
	return &Repository{items: items}
}

type Repository struct {
	items map[string]*domain.ADR
}

func (r *Repository) Update(record *domain.ADR) error {
	r.items[record.Title] = record
	return nil
}

func (r *Repository) Get(title string) (*domain.ADR, error) {
	ret, ok := r.items[title]
	if !ok {
		return nil, errors.New("not found")
	}
	return ret, nil
}

func (r *Repository) List() ([]*domain.ADR, error) {
	var ret []*domain.ADR
	for _, item := range r.items {
		ret = append(ret, item)
	}
	return ret, nil
}
