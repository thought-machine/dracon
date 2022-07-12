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

func TestParseOutTrivy(t *testing.T) {
	results, err := sarif.FromString(trivyOutput)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	expectedIssues := []*v1.Issue{
		&v1.Issue{
			Target:      "library/ubuntu",
			Type:        "Security Automation Result",
			Title:       "CVE-2016-20013",
			Severity:    v1.Severity_SEVERITY_LOW,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: "Package: libc6\nInstalled Version: 2.35-0ubuntu3\nVulnerability CVE-2016-20013\nSeverity: LOW\nFixed Version: \nLink: [CVE-2016-20013](https://avd.aquasec.com/nvd/cve-2016-20013)",
		},
	}
	for _, run := range results.Runs {
		issues := parseOut(*run)

		assert.Equal(t, expectedIssues, issues)
	}
}

var trivyOutput = `{
	"version": "2.1.0",
	"$schema": "https://json.schemastore.org/sarif-2.1.0-rtm.5.json",
	"runs": [
	  {
		"tool": {
		  "driver": {
			"fullName": "Trivy Vulnerability Scanner",
			"informationUri": "https://github.com/aquasecurity/trivy",
			"name": "Trivy",
			"version": "0.29.2"
		  }
		},
		"results": [
		  {
			"ruleId": "CVE-2016-20013",
			"ruleIndex": 3,
			"level": "note",
			"message": {
			  "text": "Package: libc6\nInstalled Version: 2.35-0ubuntu3\nVulnerability CVE-2016-20013\nSeverity: LOW\nFixed Version: \nLink: [CVE-2016-20013](https://avd.aquasec.com/nvd/cve-2016-20013)"
			},
			"locations": [
			  {
				"physicalLocation": {
				  "artifactLocation": {
					"uri": "library/ubuntu",
					"uriBaseId": "ROOTPATH"
				  },
				  "region": {
					"startLine": 1,
					"startColumn": 1,
					"endLine": 1,
					"endColumn": 1
				  }
				}
			  }
			]
		  }
		],
		"columnKind": "utf16CodeUnits",
		"originalUriBaseIds": {
		  "ROOTPATH": {
			"uri": "file:///"
		  }
		}
	  }
	]
  }`
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
