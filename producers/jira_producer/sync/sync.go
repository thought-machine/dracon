package sync

import (
	// "fmt"

	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	"github.com/thought-machine/dracon/common/jira/config"
	"github.com/thought-machine/dracon/common/jira/document"
	"github.com/thought-machine/dracon/pkg/enrichment"
	"github.com/thought-machine/dracon/pkg/enrichment/db"

	jira "github.com/andygrunwald/go-jira"
)

//GetHash  if a hash field is present in the configuration, use that
// likely to be more acurrate than calculating it
func GetHash(issue jira.Issue, config config.Config) (string, error) {
	var jiraHashField, jiraHashFieldType string
	for _, mapping := range config.Mappings {
		if mapping.DraconField == "hash" {
			jiraHashField = mapping.JiraField
			jiraHashFieldType = mapping.FieldType
		}
	}
	if jiraHashField == "" {
		return "", errors.New("Config is missing a jira field mapping for hash")
	}
	if jiraHashFieldType != "single-value" {
		return "", errors.New("Hash must be single value, each vulnerability has one hash")
	}
	// happy path is single-value
	hash, err := issue.Fields.Unknowns.String(jiraHashField)
	if err != nil {
		return "", err
	}
	return hash, nil
}

// CalculateHash assumes the Jira Ticket doesn't have a custom hash field or the field is not mapped in the config
// tries to gather the info from the rest of the ticket, last resort method of calculating hashes.
// quite inacurrate
func CalculateHash(issue jira.Issue) string {
	description := issue.Fields.Description
	rexes := map[string]*regexp.Regexp{
		"source":         regexp.MustCompile(`(source:\s*(?P<source>\S+)\s)`),
		"target":         regexp.MustCompile(`(target:\s*(?P<target>\S+)\s)`),
		"vulnType":       regexp.MustCompile(`(type:\s*(?P<type>\w\s*)+)(\r\n|\r|\n)`),
		"severityText":   regexp.MustCompile(`(severity_text:\s*(?P<severity>.+)(\r\n|\r|\n))`),
		"cvss":           regexp.MustCompile(`(cvss:\s*(?P<cvss>[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?)(\r\n|\r|\n))`),
		"confidenceText": regexp.MustCompile(`(confidence_text:\s*(?P<confidence>\w+)(\r\n|\r|\n))`),
		"scanStartTime":  regexp.MustCompile(`(scan_start_time:\s*(.+)(\r\n|\r|\n))`),
		"scan_id":        regexp.MustCompile(`(scan_id:\s*(.+)(\r\n|\r|\n))`),
		"tool_name":      regexp.MustCompile(`(tool_name:\s*(.+)(\r\n|\r|\n))`),
		"first_found":    regexp.MustCompile(`(first_found:\s*(.+)(\r\n|\r|\n))`),
		"false_positive": regexp.MustCompile(`(false_positive:\s*(.+)(\r\n|\r|\n))`),
		"jira_code_tag":  regexp.MustCompile(`({code:?})`),
		"hash":           regexp.MustCompile(`(hash:\s*(?P<hash>\w+)(\r\n|\r|\n))`),
	}

	if len(namedGroupMatching(rexes["hash"], description)["hash"]) > 0 {
		// lucky break, if the consumer config has included the original hash and
		// no human removed it
		return namedGroupMatching(rexes["hash"], description)["hash"]
	}
	// else try and match what exists

	cvssVal, err := strconv.ParseFloat(namedGroupMatching(rexes["cvss"], description)["cvss"], 64)
	if err != nil {
		log.Println("Could not extract CVSS from the description")
	}

	desc := description
	for _, regex := range rexes {
		desc = regex.ReplaceAllLiteralString(desc, "")
	}
	draconIssue := &v1.Issue{
		Target:      namedGroupMatching(rexes["target"], description)["target"],
		Type:        namedGroupMatching(rexes["vulnType"], description)["type"],
		Title:       issue.Fields.Summary,
		Severity:    document.TextToSeverity(namedGroupMatching(rexes["severityText"], description)["severity"]),
		Cvss:        cvssVal,
		Confidence:  document.TextToConfidence(namedGroupMatching(rexes["confidenceText"], description)["confidence"]),
		Description: strings.TrimSpace(desc),
		Source:      namedGroupMatching(rexes["source"], description)["source"],
	}

	hash := enrichment.GetHash(draconIssue)
	return hash
}

func createIssue(issue *v1.Issue, db db.EnrichDatabase) {
	// create issue
	log.Println("DEBUG CREATING ISSUE THIS ASSUMES LOCAL RUN")

	enrichedIssue := enrichment.NewEnrichedIssue(issue)
	err := db.CreateIssue(context.Background(), enrichedIssue)
	if err != nil {
		log.Println(err)
	}
}

// UpdateDB updates the dracon vuln db based on the following rules
// DraconStatus: Resolved -> Delete  This allows for regression to be detected again
// DraconStatus: FalsePositive -> Update to set FalsePositive flag
// DraconStatus: Duplicate -> Do Nothing
func UpdateDB(hash string, config config.Config, issue jira.Issue, dryRun bool, db db.EnrichDatabase) {
	draconStatus := ""
	for _, mapping := range config.SyncMappings {
		if issue.Fields.Status.Name == mapping.JiraStatus && issue.Fields.Resolution.Name == mapping.JiraResolution {
			draconStatus = mapping.DraconStatus
			if dryRun {
				log.Println("Issue hash: " + hash + " would set to DraconStatus: " + draconStatus)
				return
			}
			break
		}
	}
	draconIssue, err := db.GetIssueByHash(hash)
	if err != nil {
		log.Printf("Could not get issue by hash %s", err)
		if errors.Is(err, sql.ErrNoRows) {
			log.Print("Issue " + hash + " was not found in vulndb, skipping") // log

			// for local dev only
			// createIssue(&v1.Issue{
			// 	Target:     "./plz-out/gen/third_party/java/com/google/protobuf/nano/protobuf-javanano_src.jar",
			// 	Type:       "Vulnerable Dependency",
			// 	Cvss:       8.8,
			// 	Confidence: document.TextToConfidence("Medium"),
			// 	Severity:   document.TextToSeverity("Significant / Large"),
			// 	Title:      issue.Fields.Summary,
			// }, connStr)
			// draconIssue, _ = db.GetIssueByHash(hash)
		}
	}
	switch draconStatus {
	case "FalsePositive":
		draconIssue.FalsePositive = true
		if err = db.UpdateIssue(context.Background(), draconIssue); err != nil {
			log.Printf("could not mark issue %s as false positive: error: %s\n", hash, err)
		}
	case "Resolved":
		log.Printf("deleting %s\n", hash)
		if err := db.DeleteIssueByHash(hash); err != nil {
			log.Printf("Could not delete issue %s from db, error: %s\n", hash, err) // log
		}
	case "Duplicate":
		fmt.Printf("issue %s already exists in db and is a duplicate, will keep ignoring\n", hash)
	}
}

// given a regexp with a named group e.g. "(?P<group> a named group)"
// returns a string map of the matches mapped to their group names
func namedGroupMatching(reg *regexp.Regexp, str string) map[string]string {
	match := reg.FindStringSubmatch(str)
	result := make(map[string]string)
	if len(match) == len(reg.SubexpNames()) {
		for i, name := range reg.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
	}
	return result
}
