package main

import (
	"fmt"
	"log"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/producers/elasticsearch_filebeat/types"

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

	var results types.ElasticSearchFilebeatResult
	if err := producers.ParseJSON(inFile, &results); err != nil {
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
			Target:      fmt.Sprintf("%s", h.Source.Host.Name),
			Type:        "Antivirus Issue",
			Title:       fmt.Sprintf("Antivirus Issue on %s", h.Source.Host.Name),
			Severity:    v1.Severity_SEVERITY_INFO,
			Cvss:        0.0,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: h.Source.Message,
		})

	}

	for _, b := range results.Aggregations.Aggregation.Buckets {

		name := b.Name
		h := b.Metric.Hits.Hits[0]
		issues = append(issues, &v1.Issue{
			Target:      fmt.Sprintf("%s", name),
			Type:        "Antivirus Issue",
			Title:       fmt.Sprintf("Antivirus Issue on %s", name),
			Severity:    v1.Severity_SEVERITY_INFO,
			Cvss:        0.0,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: h.Source.Message,
		})

	}

	return issues
}
