// Package npm_full_audit provides types and functions for working with audit
// reports from npm's "Full Audit" endpoint (/-/npm/v1/security/audits) and
// transforming them into data structures understood by the Dracon enricher.
// These reports are JSON objects consisting primarily of "advisories" (a list
// (of vulnerabilities known to affect the packages in the dependency tree) and
// "actions" (a list of steps that can be taken to remediate those
// vulnerabilities).
package npm_full_audit

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/producers"
	atypes "github.com/thought-machine/dracon/producers/npm_audit/types"
)

const PrintableType = "npm Full Audit report"

// Report represents an npm Full Audit report. The key for Advisories represents
// an npm advisory ID (i.e. https://npmjs.com/advisories/{int}).
type Report struct {
	PackagePath string           `json:"-"`
	Advisories  map[int]Advisory `json:"advisories"`
}

// Advisory represents a subset of information from an advisory in the
// "advisories" section of an npm Full Audit report.
type Advisory struct {
	Title              string `json:"title"`
	ModuleName         string `json:"module_name"`
	VulnerableVersions string `json:"vulnerable_versions"`
	Overview           string `json:"overview"`
	Recommendation     string `json:"recommendation"`
	References         string `json:"references"`
	Severity           string `json:"severity"`
	CWE                string `json:"cwe"`
	URL                string `json:"url"`
}

// NewReport constructs a Report from an npm Full Audit report.
func NewReport(report []byte) (atypes.Report, error) {
	var r *Report
	if err := producers.ParseJSON(report, &r); err != nil {
		switch err.(type) {
		case *json.InvalidUTF8Error, *json.SyntaxError, *json.UnmarshalFieldError,
			*json.UnmarshalTypeError, *json.UnsupportedTypeError, *json.UnsupportedValueError:
			return nil, &atypes.ParsingError{
				Type:          "npm_full_audit",
				PrintableType: PrintableType,
				Err:           err,
			}
		default:
			return nil, err
		}
	}

	// Full Audit reports have no metadata that identifies them - the clearest
	// differentiator between them and Quick Audit reports is that the top-level
	// "advisories" object only exists in Full Audit reports
	if r.Advisories == nil {
		return nil, &atypes.FormatError{
			Type:          "npm_full_audit",
			PrintableType: PrintableType,
		}
	}

	return r, nil
}

func (r *Report) SetPackagePath(packagePath string) {
	r.PackagePath = packagePath
}

func (r *Report) Type() string {
	return PrintableType
}

func (r *Report) AsIssues() []*v1.Issue {
	issues := make([]*v1.Issue, 0, len(r.Advisories))

	for _, a := range r.Advisories {
		var targetName string
		if r.PackagePath != "" {
			targetName = r.PackagePath + ":"
		}
		targetName += a.ModuleName

		issues = append(issues, &v1.Issue{
			Target:     targetName,
			Type:       "Vulnerable Dependency",
			Title:      a.Title,
			Severity:   v1.Severity(v1.Severity_value[fmt.Sprintf("SEVERITY_%s", strings.ToUpper(a.Severity))]),
			Confidence: v1.Confidence_CONFIDENCE_HIGH,
			Description: fmt.Sprintf("Vulnerable Versions: %s\nRecommendation: %s\nOverview: %s\nReferences: %s\nNPM Advisory URL: %s\n",
				a.VulnerableVersions, a.Recommendation, a.Overview, a.References, a.URL),
		})
	}

	return issues
}
