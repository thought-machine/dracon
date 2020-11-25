package config

// Config contains all the data parsed from the conf.yaml file
type Config struct {
	DefaultValues     DefaultValues              `yaml:"defaultValues"`
	Mappings          []Mappings                 `yaml:"mappings"`
	DescriptionExtras []string                   `yaml:"addToDescription"`
	SyncMappings      []JiraToDraconVulnMappings `yaml:"syncMappings"`
}

type CustomField struct {
	ID        string   `yaml:"id"`
	FieldType string   `yaml:"fieldType"`
	Values    []string `yaml:"values"`
}

type DefaultValues struct {
	Project         string        `yaml:"project"`
	IssueType       string        `yaml:"issueType"`
	Components      []string      `yaml:"components"`
	AffectsVersions []string      `yaml:"affectsVersions"`
	Labels          []string      `yaml:"labels,omitempty"`
	CustomFields    []CustomField `yaml:"customFields,omitempty"`
}

type Mappings struct {
	DraconField string `yaml:"draconField"`
	JiraField   string `yaml:"jiraField"`
	FieldType   string `yaml:"fieldType"`
}

// JiraToDraconVulnMappings used by the sync utiity,
// this Mapping matches DraconStatus-es to combinations of JiraStatus and JiraResolution, look in the sample config file for examples
// supported DraconStatus values:
// * FalsePositive <-- will set the issue's FalsePositive flag to True
// * Duplicate <-- if the issue already exists in the database, will do nothing, otherwise will insert a new one
// * Resolved <-- will _REMOVE_ the finding from the database
// JiraStatus will be matched as a string
// JiraResolution will be matched as a string
type JiraToDraconVulnMappings struct {
	JiraStatus     string `yaml:"jiraStatus"`
	JiraResolution string `yaml:"jiraResolution"`
	DraconStatus   string `yaml:"draconStatus"`
}
