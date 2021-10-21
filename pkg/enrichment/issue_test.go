package enrichment

import (
	"fmt"
	"strings"
	"testing"

	v1 "github.com/thought-machine/dracon/api/proto/v1"

	"github.com/stretchr/testify/assert"
)

func TestGetHash(t *testing.T) {
	expectedIssues := &v1.Issue{
		Target:     "pkg:golang/github.com/coreos/etcd@0.5.0-alpha.5",
		Type:       "Vulnerable Dependency",
		Title:      "[CVE-2018-1099]  Improper Input Validation",
		Source:     "git.foo.com/repo.git?ref=aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		Severity:   v1.Severity_SEVERITY_MEDIUM,
		Cvss:       5.5,
		Confidence: v1.Confidence_CONFIDENCE_HIGH,
		Description: fmt.Sprintf("CVSS Score: %v\nCvssVector: %s\nCve: %s\nCwe: %s\nReference: %s\n",
			"5.5", "CVSS:3.0/AV:L/AC:L/PR:L/UI:N/S:U/C:N/I:H/A:N", "CVE-2018-1099",
			"", "https://ossindex.sonatype.org/vuln/8a190129-526c-4ee0-b663-92f38139c165"),
		Cve: "123-321",
	}
	assert.Equal(t, GetHash(expectedIssues), "ccc217a4c2fd348bc5c6c4d73ad4311a")

	expectedIssues.Source = strings.NewReplacer("aa", "bc").Replace(expectedIssues.Source)
	// Test for regression on Bug where we would calculate ?ref=<> value for enrichment
	assert.Equal(t, GetHash(expectedIssues), "ccc217a4c2fd348bc5c6c4d73ad4311a")

	expectedIssues.Source = strings.NewReplacer("git.foo.com/repo.git", "https://example.com/foo/bar").Replace(expectedIssues.Source)
	assert.NotEqual(t, GetHash(expectedIssues), "3c73dcc2f7c647a4ff460249074a8d50")
}
