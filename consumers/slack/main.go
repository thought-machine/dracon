package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/thought-machine/dracon/consumers"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
)

var (
	webhook     string
	longFormtat bool
)

func init() {
	flag.StringVar(&webhook, "webhook", "", "the webhook to push results to")
	flag.BoolVar(&longFormtat, "long", true, "post the full results to webhook, not just metrics")
}

func parseFlags() error {
	if err := consumers.ParseFlags(); err != nil {
		return err
	}
	if len(webhook) < 1 {
		return fmt.Errorf("webhook is undefined")
	}
	return nil
}

func countRawMessages(responses []*v1.LaunchToolResponse) int {
	result := 0
	for _, res := range responses {
		result += len(res.GetIssues())
	}
	return result
}

func countEnrichedMessages(responses []*v1.EnrichedLaunchToolResponse) int {
	result := 0
	for _, res := range responses {
		result += len(res.GetIssues())
	}
	return result
}
func pushMetrics(issuesNo int, scanStartTime time.Time) {
	message := fmt.Sprintf("Dracon Scan started on %s has been completed with %s issues\n", scanStartTime, issuesNo)
	push([]byte(message))
}

func main() {
	if err := consumers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	if consumers.Raw {
		responses, err := consumers.LoadToolResponse()
		if err != nil {
			log.Fatal(err)
		}
		if longFormtat == false {
			scanStartTime, _ := ptypes.Timestamp(responses[0].GetScanInfo().GetScanStartTime())
			pushMetrics(countRawMessages(responses), scanStartTime)
			return
		}
		for _, res := range responses {
			scanStartTime, _ := ptypes.Timestamp(res.GetScanInfo().GetScanStartTime())
			for _, iss := range res.GetIssues() {
				b, err := getRawIssue(scanStartTime, res, iss)
				if err != nil {
					log.Fatal(err)
				}
				push(b)
			}
		}
	} else {
		responses, err := consumers.LoadEnrichedToolResponse()
		if err != nil {
			log.Fatal(err)
		}
		if longFormtat == false {
			scanStartTime, _ := ptypes.Timestamp(responses[0].GetOriginalResults().GetScanInfo().GetScanStartTime())
			pushMetrics(countEnrichedMessages(responses), scanStartTime)
			return
		}
		for _, res := range responses {
			scanStartTime, _ := ptypes.Timestamp(res.GetOriginalResults().GetScanInfo().GetScanStartTime())
			for _, iss := range res.GetIssues() {
				b, err := getEnrichedIssue(scanStartTime, res, iss)
				if err != nil {
					log.Fatal(err)
				}
				push(b)
			}
		}
	}
}

func getRawIssue(scanStartTime time.Time, res *v1.LaunchToolResponse, iss *v1.Issue) ([]byte, error) {
	jBytes, err := json.Marshal(&fullDocument{
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
		FalsePositive: false,
	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

func getEnrichedIssue(scanStartTime time.Time, res *v1.EnrichedLaunchToolResponse, iss *v1.EnrichedIssue) ([]byte, error) {
	firstSeenTime, _ := ptypes.Timestamp(iss.GetFirstSeen())
	jBytes, err := json.Marshal(&fullDocument{
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
		FalsePositive: iss.GetFalsePositive(),
	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

type fullDocument struct {
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
	FalsePositive bool          `json:"false_positive"`
}

func push(b []byte) error {
	type SlackRequestBody struct {
		Text string `json:"text"`
	}
	var err error
	body, _ := json.Marshal(SlackRequestBody{Text: string(b)})
	req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
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
