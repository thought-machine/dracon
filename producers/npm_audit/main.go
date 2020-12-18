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
	reportConstructors := []func([]byte) (atypes.Report, error){
		npm_quick_audit.NewReport,
		npm_full_audit.NewReport,
	}

	for _, constructor := range reportConstructors {
		report, err := constructor(inFile)

		switch err.(type) {
		case nil:
			return report, nil
		case *atypes.ParsingError, *atypes.FormatError:
			// Ignore parsing and incorrect format errors from constructors -
			// we'll just attempt again with the next one
		default:
			// Any other errors returned by a constructor are likely fatal
			return nil, err
		}
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

	log.Printf("Parsed input file as %s\n", report.Type())

	report.SetPackagePath(PackagePath)

	if err := producers.WriteDraconOut(
		"npm-audit",
		report.AsIssues(),
	); err != nil {
		log.Fatal(err)
	}
}
