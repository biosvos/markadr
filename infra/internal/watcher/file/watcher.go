package file

import (
	"github.com/biosvos/markadr/flow/watcher"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"log"
)

var _ watcher.Watcher = &Watcher{}

type Watcher struct {
	workDir string
	watcher *fsnotify.Watcher
	stopCh  chan struct{}
}

func NewWatcher(workDir string) (*Watcher, error) {
	wtchr, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Watcher{
		workDir: workDir,
		watcher: wtchr,
		stopCh:  make(chan struct{}),
	}, nil
}

func (w *Watcher) callback() {
	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}
			log.Println(event)
			// TODO broker
		case err, ok := <-w.watcher.Errors:
			if !ok {
				return
			}
			log.Println(err)
		case _ = <-w.stopCh:
			return
		}
	}
}

func (w *Watcher) Start() (chan string, error) {
	ch := make(chan string)
	go w.callback()
	err := w.watcher.Add(w.workDir)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ch, nil
}

func (w *Watcher) Stop() {
	w.stopCh <- struct{}{}
	err := w.watcher.Close()
	if err != nil {
		log.Printf("%+v", err)
	}
}
