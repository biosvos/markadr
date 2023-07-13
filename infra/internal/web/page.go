package web

import (
	"bytes"
	"fmt"
	myHTML "github.com/biosvos/markadr/assets/html"
	"github.com/biosvos/markadr/flow/adr"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"io"
	"strings"
	"text/template"
)

type CustomHook struct {
	prevLevel int
	numbering int
}

func (c *CustomHook) renderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	heading, ok := node.(*ast.Heading)
	if !(entering && ok) {
		return ast.GoToNext, false
	}

	switch {
	case c.prevLevel == heading.Level:
		_, _ = w.Write([]byte("</section>"))
	case c.prevLevel > heading.Level:
		diff := c.prevLevel - heading.Level
		for i := 0; i < diff; i++ {
			_, _ = w.Write([]byte("</section>"))
		}
		_, _ = w.Write([]byte("</section>"))
	case c.prevLevel < heading.Level:
	}

	_, _ = w.Write([]byte(fmt.Sprintf("<section id='%v'>", c.numbering)))
	c.numbering++

	c.prevLevel = heading.Level
	return ast.GoToNext, false
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	hook := &CustomHook{}
	opts := html.RendererOptions{
		Flags:          htmlFlags,
		RenderNodeHook: hook.renderHook,
	}
	renderer := html.NewRenderer(opts)

	contents := markdown.Render(doc, renderer)
	for i := 0; i < hook.prevLevel; i++ {
		contents = fmt.Appendf(contents, "</section>")
	}
	return contents
}

func (r *router) page(ctx *atreugo.RequestCtx) error {
	title := ctx.UserValue("title").(string)
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

	tmpl := template.Must(template.New("page").Parse(myHTML.Page))
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
