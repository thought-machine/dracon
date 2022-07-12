package main

import (
	v1 "github.com/thought-machine/dracon/api/proto/v1"

	"log"

	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/thought-machine/dracon/producers"
)

func main() {
	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	inFile, err := producers.ReadInFile()
	if err != nil {
		log.Fatal(err)
	}

	results, err := sarif.FromString(string(inFile))
	if err != nil {
		log.Fatal(err)
	}
	for _, run := range results.Runs {
		tool := run.Tool.Driver.Name
		if err := producers.WriteDraconOut(tool, parseOut(*run)); err != nil {
			log.Fatal(err)
		}
	}
}

func parseOut(run sarif.Run) []*v1.Issue {
	issues := []*v1.Issue{}
	for _, res := range run.Results {
		for _, loc := range res.Locations {
			target := loc.PhysicalLocation.ArtifactLocation.URI
			issues = append(issues, &v1.Issue{
				Target:      *target,
				Title:       *res.RuleID,
				Description: *res.Message.Text,
				Type:        "Security Automation Result",
				Severity:    LevelToSeverity(*res.Level),
				Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
				Cvss:        0,
				Cve:         "",
			})
		}
	}
	return issues
}

func LevelToSeverity(level string) v1.Severity {
	if level == "error" {
		return v1.Severity_SEVERITY_HIGH
	} else if level == "warning" {
		return v1.Severity_SEVERITY_MEDIUM
	}
	return v1.Severity_SEVERITY_LOW
}
