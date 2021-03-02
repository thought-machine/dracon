package utils

import (
    "fmt"
    "testing"
    "time"

    "github.com/golang/protobuf/ptypes"
    "github.com/stretchr/testify/assert"

    v1 "github.com/thought-machine/dracon/api/proto/v1"
)

func TestProcessEnrichedMessages(t *testing.T) {
    tstamp, _ := time.Parse("2006-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
    startTime, _ := ptypes.TimestampProto(tstamp)
    tstamp, _ = time.Parse("2007-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
    firstSeen, _ := ptypes.TimestampProto(tstamp)
    tstamp, _ = time.Parse("2008-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
    updatedAt, _ := ptypes.TimestampProto(tstamp)

    expectedMessage := map[string]string{"scan_start_time": "0001-01-01T00:00:00Z", "scan_id": "babbb83-4627-41c6-8ba0-70ee866290e9", "tool_name": "test", "source": "//foo/bar:baz", "target": "//foo1/bar1:baz2", "type": "test type", "title": "Unit Test Title", "severity_text": "Info", "cvss": "0.000", "confidence_text": "Info", "description": "this is a test description", "first_found": "0001-01-01T00:00:00Z", "count": "2", "false_positive": "true", "hash":"cf23df2207d99a74fbe169e3eba035e633b65d94"}
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
    messages, _, err := ProcessEnrichedMessages(response, true, true, 0)
    assert.Nil(t, err)
    fmt.Printf("%s\n", messages[0])
    assert.Equal(t, messages[0], expectedMessage)
}
func TestProcessRawMessages(t *testing.T) {
    tstamp, _ := time.Parse("2006-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
    startTime, _ := ptypes.TimestampProto(tstamp)
    expectedMessage := map[string]string{"scan_start_time": "0001-01-01T00:00:00Z", "scan_id": "babbb83-4627-41c6-8ba0-70ee866290e9", "tool_name": "test", "source": "//foo/bar:baz", "target": "//foo1/bar1:baz2", "type": "test type", "title": "Unit Test Title", "severity_text": "Info", "cvss": "0.000", "confidence_text": "Info", "description": "this is a test description", "first_found": "0001-01-01T00:00:00Z", "count": "1", "false_positive": "false","hash":""}

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
        messages, _, err := ProcessRawMessages(response, 0)
        assert.Nil(t, err)
        fmt.Printf("%s\n", messages[0])
        assert.Equal(t, messages[0], expectedMessage)

    }
