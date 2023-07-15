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

func (w *Watcher) callback(fn func(filename string)) {
	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}
			log.Println(event)
			fn(event.Name)
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

func (w *Watcher) Start(fn func(filename string)) error {
	go w.callback(fn)
	err := w.watcher.Add(w.workDir)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (w *Watcher) Stop() {
	w.stopCh <- struct{}{}
	err := w.watcher.Close()
	if err != nil {
		log.Printf("%+v", err)
	}
}
