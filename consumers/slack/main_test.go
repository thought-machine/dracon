package consumers

import (
	"testing"

	v1 "github.com/thought-machine/dracon/plz-out/gen/pkg/genproto/v1"
)

func TestPushMetrics(t *testing.T) {
	expectedIssue := &v1.Issue{
		Target:      "/tmp/source/foo.go:33",
		Type:        "G304",
		Title:       "ioutil.ReadFile(path)",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_HIGH,
		Description: "Potential file inclusion via variable",
	}
	var response = v1.LaunchToolResponse[1]
	response[0].Issues = []*v1.Issue{
		&v1.Issue{
			Target:      "/dracon/source/foobar",
			Title:       "/dracon/source/barfoo",
			Description: "/dracon/source/example.yaml",
		},
	}
	`setup 1 messages with 1 result and make sure metrics tries to push the correct message
	mock push([]byte)`
	t.Fail("missing test")
}

func TestPush(t *testing.T) {
	`setup 2 messages with 2 results each and ensure their json format gets pushed`
	t.Fail("missing test")
}
