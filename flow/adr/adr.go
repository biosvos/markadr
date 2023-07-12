package adr

import (
	"bytes"
	"github.com/gomarkdown/markdown/ast"
	"github.com/pkg/errors"
	"strings"
)

type Status string

const (
	DraftStatus      = Status("Draft")
	ProposedStatus   = Status("Proposed")
	RejectedStatus   = Status("Rejected")
	AcceptedStatus   = Status("Accepted")
	DeprecatedStatus = Status("Deprecated")
	SupersededStatus = Status("Superseded")
)

type TradeOff struct {
	Pros []string
	Cons []string
}

type Option struct {
	Title    string
	TradeOff TradeOff
}

type Options struct {
	Options []Option
	Pick    int // 0은 선택하지 않음, 혹은 잘못 선택함
}

type Link struct {
	Title       string
	Destination string
}

type ADR struct {
	Title             string
	Status            Status
	ContextAndProblem string
	DecisionDrivers   []string
	Options           Options
	Links             []Link
}

const (
	ContextAndProblemStatement = "Context and Problem Statement"
	DecisionDrivers            = "Decision Drivers"
	DecisionOutcome            = "Decision Outcome"
	ProsAndCons                = "Options"
	Links                      = "Links"
)

func NewADR(section *Section) (*ADR, error) {
	statusString, err := parseStatus(section.children[0].outlines)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var drivers []string
	var options *Options
	var pick string
	var links []Link
	var context string
	for _, child := range section.children[0].children {
		switch child.title {
		case ContextAndProblemStatement:
			context = parseLiterals(child.outlines)
		case DecisionDrivers:
			drivers = parseDecisionDrivers(child.outlines)
		case DecisionOutcome:
			pick = ParsePick(child.outlines)
		case ProsAndCons:
			options = parseOptions(child.children)
		case Links:
			links = parseLinks(child.outlines)
		default:
		}
	}
	for i, option := range options.Options {
		if option.Title == pick {
			options.Pick = i + 1
		}
	}
	return &ADR{
		Title:             section.children[0].title,
		Status:            decisionStatus(statusString),
		ContextAndProblem: context,
		DecisionDrivers:   drivers,
		Options:           Deref(options),
		Links:             links,
	}, nil
}

func parseLiterals(outlines []ast.Node) string {
	var buffer bytes.Buffer
	for _, outline := range outlines {
		parseNode(outline, func(node ast.Node) bool {
			leaf := node.AsLeaf()
			if leaf != nil {
				buffer.Write(leaf.Literal)
				buffer.WriteString("\n")
				return true
			}
			return false
		})
	}
	return buffer.String()
}

func ParsePick(outlines []ast.Node) string {
	children := getSpecificTypesChildren[*ast.Paragraph](outlines)
	for _, child := range children {
		contents := parseText(child)
		if strings.Contains(contents, "Pick:") {
			contents = strings.ReplaceAll(contents, "Pick:", "")
			return strings.TrimSpace(contents)
		}
	}
	return ""
}

func Deref[T any](item *T) T {
	var ret T
	if item == nil {
		return ret
	}
	ret = *item
	return ret
}

func parseOptions(children []*Section) *Options {
	options := Options{
		Options: nil,
		Pick:    0,
	}
	for _, child := range children {
		var option Option
		option.Title = child.title
		tables := getSpecificTypesChildren[*ast.Table](child.outlines)
		tableBodies := getSpecificTypesChildren[*ast.TableBody](tables)
		tableRows := getSpecificTypesChildren[*ast.TableRow](tableBodies)
		for idx, row := range tableRows {
			text := parseText(row)
			if text == "" {
				continue
			}
			if idx%2 == 0 {
				option.TradeOff.Pros = append(option.TradeOff.Pros, text)
			} else {
				option.TradeOff.Cons = append(option.TradeOff.Cons, text)
			}
		}
		options.Options = append(options.Options, option)
	}
	return &options
}

func parseDecisionDrivers(items []ast.Node) []string {
	children := getSpecificTypesChildren[*ast.List](items)
	children = getSpecificTypesChildren[*ast.ListItem](children)
	children = getSpecificTypesChildren[*ast.Paragraph](children)
	texts := Filter(children, isType[*ast.Text])
	titles := Map(texts, func(node ast.Node) string {
		text := node.(*ast.Text)
		return string(text.Literal)
	})
	return titles
}

func decisionStatus(status string) Status {
	switch status {
	case "draft":
		return DraftStatus
	case "proposed":
		return ProposedStatus
	case "rejected":
		return RejectedStatus
	case "accepted":
		return AcceptedStatus
	case "deprecated":
		return DeprecatedStatus
	case "superseded":
		return SupersededStatus
	default:
		return ""
	}
}

func parseStatus(items []ast.Node) (string, error) {
	children := getSpecificTypesChildren[*ast.List](items)
	children = getSpecificTypesChildren[*ast.ListItem](children)
	children = getSpecificTypesChildren[*ast.Paragraph](children)
	texts := Filter(children, isType[*ast.Text])
	titles := Map(texts, func(node ast.Node) string {
		text := node.(*ast.Text)
		return string(text.Literal)
	})
	titles = Filter(titles, func(s string) bool {
		return strings.Contains(s, "Status:")
	})
	if len(titles) != 1 {
		return "", errors.New("size is wrong")
	}
	title := titles[0]
	title = strings.ReplaceAll(title, "Status:", "")
	title = strings.TrimSpace(title)
	return title, nil
}

func getSpecificTypesChildren[T any](items []ast.Node) []ast.Node {
	lists := Filter(items, isType[T])
	children := FlatMap(lists, getChildren)
	return children
}

func isType[T any](node ast.Node) bool {
	_, ok := node.(T)
	return ok
}

func getChildren(node ast.Node) []ast.Node {
	return node.GetChildren()
}

func parseLinks(nodes []ast.Node) []Link {
	children := getSpecificTypesChildren[*ast.List](nodes)
	children = getSpecificTypesChildren[*ast.ListItem](children)
	children = getSpecificTypesChildren[*ast.Paragraph](children)
	children = Filter(children, isType[*ast.Link])

	var ret []Link
	for _, child := range children {
		link := child.(*ast.Link)
		ret = append(ret, Link{
			Title:       parseText(link),
			Destination: string(link.Destination),
		})
	}
	return ret
}
