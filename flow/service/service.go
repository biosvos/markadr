package service

import (
	"encoding/json"
	"github.com/biosvos/markadr/flow/broker"
	"github.com/biosvos/markadr/flow/repository"
	"github.com/biosvos/markadr/flow/watcher"
	"github.com/pkg/errors"
	"log"
)

func NewService(watcher watcher.Watcher, broker broker.Broker, repository repository.Repository) *Service {
	return &Service{watcher: watcher, broker: broker, repository: repository}
}

type Service struct {
	watcher    watcher.Watcher
	broker     broker.Broker
	repository repository.Repository
}

func (s *Service) Start() error {
	err := s.watcher.Start(s.callback)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

type Event struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func (s *Service) callback(title string) {
	event := Event{
		Title:  title,
		Status: "",
	}
	adr, err := s.repository.Get(title)
	if err == nil {
		event.Status = string(adr.Status)
	}
	s.publishEvent(&event)
}

func (s *Service) publishEvent(event *Event) {
	marshal, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.broker.Publish("adr", string(marshal))
	if err != nil {
		log.Println(err)
		return
	}
	return
}
