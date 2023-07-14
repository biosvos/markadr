package watcher

type Watcher interface {
	Start() (chan string, error)
	Stop()
}
