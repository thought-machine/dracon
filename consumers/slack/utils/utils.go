package utils

import (
	"bytes"
	"consumers/slack/types"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	v1 "api/proto/v1"

	"github.com/golang/protobuf/ptypes"
)

func push(b string, webhook string) error {
	type SlackRequestBody struct {
		Text string `json:"text"`
	}
	var err error
	body, _ := json.Marshal(SlackRequestBody{Text: b})
	req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
func getRawIssue(scanStartTime time.Time, res *v1.LaunchToolResponse, iss *v1.Issue) ([]byte, error) {
	jBytes, err := json.Marshal(&types.FullDocument{
		ScanStartTime: scanStartTime,
		ScanID:        res.GetScanInfo().GetScanUuid(),
		ToolName:      res.GetToolName(),
		Source:        iss.GetSource(),
		Title:         iss.GetTitle(),
		Target:        iss.GetTarget(),
		Type:          iss.GetType(),
		Severity:      iss.GetSeverity(),
		CVSS:          iss.GetCvss(),
		Confidence:    iss.GetConfidence(),
		Description:   iss.GetDescription(),
		FirstFound:    scanStartTime,
		Count:         1,
		FalsePositive: false,
	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

func getEnrichedIssue(scanStartTime time.Time, res *v1.EnrichedLaunchToolResponse, iss *v1.EnrichedIssue) ([]byte, error) {
	firstSeenTime, _ := ptypes.Timestamp(iss.GetFirstSeen())
	jBytes, err := json.Marshal(&types.FullDocument{
		ScanStartTime: scanStartTime,
		ScanID:        res.GetOriginalResults().GetScanInfo().GetScanUuid(),
		ToolName:      res.GetOriginalResults().GetToolName(),
		Source:        iss.GetRawIssue().GetSource(),
		Title:         iss.GetRawIssue().GetTitle(),
		Target:        iss.GetRawIssue().GetTarget(),
		Type:          iss.GetRawIssue().GetType(),
		Severity:      iss.GetRawIssue().GetSeverity(),
		CVSS:          iss.GetRawIssue().GetCvss(),
		Confidence:    iss.GetRawIssue().GetConfidence(),
		Description:   iss.GetRawIssue().GetDescription(),
		FirstFound:    firstSeenTime,
		Count:         iss.GetCount(),
		FalsePositive: iss.GetFalsePositive(),
	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

// returns a list of stringified v1.LaunchToolResponse
func ProcessRawMessages(responses []*v1.LaunchToolResponse) ([]string, error) {
	messages := []string{}
	for _, res := range responses {
		scanStartTime, _ := ptypes.Timestamp(GetRawScanInfo(res).GetScanStartTime())
		for _, iss := range res.GetIssues() {
			b, err := getRawIssue(scanStartTime, res, iss)
			if err != nil {
				return nil, err
			}
			messages = append(messages, string(b))
		}
	}
	return messages, nil
}

// returns a list of stringified v1.EnrichedLaunchToolResponse
func ProcessEnrichedMessages(responses []*v1.EnrichedLaunchToolResponse) ([]string, error) {
	messages := []string{}
	for _, res := range responses {
		scanStartTime, _ := ptypes.Timestamp(GetEnrichedScanInfo(res).GetScanStartTime())
		for _, iss := range res.GetIssues() {
			b, err := getEnrichedIssue(scanStartTime, res, iss)
			if err != nil {
				return nil, err
			}
			messages = append(messages, string(b))
		}
	}
	return messages, nil
}
func GetRawScanInfo(response *v1.LaunchToolResponse) *v1.ScanInfo {
	return response.GetScanInfo()
}

func GetEnrichedScanInfo(response *v1.EnrichedLaunchToolResponse) *v1.ScanInfo {
	return response.GetOriginalResults().GetScanInfo()
}

func PushMetrics(scanUUID string, issuesNo int, scanStartTime time.Time, webhook string) {
	message := fmt.Sprintf("Dracon scan %s started on %s has been completed with %d issues\n", scanUUID, scanStartTime, issuesNo)
	push(message, webhook)
}
func PushMessage(b string, webhook string) {
	push(b, webhook)
}
func CountRawMessages(responses []*v1.LaunchToolResponse) int {
	result := 0
	for _, res := range responses {
		result += len(res.GetIssues())
	}
	return result
}

func CountEnrichedMessages(responses []*v1.EnrichedLaunchToolResponse) int {
	result := 0
	for _, res := range responses {
		result += len(res.GetIssues())
	}
	return result
}
