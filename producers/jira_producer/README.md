Jira Synchronizer
===
This "Producer"  for lack of a better term will read from the Jira instance it gets pointed to and synchronise triaged vulnerabilities with the internal Dracon DB, this allows Dracon to understand when a vulnerability has been resolved so it can remove it from the list of duplicates, it also allows for marking vulnerabilities as false positives.

A cronjob has been created to make this synchronisation periodic. You can find a template for it under //examples/cronjobs/jira-sync-cronjob.yaml
This component utilises the default Jira config.yaml that the Jira consumer uses.

To run this individually:
``` plz run //producers/jira_producer:sync_tickets  -- --user="<jira email>" --token="<jira api token>" --jira="<>" --query='<jql>' --config /path/to/config.yaml --dbcon "<db connection string>"
```
