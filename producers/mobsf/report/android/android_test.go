package android

import (
	v1 "github.com/thought-machine/dracon/api/proto/v1"

	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var invalidJSON = `Not a valid JSON object`

func TestParseInvalidJSON(t *testing.T) {
	report, err := NewReport([]byte(invalidJSON), map[string]bool{})
	assert.Nil(t, report)
	assert.Error(t, err)
}

var androidReport = `{
  "code_analysis": {
    "android_ip_disclosure": {
      "files": { "test/MainApplication.java": "58" },
      "metadata": {
        "id": "android_ip_disclosure",
        "description": "IP Address disclosure",
        "type": "Regex",
        "pattern": "\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}",
        "severity": "warning",
        "input_case": "exact",
        "cvss": 4.3,
        "cwe": "CWE-200 Information Exposure",
        "owasp-mobile": "",
        "masvs": "MSTG-CODE-2"
      }
    },
    "android_insecure_random": {
      "files": { "test/MainApplication.java": "26" },
      "metadata": {
        "id": "android_insecure_random",
        "description": "The App uses an insecure Random Number Generator.",
        "type": "Regex",
        "pattern": "java\\.util\\.Random;",
        "severity": "high",
        "input_case": "exact",
        "cvss": 7.5,
        "cwe": "CWE-330 Use of Insufficiently Random Values",
        "owasp-mobile": "M5: Insufficient Cryptography",
        "masvs": "MSTG-CRYPTO-6"
      }
    }
  }
}
`

func TestParseValidIosReportNoExclusions(t *testing.T) {
	report, err := NewReport([]byte(androidReport), map[string]bool{})
	report.SetRootDir("android_project")
	assert.NoError(t, err)

	issues := report.AsIssues()
	assert.Len(t, issues, 2)

	expectedIssues := []*v1.Issue{
		&v1.Issue{
			Target:      "android_project/test/MainApplication.java:26",
			Type:        "android_insecure_random",
			Title:       "CWE-330 Use of Insufficiently Random Values",
			Severity:    v1.Severity_SEVERITY_HIGH,
			Cvss:        7.5,
			Description: "The App uses an insecure Random Number Generator.",
		},
		&v1.Issue{
			Target:      "android_project/test/MainApplication.java:58",
			Type:        "android_ip_disclosure",
			Title:       "CWE-200 Information Exposure",
			Cvss:        4.3,
			Description: "IP Address disclosure",
		},
	}

	sort.Slice(issues, func(i, j int) bool {
		return issues[i].Target < issues[j].Target
	})
	assert.Equal(t, issues, expectedIssues)
}

func TestParseValidIosReportExclusions(t *testing.T) {
	report, err := NewReport([]byte(androidReport), map[string]bool{"android_ip_disclosure": true})
	report.SetRootDir("android_project")
	assert.NoError(t, err)

	issues := report.AsIssues()
	assert.Len(t, issues, 1)

	expectedIssues := []*v1.Issue{
		&v1.Issue{
			Target:      "android_project/test/MainApplication.java:26",
			Type:        "android_insecure_random",
			Title:       "CWE-330 Use of Insufficiently Random Values",
			Severity:    v1.Severity_SEVERITY_HIGH,
			Cvss:        7.5,
			Description: "The App uses an insecure Random Number Generator.",
		},
	}

	assert.Equal(t, issues, expectedIssues)
}
