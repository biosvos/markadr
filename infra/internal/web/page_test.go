package web

import (
	"github.com/biosvos/markadr/flow/adr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestADR(t *testing.T) {
	record := adr.ADR{
		Title:   "abc",
		Status:  "draft",
		Context: "이랬음",
		Problem: "문제",
		Drivers: []string{"속도"},
		Options: []*adr.Option{
			{
				Title: "캐시",
				Pros:  []string{"빠름"},
				Cons:  []string{"관리 어렵"},
			},
		},
		Outcomes: []*adr.Outcome{
			{
				Title:    "캐시",
				Contents: "빨리졌다.",
			},
		},
		References: []*adr.Reference{
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
