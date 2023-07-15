package domain

type Option struct {
	Title string   `json:"title"`
	Pros  []string `json:"pros"`
	Cons  []string `json:"cons"`
}

type Reference struct {
	Title       string `json:"title"`
	Destination string `json:"destination"`
}

type Outcome struct {
	Title    string `json:"title"`
	Contents string `json:"contents"`
}

type ADR struct {
	Title      string       `json:"title"`
	Status     Status       `json:"status"`
	Context    string       `json:"context"`
	Problem    string       `json:"problem"`
	Drivers    []string     `json:"drivers"`
	Options    []*Option    `json:"options"`
	Outcomes   []*Outcome   `json:"outcome"`
	References []*Reference `json:"references"`
}

const (
	ContextAndProblemStatement = "Context and Problem Statement"
	DecisionDrivers            = "Decision Drivers"
	DecisionOutcome            = "Decision Outcome"
	ProsAndCons                = "Options"
	Links                      = "Links"
)
