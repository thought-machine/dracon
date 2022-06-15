package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/pkg/putil"
)

func TestMain(m *testing.M) {
	for i, arg := range os.Args {
		if arg == "--test_execution" {
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
			main()
			return
		}
	}

	m.Run()
}

func TestMainFn(t *testing.T) {
	type testCase struct {
		inputFileName string
		expErr        bool
		expStdout     string
		expIssues     []*v1.Issue
	}

	testCases := []testCase{
		{
			inputFileName: "no_findings.json",
			expErr:        false,
			expStdout:     `wrote 0 issues from to ./producers/securityhub/out.pb`,
		},
		{
			inputFileName: "empty_finding.json",
			expErr:        false,
			expStdout:     `wrote 1 issues from to ./producers/securityhub/out.pb`,
			expIssues: []*v1.Issue{
				{
					Target:      "",
					Type:        "",
					Title:       "",
					Severity:    0,
					Cvss:        0.0,
					Confidence:  2,
					Description: "",
					Source:      "unknown",
					Cve:         "",
				},
			},
		},
		{
			inputFileName: "securityhub_finding.json",
			expErr:        false,
			expStdout:     `wrote 1 issues from to ./producers/securityhub/out.pb`,
			expIssues: []*v1.Issue{
				{
					Target:      "arn:aws:ec2:eu-west-2:123456789:security-group/sg-01ef4a31dbea6f188cfbf",
					Type:        "Software and Configuration Checks/Industry and Regulatory Standards/CIS AWS Foundations Benchmark",
					Title:       "4.3 Ensure the default security group of every VPC restricts all traffic",
					Severity:    3,
					Cvss:        0,
					Confidence:  2,
					Description: "A VPC comes with a default security group whose initial settings deny all inbound traffic, allow all outbound traffic, and allow all traffic between instances assigned to the security group. If you don't specify a security group when you launch an instance, the instance is automatically assigned to this default security group. It is recommended that the default security group restrict all traffic.",
					Source:      "unknown",
					Cve:         "",
				},
			},
		},
		{
			inputFileName: "inspector_finding.json",
			expErr:        false,
			expStdout:     `wrote 1 issues from to ./producers/securityhub/out.pb`,
			expIssues: []*v1.Issue{
				{
					Target:      "ami-053269b2b68617f7c",
					Type:        "Software and Configuration Checks/Vulnerabilities/CVE",
					Title:       "CVE-2022-25315 - expat",
					Severity:    4,
					Cvss:        9.8,
					Confidence:  2,
					Description: "An integer overflow was found in expat. The issue occurs in storeRawNames() by abusing the m_buffer expansion logic to allow allocations very close to INT_MAX and out-of-bounds heap writes. This flaw can cause a denial of service or potentially arbitrary code execution.",
					Source:      "unknown",
					Cve:         "CVE-2022-25315",
				},
			},
		},
		{
			inputFileName: "bad-input",
			expErr:        true,
			expStdout:     `open ./producers/securityhub/testcases/bad-input: no such file or directory`,
		},
		{
			inputFileName: "unparseable.json",
			expErr:        true,
			expStdout:     `invalid character '.' looking for beginning of value`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.inputFileName, func(t *testing.T) {
			cmd := runProducerMain(`-in=./producers/securityhub/testcases/`+tc.inputFileName, `-out=./producers/securityhub/out.pb`)
			stdoutStderr, err := cmd.CombinedOutput()

			if err != nil && !tc.expErr {
				assert.Fail(t, "unexpected err executing producer with input file '%s'. err: '%s", tc.inputFileName, err)
			} else if tc.expErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}

			expMsg := "expected output from '%s' to end with '%s'. got: '%s'"
			stdOut := strings.TrimSpace(string(stdoutStderr))
			suffixCheck := strings.HasSuffix(stdOut, tc.expStdout)

			assert.True(t, suffixCheck, expMsg, tc.inputFileName, tc.expStdout, stdOut)

			ltr, err := putil.LoadToolResponse("./producers/securityhub")
			assert.NoError(t, err)

			if tc.expIssues != nil {
				assert.EqualValues(t, ltr[0].Issues, tc.expIssues)
			} else if !tc.expErr {
				assert.Empty(t, ltr[0].Issues)
			}

			os.Remove("./producers/securityhub/out.pb")
		})
	}
}

func runProducerMain(args ...string) *exec.Cmd {
	return exec.Command(os.Args[0], append([]string{"--test_execution"}, args...)...)
}
