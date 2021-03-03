package jira

import (
	"io/ioutil"
	"log"

	"github.com/andygrunwald/go-jira"

	"github.com/thought-machine/dracon/common/jira/config"
)

type client struct {
	JiraClient    *jira.Client
	DryRunMode    bool
	Config        config.Config
	DefaultFields defaultJiraFields
}

// NewClient returns a client containing the authentication details and the configuration settings
func NewClient(user, token, url string, dryRun bool, config config.Config) *client {
	return &client{
		JiraClient:    authJiraClient(user, token, url),
		DryRunMode:    dryRun,
		Config:        config,
		DefaultFields: getDefaultFields(config),
	}
}

// authJiraClient authenticates the client with the given Username, API token, and URL domain
func authJiraClient(user, token, url string) *jira.Client {
	tp := jira.BasicAuthTransport{
		Username: user,
		Password: token,
	}
	jiraClient, err := jira.NewClient(tp.Client(), url)
	if err != nil {
		log.Fatalf("Unable to contact Jira: %s", err)
	}
	return jiraClient
}

// assembleIssue parses the Dracon message and serializes it into a Jira Issue object
func (client client) assembleIssue(draconResult map[string]string) *jira.Issue {
	// Mappings the Dracon Result fields to their corresponding Jira fields specified in the configuration
	customFields := client.DefaultFields.CustomFields.Clone()
	for _, m := range client.Config.Mappings {
		customFields[m.JiraField] = makeCustomField(m.FieldType, []string{draconResult[m.DraconField]})
	}
	return &jira.Issue{
		Fields: &jira.IssueFields{
			Project:         client.DefaultFields.Project,
			Type:            client.DefaultFields.IssueType,
			Description:     makeDescription(draconResult, client.Config.DescriptionExtras),
			Summary:         makeSummary(draconResult),
			Components:      client.DefaultFields.Components,
			AffectsVersions: client.DefaultFields.AffectsVersions,
			Labels:          client.DefaultFields.Labels,
			Unknowns:        customFields,
		},
	}
}

// CreateIssue creates a new issue in Jira
func (client client) CreateIssue(draconResult map[string]string) error {
	issue := client.assembleIssue(draconResult)

	if client.DryRunMode {
		log.Printf("Dry run mode. The following issue would have been created: '%s'", issue.Fields.Summary)
		return nil
	}

	ri, resp, err := client.JiraClient.Issue.Create(issue)
	if err != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Error occurred posting to Jira. Response body:\n%s", body)
		return err
	}
	log.Printf("Created Jira Issue ID %s. jira_key=%s", ri.ID, string(ri.Key))
	return nil
}

// SearchByJQL searches jira instance by JQL and returns results with history
func (client client) SearchByJQL(jql string) ([]jira.Issue, error) {
	var results []jira.Issue
	startAt := 0
	maxresults := 100
	expand := "names,schema,operations,editmeta,changelog,renderedFields"
	issues, response, err := client.JiraClient.Issue.Search(jql, &jira.SearchOptions{Expand: expand, StartAt: startAt, MaxResults: maxresults}) //maxresults is capped to 100 by attlasian
	if err != nil {
		log.Print(response)
		return nil, err
	}
	results = append(results, issues...)
	startAt = len(results)
	log.Print("The query returned ", response.Total, " results")
	for len(results) < response.Total {
		issues, response, err = client.JiraClient.Issue.Search(jql, &jira.SearchOptions{Expand: expand, StartAt: startAt, MaxResults: maxresults}) //maxresults is capped to 100 by attlasian
		if err != nil {
			log.Print(response)
			return nil, err
		}
		results = append(results, issues...)
		startAt = len(results)
	}
	return results, nil
}
