package sync

import (
	"context"
	"fmt"
	"testing"

	jira "github.com/andygrunwald/go-jira"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/common/jira/config"
	database_mock "github.com/thought-machine/dracon/pkg/enrichment/db/mock"
	"github.com/trivago/tgo/tcontainer"
	// "github.com/thought-machine/dracon/producers/jira_producer/sync"
)

var (
	jiraStatus     = ""
	jiraResolution = ""
	draconStatus   = ""
	sampleOut      = "b78a94da3b999f0244240e78a01a66d0"
	jiraHashField  = "customfield_12345"
	sampleConfig   = config.Config{
		DefaultValues: config.DefaultValues{
			Project:         "TOY",
			IssueType:       "Vulnerability",
			Components:      []string{"c1", "c2", "c3"},
			AffectsVersions: []string{"V1", "V2"},
			Labels:          []string(nil),
			CustomFields: []config.CustomField{{
				ID:        "customfield_10000",
				FieldType: "multi-value",
				Values:    []string{"foo", "bar"},
			}},
		},
		Mappings: []config.Mappings{
			{
				DraconField: "hash",
				JiraField:   jiraHashField,
				FieldType:   "single-value",
			},
		},
		DescriptionExtras: []string{"target", "tool_name"},
		SyncMappings: []config.JiraToDraconVulnMappings{
			{
				JiraStatus:     jiraStatus,
				JiraResolution: jiraResolution,
				DraconStatus:   draconStatus,
			},
		},
	}
	sampleDraconIssue = &v1.EnrichedIssue{
		RawIssue: &v1.Issue{
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce iaculis felis nisi, vel ultricies eros facilisis in.",
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Severity:    v1.Severity_SEVERITY_HIGH,
			Cvss:        8.88,
			Source:      "some.git.repo/path/to/code.git?ref=11111111111111111111111111111111111111111111111",
			Target:      "./plz-out/gen/third_party/java/foo/bar/baz/foo-bar.jar",
			Title:       "foo-bar.jar ",
			Type:        "Vulnerable Dependency",
		},
		Hash: sampleOut}
	issue = jira.Issue{
		Key: "FOO-1234",
		Fields: &jira.IssueFields{
			Status:     &jira.Status{Name: jiraStatus},
			Resolution: &jira.Resolution{Name: jiraResolution},
			Unknowns: tcontainer.MarshalMap{
				jiraHashField:      sampleOut,
				"customfield_1987": []interface{}{map[string]interface{}{"id": "1", "self": "http://example.com", "value": "foobar"}},
			},
			Description: `This issue was automatically generated by the Dracon security pipeline.\n
			*Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce iaculis felis nisi, vel ultricies eros facilisis in.*
			{code:}
			scan_start_time:   2020-11-11T11:11:11Z
			scan_id: dracon-asdfghjklpoiuyt
			tool_name: dependencyCheck
			source: some.git.repo/path/to/code.git?ref=11111111111111111111111111111111111111111111111
			target: ./plz-out/gen/third_party/java/foo/bar/baz/foo-bar.jar
			type:   Vulnerable Dependency
			severity_text:  Significant / Large
			cvss:  8.888
			confidence_text:  Medium
			first_found: 2020-11-11T11:11:11Z
			false_positive: false
			{code}`,
			Summary: "foo-bar.jar ./plz-out/gen/third_party/java/foo/bar/baz/foo-bar.jar"},
	}
)

func TestGetHash(t *testing.T) {
	// happy path
	hash, _ := GetHash(issue, sampleConfig)
	assert.Equal(t, hash, sampleOut)

	conf := sampleConfig
	// config doesn't have hash
	conf.Mappings[0].DraconField = ""
	hash, err := GetHash(issue, conf)
	assert.Equal(t, hash, "")
	assert.Equal(t, fmt.Sprintf("%s", err), "Config is missing a jira field mapping for hash")

	// config wrong field type
	conf.Mappings[0].DraconField = "hash"
	conf.Mappings[0].FieldType = "float"
	hash, err = GetHash(issue, conf)
	assert.Equal(t, hash, "")
	assert.Equal(t, fmt.Sprintf("%s", err), "Hash must be single value, each vulnerability has one hash")

}

func TestCalculateHash(t *testing.T) {
	assert.Equal(t, CalculateHash(issue), sampleOut)
}

func TestUpdateDBDryRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dryRunDatabase := database_mock.NewMockEnrichDatabase(ctrl)

	// Dry Run does not call anything on the database
	// following 3 lines are not required, their use is to make it explicit
	dryRunDatabase.EXPECT().GetIssueByHash(nil).Times(0)
	dryRunDatabase.EXPECT().UpdateIssue(nil, nil).Times(0)
	dryRunDatabase.EXPECT().DeleteIssueByHash(nil).Times(0)
	UpdateDB(sampleOut, sampleConfig, issue, true, dryRunDatabase)
}
func TestUpdateDBResolved(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sampleConfig.SyncMappings[0].JiraStatus = "foo"
	sampleConfig.SyncMappings[0].JiraResolution = "bar"
	sampleConfig.SyncMappings[0].DraconStatus = "Resolved"
	issue.Fields.Status.Name = "foo"
	issue.Fields.Resolution.Name = "bar"

	// Resolved Deletes
	resolvedDatabase := database_mock.NewMockEnrichDatabase(ctrl)
	resolvedDatabase.EXPECT().GetIssueByHash(sampleOut).Return(sampleDraconIssue, nil).Times(1)
	resolvedDatabase.EXPECT().DeleteIssueByHash(sampleOut).Return(nil).Times(1)

	UpdateDB(sampleOut, sampleConfig, issue, false, resolvedDatabase)
}
func TestUpdateDBFalsePositive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sampleConfig.SyncMappings[0].JiraStatus = "foo"
	sampleConfig.SyncMappings[0].JiraResolution = "bar"
	sampleConfig.SyncMappings[0].DraconStatus = "FalsePositive"
	issue.Fields.Status.Name = "foo"
	issue.Fields.Resolution.Name = "bar"

	// False Positive Updates
	updatedDatabase := database_mock.NewMockEnrichDatabase(ctrl)
	updatedDatabase.EXPECT().GetIssueByHash(sampleOut).Return(sampleDraconIssue, nil).Times(1)
	sampleDraconIssue.FalsePositive = true
	updatedDatabase.EXPECT().UpdateIssue(context.Background(), sampleDraconIssue).Return(nil).Times(1)
	UpdateDB(sampleOut, sampleConfig, issue, false, updatedDatabase)
}

func TestUpdateDBDuplicate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sampleConfig.SyncMappings[0].JiraStatus = "foo"
	sampleConfig.SyncMappings[0].JiraResolution = "bar"
	sampleConfig.SyncMappings[0].DraconStatus = "Duplicate"
	issue.Fields.Status.Name = "foo"
	issue.Fields.Resolution.Name = "bar"

	// Duplicate is a noop
	duplicateDatabase := database_mock.NewMockEnrichDatabase(ctrl)
	duplicateDatabase.EXPECT().GetIssueByHash(sampleOut).Return(sampleDraconIssue, nil).Times(1)
	duplicateDatabase.EXPECT().UpdateIssue(context.Background(), sampleDraconIssue).Return(nil).Times(0)
	UpdateDB(sampleOut, sampleConfig, issue, false, duplicateDatabase)
}
