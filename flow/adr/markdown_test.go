package adr

import (
	"encoding/json"
	"github.com/biosvos/markadr/assets/markdown"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTableOfContents(t *testing.T) {
	document, _ := NewDocument(markdown.TestMarkdownFile)
	section := DivideSection(document)

	toc := TableOfContents(section)

	require.Len(t, toc.Rows, 7)
	require.EqualValues(t, Row{"title", 0}, toc.Rows[0])
	require.EqualValues(t, Row{ContextAndProblemStatement, 1}, toc.Rows[1])
	require.EqualValues(t, Row{DecisionDrivers, 1}, toc.Rows[2])
	require.EqualValues(t, Row{ProsAndCons, 1}, toc.Rows[3])
	require.EqualValues(t, Row{"[option 1]", 2}, toc.Rows[4])
	require.EqualValues(t, Row{DecisionOutcome, 1}, toc.Rows[5])
	require.EqualValues(t, Row{Links, 1}, toc.Rows[6])
}

func TestDivideSections(t *testing.T) {
	document, _ := NewDocument(markdown.TestMarkdownFile)

	sections := DivideSection(document)

	require.Len(t, sections.children, 1)
	require.Len(t, sections.children[0].children, 5)
}

func TestNewDocument(t *testing.T) {
	document, err := NewDocument(markdown.TestMarkdownFile)
	require.NoError(t, err)
	require.NotNil(t, document)
}

func TestJSON(t *testing.T) {
	adr := ADR{
		Title:   "abc",
		Status:  "draft",
		Context: "이랬음",
		Problem: "문제",
		Drivers: []string{"속도"},
		Options: []*Option{
			{
				Title: "캐시",
				Pros:  []string{"빠름"},
				Cons:  []string{"관리 어렵"},
			},
		},
		Outcomes: []*Outcome{
			{
				Title:    "캐시",
				Contents: "빨리졌다.",
			},
		},
		References: []*Reference{
			{
				Title:       "b",
				Destination: "a",
			},
		},
	}

	marshal, err := json.Marshal(&adr)
	require.NoError(t, err)
	t.Log(string(marshal))
}
