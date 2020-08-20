package main

import (
	"flag"
	"log"
	"os"

	"github.com/thought-machine/dracon/consumers"

	"consumers/jira_c/config/config"
	"consumers/jira_c/jira"
	"consumers/jira_c/utils"
)

const (
	// EnvJiraUser the Jira Username for the authentication (user@domain.com)
	EnvJiraUser = "DRACON_JIRA_USER"
	// EnvJiraToken the Jira API token for the authentication
	EnvJiraToken = "DRACON_JIRA_TOKEN"
	// EnvJiraURL the domain to scrape
	EnvJiraURL = "DRACON_JIRA_URL"
	// EnvConfigPath the path towards the config.yaml file
	EnvConfigPath = "DRACON_JIRA_CONFIG_PATH"
)

var (
	authUser          string
	authToken         string
	jiraURL           string
	dryRunMode        bool
	allowDuplicates   bool
	allowFP           bool
	severityThreshold int
)

func init() {
	authUser = os.Getenv(EnvJiraUser)
	authToken = os.Getenv(EnvJiraToken)
	jiraURL = os.Getenv(EnvJiraURL)
	flag.BoolVar(&dryRunMode, "dry-run", false, "Dry run. Tickets will not be created.")
	flag.BoolVar(&allowDuplicates, "allow-duplicates", false, "Allow duplicate issues to be created.")
	flag.BoolVar(&allowFP, "allow-fp", false, "Allow issues tagged as 'false positive' to be created.")
	flag.IntVar(&severityThreshold, "severity-threshold", 3, "Only issues equal or above this threshold will get processed. Must be one of: {0: Info, 1: Minor / Localized, 2: Moderate / Limited, 3: Significant / Large, 4: Extensive / Widespread}")
}

func main() {
	// Parse consumer flags
	if err := consumers.ParseFlags(); err != nil {
		log.Fatal("Could not parse flags: ", err)
	}

	// Parse config.yaml
	file, err := os.Open(os.Getenv(EnvConfigPath))
	if err != nil {
		log.Fatalf("Could not open config file:   #%v ", err)
	}
	config, err := config.New(file)
	if err != nil {
		log.Fatalf("Could not parse config file: %v", err)
	}

	// Authenticate Jira client
	apiClient := jira.NewClient(authUser, authToken, jiraURL, dryRunMode, config)

	// Parse Dracon results
	draconResults, discardedIssues, err := utils.ProcessMessages(allowDuplicates, allowFP, severityThreshold)
	if err != nil {
		log.Fatalf("Could not process messages: %s", err)
	}

	// Create issues in Jira
	createdIssues := 0
	failedIssues := 0
	for _, result := range draconResults {
		if err := apiClient.CreateIssue(result); err != nil {
			failedIssues++
		} else {
			createdIssues++
		}
	}

	// Output metrics
	log.Printf("%d Issues have been discarded as duplicates or false positives\n", discardedIssues)
	log.Printf("Dracon results: %d; Created issues: %d; Issues failed to create: %d\n", len(draconResults), createdIssues, failedIssues)
	if failedIssues > 0 {
		os.Exit(1)
	}
}
