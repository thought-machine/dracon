package producers

import (
	"log"

	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
)

func ExampleParseFlags() {
	if err := ParseFlags(); err != nil {
		log.Fatal(err)
	}
}

func ExampleParseInFileJSON() {
	type GoSecOut struct {
		Issues []struct {
			Severity   string `json:"severity"`
			Confidence string `json:"confidence"`
			RuleID     string `json:"rule_id"`
			Details    string `json:"details"`
			File       string `json:"file"`
			Code       string `json:"code"`
			Line       string `json:"line"`
			Column     string `json:"column"`
		} `json:"Issues"`
	}
	var results GoSecOut
	if err := ParseInFileJSON(&results); err != nil {
		log.Fatal(err)
	}
}

func ExampleWriteDraconOut() {
	issues := []*v1.Issue{}
	if err := WriteDraconOut(
		"gosec",
		issues,
	); err != nil {
		log.Fatal(err)
	}
}
