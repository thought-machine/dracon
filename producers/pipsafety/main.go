package main

import (
	"fmt"
	"log"

	v1 "api/proto/v1"
	"producers"
	"producers/pipsafety/types"
)

func parseIssues(out []types.SafetyIssue) []*v1.Issue {
	issues := []*v1.Issue{}
	for _, r := range out {
		issues = append(issues, &v1.Issue{
			Target:      r.Name,
			Type:        "Vulnerable Dependency",
			Title:       fmt.Sprintf("%s%s", r.Name, r.VersionConstraint),
			Severity:    v1.Severity_SEVERITY_MEDIUM,
			Cvss:        0.0,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: fmt.Sprintf("%s\nCurrent Version: %s", r.Description, r.CurrentVersion),
		})
	}
	return issues
}

func main() {
	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}
	issues := []types.SafetyIssue{}
	producers.ParseInFileJSON(&issues)
	if err := producers.WriteDraconOut(
		"pipsafety",
		parseIssues(issues),
	); err != nil {
		log.Fatal(err)
	}
}
