package flow

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
	Options   []Option
	Selection int
}

type ADR struct {
	Title           string
	Status          Status
	Problem         string
	Context         string
	DecisionDrivers []string
	Options         Options
	Tags            []string
	Links           []string
}
