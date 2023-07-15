package hfilesystem

type HFile interface {
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, contents []byte) error
}

type HDirectory interface {
	ListFiles() ([]string, error)
}
