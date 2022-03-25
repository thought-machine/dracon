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

	inLines, err := producers.ReadLines()
	if err != nil {
		log.Fatal(err)
	}

	report, errors := types.NewReport(inLines)

	// Individual errors should already be printed to logs
	if len(errors) > 0 {
		errorMessage := "Errors creating Yarn Audit report: %d"
		if report != nil{
			log.Printf(errorMessage, len(errors))
		} else {
			log.Fatalf(errorMessage, len(errors))
		}
	}

	if report != nil {
		if err := producers.WriteDraconOut(
			"yarn-audit",
			report.AsIssues(),
		); err != nil {
			log.Fatal(err)
		}
	}
}
