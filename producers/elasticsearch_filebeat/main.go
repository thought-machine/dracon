package main

import (
	"fmt"
	"log"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	types "github.com/thought-machine/dracon/producers/elasticsearch_filebeat/types/elasticsearch-filebeat-issue"

	"github.com/thought-machine/dracon/producers"
)

func main() {
	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}
	var results types.ElasticSearchFilebeatResult
	if err := producers.ParseInFileJSON(&results); err != nil {
		log.Fatal(err)
	}

	issues := parseIssues(&results)
	if err := producers.WriteDraconOut(
		"elasticsearch-filebeat",
		issues,
	); err != nil {
		log.Fatal(err)
	}
}

func parseIssues(results *types.ElasticSearchFilebeatResult) []*v1.Issue {
	issues := []*v1.Issue{}

	for _, h := range results.Hits.Hits {

		issues = append(issues, &v1.Issue{
			Target:      fmt.Sprintf("%s (%s)", h.Source.Host.Name, h.Source.Host.ID),
			Type:        "Antivirus Issue",
			Title:       fmt.Sprintf("Antivirus Issue on %s", h.Source.Host.Name),
			Severity:    v1.Severity_SEVERITY_INFO,
			Cvss:        0.0,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: h.Source.Message,
		})

	}
	return issues
}
