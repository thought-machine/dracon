package main

import (
	"github.com/thought-machine/dracon/producers"
	"github.com/thought-machine/dracon/producers/yarn_audit/types"

	"log"
)

func main() {
	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	in, err := producers.ReadInFile()
	if err != nil {
		log.Fatal(err)
	}

	yarnReport, err := types.NewReport(in)
	if err != nil {
		log.Fatal(err)
	}

	if err := producers.WriteDraconOut(
		"yarn-audit",
		yarnReport.AuditAdvisories.AsIssues(),
	); err != nil {
		log.Fatal(err)
	}
}
