package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/producers/docker_trivy/types"
)

func TestParseCombinedOut(t *testing.T) {
	var results types.CombinedOut
	combinedOutput := fmt.Sprintf(`{"ubuntu:latest":%s,"alpine:latest":%s}`, exampleOutput, exampleOutput)
	err := json.Unmarshal([]byte(combinedOutput), &results)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
	issues := parseCombinedOut(results)

<<<<<<< HEAD
	expectedIssues := []*v1.Issue{{
		Target:     "ubuntu (ubuntu 18.04)",
		Type:       "Container image vulnerability",
		Title:      "[ubuntu (ubuntu 18.04)][CVE-2020-27350] apt: integer overflows and underflows while parsing .deb packages",
		Severity:   v1.Severity_SEVERITY_MEDIUM,
		Cvss:       5.7,
		Confidence: v1.Confidence_CONFIDENCE_MEDIUM,
		Description: fmt.Sprintf("CVSS Score: %v\nCvssVector: %s\nCve: %s\nCwe: %s\nReference: %s\nOriginal Description:%s\n",
			"5.7", "CVSS:3.1/AV:L/AC:L/PR:H/UI:N/S:C/C:L/I:L/A:L", "CVE-2020-27350",
			"CWE-190", "https://avd.aquasec.com/nvd/cve-2020-27350", "APT had several integer overflows and underflows while parsing .deb packages, aka GHSL-2020-168 GHSL-2020-169, in files apt-pkg/contrib/extracttar.cc, apt-pkg/deb/debfile.cc, and apt-pkg/contrib/arfile.cc. This issue affects: apt 1.2.32ubuntu0 versions prior to 1.2.32ubuntu0.2; 1.6.12ubuntu0 versions prior to 1.6.12ubuntu0.2; 2.0.2ubuntu0 versions prior to 2.0.2ubuntu0.2; 2.1.10ubuntu0 versions prior to 2.1.10ubuntu0.1;"),
	}}

	found := 0
	assert.Equal(t, 2, len(issues))
	for _, issue := range issues {
		singleMatch := 0
		for _, expected := range expectedIssues {
			if expected.Title == issue.Title {
				singleMatch++
				found++
				assert.Equal(t, singleMatch, 1) //assert no duplicates
				assert.EqualValues(t, expected.Type, issue.Type)
				assert.EqualValues(t, expected.Title, issue.Title)
				assert.EqualValues(t, expected.Severity, issue.Severity)
				assert.EqualValues(t, expected.Cvss, issue.Cvss)
				assert.EqualValues(t, expected.Confidence, issue.Confidence)
				assert.EqualValues(t, expected.Description, issue.Description)
			}
		}
	}
	assert.Equal(t, found, len(issues)) //assert everything has been found
}

