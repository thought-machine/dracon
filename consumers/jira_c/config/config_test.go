package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleConfig = Config{
	DefaultValues: DefaultValues{
		Project:         "TOY",
		IssueType:       "Vulnerability",
		Components:      []string{"c1", "c2", "c3"},
		AffectsVersions: []string{"V1", "V2"},
		Labels:          []string(nil),
		CustomFields: []CustomField{{
			ID:        "customfield_10000",
			FieldType: "multi-value",
			Values:    []string{"foo", "bar"},
		}},
	},
	Mappings: []Mappings{{
		DraconField: "cvss",
		JiraField:   "customfield_10001",
		FieldType:   "float",
	}},
	DescriptionExtras: []string{"target", "tool_name"},
}

func TestGetConfig(t *testing.T) {
	testConfig := `
defaultValues:
  project: 'TOY'
  issueType: 'Vulnerability'
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
	res, err := New(reader)

	assert.NoError(t, err)
	assert.EqualValues(t, res, sampleConfig)
}
