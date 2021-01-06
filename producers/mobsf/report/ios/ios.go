// Package ios provides types and functions for working with iOS project scan
// reports from MobSF.
package ios

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/thought-machine/dracon/api/proto/v1"
	mreport "github.com/thought-machine/dracon/producers/mobsf/report"
)

// Report represents a (partial) iOS project scan report.
type Report struct {
	RootDir                string                                 `json:"-"`
	CodeAnalysis           map[string]mreport.CodeAnalysisFinding `json:"code_analysis"`
	CodeAnalysisExclusions map[string]bool                        `json:"-"`
	ATSAnalysis            []ATSAnalysisFinding                   `json:"ats_analysis"`
}

// ATSAnalysisFinding represents the App Transport Security (ATS) findings in an
// iOS project scan report.
type ATSAnalysisFinding struct {
	Issue       string `json:"issue"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func NewReport(report []byte, exclusions map[string]bool) (mreport.Report, error) {
	var r *Report
	if err := json.Unmarshal(report, &r); err != nil {
		return nil, err
	}

	r.CodeAnalysisExclusions = exclusions

	return r, nil
}

func (r *Report) SetRootDir(path string) {
	r.RootDir = path
}

func (r *Report) AsIssues() []*v1.Issue {
	issues := make([]*v1.Issue, 0)

	for id, finding := range r.CodeAnalysis {
		if _, exists := r.CodeAnalysisExclusions[id]; exists {
			continue
		}

		for filename, linesList := range finding.Files {
			for _, line := range strings.Split(linesList, ",") {
				issues = append(issues, &v1.Issue{
					Target:      fmt.Sprintf("%s:%s", filepath.Join(r.RootDir, filename), line),
					Type:        id,
					Title:       finding.Metadata.CWE,
					Severity:    v1.Severity(v1.Severity_value[fmt.Sprintf("SEVERITY_%s", strings.ToUpper(finding.Metadata.Severity))]),
					Cvss:        finding.Metadata.CVSS,
					Confidence:  v1.Confidence_CONFIDENCE_INFO,
					Description: finding.Metadata.Description,
				})
			}
		}
	}

	// ATS analysis findings are per plist file, but projects could contain more
	// than one plist file and MobSF doesn't specify which plist file each
	// finding is for, so they could be duplicated - remove any duplicates
	// before returning them as Dracon Issues
	seen := make(map[string]bool)
	for _, finding := range r.ATSAnalysis {
		if _, exists := seen[finding.Issue]; !exists {
			seen[finding.Issue] = true

			issue := &v1.Issue{
				// MobSF doesn't report the precise source of the issue so we
				// can't be more specific about its location than this:
				Target:      r.RootDir,
				Type:        "Insecure App Transport Security policy",
				Title:       finding.Issue,
				Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
				Description: fmt.Sprintf(
					"An insecure App Transport Security policy is defined in a plist file in the iOS app project directory %s.\n\nDetails:\n\n%s\n%s",
					r.RootDir,
					finding.Issue,
					finding.Description,
				),
			}

			switch finding.Status {
			case "info":
				issue.Severity = v1.Severity_SEVERITY_INFO
			case "warning":
				issue.Severity = v1.Severity_SEVERITY_LOW
			case "insecure":
				issue.Severity = v1.Severity_SEVERITY_MEDIUM
			case "secure":
				// We don't need to report this as an issue
				continue
			}

			issues = append(issues, issue)
		}
	}

	return issues
}
