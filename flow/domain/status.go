package domain

import "github.com/pkg/errors"

type Status string

const (
	DraftStatus      = Status("Draft")
	ProposedStatus   = Status("Proposed")
	RejectedStatus   = Status("Rejected")
	AcceptedStatus   = Status("Accepted")
	DeprecatedStatus = Status("Deprecated")
	SupersededStatus = Status("Superseded")
)

func (s Status) ValidateStatus() error {
	switch s {
	case DraftStatus:
		return nil
	case ProposedStatus:
		return nil
	case RejectedStatus:
		return nil
	case AcceptedStatus:
		return nil
	case DeprecatedStatus:
		return nil
	case SupersededStatus:
		return nil
	default:
		return errors.New("unknown status")
	}
}

func ValidateStatus(s Status) error {
	return s.ValidateStatus()
}
