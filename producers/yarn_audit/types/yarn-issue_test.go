package types

import (
	"reflect"
	"testing"

	v1 "github.com/thought-machine/dracon/api/proto/v1"

	"github.com/stretchr/testify/assert"
)

var invalidJSON = `Not a valid JSON object`

func TestParseInvalidJSON(t *testing.T) {
	oneLine := []byte(invalidJSON)
	report, err := NewReport([][]byte{
		oneLine,
		oneLine,
	})

	assert.Nil(t, report)

	assert.Len(t, err, 2)
}

// In reality these would be single lines, but for readability in test these should also work
var fullYarnJSONLines [][]byte = [][]byte{
	[]byte(`{
    "type": "auditAdvisory",
    "data": {
      "resolution": {
        "id": 1004946,
        "path": "advisory1Path",
        "dev": false,
        "optional": false,
        "bundled": false
      },
      "advisory": {
        "findings": [
          {
            "version": "5.0.0",
            "paths": [
              "some/path",
              "another/path"
            ]
          },
          {
            "version": "5.0.0",
            "paths": [
              "more/findings/path"
            ]
          }
        ],
        "metadata": null,
        "vulnerable_versions": ">2.1.1 <5.0.1",
        "module_name": "super-awesome-module",
        "severity": "moderate",
        "github_advisory_id": "GHSA-93q8-gq69-wqmw",
        "cves": [
          "CVE-2022-0001"
        ],
        "access": "public",
        "patched_versions": ">=5.0.1",
        "updated": "2021-09-23T15:45:50.000Z",
        "recommendation": "Upgrade to version 5.0.1 or later",
        "cwe": "CWE-918",
        "found_by": null,
        "deleted": null,
        "id": 1004946,
        "references": "- https://advisory1.test.url/Ref1\n- https://advisory1.test.url/Ref2",
        "created": "2021-11-18T16:00:48.472Z",
        "reported_by": null,
        "title": "ADVISORY 1 TITLE",
        "npm_advisory_id": null,
        "overview": "Advisory 1 overview",
        "url": "https://advisory.1.url"
      }
    }
  }`),
	[]byte(`{
    "type": "unsupported",
    "data": {
      "vulnerabilities": {
        "info": 1,
        "low": 10,
        "moderate": 177,
        "high": 94,
        "critical": 4
      },
      "dependencies": 6274,
      "devDependencies": 0,
      "optionalDependencies": 0,
      "totalDependencies": 6274
    }
  }`),
	[]byte(`{
    "type": "auditAdvisory",
    "data": {
      "resolution": {
        "id": 1004947,
        "path": "advisory2Path",
        "dev": true,
        "optional": false,
        "bundled": false
      },
      "advisory": {
        "findings": [
          {
            "version": "1.1.0",
            "paths": [
              "some/path",
              "another/path"
            ]
          },
          {
            "version": "1.1.0",
            "paths": [
              "more/findings/path"
            ]
          }
        ],
        "metadata": null,
        "vulnerable_versions": ">1.1.1 <1.2.0",
        "module_name": "not-so-awesome-module",
        "severity": "low",
        "github_advisory_id": "GHSA-93q8-gq69-wqmw",
        "cves": [
          "CVE-2022-0002"
        ],
        "access": "public",
        "patched_versions": ">=1.2.0",
        "updated": "2021-09-23T15:45:50.000Z",
        "recommendation": "Upgrade to version 1.2.0 or later",
        "cwe": "CWE-920",
        "found_by": null,
        "deleted": null,
        "id": 1004947,
        "references": "- https://advisory2.test.url/Ref1\n- https://advisory2.test.url/Ref2\n- https://advisory2.test.url/Ref3",
        "created": "2021-11-18T16:00:48.472Z",
        "reported_by": null,
        "title": "ADVISORY 2 TITLE",
        "npm_advisory_id": null,
        "overview": "Advisory 2 overview",
        "url": "https://advisory.2.url"
      }
    }
  }`),
	[]byte(`{
    "type":"auditAction",
    "data":{
      "cmd":"action command",
      "isBreaking":false,
      "action":{
        "action":"action string",
        "module":"action module string",
        "target":"action target",
        "isMajor":true,
        "resolves":[
          {
            "id":1,
            "path":"action reolve path",
            "dev":true,
            "optional":true,
            "bundled":true
          }
        ]
      }
    }
  }`),
	[]byte(`{
    "completely": "unsupported"
  }`),
	[]byte(`{
    "type": "auditSummary",
    "data": {
      "vulnerabilities": {
        "info": 1,
        "low": 10,
        "moderate": 177,
        "high": 94,
        "critical": 4
      },
      "dependencies": 6274,
      "devDependencies": 0,
      "optionalDependencies": 0,
      "totalDependencies": 6274
    }
  }`),
}

