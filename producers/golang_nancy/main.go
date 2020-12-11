package main

import (
	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"strconv"
	types "github.com/thought-machine/dracon/producers/golang_nancy/types/nancy-issue"

	"fmt"
	"log"
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

	var results types.NancyOut
	if err := producers.ParseJSON(inFile, &results); err != nil {
		log.Fatal(err)
	}

	if err := producers.WriteDraconOut(
		"nancy",
		parseOut(&results),
	); err != nil {
		log.Fatal(err)
	}
}

func parseOut(results *types.NancyOut) []*v1.Issue {
	issues := []*v1.Issue{}
	for _, res := range results.Vulnerable {
		target := res.Coordinates
		for _, vuln := range res.Vulnerabilities {
			issues = append(issues, parseResult(&vuln, target))
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
	}
}
