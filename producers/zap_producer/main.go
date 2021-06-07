package main

import (

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/producers/zap_producer/types"

	"fmt"
	"log"

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

	var results types.ZapOut
	if err := producers.ParseJSON(inFile, &results); err != nil {
		log.Fatal(err)
	}

	if err := producers.WriteDraconOut(
		"zap",
		parseOut(&results),
	); err != nil {
		log.Fatal(err)
	}
}

func parseOut(results *types.ZapOut) []*v1.Issue {
	issues := []*v1.Issue{}
	for _, res := range results.Site{
		target:= res.Name
		for _, alert := range res.Alerts {
			issues = append (issues, parseIssue(&alert, target))
		}
	}
	return issues
}

//zap doesn't provide cvss so assigned as 0.0
func parseIssue(r *types.ZapAlerts, target string) *v1.Issue {
	var cvss = 0.0
	return &v1.Issue{
		Target:     target,
		Type:       r.CweId,
		Title:      r.Name ,
		Severity:   riskcodeToSeverity(r.RiskCode),
		Confidence: zapconfidenceToConfidence(r.Confidence),
		Cvss:       cvss,
		Description: fmt.Sprintf("Description: %s\nSolution: %s\nReference: %s\n", r.Description, r.Solution, r.Reference),
	}
}

//riskcode values are 0-INFO,1-LOW,2-MEDIUM,3-HIGH only available from ZAP. It is determined by the ZAP contributors 
func riskcodeToSeverity(riskcode string) v1.Severity {

	if (riskcode == "0") {

		return v1.Severity_SEVERITY_INFO

	} else if (riskcode == "1") {

		return v1.Severity_SEVERITY_LOW

	} else if (riskcode == "2") {

		return v1.Severity_SEVERITY_MEDIUM

	} else if (riskcode == "3") {

		return v1.Severity_SEVERITY_HIGH

	} else if (riskcode == "4") {

		return v1.Severity_SEVERITY_CRITICAL

	} else {

		return v1.Severity_SEVERITY_CRITICAL
	}	
}

//Confidence values are 0-INFO,1-LOW,2-MEDIUM,3-HIGH only available from ZAP. It is determined by the ZAP contributors 
func zapconfidenceToConfidence(confidence string) v1.Confidence {

		if (confidence == "0") {
	
			return v1.Confidence_CONFIDENCE_INFO
	
		} else if (confidence == "1") {
	
			return v1.Confidence_CONFIDENCE_LOW
	
		} else if (confidence == "2") {
	
			return v1.Confidence_CONFIDENCE_MEDIUM
	
		} else if (confidence == "3") {
	
			return v1.Confidence_CONFIDENCE_HIGH
	
		} else if (confidence == "4") {

			return v1.Confidence_CONFIDENCE_CRITICAL
	
		} else {
			
			return v1.Confidence_CONFIDENCE_CRITICAL
		}	
	}
