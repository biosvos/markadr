package domain

import (
	"encoding/json"
	"github.com/biosvos/markadr/flow/broker"
	"github.com/biosvos/markadr/flow/repository"
	"github.com/biosvos/markadr/flow/watcher"
	"github.com/pkg/errors"
	"log"
	"strings"
)

func NewService(watcher watcher.Watcher, broker broker.Broker) *Service {
	return &Service{watcher: watcher, broker: broker}
}

type Service struct {
	watcher    watcher.Watcher
	broker     broker.Broker
	workspace  string
	extension  string
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

func (s *Service) callback(filename string) {
	filename = strings.TrimLeft(filename, s.workspace)
	filename = strings.TrimLeft(filename, "/")
	filename = strings.TrimRight(filename, s.extension)
	filename = strings.TrimRight(filename, ".")

	event := Event{
		Title:  filename,
		Status: "",
	}
	adr, err := s.repository.Get(filename)
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