func TestParseSingleOut(t *testing.T) {
	var results []types.TrivyOut
	err := json.Unmarshal([]byte(exampleOutput), &results)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
	issues := parseSingleOut(results)

	expectedIssues := []*v1.Issue{{
		Target:     "ubuntu (ubuntu 18.04)",
		Type:       "Container image vulnerability",
		Title:      "[ubuntu (ubuntu 18.04)][CVE-2020-27350] apt: integer overflows and underflows while parsing .deb packages",
		Severity:   v1.Severity_SEVERITY_MEDIUM,
		Cvss:       5.7,
		Confidence: v1.Confidence_CONFIDENCE_MEDIUM,
		Description: fmt.Sprintf("CVSS Score: %v\nCvssVector: %s\nCve: %s\nCwe: %s\nReference: %s\nOriginal Description:%s\n",
			"5.7", "CVSS:3.1/AV:L/AC:L/PR:H/UI:N/S:C/C:L/I:L/A:L", "CVE-2020-27350",
			"CWE-190", "https://avd.aquasec.com/nvd/cve-2020-27350", "APT had several integer overflows and underflows while parsing .deb packages, aka GHSL-2020-168 GHSL-2020-169, in files apt-pkg/contrib/extracttar.cc, apt-pkg/deb/debfile.cc, and apt-pkg/contrib/arfile.cc. This issue affects: apt 1.2.32ubuntu0 versions prior to 1.2.32ubuntu0.2; 1.6.12ubuntu0 versions prior to 1.6.12ubuntu0.2; 2.0.2ubuntu0 versions prior to 2.0.2ubuntu0.2; 2.1.10ubuntu0 versions prior to 2.1.10ubuntu0.1;"),
	}}

	found := 0
	assert.Equal(t, len(expectedIssues), len(issues))
	for _, issue := range issues {
		singleMatch := 0
		for _, expected := range expectedIssues {
			if expected.Title == issue.Title {
				singleMatch++
				found++
				assert.Equal(t, singleMatch, 1) //assert no duplicates
				assert.EqualValues(t, expected.Type, issue.Type)
				assert.EqualValues(t, expected.Title, issue.Title)
				assert.EqualValues(t, expected.Severity, issue.Severity)
				assert.EqualValues(t, expected.Cvss, issue.Cvss)
				assert.EqualValues(t, expected.Confidence, issue.Confidence)
				assert.EqualValues(t, expected.Description, issue.Description)
			}
		}
	}
	assert.Equal(t, found, len(issues)) //assert everything has been found
}

var exampleOutput = `
[
{
    "Target": "ubuntu (ubuntu 18.04)",
    "Type": "ubuntu",
    "Vulnerabilities": [
    {
        "VulnerabilityID": "CVE-2020-27350",
        "PkgName": "apt",
        "InstalledVersion": "1.6.12",
        "FixedVersion": "1.6.12ubuntu0.2",
        "Layer": {
            "DiffID": "sha256:a090697502b8d19fbc83afb24d8fb59b01e48bf87763a00ca55cfff42423ad36"
        },
        "SeveritySource": "ubuntu",
        "PrimaryURL": "https://avd.aquasec.com/nvd/cve-2020-27350",
        "Title": "apt: integer overflows and underflows while parsing .deb packages",
        "Description": "APT had several integer overflows and underflows while parsing .deb packages, aka GHSL-2020-168 GHSL-2020-169, in files apt-pkg/contrib/extracttar.cc, apt-pkg/deb/debfile.cc, and apt-pkg/contrib/arfile.cc. This issue affects: apt 1.2.32ubuntu0 versions prior to 1.2.32ubuntu0.2; 1.6.12ubuntu0 versions prior to 1.6.12ubuntu0.2; 2.0.2ubuntu0 versions prior to 2.0.2ubuntu0.2; 2.1.10ubuntu0 versions prior to 2.1.10ubuntu0.1;",
        "Severity": "MEDIUM",
        "CweIDs": [
        "CWE-190"
        ],
        "CVSS": {
            "nvd": {
                "V2Vector": "AV:L/AC:L/Au:N/C:P/I:P/A:P",
                "V3Vector": "CVSS:3.1/AV:L/AC:L/PR:H/UI:N/S:C/C:L/I:L/A:L",
                "V2Score": 4.6,
                "V3Score": 5.7
            },
            "redhat": {
                "V3Vector": "CVSS:3.1/AV:L/AC:L/PR:H/UI:N/S:C/C:L/I:L/A:L",
                "V3Score": 5.7
            }
        },
        "References": [
        "https://bugs.launchpad.net/bugs/1899193",
        "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-27350",
        "https://security.netapp.com/advisory/ntap-20210108-0005/",
        "https://usn.ubuntu.com/usn/usn-4667-1",
        "https://usn.ubuntu.com/usn/usn-4667-2",
        "https://www.debian.org/security/2020/dsa-4808"
        ],
        "PublishedDate": "2020-12-10T04:15:00Z",
        "LastModifiedDate": "2021-01-08T12:15:00Z"
    }]}]
    `
