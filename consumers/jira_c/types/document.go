package document

import (
	"time"
)

type Document struct {
	ScanStartTime  time.Time `json:"scan_start_time"`
	ScanID         string    `json:"scan_id"`
	ToolName       string    `json:"tool_name"`
	Source         string    `json:"source"`
	Target         string    `json:"target"`
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	SeverityText   string    `json:"severity_text"`
	CVSS           string    `json:"cvss"`
	ConfidenceText string    `json:"confidence_text"`
	Description    string    `json:"description"`
	FirstFound     time.Time `json:"first_found"`
	Count          string    `json:"count"`
	FalsePositive  string    `json:"false_positive"`
	// Severity   v1.Severity   `json:"severity"`
	// Confidence v1.Confidence `json:"confidence"`
}
