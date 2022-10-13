package types

import (
	"encoding/json"
	"fmt"
	"strings"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
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

type RawMessage = json.RawMessage

type auditAction struct {
	Type string 			`json:"type"`
	Data auditActionData	`json:"data"`
}

func (audit *auditAction) unmarshal(raw RawMessage) bool {
	if err := json.Unmarshal(raw, audit); err != nil {
		return false
	}
	return audit.Type == "auditAction"
}

type auditAdvisory struct {
	Type string 			`json:"type"`
	Data auditAdvisoryData 	`json:"data"`
}

func (audit *auditAdvisory) unmarshal(raw RawMessage) bool {
	if err := json.Unmarshal(raw, audit); err != nil {
		return false
	}
	return audit.Type == "auditAdvisory"
}

type auditSummary struct {
	Type string 			`json:"type"`
	Data auditSummaryData 	`json:"data"`
}

func (audit *auditSummary) unmarshal(raw RawMessage) bool {
	if err := json.Unmarshal(raw, audit); err != nil {
		return false
	}
	return audit.Type == "auditSummary"
}

type auditActionData struct {
	Cmd        string      			`json:"cmd"`
	IsBreaking bool        			`json:"isBreaking"`
	Action     auditActionAction 	`json:"action"`
}

type auditAdvisoryData struct {
	Resolution auditResolution `json:"resolution"`
	Advisory   yarnAdvisory        `json:"advisory"`
}

type auditSummaryData struct {
	Vulnerabilities      vulnerabilities `json:"vulnerabilities"`
	Dependencies         int             `json:"dependencies"`
	DevDependencies      int             `json:"devDependencies"`
	OptionalDependencies int             `json:"optionalDependencies"`
	TotalDependencies    int             `json:"totalDependencies"`
}

type auditActionAction struct {
	Action   string            `json:"action"`
	Module   string            `json:"module"`
	Target   string            `json:"target"`
	IsMajor  bool              `json:"isMajor"`
	Resolves []auditResolution `json:"resolves"`
}

type vulnerabilities struct {
	Info     int `json:"info"`
	Low      int `json:"low"`
	Moderate int `json:"moderate"`
	High     int `json:"high"`
	Critical int `json:"critical"`
}

type yarnAdvisory struct {
	Findings           []finding         `json:"findings"`
	Metadata           *advisoryMetaData `json:"metadata"`
	VulnerableVersions string            `json:"vulnerable_versions"`
	ModuleName         string            `json:"module_name"`
	Severity           string            `json:"severity"`
	GithubAdvisoryID   string            `json:"github_advisory_id"`
	Cves               []string          `json:"cves"`
	Access             string            `json:"access"`
	PatchedVersions    string            `json:"patched_versions"`
	Cvss			   cvss				 `json:"cvss"`
	Updated            string            `json:"updated"`
	Recommendation     string            `json:"recommendation"`
	Cwe                []string          `json:"cwe"`
	FoundBy            *contact          `json:"found_by"`
	Deleted            bool          	 `json:"deleted"`
	ID                 int               `json:"id"`
	References         string            `json:"references"`
	Created            string            `json:"created"`
	ReportedBy         *contact          `json:"reported_by"`
	Title              string            `json:"title"`
	NpmAdvisoryID      *interface{}      `json:"npm_advisory_id"`
	Overview           string            `json:"overview"`
	URL                string            `json:"url"`
}

type cvss struct {
	Score 		 json.Number `json:"score"`
	VectorString string `json:"vectorString"`
}

type finding struct {
	Version  string   `json:"version"`
	Paths    []string `json:"paths"`
	Dev      bool     `json:"dev"`
	Optional bool     `json:"optional"`
	Bundled  bool     `json:"bundled"`
}

type auditResolution struct {
	ID       int    `json:"id"`
	Path     string `json:"path"`
	Dev      bool   `json:"dev"`
	Optional bool   `json:"optional"`
	Bundled  bool   `json:"bundled"`
}

type advisoryMetaData struct {
	ModuleType         string `json:"module_type"`
	Exploitability      int    `json:"exploitability"`
	AffectedComponents string `json:"affected_components"`
}

type contact struct {
	Name string `json: name`
}

func NewReport(report []byte) (*[]auditAction, *[]auditAdvisory, *[]auditSummary, error) {

	var raws []RawMessage
	if err := json.Unmarshal(report, &raws); err != nil {
		return nil, nil, nil, err
	}

	var auditActions []auditAction
	var auditAdvisories []auditAdvisory
	var auditSummaries []auditSummary

	for _, raw := range raws {
		auditAction := new(auditAction)
		if auditAction.unmarshal(raw) {
			auditActions = append(auditActions, *auditAction)
		}

		auditAdvisory := new(auditAdvisory)
		if auditAdvisory.unmarshal(raw) {
			auditAdvisories = append(auditAdvisories, *auditAdvisory)
		}

		auditSummary := new(auditSummary)
		if auditSummary.unmarshal(raw) {
			auditSummaries = append(auditSummaries, *auditSummary)
		}
	}

	return &auditActions, &auditAdvisories, &auditSummaries, nil
}

func (advisory *yarnAdvisory) GetDescription() string {
	return fmt.Sprintf(
		"Vulnerable Versions: %s\nRecommendation: %s\nOverview: %s\nReferences:\n%s\nAdvisory URL: %s\n",
		advisory.VulnerableVersions,
		advisory.Recommendation,
		advisory.Overview,
		advisory.References,
		advisory.URL,
	)
}

// AsIssue returns data as a Dracon v1.Issue
func (audit *auditAdvisoryData) AsIssue() *v1.Issue {
	var targetName string
	if audit.Resolution.Path != "" {
		targetName = audit.Resolution.Path + ": "
	}
	targetName += audit.Advisory.ModuleName

	// yarn audit now outputs CWEs as an array. if there is at least one CWE use the first, else use empty string
	cwe := ""
	if(len(audit.Advisory.Cwe) > 0) {
		cwe = audit.Advisory.Cwe[0]
	}

	return &v1.Issue{
		Target:      targetName,
		Type:        cwe,
		Title:       audit.Advisory.Title,
		Severity:    yarnToIssueSeverity(audit.Advisory.Severity),
		Confidence:  v1.Confidence_CONFIDENCE_HIGH,
		Description: fmt.Sprintf("%s", audit.Advisory.GetDescription()),
		Cve:         strings.Join(audit.Advisory.Cves, ", "),
	}
}

// AsIssues returns an auditAdvisory as Dracon v1.Issue list
func AsIssues(advisories *[]auditAdvisory) []*v1.Issue {
	issues := make([]*v1.Issue, 0)

	for _, audit := range *advisories {
		issues = append(issues, audit.Data.AsIssue())
	}

	return issues
}
