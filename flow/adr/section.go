package adr

import "github.com/gomarkdown/markdown/ast"

type Section struct {
	title    string
	outlines []ast.Node
	children []*Section
}

func DivideSection(doc ast.Node) *Section {
	stack := []*Section{
		{
			title: "",
		},
	}
	for _, node := range doc.GetChildren() {
		heading, ok := node.(*ast.Heading)
		if ok {
			stack = stack[:heading.Level]
			chapter := &Section{
				title: parseText(heading),
			}
			stack[len(stack)-1].children = append(stack[len(stack)-1].children, chapter)
			stack = append(stack, chapter)
		} else {
			stack[len(stack)-1].outlines = append(stack[len(stack)-1].outlines, node)
		}
	}
	return stack[0]
}
