package web

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
)

func (r *router) homepage(ctx *atreugo.RequestCtx) error {
	pages, err := r.navigator.ListPages()
	if err != nil {
		return errors.WithStack(err)
	}
	var buffer bytes.Buffer
	for _, page := range pages {
		buffer.WriteString(page.Name)
		buffer.WriteString("\n")
	}
	return ctx.TextResponse(buffer.String())
}
