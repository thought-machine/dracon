package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/producers"
	"github.com/thought-machine/dracon/producers/docker_trivy/types"
)

// Combined flag to indicate the producer is being fed  aggregated input from multiple images
var Combined bool

func main() {
	flag.BoolVar(&Combined, "combinedout", false, "Output is the combined output of Trivy against multiple images, expects {<img-name>:[<regular trivy output>],<other-img>:[<trivy out for 'other-img'>]}")

	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	inFile, err := producers.ReadInFile()
	if err != nil {
		log.Fatal(err)
	}

	if Combined {
		var results types.CombinedOut
		if err := producers.ParseJSON(inFile, &results); err != nil {
			log.Fatal(err)
		}
		if err := producers.WriteDraconOut(
			"trivy",
			parseCombinedOut(results),
		); err != nil {
			log.Fatal(err)
		}
	} else {
		var results []types.TrivyOut
		if err := producers.ParseJSON(inFile, &results); err != nil {
			log.Fatal(err)
		}
		if err := producers.WriteDraconOut(
			"trivy",
			parseSingleOut(results),
		); err != nil {
			log.Fatal(err)
		}

	}
}
func parseCombinedOut(results types.CombinedOut) []*v1.Issue {
	issues := []*v1.Issue{}
	for img, output := range results {
		log.Printf("Parsing Combined Output for %s\n", img)
		for _, res := range output {
			target := res.Target
			for _, vuln := range res.Vulnerable {
				issues = append(issues, parseResult(&vuln, target))
			}
		}
	}
	return issues
}

func parseSingleOut(results []types.TrivyOut) []*v1.Issue {
	issues := []*v1.Issue{}
	for _, res := range results {
		target := res.Target

		for _, vuln := range res.Vulnerable {
			issues = append(issues, parseResult(&vuln, target))
		}
	}
	return issues
}

// TrivySeverityToDracon maps Trivy Severity Strings to dracon struct
func TrivySeverityToDracon(severity string) v1.Severity {
	switch severity {
	case "LOW":
		return v1.Severity_SEVERITY_LOW
	case "MEDIUM":
		return v1.Severity_SEVERITY_MEDIUM
	case "HIGH":
		return v1.Severity_SEVERITY_HIGH
	case "CRITICAL":
		return v1.Severity_SEVERITY_CRITICAL
	default:
		return v1.Severity_SEVERITY_INFO
	}
}

func parseResult(r *types.TrivyVulnerability, target string) *v1.Issue {
	return &v1.Issue{
		Target:     target,
		Type:       "Container image vulnerability",
		Title:      fmt.Sprintf("[%s][%s] %s", target, r.CVE, r.Title),
		Severity:   TrivySeverityToDracon(r.Severity),
		Confidence: v1.Confidence_CONFIDENCE_MEDIUM,
		Cvss:       r.CVSS.Nvd.V3Score,
		Description: fmt.Sprintf("CVSS Score: %v\nCvssVector: %s\nCve: %s\nCwe: %s\nReference: %s\nOriginal Description:%s\n",
			r.CVSS.Nvd.V3Score, r.CVSS.Nvd.V3Vector, r.CVE, strings.Join(r.CweIDs[:], ","), r.PrimaryURL, r.Description),
	}
}
