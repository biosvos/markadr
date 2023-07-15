package filesystem

import (
	"github.com/biosvos/markadr/flow/domain"
	"github.com/pkg/errors"
	"strings"
)

func decisionStatus(status string) domain.Status {
	status = strings.ToLower(status)
	switch status {
	case "draft":
		return domain.DraftStatus
	case "proposed":
		return domain.ProposedStatus
	case "rejected":
		return domain.RejectedStatus
	case "accepted":
		return domain.AcceptedStatus
	case "deprecated":
		return domain.DeprecatedStatus
	case "superseded":
		return domain.SupersededStatus
	default:
		return ""
	}
}

func parseStatus(contents []byte) (domain.Status, error) {
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
