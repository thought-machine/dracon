package types

import (
	"encoding/json"
	"regexp"
)

// SafetyIssue represents a pip-safety finding
type SafetyIssue struct {
	Name              string
	VersionConstraint string
	CurrentVersion    string
	Description       string
	Cve               string
}

//UnmarshalJSON is autocalled on any JSON unmarshalling into the SafetyIssue struct
// read semi-unstructured safety json into struct
func (i *SafetyIssue) UnmarshalJSON(data []byte) error {

	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	re := regexp.MustCompile(`CVE-\d{4}-\d+`)
	match := re.FindStringSubmatch(string(data))
	cve := ""
	if len(match) > 0 {
		cve = match[0]
	}
	i.Name, _ = v[0].(string)
	i.VersionConstraint, _ = v[1].(string)
	i.CurrentVersion, _ = v[2].(string)
	i.Description, _ = v[3].(string)
	i.Cve = cve
	return nil
}
