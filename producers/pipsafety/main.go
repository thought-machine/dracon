package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
	"github.com/thought-machine/dracon/producers"
)

type SafetyIssue struct {
	Name              string
	VersionConstraint string
	CurrentVersion    string
	Description       string
}

//read semi-unstructured safety json into struct
func (i *SafetyIssue) UnmarshalJSON(data []byte) error {

	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	i.Name, _ = v[0].(string)
	i.VersionConstraint, _ = v[1].(string)
	i.CurrentVersion = v[2].(string)
	i.Description = v[3].(string)

	return nil
}

func main() {
	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := ioutil.ReadFile(producers.InResults)
	if err != nil {
		log.Fatal(err)
	}
	issues := []SafetyIssue{}
	if err := json.Unmarshal(jsonBytes, &issues); err != nil {
		log.Fatal(err)
	}
	if err := producers.WriteDraconOut(
		"pipsafety",
		parseIssues(issues),
	); err != nil {
		log.Fatal(err)
	}
}

func parseIssues(out []SafetyIssue) []*v1.Issue {
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
