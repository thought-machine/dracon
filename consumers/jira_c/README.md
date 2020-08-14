# Jira Consumer

The Jira Consumer allows you to publish Vulnerabiltiy Issues to your organisation's Jira workspace straight from the Dracon pipeline results. Result fields such as 'target', 'cvss', 'severity', can either be written in the description or mapped to specific customfields used by your workspace.

## Setting Up

### Matching your Jira Workspace needs
All you need to edit is the config.yaml file so it can match your organisation's needs. You should be thinking about what Jira Issue Fields you expect the tool to fill in (such as Project Key, Issue Type, and even customfields used by your workspace).

 The configuration has three components:
* defaultValues: Here you can specify fields with default values such as: Project Key, Issue Type, or even specific customfields that your workspace uses.
* addToDescription: Here you can specify what fields from the Dracon Results you want written in the Issue's description.
* mappings: In case your organisation has specific fields for CVSS, Severity, etc, you can also map the dracon results straight to these fields. If not, you can just add them to the description (see the point above).

There are more instructions in the config.yaml file for how to format those three components.

### Authentication through the Jira API
Authentication details for the Jira API are passed as environment variables. These can be set up in the pipeline.yaml file or in the Dockerfile.

DRACON_JIRA_USER="user@email.com" 
DRACON_JIRA_TOKEN="your api token"
DRACON_JIRA_URL="domain your jira workspace is hosted on"


## Testing locally
The following command will test that the app and configuration is working correctly.
`plz test //consumers/jira_c/...`


## Running as part of the Dracon Pipeline
// TO COMPLETE

The following arguments can be specified:
```
   --dryRun              For debugging. Tickets will not be created
   --raw                 If the non-enriched results should be used
   --allowFP             Allows issues tagged as 'false positive to be created.
   --allowDuplicates     Allows duplicate issues to be created.
   --severityThreshold   Only issues equal or above this threshold will get processed. {0: All, 1: Minor, 2: Moderate, 3: High, 4: Critical,   Default: 4}
```