func TestParseValidReportContainsAllSupportedFields(t *testing.T) {
	report, err := NewReport(
		fullYarnJSONLines,
	)

	assert.Nil(t, err)
	assert.NotNil(t, report)

	assert.NotNil(t, report.AuditSummary)
	assert.Len(t, report.AuditAdvisories, 2)
	assert.Len(t, report.AuditActions, 1)
}

func TestParseValidReportSummary(t *testing.T) {
	report, err := NewReport(
		fullYarnJSONLines,
	)

	assert.Nil(t, err)
	assert.NotNil(t, report)

	assert.NotNil(t, report.AuditSummary)

	expectedSummaryData := auditSummaryData{
		Vulnerabilities: vulnerabilities{
			Info:     1,
			Low:      10,
			Moderate: 177,
			High:     94,
			Critical: 4,
		},
		Dependencies:         6274,
		DevDependencies:      0,
		OptionalDependencies: 0,
		TotalDependencies:    6274,
	}

	assert.True(t, reflect.DeepEqual(&expectedSummaryData, report.AuditSummary), report.AuditSummary)
}

func TestParseValidReportAdvisories(t *testing.T) {
	report, err := NewReport(
		fullYarnJSONLines,
	)

	assert.Nil(t, err)
	assert.NotNil(t, report)

	assert.Len(t, report.AuditAdvisories, 2)

	expectedAdvisories := []*auditAdvisoryData{
		{
			Resolution: auditResolution{
				ID:       1004946,
				Path:     "advisory1Path",
				Dev:      false,
				Optional: false,
				Bundled:  false,
			},
			Advisory: yarnAdvisory{
				Findings: []finding{
					{
						Version: "5.0.0",
						Paths: []string{
							"some/path",
							"another/path",
						},
					},
					{
						Version: "5.0.0",
						Paths: []string{
							"more/findings/path",
						},
					},
				},
				Metadata:           nil,
				VulnerableVersions: ">2.1.1 <5.0.1",
				ModuleName:         "super-awesome-module",
				Severity:           "moderate",
				GithubAdvisoryID:   "GHSA-93q8-gq69-wqmw",
				Cves: []string{
					"CVE-2022-0001",
				},
				Access:          "public",
				PatchedVersions: ">=5.0.1",
				Updated:         "2021-09-23T15:45:50.000Z",
				Recommendation:  "Upgrade to version 5.0.1 or later",
				Cwe:             "CWE-918",
				FoundBy:         nil,
				Deleted:         false,
				ID:              1004946,
				References:      "- https://advisory1.test.url/Ref1\n- https://advisory1.test.url/Ref2",
				Created:         "2021-11-18T16:00:48.472Z",
				ReportedBy:      nil,
				Title:           "ADVISORY 1 TITLE",
				NpmAdvisoryID:   nil,
				Overview:        "Advisory 1 overview",
				URL:             "https://advisory.1.url",
			},
		},
		{
			Resolution: auditResolution{
				ID:       1004947,
				Path:     "advisory2Path",
				Dev:      true,
				Optional: false,
				Bundled:  false,
			},
			Advisory: yarnAdvisory{
				Findings: []finding{
					{
						Version: "1.1.0",
						Paths: []string{
							"some/path",
							"another/path",
						},
					},
					{
						Version: "1.1.0",
						Paths: []string{
							"more/findings/path",
						},
					},
				},
				Metadata:           nil,
				VulnerableVersions: ">1.1.1 <1.2.0",
				ModuleName:         "not-so-awesome-module",
				Severity:           "low",
				GithubAdvisoryID:   "GHSA-93q8-gq69-wqmw",
				Cves: []string{
					"CVE-2022-0002",
				},
				Access:          "public",
				PatchedVersions: ">=1.2.0",
				Updated:         "2021-09-23T15:45:50.000Z",
				Recommendation:  "Upgrade to version 1.2.0 or later",
				Cwe:             "CWE-920",
				FoundBy:         nil,
				Deleted:         false,
				ID:              1004947,
				References:      "- https://advisory2.test.url/Ref1\n- https://advisory2.test.url/Ref2\n- https://advisory2.test.url/Ref3",
				Created:         "2021-11-18T16:00:48.472Z",
				ReportedBy:      nil,
				Title:           "ADVISORY 2 TITLE",
				NpmAdvisoryID:   nil,
				Overview:        "Advisory 2 overview",
				URL:             "https://advisory.2.url",
			},
		},
	}

	assert.True(t, reflect.DeepEqual(expectedAdvisories, report.AuditAdvisories), report.AuditAdvisories)
}

