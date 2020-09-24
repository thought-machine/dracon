package types

import (
	"time"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
)

type FullDocument struct {
	ScanStartTime time.Time     `json:"scan_start_time"`
	ScanID        string        `json:"scan_id"`
	ToolName      string        `json:"tool_name"`
	Source        string        `json:"source"`
	Target        string        `json:"target"`
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	Severity      v1.Severity   `json:"severity"`
	CVSS          float64       `json:"cvss"`
	Confidence    v1.Confidence `json:"confidence"`
	Description   string        `json:"description"`
	FirstFound    time.Time     `json:"first_found"`
	Count         uint64        `json:"count"`
	FalsePositive bool          `json:"false_positive"`
}
