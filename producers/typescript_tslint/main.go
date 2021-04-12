package main

import (
	"encoding/json"
	"fmt"
	"log"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/producers/typescript_tslint/types"

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

	var results []types.TSLintIssue
	if err := producers.ParseJSON(inFile, &results); err != nil {
		log.Fatal(err)
	}
	issues := parseIssues(results)
	if err := producers.WriteDraconOut(
		"tslint",
		issues,
	); err != nil {
		log.Fatal(err)
	}
}

func parseIssues(out []types.TSLintIssue) []*v1.Issue {
	issues := []*v1.Issue{}
	for _, r := range out {

		bytes, err := json.Marshal(r)
		if err != nil {
			log.Print("Couldn't write issue ", fmt.Sprintf("%s:%v-%v", r.Name, r.StartPosition.Line, r.EndPosition.Line))
		}

		issues = append(issues, &v1.Issue{
			Target:      fmt.Sprintf("%s:%v-%v", r.Name, r.StartPosition.Line, r.EndPosition.Line),
			Type:        r.RuleName,
			Title:       r.Failure,
			Severity:    v1.Severity_SEVERITY_MEDIUM,
			Cvss:        0.0,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: string(bytes),
		})
	}
	return issues
}
