package adr

import (
	"github.com/biosvos/markadr/flow/broker"
	"github.com/biosvos/markadr/flow/watcher"
	"github.com/pkg/errors"
)

type Service struct {
	broker  broker.Broker
	watcher watcher.Watcher
}

func (s *Service) Start() error {
	_, err := s.watcher.Start()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
