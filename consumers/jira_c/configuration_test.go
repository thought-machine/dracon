package configuration

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"consumers/jira_c/types/config"
)

var sampleConfig = config.Config{
	DefaultValues: config.DefaultValues{
		IssueFields: map[string][]string{
			"project":         []string{"TOY"},
			"issueType":       []string{"Vulnerability"},
			"components":      []string{"c1", "c2", "c3"},
			"affectsVersions": []string{"V1", "V2"},
		},
		CustomFields: []config.CustomField{{
			ID:        "customfield_10000",
			FieldType: "multi-value",
			Values:    []string{"foo", "bar"},
		}},
	},
	Mappings: []config.Mappings{{
		DraconField: "cvss",
		JiraField:   "customfield_10001",
		FieldType:   "float",
	}},
	DescriptionExtras: []string{"target", "tool_name"},
}

func TestGetConfig(t *testing.T) {
	testConfig := `
defaultValues:
  issueFields:
    project: ['TOY']
    issueType: ['Vulnerability']
    components: ['c1', 'c2', 'c3']
    affectsVersions: ['V1', 'V2']

  customFields:
    - id: 'customfield_10000'
      fieldType: multi-value
      values: ['foo', 'bar']

addToDescription:
  - target
  - tool_name

mappings:
  - draconField: cvss
    jiraField: customfield_10001
    fieldType: float
`

	reader := strings.NewReader(testConfig)
	res, err := GetConfig(reader)

	assert.NoError(t, err)
	assert.EqualValues(t, res, sampleConfig)
}
