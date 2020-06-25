package types

// NpmAuditOut represents the output of an npm-audit run that we care about
type NpmAuditOut struct {
	Advisories map[int]NpmAuditAdvisories `json:"advisories"`
}

// NpmAuditAdvisories represents an npm-audit advisory section that we care about
type NpmAuditAdvisories struct {
	Title              string `json:"title"`
	ModuleName         string `json:"module_name"`
	VulnerableVersions string `json:"vulnerable_versions"`
	Overview           string `json:"overview"`
	Recommendation     string `json:"recommendation"`
	References         string `json:"references"`
	Severity           string `json:"severity"`
	CWE                string `json:"cwe"`
	URL                string `json:"url"`
}
