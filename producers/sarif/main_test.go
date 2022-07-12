package main

import (
	v1 "github.com/thought-machine/dracon/api/proto/v1"

	"testing"

	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
)

func TestParseOut(t *testing.T) {
	results, err := sarif.FromString(exampleOutput)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	expectedIssues := []*v1.Issue{
		&v1.Issue{
			Target:      "main.go",
			Type:        "Security Automation Result",
			Title:       "G404",
			Severity:    v1.Severity_SEVERITY_HIGH,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: "Use of weak random number generator (math/rand instead of crypto/rand)"},
		&v1.Issue{
			Target:      "main.go",
			Type:        "Security Automation Result",
			Title:       "G104",
			Severity:    v1.Severity_SEVERITY_MEDIUM,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: "Errors unhandled."},
	}
	for _, run := range results.Runs {
		issues := parseOut(*run)

		assert.Equal(t, expectedIssues, issues)
	}
}

var exampleOutput = `{
	"runs": [{
		"results": [{
				"level": "error",
				"locations": [{
					"physicalLocation": {
						"artifactLocation": {
							"uri": "main.go"
						},
						"region": {
							"endColumn": 7,
							"endLine": 83,
							"snippet": {
								"text": "r := rand.New(rand.NewSource(time.Now().UnixNano()))"
							},
							"sourceLanguage": "go",
							"startColumn": 7,
							"startLine": 83
						}
					}
				}],
				"message": {
					"text": "Use of weak random number generator (math/rand instead of crypto/rand)"
				},
				"ruleId": "G404"
			},
			{
				"level": "warning",
				"locations": [{
					"physicalLocation": {
						"artifactLocation": {
							"uri": "main.go"
						},
						"region": {
							"endColumn": 2,
							"endLine": 347,
							"snippet": {
								"text": "zipWriter.Close()"
							},
							"sourceLanguage": "go",
							"startColumn": 2,
							"startLine": 347
						}
					}
				}],
				"message": {
					"text": "Errors unhandled."
				},
				"ruleId": "G104",
				"ruleIndex": 3
			}
		],
		"tool": {
			"driver": {
				"guid": "8b518d5f-906d-39f9-894b-d327b1a421c5",
				"informationUri": "https://github.com/securego/gosec/",
				"name": "gosec"
			}
		}
	}],
	"$schema": "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json",
	"version": "2.1.0"
}`
