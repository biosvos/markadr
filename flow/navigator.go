package flow

type Navigator struct {
	pager Pager
}

func NewNavigator(pager Pager) *Navigator {
	return &Navigator{pager: pager}
}

func (n *Navigator) ListPages() ([]Page, error) {
	return n.pager.List()
}

func (n *Navigator) GetPage(title string) (Page, error) {
	return n.pager.Get(title)
}
