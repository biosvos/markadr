package flow

type Page struct {
	Name string
}

type Pager interface {
	List() ([]*Page, error)
}
