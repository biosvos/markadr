package flow

type Navigator struct {
	pager Pager
}

func NewNavigator(pager Pager) *Navigator {
	return &Navigator{pager: pager}
}

func (n *Navigator) ListPages() ([]*Page, error) {
	return n.pager.List()
}