func TestParseValidReportActions(t *testing.T) {
	report, err := NewReport(
		fullYarnJSONLines,
	)

	assert.Nil(t, err)
	assert.NotNil(t, report)

	assert.Len(t, report.AuditActions, 1)

	expectedActionData := auditActionData{
		Cmd:        "action command",
		IsBreaking: false,
		Action: auditAction{
			Action:  "action string",
			Module:  "action module string",
			Target:  "action target",
			IsMajor: true,
			Resolves: []auditResolution{
				{
					ID:       1,
					Path:     "action reolve path",
					Dev:      true,
					Optional: true,
					Bundled:  true,
				},
			},
		},
	}

	assert.True(t, reflect.DeepEqual(&expectedActionData, report.AuditActions[0]), report.AuditActions[0])
}

func TestParseValidReportAsIssues(t *testing.T) {
	report, err := NewReport(
		fullYarnJSONLines,
	)

	assert.Nil(t, err)

	assert.Len(t, report.AuditAdvisories, 2)

	issues := report.AsIssues()
	assert.Len(t, issues, 2)

	expectedIssues := []*v1.Issue{
		&v1.Issue{
			Target:     "advisory1Path: super-awesome-module",
			Type:       "CWE-918",
			Title:      "ADVISORY 1 TITLE",
			Severity:   v1.Severity_SEVERITY_MEDIUM,
			Confidence: v1.Confidence_CONFIDENCE_HIGH,
			Description: `Vulnerable Versions: >2.1.1 <5.0.1
Recommendation: Upgrade to version 5.0.1 or later
Overview: Advisory 1 overview
References:
- https://advisory1.test.url/Ref1
- https://advisory1.test.url/Ref2
Advisory URL: https://advisory.1.url
`,
			Cve: "CVE-2022-0001",
		},
		&v1.Issue{
			Target:     "advisory2Path: not-so-awesome-module",
			Type:       "CWE-920",
			Title:      "ADVISORY 2 TITLE",
			Severity:   v1.Severity_SEVERITY_LOW,
			Confidence: v1.Confidence_CONFIDENCE_HIGH,
			Description: `Vulnerable Versions: >1.1.1 <1.2.0
Recommendation: Upgrade to version 1.2.0 or later
Overview: Advisory 2 overview
References:
- https://advisory2.test.url/Ref1
- https://advisory2.test.url/Ref2
- https://advisory2.test.url/Ref3
Advisory URL: https://advisory.2.url
`,
			Cve: "CVE-2022-0002",
		},
	}

	for i := range expectedIssues {
		assert.Equal(t, expectedIssues[i].Target, issues[i].Target)
		assert.Equal(t, expectedIssues[i].Type, issues[i].Type)
		assert.Equal(t, expectedIssues[i].Title, issues[i].Title)
		assert.Equal(t, expectedIssues[i].Severity, issues[i].Severity)
		assert.Equal(t, expectedIssues[i].Cvss, issues[i].Cvss)
		assert.Equal(t, expectedIssues[i].Confidence, issues[i].Confidence)
		assert.Equal(t, expectedIssues[i].Description, issues[i].Description)
		assert.Equal(t, expectedIssues[i].Cve, issues[i].Cve)
	}
}
