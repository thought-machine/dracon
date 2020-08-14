package gojira

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleGoJiraClient = &goJiraClient{
	JiraClient: authJiraClient("test_user", "test_token", "test_url"),
	DryRunMode: true,
	Configs: config{
		DefaultValues: defaultValues{
			IssueFields: map[string][]string{
				"project":         []string{"TOY"},
				"issueType":       []string{"Vulnerability"},
				"components":      []string{"c1", "c2", "c3"},
				"affectsVersions": []string{"V1", "V2"},
			},
			CustomFields: []customField{{
				ID:        "customfield_1000",
				FieldType: "multi-value",
				Values:    []string{"foo", "bar"},
			}},
		},
		Mappings: []mappings{{
			DraconField: "cvss",
			JiraField:   "customfield_1001",
			FieldType:   "float",
		}},
		DescriptionExtras: []string{"target", "tool_name"},
	},
}

var sampleMessage = `{"scan_start_time":"0001-01-01T00:00:00Z","scan_id":"babbb83-4627-41c6-8ba0-70ee866290e9","tool_name":"spotbugs","source":"//foo/bar:baz","target":"//foo1/bar1:baz2","type":"test type","title":"Unit Test Title","severity_text":"Info","cvss":"0.000","confidence_text":"Info","description":"this is a test description","first_found":"0001-01-01T00:00:00Z","false_positive":"true"}`

func TestNewGoJiraClient(t *testing.T) {
	goJiraClient := NewGoJiraClient("test_user", "test_token", "test_url", false)
	assert.NotEmpty(t, goJiraClient)
}

func TestAuthJiraClient(t *testing.T) {
	client := authJiraClient("test_user", "test_token", "test_url")
	assert.NotEmpty(t, client)
}

func TestAssembleIssue(t *testing.T) {
	issue := sampleGoJiraClient.assembleIssue(sampleMessage)
	assert.NotEmpty(t, issue)
}

func TestCreateIssue(t *testing.T) {
	err := sampleGoJiraClient.CreateIssue(sampleMessage)
	assert.NoError(t, err)
}
