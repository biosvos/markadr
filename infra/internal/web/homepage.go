package web

import (
	"bytes"
	"github.com/biosvos/markadr/assets/html"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"text/template"
)

func (r *router) homepage(ctx *atreugo.RequestCtx) error {
	pages, err := r.navigator.ListPages()
	if err != nil {
		return errors.WithStack(err)
	}
	var stringPages []string
	for _, page := range pages {
		stringPages = append(stringPages, page.Title())
	}
	tmpl := template.Must(template.New("index").Parse(html.Index))
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, map[string]any{
		"pages": stringPages,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return ctx.HTTPResponse(buffer.String())
}
