package types

import (
	"encoding/json"
	"fmt"
	"strings"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/producers"

	"log"
)

func yarnToIssueSeverity(severity string) v1.Severity {

	switch severity {
	case "low":
		return v1.Severity_SEVERITY_LOW
	case "moderate":
		return v1.Severity_SEVERITY_MEDIUM
	case "high":
		return v1.Severity_SEVERITY_HIGH
	case "critical":
		return v1.Severity_SEVERITY_CRITICAL
	default:
		return v1.Severity_SEVERITY_INFO

	}
}

type YarnAuditLine struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func (yl *YarnAuditLine) UnmarshalJSON(data []byte) error {
	var typ struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(data, &typ); err != nil {
		return err
	}

	switch typ.Type {
	case "auditSummary":
		yl.Data = new(SummaryData)
	case "auditAdvisory":
		yl.Data = new(AuditData)
	case "auditAction":
		yl.Data = new(AuditActionData)
	default:
		log.Printf("Parsed unsupported type: %s", typ.Type)
	}

	type tmp YarnAuditLine // avoids infinite recursion
	return json.Unmarshal(data, (*tmp)(yl))

}

type AuditActionData struct {
	Cmd        string      `json:"cmd"`
	IsBreaking bool        `json:"isBreaking"`
	Action     AuditAction `json:"action"`
}

type AuditData struct {
	Resolution AuditResolution `json:"resolution"`
	Advisory   Advisory        `json:"advisory"`
}

func (audit *AuditData) AsIssue() *v1.Issue {
	var targetName string
	if audit.Resolution.Path != "" {
		targetName = audit.Resolution.Path + ": "
	}
	targetName += audit.Advisory.ModuleName

	return &v1.Issue{
		Target:      targetName,
		Type:        audit.Advisory.Cwe,
		Title:       audit.Advisory.Title,
		Severity:    yarnToIssueSeverity(audit.Advisory.Severity),
		Confidence:  v1.Confidence_CONFIDENCE_HIGH,
		Description: fmt.Sprintf("%s", audit.Advisory.GetDescription()),
		Cve:         strings.Join(audit.Advisory.Cves, ", "),
	}
}

type SummaryData struct {
	Vulnerabilities      Vulnerabilities `json:"vulnerabilities"`
	Dependencies         int             `json:"dependencies"`
	DevDependencies      int             `json:"devDependencies"`
	OptionalDependencies int             `json:"optionalDependencies"`
	TotalDependencies    int             `json:"totalDependencies"`
}

type AuditAction struct {
	Action   string            `json:"action"`
	Module   string            `json:"module"`
	Target   string            `json:"target"`
	IsMajor  bool              `json:"isMajor"`
	Resolves []AuditResolution `json:"resolves"`
}

type Vulnerabilities struct {
	Info     int `json:"info"`
	Low      int `json:"low"`
	Moderate int `json:"moderate"`
	High     int `json:"high"`
	Critical int `json:"critical"`
}

type Advisory struct {
	Findings           []Finding         `json:"findings"`
	Metadata           *AdvisoryMetaData `json:"metadata"`
	VulnerableVersions string            `json:"vulnerable_versions"`
	ModuleName         string            `json:"module_name"`
	Severity           string            `json:"severity"`
	GithubAdvisoryId   string            `json:"github_advisory_id"`
	Cves               []string          `json:"cves"`
	Access             string            `json:"access"`
	PatchedVersions    string            `json:"patched_versions"`
	Updated            string            `json:"updated"`
	Recommendation     string            `json:"recommendation"`
	Cwe                string            `json:"cwe"`
	FoundBy            *Contact          `json:"found_by"`
	Deleted            bool              `json:"deleted"`
	Id                 int               `json:"id"`
	References         string            `json:"references"`
	Created            string            `json:"created"`
	ReportedBy         *Contact          `json:"reported_by"`
	Title              string            `json:"title"`
	NpmAdvisoryId      interface{}       `json:"npm_advisory_id"`
	Overview           string            `json:"overview"`
	URL                string            `json:"url"`
}

func (advisory *Advisory) GetDescription() string {
	return fmt.Sprintf(
		"Vulnerable Versions: %s\nRecommendation: %s\nOverview: %s\nReferences:\n%s\nAdvisory URL: %s\n",
		advisory.VulnerableVersions,
		advisory.Recommendation,
		advisory.Overview,
		advisory.References,
		advisory.URL,
	)
}

type Finding struct {
	Version  string   `json:"version"`
	Paths    []string `json:"paths"`
	Dev      bool     `json:"dev"`
	Optional bool     `json:"optional"`
	Bundled  bool     `json:"bundled"`
}

type AuditResolution struct {
	Id       int    `json:"id"`
	Path     string `json:"path"`
	Dev      bool   `json:"dev"`
	Optional bool   `json:"optional"`
	Bundled  bool   `json:"bundled"`
}

type AdvisoryMetaData struct {
	Module_type         string `json:"module_type"`
	Exploitability      int    `json:"exploitability"`
	Affected_components string `json:"affected_components"`
}

type Contact struct {
	Name string `json: name`
}

type YarnAuditReport struct {
	AuditAdvisory []*AuditData
	AuditActions  []*AuditActionData
	Summary       *SummaryData
}

func NewReport(reportLines [][]byte) (*YarnAuditReport, []error) {

	var report YarnAuditReport

	var errors []error

	for _, line := range reportLines {
		var auditLine YarnAuditLine
		if err := producers.ParseJSON(line, &auditLine); err != nil {
			log.Printf("Error parsing JSON line '%s': %s\n", line, err)
			errors = append(errors, err)
		} else {

			switch auditLine.Data.(type) {
			case *SummaryData:
				report.Summary = auditLine.Data.(*SummaryData)
			case *AuditData:
				report.AuditAdvisory = append(report.AuditAdvisory, auditLine.Data.(*AuditData))
			case *AuditActionData:
				report.AuditActions = append(report.AuditActions, auditLine.Data.(*AuditActionData))
			}
		}
	}

	if len(report.AuditAdvisory) > 0 {
		return &report, errors
	}

	return nil, errors
}

func (r *YarnAuditReport) AsIssues() []*v1.Issue {
	issues := make([]*v1.Issue, 0)

	for _, audit := range r.AuditAdvisory {
		issues = append(issues, audit.AsIssue())
	}

	return issues
}
