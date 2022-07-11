package main

import (
	"strconv"

	"github.com/securego/gosec/report/sarif"
	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/producers/golang_nancy/types"

	"fmt"
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

	results, err := sarif.FromString(inFile)
	if err != nil {
		log.Fatal(err)
	}
	for _, run := range results.Runs {
		tool := run.Tool.Driver.Name
		if err := producers.WriteDraconOut(tool,
			parseOut(&results),
		); err != nil {
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
					Target:      target,
					Title:       res.RuleID,
					Description: res.Message.Text,
					Type: "",
					Severity: ,
					Confidence: ,
					Cvss: ,
					Cve: ,
				})
			}
		}
	return issues
}

func cvssToSeverity(score string) v1.Severity {

	switch s, err := strconv.ParseFloat(score, 64); err == nil {
	case 0.1 <= s && s <= 3.9:
		return v1.Severity_SEVERITY_LOW
	case 4.0 <= s && s <= 6.9:
		return v1.Severity_SEVERITY_MEDIUM
	case 7.0 <= s && s <= 8.9:
		return v1.Severity_SEVERITY_HIGH
	case 9.0 <= s && s <= 10.0:
		return v1.Severity_SEVERITY_CRITICAL
	default:
		return v1.Severity_SEVERITY_INFO

	}
}
func parseResult(r *types.NancyVulnerabilities, target string) *v1.Issue {
	cvss, err := strconv.ParseFloat(r.CvssScore, 64)
	if err != nil {
		cvss = 0.0
	}
	return &v1.Issue{
		Target:     target,
		Type:       "Vulnerable Dependency",
		Title:      r.Title,
		Severity:   cvssToSeverity(r.CvssScore),
		Confidence: v1.Confidence_CONFIDENCE_HIGH,
		Cvss:       cvss,
		Description: fmt.Sprintf("CVSS Score: %v\nCvssVector: %s\nCve: %s\nCwe: %s\nReference: %s\n",
			r.CvssScore, r.CvssVector, r.Cve, r.Cwe, r.Reference),
		Cve: r.Cve,
	}
}
