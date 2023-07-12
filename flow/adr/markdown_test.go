package adr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTableOfContents(t *testing.T) {
	document, _ := newDocument("../../assets/markdown/template.md")
	section := divideSection(document)

	toc := TableOfContents(section)

	require.Len(t, toc.children, 1)
	require.EqualValues(t, "title", toc.children[0].title)
	require.Len(t, toc.children[0].children, 5)
	require.EqualValues(t, ContextAndProblemStatement, toc.children[0].children[0].title)
	require.EqualValues(t, DecisionDrivers, toc.children[0].children[1].title)
	require.EqualValues(t, ProsAndCons, toc.children[0].children[2].title)
	require.EqualValues(t, DecisionOutcome, toc.children[0].children[3].title)
	require.EqualValues(t, Links, toc.children[0].children[4].title)
}

func TestDivideSections(t *testing.T) {
	document, _ := newDocument("../../assets/markdown/template.md")

	sections := divideSection(document)

	require.Len(t, sections.children, 1)
	require.Len(t, sections.children[0].children, 5)
}

func TestNewDocument(t *testing.T) {
	document, err := newDocument("../../assets/markdown/template.md")
	require.NoError(t, err)
	require.NotNil(t, document)
}

func TestNewADR(t *testing.T) {
	document, _ := newDocument("../../assets/markdown/template.md")
	sections := divideSection(document)

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
