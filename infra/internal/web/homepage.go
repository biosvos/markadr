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
	summaries, err := r.repository.ListSummaries()
	if err != nil {
		return errors.WithStack(err)
	}

	var draftSlice []string
	var proposedSlice []string
	var rejectedSlice []string
	var acceptedSlice []string
	var deprecatedSlice []string
	var supersededSlice []string

	for _, summary := range summaries {
		switch summary.Status {
		case adr.DraftStatus:
			draftSlice = append(draftSlice, summary.Title)
		case adr.ProposedStatus:
			proposedSlice = append(proposedSlice, summary.Title)
		case adr.RejectedStatus:
			rejectedSlice = append(rejectedSlice, summary.Title)
		case adr.AcceptedStatus:
			acceptedSlice = append(acceptedSlice, summary.Title)
		case adr.DeprecatedStatus:
			deprecatedSlice = append(deprecatedSlice, summary.Title)
		case adr.SupersededStatus:
			supersededSlice = append(supersededSlice, summary.Title)
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
