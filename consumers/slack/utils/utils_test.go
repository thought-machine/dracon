package utils

import (
	"bytes"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	v1 "github.com/thought-machine/dracon/api/proto/v1"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
)

//TODO tests: count* get*
func TestCountEnrichedMessages(t *testing.T) {
	eIssue := &v1.EnrichedIssue{
		RawIssue: &v1.Issue{
			Description: "this is a test description",
			Confidence:  v1.Confidence_CONFIDENCE_INFO,
			Severity:    v1.Severity_SEVERITY_INFO,
			Cvss:        0.0,
			Source:      "//foo/bar:baz",
			Target:      "//foo1/bar1:baz2",
			Title:       "Unit Test Title",
			Type:        "test type",
		}}
	expectedMessage := 2
	response := []*v1.EnrichedLaunchToolResponse{
		&v1.EnrichedLaunchToolResponse{
			OriginalResults: &v1.LaunchToolResponse{
				ToolName: "test",
				Issues:   []*v1.Issue{&v1.Issue{}},
				ScanInfo: &v1.ScanInfo{},
			},
			Issues: []*v1.EnrichedIssue{eIssue, eIssue},
		},
	}
	assert.Equal(t, expectedMessage, CountEnrichedMessages(response))

}
func TestCountRawMessages(t *testing.T) {
	eIssue := &v1.Issue{
		Description: "this is a test description",
		Confidence:  v1.Confidence_CONFIDENCE_INFO,
		Severity:    v1.Severity_SEVERITY_INFO,
		Cvss:        0.0,
		Source:      "//foo/bar:baz",
		Target:      "//foo1/bar1:baz2",
		Title:       "Unit Test Title",
		Type:        "test type",
	}
	expectedMessage := 3
	response := []*v1.LaunchToolResponse{&v1.LaunchToolResponse{Issues: []*v1.Issue{eIssue, eIssue, eIssue}, ScanInfo: &v1.ScanInfo{}}}
	assert.Equal(t, expectedMessage, CountRawMessages(response))
}
func TestProcessEnrichedMessages(t *testing.T) {
	tstamp, _ := time.Parse("2006-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
	startTime, _ := ptypes.TimestampProto(tstamp)
	tstamp, _ = time.Parse("2007-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
	firstSeen, _ := ptypes.TimestampProto(tstamp)
	tstamp, _ = time.Parse("2008-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
	updatedAt, _ := ptypes.TimestampProto(tstamp)

	expectedMessage := `{"scan_start_time":"0001-01-01T00:00:00Z","scan_id":"babbb83-4627-41c6-8ba0-70ee866290e9","tool_name":"test","source":"//foo/bar:baz","target":"//foo1/bar1:baz2","type":"test type","title":"Unit Test Title","severity":0,"cvss":0,"confidence":0,"description":"this is a test description","first_found":"0001-01-01T00:00:00Z","count":2,"false_positive":true}`
	response := []*v1.EnrichedLaunchToolResponse{
		&v1.EnrichedLaunchToolResponse{
			OriginalResults: &v1.LaunchToolResponse{
				ToolName: "test",
				Issues: []*v1.Issue{
					&v1.Issue{
						Description: "this is a test description",
						Confidence:  v1.Confidence_CONFIDENCE_INFO,
						Severity:    v1.Severity_SEVERITY_INFO,
						Cvss:        0.0,
						Source:      "//foo/bar:baz",
						Target:      "//foo1/bar1:baz2",
						Title:       "Unit Test Title",
						Type:        "test type",
					},
				},
				ScanInfo: &v1.ScanInfo{
					ScanUuid:      "babbb83-4627-41c6-8ba0-70ee866290e9",
					ScanStartTime: startTime,
				},
			},
			Issues: []*v1.EnrichedIssue{
				&v1.EnrichedIssue{
					FirstSeen:     firstSeen,
					UpdatedAt:     updatedAt,
					Hash:          "cf23df2207d99a74fbe169e3eba035e633b65d94",
					FalsePositive: true,
					Count:         2,
					RawIssue: &v1.Issue{
						Description: "this is a test description",
						Confidence:  v1.Confidence_CONFIDENCE_INFO,
						Severity:    v1.Severity_SEVERITY_INFO,
						Cvss:        0.0,
						Source:      "//foo/bar:baz",
						Target:      "//foo1/bar1:baz2",
						Title:       "Unit Test Title",
						Type:        "test type",
					},
				},
			},
		},
	}
	messages, err := ProcessEnrichedMessages(response)
	assert.Nil(t, err)
	fmt.Printf("%s\n", messages[0])
	assert.Equal(t, messages[0], expectedMessage)
}
func TestProcessRawMessages(t *testing.T) {
	tstamp, _ := time.Parse("2006-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
	startTime, _ := ptypes.TimestampProto(tstamp)
	expectedMessage := `{"scan_start_time":"0001-01-01T00:00:00Z","scan_id":"babbb83-4627-41c6-8ba0-70ee866290e9","tool_name":"test","source":"//foo/bar:baz","target":"//foo1/bar1:baz2","type":"test type","title":"Unit Test Title","severity":0,"cvss":0,"confidence":0,"description":"this is a test description","first_found":"0001-01-01T00:00:00Z","count":1,"false_positive":false}`

	response := []*v1.LaunchToolResponse{
		&v1.LaunchToolResponse{
			ToolName: "test",
			Issues: []*v1.Issue{
				&v1.Issue{
					Description: "this is a test description",
					Confidence:  v1.Confidence_CONFIDENCE_INFO,
					Severity:    v1.Severity_SEVERITY_INFO,
					Cvss:        0.0,
					Source:      "//foo/bar:baz",
					Target:      "//foo1/bar1:baz2",
					Title:       "Unit Test Title",
					Type:        "test type",
				},
			},
			ScanInfo: &v1.ScanInfo{
				ScanUuid:      "babbb83-4627-41c6-8ba0-70ee866290e9",
				ScanStartTime: startTime,
			},
		}}
	messages, err := ProcessRawMessages(response)
	assert.Nil(t, err)
	fmt.Printf("%s\n", messages[0])
	assert.Equal(t, messages[0], expectedMessage)

}
func TestPushMetrics(t *testing.T) {
	want := "OK"
	scanUUID := "test-uuid"
	scanStartTime, _ := time.Parse("2006-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
	issuesNo := 1234
	slackIn := `{"text":"Dracon scan test-uuid started on 0001-01-01 00:00:00 +0000 UTC has been completed with 1234 issues\n"}`
	slackStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		assert.Equal(t, buf.String(), slackIn)
		w.WriteHeader(200)
		w.Write([]byte(want))
	}))
	defer slackStub.Close()
	PushMetrics(scanUUID, issuesNo, scanStartTime, slackStub.URL)

}

func TestPush(t *testing.T) {
	testMessage := "test Message"
	want := "OK"
	slackIn := `{"text":"` + testMessage + `"}`
	slackStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		assert.Equal(t, buf.String(), slackIn)
		w.WriteHeader(200)
		w.Write([]byte(want))
	}))
	defer slackStub.Close()

	PushMessage(testMessage, slackStub.URL)

}
