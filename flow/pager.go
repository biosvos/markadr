package flow

type Page interface {
	Title() string
	Get() ([]byte, error)
}

type Pager interface {
	List() ([]Page, error)
}
