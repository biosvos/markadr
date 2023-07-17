package watcher

type Watcher interface {
	Start(fn func(filename string)) error
	Stop()
}
