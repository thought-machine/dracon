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

	report, err := types.NewReport(inLines)

	if len(err) > 0 {
		errorMessage := "Errors creating Yarn Audit report: %s"
		if report != nil{
			log.Printf(errorMessage, err)
		} else {
			log.Fatalf(errorMessage, err)
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
