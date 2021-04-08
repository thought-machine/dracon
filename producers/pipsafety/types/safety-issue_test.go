package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleOutput = `[
	[
	"aegea",
	"<2.2.7",
	"2.2.6",
	"Aegea 2.2.7 avoids CVE-2018-1000805.",
	"37611"
],[
	"aegea",
	"<2.2.7",
	"2.2.6",
	"Aegea 2.2.7 avoids CVE-2018-1000805.",
	"37611"
]
]
`

func TestUnmarshalJSON(t *testing.T) {
	expectedOutput := []SafetyIssue{{
		Name:              "aegea",
		CurrentVersion:    "2.2.6",
		Description:       "Aegea 2.2.7 avoids CVE-2018-1000805.",
		VersionConstraint: "<2.2.7",
	}, {
		Name:              "aegea",
		CurrentVersion:    "2.2.6",
		Description:       "Aegea 2.2.7 avoids CVE-2018-1000805.",
		VersionConstraint: "<2.2.7",
	}}
	safetyIssues := []SafetyIssue{}
	err := json.Unmarshal([]byte(exampleOutput), &safetyIssues)
	assert.Nil(t, err)
	assert.Equal(t, safetyIssues, expectedOutput)

}
