package web

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"html/template"
)

func (r *router) homepage(ctx *atreugo.RequestCtx) error {
	pages, err := r.navigator.ListPages()
	if err != nil {
		return errors.WithStack(err)
	}
	var stringPages []string
	for _, page := range pages {
		stringPages = append(stringPages, page.Name)
	}
	tmpl, err := template.ParseFiles("assets/html/index.html")
	if err != nil {
		return errors.WithStack(err)
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, map[string]any{
		"pages": stringPages,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return ctx.HTTPResponse(buffer.String())
}
