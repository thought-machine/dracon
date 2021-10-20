package document

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
)

// NewRaw returns an []byte containing the parsed Document (dracon result) from the given raw issue
func NewRaw(scanStartTime time.Time, res *v1.LaunchToolResponse, iss *v1.Issue) ([]byte, error) {
	jBytes, err := json.Marshal(&Document{
		ScanStartTime:  scanStartTime,
		ScanID:         res.GetScanInfo().GetScanUuid(),
		ToolName:       res.GetToolName(),
		Source:         iss.GetSource(),
		Title:          iss.GetTitle(),
		Target:         iss.GetTarget(),
		Type:           iss.GetType(),
		SeverityText:   severtiyToText(iss.GetSeverity()),
		CVSS:           strconv.FormatFloat(iss.GetCvss(), 'f', 3, 64), // formatted as string
		ConfidenceText: confidenceToText(iss.GetConfidence()),
		Description:    iss.GetDescription(),
		FirstFound:     scanStartTime,
		Count:          "1",
		FalsePositive:  "false",
		Hash:           "",
		// The fields below are not used in this consumer. We use the text versions instead.
		// Severity:       iss.GetSeverity(),
		// Confidence:     iss.GetConfidence(),

	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

// NewEnriched returns an []byte containing the parsed Document (dracon result) from the given enriched issue
func NewEnriched(scanStartTime time.Time, res *v1.EnrichedLaunchToolResponse, iss *v1.EnrichedIssue) ([]byte, error) {
	firstSeenTime, _ := ptypes.Timestamp(iss.GetFirstSeen())
	jBytes, err := json.Marshal(&Document{
		ScanStartTime:  scanStartTime,
		ScanID:         res.GetOriginalResults().GetScanInfo().GetScanUuid(),
		ToolName:       res.GetOriginalResults().GetToolName(),
		Source:         iss.GetRawIssue().GetSource(),
		Title:          iss.GetRawIssue().GetTitle(),
		Target:         iss.GetRawIssue().GetTarget(),
		Type:           iss.GetRawIssue().GetType(),
		SeverityText:   severtiyToText(iss.GetRawIssue().GetSeverity()),
		CVSS:           strconv.FormatFloat(iss.GetRawIssue().GetCvss(), 'f', 3, 64), // formatted as string
		ConfidenceText: confidenceToText(iss.GetRawIssue().GetConfidence()),
		Description:    iss.GetRawIssue().GetDescription(),
		FirstFound:     firstSeenTime,
		Count:          strconv.Itoa(int(iss.GetCount())),          // formatted as string
		FalsePositive:  strconv.FormatBool(iss.GetFalsePositive()), // formatted as string
		Hash:           iss.GetHash(),
		// The fields below are not used in this consumer. We use the text versions instead.
		// Severity:       iss.GetRawIssue().GetSeverity(),
		// Confidence:     iss.GetRawIssue().GetConfidence(),

	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

// The Severity field is normally mapped into the jira 'Impact' field, so the assumption
// is that Severity = Impact; which in practice is generally true with small exceptions
func severtiyToText(severity v1.Severity) string {
	switch severity {
	case v1.Severity_SEVERITY_INFO:
		return "Info"
	case v1.Severity_SEVERITY_LOW:
		return "Minor / Localized"
	case v1.Severity_SEVERITY_MEDIUM:
		return "Moderate / Limited"
	case v1.Severity_SEVERITY_HIGH:
		return "Significant / Large"
	case v1.Severity_SEVERITY_CRITICAL:
		return "Extensive / Widespread"
	default:
		return "N/A"
	}
}

func confidenceToText(confidence v1.Confidence) string {
	switch confidence {
	case v1.Confidence_CONFIDENCE_INFO:
		return "Info"
	case v1.Confidence_CONFIDENCE_LOW:
		return "Low"
	case v1.Confidence_CONFIDENCE_MEDIUM:
		return "Medium"
	case v1.Confidence_CONFIDENCE_HIGH:
		return "High"
	case v1.Confidence_CONFIDENCE_CRITICAL:
		return "Critical"
	default:
		return "N/A"
	}
}

// TextToSeverity maps between dracon Severity and Jira severity textual fields
func TextToSeverity(severity string) v1.Severity {

	// The Severity field is normally mapped into the jira 'Impact' field, so the assumption
	// is that Severity = Impact; which in practice is generally true with small exceptions
	switch severity {
	case "Minor / Localized":
		return v1.Severity_SEVERITY_LOW
	case "Moderate / Limited":
		return v1.Severity_SEVERITY_MEDIUM
	case "Significant / Large":
		return v1.Severity_SEVERITY_HIGH
	case "Extensive / Widespread":
		return v1.Severity_SEVERITY_CRITICAL
	default:
		return v1.Severity_SEVERITY_INFO
	}
}

// TextToConfidence maps between dracon Confidence and a it's ext representation, used for adding the Confidence to Jira description.
func TextToConfidence(confidence string) v1.Confidence {
	switch confidence {

	case "Low":
		return v1.Confidence_CONFIDENCE_LOW
	case "Medium":
		return v1.Confidence_CONFIDENCE_MEDIUM
	case "High":
		return v1.Confidence_CONFIDENCE_HIGH
	case "Critical":
		return v1.Confidence_CONFIDENCE_CRITICAL
	default:
		return v1.Confidence_CONFIDENCE_INFO
	}
}
