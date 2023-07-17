package repository

import "github.com/biosvos/markadr/flow/domain"

type Repository interface {
	Get(title string) (*domain.ADR, error)
	List() ([]*domain.ADR, error)
	Update(record *domain.ADR) error
}
