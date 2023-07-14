package filesystem

import (
	"github.com/biosvos/markadr/flow/adr"
	"github.com/pkg/errors"
	"strings"
)

func decisionStatus(status string) adr.Status {
	status = strings.ToLower(status)
	switch status {
	case "draft":
		return adr.DraftStatus
	case "proposed":
		return adr.ProposedStatus
	case "rejected":
		return adr.RejectedStatus
	case "accepted":
		return adr.AcceptedStatus
	case "deprecated":
		return adr.DeprecatedStatus
	case "superseded":
		return adr.SupersededStatus
	default:
		return ""
	}
}

func parseStatus(contents []byte) (adr.Status, error) {
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		const statusSection = "- Status:"
		if !strings.Contains(line, statusSection) {
			continue
		}
		line = strings.ReplaceAll(line, statusSection, "")
		line = strings.TrimSpace(line)
		status := decisionStatus(line)
		return status, nil
	}
	return "", errors.New("not found status section")
}
