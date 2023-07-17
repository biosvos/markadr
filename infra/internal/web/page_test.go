package web

import (
	"github.com/biosvos/markadr/flow/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestADR(t *testing.T) {
	record := domain.ADR{
		Title:   "abc",
		Status:  "draft",
		Context: "이랬음",
		Problem: "문제",
		Drivers: []string{"속도"},
		Options: []*domain.Option{
			{
				Title: "캐시",
				Pros:  []string{"빠름"},
				Cons:  []string{"관리 어렵"},
			},
		},
		Outcomes: []*domain.Outcome{
			{
				Title:    "캐시",
				Contents: "빨리졌다.",
			},
		},
		References: []*domain.Reference{
			{
				Title:       "b",
				Destination: "a",
			},
		},
	}

	adrhtml, err := makeADRHTML(&record)
	require.NoError(t, err)
	t.Log(adrhtml)
}
