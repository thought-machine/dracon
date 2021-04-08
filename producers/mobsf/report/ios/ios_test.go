package ios

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

var iOSReport = `{
  "ats_analysis": [
    {
      "issue": "App Transport Security AllowsArbitraryLoads is allowed",
      "status": "insecure",
      "description": "App Transport Security restrictions are disabled for all network connections. Disabling ATS means that unsecured HTTP connections are allowed. HTTPS connections are also allowed, and are still subject to default server trust evaluation. However, extended security checks like requiring a minimum Transport Layer Security (TLS) protocol version\u2014are disabled. This setting is not applicable to domains listed in NSExceptionDomains."
    },
    {
      "issue": "NSExceptionRequiresForwardSecrecy set to NO for localhost",
      "status": "insecure",
      "description": "NSExceptionRequiresForwardSecrecy limits the accepted ciphers to those that support perfect forward secrecy (PFS) through the Elliptic Curve Diffie-Hellman Ephemeral (ECDHE) key exchange. Set the value for this key to NO to override the requirement that a server must support PFS for the given domain. This key is optional. The default value is YES, which limits the accepted ciphers to those that support PFS through Elliptic Curve Diffie-Hellman Ephemeral (ECDHE) key exchange."
    },
    {
      "issue": "App Transport Security AllowsArbitraryLoads is allowed",
      "status": "insecure",
      "description": "App Transport Security restrictions are disabled for all network connections. Disabling ATS means that unsecured HTTP connections are allowed. HTTPS connections are also allowed, and are still subject to default server trust evaluation. However, extended security checks like requiring a minimum Transport Layer Security (TLS) protocol version\u2014are disabled. This setting is not applicable to domains listed in NSExceptionDomains."
    }
  ],
  "code_analysis": {
    "ios_app_logging": {
      "files": { "file.m": "31" },
      "metadata": {
        "id": "ios_app_logging",
        "cvss": 7.5,
        "cwe": "CWE-532 Insertion of Sensitive Information into Log File",
        "description": "The App logs information. Sensitive information should never be logged.",
        "input_case": "exact",
        "masvs": "MSTG-STORAGE-3",
        "owasp-mobile": "",
        "pattern": "NSLog|NSAssert|fprintf|fprintf|Logging",
        "severity": "info",
        "type": "Regex"
      }
    },
    "ios_swift_log": {
      "files": {
        "file1.swift": "62",
        "file2.swift": "37,16"
      },
      "metadata": {
        "id": "ios_swift_log",
        "cvss": 7.5,
        "cwe": "CWE-532",
        "description": "The App logs information. Sensitive information should never be logged.",
        "input_case": "exact",
        "masvs": "MSTG-STORAGE-3",
        "owasp-mobile": "",
        "pattern": "(print|NSLog|os_log|OSLog|os_signpost)\\(.*\\)",
        "severity": "info",
        "type": "Regex"
      }
    }
  }
}
`

