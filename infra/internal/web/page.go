package web

import (
	"bytes"
	"fmt"
	"github.com/biosvos/markadr/flow/adr"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"log"
	"strings"
	"text/template"
)

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags: htmlFlags,
	}
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
	contents, err := page.Get()
	if err != nil {
		return errors.WithStack(err)
	}
	ret := mdToHTML(contents)

	document, err := adr.NewDocument(contents)
	if err != nil {
		return errors.WithStack(err)
	}

	section := adr.DivideSection(document)
	toc := adr.TableOfContents(section)

	navigator := newNavigator(toc)
	if err != nil {
		return errors.WithStack(err)
	}

	tmpl, err := template.ParseFiles("assets/html/page.html")
	if err != nil {
		return errors.WithStack(err)
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, map[string]string{
		"contents":   string(ret),
		"navigation": navigator,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.HTTPResponse(buffer.String())
}

func newNavigator(toc *adr.TOC) string {
	if len(toc.Rows) == 0 {
		return ""
	}

	const defaultFormat = "<li><a href='#%v'>%v</a>INSIDE</li>NEXT"
	const insideFormat = "<ul><li><a href='#%v'>%v</a>INSIDE</li>NEXT</ul>"
	var ret = fmt.Sprintf(defaultFormat, 0, toc.Rows[0].Title)
	for prev, row := range toc.Rows[1:] {
		switch {
		case toc.Rows[prev].Depth == row.Depth:
			ret = strings.Replace(ret, "INSIDE", "", 1)
			ret = strings.Replace(ret, "NEXT", fmt.Sprintf(defaultFormat, prev+1, row.Title), 1)
		case toc.Rows[prev].Depth < row.Depth:
			ret = strings.Replace(ret, "INSIDE", fmt.Sprintf(insideFormat, prev+1, row.Title), 1)
		case toc.Rows[prev].Depth > row.Depth:
			diff := toc.Rows[prev].Depth - row.Depth
			ret = strings.Replace(ret, "NEXT", "", diff)
			ret = strings.Replace(ret, "NEXT", fmt.Sprintf(defaultFormat, prev+1, row.Title), 1)
		}
	}
	ret = strings.ReplaceAll(ret, "NEXT", "")
	ret = strings.ReplaceAll(ret, "INSIDE", "")
	return ret
}
