// Package android provides types and functions for working with Android project
// scan reports from MobSF.
package android

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/thought-machine/dracon/api/proto/v1"
	mreport "github.com/thought-machine/dracon/producers/mobsf/report"
)

// Report represents a (partial) Android project scan report.
type Report struct {
	RootDir                string                                 `json:"-"`
	CodeAnalysis           map[string]mreport.CodeAnalysisFinding `json:"code_analysis"`
	CodeAnalysisExclusions map[string]bool                        `json:"-"`
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

	return issues
}