func TestParseValidIosReportNoExclusions(t *testing.T) {
	report, err := NewReport([]byte(iOSReport), map[string]bool{})
	report.SetRootDir("ios_project")
	assert.NoError(t, err)

	issues := report.AsIssues()
	assert.Len(t, issues, 6)

	expectedIssues := []*v1.Issue{
		{
			Target:      "ios_project",
			Type:        "Insecure App Transport Security policy",
			Title:       "App Transport Security AllowsArbitraryLoads is allowed",
			Severity:    v1.Severity_SEVERITY_MEDIUM,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: "An insecure App Transport Security policy is defined in a plist file in the iOS app project directory ios_project.\n\nDetails:\n\nApp Transport Security AllowsArbitraryLoads is allowed\nApp Transport Security restrictions are disabled for all network connections. Disabling ATS means that unsecured HTTP connections are allowed. HTTPS connections are also allowed, and are still subject to default server trust evaluation. However, extended security checks like requiring a minimum Transport Layer Security (TLS) protocol version\342\200\224are disabled. This setting is not applicable to domains listed in NSExceptionDomains.",
		},
		{
			Target:      "ios_project",
			Type:        "Insecure App Transport Security policy",
			Title:       "NSExceptionRequiresForwardSecrecy set to NO for localhost",
			Severity:    v1.Severity_SEVERITY_MEDIUM,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: "An insecure App Transport Security policy is defined in a plist file in the iOS app project directory ios_project.\n\nDetails:\n\nNSExceptionRequiresForwardSecrecy set to NO for localhost\nNSExceptionRequiresForwardSecrecy limits the accepted ciphers to those that support perfect forward secrecy (PFS) through the Elliptic Curve Diffie-Hellman Ephemeral (ECDHE) key exchange. Set the value for this key to NO to override the requirement that a server must support PFS for the given domain. This key is optional. The default value is YES, which limits the accepted ciphers to those that support PFS through Elliptic Curve Diffie-Hellman Ephemeral (ECDHE) key exchange.",
		},
		{
			Target:      "ios_project/file.m:31",
			Type:        "ios_app_logging",
			Title:       "CWE-532 Insertion of Sensitive Information into Log File",
			Cvss:        7.5,
			Description: "The App logs information. Sensitive information should never be logged.",
		},
		{
			Target:      "ios_project/file1.swift:62",
			Type:        "ios_swift_log",
			Title:       "CWE-532",
			Cvss:        7.5,
			Description: "The App logs information. Sensitive information should never be logged.",
		},
		{
			Target:      "ios_project/file2.swift:16",
			Type:        "ios_swift_log",
			Title:       "CWE-532",
			Cvss:        7.5,
			Description: "The App logs information. Sensitive information should never be logged.",
		},
		{
			Target:      "ios_project/file2.swift:37",
			Type:        "ios_swift_log",
			Title:       "CWE-532",
			Cvss:        7.5,
			Description: "The App logs information. Sensitive information should never be logged.",
		},
	}

	sort.Slice(issues, func(i, j int) bool {
		return issues[i].Target < issues[j].Target || issues[i].Description < issues[j].Description
	})
	assert.Equal(t, issues, expectedIssues)
}

func TestParseValidIosReportExclusions(t *testing.T) {
	report, err := NewReport([]byte(iOSReport), map[string]bool{"ios_swift_log": true})
	report.SetRootDir("ios_project")
	assert.NoError(t, err)

	issues := report.AsIssues()
	assert.Len(t, issues, 3)

	expectedIssues := []*v1.Issue{
		{
			Target:      "ios_project",
			Type:        "Insecure App Transport Security policy",
			Title:       "App Transport Security AllowsArbitraryLoads is allowed",
			Severity:    v1.Severity_SEVERITY_MEDIUM,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: "An insecure App Transport Security policy is defined in a plist file in the iOS app project directory ios_project.\n\nDetails:\n\nApp Transport Security AllowsArbitraryLoads is allowed\nApp Transport Security restrictions are disabled for all network connections. Disabling ATS means that unsecured HTTP connections are allowed. HTTPS connections are also allowed, and are still subject to default server trust evaluation. However, extended security checks like requiring a minimum Transport Layer Security (TLS) protocol version\342\200\224are disabled. This setting is not applicable to domains listed in NSExceptionDomains.",
		},
		{
			Target:      "ios_project",
			Type:        "Insecure App Transport Security policy",
			Title:       "NSExceptionRequiresForwardSecrecy set to NO for localhost",
			Severity:    v1.Severity_SEVERITY_MEDIUM,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: "An insecure App Transport Security policy is defined in a plist file in the iOS app project directory ios_project.\n\nDetails:\n\nNSExceptionRequiresForwardSecrecy set to NO for localhost\nNSExceptionRequiresForwardSecrecy limits the accepted ciphers to those that support perfect forward secrecy (PFS) through the Elliptic Curve Diffie-Hellman Ephemeral (ECDHE) key exchange. Set the value for this key to NO to override the requirement that a server must support PFS for the given domain. This key is optional. The default value is YES, which limits the accepted ciphers to those that support PFS through Elliptic Curve Diffie-Hellman Ephemeral (ECDHE) key exchange.",
		},
		{
			Target:      "ios_project/file.m:31",
			Type:        "ios_app_logging",
			Title:       "CWE-532 Insertion of Sensitive Information into Log File",
			Cvss:        7.5,
			Description: "The App logs information. Sensitive information should never be logged.",
		},
	}

	sort.Slice(issues, func(i, j int) bool {
		return issues[i].Target < issues[j].Target || issues[i].Description < issues[j].Description
	})
	assert.Equal(t, issues, expectedIssues)
}
