package types

// ZapOut represents the output of a zap scan
type ZapOut struct {
	Version   string     `json:"@version"`
	Generated string     `json:"@generated"`
	Site      []ZapSites `json:"site"`
}

// ZapSite represents a zap site section
type ZapSites struct {
	Name   string      `json:"@name"`
	Host   string      `json:"@host"`
	Port   string      `json:"@port"`
	Ssl    string      `json:"@ssl"`
	Alerts []ZapAlerts `json:"alerts"`
}

// ZapInstances represents a zap occurrence for a specific alert
type ZapInstances struct {
	Uri    string `json:"uri"`
	Method string `json:"method"`
}

// ZapAlert represents a zap vulnerability
type ZapAlerts struct {
	PluginId    string         `json:"Id"`
	AlertRef    string         `json:"alertRef"`
	Alert       string         `json:"alert"`
	Name        string         `json:"name"`
	RiskCode    string         `json:"riskcode"`
	Confidence  string         `json:"confidence"`
	RiskDesc    string         `json:"riskdesc"`
	Description string         `json:"desc"`
	Instances   []ZapInstances `json:"instances"`
	Count       string         `json:"count"`
	Solution    string         `json:"solution"`
	OtherInfo   string         `json:"otherinfo"`
	Reference   string         `json:"reference"`
	CweId       string         `json:"cweid"`
	WascId      string         `json:"wascid"`
	SourceId    string         `json:"sourceid"`
}
