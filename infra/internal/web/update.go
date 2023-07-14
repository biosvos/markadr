package web

import (
	"encoding/json"
	"github.com/biosvos/markadr/flow/adr"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (r *router) updateADRStatus(ctx *atreugo.RequestCtx) error {
	title := ctx.UserValue("title").(string)
	record, err := r.repository.Get(title)
	if err != nil {
		return errors.WithStack(err)
	}

	var tmp adr.ADR
	err = json.Unmarshal(ctx.PostBody(), &tmp)
	if err != nil {
		return errors.WithStack(err)
	}
	caser := cases.Title(language.English)
	record.Status = adr.Status(caser.String(string(tmp.Status)))

	err = r.repository.Update(record)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
