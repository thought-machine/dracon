package main

import (
	"encoding/json"
	"github.com/thought-machine/dracon/producers/pipsafety/types"
	"testing"

	v1 "github.com/thought-machine/dracon/api/proto/v1"

	"github.com/stretchr/testify/assert"
)

const exampleOutput = `[
	[
	"aegea",
	"<2.2.7",
	"2.2.6",
	"Aegea 2.2.7 avoids CVE-2018-1000805.",
	"37611"
],
[
	"pyyaml",
	"<5.3.1",
	"3.13",
	"A vulnerability was discovered in the PyYAML library in versions before 5.3.1, where it is susceptible to arbitrary code execution when it processes untrusted YAML files through the full_load method or with the FullLoader loader. Applications that use the library to process untrusted input may be vulnerable to this flaw. An attacker could use this flaw to execute arbitrary code on the system by abusing the python/object/new constructor. See: CVE-2020-1747.",
	"38100"
]
]
`

func TestParseIssues(t *testing.T) {
	safetyIssues := []types.SafetyIssue{}
	err := json.Unmarshal([]byte(exampleOutput), &safetyIssues)
	assert.Nil(t, err)
	draconIssues := parseIssues(safetyIssues)

	expectedIssue := &v1.Issue{
		Target:      "aegea",
		Type:        "Vulnerable Dependency",
		Title:       "aegea<2.2.7",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: "Aegea 2.2.7 avoids CVE-2018-1000805.\nCurrent Version: 2.2.6",
		Cve:         "CVE-2018-1000805",
	}
	issue2 := &v1.Issue{
		Target:      "pyyaml",
		Type:        "Vulnerable Dependency",
		Title:       "pyyaml<5.3.1",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: "A vulnerability was discovered in the PyYAML library in versions before 5.3.1, where it is susceptible to arbitrary code execution when it processes untrusted YAML files through the full_load method or with the FullLoader loader. Applications that use the library to process untrusted input may be vulnerable to this flaw. An attacker could use this flaw to execute arbitrary code on the system by abusing the python/object/new constructor. See: CVE-2020-1747.\nCurrent Version: 3.13",
		Cve:         "CVE-2020-1747",
	}
	assert.Equal(t, draconIssues[0], expectedIssue)
	assert.Equal(t, draconIssues[1], issue2)
}
