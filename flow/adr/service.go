package adr

import (
	"github.com/biosvos/markadr/flow"
	"github.com/pkg/errors"
)

type Service struct {
	pager flow.Pager
}

func (s *Service) ListADRs() ([]*ADR, error) {
	pages, err := s.pager.List()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, page := range pages {
		_, err := page.Get()
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return nil, nil
}
