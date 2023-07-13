package adr

import (
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

func TestNewADR(t *testing.T) {
	document, _ := NewDocument(markdown.TestMarkdownFile)
	sections := DivideSection(document)

	ret, err := NewADR(sections)

	require.NoError(t, err)
	require.NotNil(t, ret)
	require.EqualValues(t, "title", ret.Title)
	require.EqualValues(t, DraftStatus, ret.Status)
	require.Len(t, ret.DecisionDrivers, 2)
	require.EqualValues(t, "driver1", ret.DecisionDrivers[0])
	require.EqualValues(t, "driver2", ret.DecisionDrivers[1])
	require.Len(t, ret.Options.Options, 1)
	require.EqualValues(t, 1, ret.Options.Pick)
	require.Len(t, ret.Options.Options[0].TradeOff.Pros, 2)
	require.Len(t, ret.Options.Options[0].TradeOff.Cons, 3)
	require.Len(t, ret.Links, 1)
	require.NotEmpty(t, ret.ContextAndProblem)
}
