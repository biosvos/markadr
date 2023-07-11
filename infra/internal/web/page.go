package web

import (
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"log"
)

func (r *router) page(ctx *atreugo.RequestCtx) error {
	title := ctx.UserValue("title").(string)
	log.Println(title)
	page, err := r.navigator.GetPage(title)
	if err != nil {
		return errors.WithStack(err)
	}
	bytes, err := page.Get()
	if err != nil {
		return errors.WithStack(err)
	}
	return ctx.TextResponseBytes(bytes)
}
