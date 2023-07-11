package web

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"log"
)

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

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
	ret := mdToHTML(bytes)
	return ctx.HTTPResponseBytes(ret)
}
