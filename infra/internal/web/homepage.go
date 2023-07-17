package web

import (
	"github.com/biosvos/markadr/assets/html"
	"github.com/savsgio/atreugo/v11"
)

func (r *router) homepage(ctx *atreugo.RequestCtx) error {
	return ctx.HTTPResponse(html.Index)
}
