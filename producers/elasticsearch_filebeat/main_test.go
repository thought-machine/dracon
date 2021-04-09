package main

import (
	"encoding/json"
	"testing"

	v1 "github.com/thought-machine/dracon/api/proto/v1"
	types "github.com/thought-machine/dracon/producers/elasticsearch_filebeat/types/elasticsearch-filebeat-issue"

	"github.com/stretchr/testify/assert"
)

func TestParseIssues(t *testing.T) {
	var results types.ElasticSearchFilebeatResult
	json.Unmarshal([]byte(exampleOutput), &results)

	issues := parseIssues(&results)

	expectedIssues := make([]*v1.Issue, 2)
	expectedIssues[0] = &v1.Issue{
		Target:      "foo-01234.example.com",
		Type:        "Antivirus Issue",
		Title:       "Antivirus Issue on foo-01234.example.com",
		Severity:    v1.Severity_SEVERITY_INFO,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: "error[013a0f06]: ESET Daemon: Error updating Antivirus modules: Server not found.",
	}
	expectedIssues[1] = &v1.Issue{
		Target:      "bar-56789.example.com",
		Type:        "Antivirus Issue",
		Title:       "Antivirus Issue on bar-56789.example.com",
		Severity:    v1.Severity_SEVERITY_INFO,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: "Tue Nov  24 06:52:07 2020 -> ERROR: Update failed.",
	}

	for i, issue := range issues {
		assert.Equal(t, issue.Target, expectedIssues[i].Target)
		assert.Equal(t, issue.Type, expectedIssues[i].Type)
		assert.Equal(t, issue.Title, expectedIssues[i].Title)
		assert.Equal(t, issue.Severity, expectedIssues[i].Severity)
		assert.Equal(t, issue.Cvss, expectedIssues[i].Cvss)
		assert.Equal(t, issue.Confidence, expectedIssues[i].Confidence)
		assert.Equal(t, issue.Description, expectedIssues[i].Description)
	}
}

var exampleOutput = `{
    "took": 187,
    "timed_out": false,
    "_shards": {
        "total": 45,
        "successful": 45,
        "skipped": 40,
        "failed": 0
    },
    "hits": {
        "total": 1,
        "max_score": null,
        "hits": [
            {
                "_index": "filebeat-7.6.0-2020.11.24",
                "_type": "_doc",
                "_id": "MjGs1VOXARhKg6Ac5tQM",
                "_version": 1,
                "_score": null,
                "_source": {
                "agent": {
                    "hostname": "foo-01234.example.com",
                    "id": "0fabad50-a690-4193-8714-370b523d2b04",
                    "type": "filebeat",
                    "ephemeral_id": "f9290b4a-3500-4a55-a870-0d5611f04e53",
                    "version": "7.6.0"
                },
                "process": {
                    "name": "esets",
                    "pid": 314
                },
                "log": {
                    "file": {
                        "path": "/var/log/system.log"
                    },
                    "offset": 186662
                },
                "fileset": {
                    "name": "syslog"
                },
                "message": "error[013a0f06]: ESET Daemon: Error updating Antivirus modules: Server not found.",
                "tags": [
                    "testENV",
                    "tls",
                    "v1.0",
                    "clientcert",
                    "beats_input_codec_plain_applied"
                ],
                "input": {
                    "type": "log"
                },
                "@timestamp": "2020-11-24T06:52:07.000Z",
                "system": {
                    "syslog": {}
                },
                "ecs": {
                    "version": "1.4.0"
                },
                "service": {
                    "type": "system"
                },
                "host": {
                    "hostname": "foo-01234",
                    "os": {
                        "build": "19H15",
                        "kernel": "19.6.0",
                        "name": "Mac OS X",
                        "family": "darwin",
                        "version": "10.15.7",
                        "platform": "darwin"
                    },
                    "name": "foo-01234.example.com",
                    "id": "5636DD9E-761B-41AB-8B40-BDDD8626988D",
                    "architecture": "x86_64"
                },
                "@version": "1",
                "event": {
                    "timezone": "+00:00",
                    "module": "system",
                    "dataset": "system.syslog"
                }
                },
                "fields": {
                    "@timestamp": [
                        "2020-11-24T06:52:07.000Z"
                    ]
                },
                "sort": [
                    1606200727000
                ]
            }
        ]
    },
    "aggregations": {
        "aggregation": {
            "doc_count": 5,
            "bg_count": 123,
            "buckets": [
                {
                    "metric": {
                    "hits": {
                        "total": 4,
                        "max_score": null,
                        "hits": [
                            {
                                "_index": "filebeat-7.6.0-2020.11.24",
                                "_type": "_doc",
                                "_id": "AxyDineddNSdSJTWCyDS",
                                "_score": null,
                                "_source": {
                                    "message": "Tue Nov  24 06:52:07 2020 -> ERROR: Update failed."
                                },
                                "sort": [
                                    1606200727000
                                ]
                            }
                        ]
                    },
                    "key": "bar-56789.example.com",
                    "doc_count": 4,
                    "score": 1092.342718382911,
                    "bg_count": 123
                }
            ]
        }
    }
}`
