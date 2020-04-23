package main

import (
	"encoding/json"
	"testing"

	v1 "api/proto/v1"

	"github.com/stretchr/testify/assert"
)

const exampleOutput = `
{
	"Issues": [
		{
			"severity": "MEDIUM",
			"confidence": "HIGH",
			"rule_id": "G304",
			"details": "Potential file inclusion via variable",
			"file": "/tmp/source/foo.go",
			"code": "ioutil.ReadFile(path)",
			"line": "33",
			"column": "44"
		}
	],
	"Stats": {
		"files": 1,
		"lines": 60,
		"nosec": 0,
		"found": 1
	}
}`

func TestParseIssues(t *testing.T) {
	var results GoSecOut
	err := json.Unmarshal([]byte(exampleOutput), &results)
	assert.Nil(t, err)

	issues := parseIssues(&results)

	expectedIssue := &v1.Issue{
		Target:      "/tmp/source/foo.go:33",
		Type:        "G304",
		Title:       "ioutil.ReadFile(path)",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_HIGH,
		Description: "Potential file inclusion via variable",
	}

	assert.Equal(t, []*v1.Issue{expectedIssue}, issues)
}
