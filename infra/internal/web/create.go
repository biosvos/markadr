package web

import (
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"log"
	"os"
)

func (r *router) create(ctx *atreugo.RequestCtx) error {
	args := ctx.PostArgs()
	log.Println(args)
	return ctx.RedirectResponse("/pages", 200)
}

func (r *router) createForm(ctx *atreugo.RequestCtx) error {
	file, err := os.ReadFile("assets/html/create.html")
	if err != nil {
		return errors.WithStack(err)
	}
	return ctx.HTTPResponseBytes(file)
}
