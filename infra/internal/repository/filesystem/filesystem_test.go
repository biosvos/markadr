package filesystem

import (
	"github.com/biosvos/markadr/flow/adr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseStatus(t *testing.T) {
	status, err := parseStatus([]byte(`# title

- Status: draft

## Context and Problem Statement

[Describe the context and problem statement, e.g., in free form using two to three sentences. You may want to articulate the problem in form of a question.]

## Decision Drivers

- driver1
- driver2

## Options

### [option 1]

| Pros | Cons |
|------|------|
| aa   | bb   |
| aa   | bb   |
|      | bb   |

## Decision Outcome

Pick: [option 1]  
- ...

## Links

- [Link type](link to adr)
`))
	require.NoError(t, err)
	require.Equal(t, adr.DraftStatus, status)
}