package main

import (
	"encoding/json"
	"testing"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/producers/typescript_tslint/types"

	"github.com/stretchr/testify/assert"
)

const exampleOutput = `
[
	{
	"endPosition": {
	  "character": 63,
	  "line": 21,
	  "position": 774
	},
	"failure": "== should be ===",
	"fix": [
	  {
		"innerStart": 760,
		"innerLength": 6,
		"innerText": ""
	  },
	  {
		"innerStart": 773,
		"innerLength": 1,
		"innerText": "[]"
	  }
	],
	"name": "/foo/bar/js/types/File.ts",
	"ruleName": "triple-equals",
	"ruleSeverity": "error",
	"startPosition": {
	  "character": 49,
	  "line": 21,
	  "position": 760
	}
  },
  {
	"endPosition": {
	  "character": 63,
	  "line": 23,
	  "position": 774
	},
	"failure": "fail title",
	"name": "/some/path/foo/types/File.ts",
	"ruleName": "rule-name",
	"ruleSeverity": "error",
	"startPosition": {
	  "character": 49,
	  "line": 20,
	  "position": 760
	}
  }
  ]`

func TestParseIssues(t *testing.T) {
	var results []types.TSLintIssue
	err := json.Unmarshal([]byte(exampleOutput), &results)
	assert.Nil(t, err)
	issues := parseIssues(results)

	desc, _ := json.Marshal(results[0])
	expectedIssue := &v1.Issue{
		Target:      "/foo/bar/js/types/File.ts:21-21",
		Type:        "triple-equals",
		Title:       "== should be ===",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: string(desc),
	}
	desc, _ = json.Marshal(results[1])
	issue2 := &v1.Issue{
		Target:      "/some/path/foo/types/File.ts:20-23",
		Type:        "rule-name",
		Title:       "fail title",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: string(desc),
	}
	assert.Equal(t, expectedIssue, issues[0])
	assert.Equal(t, issue2, issues[1])
}
