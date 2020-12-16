package main

import (
	"github.com/thought-machine/dracon/producers"
	atypes "github.com/thought-machine/dracon/producers/npm_audit/types"
	"github.com/thought-machine/dracon/producers/npm_audit/types/npm_full_audit"
	"github.com/thought-machine/dracon/producers/npm_audit/types/npm_quick_audit"

	"errors"
	"flag"
	"log"
)

var (
	PackagePath string
)

func inFileToReport(inFile []byte) (atypes.Report, error) {
	if report, err := npm_quick_audit.NewReport(inFile); err == nil {
		return report, nil
	}

	if report, err := npm_full_audit.NewReport(inFile); err == nil {
		return report, nil
	}

	return nil, errors.New("input file is not a supported npm audit report format")
}

func main() {
	flag.StringVar(&PackagePath, "package-path", "", "Path to the package.json file corresponding to this audit report; will be prepended to vulnerable dependency names in issue reports if specified")

	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	inFile, err := producers.ReadInFile()
	if err != nil {
		log.Fatal(err)
	}

	report, err := inFileToReport(inFile)
	if err != nil {
		log.Fatal(err)
	}

	report.SetPackagePath(PackagePath)

	if err := producers.WriteDraconOut(
		"npm-audit",
		report.AsIssues(),
	); err != nil {
		log.Fatal(err)
	}
}
