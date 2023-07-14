package web

import (
	"bytes"
	"github.com/biosvos/markadr/assets/html"
	"github.com/biosvos/markadr/flow/adr"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"text/template"
)

func (r *router) homepage(ctx *atreugo.RequestCtx) error {
	records, err := r.repository.List()
	if err != nil {
		return errors.WithStack(err)
	}

	var draftSlice []string
	var proposedSlice []string
	var rejectedSlice []string
	var acceptedSlice []string
	var deprecatedSlice []string
	var supersededSlice []string

	for _, record := range records {
		switch record.Status {
		case adr.DraftStatus:
			draftSlice = append(draftSlice, record.Title)
		case adr.ProposedStatus:
			proposedSlice = append(proposedSlice, record.Title)
		case adr.RejectedStatus:
			rejectedSlice = append(rejectedSlice, record.Title)
		case adr.AcceptedStatus:
			acceptedSlice = append(acceptedSlice, record.Title)
		case adr.DeprecatedStatus:
			deprecatedSlice = append(deprecatedSlice, record.Title)
		case adr.SupersededStatus:
			supersededSlice = append(supersededSlice, record.Title)
		}
	}

	tmpl := template.Must(template.New("index").Parse(html.Index))
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, map[string]any{
		"draft":      draftSlice,
		"proposed":   proposedSlice,
		"rejected":   rejectedSlice,
		"accepted":   acceptedSlice,
		"deprecated": deprecatedSlice,
		"superseded": supersededSlice,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return ctx.HTTPResponse(buffer.String())
}
