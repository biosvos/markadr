package domain

import "github.com/pkg/errors"

type Summary struct {
	Title  string
	Status Status
}

func NewSummary(title string, status Status) (*Summary, error) {
	err := validateSummary(title, status)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Summary{Title: title, Status: status}, nil
}

func validateSummary(title string, status Status) error {
	if title == "" {
		return errors.New("title is empty")
	}
	err := ValidateStatus(status)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
