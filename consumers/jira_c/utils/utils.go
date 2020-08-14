package utils

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/thought-machine/dracon/consumers"

	v1 "api/proto/v1"
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
	FalsePositive  string    `json:"false_positive"`
	// Severity   v1.Severity   `json:"severity"`
	// Confidence v1.Confidence `json:"confidence"`
}

func float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 3, 64)
}

func boolToString(b bool) string {
	return strconv.FormatBool(b)
}

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

func getRawIssue(scanStartTime time.Time, res *v1.LaunchToolResponse, iss *v1.Issue) ([]byte, error) {
	jBytes, err := json.Marshal(&Document{
		ScanStartTime:  scanStartTime,
		ScanID:         res.GetScanInfo().GetScanUuid(),
		ToolName:       res.GetToolName(),
		Source:         iss.GetSource(),
		Title:          iss.GetTitle(),
		Target:         iss.GetTarget(),
		Type:           iss.GetType(),
		SeverityText:   severtiyToText(iss.GetSeverity()),
		CVSS:           float64ToString(iss.GetCvss()),
		ConfidenceText: confidenceToText(iss.GetConfidence()),
		Description:    iss.GetDescription(),
		FirstFound:     scanStartTime,
		FalsePositive:  "false",
		// Severity:       iss.GetSeverity(),
		// Confidence:     iss.GetConfidence(),

	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

func getEnrichedIssue(scanStartTime time.Time, res *v1.EnrichedLaunchToolResponse, iss *v1.EnrichedIssue) ([]byte, error) {
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
		CVSS:           float64ToString(iss.GetRawIssue().GetCvss()),
		ConfidenceText: confidenceToText(iss.GetRawIssue().GetConfidence()),
		Description:    iss.GetRawIssue().GetDescription(),
		FirstFound:     firstSeenTime,
		FalsePositive:  boolToString(iss.GetFalsePositive()),
		// Severity:       iss.GetRawIssue().GetSeverity(),
		// Confidence:     iss.GetRawIssue().GetConfidence(),

	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

// ProcessMessages returns a list of stringified v1.LaunchToolResponse if consumers.Raw is true, or v1.EnrichedLaunchToolResponse otherwise
// This esentially avoids if/else statements, since the return type is the same in both scenarios
// :param responses: list of LaunchToolResponse protobufs
func ProcessMessages(allowDuplicates, allowFP bool, sevThreshold int) ([]string, int, error) {
	if consumers.Raw {
		log.Print("Parsing Raw results")
		responses, err := consumers.LoadToolResponse()
		if err != nil {
			log.Print("Could not load Raw tool response: ", err)
			return nil, 0, err
		}
		messages, discarded, err := ProcessRawMessages(responses, sevThreshold)
		if err != nil {
			log.Print("Could not Process Raw Messages: ", err)
			return nil, 0, err
		}
		return messages, discarded, nil
	} else {
		log.Print("Parsing Enriched results")
		responses, err := consumers.LoadEnrichedToolResponse()
		if err != nil {
			log.Print("Could not load Enriched tool response: ", err)
			return nil, 0, err
		}
		messages, discarded, err := ProcessEnrichedMessages(responses, allowDuplicates, allowFP, sevThreshold)
		if err != nil {
			log.Print("Could not Process Enriched messages: ", err)
			return nil, 0, err
		}
		return messages, discarded, nil
	}
}

// returns a list of stringified v1.LaunchToolResponse
func ProcessRawMessages(responses []*v1.LaunchToolResponse, sevThreshold int) ([]string, int, error) {
	messages := []string{}
	for _, res := range responses {
		scanStartTime, _ := ptypes.Timestamp(GetRawScanInfo(res).GetScanStartTime())
		for _, iss := range res.GetIssues() {
			// Discard issues that don't pass the severity threshold
			if iss.GetSeverity() < v1.Severity(sevThreshold) {
				continue
			}
			b, err := getRawIssue(scanStartTime, res, iss)
			if err != nil {
				return nil, 0, err
			}
			messages = append(messages, string(b))
		}
	}
	return messages, 0, nil
}

// returns a list of stringified v1.EnrichedLaunchToolResponse
func ProcessEnrichedMessages(responses []*v1.EnrichedLaunchToolResponse, allowDuplicate, allowFP bool, sevThreshold int) ([]string, int, error) {
	discardedMsgs := 0
	messages := []string{}
	for _, res := range responses {
		scanStartTime, _ := ptypes.Timestamp(GetEnrichedScanInfo(res).GetScanStartTime())
		for _, iss := range res.GetIssues() {
			// Discard issues that don't pass the severity threshold
			if iss.GetRawIssue().GetSeverity() < v1.Severity(sevThreshold) {
				continue
				// Discard issues that are duplicates or false positives, according to the policy
			} else if (!allowDuplicate && iss.GetCount() > 1) || (!allowFP && iss.GetFalsePositive()) {
				discardedMsgs++
				continue
			}
			b, err := getEnrichedIssue(scanStartTime, res, iss)
			if err != nil {
				return nil, 0, err
			}
			messages = append(messages, string(b))
		}
	}
	return messages, discardedMsgs, nil
}

func GetRawScanInfo(response *v1.LaunchToolResponse) *v1.ScanInfo {
	return response.GetScanInfo()
}

func GetEnrichedScanInfo(response *v1.EnrichedLaunchToolResponse) *v1.ScanInfo {
	return response.GetOriginalResults().GetScanInfo()
}
