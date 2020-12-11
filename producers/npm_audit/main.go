package main

import (
	"github.com/thought-machine/dracon/producers"
	"github.com/thought-machine/dracon/producers/npm_audit/types/npm_full_audit"

	"flag"
	"log"
)

var (
	PackagePath string
)

func main() {
	flag.StringVar(&PackagePath, "package-path", "", "Path to the package.json file corresponding to this audit report; will be prepended to vulnerable dependency names in issue reports if specified")

	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

    inFile, err := producers.ReadInFile()
    if err != nil {
        log.Fatal(err)
    }

    report, err := npm_full_audit.NewReport(inFile, PackagePath)
    if err != nil {
		log.Fatal(err)
    }

	if err := producers.WriteDraconOut(
		"npm-audit",
		report.AsIssues(),
	); err != nil {
		log.Fatal(err)
	}
}
