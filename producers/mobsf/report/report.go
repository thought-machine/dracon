// Package report provides common types for scan report formats.
package report

import (
	v1 "github.com/thought-machine/dracon/api/proto/v1"
)

// Report is an interface for scan report formats.
type Report interface {
	// SetRootDir sets the path to this project's root directory.
	SetRootDir(string)

	// AsIssues transforms this Report into a slice of Dracon Issues that can be
	// processed by the Dracon enricher.
	AsIssues() []*v1.Issue
}

// CodeAnalysisFinding Describes a mobsf finding
type CodeAnalysisFinding struct {
	Files    map[string]string    `json:"files"`
	Metadata codeAnalysisMetadata `json:"metadata"`
}

type codeAnalysisMetadata struct {
	CVSS        float64 `json:"cvss"`
	CWE         string  `json:"cwe"`
	Description string  `json:"description"`
	Severity    string  `json:"severity"`
}
