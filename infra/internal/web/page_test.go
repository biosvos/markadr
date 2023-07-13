package web

import (
	"github.com/biosvos/markadr/flow/adr"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_newNavigator(t *testing.T) {
	tests := []struct {
		name string
		toc  *adr.TOC
		want string
	}{
		{
			name: "",
			toc:  &adr.TOC{},
			want: "",
		},
		{
			name: "",
			toc: &adr.TOC{
				Rows: []adr.Row{
					{
						Title: "title",
						Depth: 0,
					},
				},
			},
			want: "<li><a href='#0'>title</a></li>",
		},
		{
			name: "",
			toc: &adr.TOC{
				Rows: []adr.Row{
					{
						Title: "title",
						Depth: 0,
					},
					{
						Title: "title",
						Depth: 0,
					},
				},
			},
			want: "<li><a href='#0'>title</a></li><li><a href='#1'>title</a></li>",
		},
		{
			name: "",
			toc: &adr.TOC{
				Rows: []adr.Row{
					{
						Title: "title",
						Depth: 0,
					},
					{
						Title: "title",
						Depth: 1,
					},
				},
			},
			want: "<li><a href='#0'>title</a><ul><li><a href='#1'>title</a></li></ul></li>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newNavigator(tt.toc)
			require.Equal(t, tt.want, got)
		})
	}
}
