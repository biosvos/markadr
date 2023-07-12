package adr

import (
	"bytes"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/pkg/errors"
	"os"
)

func newDocument(filename string) (ast.Node, error) {
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	ret := p.Parse(body)
	return ret, nil
}

func parseText(doc ast.Node) string {
	var buffer bytes.Buffer
	parseNode(doc, func(node ast.Node) bool {
		switch cur := node.(type) {
		case *ast.Text:
			buffer.WriteString(string(cur.Literal))
			return true
		}
		return false
	})
	return buffer.String()
}

func parseNode(doc ast.Node, fn func(node ast.Node) bool) {
	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}
		ret := fn(node)
		if ret {
			return ast.SkipChildren
		}
		return ast.GoToNext
	})
}

type TOC struct {
	title    string
	children []*TOC
}

func TableOfContents(section *Section) *TOC {
	var ret TOC
	ret.title = section.title
	for _, child := range section.children {
		ret.children = append(ret.children, TableOfContents(child))
	}
	return &ret
}
