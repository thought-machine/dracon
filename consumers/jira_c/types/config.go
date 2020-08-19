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
	IssueFields  map[string][]string `yaml:"issueFields,omitempty"`
	CustomFields []CustomField       `yaml:"customFields,omitempty"`
}

type Mappings struct {
	DraconField string `yaml:"draconField"`
	JiraField   string `yaml:"jiraField"`
	FieldType   string `yaml:"fieldType"`
}
