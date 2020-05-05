package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
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
	results := []SafetyIssue{}
	err := json.Unmarshal([]byte(exampleOutput), &results)
	assert.Nil(t, err)

	issues := parseIssues(results)

	expectedIssue := &v1.Issue{
		Target:      "aegea",
		Type:        "Vulnerable Dependency",
		Title:       "aegea<2.2.7",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: "Aegea 2.2.7 avoids CVE-2018-1000805.\nCurrent Version: 2.2.6",
	}
	issue2 := &v1.Issue{
		Target:      "pyyaml",
		Type:        "Vulnerable Dependency",
		Title:       "pyyaml<5.3.1",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: "A vulnerability was discovered in the PyYAML library in versions before 5.3.1, where it is susceptible to arbitrary code execution when it processes untrusted YAML files through the full_load method or with the FullLoader loader. Applications that use the library to process untrusted input may be vulnerable to this flaw. An attacker could use this flaw to execute arbitrary code on the system by abusing the python/object/new constructor. See: CVE-2020-1747.\nCurrent Version: 3.13",
	}
	assert.Equal(t, issues[0].Target, expectedIssue.Target)
	assert.Equal(t, issues[0].Type, expectedIssue.Type)
	assert.Equal(t, issues[0].Title, expectedIssue.Title)
	assert.Equal(t, issues[0].Severity, expectedIssue.Severity)
	assert.Equal(t, issues[0].Cvss, expectedIssue.Cvss)
	assert.Equal(t, issues[0].Confidence, expectedIssue.Confidence)
	assert.Equal(t, issues[0].Description, expectedIssue.Description)

	assert.Equal(t, issues[1].Target, issue2.Target)
	assert.Equal(t, issues[1].Type, issue2.Type)
	assert.Equal(t, issues[1].Title, issue2.Title)
	assert.Equal(t, issues[1].Severity, issue2.Severity)
	assert.Equal(t, issues[1].Cvss, issue2.Cvss)
	assert.Equal(t, issues[1].Confidence, issue2.Confidence)
	assert.Equal(t, issues[1].Description, issue2.Description)

	// assert.Equal(t, []*v1.Issue{expectedIssue}, issues)
}
