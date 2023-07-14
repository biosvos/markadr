package web

import (
	"bytes"
	myHTML "github.com/biosvos/markadr/assets/html"
	"github.com/biosvos/markadr/flow/adr"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"text/template"
)

func (r *router) page(ctx *atreugo.RequestCtx) error {
	title := ctx.UserValue("title").(string)
	record, err := r.repository.Get(title)
	if err != nil {
		return errors.WithStack(err)
	}

	result, err := makeADRHTML(record)
	if err != nil {
		return errors.WithStack(err)
	}

	navigation, err := makeNavigation(record)
	if err != nil {
		return errors.WithStack(err)
	}

	tmpl := template.Must(template.New("page").Parse(myHTML.Page))
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, map[string]string{
		"contents":   result,
		"navigation": navigation,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.HTTPResponse(buffer.String())
}

func makeADRHTML(record *adr.ADR) (string, error) {
	tmpl := template.Must(template.New("adr").Parse(myHTML.ADR))
	var buffer bytes.Buffer
	err := tmpl.Execute(&buffer, record)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return buffer.String(), nil
}

func makeNavigation(record *adr.ADR) (string, error) {
	tmpl := template.Must(template.New("navigation").Parse(myHTML.Navigation))
	var buffer bytes.Buffer
	err := tmpl.Execute(&buffer, record)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return buffer.String(), nil
}
