package main

import (
	"log"

	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
	"github.com/thought-machine/dracon/producers"
)

func main() {
	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	var results SpotBugsOut
	// TODO(vj): XML parse and implement
	if err := producers.ParseInFileJSON(&results); err != nil {
		log.Fatal(err)
	}

	issues := []*v1.Issue{}
	for _, res := range results.Issues {
		issues = append(issues, parseResult(&res))
	}

	if err := producers.WriteDraconOut(
		"spotbugs",
		issues,
	); err != nil {
		log.Fatal(err)
	}
}

func parseResult(r *SpotBugsIssue) *v1.Issue {
	return &v1.Issue{
		// Target:      fmt.Sprintf("%s:%v", r.File, r.Line),
		// Type:        r.RuleID,
		// Title:       r.Code,
		// Severity:    v1.Severity(v1.Severity_value[fmt.Sprintf("SEVERITY_%s", r.Severity)]),
		// Cvss:        0.0,
		// Confidence:  v1.Confidence(v1.Confidence_value[fmt.Sprintf("CONFIDENCE_%s", r.Confidence)]),
		// Description: r.Details,
	}
}

// SpotBugsOut represents the output of a SpotBugs run
type SpotBugsOut struct {
	Issues []SpotBugsIssue `json:"Issues"`
	// Stats  SpotBugsStats   `json:"Stats"`
}

// SpotBugsIssue represents a SpotBugs Result
type SpotBugsIssue struct {
	Severity   string `json:"severity"`
	Confidence string `json:"confidence"`
	RuleID     string `json:"rule_id"`
	Details    string `json:"details"`
	File       string `json:"file"`
	Code       string `json:"code"`
	Line       string `json:"line"`
	Column     string `json:"column"`
}
