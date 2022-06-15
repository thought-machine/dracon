package main

import (
	"log"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/producers"

	securityhub "github.com/aws/aws-sdk-go-v2/service/securityhub/types"
)

func main() {
	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	inFile, err := producers.ReadInFile()
	if err != nil {
		log.Fatal(err)
	}

	var results securityHubFindings
	if err := producers.ParseJSON(inFile, &results); err != nil {
		log.Fatal(err)
	}

	issues := parseIssues(&results)

	if err := producers.WriteDraconOut("securityhub", issues); err != nil {
		log.Fatal(err)
	}
}

type securityHubFindings struct {
	Findings []securityhub.AwsSecurityFinding `json:"Findings"`
}

func parseIssues(shf *securityHubFindings) []*v1.Issue {
	severityMap := map[securityhub.SeverityLabel]v1.Severity{
		securityhub.SeverityLabelInformational: v1.Severity_SEVERITY_INFO,
		securityhub.SeverityLabelLow:           v1.Severity_SEVERITY_LOW,
		securityhub.SeverityLabelMedium:        v1.Severity_SEVERITY_MEDIUM,
		securityhub.SeverityLabelHigh:          v1.Severity_SEVERITY_HIGH,
		securityhub.SeverityLabelCritical:      v1.Severity_SEVERITY_CRITICAL,
	}

	issues := make([]*v1.Issue, len(shf.Findings))
	for i, r := range shf.Findings {
		issue := &v1.Issue{Confidence: v1.Confidence_CONFIDENCE_MEDIUM}

		if r.Title != nil {
			issue.Title = *r.Title
		}

		if r.Description != nil {
			issue.Description = *r.Description
		}

		if r.SourceUrl != nil {
			issue.Source = *r.SourceUrl
		}

		switch {
		case r.ProductName != nil && *r.ProductName == "Inspector" && len(r.Resources) > 0:
			if r.Resources[0].Details != nil && r.Resources[0].Details.AwsEc2Instance != nil {
				issue.Target = *r.Resources[0].Details.AwsEc2Instance.ImageId
			}
		case len(r.Resources) > 0:
			issue.Target = *r.Resources[0].Id
		case r.AwsAccountId != nil:
			issue.Target = *r.AwsAccountId
		}

		switch {
		case len(r.Types) > 0:
			issue.Type = r.Types[0]
		case r.ProductName != nil:
			issue.Type = *r.ProductName
		}

		switch {
		case r.Severity != nil:
			issue.Severity = severityMap[r.Severity.Label]
		case r.FindingProviderFields != nil && r.FindingProviderFields.Severity != nil:
			issue.Severity = severityMap[r.FindingProviderFields.Severity.Label]
		}

		if len(r.Vulnerabilities) > 0 {
			issue.Cve = *r.Vulnerabilities[0].Id

			highestCvss := 0.0
			for _, cvss := range r.Vulnerabilities[0].Cvss {
				if cvss.BaseScore > highestCvss {
					highestCvss = cvss.BaseScore
				}
			}
			issue.Cvss = highestCvss
		}

		issues[i] = issue
	}

	return issues
}
