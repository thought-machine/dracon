package types

import (
	"encoding/json"
)

type SafetyIssue struct {
	Name              string
	VersionConstraint string
	CurrentVersion    string
	Description       string
}

//read semi-unstructured safety json into struct
func (i *SafetyIssue) UnmarshalJSON(data []byte) error {

	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	i.Name, _ = v[0].(string)
	i.VersionConstraint, _ = v[1].(string)
	i.CurrentVersion = v[2].(string)
	i.Description = v[3].(string)
	return nil
}
