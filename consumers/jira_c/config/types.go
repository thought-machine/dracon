package config

// Config contains all the data parsed from the conf.yaml file
type Config struct {
	DefaultValues     DefaultValues `yaml:"defaultValues"`
	Mappings          []Mappings    `yaml:"mappings"`
	DescriptionExtras []string      `yaml:"addToDescription"`
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
