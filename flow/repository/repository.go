package repository

import "github.com/biosvos/markadr/flow/adr"

type Repository interface {
	Get(title string) (*adr.ADR, error)
	List() ([]*adr.ADR, error)
}
