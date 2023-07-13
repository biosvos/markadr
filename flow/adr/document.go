package adr

import (
	"bytes"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

func NewDocument(contents []byte) (ast.Node, error) {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	ret := p.Parse(contents)
	return ret, nil
}

func parseText(doc ast.Node) string {
	var buffer bytes.Buffer
	parseNode(doc, func(node ast.Node) bool {
		if cur, ok := node.(*ast.Text); ok {
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

type Row struct {
	Title string
	Depth int
}

type TOC struct {
	Rows []Row
}

func TableOfContents(section *Section) *TOC {
	var ret TOC

	type Node struct {
		section *Section
		depth   int
	}
	var stack []*Node

	for i := len(section.children); i > 0; i-- {
		stack = append(stack, &Node{
			section: section.children[i-1],
			depth:   0,
		})
	}

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ret.Rows = append(ret.Rows, Row{
			Title: pop.section.title,
			Depth: pop.depth,
		})
		for i := len(pop.section.children); i > 0; i-- {
			stack = append(stack, &Node{
				section: pop.section.children[i-1],
				depth:   pop.depth + 1,
			})
		}
	}
	return &ret
}
